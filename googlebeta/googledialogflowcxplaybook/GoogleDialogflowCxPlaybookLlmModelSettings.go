// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledialogflowcxplaybook


type GoogleDialogflowCxPlaybookLlmModelSettings struct {
	// The selected LLM model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.0/docs/resources/google_dialogflow_cx_playbook#model GoogleDialogflowCxPlaybook#model}
	Model *string `field:"optional" json:"model" yaml:"model"`
	// The custom prompt to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.0/docs/resources/google_dialogflow_cx_playbook#prompt_text GoogleDialogflowCxPlaybook#prompt_text}
	PromptText *string `field:"optional" json:"promptText" yaml:"promptText"`
}

