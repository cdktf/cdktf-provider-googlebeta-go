// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledatastreamstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googledatastreamstream/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference interface {
	cdktf.ComplexObject
	BinaryLogPosition() GoogleDatastreamStreamSourceConfigMysqlSourceConfigBinaryLogPositionOutputReference
	BinaryLogPositionInput() *GoogleDatastreamStreamSourceConfigMysqlSourceConfigBinaryLogPosition
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
	ExcludeObjects() GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsOutputReference
	ExcludeObjectsInput() *GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjects
	// Experimental.
	Fqn() *string
	Gtid() GoogleDatastreamStreamSourceConfigMysqlSourceConfigGtidOutputReference
	GtidInput() *GoogleDatastreamStreamSourceConfigMysqlSourceConfigGtid
	IncludeObjects() GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsOutputReference
	IncludeObjectsInput() *GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjects
	InternalValue() *GoogleDatastreamStreamSourceConfigMysqlSourceConfig
	SetInternalValue(val *GoogleDatastreamStreamSourceConfigMysqlSourceConfig)
	MaxConcurrentBackfillTasks() *float64
	SetMaxConcurrentBackfillTasks(val *float64)
	MaxConcurrentBackfillTasksInput() *float64
	MaxConcurrentCdcTasks() *float64
	SetMaxConcurrentCdcTasks(val *float64)
	MaxConcurrentCdcTasksInput() *float64
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
	PutBinaryLogPosition(value *GoogleDatastreamStreamSourceConfigMysqlSourceConfigBinaryLogPosition)
	PutExcludeObjects(value *GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjects)
	PutGtid(value *GoogleDatastreamStreamSourceConfigMysqlSourceConfigGtid)
	PutIncludeObjects(value *GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjects)
	ResetBinaryLogPosition()
	ResetExcludeObjects()
	ResetGtid()
	ResetIncludeObjects()
	ResetMaxConcurrentBackfillTasks()
	ResetMaxConcurrentCdcTasks()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference
type jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) BinaryLogPosition() GoogleDatastreamStreamSourceConfigMysqlSourceConfigBinaryLogPositionOutputReference {
	var returns GoogleDatastreamStreamSourceConfigMysqlSourceConfigBinaryLogPositionOutputReference
	_jsii_.Get(
		j,
		"binaryLogPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) BinaryLogPositionInput() *GoogleDatastreamStreamSourceConfigMysqlSourceConfigBinaryLogPosition {
	var returns *GoogleDatastreamStreamSourceConfigMysqlSourceConfigBinaryLogPosition
	_jsii_.Get(
		j,
		"binaryLogPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) ExcludeObjects() GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsOutputReference {
	var returns GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsOutputReference
	_jsii_.Get(
		j,
		"excludeObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) ExcludeObjectsInput() *GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjects {
	var returns *GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjects
	_jsii_.Get(
		j,
		"excludeObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) Gtid() GoogleDatastreamStreamSourceConfigMysqlSourceConfigGtidOutputReference {
	var returns GoogleDatastreamStreamSourceConfigMysqlSourceConfigGtidOutputReference
	_jsii_.Get(
		j,
		"gtid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) GtidInput() *GoogleDatastreamStreamSourceConfigMysqlSourceConfigGtid {
	var returns *GoogleDatastreamStreamSourceConfigMysqlSourceConfigGtid
	_jsii_.Get(
		j,
		"gtidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) IncludeObjects() GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsOutputReference {
	var returns GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsOutputReference
	_jsii_.Get(
		j,
		"includeObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) IncludeObjectsInput() *GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjects {
	var returns *GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjects
	_jsii_.Get(
		j,
		"includeObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) InternalValue() *GoogleDatastreamStreamSourceConfigMysqlSourceConfig {
	var returns *GoogleDatastreamStreamSourceConfigMysqlSourceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) MaxConcurrentBackfillTasks() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentBackfillTasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) MaxConcurrentBackfillTasksInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentBackfillTasksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) MaxConcurrentCdcTasks() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentCdcTasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) MaxConcurrentCdcTasksInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentCdcTasksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference_Override(g GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference)SetInternalValue(val *GoogleDatastreamStreamSourceConfigMysqlSourceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference)SetMaxConcurrentBackfillTasks(val *float64) {
	if err := j.validateSetMaxConcurrentBackfillTasksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentBackfillTasks",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference)SetMaxConcurrentCdcTasks(val *float64) {
	if err := j.validateSetMaxConcurrentCdcTasksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentCdcTasks",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) PutBinaryLogPosition(value *GoogleDatastreamStreamSourceConfigMysqlSourceConfigBinaryLogPosition) {
	if err := g.validatePutBinaryLogPositionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBinaryLogPosition",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) PutExcludeObjects(value *GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjects) {
	if err := g.validatePutExcludeObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExcludeObjects",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) PutGtid(value *GoogleDatastreamStreamSourceConfigMysqlSourceConfigGtid) {
	if err := g.validatePutGtidParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGtid",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) PutIncludeObjects(value *GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjects) {
	if err := g.validatePutIncludeObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIncludeObjects",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) ResetBinaryLogPosition() {
	_jsii_.InvokeVoid(
		g,
		"resetBinaryLogPosition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) ResetExcludeObjects() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludeObjects",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) ResetGtid() {
	_jsii_.InvokeVoid(
		g,
		"resetGtid",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) ResetIncludeObjects() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeObjects",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) ResetMaxConcurrentBackfillTasks() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxConcurrentBackfillTasks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) ResetMaxConcurrentCdcTasks() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxConcurrentCdcTasks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

