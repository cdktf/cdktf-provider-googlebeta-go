// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleprivatecacertificate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googleprivatecacertificate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GooglePrivatecaCertificateConfigAOutputReference interface {
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
	InternalValue() *GooglePrivatecaCertificateConfigA
	SetInternalValue(val *GooglePrivatecaCertificateConfigA)
	PublicKey() GooglePrivatecaCertificateConfigPublicKeyOutputReference
	PublicKeyInput() *GooglePrivatecaCertificateConfigPublicKey
	SubjectConfig() GooglePrivatecaCertificateConfigSubjectConfigOutputReference
	SubjectConfigInput() *GooglePrivatecaCertificateConfigSubjectConfig
	SubjectKeyId() GooglePrivatecaCertificateConfigSubjectKeyIdOutputReference
	SubjectKeyIdInput() *GooglePrivatecaCertificateConfigSubjectKeyId
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	X509Config() GooglePrivatecaCertificateConfigX509ConfigOutputReference
	X509ConfigInput() *GooglePrivatecaCertificateConfigX509Config
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
	PutPublicKey(value *GooglePrivatecaCertificateConfigPublicKey)
	PutSubjectConfig(value *GooglePrivatecaCertificateConfigSubjectConfig)
	PutSubjectKeyId(value *GooglePrivatecaCertificateConfigSubjectKeyId)
	PutX509Config(value *GooglePrivatecaCertificateConfigX509Config)
	ResetSubjectKeyId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePrivatecaCertificateConfigAOutputReference
type jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) InternalValue() *GooglePrivatecaCertificateConfigA {
	var returns *GooglePrivatecaCertificateConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) PublicKey() GooglePrivatecaCertificateConfigPublicKeyOutputReference {
	var returns GooglePrivatecaCertificateConfigPublicKeyOutputReference
	_jsii_.Get(
		j,
		"publicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) PublicKeyInput() *GooglePrivatecaCertificateConfigPublicKey {
	var returns *GooglePrivatecaCertificateConfigPublicKey
	_jsii_.Get(
		j,
		"publicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) SubjectConfig() GooglePrivatecaCertificateConfigSubjectConfigOutputReference {
	var returns GooglePrivatecaCertificateConfigSubjectConfigOutputReference
	_jsii_.Get(
		j,
		"subjectConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) SubjectConfigInput() *GooglePrivatecaCertificateConfigSubjectConfig {
	var returns *GooglePrivatecaCertificateConfigSubjectConfig
	_jsii_.Get(
		j,
		"subjectConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) SubjectKeyId() GooglePrivatecaCertificateConfigSubjectKeyIdOutputReference {
	var returns GooglePrivatecaCertificateConfigSubjectKeyIdOutputReference
	_jsii_.Get(
		j,
		"subjectKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) SubjectKeyIdInput() *GooglePrivatecaCertificateConfigSubjectKeyId {
	var returns *GooglePrivatecaCertificateConfigSubjectKeyId
	_jsii_.Get(
		j,
		"subjectKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) X509Config() GooglePrivatecaCertificateConfigX509ConfigOutputReference {
	var returns GooglePrivatecaCertificateConfigX509ConfigOutputReference
	_jsii_.Get(
		j,
		"x509Config",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) X509ConfigInput() *GooglePrivatecaCertificateConfigX509Config {
	var returns *GooglePrivatecaCertificateConfigX509Config
	_jsii_.Get(
		j,
		"x509ConfigInput",
		&returns,
	)
	return returns
}


func NewGooglePrivatecaCertificateConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GooglePrivatecaCertificateConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePrivatecaCertificateConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePrivatecaCertificate.GooglePrivatecaCertificateConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGooglePrivatecaCertificateConfigAOutputReference_Override(g GooglePrivatecaCertificateConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googlePrivatecaCertificate.GooglePrivatecaCertificateConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference)SetInternalValue(val *GooglePrivatecaCertificateConfigA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) PutPublicKey(value *GooglePrivatecaCertificateConfigPublicKey) {
	if err := g.validatePutPublicKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPublicKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) PutSubjectConfig(value *GooglePrivatecaCertificateConfigSubjectConfig) {
	if err := g.validatePutSubjectConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSubjectConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) PutSubjectKeyId(value *GooglePrivatecaCertificateConfigSubjectKeyId) {
	if err := g.validatePutSubjectKeyIdParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSubjectKeyId",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) PutX509Config(value *GooglePrivatecaCertificateConfigX509Config) {
	if err := g.validatePutX509ConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putX509Config",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) ResetSubjectKeyId() {
	_jsii_.InvokeVoid(
		g,
		"resetSubjectKeyId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GooglePrivatecaCertificateConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

