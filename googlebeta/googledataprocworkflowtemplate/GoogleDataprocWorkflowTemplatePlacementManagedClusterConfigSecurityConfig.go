// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledataprocworkflowtemplate


type GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfig struct {
	// kerberos_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.11.1/docs/resources/google_dataproc_workflow_template#kerberos_config GoogleDataprocWorkflowTemplate#kerberos_config}
	KerberosConfig *GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig `field:"optional" json:"kerberosConfig" yaml:"kerberosConfig"`
}

