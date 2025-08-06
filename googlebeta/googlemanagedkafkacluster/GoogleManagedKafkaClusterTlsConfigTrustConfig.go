// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlemanagedkafkacluster


type GoogleManagedKafkaClusterTlsConfigTrustConfig struct {
	// cas_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.47.0/docs/resources/google_managed_kafka_cluster#cas_configs GoogleManagedKafkaCluster#cas_configs}
	CasConfigs interface{} `field:"optional" json:"casConfigs" yaml:"casConfigs"`
}

