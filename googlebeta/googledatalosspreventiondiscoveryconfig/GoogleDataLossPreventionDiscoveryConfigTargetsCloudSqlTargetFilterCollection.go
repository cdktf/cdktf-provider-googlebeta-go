// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledatalosspreventiondiscoveryconfig


type GoogleDataLossPreventionDiscoveryConfigTargetsCloudSqlTargetFilterCollection struct {
	// include_regexes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.8.0/docs/resources/google_data_loss_prevention_discovery_config#include_regexes GoogleDataLossPreventionDiscoveryConfig#include_regexes}
	IncludeRegexes *GoogleDataLossPreventionDiscoveryConfigTargetsCloudSqlTargetFilterCollectionIncludeRegexes `field:"optional" json:"includeRegexes" yaml:"includeRegexes"`
}

