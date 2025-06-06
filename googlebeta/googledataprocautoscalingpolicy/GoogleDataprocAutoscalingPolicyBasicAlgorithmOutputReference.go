// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledataprocautoscalingpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googledataprocautoscalingpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference interface {
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
	CooldownPeriod() *string
	SetCooldownPeriod(val *string)
	CooldownPeriodInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDataprocAutoscalingPolicyBasicAlgorithm
	SetInternalValue(val *GoogleDataprocAutoscalingPolicyBasicAlgorithm)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	YarnConfig() GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference
	YarnConfigInput() *GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfig
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
	PutYarnConfig(value *GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfig)
	ResetCooldownPeriod()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference
type jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) CooldownPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cooldownPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) CooldownPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cooldownPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) InternalValue() *GoogleDataprocAutoscalingPolicyBasicAlgorithm {
	var returns *GoogleDataprocAutoscalingPolicyBasicAlgorithm
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) YarnConfig() GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference {
	var returns GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfigOutputReference
	_jsii_.Get(
		j,
		"yarnConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) YarnConfigInput() *GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfig {
	var returns *GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfig
	_jsii_.Get(
		j,
		"yarnConfigInput",
		&returns,
	)
	return returns
}


func NewGoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocAutoscalingPolicy.GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference_Override(g GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocAutoscalingPolicy.GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference)SetCooldownPeriod(val *string) {
	if err := j.validateSetCooldownPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cooldownPeriod",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference)SetInternalValue(val *GoogleDataprocAutoscalingPolicyBasicAlgorithm) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) PutYarnConfig(value *GoogleDataprocAutoscalingPolicyBasicAlgorithmYarnConfig) {
	if err := g.validatePutYarnConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putYarnConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) ResetCooldownPeriod() {
	_jsii_.InvokeVoid(
		g,
		"resetCooldownPeriod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataprocAutoscalingPolicyBasicAlgorithmOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

