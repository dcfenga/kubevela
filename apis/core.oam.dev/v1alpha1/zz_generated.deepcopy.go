//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023 The KubeVela Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/oam-dev/kubevela/apis/core.oam.dev/common"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplyOncePolicyRule) DeepCopyInto(out *ApplyOncePolicyRule) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	if in.Strategy != nil {
		in, out := &in.Strategy, &out.Strategy
		*out = new(ApplyOnceStrategy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplyOncePolicyRule.
func (in *ApplyOncePolicyRule) DeepCopy() *ApplyOncePolicyRule {
	if in == nil {
		return nil
	}
	out := new(ApplyOncePolicyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplyOncePolicySpec) DeepCopyInto(out *ApplyOncePolicySpec) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]ApplyOncePolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplyOncePolicySpec.
func (in *ApplyOncePolicySpec) DeepCopy() *ApplyOncePolicySpec {
	if in == nil {
		return nil
	}
	out := new(ApplyOncePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplyOnceStrategy) DeepCopyInto(out *ApplyOnceStrategy) {
	*out = *in
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplyOnceStrategy.
func (in *ApplyOnceStrategy) DeepCopy() *ApplyOnceStrategy {
	if in == nil {
		return nil
	}
	out := new(ApplyOnceStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConnection) DeepCopyInto(out *ClusterConnection) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConnection.
func (in *ClusterConnection) DeepCopy() *ClusterConnection {
	if in == nil {
		return nil
	}
	out := new(ClusterConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvBindingSpec) DeepCopyInto(out *EnvBindingSpec) {
	*out = *in
	if in.Envs != nil {
		in, out := &in.Envs, &out.Envs
		*out = make([]EnvConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvBindingSpec.
func (in *EnvBindingSpec) DeepCopy() *EnvBindingSpec {
	if in == nil {
		return nil
	}
	out := new(EnvBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvBindingStatus) DeepCopyInto(out *EnvBindingStatus) {
	*out = *in
	if in.Envs != nil {
		in, out := &in.Envs, &out.Envs
		*out = make([]EnvStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClusterConnections != nil {
		in, out := &in.ClusterConnections, &out.ClusterConnections
		*out = make([]ClusterConnection, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvBindingStatus.
func (in *EnvBindingStatus) DeepCopy() *EnvBindingStatus {
	if in == nil {
		return nil
	}
	out := new(EnvBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvComponentPatch) DeepCopyInto(out *EnvComponentPatch) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Traits != nil {
		in, out := &in.Traits, &out.Traits
		*out = make([]EnvTraitPatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvComponentPatch.
func (in *EnvComponentPatch) DeepCopy() *EnvComponentPatch {
	if in == nil {
		return nil
	}
	out := new(EnvComponentPatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvConfig) DeepCopyInto(out *EnvConfig) {
	*out = *in
	in.Placement.DeepCopyInto(&out.Placement)
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(EnvSelector)
		(*in).DeepCopyInto(*out)
	}
	in.Patch.DeepCopyInto(&out.Patch)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvConfig.
func (in *EnvConfig) DeepCopy() *EnvConfig {
	if in == nil {
		return nil
	}
	out := new(EnvConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvPatch) DeepCopyInto(out *EnvPatch) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]EnvComponentPatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvPatch.
func (in *EnvPatch) DeepCopy() *EnvPatch {
	if in == nil {
		return nil
	}
	out := new(EnvPatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvPlacement) DeepCopyInto(out *EnvPlacement) {
	*out = *in
	if in.ClusterSelector != nil {
		in, out := &in.ClusterSelector, &out.ClusterSelector
		*out = new(common.ClusterSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = new(NamespaceSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvPlacement.
func (in *EnvPlacement) DeepCopy() *EnvPlacement {
	if in == nil {
		return nil
	}
	out := new(EnvPlacement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvSelector) DeepCopyInto(out *EnvSelector) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvSelector.
func (in *EnvSelector) DeepCopy() *EnvSelector {
	if in == nil {
		return nil
	}
	out := new(EnvSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvStatus) DeepCopyInto(out *EnvStatus) {
	*out = *in
	if in.Placements != nil {
		in, out := &in.Placements, &out.Placements
		*out = make([]PlacementDecision, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvStatus.
func (in *EnvStatus) DeepCopy() *EnvStatus {
	if in == nil {
		return nil
	}
	out := new(EnvStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvTraitPatch) DeepCopyInto(out *EnvTraitPatch) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvTraitPatch.
func (in *EnvTraitPatch) DeepCopy() *EnvTraitPatch {
	if in == nil {
		return nil
	}
	out := new(EnvTraitPatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GarbageCollectPolicyRule) DeepCopyInto(out *GarbageCollectPolicyRule) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GarbageCollectPolicyRule.
func (in *GarbageCollectPolicyRule) DeepCopy() *GarbageCollectPolicyRule {
	if in == nil {
		return nil
	}
	out := new(GarbageCollectPolicyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GarbageCollectPolicySpec) DeepCopyInto(out *GarbageCollectPolicySpec) {
	*out = *in
	if in.ApplicationRevisionLimit != nil {
		in, out := &in.ApplicationRevisionLimit, &out.ApplicationRevisionLimit
		*out = new(int)
		**out = **in
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]GarbageCollectPolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GarbageCollectPolicySpec.
func (in *GarbageCollectPolicySpec) DeepCopy() *GarbageCollectPolicySpec {
	if in == nil {
		return nil
	}
	out := new(GarbageCollectPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LegacyObjectTypeIdentifier) DeepCopyInto(out *LegacyObjectTypeIdentifier) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LegacyObjectTypeIdentifier.
func (in *LegacyObjectTypeIdentifier) DeepCopy() *LegacyObjectTypeIdentifier {
	if in == nil {
		return nil
	}
	out := new(LegacyObjectTypeIdentifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceSelector) DeepCopyInto(out *NamespaceSelector) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceSelector.
func (in *NamespaceSelector) DeepCopy() *NamespaceSelector {
	if in == nil {
		return nil
	}
	out := new(NamespaceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReferrer) DeepCopyInto(out *ObjectReferrer) {
	*out = *in
	out.ObjectTypeIdentifier = in.ObjectTypeIdentifier
	in.ObjectSelector.DeepCopyInto(&out.ObjectSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReferrer.
func (in *ObjectReferrer) DeepCopy() *ObjectReferrer {
	if in == nil {
		return nil
	}
	out := new(ObjectReferrer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectSelector) DeepCopyInto(out *ObjectSelector) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DeprecatedLabelSelector != nil {
		in, out := &in.DeprecatedLabelSelector, &out.DeprecatedLabelSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectSelector.
func (in *ObjectSelector) DeepCopy() *ObjectSelector {
	if in == nil {
		return nil
	}
	out := new(ObjectSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectTypeIdentifier) DeepCopyInto(out *ObjectTypeIdentifier) {
	*out = *in
	out.LegacyObjectTypeIdentifier = in.LegacyObjectTypeIdentifier
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectTypeIdentifier.
func (in *ObjectTypeIdentifier) DeepCopy() *ObjectTypeIdentifier {
	if in == nil {
		return nil
	}
	out := new(ObjectTypeIdentifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OverridePolicySpec) DeepCopyInto(out *OverridePolicySpec) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]EnvComponentPatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OverridePolicySpec.
func (in *OverridePolicySpec) DeepCopy() *OverridePolicySpec {
	if in == nil {
		return nil
	}
	out := new(OverridePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Placement) DeepCopyInto(out *Placement) {
	*out = *in
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ClusterLabelSelector != nil {
		in, out := &in.ClusterLabelSelector, &out.ClusterLabelSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DeprecatedClusterSelector != nil {
		in, out := &in.DeprecatedClusterSelector, &out.DeprecatedClusterSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Placement.
func (in *Placement) DeepCopy() *Placement {
	if in == nil {
		return nil
	}
	out := new(Placement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementDecision) DeepCopyInto(out *PlacementDecision) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementDecision.
func (in *PlacementDecision) DeepCopy() *PlacementDecision {
	if in == nil {
		return nil
	}
	out := new(PlacementDecision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Policy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyList) DeepCopyInto(out *PolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Policy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyList.
func (in *PolicyList) DeepCopy() *PolicyList {
	if in == nil {
		return nil
	}
	out := new(PolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadOnlyPolicyRule) DeepCopyInto(out *ReadOnlyPolicyRule) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadOnlyPolicyRule.
func (in *ReadOnlyPolicyRule) DeepCopy() *ReadOnlyPolicyRule {
	if in == nil {
		return nil
	}
	out := new(ReadOnlyPolicyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadOnlyPolicySpec) DeepCopyInto(out *ReadOnlyPolicySpec) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]ReadOnlyPolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadOnlyPolicySpec.
func (in *ReadOnlyPolicySpec) DeepCopy() *ReadOnlyPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ReadOnlyPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RefObjectsComponentSpec) DeepCopyInto(out *RefObjectsComponentSpec) {
	*out = *in
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]ObjectReferrer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.URLs != nil {
		in, out := &in.URLs, &out.URLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RefObjectsComponentSpec.
func (in *RefObjectsComponentSpec) DeepCopy() *RefObjectsComponentSpec {
	if in == nil {
		return nil
	}
	out := new(RefObjectsComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationPolicySpec) DeepCopyInto(out *ReplicationPolicySpec) {
	*out = *in
	if in.Keys != nil {
		in, out := &in.Keys, &out.Keys
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationPolicySpec.
func (in *ReplicationPolicySpec) DeepCopy() *ReplicationPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ReplicationPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcePolicyRuleSelector) DeepCopyInto(out *ResourcePolicyRuleSelector) {
	*out = *in
	if in.CompNames != nil {
		in, out := &in.CompNames, &out.CompNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CompTypes != nil {
		in, out := &in.CompTypes, &out.CompTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OAMResourceTypes != nil {
		in, out := &in.OAMResourceTypes, &out.OAMResourceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TraitTypes != nil {
		in, out := &in.TraitTypes, &out.TraitTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ResourceTypes != nil {
		in, out := &in.ResourceTypes, &out.ResourceTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ResourceNames != nil {
		in, out := &in.ResourceNames, &out.ResourceNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcePolicyRuleSelector.
func (in *ResourcePolicyRuleSelector) DeepCopy() *ResourcePolicyRuleSelector {
	if in == nil {
		return nil
	}
	out := new(ResourcePolicyRuleSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceUpdatePolicyRule) DeepCopyInto(out *ResourceUpdatePolicyRule) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	in.Strategy.DeepCopyInto(&out.Strategy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceUpdatePolicyRule.
func (in *ResourceUpdatePolicyRule) DeepCopy() *ResourceUpdatePolicyRule {
	if in == nil {
		return nil
	}
	out := new(ResourceUpdatePolicyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceUpdatePolicySpec) DeepCopyInto(out *ResourceUpdatePolicySpec) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]ResourceUpdatePolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceUpdatePolicySpec.
func (in *ResourceUpdatePolicySpec) DeepCopy() *ResourceUpdatePolicySpec {
	if in == nil {
		return nil
	}
	out := new(ResourceUpdatePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceUpdateStrategy) DeepCopyInto(out *ResourceUpdateStrategy) {
	*out = *in
	if in.RecreateFields != nil {
		in, out := &in.RecreateFields, &out.RecreateFields
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceUpdateStrategy.
func (in *ResourceUpdateStrategy) DeepCopy() *ResourceUpdateStrategy {
	if in == nil {
		return nil
	}
	out := new(ResourceUpdateStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedResourcePolicyRule) DeepCopyInto(out *SharedResourcePolicyRule) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedResourcePolicyRule.
func (in *SharedResourcePolicyRule) DeepCopy() *SharedResourcePolicyRule {
	if in == nil {
		return nil
	}
	out := new(SharedResourcePolicyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedResourcePolicySpec) DeepCopyInto(out *SharedResourcePolicySpec) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]SharedResourcePolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedResourcePolicySpec.
func (in *SharedResourcePolicySpec) DeepCopy() *SharedResourcePolicySpec {
	if in == nil {
		return nil
	}
	out := new(SharedResourcePolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TakeOverPolicyRule) DeepCopyInto(out *TakeOverPolicyRule) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TakeOverPolicyRule.
func (in *TakeOverPolicyRule) DeepCopy() *TakeOverPolicyRule {
	if in == nil {
		return nil
	}
	out := new(TakeOverPolicyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TakeOverPolicySpec) DeepCopyInto(out *TakeOverPolicySpec) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]TakeOverPolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TakeOverPolicySpec.
func (in *TakeOverPolicySpec) DeepCopy() *TakeOverPolicySpec {
	if in == nil {
		return nil
	}
	out := new(TakeOverPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TopologyPolicySpec) DeepCopyInto(out *TopologyPolicySpec) {
	*out = *in
	in.Placement.DeepCopyInto(&out.Placement)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopologyPolicySpec.
func (in *TopologyPolicySpec) DeepCopy() *TopologyPolicySpec {
	if in == nil {
		return nil
	}
	out := new(TopologyPolicySpec)
	in.DeepCopyInto(out)
	return out
}
