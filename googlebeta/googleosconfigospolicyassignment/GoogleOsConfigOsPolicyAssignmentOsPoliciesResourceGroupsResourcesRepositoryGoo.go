package googleosconfigospolicyassignment


type GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo struct {
	// Required. The name of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.63.0/docs/resources/google_os_config_os_policy_assignment#name GoogleOsConfigOsPolicyAssignment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Required. The url of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.63.0/docs/resources/google_os_config_os_policy_assignment#url GoogleOsConfigOsPolicyAssignment#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

