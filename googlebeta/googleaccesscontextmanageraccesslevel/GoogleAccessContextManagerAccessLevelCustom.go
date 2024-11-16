// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleaccesscontextmanageraccesslevel


type GoogleAccessContextManagerAccessLevelCustom struct {
	// expr block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.11.2/docs/resources/google_access_context_manager_access_level#expr GoogleAccessContextManagerAccessLevel#expr}
	Expr *GoogleAccessContextManagerAccessLevelCustomExpr `field:"required" json:"expr" yaml:"expr"`
}

