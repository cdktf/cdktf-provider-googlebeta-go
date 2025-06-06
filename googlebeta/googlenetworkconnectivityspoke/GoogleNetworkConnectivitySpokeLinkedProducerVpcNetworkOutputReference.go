// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlenetworkconnectivityspoke

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googlenetworkconnectivityspoke/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference interface {
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
	ExcludeExportRanges() *[]*string
	SetExcludeExportRanges(val *[]*string)
	ExcludeExportRangesInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludeExportRanges() *[]*string
	SetIncludeExportRanges(val *[]*string)
	IncludeExportRangesInput() *[]*string
	InternalValue() *GoogleNetworkConnectivitySpokeLinkedProducerVpcNetwork
	SetInternalValue(val *GoogleNetworkConnectivitySpokeLinkedProducerVpcNetwork)
	Network() *string
	SetNetwork(val *string)
	NetworkInput() *string
	Peering() *string
	SetPeering(val *string)
	PeeringInput() *string
	ProducerNetwork() *string
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
	ResetExcludeExportRanges()
	ResetIncludeExportRanges()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference
type jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) ExcludeExportRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeExportRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) ExcludeExportRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeExportRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) IncludeExportRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeExportRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) IncludeExportRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includeExportRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) InternalValue() *GoogleNetworkConnectivitySpokeLinkedProducerVpcNetwork {
	var returns *GoogleNetworkConnectivitySpokeLinkedProducerVpcNetwork
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) Peering() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) PeeringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peeringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) ProducerNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"producerNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkConnectivitySpoke.GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference_Override(g GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkConnectivitySpoke.GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference)SetExcludeExportRanges(val *[]*string) {
	if err := j.validateSetExcludeExportRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeExportRanges",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference)SetIncludeExportRanges(val *[]*string) {
	if err := j.validateSetIncludeExportRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeExportRanges",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference)SetInternalValue(val *GoogleNetworkConnectivitySpokeLinkedProducerVpcNetwork) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference)SetPeering(val *string) {
	if err := j.validateSetPeeringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peering",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) ResetExcludeExportRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludeExportRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) ResetIncludeExportRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeExportRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetworkConnectivitySpokeLinkedProducerVpcNetworkOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

