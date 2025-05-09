// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomputeurlmap


type GoogleComputeUrlMapPathMatcherDefaultRouteActionRequestMirrorPolicy struct {
	// The full or partial URL to the BackendService resource being mirrored to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_url_map#backend_service GoogleComputeUrlMap#backend_service}
	BackendService *string `field:"required" json:"backendService" yaml:"backendService"`
}

