// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googlecontainercluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference interface {
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
	Count() *float64
	SetCount(val *float64)
	CountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GpuDriverInstallationConfig() GoogleContainerClusterNodeConfigGuestAcceleratorGpuDriverInstallationConfigOutputReference
	GpuDriverInstallationConfigInput() *GoogleContainerClusterNodeConfigGuestAcceleratorGpuDriverInstallationConfig
	GpuPartitionSize() *string
	SetGpuPartitionSize(val *string)
	GpuPartitionSizeInput() *string
	GpuSharingConfig() GoogleContainerClusterNodeConfigGuestAcceleratorGpuSharingConfigOutputReference
	GpuSharingConfigInput() *GoogleContainerClusterNodeConfigGuestAcceleratorGpuSharingConfig
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutGpuDriverInstallationConfig(value *GoogleContainerClusterNodeConfigGuestAcceleratorGpuDriverInstallationConfig)
	PutGpuSharingConfig(value *GoogleContainerClusterNodeConfigGuestAcceleratorGpuSharingConfig)
	ResetGpuDriverInstallationConfig()
	ResetGpuPartitionSize()
	ResetGpuSharingConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference
type jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) CountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) GpuDriverInstallationConfig() GoogleContainerClusterNodeConfigGuestAcceleratorGpuDriverInstallationConfigOutputReference {
	var returns GoogleContainerClusterNodeConfigGuestAcceleratorGpuDriverInstallationConfigOutputReference
	_jsii_.Get(
		j,
		"gpuDriverInstallationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) GpuDriverInstallationConfigInput() *GoogleContainerClusterNodeConfigGuestAcceleratorGpuDriverInstallationConfig {
	var returns *GoogleContainerClusterNodeConfigGuestAcceleratorGpuDriverInstallationConfig
	_jsii_.Get(
		j,
		"gpuDriverInstallationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) GpuPartitionSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gpuPartitionSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) GpuPartitionSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gpuPartitionSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) GpuSharingConfig() GoogleContainerClusterNodeConfigGuestAcceleratorGpuSharingConfigOutputReference {
	var returns GoogleContainerClusterNodeConfigGuestAcceleratorGpuSharingConfigOutputReference
	_jsii_.Get(
		j,
		"gpuSharingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) GpuSharingConfigInput() *GoogleContainerClusterNodeConfigGuestAcceleratorGpuSharingConfig {
	var returns *GoogleContainerClusterNodeConfigGuestAcceleratorGpuSharingConfig
	_jsii_.Get(
		j,
		"gpuSharingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewGoogleContainerClusterNodeConfigGuestAcceleratorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleContainerClusterNodeConfigGuestAcceleratorOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContainerCluster.GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleContainerClusterNodeConfigGuestAcceleratorOutputReference_Override(g GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleContainerCluster.GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference)SetCount(val *float64) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference)SetGpuPartitionSize(val *string) {
	if err := j.validateSetGpuPartitionSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gpuPartitionSize",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) PutGpuDriverInstallationConfig(value *GoogleContainerClusterNodeConfigGuestAcceleratorGpuDriverInstallationConfig) {
	if err := g.validatePutGpuDriverInstallationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGpuDriverInstallationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) PutGpuSharingConfig(value *GoogleContainerClusterNodeConfigGuestAcceleratorGpuSharingConfig) {
	if err := g.validatePutGpuSharingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGpuSharingConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) ResetGpuDriverInstallationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGpuDriverInstallationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) ResetGpuPartitionSize() {
	_jsii_.InvokeVoid(
		g,
		"resetGpuPartitionSize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) ResetGpuSharingConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetGpuSharingConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleContainerClusterNodeConfigGuestAcceleratorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

