// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlenetworkconnectivityspoke

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googlenetworkconnectivityspoke/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference interface {
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
	IncludeImportRanges() *[]*string
	SetIncludeImportRanges(val *[]*string)
	IncludeImportRangesInput() *[]*string
	InternalValue() *GoogleNetworkConnectivitySpokeLinkedInterconnectAttachments
	SetInternalValue(val *GoogleNetworkConnectivitySpokeLinkedInterconnectAttachments)
	SiteToSiteDataTransfer() interface{}
	SetSiteToSiteDataTransfer(val interface{})
	SiteToSiteDataTransferInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uris() *[]*string
	SetUris(val *[]*string)
	UrisInput() *[]*string
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
	ResetIncludeImportRanges()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference
type jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) IncludeImportRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeImportRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) IncludeImportRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeImportRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) InternalValue() *GoogleNetworkConnectivitySpokeLinkedInterconnectAttachments {
	var returns *GoogleNetworkConnectivitySpokeLinkedInterconnectAttachments
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) SiteToSiteDataTransfer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"siteToSiteDataTransfer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) SiteToSiteDataTransferInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"siteToSiteDataTransferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) Uris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"uris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) UrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"urisInput",
		&returns,
	)
	return returns
}


func NewGoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkConnectivitySpoke.GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference_Override(g GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkConnectivitySpoke.GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference)SetIncludeImportRanges(val *[]*string) {
	if err := j.validateSetIncludeImportRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeImportRanges",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference)SetInternalValue(val *GoogleNetworkConnectivitySpokeLinkedInterconnectAttachments) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference)SetSiteToSiteDataTransfer(val interface{}) {
	if err := j.validateSetSiteToSiteDataTransferParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"siteToSiteDataTransfer",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference)SetUris(val *[]*string) {
	if err := j.validateSetUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uris",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) ResetIncludeImportRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeImportRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedInterconnectAttachmentsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

