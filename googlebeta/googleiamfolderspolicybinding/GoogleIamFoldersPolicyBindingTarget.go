// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleiamfolderspolicybinding


type GoogleIamFoldersPolicyBindingTarget struct {
	// Required.
	//
	// Immutable. The resource name of the policy to be bound.
	// The binding parent and policy must belong to the same Organization (or Project).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.17.0/docs/resources/google_iam_folders_policy_binding#principal_set GoogleIamFoldersPolicyBinding#principal_set}
	PrincipalSet *string `field:"optional" json:"principalSet" yaml:"principalSet"`
}

