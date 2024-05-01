// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlealloydbcluster


type GoogleAlloydbClusterAutomatedBackupPolicyQuantityBasedRetention struct {
	// The number of backups to retain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.27.0/docs/resources/google_alloydb_cluster#count GoogleAlloydbCluster#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
}

