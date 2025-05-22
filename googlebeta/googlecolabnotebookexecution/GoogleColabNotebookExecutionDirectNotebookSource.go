// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecolabnotebookexecution


type GoogleColabNotebookExecutionDirectNotebookSource struct {
	// The base64-encoded contents of the input notebook file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_colab_notebook_execution#content GoogleColabNotebookExecution#content}
	Content *string `field:"required" json:"content" yaml:"content"`
}

