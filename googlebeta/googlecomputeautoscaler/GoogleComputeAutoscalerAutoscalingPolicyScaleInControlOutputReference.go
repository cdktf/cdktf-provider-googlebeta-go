// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputeautoscaler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googlecomputeautoscaler/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference interface {
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
	InternalValue() *GoogleComputeAutoscalerAutoscalingPolicyScaleInControl
	SetInternalValue(val *GoogleComputeAutoscalerAutoscalingPolicyScaleInControl)
	MaxScaledInReplicas() GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference
	MaxScaledInReplicasInput() *GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeWindowSec() *float64
	SetTimeWindowSec(val *float64)
	TimeWindowSecInput() *float64
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
	PutMaxScaledInReplicas(value *GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas)
	ResetMaxScaledInReplicas()
	ResetTimeWindowSec()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference
type jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) InternalValue() *GoogleComputeAutoscalerAutoscalingPolicyScaleInControl {
	var returns *GoogleComputeAutoscalerAutoscalingPolicyScaleInControl
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) MaxScaledInReplicas() GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference {
	var returns GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicasOutputReference
	_jsii_.Get(
		j,
		"maxScaledInReplicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) MaxScaledInReplicasInput() *GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas {
	var returns *GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas
	_jsii_.Get(
		j,
		"maxScaledInReplicasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) TimeWindowSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeWindowSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) TimeWindowSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeWindowSecInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeAutoscaler.GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference_Override(g GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeAutoscaler.GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference)SetInternalValue(val *GoogleComputeAutoscalerAutoscalingPolicyScaleInControl) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference)SetTimeWindowSec(val *float64) {
	if err := j.validateSetTimeWindowSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeWindowSec",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) PutMaxScaledInReplicas(value *GoogleComputeAutoscalerAutoscalingPolicyScaleInControlMaxScaledInReplicas) {
	if err := g.validatePutMaxScaledInReplicasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaxScaledInReplicas",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) ResetMaxScaledInReplicas() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxScaledInReplicas",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) ResetTimeWindowSec() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeWindowSec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeAutoscalerAutoscalingPolicyScaleInControlOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

