// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleOsConfigPatchDeploymentRolloutDisruptionBudget struct {
	// Specifies a fixed value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_os_config_patch_deployment#fixed GoogleOsConfigPatchDeployment#fixed}
	Fixed *float64 `field:"optional" json:"fixed" yaml:"fixed"`
	// Specifies the relative value defined as a percentage, which will be multiplied by a reference value.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_os_config_patch_deployment#percentage GoogleOsConfigPatchDeployment#percentage}
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
}

