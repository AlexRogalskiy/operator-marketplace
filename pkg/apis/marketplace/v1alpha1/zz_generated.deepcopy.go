// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSourceConfig) DeepCopyInto(out *CatalogSourceConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSourceConfig.
func (in *CatalogSourceConfig) DeepCopy() *CatalogSourceConfig {
	if in == nil {
		return nil
	}
	out := new(CatalogSourceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CatalogSourceConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSourceConfigList) DeepCopyInto(out *CatalogSourceConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CatalogSourceConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSourceConfigList.
func (in *CatalogSourceConfigList) DeepCopy() *CatalogSourceConfigList {
	if in == nil {
		return nil
	}
	out := new(CatalogSourceConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CatalogSourceConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSourceConfigSpec) DeepCopyInto(out *CatalogSourceConfigSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSourceConfigSpec.
func (in *CatalogSourceConfigSpec) DeepCopy() *CatalogSourceConfigSpec {
	if in == nil {
		return nil
	}
	out := new(CatalogSourceConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogSourceConfigStatus) DeepCopyInto(out *CatalogSourceConfigStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogSourceConfigStatus.
func (in *CatalogSourceConfigStatus) DeepCopy() *CatalogSourceConfigStatus {
	if in == nil {
		return nil
	}
	out := new(CatalogSourceConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorSource) DeepCopyInto(out *OperatorSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorSource.
func (in *OperatorSource) DeepCopy() *OperatorSource {
	if in == nil {
		return nil
	}
	out := new(OperatorSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OperatorSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorSourceList) DeepCopyInto(out *OperatorSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OperatorSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorSourceList.
func (in *OperatorSourceList) DeepCopy() *OperatorSourceList {
	if in == nil {
		return nil
	}
	out := new(OperatorSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OperatorSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorSourceSpec) DeepCopyInto(out *OperatorSourceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorSourceSpec.
func (in *OperatorSourceSpec) DeepCopy() *OperatorSourceSpec {
	if in == nil {
		return nil
	}
	out := new(OperatorSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperatorSourceStatus) DeepCopyInto(out *OperatorSourceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperatorSourceStatus.
func (in *OperatorSourceStatus) DeepCopy() *OperatorSourceStatus {
	if in == nil {
		return nil
	}
	out := new(OperatorSourceStatus)
	in.DeepCopyInto(out)
	return out
}
