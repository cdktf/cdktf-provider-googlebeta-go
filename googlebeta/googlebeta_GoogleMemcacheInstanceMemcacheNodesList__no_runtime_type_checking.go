//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GoogleMemcacheInstanceMemcacheNodesList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GoogleMemcacheInstanceMemcacheNodesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GoogleMemcacheInstanceMemcacheNodesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GoogleMemcacheInstanceMemcacheNodesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GoogleMemcacheInstanceMemcacheNodesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGoogleMemcacheInstanceMemcacheNodesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

