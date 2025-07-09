// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleprivatecacapool


type GooglePrivatecaCaPoolIssuancePolicyAllowedIssuanceModes struct {
	// When true, allows callers to create Certificates by specifying a CertificateConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.43.0/docs/resources/google_privateca_ca_pool#allow_config_based_issuance GooglePrivatecaCaPool#allow_config_based_issuance}
	AllowConfigBasedIssuance interface{} `field:"required" json:"allowConfigBasedIssuance" yaml:"allowConfigBasedIssuance"`
	// When true, allows callers to create Certificates by specifying a CSR.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.43.0/docs/resources/google_privateca_ca_pool#allow_csr_based_issuance GooglePrivatecaCaPool#allow_csr_based_issuance}
	AllowCsrBasedIssuance interface{} `field:"required" json:"allowCsrBasedIssuance" yaml:"allowCsrBasedIssuance"`
}

