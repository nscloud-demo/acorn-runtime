//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	internal_acorn_iov1 "github.com/acorn-io/runtime/pkg/apis/internal.acorn.io/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaseResources) DeepCopyInto(out *BaseResources) {
	*out = *in
	out.VolumeStorage = in.VolumeStorage.DeepCopy()
	out.Memory = in.Memory.DeepCopy()
	out.CPU = in.CPU.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseResources.
func (in *BaseResources) DeepCopy() *BaseResources {
	if in == nil {
		return nil
	}
	out := new(BaseResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterComputeClassInstance) DeepCopyInto(out *ClusterComputeClassInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Memory.DeepCopyInto(&out.Memory)
	if in.SupportedRegions != nil {
		in, out := &in.SupportedRegions, &out.SupportedRegions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterComputeClassInstance.
func (in *ClusterComputeClassInstance) DeepCopy() *ClusterComputeClassInstance {
	if in == nil {
		return nil
	}
	out := new(ClusterComputeClassInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterComputeClassInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterComputeClassInstanceList) DeepCopyInto(out *ClusterComputeClassInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterComputeClassInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterComputeClassInstanceList.
func (in *ClusterComputeClassInstanceList) DeepCopy() *ClusterComputeClassInstanceList {
	if in == nil {
		return nil
	}
	out := new(ClusterComputeClassInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterComputeClassInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterImageRoleAuthorizationInstance) DeepCopyInto(out *ClusterImageRoleAuthorizationInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.ImageSelector.DeepCopyInto(&out.ImageSelector)
	in.Roles.DeepCopyInto(&out.Roles)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterImageRoleAuthorizationInstance.
func (in *ClusterImageRoleAuthorizationInstance) DeepCopy() *ClusterImageRoleAuthorizationInstance {
	if in == nil {
		return nil
	}
	out := new(ClusterImageRoleAuthorizationInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterImageRoleAuthorizationInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterImageRoleAuthorizationInstanceList) DeepCopyInto(out *ClusterImageRoleAuthorizationInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterImageRoleAuthorizationInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterImageRoleAuthorizationInstanceList.
func (in *ClusterImageRoleAuthorizationInstanceList) DeepCopy() *ClusterImageRoleAuthorizationInstanceList {
	if in == nil {
		return nil
	}
	out := new(ClusterImageRoleAuthorizationInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterImageRoleAuthorizationInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVolumeClassInstance) DeepCopyInto(out *ClusterVolumeClassInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.AllowedAccessModes != nil {
		in, out := &in.AllowedAccessModes, &out.AllowedAccessModes
		*out = make(internal_acorn_iov1.AccessModes, len(*in))
		copy(*out, *in)
	}
	out.Size = in.Size
	if in.SupportedRegions != nil {
		in, out := &in.SupportedRegions, &out.SupportedRegions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVolumeClassInstance.
func (in *ClusterVolumeClassInstance) DeepCopy() *ClusterVolumeClassInstance {
	if in == nil {
		return nil
	}
	out := new(ClusterVolumeClassInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterVolumeClassInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterVolumeClassInstanceList) DeepCopyInto(out *ClusterVolumeClassInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterVolumeClassInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterVolumeClassInstanceList.
func (in *ClusterVolumeClassInstanceList) DeepCopy() *ClusterVolumeClassInstanceList {
	if in == nil {
		return nil
	}
	out := new(ClusterVolumeClassInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterVolumeClassInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeClassMemory) DeepCopyInto(out *ComputeClassMemory) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeClassMemory.
func (in *ComputeClassMemory) DeepCopy() *ComputeClassMemory {
	if in == nil {
		return nil
	}
	out := new(ComputeClassMemory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRoleAuthorizationInstance) DeepCopyInto(out *ImageRoleAuthorizationInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.ImageSelector.DeepCopyInto(&out.ImageSelector)
	in.Roles.DeepCopyInto(&out.Roles)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRoleAuthorizationInstance.
func (in *ImageRoleAuthorizationInstance) DeepCopy() *ImageRoleAuthorizationInstance {
	if in == nil {
		return nil
	}
	out := new(ImageRoleAuthorizationInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageRoleAuthorizationInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageRoleAuthorizationInstanceList) DeepCopyInto(out *ImageRoleAuthorizationInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImageRoleAuthorizationInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageRoleAuthorizationInstanceList.
func (in *ImageRoleAuthorizationInstanceList) DeepCopy() *ImageRoleAuthorizationInstanceList {
	if in == nil {
		return nil
	}
	out := new(ImageRoleAuthorizationInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageRoleAuthorizationInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectComputeClassInstance) DeepCopyInto(out *ProjectComputeClassInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Memory.DeepCopyInto(&out.Memory)
	if in.SupportedRegions != nil {
		in, out := &in.SupportedRegions, &out.SupportedRegions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectComputeClassInstance.
func (in *ProjectComputeClassInstance) DeepCopy() *ProjectComputeClassInstance {
	if in == nil {
		return nil
	}
	out := new(ProjectComputeClassInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectComputeClassInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectComputeClassInstanceList) DeepCopyInto(out *ProjectComputeClassInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProjectComputeClassInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectComputeClassInstanceList.
func (in *ProjectComputeClassInstanceList) DeepCopy() *ProjectComputeClassInstanceList {
	if in == nil {
		return nil
	}
	out := new(ProjectComputeClassInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectComputeClassInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectVolumeClassInstance) DeepCopyInto(out *ProjectVolumeClassInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.AllowedAccessModes != nil {
		in, out := &in.AllowedAccessModes, &out.AllowedAccessModes
		*out = make(internal_acorn_iov1.AccessModes, len(*in))
		copy(*out, *in)
	}
	out.Size = in.Size
	if in.SupportedRegions != nil {
		in, out := &in.SupportedRegions, &out.SupportedRegions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectVolumeClassInstance.
func (in *ProjectVolumeClassInstance) DeepCopy() *ProjectVolumeClassInstance {
	if in == nil {
		return nil
	}
	out := new(ProjectVolumeClassInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectVolumeClassInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectVolumeClassInstanceList) DeepCopyInto(out *ProjectVolumeClassInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProjectVolumeClassInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectVolumeClassInstanceList.
func (in *ProjectVolumeClassInstanceList) DeepCopy() *ProjectVolumeClassInstanceList {
	if in == nil {
		return nil
	}
	out := new(ProjectVolumeClassInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectVolumeClassInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaRequestInstance) DeepCopyInto(out *QuotaRequestInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaRequestInstance.
func (in *QuotaRequestInstance) DeepCopy() *QuotaRequestInstance {
	if in == nil {
		return nil
	}
	out := new(QuotaRequestInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QuotaRequestInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaRequestInstanceList) DeepCopyInto(out *QuotaRequestInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]QuotaRequestInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaRequestInstanceList.
func (in *QuotaRequestInstanceList) DeepCopy() *QuotaRequestInstanceList {
	if in == nil {
		return nil
	}
	out := new(QuotaRequestInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QuotaRequestInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaRequestInstanceSpec) DeepCopyInto(out *QuotaRequestInstanceSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaRequestInstanceSpec.
func (in *QuotaRequestInstanceSpec) DeepCopy() *QuotaRequestInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(QuotaRequestInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaRequestInstanceStatus) DeepCopyInto(out *QuotaRequestInstanceStatus) {
	*out = *in
	in.AllocatedResources.DeepCopyInto(&out.AllocatedResources)
	if in.FailedResources != nil {
		in, out := &in.FailedResources, &out.FailedResources
		*out = new(QuotaRequestResources)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]internal_acorn_iov1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaRequestInstanceStatus.
func (in *QuotaRequestInstanceStatus) DeepCopy() *QuotaRequestInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(QuotaRequestInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaRequestResources) DeepCopyInto(out *QuotaRequestResources) {
	*out = *in
	in.BaseResources.DeepCopyInto(&out.BaseResources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaRequestResources.
func (in *QuotaRequestResources) DeepCopy() *QuotaRequestResources {
	if in == nil {
		return nil
	}
	out := new(QuotaRequestResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleAuthorizations) DeepCopyInto(out *RoleAuthorizations) {
	*out = *in
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RoleRefs != nil {
		in, out := &in.RoleRefs, &out.RoleRefs
		*out = make([]RoleRef, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleAuthorizations.
func (in *RoleAuthorizations) DeepCopy() *RoleAuthorizations {
	if in == nil {
		return nil
	}
	out := new(RoleAuthorizations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RoleRef) DeepCopyInto(out *RoleRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RoleRef.
func (in *RoleRef) DeepCopy() *RoleRef {
	if in == nil {
		return nil
	}
	out := new(RoleRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeClassSize) DeepCopyInto(out *VolumeClassSize) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeClassSize.
func (in *VolumeClassSize) DeepCopy() *VolumeClassSize {
	if in == nil {
		return nil
	}
	out := new(VolumeClassSize)
	in.DeepCopyInto(out)
	return out
}
