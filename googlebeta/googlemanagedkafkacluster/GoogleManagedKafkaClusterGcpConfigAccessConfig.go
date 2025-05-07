// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlemanagedkafkacluster


type GoogleManagedKafkaClusterGcpConfigAccessConfig struct {
	// network_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_managed_kafka_cluster#network_configs GoogleManagedKafkaCluster#network_configs}
	NetworkConfigs interface{} `field:"required" json:"networkConfigs" yaml:"networkConfigs"`
}

