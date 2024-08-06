// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputebackendservice


type GoogleComputeBackendServiceCdnPolicyBypassCacheOnRequestHeaders struct {
	// The header field name to match on when bypassing cache. Values are case-insensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.40.0/docs/resources/google_compute_backend_service#header_name GoogleComputeBackendService#header_name}
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
}

