package googleosconfigpatchdeployment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v3/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v3/googleosconfigpatchdeployment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference interface {
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
	InternalValue() *GoogleOsConfigPatchDeploymentPatchConfigPreStep
	SetInternalValue(val *GoogleOsConfigPatchDeploymentPatchConfigPreStep)
	LinuxExecStepConfig() GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference
	LinuxExecStepConfigInput() *GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfig
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WindowsExecStepConfig() GoogleOsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference
	WindowsExecStepConfigInput() *GoogleOsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfig
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
	PutLinuxExecStepConfig(value *GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfig)
	PutWindowsExecStepConfig(value *GoogleOsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfig)
	ResetLinuxExecStepConfig()
	ResetWindowsExecStepConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference
type jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) InternalValue() *GoogleOsConfigPatchDeploymentPatchConfigPreStep {
	var returns *GoogleOsConfigPatchDeploymentPatchConfigPreStep
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) LinuxExecStepConfig() GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference {
	var returns GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfigOutputReference
	_jsii_.Get(
		j,
		"linuxExecStepConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) LinuxExecStepConfigInput() *GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfig {
	var returns *GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfig
	_jsii_.Get(
		j,
		"linuxExecStepConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) WindowsExecStepConfig() GoogleOsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference {
	var returns GoogleOsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfigOutputReference
	_jsii_.Get(
		j,
		"windowsExecStepConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) WindowsExecStepConfigInput() *GoogleOsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfig {
	var returns *GoogleOsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfig
	_jsii_.Get(
		j,
		"windowsExecStepConfigInput",
		&returns,
	)
	return returns
}


func NewGoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOsConfigPatchDeployment.GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference_Override(g GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleOsConfigPatchDeployment.GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference)SetInternalValue(val *GoogleOsConfigPatchDeploymentPatchConfigPreStep) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) PutLinuxExecStepConfig(value *GoogleOsConfigPatchDeploymentPatchConfigPreStepLinuxExecStepConfig) {
	if err := g.validatePutLinuxExecStepConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLinuxExecStepConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) PutWindowsExecStepConfig(value *GoogleOsConfigPatchDeploymentPatchConfigPreStepWindowsExecStepConfig) {
	if err := g.validatePutWindowsExecStepConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWindowsExecStepConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) ResetLinuxExecStepConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLinuxExecStepConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) ResetWindowsExecStepConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetWindowsExecStepConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleOsConfigPatchDeploymentPatchConfigPreStepOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
