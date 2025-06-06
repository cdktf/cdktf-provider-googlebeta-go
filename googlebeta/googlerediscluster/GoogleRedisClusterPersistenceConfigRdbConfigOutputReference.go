// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlerediscluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googlerediscluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleRedisClusterPersistenceConfigRdbConfigOutputReference interface {
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
	InternalValue() *GoogleRedisClusterPersistenceConfigRdbConfig
	SetInternalValue(val *GoogleRedisClusterPersistenceConfigRdbConfig)
	RdbSnapshotPeriod() *string
	SetRdbSnapshotPeriod(val *string)
	RdbSnapshotPeriodInput() *string
	RdbSnapshotStartTime() *string
	SetRdbSnapshotStartTime(val *string)
	RdbSnapshotStartTimeInput() *string
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
	ResetRdbSnapshotPeriod()
	ResetRdbSnapshotStartTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleRedisClusterPersistenceConfigRdbConfigOutputReference
type jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) InternalValue() *GoogleRedisClusterPersistenceConfigRdbConfig {
	var returns *GoogleRedisClusterPersistenceConfigRdbConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) RdbSnapshotPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdbSnapshotPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) RdbSnapshotPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdbSnapshotPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) RdbSnapshotStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdbSnapshotStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) RdbSnapshotStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rdbSnapshotStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleRedisClusterPersistenceConfigRdbConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleRedisClusterPersistenceConfigRdbConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleRedisClusterPersistenceConfigRdbConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleRedisCluster.GoogleRedisClusterPersistenceConfigRdbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleRedisClusterPersistenceConfigRdbConfigOutputReference_Override(g GoogleRedisClusterPersistenceConfigRdbConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleRedisCluster.GoogleRedisClusterPersistenceConfigRdbConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference)SetInternalValue(val *GoogleRedisClusterPersistenceConfigRdbConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference)SetRdbSnapshotPeriod(val *string) {
	if err := j.validateSetRdbSnapshotPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdbSnapshotPeriod",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference)SetRdbSnapshotStartTime(val *string) {
	if err := j.validateSetRdbSnapshotStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rdbSnapshotStartTime",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) ResetRdbSnapshotPeriod() {
	_jsii_.InvokeVoid(
		g,
		"resetRdbSnapshotPeriod",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) ResetRdbSnapshotStartTime() {
	_jsii_.InvokeVoid(
		g,
		"resetRdbSnapshotStartTime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleRedisClusterPersistenceConfigRdbConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

