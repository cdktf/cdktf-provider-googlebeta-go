// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlegkebackuprestoreplan


type GoogleGkeBackupRestorePlanRestoreConfigSelectedApplications struct {
	// namespaced_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.26.0/docs/resources/google_gke_backup_restore_plan#namespaced_names GoogleGkeBackupRestorePlan#namespaced_names}
	NamespacedNames interface{} `field:"required" json:"namespacedNames" yaml:"namespacedNames"`
}

