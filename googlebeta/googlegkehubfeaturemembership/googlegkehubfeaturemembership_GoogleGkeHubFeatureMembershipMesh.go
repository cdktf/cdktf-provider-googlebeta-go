package googlegkehubfeaturemembership


type GoogleGkeHubFeatureMembershipMesh struct {
	// Whether to automatically manage Service Mesh. Possible values: MANAGEMENT_UNSPECIFIED, MANAGEMENT_AUTOMATIC, MANAGEMENT_MANUAL.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_hub_feature_membership#management GoogleGkeHubFeatureMembership#management}
	Management *string `field:"optional" json:"management" yaml:"management"`
}

