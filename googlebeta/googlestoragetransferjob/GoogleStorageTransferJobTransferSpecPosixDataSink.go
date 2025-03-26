// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlestoragetransferjob


type GoogleStorageTransferJobTransferSpecPosixDataSink struct {
	// Root directory path to the filesystem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.27.0/docs/resources/google_storage_transfer_job#root_directory GoogleStorageTransferJob#root_directory}
	RootDirectory *string `field:"required" json:"rootDirectory" yaml:"rootDirectory"`
}

