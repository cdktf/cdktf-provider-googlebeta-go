package googleclouddeploydeliverypipeline


type GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetes struct {
	// gateway_service_mesh block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_clouddeploy_delivery_pipeline#gateway_service_mesh GoogleClouddeployDeliveryPipeline#gateway_service_mesh}
	GatewayServiceMesh *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesGatewayServiceMesh `field:"optional" json:"gatewayServiceMesh" yaml:"gatewayServiceMesh"`
	// service_networking block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_clouddeploy_delivery_pipeline#service_networking GoogleClouddeployDeliveryPipeline#service_networking}
	ServiceNetworking *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking `field:"optional" json:"serviceNetworking" yaml:"serviceNetworking"`
}

