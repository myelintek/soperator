//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FailureActions) DeepCopyInto(out *FailureActions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FailureActions.
func (in *FailureActions) DeepCopy() *FailureActions {
	if in == nil {
		return nil
	}
	out := new(FailureActions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sNodeFilter) DeepCopyInto(out *K8sNodeFilter) {
	*out = *in
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
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sNodeFilter.
func (in *K8sNodeFilter) DeepCopy() *K8sNodeFilter {
	if in == nil {
		return nil
	}
	out := new(K8sNodeFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NCCLBenchmark) DeepCopyInto(out *NCCLBenchmark) {
	*out = *in
	out.NCCLSettings = in.NCCLSettings
	out.FailureActions = in.FailureActions
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NCCLBenchmark.
func (in *NCCLBenchmark) DeepCopy() *NCCLBenchmark {
	if in == nil {
		return nil
	}
	out := new(NCCLBenchmark)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NCCLSettings) DeepCopyInto(out *NCCLSettings) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NCCLSettings.
func (in *NCCLSettings) DeepCopy() *NCCLSettings {
	if in == nil {
		return nil
	}
	out := new(NCCLSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeContainer) DeepCopyInto(out *NodeContainer) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeContainer.
func (in *NodeContainer) DeepCopy() *NodeContainer {
	if in == nil {
		return nil
	}
	out := new(NodeContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeVolume) DeepCopyInto(out *NodeVolume) {
	*out = *in
	if in.VolumeSourceName != nil {
		in, out := &in.VolumeSourceName, &out.VolumeSourceName
		*out = new(string)
		**out = **in
	}
	if in.VolumeClaimTemplateSpec != nil {
		in, out := &in.VolumeClaimTemplateSpec, &out.VolumeClaimTemplateSpec
		*out = new(corev1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeVolume.
func (in *NodeVolume) DeepCopy() *NodeVolume {
	if in == nil {
		return nil
	}
	out := new(NodeVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeVolumeJailSubMount) DeepCopyInto(out *NodeVolumeJailSubMount) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeVolumeJailSubMount.
func (in *NodeVolumeJailSubMount) DeepCopy() *NodeVolumeJailSubMount {
	if in == nil {
		return nil
	}
	out := new(NodeVolumeJailSubMount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeriodicChecks) DeepCopyInto(out *PeriodicChecks) {
	*out = *in
	out.NCCLBenchmark = in.NCCLBenchmark
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeriodicChecks.
func (in *PeriodicChecks) DeepCopy() *PeriodicChecks {
	if in == nil {
		return nil
	}
	out := new(PeriodicChecks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PopulateJail) DeepCopyInto(out *PopulateJail) {
	*out = *in
	if in.JailSnapshotVolume != nil {
		in, out := &in.JailSnapshotVolume, &out.JailSnapshotVolume
		*out = new(NodeVolume)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PopulateJail.
func (in *PopulateJail) DeepCopy() *PopulateJail {
	if in == nil {
		return nil
	}
	out := new(PopulateJail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKey) DeepCopyInto(out *SecretKey) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKey.
func (in *SecretKey) DeepCopy() *SecretKey {
	if in == nil {
		return nil
	}
	out := new(SecretKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Secrets) DeepCopyInto(out *Secrets) {
	*out = *in
	out.MungeKey = in.MungeKey
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Secrets.
func (in *Secrets) DeepCopy() *Secrets {
	if in == nil {
		return nil
	}
	out := new(Secrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmCluster) DeepCopyInto(out *SlurmCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmCluster.
func (in *SlurmCluster) DeepCopy() *SlurmCluster {
	if in == nil {
		return nil
	}
	out := new(SlurmCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SlurmCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmClusterList) DeepCopyInto(out *SlurmClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SlurmCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmClusterList.
func (in *SlurmClusterList) DeepCopy() *SlurmClusterList {
	if in == nil {
		return nil
	}
	out := new(SlurmClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SlurmClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmClusterSpec) DeepCopyInto(out *SlurmClusterSpec) {
	*out = *in
	in.PopulateJail.DeepCopyInto(&out.PopulateJail)
	out.PeriodicChecks = in.PeriodicChecks
	if in.K8sNodeFilters != nil {
		in, out := &in.K8sNodeFilters, &out.K8sNodeFilters
		*out = make([]K8sNodeFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeSources != nil {
		in, out := &in.VolumeSources, &out.VolumeSources
		*out = make([]VolumeSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Secrets = in.Secrets
	in.SlurmNodes.DeepCopyInto(&out.SlurmNodes)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmClusterSpec.
func (in *SlurmClusterSpec) DeepCopy() *SlurmClusterSpec {
	if in == nil {
		return nil
	}
	out := new(SlurmClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmClusterStatus) DeepCopyInto(out *SlurmClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Phase != nil {
		in, out := &in.Phase, &out.Phase
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmClusterStatus.
func (in *SlurmClusterStatus) DeepCopy() *SlurmClusterStatus {
	if in == nil {
		return nil
	}
	out := new(SlurmClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmNode) DeepCopyInto(out *SlurmNode) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmNode.
func (in *SlurmNode) DeepCopy() *SlurmNode {
	if in == nil {
		return nil
	}
	out := new(SlurmNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmNodeController) DeepCopyInto(out *SlurmNodeController) {
	*out = *in
	out.SlurmNode = in.SlurmNode
	in.Slurmctld.DeepCopyInto(&out.Slurmctld)
	in.Munge.DeepCopyInto(&out.Munge)
	in.Volumes.DeepCopyInto(&out.Volumes)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmNodeController.
func (in *SlurmNodeController) DeepCopy() *SlurmNodeController {
	if in == nil {
		return nil
	}
	out := new(SlurmNodeController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmNodeControllerVolumes) DeepCopyInto(out *SlurmNodeControllerVolumes) {
	*out = *in
	in.Spool.DeepCopyInto(&out.Spool)
	in.Jail.DeepCopyInto(&out.Jail)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmNodeControllerVolumes.
func (in *SlurmNodeControllerVolumes) DeepCopy() *SlurmNodeControllerVolumes {
	if in == nil {
		return nil
	}
	out := new(SlurmNodeControllerVolumes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmNodeLogin) DeepCopyInto(out *SlurmNodeLogin) {
	*out = *in
	out.SlurmNode = in.SlurmNode
	in.Sshd.DeepCopyInto(&out.Sshd)
	in.Munge.DeepCopyInto(&out.Munge)
	if in.SshdServiceAnnotations != nil {
		in, out := &in.SshdServiceAnnotations, &out.SshdServiceAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.SshRootPublicKeysConfigMap = in.SshRootPublicKeysConfigMap
	in.Volumes.DeepCopyInto(&out.Volumes)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmNodeLogin.
func (in *SlurmNodeLogin) DeepCopy() *SlurmNodeLogin {
	if in == nil {
		return nil
	}
	out := new(SlurmNodeLogin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmNodeLoginVolumes) DeepCopyInto(out *SlurmNodeLoginVolumes) {
	*out = *in
	in.Jail.DeepCopyInto(&out.Jail)
	if in.JailSubMounts != nil {
		in, out := &in.JailSubMounts, &out.JailSubMounts
		*out = make([]NodeVolumeJailSubMount, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmNodeLoginVolumes.
func (in *SlurmNodeLoginVolumes) DeepCopy() *SlurmNodeLoginVolumes {
	if in == nil {
		return nil
	}
	out := new(SlurmNodeLoginVolumes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmNodeWorker) DeepCopyInto(out *SlurmNodeWorker) {
	*out = *in
	out.SlurmNode = in.SlurmNode
	in.Slurmd.DeepCopyInto(&out.Slurmd)
	in.Munge.DeepCopyInto(&out.Munge)
	in.Volumes.DeepCopyInto(&out.Volumes)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmNodeWorker.
func (in *SlurmNodeWorker) DeepCopy() *SlurmNodeWorker {
	if in == nil {
		return nil
	}
	out := new(SlurmNodeWorker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmNodeWorkerVolumes) DeepCopyInto(out *SlurmNodeWorkerVolumes) {
	*out = *in
	in.Spool.DeepCopyInto(&out.Spool)
	in.Jail.DeepCopyInto(&out.Jail)
	if in.JailSubMounts != nil {
		in, out := &in.JailSubMounts, &out.JailSubMounts
		*out = make([]NodeVolumeJailSubMount, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmNodeWorkerVolumes.
func (in *SlurmNodeWorkerVolumes) DeepCopy() *SlurmNodeWorkerVolumes {
	if in == nil {
		return nil
	}
	out := new(SlurmNodeWorkerVolumes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SlurmNodes) DeepCopyInto(out *SlurmNodes) {
	*out = *in
	in.Controller.DeepCopyInto(&out.Controller)
	in.Worker.DeepCopyInto(&out.Worker)
	in.Login.DeepCopyInto(&out.Login)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SlurmNodes.
func (in *SlurmNodes) DeepCopy() *SlurmNodes {
	if in == nil {
		return nil
	}
	out := new(SlurmNodes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SshRootPublicKeysConfigMap) DeepCopyInto(out *SshRootPublicKeysConfigMap) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SshRootPublicKeysConfigMap.
func (in *SshRootPublicKeysConfigMap) DeepCopy() *SshRootPublicKeysConfigMap {
	if in == nil {
		return nil
	}
	out := new(SshRootPublicKeysConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSource) DeepCopyInto(out *VolumeSource) {
	*out = *in
	in.VolumeSource.DeepCopyInto(&out.VolumeSource)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSource.
func (in *VolumeSource) DeepCopy() *VolumeSource {
	if in == nil {
		return nil
	}
	out := new(VolumeSource)
	in.DeepCopyInto(out)
	return out
}
