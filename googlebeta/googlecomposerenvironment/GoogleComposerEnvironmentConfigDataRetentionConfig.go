// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecomposerenvironment


type GoogleComposerEnvironmentConfigDataRetentionConfig struct {
	// task_logs_retention_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.37.0/docs/resources/google_composer_environment#task_logs_retention_config GoogleComposerEnvironment#task_logs_retention_config}
	TaskLogsRetentionConfig interface{} `field:"required" json:"taskLogsRetentionConfig" yaml:"taskLogsRetentionConfig"`
}

