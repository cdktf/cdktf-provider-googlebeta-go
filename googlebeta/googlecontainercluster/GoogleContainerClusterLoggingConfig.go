// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterLoggingConfig struct {
	// GKE components exposing logs. Valid values include SYSTEM_COMPONENTS, APISERVER, CONTROLLER_MANAGER, SCHEDULER, and WORKLOADS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.80.0/docs/resources/google_container_cluster#enable_components GoogleContainerCluster#enable_components}
	EnableComponents *[]*string `field:"required" json:"enableComponents" yaml:"enableComponents"`
}

