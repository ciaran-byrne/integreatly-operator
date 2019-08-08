// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1beta2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressPlan) DeepCopyInto(out *AddressPlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressPlan.
func (in *AddressPlan) DeepCopy() *AddressPlan {
	if in == nil {
		return nil
	}
	out := new(AddressPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddressPlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressPlanList) DeepCopyInto(out *AddressPlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AddressPlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressPlanList.
func (in *AddressPlanList) DeepCopy() *AddressPlanList {
	if in == nil {
		return nil
	}
	out := new(AddressPlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddressPlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressPlanResources) DeepCopyInto(out *AddressPlanResources) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressPlanResources.
func (in *AddressPlanResources) DeepCopy() *AddressPlanResources {
	if in == nil {
		return nil
	}
	out := new(AddressPlanResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressPlanSpec) DeepCopyInto(out *AddressPlanSpec) {
	*out = *in
	out.Resources = in.Resources
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressPlanSpec.
func (in *AddressPlanSpec) DeepCopy() *AddressPlanSpec {
	if in == nil {
		return nil
	}
	out := new(AddressPlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressPlanStatus) DeepCopyInto(out *AddressPlanStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressPlanStatus.
func (in *AddressPlanStatus) DeepCopy() *AddressPlanStatus {
	if in == nil {
		return nil
	}
	out := new(AddressPlanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressSpacePlan) DeepCopyInto(out *AddressSpacePlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressSpacePlan.
func (in *AddressSpacePlan) DeepCopy() *AddressSpacePlan {
	if in == nil {
		return nil
	}
	out := new(AddressSpacePlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddressSpacePlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressSpacePlanList) DeepCopyInto(out *AddressSpacePlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AddressSpacePlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressSpacePlanList.
func (in *AddressSpacePlanList) DeepCopy() *AddressSpacePlanList {
	if in == nil {
		return nil
	}
	out := new(AddressSpacePlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddressSpacePlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressSpacePlanResourceLimits) DeepCopyInto(out *AddressSpacePlanResourceLimits) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressSpacePlanResourceLimits.
func (in *AddressSpacePlanResourceLimits) DeepCopy() *AddressSpacePlanResourceLimits {
	if in == nil {
		return nil
	}
	out := new(AddressSpacePlanResourceLimits)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressSpacePlanSpec) DeepCopyInto(out *AddressSpacePlanSpec) {
	*out = *in
	out.ResourceLimits = in.ResourceLimits
	if in.AddressPlans != nil {
		in, out := &in.AddressPlans, &out.AddressPlans
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressSpacePlanSpec.
func (in *AddressSpacePlanSpec) DeepCopy() *AddressSpacePlanSpec {
	if in == nil {
		return nil
	}
	out := new(AddressSpacePlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddressSpacePlanStatus) DeepCopyInto(out *AddressSpacePlanStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddressSpacePlanStatus.
func (in *AddressSpacePlanStatus) DeepCopy() *AddressSpacePlanStatus {
	if in == nil {
		return nil
	}
	out := new(AddressSpacePlanStatus)
	in.DeepCopyInto(out)
	return out
}