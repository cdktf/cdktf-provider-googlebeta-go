// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlenetappvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googlenetappvolume/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Hour() *float64
	SetHour(val *float64)
	HourInput() *float64
	InternalValue() *GoogleNetappVolumeSnapshotPolicyDailySchedule
	SetInternalValue(val *GoogleNetappVolumeSnapshotPolicyDailySchedule)
	Minute() *float64
	SetMinute(val *float64)
	MinuteInput() *float64
	SnapshotsToKeep() *float64
	SetSnapshotsToKeep(val *float64)
	SnapshotsToKeepInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetHour()
	ResetMinute()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference
type jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) Hour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) HourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) InternalValue() *GoogleNetappVolumeSnapshotPolicyDailySchedule {
	var returns *GoogleNetappVolumeSnapshotPolicyDailySchedule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) Minute() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) MinuteInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minuteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) SnapshotsToKeep() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotsToKeep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) SnapshotsToKeepInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotsToKeepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetappVolume.GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference_Override(g GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetappVolume.GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference)SetHour(val *float64) {
	if err := j.validateSetHourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hour",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference)SetInternalValue(val *GoogleNetappVolumeSnapshotPolicyDailySchedule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference)SetMinute(val *float64) {
	if err := j.validateSetMinuteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference)SetSnapshotsToKeep(val *float64) {
	if err := j.validateSetSnapshotsToKeepParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotsToKeep",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) ResetHour() {
	_jsii_.InvokeVoid(
		g,
		"resetHour",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) ResetMinute() {
	_jsii_.InvokeVoid(
		g,
		"resetMinute",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolumeSnapshotPolicyDailyScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

