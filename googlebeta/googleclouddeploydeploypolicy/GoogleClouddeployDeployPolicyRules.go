// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleclouddeploydeploypolicy


type GoogleClouddeployDeployPolicyRules struct {
	// rollout_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.46.0/docs/resources/google_clouddeploy_deploy_policy#rollout_restriction GoogleClouddeployDeployPolicy#rollout_restriction}
	RolloutRestriction *GoogleClouddeployDeployPolicyRulesRolloutRestriction `field:"optional" json:"rolloutRestriction" yaml:"rolloutRestriction"`
}

