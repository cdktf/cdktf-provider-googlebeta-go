// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlediscoveryenginerecommendationengine


type GoogleDiscoveryEngineRecommendationEngineCommonConfig struct {
	// The name of the company, business or entity that is associated with the engine.
	//
	// Setting this may help improve LLM related features.cd
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.50.0/docs/resources/google_discovery_engine_recommendation_engine#company_name GoogleDiscoveryEngineRecommendationEngine#company_name}
	CompanyName *string `field:"optional" json:"companyName" yaml:"companyName"`
}

