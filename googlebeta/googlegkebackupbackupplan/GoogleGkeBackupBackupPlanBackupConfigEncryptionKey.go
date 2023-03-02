package googlegkebackupbackupplan


type GoogleGkeBackupBackupPlanBackupConfigEncryptionKey struct {
	// Google Cloud KMS encryption key. Format: projects/*\/locations/*\/keyRings/*\/cryptoKeys/*.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_gke_backup_backup_plan#gcp_kms_encryption_key GoogleGkeBackupBackupPlan#gcp_kms_encryption_key}
	GcpKmsEncryptionKey *string `field:"required" json:"gcpKmsEncryptionKey" yaml:"gcpKmsEncryptionKey"`
}

