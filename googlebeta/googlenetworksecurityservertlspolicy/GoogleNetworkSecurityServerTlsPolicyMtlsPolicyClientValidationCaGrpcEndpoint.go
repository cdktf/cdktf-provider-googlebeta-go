// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlenetworksecurityservertlspolicy


type GoogleNetworkSecurityServerTlsPolicyMtlsPolicyClientValidationCaGrpcEndpoint struct {
	// The target URI of the gRPC endpoint. Only UDS path is supported, and should start with "unix:".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.81.0/docs/resources/google_network_security_server_tls_policy#target_uri GoogleNetworkSecurityServerTlsPolicy#target_uri}
	TargetUri *string `field:"required" json:"targetUri" yaml:"targetUri"`
}

