package googlecomputeregioninstancetemplate


type GoogleComputeRegionInstanceTemplateDiskDiskEncryptionKey struct {
	// The self link of the encryption key that is stored in Google Cloud KMS.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_region_instance_template#kms_key_self_link GoogleComputeRegionInstanceTemplate#kms_key_self_link}
	KmsKeySelfLink *string `field:"required" json:"kmsKeySelfLink" yaml:"kmsKeySelfLink"`
}

