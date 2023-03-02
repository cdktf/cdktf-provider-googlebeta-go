package googlegkebackupbackupplaniambinding


type GoogleGkeBackupBackupPlanIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_backup_backup_plan_iam_binding#expression GoogleGkeBackupBackupPlanIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_backup_backup_plan_iam_binding#title GoogleGkeBackupBackupPlanIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_backup_backup_plan_iam_binding#description GoogleGkeBackupBackupPlanIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

