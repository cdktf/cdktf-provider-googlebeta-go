// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledialogflowcxwebhook


type GoogleDialogflowCxWebhookGenericWebServiceSecretVersionsForRequestHeaders struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.1/docs/resources/google_dialogflow_cx_webhook#key GoogleDialogflowCxWebhook#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The SecretManager secret version resource storing the header value. Format: 'projects/{project}/secrets/{secret}/versions/{version}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.1/docs/resources/google_dialogflow_cx_webhook#secret_version GoogleDialogflowCxWebhook#secret_version}
	SecretVersion *string `field:"required" json:"secretVersion" yaml:"secretVersion"`
}

