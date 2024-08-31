// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleaccesscontextmanageraccesslevelcondition


type GoogleAccessContextManagerAccessLevelConditionVpcNetworkSources struct {
	// vpc_subnetwork block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.43.1/docs/resources/google_access_context_manager_access_level_condition#vpc_subnetwork GoogleAccessContextManagerAccessLevelCondition#vpc_subnetwork}
	VpcSubnetwork *GoogleAccessContextManagerAccessLevelConditionVpcNetworkSourcesVpcSubnetwork `field:"optional" json:"vpcSubnetwork" yaml:"vpcSubnetwork"`
}

