// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlestorageinsightsreportconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googlestorageinsightsreportconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference interface {
	cdktf.ComplexObject
	Bucket() *string
	SetBucket(val *string)
	BucketInput() *string
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
	DestinationPath() *string
	SetDestinationPath(val *string)
	DestinationPathInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptions
	SetInternalValue(val *GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptions)
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
	ResetDestinationPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference
type jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) DestinationPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) DestinationPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) InternalValue() *GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptions {
	var returns *GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleStorageInsightsReportConfig.GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference_Override(g GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleStorageInsightsReportConfig.GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference)SetDestinationPath(val *string) {
	if err := j.validateSetDestinationPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationPath",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference)SetInternalValue(val *GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) ResetDestinationPath() {
	_jsii_.InvokeVoid(
		g,
		"resetDestinationPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleStorageInsightsReportConfigObjectMetadataReportOptionsStorageDestinationOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

