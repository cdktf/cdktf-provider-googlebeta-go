// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleosconfigpatchdeployment


type GoogleOsConfigPatchDeploymentInstanceFilterGroupLabels struct {
	// Compute Engine instance labels that must be present for a VM instance to be targeted by this filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.9.0/docs/resources/google_os_config_patch_deployment#labels GoogleOsConfigPatchDeployment#labels}
	Labels *map[string]*string `field:"required" json:"labels" yaml:"labels"`
}

