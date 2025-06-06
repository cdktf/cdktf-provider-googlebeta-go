// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googletranscoderjobtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googletranscoderjobtemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference interface {
	cdktf.ComplexObject
	BitrateBps() *float64
	SetBitrateBps(val *float64)
	BitrateBpsInput() *float64
	ChannelCount() *float64
	SetChannelCount(val *float64)
	ChannelCountInput() *float64
	ChannelLayout() *[]*string
	SetChannelLayout(val *[]*string)
	ChannelLayoutInput() *[]*string
	Codec() *string
	SetCodec(val *string)
	CodecInput() *string
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
	InternalValue() *GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStream
	SetInternalValue(val *GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStream)
	SampleRateHertz() *float64
	SetSampleRateHertz(val *float64)
	SampleRateHertzInput() *float64
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
	ResetChannelCount()
	ResetChannelLayout()
	ResetCodec()
	ResetSampleRateHertz()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference
type jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) BitrateBps() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrateBps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) BitrateBpsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bitrateBpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) ChannelCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"channelCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) ChannelCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"channelCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) ChannelLayout() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"channelLayout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) ChannelLayoutInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"channelLayoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) Codec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) CodecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) InternalValue() *GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStream {
	var returns *GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStream
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) SampleRateHertz() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleRateHertz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) SampleRateHertzInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sampleRateHertzInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleTranscoderJobTemplate.GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference_Override(g GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleTranscoderJobTemplate.GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference)SetBitrateBps(val *float64) {
	if err := j.validateSetBitrateBpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitrateBps",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference)SetChannelCount(val *float64) {
	if err := j.validateSetChannelCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelCount",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference)SetChannelLayout(val *[]*string) {
	if err := j.validateSetChannelLayoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channelLayout",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference)SetCodec(val *string) {
	if err := j.validateSetCodecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codec",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference)SetInternalValue(val *GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStream) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference)SetSampleRateHertz(val *float64) {
	if err := j.validateSetSampleRateHertzParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleRateHertz",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) ResetChannelCount() {
	_jsii_.InvokeVoid(
		g,
		"resetChannelCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) ResetChannelLayout() {
	_jsii_.InvokeVoid(
		g,
		"resetChannelLayout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) ResetCodec() {
	_jsii_.InvokeVoid(
		g,
		"resetCodec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) ResetSampleRateHertz() {
	_jsii_.InvokeVoid(
		g,
		"resetSampleRateHertz",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigElementaryStreamsAudioStreamOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

