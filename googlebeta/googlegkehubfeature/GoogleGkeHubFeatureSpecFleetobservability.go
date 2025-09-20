// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlegkehubfeature


type GoogleGkeHubFeatureSpecFleetobservability struct {
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.50.0/docs/resources/google_gke_hub_feature#logging_config GoogleGkeHubFeature#logging_config}
	LoggingConfig *GoogleGkeHubFeatureSpecFleetobservabilityLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
}

