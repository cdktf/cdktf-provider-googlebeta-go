// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleclouddeploytarget


type GoogleClouddeployTargetAnthosCluster struct {
	// Membership of the GKE Hub-registered cluster to which to apply the Skaffold configuration. Format is `projects/{project}/locations/{location}/memberships/{membership_name}`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.40.0/docs/resources/google_clouddeploy_target#membership GoogleClouddeployTarget#membership}
	Membership *string `field:"optional" json:"membership" yaml:"membership"`
}

