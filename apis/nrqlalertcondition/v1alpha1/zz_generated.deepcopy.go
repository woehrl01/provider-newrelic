//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertCondition) DeepCopyInto(out *AlertCondition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertCondition.
func (in *AlertCondition) DeepCopy() *AlertCondition {
	if in == nil {
		return nil
	}
	out := new(AlertCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertCondition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertConditionList) DeepCopyInto(out *AlertConditionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlertCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertConditionList.
func (in *AlertConditionList) DeepCopy() *AlertConditionList {
	if in == nil {
		return nil
	}
	out := new(AlertConditionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertConditionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertConditionObservation) DeepCopyInto(out *AlertConditionObservation) {
	*out = *in
	if in.EntityGUID != nil {
		in, out := &in.EntityGUID, &out.EntityGUID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertConditionObservation.
func (in *AlertConditionObservation) DeepCopy() *AlertConditionObservation {
	if in == nil {
		return nil
	}
	out := new(AlertConditionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertConditionParameters) DeepCopyInto(out *AlertConditionParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(float64)
		**out = **in
	}
	if in.AggregationDelay != nil {
		in, out := &in.AggregationDelay, &out.AggregationDelay
		*out = new(string)
		**out = **in
	}
	if in.AggregationMethod != nil {
		in, out := &in.AggregationMethod, &out.AggregationMethod
		*out = new(string)
		**out = **in
	}
	if in.AggregationTimer != nil {
		in, out := &in.AggregationTimer, &out.AggregationTimer
		*out = new(string)
		**out = **in
	}
	if in.AggregationWindow != nil {
		in, out := &in.AggregationWindow, &out.AggregationWindow
		*out = new(float64)
		**out = **in
	}
	if in.BaselineDirection != nil {
		in, out := &in.BaselineDirection, &out.BaselineDirection
		*out = new(string)
		**out = **in
	}
	if in.CloseViolationsOnExpiration != nil {
		in, out := &in.CloseViolationsOnExpiration, &out.CloseViolationsOnExpiration
		*out = new(bool)
		**out = **in
	}
	if in.Critical != nil {
		in, out := &in.Critical, &out.Critical
		*out = make([]CriticalParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ExpirationDuration != nil {
		in, out := &in.ExpirationDuration, &out.ExpirationDuration
		*out = new(float64)
		**out = **in
	}
	if in.FillOption != nil {
		in, out := &in.FillOption, &out.FillOption
		*out = new(string)
		**out = **in
	}
	if in.FillValue != nil {
		in, out := &in.FillValue, &out.FillValue
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Nrql != nil {
		in, out := &in.Nrql, &out.Nrql
		*out = make([]NrqlParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OpenViolationOnExpiration != nil {
		in, out := &in.OpenViolationOnExpiration, &out.OpenViolationOnExpiration
		*out = new(bool)
		**out = **in
	}
	if in.PolicyID != nil {
		in, out := &in.PolicyID, &out.PolicyID
		*out = new(float64)
		**out = **in
	}
	if in.RunbookURL != nil {
		in, out := &in.RunbookURL, &out.RunbookURL
		*out = new(string)
		**out = **in
	}
	if in.SlideBy != nil {
		in, out := &in.SlideBy, &out.SlideBy
		*out = new(float64)
		**out = **in
	}
	if in.Term != nil {
		in, out := &in.Term, &out.Term
		*out = make([]TermParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.ValueFunction != nil {
		in, out := &in.ValueFunction, &out.ValueFunction
		*out = new(string)
		**out = **in
	}
	if in.ViolationTimeLimit != nil {
		in, out := &in.ViolationTimeLimit, &out.ViolationTimeLimit
		*out = new(string)
		**out = **in
	}
	if in.ViolationTimeLimitSeconds != nil {
		in, out := &in.ViolationTimeLimitSeconds, &out.ViolationTimeLimitSeconds
		*out = new(float64)
		**out = **in
	}
	if in.Warning != nil {
		in, out := &in.Warning, &out.Warning
		*out = make([]WarningParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertConditionParameters.
func (in *AlertConditionParameters) DeepCopy() *AlertConditionParameters {
	if in == nil {
		return nil
	}
	out := new(AlertConditionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertConditionSpec) DeepCopyInto(out *AlertConditionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertConditionSpec.
func (in *AlertConditionSpec) DeepCopy() *AlertConditionSpec {
	if in == nil {
		return nil
	}
	out := new(AlertConditionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertConditionStatus) DeepCopyInto(out *AlertConditionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertConditionStatus.
func (in *AlertConditionStatus) DeepCopy() *AlertConditionStatus {
	if in == nil {
		return nil
	}
	out := new(AlertConditionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CriticalObservation) DeepCopyInto(out *CriticalObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CriticalObservation.
func (in *CriticalObservation) DeepCopy() *CriticalObservation {
	if in == nil {
		return nil
	}
	out := new(CriticalObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CriticalParameters) DeepCopyInto(out *CriticalParameters) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(float64)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.ThresholdDuration != nil {
		in, out := &in.ThresholdDuration, &out.ThresholdDuration
		*out = new(float64)
		**out = **in
	}
	if in.ThresholdOccurrences != nil {
		in, out := &in.ThresholdOccurrences, &out.ThresholdOccurrences
		*out = new(string)
		**out = **in
	}
	if in.TimeFunction != nil {
		in, out := &in.TimeFunction, &out.TimeFunction
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CriticalParameters.
func (in *CriticalParameters) DeepCopy() *CriticalParameters {
	if in == nil {
		return nil
	}
	out := new(CriticalParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NrqlObservation) DeepCopyInto(out *NrqlObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NrqlObservation.
func (in *NrqlObservation) DeepCopy() *NrqlObservation {
	if in == nil {
		return nil
	}
	out := new(NrqlObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NrqlParameters) DeepCopyInto(out *NrqlParameters) {
	*out = *in
	if in.EvaluationOffset != nil {
		in, out := &in.EvaluationOffset, &out.EvaluationOffset
		*out = new(float64)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(string)
		**out = **in
	}
	if in.SinceValue != nil {
		in, out := &in.SinceValue, &out.SinceValue
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NrqlParameters.
func (in *NrqlParameters) DeepCopy() *NrqlParameters {
	if in == nil {
		return nil
	}
	out := new(NrqlParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TermObservation) DeepCopyInto(out *TermObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TermObservation.
func (in *TermObservation) DeepCopy() *TermObservation {
	if in == nil {
		return nil
	}
	out := new(TermObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TermParameters) DeepCopyInto(out *TermParameters) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(float64)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(string)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.ThresholdDuration != nil {
		in, out := &in.ThresholdDuration, &out.ThresholdDuration
		*out = new(float64)
		**out = **in
	}
	if in.ThresholdOccurrences != nil {
		in, out := &in.ThresholdOccurrences, &out.ThresholdOccurrences
		*out = new(string)
		**out = **in
	}
	if in.TimeFunction != nil {
		in, out := &in.TimeFunction, &out.TimeFunction
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TermParameters.
func (in *TermParameters) DeepCopy() *TermParameters {
	if in == nil {
		return nil
	}
	out := new(TermParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WarningObservation) DeepCopyInto(out *WarningObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WarningObservation.
func (in *WarningObservation) DeepCopy() *WarningObservation {
	if in == nil {
		return nil
	}
	out := new(WarningObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WarningParameters) DeepCopyInto(out *WarningParameters) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(float64)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.ThresholdDuration != nil {
		in, out := &in.ThresholdDuration, &out.ThresholdDuration
		*out = new(float64)
		**out = **in
	}
	if in.ThresholdOccurrences != nil {
		in, out := &in.ThresholdOccurrences, &out.ThresholdOccurrences
		*out = new(string)
		**out = **in
	}
	if in.TimeFunction != nil {
		in, out := &in.TimeFunction, &out.TimeFunction
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WarningParameters.
func (in *WarningParameters) DeepCopy() *WarningParameters {
	if in == nil {
		return nil
	}
	out := new(WarningParameters)
	in.DeepCopyInto(out)
	return out
}
