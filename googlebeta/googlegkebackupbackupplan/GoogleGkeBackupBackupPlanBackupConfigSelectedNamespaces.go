package googlegkebackupbackupplan


type GoogleGkeBackupBackupPlanBackupConfigSelectedNamespaces struct {
	// A list of Kubernetes Namespaces.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_backup_backup_plan#namespaces GoogleGkeBackupBackupPlan#namespaces}
	Namespaces *[]*string `field:"required" json:"namespaces" yaml:"namespaces"`
}

