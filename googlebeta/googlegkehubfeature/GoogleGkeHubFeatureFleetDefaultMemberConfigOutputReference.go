// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlegkehubfeature

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googlegkehubfeature/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference interface {
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
	Configmanagement() GoogleGkeHubFeatureFleetDefaultMemberConfigConfigmanagementOutputReference
	ConfigmanagementInput() *GoogleGkeHubFeatureFleetDefaultMemberConfigConfigmanagement
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleGkeHubFeatureFleetDefaultMemberConfig
	SetInternalValue(val *GoogleGkeHubFeatureFleetDefaultMemberConfig)
	Mesh() GoogleGkeHubFeatureFleetDefaultMemberConfigMeshOutputReference
	MeshInput() *GoogleGkeHubFeatureFleetDefaultMemberConfigMesh
	Policycontroller() GoogleGkeHubFeatureFleetDefaultMemberConfigPolicycontrollerOutputReference
	PolicycontrollerInput() *GoogleGkeHubFeatureFleetDefaultMemberConfigPolicycontroller
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
	PutConfigmanagement(value *GoogleGkeHubFeatureFleetDefaultMemberConfigConfigmanagement)
	PutMesh(value *GoogleGkeHubFeatureFleetDefaultMemberConfigMesh)
	PutPolicycontroller(value *GoogleGkeHubFeatureFleetDefaultMemberConfigPolicycontroller)
	ResetConfigmanagement()
	ResetMesh()
	ResetPolicycontroller()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference
type jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) Configmanagement() GoogleGkeHubFeatureFleetDefaultMemberConfigConfigmanagementOutputReference {
	var returns GoogleGkeHubFeatureFleetDefaultMemberConfigConfigmanagementOutputReference
	_jsii_.Get(
		j,
		"configmanagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) ConfigmanagementInput() *GoogleGkeHubFeatureFleetDefaultMemberConfigConfigmanagement {
	var returns *GoogleGkeHubFeatureFleetDefaultMemberConfigConfigmanagement
	_jsii_.Get(
		j,
		"configmanagementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) InternalValue() *GoogleGkeHubFeatureFleetDefaultMemberConfig {
	var returns *GoogleGkeHubFeatureFleetDefaultMemberConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) Mesh() GoogleGkeHubFeatureFleetDefaultMemberConfigMeshOutputReference {
	var returns GoogleGkeHubFeatureFleetDefaultMemberConfigMeshOutputReference
	_jsii_.Get(
		j,
		"mesh",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) MeshInput() *GoogleGkeHubFeatureFleetDefaultMemberConfigMesh {
	var returns *GoogleGkeHubFeatureFleetDefaultMemberConfigMesh
	_jsii_.Get(
		j,
		"meshInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) Policycontroller() GoogleGkeHubFeatureFleetDefaultMemberConfigPolicycontrollerOutputReference {
	var returns GoogleGkeHubFeatureFleetDefaultMemberConfigPolicycontrollerOutputReference
	_jsii_.Get(
		j,
		"policycontroller",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) PolicycontrollerInput() *GoogleGkeHubFeatureFleetDefaultMemberConfigPolicycontroller {
	var returns *GoogleGkeHubFeatureFleetDefaultMemberConfigPolicycontroller
	_jsii_.Get(
		j,
		"policycontrollerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGkeHubFeatureFleetDefaultMemberConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeHubFeature.GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference_Override(g GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleGkeHubFeature.GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference)SetInternalValue(val *GoogleGkeHubFeatureFleetDefaultMemberConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) PutConfigmanagement(value *GoogleGkeHubFeatureFleetDefaultMemberConfigConfigmanagement) {
	if err := g.validatePutConfigmanagementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConfigmanagement",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) PutMesh(value *GoogleGkeHubFeatureFleetDefaultMemberConfigMesh) {
	if err := g.validatePutMeshParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMesh",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) PutPolicycontroller(value *GoogleGkeHubFeatureFleetDefaultMemberConfigPolicycontroller) {
	if err := g.validatePutPolicycontrollerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPolicycontroller",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) ResetConfigmanagement() {
	_jsii_.InvokeVoid(
		g,
		"resetConfigmanagement",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) ResetMesh() {
	_jsii_.InvokeVoid(
		g,
		"resetMesh",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) ResetPolicycontroller() {
	_jsii_.InvokeVoid(
		g,
		"resetPolicycontroller",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleGkeHubFeatureFleetDefaultMemberConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

