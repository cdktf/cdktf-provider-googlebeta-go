// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlehealthcarepipelinejob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googlehealthcarepipelinejob/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference interface {
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
	FhirStoreDestination() *string
	SetFhirStoreDestination(val *string)
	FhirStoreDestinationInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleHealthcarePipelineJobReconciliationPipelineJob
	SetInternalValue(val *GoogleHealthcarePipelineJobReconciliationPipelineJob)
	MatchingUriPrefix() *string
	SetMatchingUriPrefix(val *string)
	MatchingUriPrefixInput() *string
	MergeConfig() GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigOutputReference
	MergeConfigInput() *GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfig
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
	PutMergeConfig(value *GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfig)
	ResetFhirStoreDestination()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference
type jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) FhirStoreDestination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fhirStoreDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) FhirStoreDestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fhirStoreDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) InternalValue() *GoogleHealthcarePipelineJobReconciliationPipelineJob {
	var returns *GoogleHealthcarePipelineJobReconciliationPipelineJob
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) MatchingUriPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchingUriPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) MatchingUriPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchingUriPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) MergeConfig() GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigOutputReference {
	var returns GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigOutputReference
	_jsii_.Get(
		j,
		"mergeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) MergeConfigInput() *GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfig {
	var returns *GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfig
	_jsii_.Get(
		j,
		"mergeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleHealthcarePipelineJobReconciliationPipelineJobOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleHealthcarePipelineJob.GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference_Override(g GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleHealthcarePipelineJob.GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference)SetFhirStoreDestination(val *string) {
	if err := j.validateSetFhirStoreDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fhirStoreDestination",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference)SetInternalValue(val *GoogleHealthcarePipelineJobReconciliationPipelineJob) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference)SetMatchingUriPrefix(val *string) {
	if err := j.validateSetMatchingUriPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchingUriPrefix",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) PutMergeConfig(value *GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfig) {
	if err := g.validatePutMergeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMergeConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) ResetFhirStoreDestination() {
	_jsii_.InvokeVoid(
		g,
		"resetFhirStoreDestination",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

