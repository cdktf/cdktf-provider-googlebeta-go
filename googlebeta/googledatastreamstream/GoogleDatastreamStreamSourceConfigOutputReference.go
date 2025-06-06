// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledatastreamstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googledatastreamstream/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDatastreamStreamSourceConfigOutputReference interface {
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
	InternalValue() *GoogleDatastreamStreamSourceConfig
	SetInternalValue(val *GoogleDatastreamStreamSourceConfig)
	MysqlSourceConfig() GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference
	MysqlSourceConfigInput() *GoogleDatastreamStreamSourceConfigMysqlSourceConfig
	OracleSourceConfig() GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference
	OracleSourceConfigInput() *GoogleDatastreamStreamSourceConfigOracleSourceConfig
	PostgresqlSourceConfig() GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigOutputReference
	PostgresqlSourceConfigInput() *GoogleDatastreamStreamSourceConfigPostgresqlSourceConfig
	SalesforceSourceConfig() GoogleDatastreamStreamSourceConfigSalesforceSourceConfigOutputReference
	SalesforceSourceConfigInput() *GoogleDatastreamStreamSourceConfigSalesforceSourceConfig
	SourceConnectionProfile() *string
	SetSourceConnectionProfile(val *string)
	SourceConnectionProfileInput() *string
	SqlServerSourceConfig() GoogleDatastreamStreamSourceConfigSqlServerSourceConfigOutputReference
	SqlServerSourceConfigInput() *GoogleDatastreamStreamSourceConfigSqlServerSourceConfig
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
	PutMysqlSourceConfig(value *GoogleDatastreamStreamSourceConfigMysqlSourceConfig)
	PutOracleSourceConfig(value *GoogleDatastreamStreamSourceConfigOracleSourceConfig)
	PutPostgresqlSourceConfig(value *GoogleDatastreamStreamSourceConfigPostgresqlSourceConfig)
	PutSalesforceSourceConfig(value *GoogleDatastreamStreamSourceConfigSalesforceSourceConfig)
	PutSqlServerSourceConfig(value *GoogleDatastreamStreamSourceConfigSqlServerSourceConfig)
	ResetMysqlSourceConfig()
	ResetOracleSourceConfig()
	ResetPostgresqlSourceConfig()
	ResetSalesforceSourceConfig()
	ResetSqlServerSourceConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDatastreamStreamSourceConfigOutputReference
type jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) InternalValue() *GoogleDatastreamStreamSourceConfig {
	var returns *GoogleDatastreamStreamSourceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) MysqlSourceConfig() GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference {
	var returns GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference
	_jsii_.Get(
		j,
		"mysqlSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) MysqlSourceConfigInput() *GoogleDatastreamStreamSourceConfigMysqlSourceConfig {
	var returns *GoogleDatastreamStreamSourceConfigMysqlSourceConfig
	_jsii_.Get(
		j,
		"mysqlSourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) OracleSourceConfig() GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference {
	var returns GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference
	_jsii_.Get(
		j,
		"oracleSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) OracleSourceConfigInput() *GoogleDatastreamStreamSourceConfigOracleSourceConfig {
	var returns *GoogleDatastreamStreamSourceConfigOracleSourceConfig
	_jsii_.Get(
		j,
		"oracleSourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) PostgresqlSourceConfig() GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigOutputReference {
	var returns GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigOutputReference
	_jsii_.Get(
		j,
		"postgresqlSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) PostgresqlSourceConfigInput() *GoogleDatastreamStreamSourceConfigPostgresqlSourceConfig {
	var returns *GoogleDatastreamStreamSourceConfigPostgresqlSourceConfig
	_jsii_.Get(
		j,
		"postgresqlSourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) SalesforceSourceConfig() GoogleDatastreamStreamSourceConfigSalesforceSourceConfigOutputReference {
	var returns GoogleDatastreamStreamSourceConfigSalesforceSourceConfigOutputReference
	_jsii_.Get(
		j,
		"salesforceSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) SalesforceSourceConfigInput() *GoogleDatastreamStreamSourceConfigSalesforceSourceConfig {
	var returns *GoogleDatastreamStreamSourceConfigSalesforceSourceConfig
	_jsii_.Get(
		j,
		"salesforceSourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) SourceConnectionProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceConnectionProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) SourceConnectionProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceConnectionProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) SqlServerSourceConfig() GoogleDatastreamStreamSourceConfigSqlServerSourceConfigOutputReference {
	var returns GoogleDatastreamStreamSourceConfigSqlServerSourceConfigOutputReference
	_jsii_.Get(
		j,
		"sqlServerSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) SqlServerSourceConfigInput() *GoogleDatastreamStreamSourceConfigSqlServerSourceConfig {
	var returns *GoogleDatastreamStreamSourceConfigSqlServerSourceConfig
	_jsii_.Get(
		j,
		"sqlServerSourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDatastreamStreamSourceConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDatastreamStreamSourceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDatastreamStreamSourceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDatastreamStreamSourceConfigOutputReference_Override(g GoogleDatastreamStreamSourceConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference)SetInternalValue(val *GoogleDatastreamStreamSourceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference)SetSourceConnectionProfile(val *string) {
	if err := j.validateSetSourceConnectionProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceConnectionProfile",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) PutMysqlSourceConfig(value *GoogleDatastreamStreamSourceConfigMysqlSourceConfig) {
	if err := g.validatePutMysqlSourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMysqlSourceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) PutOracleSourceConfig(value *GoogleDatastreamStreamSourceConfigOracleSourceConfig) {
	if err := g.validatePutOracleSourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOracleSourceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) PutPostgresqlSourceConfig(value *GoogleDatastreamStreamSourceConfigPostgresqlSourceConfig) {
	if err := g.validatePutPostgresqlSourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPostgresqlSourceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) PutSalesforceSourceConfig(value *GoogleDatastreamStreamSourceConfigSalesforceSourceConfig) {
	if err := g.validatePutSalesforceSourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSalesforceSourceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) PutSqlServerSourceConfig(value *GoogleDatastreamStreamSourceConfigSqlServerSourceConfig) {
	if err := g.validatePutSqlServerSourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSqlServerSourceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) ResetMysqlSourceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetMysqlSourceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) ResetOracleSourceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetOracleSourceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) ResetPostgresqlSourceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPostgresqlSourceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) ResetSalesforceSourceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSalesforceSourceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) ResetSqlServerSourceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSqlServerSourceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

