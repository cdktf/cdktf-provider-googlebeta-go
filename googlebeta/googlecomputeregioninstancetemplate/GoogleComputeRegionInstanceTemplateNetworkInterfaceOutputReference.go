// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputeregioninstancetemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googlecomputeregioninstancetemplate/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference interface {
	cdktf.ComplexObject
	AccessConfig() GoogleComputeRegionInstanceTemplateNetworkInterfaceAccessConfigList
	AccessConfigInput() interface{}
	AliasIpRange() GoogleComputeRegionInstanceTemplateNetworkInterfaceAliasIpRangeList
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
	Ipv6AccessConfig() GoogleComputeRegionInstanceTemplateNetworkInterfaceIpv6AccessConfigList
	Ipv6AccessConfigInput() interface{}
	Ipv6AccessType() *string
	Ipv6Address() *string
	SetIpv6Address(val *string)
	Ipv6AddressInput() *string
	Name() *string
	Network() *string
	SetNetwork(val *string)
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
	ResetNetworkIp()
	ResetNicType()
	ResetQueueCount()
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

// The jsii proxy struct for GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference
type jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) AccessConfig() GoogleComputeRegionInstanceTemplateNetworkInterfaceAccessConfigList {
	var returns GoogleComputeRegionInstanceTemplateNetworkInterfaceAccessConfigList
	_jsii_.Get(
		j,
		"accessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) AccessConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) AliasIpRange() GoogleComputeRegionInstanceTemplateNetworkInterfaceAliasIpRangeList {
	var returns GoogleComputeRegionInstanceTemplateNetworkInterfaceAliasIpRangeList
	_jsii_.Get(
		j,
		"aliasIpRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) AliasIpRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aliasIpRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) InternalIpv6PrefixLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalIpv6PrefixLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) InternalIpv6PrefixLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"internalIpv6PrefixLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) Ipv6AccessConfig() GoogleComputeRegionInstanceTemplateNetworkInterfaceIpv6AccessConfigList {
	var returns GoogleComputeRegionInstanceTemplateNetworkInterfaceIpv6AccessConfigList
	_jsii_.Get(
		j,
		"ipv6AccessConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) Ipv6AccessConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6AccessConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) Ipv6AccessType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6AccessType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) Ipv6Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6Address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) Ipv6AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6AddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) NetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) NetworkIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) NetworkIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) NicType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nicType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) NicTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nicTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) QueueCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queueCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) QueueCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queueCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) StackType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) StackTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) Subnetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) SubnetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) SubnetworkProject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) SubnetworkProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetworkProjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeRegionInstanceTemplate.GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference_Override(g GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleComputeRegionInstanceTemplate.GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference)SetInternalIpv6PrefixLength(val *float64) {
	if err := j.validateSetInternalIpv6PrefixLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalIpv6PrefixLength",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference)SetIpv6Address(val *string) {
	if err := j.validateSetIpv6AddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Address",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference)SetNetwork(val *string) {
	if err := j.validateSetNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"network",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference)SetNetworkIp(val *string) {
	if err := j.validateSetNetworkIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkIp",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference)SetNicType(val *string) {
	if err := j.validateSetNicTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nicType",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference)SetQueueCount(val *float64) {
	if err := j.validateSetQueueCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queueCount",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference)SetStackType(val *string) {
	if err := j.validateSetStackTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackType",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference)SetSubnetwork(val *string) {
	if err := j.validateSetSubnetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetwork",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference)SetSubnetworkProject(val *string) {
	if err := j.validateSetSubnetworkProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetworkProject",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) PutAccessConfig(value interface{}) {
	if err := g.validatePutAccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAccessConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) PutAliasIpRange(value interface{}) {
	if err := g.validatePutAliasIpRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAliasIpRange",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) PutIpv6AccessConfig(value interface{}) {
	if err := g.validatePutIpv6AccessConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIpv6AccessConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) ResetAccessConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) ResetAliasIpRange() {
	_jsii_.InvokeVoid(
		g,
		"resetAliasIpRange",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) ResetInternalIpv6PrefixLength() {
	_jsii_.InvokeVoid(
		g,
		"resetInternalIpv6PrefixLength",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) ResetIpv6AccessConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetIpv6AccessConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) ResetIpv6Address() {
	_jsii_.InvokeVoid(
		g,
		"resetIpv6Address",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) ResetNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) ResetNetworkIp() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) ResetNicType() {
	_jsii_.InvokeVoid(
		g,
		"resetNicType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) ResetQueueCount() {
	_jsii_.InvokeVoid(
		g,
		"resetQueueCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) ResetStackType() {
	_jsii_.InvokeVoid(
		g,
		"resetStackType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) ResetSubnetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetSubnetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) ResetSubnetworkProject() {
	_jsii_.InvokeVoid(
		g,
		"resetSubnetworkProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateNetworkInterfaceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

