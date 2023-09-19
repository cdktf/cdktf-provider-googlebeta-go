// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlesecretmanagersecret


type GoogleSecretManagerSecretReplication struct {
	// auto block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.83.0/docs/resources/google_secret_manager_secret#auto GoogleSecretManagerSecret#auto}
	Auto *GoogleSecretManagerSecretReplicationAuto `field:"optional" json:"auto" yaml:"auto"`
	// The Secret will automatically be replicated without any restrictions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.83.0/docs/resources/google_secret_manager_secret#automatic GoogleSecretManagerSecret#automatic}
	Automatic interface{} `field:"optional" json:"automatic" yaml:"automatic"`
	// user_managed block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.83.0/docs/resources/google_secret_manager_secret#user_managed GoogleSecretManagerSecret#user_managed}
	UserManaged *GoogleSecretManagerSecretReplicationUserManaged `field:"optional" json:"userManaged" yaml:"userManaged"`
}

