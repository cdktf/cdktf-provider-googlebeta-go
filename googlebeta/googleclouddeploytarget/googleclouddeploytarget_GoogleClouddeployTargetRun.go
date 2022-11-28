package googleclouddeploytarget


type GoogleClouddeployTargetRun struct {
	// Required. The location where the Cloud Run Service should be located. Format is `projects/{project}/locations/{location}`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_clouddeploy_target#location GoogleClouddeployTarget#location}
	Location *string `field:"required" json:"location" yaml:"location"`
}

