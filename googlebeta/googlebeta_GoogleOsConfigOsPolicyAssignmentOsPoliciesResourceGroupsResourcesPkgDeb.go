// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb struct {
	// source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_os_config_os_policy_assignment#source GoogleOsConfigOsPolicyAssignment#source}
	Source *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource `field:"required" json:"source" yaml:"source"`
	// Whether dependencies should also be installed.
	//
	// - install when false: `dpkg -i package` - install when true: `apt-get update && apt-get -y install package.deb`
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_os_config_os_policy_assignment#pull_deps GoogleOsConfigOsPolicyAssignment#pull_deps}
	PullDeps interface{} `field:"optional" json:"pullDeps" yaml:"pullDeps"`
}

