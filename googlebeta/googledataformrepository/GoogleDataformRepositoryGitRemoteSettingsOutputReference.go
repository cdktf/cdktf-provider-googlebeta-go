// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledataformrepository

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googledataformrepository/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDataformRepositoryGitRemoteSettingsOutputReference interface {
	cdktf.ComplexObject
	AuthenticationTokenSecretVersion() *string
	SetAuthenticationTokenSecretVersion(val *string)
	AuthenticationTokenSecretVersionInput() *string
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
	DefaultBranch() *string
	SetDefaultBranch(val *string)
	DefaultBranchInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDataformRepositoryGitRemoteSettings
	SetInternalValue(val *GoogleDataformRepositoryGitRemoteSettings)
	SshAuthenticationConfig() GoogleDataformRepositoryGitRemoteSettingsSshAuthenticationConfigOutputReference
	SshAuthenticationConfigInput() *GoogleDataformRepositoryGitRemoteSettingsSshAuthenticationConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TokenStatus() *string
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
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
	PutSshAuthenticationConfig(value *GoogleDataformRepositoryGitRemoteSettingsSshAuthenticationConfig)
	ResetAuthenticationTokenSecretVersion()
	ResetSshAuthenticationConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataformRepositoryGitRemoteSettingsOutputReference
type jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) AuthenticationTokenSecretVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTokenSecretVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) AuthenticationTokenSecretVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTokenSecretVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) DefaultBranch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) DefaultBranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBranchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) InternalValue() *GoogleDataformRepositoryGitRemoteSettings {
	var returns *GoogleDataformRepositoryGitRemoteSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) SshAuthenticationConfig() GoogleDataformRepositoryGitRemoteSettingsSshAuthenticationConfigOutputReference {
	var returns GoogleDataformRepositoryGitRemoteSettingsSshAuthenticationConfigOutputReference
	_jsii_.Get(
		j,
		"sshAuthenticationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) SshAuthenticationConfigInput() *GoogleDataformRepositoryGitRemoteSettingsSshAuthenticationConfig {
	var returns *GoogleDataformRepositoryGitRemoteSettingsSshAuthenticationConfig
	_jsii_.Get(
		j,
		"sshAuthenticationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) TokenStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewGoogleDataformRepositoryGitRemoteSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataformRepositoryGitRemoteSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataformRepositoryGitRemoteSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataformRepository.GoogleDataformRepositoryGitRemoteSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataformRepositoryGitRemoteSettingsOutputReference_Override(g GoogleDataformRepositoryGitRemoteSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDataformRepository.GoogleDataformRepositoryGitRemoteSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference)SetAuthenticationTokenSecretVersion(val *string) {
	if err := j.validateSetAuthenticationTokenSecretVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationTokenSecretVersion",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference)SetDefaultBranch(val *string) {
	if err := j.validateSetDefaultBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultBranch",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference)SetInternalValue(val *GoogleDataformRepositoryGitRemoteSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) PutSshAuthenticationConfig(value *GoogleDataformRepositoryGitRemoteSettingsSshAuthenticationConfig) {
	if err := g.validatePutSshAuthenticationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSshAuthenticationConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) ResetAuthenticationTokenSecretVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetAuthenticationTokenSecretVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) ResetSshAuthenticationConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSshAuthenticationConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataformRepositoryGitRemoteSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

