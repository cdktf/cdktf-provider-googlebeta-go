// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlealloydbcluster


type GoogleAlloydbClusterPscConfig struct {
	// Create an instance that allows connections from Private Service Connect endpoints to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.33.0/docs/resources/google_alloydb_cluster#psc_enabled GoogleAlloydbCluster#psc_enabled}
	PscEnabled interface{} `field:"optional" json:"pscEnabled" yaml:"pscEnabled"`
}

