// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterNodeConfigSoleTenantConfigNodeAffinity struct {
	// .
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.14.0/docs/resources/google_container_cluster#key GoogleContainerCluster#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// .
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.14.0/docs/resources/google_container_cluster#operator GoogleContainerCluster#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// .
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.14.0/docs/resources/google_container_cluster#values GoogleContainerCluster#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

