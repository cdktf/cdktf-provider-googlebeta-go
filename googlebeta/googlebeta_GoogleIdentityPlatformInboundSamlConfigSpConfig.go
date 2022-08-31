// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleIdentityPlatformInboundSamlConfigSpConfig struct {
	// Callback URI where responses from IDP are handled. Must start with 'https://'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_identity_platform_inbound_saml_config#callback_uri GoogleIdentityPlatformInboundSamlConfig#callback_uri}
	CallbackUri *string `field:"optional" json:"callbackUri" yaml:"callbackUri"`
	// Unique identifier for all SAML entities.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_identity_platform_inbound_saml_config#sp_entity_id GoogleIdentityPlatformInboundSamlConfig#sp_entity_id}
	SpEntityId *string `field:"optional" json:"spEntityId" yaml:"spEntityId"`
}

