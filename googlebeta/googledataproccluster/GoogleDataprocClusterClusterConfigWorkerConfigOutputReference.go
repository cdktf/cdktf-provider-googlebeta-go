// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledataproccluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googledataproccluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDataprocClusterClusterConfigWorkerConfigOutputReference interface {
	cdktf.ComplexObject
	Accelerators() GoogleDataprocClusterClusterConfigWorkerConfigAcceleratorsList
	AcceleratorsInput() interface{}
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
	DiskConfig() GoogleDataprocClusterClusterConfigWorkerConfigDiskConfigOutputReference
	DiskConfigInput() *GoogleDataprocClusterClusterConfigWorkerConfigDiskConfig
	// Experimental.
	Fqn() *string
	ImageUri() *string
	SetImageUri(val *string)
	ImageUriInput() *string
	InstanceNames() *[]*string
	InternalValue() *GoogleDataprocClusterClusterConfigWorkerConfig
	SetInternalValue(val *GoogleDataprocClusterClusterConfigWorkerConfig)
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	MinCpuPlatform() *string
	SetMinCpuPlatform(val *string)
	MinCpuPlatformInput() *string
	MinNumInstances() *float64
	SetMinNumInstances(val *float64)
	MinNumInstancesInput() *float64
	NumInstances() *float64
	SetNumInstances(val *float64)
	NumInstancesInput() *float64
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
	PutAccelerators(value interface{})
	PutDiskConfig(value *GoogleDataprocClusterClusterConfigWorkerConfigDiskConfig)
	ResetAccelerators()
	ResetDiskConfig()
	ResetImageUri()
	ResetMachineType()
	ResetMinCpuPlatform()
	ResetMinNumInstances()
	ResetNumInstances()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocClusterClusterConfigWorkerConfigOutputReference
type jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) Accelerators() GoogleDataprocClusterClusterConfigWorkerConfigAcceleratorsList {
	var returns GoogleDataprocClusterClusterConfigWorkerConfigAcceleratorsList
	_jsii_.Get(
		j,
		"accelerators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) AcceleratorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceleratorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) DiskConfig() GoogleDataprocClusterClusterConfigWorkerConfigDiskConfigOutputReference {
	var returns GoogleDataprocClusterClusterConfigWorkerConfigDiskConfigOutputReference
	_jsii_.Get(
		j,
		"diskConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) DiskConfigInput() *GoogleDataprocClusterClusterConfigWorkerConfigDiskConfig {
	var returns *GoogleDataprocClusterClusterConfigWorkerConfigDiskConfig
	_jsii_.Get(
		j,
		"diskConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) ImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) ImageUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) InstanceNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) InternalValue() *GoogleDataprocClusterClusterConfigWorkerConfig {
	var returns *GoogleDataprocClusterClusterConfigWorkerConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) MinCpuPlatform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) MinCpuPlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minCpuPlatformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) MinNumInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) MinNumInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) NumInstances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) NumInstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataprocClusterClusterConfigWorkerConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataprocClusterClusterConfigWorkerConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocClusterClusterConfigWorkerConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocCluster.GoogleDataprocClusterClusterConfigWorkerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocClusterClusterConfigWorkerConfigOutputReference_Override(g GoogleDataprocClusterClusterConfigWorkerConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataprocCluster.GoogleDataprocClusterClusterConfigWorkerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference)SetImageUri(val *string) {
	if err := j.validateSetImageUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageUri",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference)SetInternalValue(val *GoogleDataprocClusterClusterConfigWorkerConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference)SetMinCpuPlatform(val *string) {
	if err := j.validateSetMinCpuPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCpuPlatform",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference)SetMinNumInstances(val *float64) {
	if err := j.validateSetMinNumInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNumInstances",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference)SetNumInstances(val *float64) {
	if err := j.validateSetNumInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numInstances",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) PutAccelerators(value interface{}) {
	if err := g.validatePutAcceleratorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAccelerators",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) PutDiskConfig(value *GoogleDataprocClusterClusterConfigWorkerConfigDiskConfig) {
	if err := g.validatePutDiskConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDiskConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) ResetAccelerators() {
	_jsii_.InvokeVoid(
		g,
		"resetAccelerators",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) ResetDiskConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetDiskConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) ResetImageUri() {
	_jsii_.InvokeVoid(
		g,
		"resetImageUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) ResetMachineType() {
	_jsii_.InvokeVoid(
		g,
		"resetMachineType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) ResetMinCpuPlatform() {
	_jsii_.InvokeVoid(
		g,
		"resetMinCpuPlatform",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) ResetMinNumInstances() {
	_jsii_.InvokeVoid(
		g,
		"resetMinNumInstances",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) ResetNumInstances() {
	_jsii_.InvokeVoid(
		g,
		"resetNumInstances",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataprocClusterClusterConfigWorkerConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

