// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlespannerinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googlespannerinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference interface {
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
	InternalValue() *GoogleSpannerInstanceAutoscalingConfigAutoscalingLimits
	SetInternalValue(val *GoogleSpannerInstanceAutoscalingConfigAutoscalingLimits)
	MaxNodes() *float64
	SetMaxNodes(val *float64)
	MaxNodesInput() *float64
	MaxProcessingUnits() *float64
	SetMaxProcessingUnits(val *float64)
	MaxProcessingUnitsInput() *float64
	MinNodes() *float64
	SetMinNodes(val *float64)
	MinNodesInput() *float64
	MinProcessingUnits() *float64
	SetMinProcessingUnits(val *float64)
	MinProcessingUnitsInput() *float64
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
	ResetMaxNodes()
	ResetMaxProcessingUnits()
	ResetMinNodes()
	ResetMinProcessingUnits()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference
type jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) InternalValue() *GoogleSpannerInstanceAutoscalingConfigAutoscalingLimits {
	var returns *GoogleSpannerInstanceAutoscalingConfigAutoscalingLimits
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) MaxNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) MaxNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) MaxProcessingUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxProcessingUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) MaxProcessingUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxProcessingUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) MinNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) MinNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) MinProcessingUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minProcessingUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) MinProcessingUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minProcessingUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSpannerInstance.GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference_Override(g GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSpannerInstance.GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference)SetInternalValue(val *GoogleSpannerInstanceAutoscalingConfigAutoscalingLimits) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference)SetMaxNodes(val *float64) {
	if err := j.validateSetMaxNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxNodes",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference)SetMaxProcessingUnits(val *float64) {
	if err := j.validateSetMaxProcessingUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxProcessingUnits",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference)SetMinNodes(val *float64) {
	if err := j.validateSetMinNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNodes",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference)SetMinProcessingUnits(val *float64) {
	if err := j.validateSetMinProcessingUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minProcessingUnits",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) ResetMaxNodes() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxNodes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) ResetMaxProcessingUnits() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxProcessingUnits",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) ResetMinNodes() {
	_jsii_.InvokeVoid(
		g,
		"resetMinNodes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) ResetMinProcessingUnits() {
	_jsii_.InvokeVoid(
		g,
		"resetMinProcessingUnits",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleSpannerInstanceAutoscalingConfigAutoscalingLimitsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

