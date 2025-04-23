// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledataproccluster


type GoogleDataprocClusterClusterConfigSecurityConfig struct {
	// kerberos_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.31.0/docs/resources/google_dataproc_cluster#kerberos_config GoogleDataprocCluster#kerberos_config}
	KerberosConfig *GoogleDataprocClusterClusterConfigSecurityConfigKerberosConfig `field:"required" json:"kerberosConfig" yaml:"kerberosConfig"`
}

