// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleprivatecacertificate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googleprivatecacertificate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference interface {
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
	Critical() interface{}
	SetCritical(val interface{})
	CriticalInput() interface{}
	ExcludedDnsNames() *[]*string
	SetExcludedDnsNames(val *[]*string)
	ExcludedDnsNamesInput() *[]*string
	ExcludedEmailAddresses() *[]*string
	SetExcludedEmailAddresses(val *[]*string)
	ExcludedEmailAddressesInput() *[]*string
	ExcludedIpRanges() *[]*string
	SetExcludedIpRanges(val *[]*string)
	ExcludedIpRangesInput() *[]*string
	ExcludedUris() *[]*string
	SetExcludedUris(val *[]*string)
	ExcludedUrisInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GooglePrivatecaCertificateConfigX509ConfigNameConstraints
	SetInternalValue(val *GooglePrivatecaCertificateConfigX509ConfigNameConstraints)
	PermittedDnsNames() *[]*string
	SetPermittedDnsNames(val *[]*string)
	PermittedDnsNamesInput() *[]*string
	PermittedEmailAddresses() *[]*string
	SetPermittedEmailAddresses(val *[]*string)
	PermittedEmailAddressesInput() *[]*string
	PermittedIpRanges() *[]*string
	SetPermittedIpRanges(val *[]*string)
	PermittedIpRangesInput() *[]*string
	PermittedUris() *[]*string
	SetPermittedUris(val *[]*string)
	PermittedUrisInput() *[]*string
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
	ResetExcludedDnsNames()
	ResetExcludedEmailAddresses()
	ResetExcludedIpRanges()
	ResetExcludedUris()
	ResetPermittedDnsNames()
	ResetPermittedEmailAddresses()
	ResetPermittedIpRanges()
	ResetPermittedUris()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference
type jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) Critical() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"critical",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) CriticalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"criticalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) ExcludedDnsNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedDnsNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) ExcludedDnsNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedDnsNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) ExcludedEmailAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedEmailAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) ExcludedEmailAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedEmailAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) ExcludedIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) ExcludedIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) ExcludedUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) ExcludedUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) InternalValue() *GooglePrivatecaCertificateConfigX509ConfigNameConstraints {
	var returns *GooglePrivatecaCertificateConfigX509ConfigNameConstraints
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) PermittedDnsNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedDnsNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) PermittedDnsNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedDnsNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) PermittedEmailAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedEmailAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) PermittedEmailAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedEmailAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) PermittedIpRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedIpRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) PermittedIpRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedIpRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) PermittedUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) PermittedUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePrivatecaCertificate.GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference_Override(g GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePrivatecaCertificate.GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference)SetCritical(val interface{}) {
	if err := j.validateSetCriticalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"critical",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference)SetExcludedDnsNames(val *[]*string) {
	if err := j.validateSetExcludedDnsNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedDnsNames",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference)SetExcludedEmailAddresses(val *[]*string) {
	if err := j.validateSetExcludedEmailAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedEmailAddresses",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference)SetExcludedIpRanges(val *[]*string) {
	if err := j.validateSetExcludedIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedIpRanges",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference)SetExcludedUris(val *[]*string) {
	if err := j.validateSetExcludedUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedUris",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference)SetInternalValue(val *GooglePrivatecaCertificateConfigX509ConfigNameConstraints) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference)SetPermittedDnsNames(val *[]*string) {
	if err := j.validateSetPermittedDnsNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permittedDnsNames",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference)SetPermittedEmailAddresses(val *[]*string) {
	if err := j.validateSetPermittedEmailAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permittedEmailAddresses",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference)SetPermittedIpRanges(val *[]*string) {
	if err := j.validateSetPermittedIpRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permittedIpRanges",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference)SetPermittedUris(val *[]*string) {
	if err := j.validateSetPermittedUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permittedUris",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) ResetExcludedDnsNames() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludedDnsNames",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) ResetExcludedEmailAddresses() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludedEmailAddresses",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) ResetExcludedIpRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludedIpRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) ResetExcludedUris() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludedUris",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) ResetPermittedDnsNames() {
	_jsii_.InvokeVoid(
		g,
		"resetPermittedDnsNames",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) ResetPermittedEmailAddresses() {
	_jsii_.InvokeVoid(
		g,
		"resetPermittedEmailAddresses",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) ResetPermittedIpRanges() {
	_jsii_.InvokeVoid(
		g,
		"resetPermittedIpRanges",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) ResetPermittedUris() {
	_jsii_.InvokeVoid(
		g,
		"resetPermittedUris",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigX509ConfigNameConstraintsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

