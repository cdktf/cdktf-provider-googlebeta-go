// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecertificatemanagercertificateissuanceconfig


type GoogleCertificateManagerCertificateIssuanceConfigCertificateAuthorityConfig struct {
	// certificate_authority_service_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.84.0/docs/resources/google_certificate_manager_certificate_issuance_config#certificate_authority_service_config GoogleCertificateManagerCertificateIssuanceConfig#certificate_authority_service_config}
	CertificateAuthorityServiceConfig *GoogleCertificateManagerCertificateIssuanceConfigCertificateAuthorityConfigCertificateAuthorityServiceConfig `field:"optional" json:"certificateAuthorityServiceConfig" yaml:"certificateAuthorityServiceConfig"`
}

