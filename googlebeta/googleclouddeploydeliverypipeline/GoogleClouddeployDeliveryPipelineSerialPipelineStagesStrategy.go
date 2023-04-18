package googleclouddeploydeliverypipeline


type GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategy struct {
	// canary block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_clouddeploy_delivery_pipeline#canary GoogleClouddeployDeliveryPipeline#canary}
	Canary *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanary `field:"optional" json:"canary" yaml:"canary"`
	// standard block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_clouddeploy_delivery_pipeline#standard GoogleClouddeployDeliveryPipeline#standard}
	Standard *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandard `field:"optional" json:"standard" yaml:"standard"`
}

