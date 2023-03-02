package googlegkebackupbackupplaniammember


type GoogleGkeBackupBackupPlanIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_backup_backup_plan_iam_member#expression GoogleGkeBackupBackupPlanIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_backup_backup_plan_iam_member#title GoogleGkeBackupBackupPlanIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_backup_backup_plan_iam_member#description GoogleGkeBackupBackupPlanIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

