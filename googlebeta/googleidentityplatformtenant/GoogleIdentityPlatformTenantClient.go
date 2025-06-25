// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleidentityplatformtenant


type GoogleIdentityPlatformTenantClient struct {
	// permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.41.0/docs/resources/google_identity_platform_tenant#permissions GoogleIdentityPlatformTenant#permissions}
	Permissions *GoogleIdentityPlatformTenantClientPermissions `field:"optional" json:"permissions" yaml:"permissions"`
}

