// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlechronicledataaccessscope

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v16/googlechronicledataaccessscope/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference interface {
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
	IngestionLabelKey() *string
	SetIngestionLabelKey(val *string)
	IngestionLabelKeyInput() *string
	IngestionLabelValue() *string
	SetIngestionLabelValue(val *string)
	IngestionLabelValueInput() *string
	InternalValue() *GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabel
	SetInternalValue(val *GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabel)
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
	ResetIngestionLabelValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference
type jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) IngestionLabelKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingestionLabelKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) IngestionLabelKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingestionLabelKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) IngestionLabelValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingestionLabelValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) IngestionLabelValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ingestionLabelValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) InternalValue() *GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabel {
	var returns *GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabel
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleChronicleDataAccessScope.GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference_Override(g GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleChronicleDataAccessScope.GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference)SetIngestionLabelKey(val *string) {
	if err := j.validateSetIngestionLabelKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingestionLabelKey",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference)SetIngestionLabelValue(val *string) {
	if err := j.validateSetIngestionLabelValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingestionLabelValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference)SetInternalValue(val *GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabel) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) ResetIngestionLabelValue() {
	_jsii_.InvokeVoid(
		g,
		"resetIngestionLabelValue",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

