// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlegkebackupbackupplan


type GoogleGkeBackupBackupPlanBackupConfigSelectedNamespaces struct {
	// A list of Kubernetes Namespaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_gke_backup_backup_plan#namespaces GoogleGkeBackupBackupPlan#namespaces}
	Namespaces *[]*string `field:"required" json:"namespaces" yaml:"namespaces"`
}

