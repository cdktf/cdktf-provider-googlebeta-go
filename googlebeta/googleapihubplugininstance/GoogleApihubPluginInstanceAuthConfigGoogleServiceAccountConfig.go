// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleapihubplugininstance


type GoogleApihubPluginInstanceAuthConfigGoogleServiceAccountConfig struct {
	// The service account to be used for authenticating request.
	//
	// The 'iam.serviceAccounts.getAccessToken' permission should be granted on
	// this service account to the impersonator service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.44.0/docs/resources/google_apihub_plugin_instance#service_account GoogleApihubPluginInstance#service_account}
	ServiceAccount *string `field:"required" json:"serviceAccount" yaml:"serviceAccount"`
}

