package googlevertexaiendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v4/jsii"

	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v4/googlevertexaiendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) GoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecList
type jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecList {
	_init_.Initialize()

	if err := validateNewGoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecList{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiEndpoint.GoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecList_Override(g GoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleVertexAiEndpoint.GoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecList) Get(index *float64) GoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleVertexAiEndpointDeployedModelsDedicatedResourcesMachineSpecList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
