package googlefirebasehostingversion


type GoogleFirebaseHostingVersionConfigRewritesRun struct {
	// User-defined ID of the Cloud Run service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firebase_hosting_version#service_id GoogleFirebaseHostingVersion#service_id}
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
	// Optional. User-provided region where the Cloud Run service is hosted. Defaults to 'us-central1' if not supplied.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firebase_hosting_version#region GoogleFirebaseHostingVersion#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

