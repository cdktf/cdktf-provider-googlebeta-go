// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputeinstancefromtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googlecomputeinstancefromtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference interface {
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
	EnableNestedVirtualization() interface{}
	SetEnableNestedVirtualization(val interface{})
	EnableNestedVirtualizationInput() interface{}
	EnableUefiNetworking() interface{}
	SetEnableUefiNetworking(val interface{})
	EnableUefiNetworkingInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleComputeInstanceFromTemplateAdvancedMachineFeatures
	SetInternalValue(val *GoogleComputeInstanceFromTemplateAdvancedMachineFeatures)
	PerformanceMonitoringUnit() *string
	SetPerformanceMonitoringUnit(val *string)
	PerformanceMonitoringUnitInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ThreadsPerCore() *float64
	SetThreadsPerCore(val *float64)
	ThreadsPerCoreInput() *float64
	TurboMode() *string
	SetTurboMode(val *string)
	TurboModeInput() *string
	VisibleCoreCount() *float64
	SetVisibleCoreCount(val *float64)
	VisibleCoreCountInput() *float64
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
	ResetEnableNestedVirtualization()
	ResetEnableUefiNetworking()
	ResetPerformanceMonitoringUnit()
	ResetThreadsPerCore()
	ResetTurboMode()
	ResetVisibleCoreCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference
type jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) EnableNestedVirtualization() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNestedVirtualization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) EnableNestedVirtualizationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNestedVirtualizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) EnableUefiNetworking() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUefiNetworking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) EnableUefiNetworkingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableUefiNetworkingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) InternalValue() *GoogleComputeInstanceFromTemplateAdvancedMachineFeatures {
	var returns *GoogleComputeInstanceFromTemplateAdvancedMachineFeatures
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) PerformanceMonitoringUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceMonitoringUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) PerformanceMonitoringUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceMonitoringUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) ThreadsPerCore() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadsPerCore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) ThreadsPerCoreInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"threadsPerCoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) TurboMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"turboMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) TurboModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"turboModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) VisibleCoreCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"visibleCoreCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) VisibleCoreCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"visibleCoreCountInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeInstanceFromTemplate.GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference_Override(g GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeInstanceFromTemplate.GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference)SetEnableNestedVirtualization(val interface{}) {
	if err := j.validateSetEnableNestedVirtualizationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNestedVirtualization",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference)SetEnableUefiNetworking(val interface{}) {
	if err := j.validateSetEnableUefiNetworkingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableUefiNetworking",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference)SetInternalValue(val *GoogleComputeInstanceFromTemplateAdvancedMachineFeatures) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference)SetPerformanceMonitoringUnit(val *string) {
	if err := j.validateSetPerformanceMonitoringUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceMonitoringUnit",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference)SetThreadsPerCore(val *float64) {
	if err := j.validateSetThreadsPerCoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threadsPerCore",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference)SetTurboMode(val *string) {
	if err := j.validateSetTurboModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"turboMode",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference)SetVisibleCoreCount(val *float64) {
	if err := j.validateSetVisibleCoreCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibleCoreCount",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) ResetEnableNestedVirtualization() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableNestedVirtualization",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) ResetEnableUefiNetworking() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableUefiNetworking",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) ResetPerformanceMonitoringUnit() {
	_jsii_.InvokeVoid(
		g,
		"resetPerformanceMonitoringUnit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) ResetThreadsPerCore() {
	_jsii_.InvokeVoid(
		g,
		"resetThreadsPerCore",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) ResetTurboMode() {
	_jsii_.InvokeVoid(
		g,
		"resetTurboMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) ResetVisibleCoreCount() {
	_jsii_.InvokeVoid(
		g,
		"resetVisibleCoreCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateAdvancedMachineFeaturesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

