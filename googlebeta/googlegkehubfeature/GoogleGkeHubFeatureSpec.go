package googlegkehubfeature


type GoogleGkeHubFeatureSpec struct {
	// multiclusteringress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.63.0/docs/resources/google_gke_hub_feature#multiclusteringress GoogleGkeHubFeature#multiclusteringress}
	Multiclusteringress *GoogleGkeHubFeatureSpecMulticlusteringress `field:"optional" json:"multiclusteringress" yaml:"multiclusteringress"`
}

