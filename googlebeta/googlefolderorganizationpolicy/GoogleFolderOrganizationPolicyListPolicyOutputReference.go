// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlefolderorganizationpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googlefolderorganizationpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleFolderOrganizationPolicyListPolicyOutputReference interface {
	cdktf.ComplexObject
	Allow() GoogleFolderOrganizationPolicyListPolicyAllowOutputReference
	AllowInput() *GoogleFolderOrganizationPolicyListPolicyAllow
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
	Deny() GoogleFolderOrganizationPolicyListPolicyDenyOutputReference
	DenyInput() *GoogleFolderOrganizationPolicyListPolicyDeny
	// Experimental.
	Fqn() *string
	InheritFromParent() interface{}
	SetInheritFromParent(val interface{})
	InheritFromParentInput() interface{}
	InternalValue() *GoogleFolderOrganizationPolicyListPolicy
	SetInternalValue(val *GoogleFolderOrganizationPolicyListPolicy)
	SuggestedValue() *string
	SetSuggestedValue(val *string)
	SuggestedValueInput() *string
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
	PutAllow(value *GoogleFolderOrganizationPolicyListPolicyAllow)
	PutDeny(value *GoogleFolderOrganizationPolicyListPolicyDeny)
	ResetAllow()
	ResetDeny()
	ResetInheritFromParent()
	ResetSuggestedValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleFolderOrganizationPolicyListPolicyOutputReference
type jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) Allow() GoogleFolderOrganizationPolicyListPolicyAllowOutputReference {
	var returns GoogleFolderOrganizationPolicyListPolicyAllowOutputReference
	_jsii_.Get(
		j,
		"allow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) AllowInput() *GoogleFolderOrganizationPolicyListPolicyAllow {
	var returns *GoogleFolderOrganizationPolicyListPolicyAllow
	_jsii_.Get(
		j,
		"allowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) Deny() GoogleFolderOrganizationPolicyListPolicyDenyOutputReference {
	var returns GoogleFolderOrganizationPolicyListPolicyDenyOutputReference
	_jsii_.Get(
		j,
		"deny",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) DenyInput() *GoogleFolderOrganizationPolicyListPolicyDeny {
	var returns *GoogleFolderOrganizationPolicyListPolicyDeny
	_jsii_.Get(
		j,
		"denyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) InheritFromParent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inheritFromParent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) InheritFromParentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inheritFromParentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) InternalValue() *GoogleFolderOrganizationPolicyListPolicy {
	var returns *GoogleFolderOrganizationPolicyListPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) SuggestedValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suggestedValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) SuggestedValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"suggestedValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleFolderOrganizationPolicyListPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleFolderOrganizationPolicyListPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleFolderOrganizationPolicyListPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleFolderOrganizationPolicy.GoogleFolderOrganizationPolicyListPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleFolderOrganizationPolicyListPolicyOutputReference_Override(g GoogleFolderOrganizationPolicyListPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleFolderOrganizationPolicy.GoogleFolderOrganizationPolicyListPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference)SetInheritFromParent(val interface{}) {
	if err := j.validateSetInheritFromParentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inheritFromParent",
		val,
	)
}

func (j *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference)SetInternalValue(val *GoogleFolderOrganizationPolicyListPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference)SetSuggestedValue(val *string) {
	if err := j.validateSetSuggestedValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suggestedValue",
		val,
	)
}

func (j *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) PutAllow(value *GoogleFolderOrganizationPolicyListPolicyAllow) {
	if err := g.validatePutAllowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAllow",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) PutDeny(value *GoogleFolderOrganizationPolicyListPolicyDeny) {
	if err := g.validatePutDenyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeny",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) ResetAllow() {
	_jsii_.InvokeVoid(
		g,
		"resetAllow",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) ResetDeny() {
	_jsii_.InvokeVoid(
		g,
		"resetDeny",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) ResetInheritFromParent() {
	_jsii_.InvokeVoid(
		g,
		"resetInheritFromParent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) ResetSuggestedValue() {
	_jsii_.InvokeVoid(
		g,
		"resetSuggestedValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleFolderOrganizationPolicyListPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

