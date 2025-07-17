// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledataprocworkflowtemplate


type GoogleDataprocWorkflowTemplateJobsSparkRJobLoggingConfig struct {
	// The per-package log levels for the driver.
	//
	// This may include "root" package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.44.0/docs/resources/google_dataproc_workflow_template#driver_log_levels GoogleDataprocWorkflowTemplate#driver_log_levels}
	DriverLogLevels *map[string]*string `field:"optional" json:"driverLogLevels" yaml:"driverLogLevels"`
}

