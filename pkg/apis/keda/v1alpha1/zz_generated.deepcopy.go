// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/batch/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthEnvironment) DeepCopyInto(out *AuthEnvironment) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthEnvironment.
func (in *AuthEnvironment) DeepCopy() *AuthEnvironment {
	if in == nil {
		return nil
	}
	out := new(AuthEnvironment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthPodIdentity) DeepCopyInto(out *AuthPodIdentity) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthPodIdentity.
func (in *AuthPodIdentity) DeepCopy() *AuthPodIdentity {
	if in == nil {
		return nil
	}
	out := new(AuthPodIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthSecretTargetRef) DeepCopyInto(out *AuthSecretTargetRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthSecretTargetRef.
func (in *AuthSecretTargetRef) DeepCopy() *AuthSecretTargetRef {
	if in == nil {
		return nil
	}
	out := new(AuthSecretTargetRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Credential) DeepCopyInto(out *Credential) {
	*out = *in
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
func (in *HashiCorpVault) DeepCopyInto(out *HashiCorpVault) {
	*out = *in
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]VaultSecret, len(*in))
		copy(*out, *in)
	}
	out.Credential = in.Credential
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HashiCorpVault.
func (in *HashiCorpVault) DeepCopy() *HashiCorpVault {
	if in == nil {
		return nil
	}
	out := new(HashiCorpVault)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectReference) DeepCopyInto(out *ObjectReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectReference.
func (in *ObjectReference) DeepCopy() *ObjectReference {
	if in == nil {
		return nil
	}
	out := new(ObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleTriggers) DeepCopyInto(out *ScaleTriggers) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AuthenticationRef != nil {
		in, out := &in.AuthenticationRef, &out.AuthenticationRef
		*out = new(ScaledObjectAuthRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleTriggers.
func (in *ScaleTriggers) DeepCopy() *ScaleTriggers {
	if in == nil {
		return nil
	}
	out := new(ScaleTriggers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledObject) DeepCopyInto(out *ScaledObject) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledObject.
func (in *ScaledObject) DeepCopy() *ScaledObject {
	if in == nil {
		return nil
	}
	out := new(ScaledObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScaledObject) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledObjectAuthRef) DeepCopyInto(out *ScaledObjectAuthRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledObjectAuthRef.
func (in *ScaledObjectAuthRef) DeepCopy() *ScaledObjectAuthRef {
	if in == nil {
		return nil
	}
	out := new(ScaledObjectAuthRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledObjectList) DeepCopyInto(out *ScaledObjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScaledObject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledObjectList.
func (in *ScaledObjectList) DeepCopy() *ScaledObjectList {
	if in == nil {
		return nil
	}
	out := new(ScaledObjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScaledObjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledObjectSpec) DeepCopyInto(out *ScaledObjectSpec) {
	*out = *in
	if in.ScaleTargetRef != nil {
		in, out := &in.ScaleTargetRef, &out.ScaleTargetRef
		*out = new(ObjectReference)
		**out = **in
	}
	if in.JobTargetRef != nil {
		in, out := &in.JobTargetRef, &out.JobTargetRef
		*out = new(v1.JobSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.PollingInterval != nil {
		in, out := &in.PollingInterval, &out.PollingInterval
		*out = new(int32)
		**out = **in
	}
	if in.CooldownPeriod != nil {
		in, out := &in.CooldownPeriod, &out.CooldownPeriod
		*out = new(int32)
		**out = **in
	}
	if in.MinReplicaCount != nil {
		in, out := &in.MinReplicaCount, &out.MinReplicaCount
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicaCount != nil {
		in, out := &in.MaxReplicaCount, &out.MaxReplicaCount
		*out = new(int32)
		**out = **in
	}
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]ScaleTriggers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledObjectSpec.
func (in *ScaledObjectSpec) DeepCopy() *ScaledObjectSpec {
	if in == nil {
		return nil
	}
	out := new(ScaledObjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaledObjectStatus) DeepCopyInto(out *ScaledObjectStatus) {
	*out = *in
	if in.LastActiveTime != nil {
		in, out := &in.LastActiveTime, &out.LastActiveTime
		*out = (*in).DeepCopy()
	}
	if in.ExternalMetricNames != nil {
		in, out := &in.ExternalMetricNames, &out.ExternalMetricNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaledObjectStatus.
func (in *ScaledObjectStatus) DeepCopy() *ScaledObjectStatus {
	if in == nil {
		return nil
	}
	out := new(ScaledObjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerAuthentication) DeepCopyInto(out *TriggerAuthentication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerAuthentication.
func (in *TriggerAuthentication) DeepCopy() *TriggerAuthentication {
	if in == nil {
		return nil
	}
	out := new(TriggerAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TriggerAuthentication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerAuthenticationList) DeepCopyInto(out *TriggerAuthenticationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TriggerAuthentication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerAuthenticationList.
func (in *TriggerAuthenticationList) DeepCopy() *TriggerAuthenticationList {
	if in == nil {
		return nil
	}
	out := new(TriggerAuthenticationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TriggerAuthenticationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerAuthenticationSpec) DeepCopyInto(out *TriggerAuthenticationSpec) {
	*out = *in
	out.PodIdentity = in.PodIdentity
	if in.SecretTargetRef != nil {
		in, out := &in.SecretTargetRef, &out.SecretTargetRef
		*out = make([]AuthSecretTargetRef, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]AuthEnvironment, len(*in))
		copy(*out, *in)
	}
	in.HashiCorpVault.DeepCopyInto(&out.HashiCorpVault)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerAuthenticationSpec.
func (in *TriggerAuthenticationSpec) DeepCopy() *TriggerAuthenticationSpec {
	if in == nil {
		return nil
	}
	out := new(TriggerAuthenticationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VaultSecret) DeepCopyInto(out *VaultSecret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VaultSecret.
func (in *VaultSecret) DeepCopy() *VaultSecret {
	if in == nil {
		return nil
	}
	out := new(VaultSecret)
	in.DeepCopyInto(out)
	return out
}
