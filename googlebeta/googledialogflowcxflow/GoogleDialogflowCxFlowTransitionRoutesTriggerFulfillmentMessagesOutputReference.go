// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledialogflowcxflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googledialogflowcxflow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference interface {
	cdktf.ComplexObject
	Channel() *string
	SetChannel(val *string)
	ChannelInput() *string
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
	ConversationSuccess() GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesConversationSuccessOutputReference
	ConversationSuccessInput() *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesConversationSuccess
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LiveAgentHandoff() GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesLiveAgentHandoffOutputReference
	LiveAgentHandoffInput() *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesLiveAgentHandoff
	OutputAudioText() GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputAudioTextOutputReference
	OutputAudioTextInput() *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputAudioText
	Payload() *string
	SetPayload(val *string)
	PayloadInput() *string
	PlayAudio() GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesPlayAudioOutputReference
	PlayAudioInput() *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesPlayAudio
	TelephonyTransferCall() GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesTelephonyTransferCallOutputReference
	TelephonyTransferCallInput() *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesTelephonyTransferCall
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Text() GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesTextOutputReference
	TextInput() *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesText
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
	PutConversationSuccess(value *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesConversationSuccess)
	PutLiveAgentHandoff(value *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesLiveAgentHandoff)
	PutOutputAudioText(value *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputAudioText)
	PutPlayAudio(value *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesPlayAudio)
	PutTelephonyTransferCall(value *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesTelephonyTransferCall)
	PutText(value *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesText)
	ResetChannel()
	ResetConversationSuccess()
	ResetLiveAgentHandoff()
	ResetOutputAudioText()
	ResetPayload()
	ResetPlayAudio()
	ResetTelephonyTransferCall()
	ResetText()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference
type jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) Channel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) ChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"channelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) ConversationSuccess() GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesConversationSuccessOutputReference {
	var returns GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesConversationSuccessOutputReference
	_jsii_.Get(
		j,
		"conversationSuccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) ConversationSuccessInput() *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesConversationSuccess {
	var returns *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesConversationSuccess
	_jsii_.Get(
		j,
		"conversationSuccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) LiveAgentHandoff() GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesLiveAgentHandoffOutputReference {
	var returns GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesLiveAgentHandoffOutputReference
	_jsii_.Get(
		j,
		"liveAgentHandoff",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) LiveAgentHandoffInput() *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesLiveAgentHandoff {
	var returns *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesLiveAgentHandoff
	_jsii_.Get(
		j,
		"liveAgentHandoffInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) OutputAudioText() GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputAudioTextOutputReference {
	var returns GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputAudioTextOutputReference
	_jsii_.Get(
		j,
		"outputAudioText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) OutputAudioTextInput() *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputAudioText {
	var returns *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputAudioText
	_jsii_.Get(
		j,
		"outputAudioTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) Payload() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) PayloadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"payloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) PlayAudio() GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesPlayAudioOutputReference {
	var returns GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesPlayAudioOutputReference
	_jsii_.Get(
		j,
		"playAudio",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) PlayAudioInput() *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesPlayAudio {
	var returns *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesPlayAudio
	_jsii_.Get(
		j,
		"playAudioInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) TelephonyTransferCall() GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesTelephonyTransferCallOutputReference {
	var returns GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesTelephonyTransferCallOutputReference
	_jsii_.Get(
		j,
		"telephonyTransferCall",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) TelephonyTransferCallInput() *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesTelephonyTransferCall {
	var returns *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesTelephonyTransferCall
	_jsii_.Get(
		j,
		"telephonyTransferCallInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) Text() GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesTextOutputReference {
	var returns GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesTextOutputReference
	_jsii_.Get(
		j,
		"text",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) TextInput() *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesText {
	var returns *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesText
	_jsii_.Get(
		j,
		"textInput",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxFlow.GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference_Override(g GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleDialogflowCxFlow.GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference)SetChannel(val *string) {
	if err := j.validateSetChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"channel",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference)SetPayload(val *string) {
	if err := j.validateSetPayloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"payload",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) PutConversationSuccess(value *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesConversationSuccess) {
	if err := g.validatePutConversationSuccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConversationSuccess",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) PutLiveAgentHandoff(value *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesLiveAgentHandoff) {
	if err := g.validatePutLiveAgentHandoffParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLiveAgentHandoff",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) PutOutputAudioText(value *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputAudioText) {
	if err := g.validatePutOutputAudioTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOutputAudioText",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) PutPlayAudio(value *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesPlayAudio) {
	if err := g.validatePutPlayAudioParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPlayAudio",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) PutTelephonyTransferCall(value *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesTelephonyTransferCall) {
	if err := g.validatePutTelephonyTransferCallParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTelephonyTransferCall",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) PutText(value *GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesText) {
	if err := g.validatePutTextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putText",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) ResetChannel() {
	_jsii_.InvokeVoid(
		g,
		"resetChannel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) ResetConversationSuccess() {
	_jsii_.InvokeVoid(
		g,
		"resetConversationSuccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) ResetLiveAgentHandoff() {
	_jsii_.InvokeVoid(
		g,
		"resetLiveAgentHandoff",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) ResetOutputAudioText() {
	_jsii_.InvokeVoid(
		g,
		"resetOutputAudioText",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) ResetPayload() {
	_jsii_.InvokeVoid(
		g,
		"resetPayload",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) ResetPlayAudio() {
	_jsii_.InvokeVoid(
		g,
		"resetPlayAudio",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) ResetTelephonyTransferCall() {
	_jsii_.InvokeVoid(
		g,
		"resetTelephonyTransferCall",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) ResetText() {
	_jsii_.InvokeVoid(
		g,
		"resetText",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowTransitionRoutesTriggerFulfillmentMessagesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

