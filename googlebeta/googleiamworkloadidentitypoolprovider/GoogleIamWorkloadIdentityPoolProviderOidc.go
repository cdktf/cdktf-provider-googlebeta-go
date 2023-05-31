package googleiamworkloadidentitypoolprovider


type GoogleIamWorkloadIdentityPoolProviderOidc struct {
	// The OIDC issuer URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.67.0/docs/resources/google_iam_workload_identity_pool_provider#issuer_uri GoogleIamWorkloadIdentityPoolProvider#issuer_uri}
	IssuerUri *string `field:"required" json:"issuerUri" yaml:"issuerUri"`
	// Acceptable values for the 'aud' field (audience) in the OIDC token.
	//
	// Token exchange
	// requests are rejected if the token audience does not match one of the configured
	// values. Each audience may be at most 256 characters. A maximum of 10 audiences may
	// be configured.
	//
	// If this list is empty, the OIDC token audience must be equal to the full canonical
	// resource name of the WorkloadIdentityPoolProvider, with or without the HTTPS prefix.
	// For example:
	// ```
	// //iam.googleapis.com/projects/<project-number>/locations/<location>/workloadIdentityPools/<pool-id>/providers/<provider-id>
	// https://iam.googleapis.com/projects/<project-number>/locations/<location>/workloadIdentityPools/<pool-id>/providers/<provider-id>
	// ```
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.67.0/docs/resources/google_iam_workload_identity_pool_provider#allowed_audiences GoogleIamWorkloadIdentityPoolProvider#allowed_audiences}
	AllowedAudiences *[]*string `field:"optional" json:"allowedAudiences" yaml:"allowedAudiences"`
}

