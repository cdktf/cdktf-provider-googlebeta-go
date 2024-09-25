// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleidentityplatformconfig


type GoogleIdentityPlatformConfigMonitoring struct {
	// request_logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.4.0/docs/resources/google_identity_platform_config#request_logging GoogleIdentityPlatformConfig#request_logging}
	RequestLogging *GoogleIdentityPlatformConfigMonitoringRequestLogging `field:"optional" json:"requestLogging" yaml:"requestLogging"`
}

