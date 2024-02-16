//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CIOperatorMetadata) DeepCopyInto(out *CIOperatorMetadata) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CIOperatorMetadata.
func (in *CIOperatorMetadata) DeepCopy() *CIOperatorMetadata {
	if in == nil {
		return nil
	}
	out := new(CIOperatorMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageTagOverride) DeepCopyInto(out *ImageTagOverride) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageTagOverride.
func (in *ImageTagOverride) DeepCopy() *ImageTagOverride {
	if in == nil {
		return nil
	}
	out := new(ImageTagOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PayloadOverrides) DeepCopyInto(out *PayloadOverrides) {
	*out = *in
	if in.ImageTagOverrides != nil {
		in, out := &in.ImageTagOverrides, &out.ImageTagOverrides
		*out = make([]ImageTagOverride, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PayloadOverrides.
func (in *PayloadOverrides) DeepCopy() *PayloadOverrides {
	if in == nil {
		return nil
	}
	out := new(PayloadOverrides)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequest) DeepCopyInto(out *PullRequest) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequest.
func (in *PullRequest) DeepCopy() *PullRequest {
	if in == nil {
		return nil
	}
	out := new(PullRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestPayloadJobSpec) DeepCopyInto(out *PullRequestPayloadJobSpec) {
	*out = *in
	out.ReleaseControllerConfig = in.ReleaseControllerConfig
	if in.Jobs != nil {
		in, out := &in.Jobs, &out.Jobs
		*out = make([]ReleaseJobSpec, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestPayloadJobSpec.
func (in *PullRequestPayloadJobSpec) DeepCopy() *PullRequestPayloadJobSpec {
	if in == nil {
		return nil
	}
	out := new(PullRequestPayloadJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestPayloadJobStatus) DeepCopyInto(out *PullRequestPayloadJobStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestPayloadJobStatus.
func (in *PullRequestPayloadJobStatus) DeepCopy() *PullRequestPayloadJobStatus {
	if in == nil {
		return nil
	}
	out := new(PullRequestPayloadJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestPayloadQualificationRun) DeepCopyInto(out *PullRequestPayloadQualificationRun) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestPayloadQualificationRun.
func (in *PullRequestPayloadQualificationRun) DeepCopy() *PullRequestPayloadQualificationRun {
	if in == nil {
		return nil
	}
	out := new(PullRequestPayloadQualificationRun)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PullRequestPayloadQualificationRun) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestPayloadQualificationRunList) DeepCopyInto(out *PullRequestPayloadQualificationRunList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PullRequestPayloadQualificationRun, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestPayloadQualificationRunList.
func (in *PullRequestPayloadQualificationRunList) DeepCopy() *PullRequestPayloadQualificationRunList {
	if in == nil {
		return nil
	}
	out := new(PullRequestPayloadQualificationRunList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PullRequestPayloadQualificationRunList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestPayloadTestSpec) DeepCopyInto(out *PullRequestPayloadTestSpec) {
	*out = *in
	if in.PullRequests != nil {
		in, out := &in.PullRequests, &out.PullRequests
		*out = make([]PullRequestUnderTest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Jobs.DeepCopyInto(&out.Jobs)
	in.PayloadOverrides.DeepCopyInto(&out.PayloadOverrides)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestPayloadTestSpec.
func (in *PullRequestPayloadTestSpec) DeepCopy() *PullRequestPayloadTestSpec {
	if in == nil {
		return nil
	}
	out := new(PullRequestPayloadTestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestPayloadTestStatus) DeepCopyInto(out *PullRequestPayloadTestStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Jobs != nil {
		in, out := &in.Jobs, &out.Jobs
		*out = make([]PullRequestPayloadJobStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestPayloadTestStatus.
func (in *PullRequestPayloadTestStatus) DeepCopy() *PullRequestPayloadTestStatus {
	if in == nil {
		return nil
	}
	out := new(PullRequestPayloadTestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestUnderTest) DeepCopyInto(out *PullRequestUnderTest) {
	*out = *in
	if in.PullRequest != nil {
		in, out := &in.PullRequest, &out.PullRequest
		*out = new(PullRequest)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestUnderTest.
func (in *PullRequestUnderTest) DeepCopy() *PullRequestUnderTest {
	if in == nil {
		return nil
	}
	out := new(PullRequestUnderTest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseControllerConfig) DeepCopyInto(out *ReleaseControllerConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseControllerConfig.
func (in *ReleaseControllerConfig) DeepCopy() *ReleaseControllerConfig {
	if in == nil {
		return nil
	}
	out := new(ReleaseControllerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseJobSpec) DeepCopyInto(out *ReleaseJobSpec) {
	*out = *in
	out.CIOperatorConfig = in.CIOperatorConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseJobSpec.
func (in *ReleaseJobSpec) DeepCopy() *ReleaseJobSpec {
	if in == nil {
		return nil
	}
	out := new(ReleaseJobSpec)
	in.DeepCopyInto(out)
	return out
}
