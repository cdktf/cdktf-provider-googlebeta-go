// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlenetworkservicesgrpcroute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googlenetworkservicesgrpcroute/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkServicesGrpcRouteRulesActionOutputReference interface {
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
	Destinations() GoogleNetworkServicesGrpcRouteRulesActionDestinationsList
	DestinationsInput() interface{}
	FaultInjectionPolicy() GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference
	FaultInjectionPolicyInput() *GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicy
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleNetworkServicesGrpcRouteRulesAction
	SetInternalValue(val *GoogleNetworkServicesGrpcRouteRulesAction)
	RetryPolicy() GoogleNetworkServicesGrpcRouteRulesActionRetryPolicyOutputReference
	RetryPolicyInput() *GoogleNetworkServicesGrpcRouteRulesActionRetryPolicy
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Timeout() *string
	SetTimeout(val *string)
	TimeoutInput() *string
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
	PutDestinations(value interface{})
	PutFaultInjectionPolicy(value *GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicy)
	PutRetryPolicy(value *GoogleNetworkServicesGrpcRouteRulesActionRetryPolicy)
	ResetDestinations()
	ResetFaultInjectionPolicy()
	ResetRetryPolicy()
	ResetTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkServicesGrpcRouteRulesActionOutputReference
type jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) Destinations() GoogleNetworkServicesGrpcRouteRulesActionDestinationsList {
	var returns GoogleNetworkServicesGrpcRouteRulesActionDestinationsList
	_jsii_.Get(
		j,
		"destinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) DestinationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) FaultInjectionPolicy() GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference {
	var returns GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicyOutputReference
	_jsii_.Get(
		j,
		"faultInjectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) FaultInjectionPolicyInput() *GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicy {
	var returns *GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicy
	_jsii_.Get(
		j,
		"faultInjectionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) InternalValue() *GoogleNetworkServicesGrpcRouteRulesAction {
	var returns *GoogleNetworkServicesGrpcRouteRulesAction
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) RetryPolicy() GoogleNetworkServicesGrpcRouteRulesActionRetryPolicyOutputReference {
	var returns GoogleNetworkServicesGrpcRouteRulesActionRetryPolicyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) RetryPolicyInput() *GoogleNetworkServicesGrpcRouteRulesActionRetryPolicy {
	var returns *GoogleNetworkServicesGrpcRouteRulesActionRetryPolicy
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) Timeout() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) TimeoutInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}


func NewGoogleNetworkServicesGrpcRouteRulesActionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleNetworkServicesGrpcRouteRulesActionOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkServicesGrpcRouteRulesActionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkServicesGrpcRoute.GoogleNetworkServicesGrpcRouteRulesActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleNetworkServicesGrpcRouteRulesActionOutputReference_Override(g GoogleNetworkServicesGrpcRouteRulesActionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleNetworkServicesGrpcRoute.GoogleNetworkServicesGrpcRouteRulesActionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference)SetInternalValue(val *GoogleNetworkServicesGrpcRouteRulesAction) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference)SetTimeout(val *string) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) PutDestinations(value interface{}) {
	if err := g.validatePutDestinationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDestinations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) PutFaultInjectionPolicy(value *GoogleNetworkServicesGrpcRouteRulesActionFaultInjectionPolicy) {
	if err := g.validatePutFaultInjectionPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFaultInjectionPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) PutRetryPolicy(value *GoogleNetworkServicesGrpcRouteRulesActionRetryPolicy) {
	if err := g.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) ResetDestinations() {
	_jsii_.InvokeVoid(
		g,
		"resetDestinations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) ResetFaultInjectionPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetFaultInjectionPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeout",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleNetworkServicesGrpcRouteRulesActionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

