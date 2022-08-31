// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleIdentityPlatformInboundSamlConfigIdpConfig struct {
	// idp_certificates block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_identity_platform_inbound_saml_config#idp_certificates GoogleIdentityPlatformInboundSamlConfig#idp_certificates}
	IdpCertificates interface{} `field:"required" json:"idpCertificates" yaml:"idpCertificates"`
	// Unique identifier for all SAML entities.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_identity_platform_inbound_saml_config#idp_entity_id GoogleIdentityPlatformInboundSamlConfig#idp_entity_id}
	IdpEntityId *string `field:"required" json:"idpEntityId" yaml:"idpEntityId"`
	// URL to send Authentication request to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_identity_platform_inbound_saml_config#sso_url GoogleIdentityPlatformInboundSamlConfig#sso_url}
	SsoUrl *string `field:"required" json:"ssoUrl" yaml:"ssoUrl"`
	// Indicates if outbounding SAMLRequest should be signed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_identity_platform_inbound_saml_config#sign_request GoogleIdentityPlatformInboundSamlConfig#sign_request}
	SignRequest interface{} `field:"optional" json:"signRequest" yaml:"signRequest"`
}

