// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputeinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googlecomputeinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeInstanceNetworkInterfaceOutputReference interface {
	cdktf.ComplexObject
	AccessConfig() GoogleComputeInstanceNetworkInterfaceAccessConfigList
	AccessConfigInput() interface{}
	AliasIpRange() GoogleComputeInstanceNetworkInterfaceAliasIpRangeList
	AliasIpRangeInput() interface{}
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
	InternalIpv6PrefixLength() *float64
	SetInternalIpv6PrefixLength(val *float64)
	InternalIpv6PrefixLengthInput() *float64
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Ipv6AccessConfig() GoogleComputeInstanceNetworkInterfaceIpv6AccessConfigList
	Ipv6AccessConfigInput() interface{}
	Ipv6AccessType() *string
	Ipv6Address() *string
	SetIpv6Address(val *string)
	Ipv6AddressInput() *string
	Name() *string
	Network() *string
	SetNetwork(val *string)
	NetworkAttachment() *string
	SetNetworkAttachment(val *string)
	NetworkAttachmentInput() *string
	NetworkInput() *string
	NetworkIp() *string
	SetNetworkIp(val *string)
	NetworkIpInput() *string
	NicType() *string
	SetNicType(val *string)
	NicTypeInput() *string
	QueueCount() *float64
	SetQueueCount(val *float64)
	QueueCountInput() *float64
	SecurityPolicy() *string
	SetSecurityPolicy(val *string)
	SecurityPolicyInput() *string
	StackType() *string
	SetStackType(val *string)
	StackTypeInput() *string
	Subnetwork() *string
	SetSubnetwork(val *string)
	SubnetworkInput() *string
	SubnetworkProject() *string
	SetSubnetworkProject(val *string)
	SubnetworkProjectInput() *string
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
	PutAccessConfig(value interface{})
	PutAliasIpRange(value interface{})
	PutIpv6AccessConfig(value interface{})
	ResetAccessConfig()
	ResetAliasIpRange()
	ResetInternalIpv6PrefixLength()
	ResetIpv6AccessConfig()
	ResetIpv6Address()
	ResetNetwork()
	ResetNetworkAttachment()
	ResetNetworkIp()
	ResetNicType()
	ResetQueueCount()
	ResetSecurityPolicy()
	ResetStackType()
	ResetSubnetwork()
	ResetSubnetworkProject()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeInstanceNetworkInterfaceOutputReference
type jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) AccessConfig() GoogleComputeInstanceNetworkInterfaceAccessConfigList {
	var returns GoogleComputeInstanceNetworkInterfaceAccessConfigList
	_jsii_.Get(
		j,
		"accessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) AccessConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) AliasIpRange() GoogleComputeInstanceNetworkInterfaceAliasIpRangeList {
	var returns GoogleComputeInstanceNetworkInterfaceAliasIpRangeList
	_jsii_.Get(
		j,
		"aliasIpRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) AliasIpRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aliasIpRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) InternalIpv6PrefixLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalIpv6PrefixLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) InternalIpv6PrefixLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalIpv6PrefixLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) Ipv6AccessConfig() GoogleComputeInstanceNetworkInterfaceIpv6AccessConfigList {
	var returns GoogleComputeInstanceNetworkInterfaceIpv6AccessConfigList
	_jsii_.Get(
		j,
		"ipv6AccessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) Ipv6AccessConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6AccessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) Ipv6AccessType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6AccessType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) Ipv6Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) Ipv6AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6AddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) NetworkAttachment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAttachment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) NetworkAttachmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkAttachmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) NetworkIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) NetworkIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) NicType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nicType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) NicTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nicTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) QueueCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queueCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) QueueCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queueCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) SecurityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) SecurityPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) StackType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) StackTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) Subnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) SubnetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) SubnetworkProject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) SubnetworkProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkProjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeInstanceNetworkInterfaceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleComputeInstanceNetworkInterfaceOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeInstanceNetworkInterfaceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeInstance.GoogleComputeInstanceNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleComputeInstanceNetworkInterfaceOutputReference_Override(g GoogleComputeInstanceNetworkInterfaceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeInstance.GoogleComputeInstanceNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference)SetInternalIpv6PrefixLength(val *float64) {
	if err := j.validateSetInternalIpv6PrefixLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalIpv6PrefixLength",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference)SetIpv6Address(val *string) {
	if err := j.validateSetIpv6AddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Address",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference)SetNetworkAttachment(val *string) {
	if err := j.validateSetNetworkAttachmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkAttachment",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference)SetNetworkIp(val *string) {
	if err := j.validateSetNetworkIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkIp",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference)SetNicType(val *string) {
	if err := j.validateSetNicTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nicType",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference)SetQueueCount(val *float64) {
	if err := j.validateSetQueueCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queueCount",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference)SetSecurityPolicy(val *string) {
	if err := j.validateSetSecurityPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference)SetStackType(val *string) {
	if err := j.validateSetStackTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackType",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference)SetSubnetwork(val *string) {
	if err := j.validateSetSubnetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetwork",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference)SetSubnetworkProject(val *string) {
	if err := j.validateSetSubnetworkProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetworkProject",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) PutAccessConfig(value interface{}) {
	if err := g.validatePutAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAccessConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) PutAliasIpRange(value interface{}) {
	if err := g.validatePutAliasIpRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAliasIpRange",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) PutIpv6AccessConfig(value interface{}) {
	if err := g.validatePutIpv6AccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIpv6AccessConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) ResetAccessConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) ResetAliasIpRange() {
	_jsii_.InvokeVoid(
		g,
		"resetAliasIpRange",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) ResetInternalIpv6PrefixLength() {
	_jsii_.InvokeVoid(
		g,
		"resetInternalIpv6PrefixLength",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) ResetIpv6AccessConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetIpv6AccessConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) ResetIpv6Address() {
	_jsii_.InvokeVoid(
		g,
		"resetIpv6Address",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) ResetNetworkAttachment() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkAttachment",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) ResetNetworkIp() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) ResetNicType() {
	_jsii_.InvokeVoid(
		g,
		"resetNicType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) ResetQueueCount() {
	_jsii_.InvokeVoid(
		g,
		"resetQueueCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) ResetSecurityPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) ResetStackType() {
	_jsii_.InvokeVoid(
		g,
		"resetStackType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) ResetSubnetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetSubnetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) ResetSubnetworkProject() {
	_jsii_.InvokeVoid(
		g,
		"resetSubnetworkProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeInstanceNetworkInterfaceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

