// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontaineranalysisnote


type GoogleContainerAnalysisNoteAttestationAuthority struct {
	// hint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.80.0/docs/resources/google_container_analysis_note#hint GoogleContainerAnalysisNote#hint}
	Hint *GoogleContainerAnalysisNoteAttestationAuthorityHint `field:"required" json:"hint" yaml:"hint"`
}

