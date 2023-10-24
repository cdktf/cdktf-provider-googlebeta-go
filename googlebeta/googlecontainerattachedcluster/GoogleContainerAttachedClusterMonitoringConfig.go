// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainerattachedcluster


type GoogleContainerAttachedClusterMonitoringConfig struct {
	// managed_prometheus_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.3.0/docs/resources/google_container_attached_cluster#managed_prometheus_config GoogleContainerAttachedCluster#managed_prometheus_config}
	ManagedPrometheusConfig *GoogleContainerAttachedClusterMonitoringConfigManagedPrometheusConfig `field:"optional" json:"managedPrometheusConfig" yaml:"managedPrometheusConfig"`
}

