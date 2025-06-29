// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecloudrunv2job


type GoogleCloudRunV2JobTemplateTemplateContainersResources struct {
	// Only memory and CPU are supported.
	//
	// Use key 'cpu' for CPU limit and 'memory' for memory limit. Note: The only supported values for CPU are '1', '2', '4', and '8'. Setting 4 CPU requires at least 2Gi of memory. The values of the map is string form of the 'quantity' k8s type: https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.41.0/docs/resources/google_cloud_run_v2_job#limits GoogleCloudRunV2Job#limits}
	Limits *map[string]*string `field:"optional" json:"limits" yaml:"limits"`
}

