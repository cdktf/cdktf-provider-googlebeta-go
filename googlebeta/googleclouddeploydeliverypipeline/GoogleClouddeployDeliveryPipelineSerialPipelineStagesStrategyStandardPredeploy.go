// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleclouddeploydeliverypipeline


type GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy struct {
	// Optional. A sequence of skaffold custom actions to invoke during execution of the predeploy job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.46.0/docs/resources/google_clouddeploy_delivery_pipeline#actions GoogleClouddeployDeliveryPipeline#actions}
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
}

