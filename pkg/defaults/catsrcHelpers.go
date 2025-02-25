package defaults

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"

	olmv1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
	wrapper "github.com/operator-framework/operator-marketplace/pkg/client"
	"github.com/sirupsen/logrus"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func ensureCatsrc(
	ctx context.Context,
	client wrapper.Client,
	config map[string]bool,
	catsrc olmv1alpha1.CatalogSource,
) error {
	disable, present := config[catsrc.Name]
	if !present {
		disable = false
	}

	err := processCatsrc(ctx, client, catsrc, disable)
	if err != nil {
		return err
	}

	return nil
}

// getCatsrcDefinition returns a CatalogSource definition from the given file
// in the defaults directory. It only supports decoding CatalogSources. Any
// other resource type will result in an error.
func getCatsrcDefinition(fileName string) (*olmv1alpha1.CatalogSource, error) {
	file, err := os.Open(filepath.Join(Dir, fileName))
	if err != nil {
		return nil, err
	}

	catsrc := &olmv1alpha1.CatalogSource{}
	decoder := yaml.NewYAMLOrJSONDecoder(file, 1024)
	err = decoder.Decode(catsrc)
	if err != nil {
		return nil, err
	}
	if strings.Compare(catsrc.Kind, "CatalogSource") != 0 {
		return nil, errors.New("Not an CatalogSource")
	}
	return catsrc, nil
}

// processCatsrc will ensure that the given CatalogSource is present or not on
// the cluster based on the disable flag.
func processCatsrc(ctx context.Context, client wrapper.Client, def olmv1alpha1.CatalogSource, disable bool) error {
	// Get CatalogSource on the cluster
	cluster := &olmv1alpha1.CatalogSource{}
	err := client.Get(ctx, wrapper.ObjectKey{
		Name:      def.Name,
		Namespace: def.Namespace,
	}, cluster)
	if err != nil && !k8sErrors.IsNotFound(err) {
		logrus.Errorf("[defaults] Error getting CatalogSource %s - %v", def.Name, err)
		return err
	}

	if disable {
		if cluster.Annotations[defaultCatsrcAnnotationKey] == defaultCatsrcAnnotationValue {
			err = ensureCatsrcAbsent(ctx, client, def, cluster)
		}
	} else {
		err = ensureCatsrcPresent(ctx, client, def, cluster)
	}

	if err != nil {
		logrus.Errorf("[defaults] Error processing CatalogSource %s - %v", def.Name, err)
	}

	return err
}

// ensureCatsrcAbsent ensure that that the default CatalogSource is not present on the cluster
func ensureCatsrcAbsent(
	ctx context.Context,
	client wrapper.Client,
	def olmv1alpha1.CatalogSource,
	cluster *olmv1alpha1.CatalogSource,
) error {
	// CatalogSource is not present on the cluster or has been marked for deletion
	if cluster.Name == "" || !cluster.ObjectMeta.DeletionTimestamp.IsZero() {
		logrus.Infof("[defaults] CatalogSource %s not present or has been marked for deletion", def.Name)
		return nil
	}

	if err := client.Delete(ctx, cluster); err != nil {
		return err
	}
	logrus.Infof("[defaults] Deleting CatalogSource %s", def.Name)

	return nil
}

// ensureCatsrcPresent ensure that that the default CatalogSource is present on the cluster
func ensureCatsrcPresent(
	ctx context.Context,
	client wrapper.Client,
	def olmv1alpha1.CatalogSource,
	cluster *olmv1alpha1.CatalogSource,
) error {
	def = *def.DeepCopy()
	if def.Annotations == nil {
		def.Annotations = make(map[string]string)
	}
	def.Annotations[defaultCatsrcAnnotationKey] = defaultCatsrcAnnotationValue

	// Create if not present or is deleted
	if cluster.Name == "" || (!cluster.ObjectMeta.DeletionTimestamp.IsZero() && len(cluster.Finalizers) == 0) {
		err := client.Create(ctx, &def)
		if err != nil {
			return err
		}
		logrus.Infof("[defaults] Creating CatalogSource %s", def.Name)
		return nil
	}

	if cluster.Annotations[defaultCatsrcAnnotationKey] == defaultCatsrcAnnotationValue && AreCatsrcSpecsEqual(&def.Spec, &cluster.Spec) {
		logrus.Infof("[defaults] CatalogSource %s is annotated and its spec is the same as the default spec", def.Name)
		return nil
	}

	// Update if the spec has changed
	cluster.Spec = def.Spec
	if cluster.Annotations == nil {
		cluster.Annotations = make(map[string]string)
	}
	cluster.Annotations[defaultCatsrcAnnotationKey] = defaultCatsrcAnnotationValue
	err := client.Update(ctx, cluster)
	if err != nil {
		return err
	}

	logrus.Infof("[defaults] Restoring CatalogSource %s", def.Name)

	return nil
}

// AreCatsrcSpecsEqual returns true if the Specs it receives are the same.
// Otherwise, the function returns false.
//
// The function performs a case insensitive comparison of corresponding
// attributes.
//
// If either of the Specs received is nil, then the function returns false.
func AreCatsrcSpecsEqual(spec1 *olmv1alpha1.CatalogSourceSpec, spec2 *olmv1alpha1.CatalogSourceSpec) bool {
	if spec1 == nil || spec2 == nil {
		return false
	}
	if !strings.EqualFold(string(spec1.SourceType), string(spec2.SourceType)) ||
		!strings.EqualFold(spec1.ConfigMap, spec2.ConfigMap) ||
		!strings.EqualFold(spec1.Address, spec2.Address) ||
		!strings.EqualFold(spec1.DisplayName, spec2.DisplayName) ||
		!strings.EqualFold(spec1.Publisher, spec2.Publisher) ||
		!strings.EqualFold(spec1.Image, spec2.Image) {
		return false
	}
	if spec1.UpdateStrategy != nil && spec2.UpdateStrategy == nil {
		return false
	}
	if spec1.UpdateStrategy == nil && spec2.UpdateStrategy != nil {
		return false
	}
	if spec1.UpdateStrategy != nil && spec2.UpdateStrategy != nil {
		if spec1.UpdateStrategy.RegistryPoll != spec1.UpdateStrategy.RegistryPoll {
			return false
		}
	}
	return true
}
