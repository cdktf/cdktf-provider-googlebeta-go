package googleclouddeploytarget


type GoogleClouddeployTargetMultiTarget struct {
	// Required. The target_ids of this multiTarget.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_clouddeploy_target#target_ids GoogleClouddeployTarget#target_ids}
	TargetIds *[]*string `field:"required" json:"targetIds" yaml:"targetIds"`
}

