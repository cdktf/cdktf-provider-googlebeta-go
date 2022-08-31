// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GooglePrivatecaCaPoolIssuancePolicy struct {
	// allowed_issuance_modes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_ca_pool#allowed_issuance_modes GooglePrivatecaCaPool#allowed_issuance_modes}
	AllowedIssuanceModes *GooglePrivatecaCaPoolIssuancePolicyAllowedIssuanceModes `field:"optional" json:"allowedIssuanceModes" yaml:"allowedIssuanceModes"`
	// allowed_key_types block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_ca_pool#allowed_key_types GooglePrivatecaCaPool#allowed_key_types}
	AllowedKeyTypes interface{} `field:"optional" json:"allowedKeyTypes" yaml:"allowedKeyTypes"`
	// baseline_values block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_ca_pool#baseline_values GooglePrivatecaCaPool#baseline_values}
	BaselineValues *GooglePrivatecaCaPoolIssuancePolicyBaselineValues `field:"optional" json:"baselineValues" yaml:"baselineValues"`
	// identity_constraints block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_ca_pool#identity_constraints GooglePrivatecaCaPool#identity_constraints}
	IdentityConstraints *GooglePrivatecaCaPoolIssuancePolicyIdentityConstraints `field:"optional" json:"identityConstraints" yaml:"identityConstraints"`
	// The maximum lifetime allowed for issued Certificates.
	//
	// Note that if the issuing CertificateAuthority
	// expires before a Certificate's requested maximumLifetime, the effective lifetime will be explicitly truncated to match it.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_ca_pool#maximum_lifetime GooglePrivatecaCaPool#maximum_lifetime}
	MaximumLifetime *string `field:"optional" json:"maximumLifetime" yaml:"maximumLifetime"`
}

