// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleosconfigpatchdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googleosconfigpatchdeployment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleOsConfigPatchDeploymentPatchConfigOutputReference interface {
	cdktf.ComplexObject
	Apt() GoogleOsConfigPatchDeploymentPatchConfigAptOutputReference
	AptInput() *GoogleOsConfigPatchDeploymentPatchConfigApt
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
	Goo() GoogleOsConfigPatchDeploymentPatchConfigGooOutputReference
	GooInput() *GoogleOsConfigPatchDeploymentPatchConfigGoo
	InternalValue() *GoogleOsConfigPatchDeploymentPatchConfig
	SetInternalValue(val *GoogleOsConfigPatchDeploymentPatchConfig)
	MigInstancesAllowed() interface{}
	SetMigInstancesAllowed(val interface{})
	MigInstancesAllowedInput() interface{}
	PostStep() GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference
	PostStepInput() *GoogleOsConfigPatchDeploymentPatchConfigPostStep
	PreStep() GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference
	PreStepInput() *GoogleOsConfigPatchDeploymentPatchConfigPreStep
	RebootConfig() *string
	SetRebootConfig(val *string)
	RebootConfigInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WindowsUpdate() GoogleOsConfigPatchDeploymentPatchConfigWindowsUpdateOutputReference
	WindowsUpdateInput() *GoogleOsConfigPatchDeploymentPatchConfigWindowsUpdate
	Yum() GoogleOsConfigPatchDeploymentPatchConfigYumOutputReference
	YumInput() *GoogleOsConfigPatchDeploymentPatchConfigYum
	Zypper() GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference
	ZypperInput() *GoogleOsConfigPatchDeploymentPatchConfigZypper
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
	PutApt(value *GoogleOsConfigPatchDeploymentPatchConfigApt)
	PutGoo(value *GoogleOsConfigPatchDeploymentPatchConfigGoo)
	PutPostStep(value *GoogleOsConfigPatchDeploymentPatchConfigPostStep)
	PutPreStep(value *GoogleOsConfigPatchDeploymentPatchConfigPreStep)
	PutWindowsUpdate(value *GoogleOsConfigPatchDeploymentPatchConfigWindowsUpdate)
	PutYum(value *GoogleOsConfigPatchDeploymentPatchConfigYum)
	PutZypper(value *GoogleOsConfigPatchDeploymentPatchConfigZypper)
	ResetApt()
	ResetGoo()
	ResetMigInstancesAllowed()
	ResetPostStep()
	ResetPreStep()
	ResetRebootConfig()
	ResetWindowsUpdate()
	ResetYum()
	ResetZypper()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOsConfigPatchDeploymentPatchConfigOutputReference
type jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) Apt() GoogleOsConfigPatchDeploymentPatchConfigAptOutputReference {
	var returns GoogleOsConfigPatchDeploymentPatchConfigAptOutputReference
	_jsii_.Get(
		j,
		"apt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) AptInput() *GoogleOsConfigPatchDeploymentPatchConfigApt {
	var returns *GoogleOsConfigPatchDeploymentPatchConfigApt
	_jsii_.Get(
		j,
		"aptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) Goo() GoogleOsConfigPatchDeploymentPatchConfigGooOutputReference {
	var returns GoogleOsConfigPatchDeploymentPatchConfigGooOutputReference
	_jsii_.Get(
		j,
		"goo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) GooInput() *GoogleOsConfigPatchDeploymentPatchConfigGoo {
	var returns *GoogleOsConfigPatchDeploymentPatchConfigGoo
	_jsii_.Get(
		j,
		"gooInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) InternalValue() *GoogleOsConfigPatchDeploymentPatchConfig {
	var returns *GoogleOsConfigPatchDeploymentPatchConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) MigInstancesAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"migInstancesAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) MigInstancesAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"migInstancesAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) PostStep() GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference {
	var returns GoogleOsConfigPatchDeploymentPatchConfigPostStepOutputReference
	_jsii_.Get(
		j,
		"postStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) PostStepInput() *GoogleOsConfigPatchDeploymentPatchConfigPostStep {
	var returns *GoogleOsConfigPatchDeploymentPatchConfigPostStep
	_jsii_.Get(
		j,
		"postStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) PreStep() GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference {
	var returns GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference
	_jsii_.Get(
		j,
		"preStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) PreStepInput() *GoogleOsConfigPatchDeploymentPatchConfigPreStep {
	var returns *GoogleOsConfigPatchDeploymentPatchConfigPreStep
	_jsii_.Get(
		j,
		"preStepInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) RebootConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rebootConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) RebootConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rebootConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) WindowsUpdate() GoogleOsConfigPatchDeploymentPatchConfigWindowsUpdateOutputReference {
	var returns GoogleOsConfigPatchDeploymentPatchConfigWindowsUpdateOutputReference
	_jsii_.Get(
		j,
		"windowsUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) WindowsUpdateInput() *GoogleOsConfigPatchDeploymentPatchConfigWindowsUpdate {
	var returns *GoogleOsConfigPatchDeploymentPatchConfigWindowsUpdate
	_jsii_.Get(
		j,
		"windowsUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) Yum() GoogleOsConfigPatchDeploymentPatchConfigYumOutputReference {
	var returns GoogleOsConfigPatchDeploymentPatchConfigYumOutputReference
	_jsii_.Get(
		j,
		"yum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) YumInput() *GoogleOsConfigPatchDeploymentPatchConfigYum {
	var returns *GoogleOsConfigPatchDeploymentPatchConfigYum
	_jsii_.Get(
		j,
		"yumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) Zypper() GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference {
	var returns GoogleOsConfigPatchDeploymentPatchConfigZypperOutputReference
	_jsii_.Get(
		j,
		"zypper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) ZypperInput() *GoogleOsConfigPatchDeploymentPatchConfigZypper {
	var returns *GoogleOsConfigPatchDeploymentPatchConfigZypper
	_jsii_.Get(
		j,
		"zypperInput",
		&returns,
	)
	return returns
}


func NewGoogleOsConfigPatchDeploymentPatchConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleOsConfigPatchDeploymentPatchConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOsConfigPatchDeploymentPatchConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOsConfigPatchDeployment.GoogleOsConfigPatchDeploymentPatchConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOsConfigPatchDeploymentPatchConfigOutputReference_Override(g GoogleOsConfigPatchDeploymentPatchConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOsConfigPatchDeployment.GoogleOsConfigPatchDeploymentPatchConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference)SetInternalValue(val *GoogleOsConfigPatchDeploymentPatchConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference)SetMigInstancesAllowed(val interface{}) {
	if err := j.validateSetMigInstancesAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"migInstancesAllowed",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference)SetRebootConfig(val *string) {
	if err := j.validateSetRebootConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rebootConfig",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) PutApt(value *GoogleOsConfigPatchDeploymentPatchConfigApt) {
	if err := g.validatePutAptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApt",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) PutGoo(value *GoogleOsConfigPatchDeploymentPatchConfigGoo) {
	if err := g.validatePutGooParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGoo",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) PutPostStep(value *GoogleOsConfigPatchDeploymentPatchConfigPostStep) {
	if err := g.validatePutPostStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPostStep",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) PutPreStep(value *GoogleOsConfigPatchDeploymentPatchConfigPreStep) {
	if err := g.validatePutPreStepParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPreStep",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) PutWindowsUpdate(value *GoogleOsConfigPatchDeploymentPatchConfigWindowsUpdate) {
	if err := g.validatePutWindowsUpdateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWindowsUpdate",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) PutYum(value *GoogleOsConfigPatchDeploymentPatchConfigYum) {
	if err := g.validatePutYumParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putYum",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) PutZypper(value *GoogleOsConfigPatchDeploymentPatchConfigZypper) {
	if err := g.validatePutZypperParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putZypper",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) ResetApt() {
	_jsii_.InvokeVoid(
		g,
		"resetApt",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) ResetGoo() {
	_jsii_.InvokeVoid(
		g,
		"resetGoo",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) ResetMigInstancesAllowed() {
	_jsii_.InvokeVoid(
		g,
		"resetMigInstancesAllowed",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) ResetPostStep() {
	_jsii_.InvokeVoid(
		g,
		"resetPostStep",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) ResetPreStep() {
	_jsii_.InvokeVoid(
		g,
		"resetPreStep",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) ResetRebootConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetRebootConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) ResetWindowsUpdate() {
	_jsii_.InvokeVoid(
		g,
		"resetWindowsUpdate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) ResetYum() {
	_jsii_.InvokeVoid(
		g,
		"resetYum",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) ResetZypper() {
	_jsii_.InvokeVoid(
		g,
		"resetZypper",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

