// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigActionsPubSubNotificationPubsubCondition struct {
	// expressions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.47.0/docs/resources/google_data_loss_prevention_discovery_config#expressions GoogleDataLossPreventionDiscoveryConfig#expressions}
	Expressions *GoogleDataLossPreventionDiscoveryConfigActionsPubSubNotificationPubsubConditionExpressions `field:"optional" json:"expressions" yaml:"expressions"`
}

