// +build !ignore_autogenerated

/*
Copyright 2019 Agoda DevOps Container.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivePromotion) DeepCopyInto(out *ActivePromotion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivePromotion.
func (in *ActivePromotion) DeepCopy() *ActivePromotion {
	if in == nil {
		return nil
	}
	out := new(ActivePromotion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActivePromotion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivePromotionCondition) DeepCopyInto(out *ActivePromotionCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivePromotionCondition.
func (in *ActivePromotionCondition) DeepCopy() *ActivePromotionCondition {
	if in == nil {
		return nil
	}
	out := new(ActivePromotionCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivePromotionHistory) DeepCopyInto(out *ActivePromotionHistory) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivePromotionHistory.
func (in *ActivePromotionHistory) DeepCopy() *ActivePromotionHistory {
	if in == nil {
		return nil
	}
	out := new(ActivePromotionHistory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActivePromotionHistory) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivePromotionHistoryDeployment) DeepCopyInto(out *ActivePromotionHistoryDeployment) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivePromotionHistoryDeployment.
func (in *ActivePromotionHistoryDeployment) DeepCopy() *ActivePromotionHistoryDeployment {
	if in == nil {
		return nil
	}
	out := new(ActivePromotionHistoryDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivePromotionHistoryK8SResources) DeepCopyInto(out *ActivePromotionHistoryK8SResources) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]ActivePromotionHistoryPod, len(*in))
		copy(*out, *in)
	}
	if in.Deployments != nil {
		in, out := &in.Deployments, &out.Deployments
		*out = make([]ActivePromotionHistoryDeployment, len(*in))
		copy(*out, *in)
	}
	if in.StatefulSets != nil {
		in, out := &in.StatefulSets, &out.StatefulSets
		*out = make([]ActivePromotionHistoryStatefulSet, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivePromotionHistoryK8SResources.
func (in *ActivePromotionHistoryK8SResources) DeepCopy() *ActivePromotionHistoryK8SResources {
	if in == nil {
		return nil
	}
	out := new(ActivePromotionHistoryK8SResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivePromotionHistoryList) DeepCopyInto(out *ActivePromotionHistoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ActivePromotionHistory, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivePromotionHistoryList.
func (in *ActivePromotionHistoryList) DeepCopy() *ActivePromotionHistoryList {
	if in == nil {
		return nil
	}
	out := new(ActivePromotionHistoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActivePromotionHistoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivePromotionHistoryPod) DeepCopyInto(out *ActivePromotionHistoryPod) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivePromotionHistoryPod.
func (in *ActivePromotionHistoryPod) DeepCopy() *ActivePromotionHistoryPod {
	if in == nil {
		return nil
	}
	out := new(ActivePromotionHistoryPod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivePromotionHistorySpec) DeepCopyInto(out *ActivePromotionHistorySpec) {
	*out = *in
	if in.ActivePromotion != nil {
		in, out := &in.ActivePromotion, &out.ActivePromotion
		*out = new(ActivePromotion)
		(*in).DeepCopyInto(*out)
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivePromotionHistorySpec.
func (in *ActivePromotionHistorySpec) DeepCopy() *ActivePromotionHistorySpec {
	if in == nil {
		return nil
	}
	out := new(ActivePromotionHistorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivePromotionHistoryStatefulSet) DeepCopyInto(out *ActivePromotionHistoryStatefulSet) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivePromotionHistoryStatefulSet.
func (in *ActivePromotionHistoryStatefulSet) DeepCopy() *ActivePromotionHistoryStatefulSet {
	if in == nil {
		return nil
	}
	out := new(ActivePromotionHistoryStatefulSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivePromotionHistoryStatus) DeepCopyInto(out *ActivePromotionHistoryStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivePromotionHistoryStatus.
func (in *ActivePromotionHistoryStatus) DeepCopy() *ActivePromotionHistoryStatus {
	if in == nil {
		return nil
	}
	out := new(ActivePromotionHistoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivePromotionList) DeepCopyInto(out *ActivePromotionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ActivePromotion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivePromotionList.
func (in *ActivePromotionList) DeepCopy() *ActivePromotionList {
	if in == nil {
		return nil
	}
	out := new(ActivePromotionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActivePromotionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivePromotionSpec) DeepCopyInto(out *ActivePromotionSpec) {
	*out = *in
	if in.TearDownDuration != nil {
		in, out := &in.TearDownDuration, &out.TearDownDuration
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivePromotionSpec.
func (in *ActivePromotionSpec) DeepCopy() *ActivePromotionSpec {
	if in == nil {
		return nil
	}
	out := new(ActivePromotionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivePromotionStatus) DeepCopyInto(out *ActivePromotionStatus) {
	*out = *in
	if in.StartedAt != nil {
		in, out := &in.StartedAt, &out.StartedAt
		*out = (*in).DeepCopy()
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = (*in).DeepCopy()
	}
	if in.DestroyTime != nil {
		in, out := &in.DestroyTime, &out.DestroyTime
		*out = (*in).DeepCopy()
	}
	if in.OutdatedComponents != nil {
		in, out := &in.OutdatedComponents, &out.OutdatedComponents
		*out = make([]*OutdatedComponent, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(OutdatedComponent)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	in.PreActiveQueue.DeepCopyInto(&out.PreActiveQueue)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ActivePromotionCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivePromotionStatus.
func (in *ActivePromotionStatus) DeepCopy() *ActivePromotionStatus {
	if in == nil {
		return nil
	}
	out := new(ActivePromotionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Credential) DeepCopyInto(out *Credential) {
	*out = *in
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(UsernamePasswordCredential)
		(*in).DeepCopyInto(*out)
	}
	if in.Teamcity != nil {
		in, out := &in.Teamcity, &out.Teamcity
		*out = new(UsernamePasswordCredential)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Credential.
func (in *Credential) DeepCopy() *Credential {
	if in == nil {
		return nil
	}
	out := new(Credential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesiredComponent) DeepCopyInto(out *DesiredComponent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesiredComponent.
func (in *DesiredComponent) DeepCopy() *DesiredComponent {
	if in == nil {
		return nil
	}
	out := new(DesiredComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DesiredComponent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesiredComponentList) DeepCopyInto(out *DesiredComponentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DesiredComponent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesiredComponentList.
func (in *DesiredComponentList) DeepCopy() *DesiredComponentList {
	if in == nil {
		return nil
	}
	out := new(DesiredComponentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DesiredComponentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesiredComponentSpec) DeepCopyInto(out *DesiredComponentSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesiredComponentSpec.
func (in *DesiredComponentSpec) DeepCopy() *DesiredComponentSpec {
	if in == nil {
		return nil
	}
	out := new(DesiredComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesiredComponentStatus) DeepCopyInto(out *DesiredComponentStatus) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesiredComponentStatus.
func (in *DesiredComponentStatus) DeepCopy() *DesiredComponentStatus {
	if in == nil {
		return nil
	}
	out := new(DesiredComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DesiredImageTime) DeepCopyInto(out *DesiredImageTime) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(Image)
		**out = **in
	}
	in.CreatedTime.DeepCopyInto(&out.CreatedTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesiredImageTime.
func (in *DesiredImageTime) DeepCopy() *DesiredImageTime {
	if in == nil {
		return nil
	}
	out := new(DesiredImageTime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitStorage) DeepCopyInto(out *GitStorage) {
	*out = *in
	if in.CloneTimeout != nil {
		in, out := &in.CloneTimeout, &out.CloneTimeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.PullTimeout != nil {
		in, out := &in.PullTimeout, &out.PullTimeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.PushTimeout != nil {
		in, out := &in.PushTimeout, &out.PushTimeout
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitStorage.
func (in *GitStorage) DeepCopy() *GitStorage {
	if in == nil {
		return nil
	}
	out := new(GitStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutdatedComponent) DeepCopyInto(out *OutdatedComponent) {
	*out = *in
	if in.CurrentImage != nil {
		in, out := &in.CurrentImage, &out.CurrentImage
		*out = new(Image)
		**out = **in
	}
	if in.LatestImage != nil {
		in, out := &in.LatestImage, &out.LatestImage
		*out = new(Image)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutdatedComponent.
func (in *OutdatedComponent) DeepCopy() *OutdatedComponent {
	if in == nil {
		return nil
	}
	out := new(OutdatedComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Queue) DeepCopyInto(out *Queue) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Queue.
func (in *Queue) DeepCopy() *Queue {
	if in == nil {
		return nil
	}
	out := new(Queue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Queue) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueCondition) DeepCopyInto(out *QueueCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueCondition.
func (in *QueueCondition) DeepCopy() *QueueCondition {
	if in == nil {
		return nil
	}
	out := new(QueueCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueHistory) DeepCopyInto(out *QueueHistory) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueHistory.
func (in *QueueHistory) DeepCopy() *QueueHistory {
	if in == nil {
		return nil
	}
	out := new(QueueHistory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QueueHistory) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueHistoryList) DeepCopyInto(out *QueueHistoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]QueueHistory, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueHistoryList.
func (in *QueueHistoryList) DeepCopy() *QueueHistoryList {
	if in == nil {
		return nil
	}
	out := new(QueueHistoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QueueHistoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueHistorySpec) DeepCopyInto(out *QueueHistorySpec) {
	*out = *in
	if in.Queue != nil {
		in, out := &in.Queue, &out.Queue
		*out = new(Queue)
		(*in).DeepCopyInto(*out)
	}
	in.AppliedValues.DeepCopyInto(&out.AppliedValues)
	if in.StableComponents != nil {
		in, out := &in.StableComponents, &out.StableComponents
		*out = make([]StableComponent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueHistorySpec.
func (in *QueueHistorySpec) DeepCopy() *QueueHistorySpec {
	if in == nil {
		return nil
	}
	out := new(QueueHistorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueHistoryStatus) DeepCopyInto(out *QueueHistoryStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueHistoryStatus.
func (in *QueueHistoryStatus) DeepCopy() *QueueHistoryStatus {
	if in == nil {
		return nil
	}
	out := new(QueueHistoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueList) DeepCopyInto(out *QueueList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Queue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueList.
func (in *QueueList) DeepCopy() *QueueList {
	if in == nil {
		return nil
	}
	out := new(QueueList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QueueList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueSpec) DeepCopyInto(out *QueueSpec) {
	*out = *in
	if in.NextProcessAt != nil {
		in, out := &in.NextProcessAt, &out.NextProcessAt
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueSpec.
func (in *QueueSpec) DeepCopy() *QueueSpec {
	if in == nil {
		return nil
	}
	out := new(QueueSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QueueStatus) DeepCopyInto(out *QueueStatus) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = (*in).DeepCopy()
	}
	if in.NextProcessAt != nil {
		in, out := &in.NextProcessAt, &out.NextProcessAt
		*out = (*in).DeepCopy()
	}
	if in.StartDeployTime != nil {
		in, out := &in.StartDeployTime, &out.StartDeployTime
		*out = (*in).DeepCopy()
	}
	if in.StartTestingTime != nil {
		in, out := &in.StartTestingTime, &out.StartTestingTime
		*out = (*in).DeepCopy()
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]QueueCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.TestRunner = in.TestRunner
	if in.ImageMissingList != nil {
		in, out := &in.ImageMissingList, &out.ImageMissingList
		*out = make([]Image, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QueueStatus.
func (in *QueueStatus) DeepCopy() *QueueStatus {
	if in == nil {
		return nil
	}
	out := new(QueueStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StableComponent) DeepCopyInto(out *StableComponent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StableComponent.
func (in *StableComponent) DeepCopy() *StableComponent {
	if in == nil {
		return nil
	}
	out := new(StableComponent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StableComponent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StableComponentList) DeepCopyInto(out *StableComponentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StableComponent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StableComponentList.
func (in *StableComponentList) DeepCopy() *StableComponentList {
	if in == nil {
		return nil
	}
	out := new(StableComponentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StableComponentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StableComponentSpec) DeepCopyInto(out *StableComponentSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StableComponentSpec.
func (in *StableComponentSpec) DeepCopy() *StableComponentSpec {
	if in == nil {
		return nil
	}
	out := new(StableComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StableComponentStatus) DeepCopyInto(out *StableComponentStatus) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StableComponentStatus.
func (in *StableComponentStatus) DeepCopy() *StableComponentStatus {
	if in == nil {
		return nil
	}
	out := new(StableComponentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StagingCtrl) DeepCopyInto(out *StagingCtrl) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StagingCtrl.
func (in *StagingCtrl) DeepCopy() *StagingCtrl {
	if in == nil {
		return nil
	}
	out := new(StagingCtrl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Team) DeepCopyInto(out *Team) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Team.
func (in *Team) DeepCopy() *Team {
	if in == nil {
		return nil
	}
	out := new(Team)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Team) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamCondition) DeepCopyInto(out *TeamCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamCondition.
func (in *TeamCondition) DeepCopy() *TeamCondition {
	if in == nil {
		return nil
	}
	out := new(TeamCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamDesiredImageTime) DeepCopyInto(out *TeamDesiredImageTime) {
	*out = *in
	in.ImageTime.DeepCopyInto(&out.ImageTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamDesiredImageTime.
func (in *TeamDesiredImageTime) DeepCopy() *TeamDesiredImageTime {
	if in == nil {
		return nil
	}
	out := new(TeamDesiredImageTime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in TeamDesiredImageTimeList) DeepCopyInto(out *TeamDesiredImageTimeList) {
	{
		in := &in
		*out = make(TeamDesiredImageTimeList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamDesiredImageTimeList.
func (in TeamDesiredImageTimeList) DeepCopy() TeamDesiredImageTimeList {
	if in == nil {
		return nil
	}
	out := new(TeamDesiredImageTimeList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamList) DeepCopyInto(out *TeamList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Team, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamList.
func (in *TeamList) DeepCopy() *TeamList {
	if in == nil {
		return nil
	}
	out := new(TeamList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TeamList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamNamespace) DeepCopyInto(out *TeamNamespace) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamNamespace.
func (in *TeamNamespace) DeepCopy() *TeamNamespace {
	if in == nil {
		return nil
	}
	out := new(TeamNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamSpec) DeepCopyInto(out *TeamSpec) {
	*out = *in
	if in.Owners != nil {
		in, out := &in.Owners, &out.Owners
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	in.GitStorage.DeepCopyInto(&out.GitStorage)
	if in.StagingCtrl != nil {
		in, out := &in.StagingCtrl, &out.StagingCtrl
		*out = new(StagingCtrl)
		(*in).DeepCopyInto(*out)
	}
	in.Credential.DeepCopyInto(&out.Credential)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamSpec.
func (in *TeamSpec) DeepCopy() *TeamSpec {
	if in == nil {
		return nil
	}
	out := new(TeamSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamStatus) DeepCopyInto(out *TeamStatus) {
	*out = *in
	out.Namespace = in.Namespace
	if in.StableComponents != nil {
		in, out := &in.StableComponents, &out.StableComponents
		*out = make([]StableComponent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]TeamCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DesiredComponentImageCreatedTime != nil {
		in, out := &in.DesiredComponentImageCreatedTime, &out.DesiredComponentImageCreatedTime
		*out = make(map[string]map[string]DesiredImageTime, len(*in))
		for key, val := range *in {
			var outVal map[string]DesiredImageTime
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(map[string]DesiredImageTime, len(*in))
				for key, val := range *in {
					(*out)[key] = *val.DeepCopy()
				}
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamStatus.
func (in *TeamStatus) DeepCopy() *TeamStatus {
	if in == nil {
		return nil
	}
	out := new(TeamStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Teamcity) DeepCopyInto(out *Teamcity) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Teamcity.
func (in *Teamcity) DeepCopy() *Teamcity {
	if in == nil {
		return nil
	}
	out := new(Teamcity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestRunner) DeepCopyInto(out *TestRunner) {
	*out = *in
	out.Teamcity = in.Teamcity
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestRunner.
func (in *TestRunner) DeepCopy() *TestRunner {
	if in == nil {
		return nil
	}
	out := new(TestRunner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenCredential) DeepCopyInto(out *TokenCredential) {
	*out = *in
	if in.TokenRef != nil {
		in, out := &in.TokenRef, &out.TokenRef
		*out = new(corev1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenCredential.
func (in *TokenCredential) DeepCopy() *TokenCredential {
	if in == nil {
		return nil
	}
	out := new(TokenCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UsernamePasswordCredential) DeepCopyInto(out *UsernamePasswordCredential) {
	*out = *in
	if in.UsernameRef != nil {
		in, out := &in.UsernameRef, &out.UsernameRef
		*out = new(corev1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.PasswordRef != nil {
		in, out := &in.PasswordRef, &out.PasswordRef
		*out = new(corev1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UsernamePasswordCredential.
func (in *UsernamePasswordCredential) DeepCopy() *UsernamePasswordCredential {
	if in == nil {
		return nil
	}
	out := new(UsernamePasswordCredential)
	in.DeepCopyInto(out)
	return out
}
