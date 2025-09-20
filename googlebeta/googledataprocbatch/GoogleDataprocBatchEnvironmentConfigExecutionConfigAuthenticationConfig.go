// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledataprocbatch


type GoogleDataprocBatchEnvironmentConfigExecutionConfigAuthenticationConfig struct {
	// Authentication type for the user workload running in containers. Possible values: ["SERVICE_ACCOUNT", "END_USER_CREDENTIALS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.50.0/docs/resources/google_dataproc_batch#user_workload_authentication_type GoogleDataprocBatch#user_workload_authentication_type}
	UserWorkloadAuthenticationType *string `field:"optional" json:"userWorkloadAuthenticationType" yaml:"userWorkloadAuthenticationType"`
}

