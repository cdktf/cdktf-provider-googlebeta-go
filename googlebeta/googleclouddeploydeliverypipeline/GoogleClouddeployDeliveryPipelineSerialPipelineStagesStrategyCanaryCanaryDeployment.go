package googleclouddeploydeliverypipeline


type GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment struct {
	// Required.
	//
	// The percentage based deployments that will occur as a part of a `Rollout`. List is expected in ascending order and each integer n is 0 <= n < 100.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_clouddeploy_delivery_pipeline#percentages GoogleClouddeployDeliveryPipeline#percentages}
	Percentages *[]*float64 `field:"required" json:"percentages" yaml:"percentages"`
	// Whether to run verify tests after each percentage deployment.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_clouddeploy_delivery_pipeline#verify GoogleClouddeployDeliveryPipeline#verify}
	Verify interface{} `field:"optional" json:"verify" yaml:"verify"`
}

