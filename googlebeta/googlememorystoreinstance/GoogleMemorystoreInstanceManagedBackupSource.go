// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlememorystoreinstance


type GoogleMemorystoreInstanceManagedBackupSource struct {
	// Example: //memorystore.googleapis.com/projects/{project}/locations/{location}/backups/{backupId}. In this case, it assumes the backup is under memorystore.googleapis.com.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.0/docs/resources/google_memorystore_instance#backup GoogleMemorystoreInstance#backup}
	Backup *string `field:"required" json:"backup" yaml:"backup"`
}

