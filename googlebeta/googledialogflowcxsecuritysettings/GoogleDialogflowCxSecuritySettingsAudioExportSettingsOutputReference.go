// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledialogflowcxsecuritysettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googledialogflowcxsecuritysettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference interface {
	cdktf.ComplexObject
	AudioExportPattern() *string
	SetAudioExportPattern(val *string)
	AudioExportPatternInput() *string
	AudioFormat() *string
	SetAudioFormat(val *string)
	AudioFormatInput() *string
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
	EnableAudioRedaction() interface{}
	SetEnableAudioRedaction(val interface{})
	EnableAudioRedactionInput() interface{}
	// Experimental.
	Fqn() *string
	GcsBucket() *string
	SetGcsBucket(val *string)
	GcsBucketInput() *string
	InternalValue() *GoogleDialogflowCxSecuritySettingsAudioExportSettings
	SetInternalValue(val *GoogleDialogflowCxSecuritySettingsAudioExportSettings)
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
	ResetAudioExportPattern()
	ResetAudioFormat()
	ResetEnableAudioRedaction()
	ResetGcsBucket()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference
type jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) AudioExportPattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioExportPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) AudioExportPatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioExportPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) AudioFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) AudioFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"audioFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) EnableAudioRedaction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAudioRedaction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) EnableAudioRedactionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAudioRedactionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) GcsBucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsBucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) GcsBucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsBucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) InternalValue() *GoogleDialogflowCxSecuritySettingsAudioExportSettings {
	var returns *GoogleDialogflowCxSecuritySettingsAudioExportSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxSecuritySettings.GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference_Override(g GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxSecuritySettings.GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference)SetAudioExportPattern(val *string) {
	if err := j.validateSetAudioExportPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioExportPattern",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference)SetAudioFormat(val *string) {
	if err := j.validateSetAudioFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"audioFormat",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference)SetEnableAudioRedaction(val interface{}) {
	if err := j.validateSetEnableAudioRedactionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAudioRedaction",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference)SetGcsBucket(val *string) {
	if err := j.validateSetGcsBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcsBucket",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference)SetInternalValue(val *GoogleDialogflowCxSecuritySettingsAudioExportSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) ResetAudioExportPattern() {
	_jsii_.InvokeVoid(
		g,
		"resetAudioExportPattern",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) ResetAudioFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetAudioFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) ResetEnableAudioRedaction() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableAudioRedaction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) ResetGcsBucket() {
	_jsii_.InvokeVoid(
		g,
		"resetGcsBucket",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxSecuritySettingsAudioExportSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

