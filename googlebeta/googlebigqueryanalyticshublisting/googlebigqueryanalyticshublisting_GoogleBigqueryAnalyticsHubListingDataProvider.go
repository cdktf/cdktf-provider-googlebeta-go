package googlebigqueryanalyticshublisting


type GoogleBigqueryAnalyticsHubListingDataProvider struct {
	// Name of the data provider.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_analytics_hub_listing#name GoogleBigqueryAnalyticsHubListing#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Email or URL of the data provider.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_analytics_hub_listing#primary_contact GoogleBigqueryAnalyticsHubListing#primary_contact}
	PrimaryContact *string `field:"optional" json:"primaryContact" yaml:"primaryContact"`
}
