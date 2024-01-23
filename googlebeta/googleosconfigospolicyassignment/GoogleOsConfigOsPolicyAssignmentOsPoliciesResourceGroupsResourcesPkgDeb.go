// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleosconfigospolicyassignment


type GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDeb struct {
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.13.0/docs/resources/google_os_config_os_policy_assignment#source GoogleOsConfigOsPolicyAssignment#source}
	Source *GoogleOsConfigOsPolicyAssignmentOsPoliciesResourceGroupsResourcesPkgDebSource `field:"required" json:"source" yaml:"source"`
	// Whether dependencies should also be installed.
	//
	// - install when false: 'dpkg -i package' - install when true: 'apt-get update && apt-get -y install package.deb'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.13.0/docs/resources/google_os_config_os_policy_assignment#pull_deps GoogleOsConfigOsPolicyAssignment#pull_deps}
	PullDeps interface{} `field:"optional" json:"pullDeps" yaml:"pullDeps"`
}

