// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlesqldatabaseinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googlesqldatabaseinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference interface {
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
	ConsumerNetwork() *string
	SetConsumerNetwork(val *string)
	ConsumerNetworkInput() *string
	ConsumerServiceProjectId() *string
	SetConsumerServiceProjectId(val *string)
	ConsumerServiceProjectIdInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
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
	ResetConsumerServiceProjectId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference
type jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) ConsumerNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) ConsumerNetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) ConsumerServiceProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerServiceProjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) ConsumerServiceProjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumerServiceProjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference_Override(g GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference)SetConsumerNetwork(val *string) {
	if err := j.validateSetConsumerNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerNetwork",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference)SetConsumerServiceProjectId(val *string) {
	if err := j.validateSetConsumerServiceProjectIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumerServiceProjectId",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) ResetConsumerServiceProjectId() {
	_jsii_.InvokeVoid(
		g,
		"resetConsumerServiceProjectId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsIpConfigurationPscConfigPscAutoConnectionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

