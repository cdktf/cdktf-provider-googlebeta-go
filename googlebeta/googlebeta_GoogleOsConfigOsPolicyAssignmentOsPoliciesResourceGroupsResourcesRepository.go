// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepository struct {
	// apt block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_os_config_os_policy_assignment#apt GoogleOsConfigOsPolicyAssignment#apt}
	Apt *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryApt `field:"optional" json:"apt" yaml:"apt"`
	// goo block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_os_config_os_policy_assignment#goo GoogleOsConfigOsPolicyAssignment#goo}
	Goo *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryGoo `field:"optional" json:"goo" yaml:"goo"`
	// yum block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_os_config_os_policy_assignment#yum GoogleOsConfigOsPolicyAssignment#yum}
	Yum *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryYum `field:"optional" json:"yum" yaml:"yum"`
	// zypper block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_os_config_os_policy_assignment#zypper GoogleOsConfigOsPolicyAssignment#zypper}
	Zypper *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesRepositoryZypper `field:"optional" json:"zypper" yaml:"zypper"`
}

