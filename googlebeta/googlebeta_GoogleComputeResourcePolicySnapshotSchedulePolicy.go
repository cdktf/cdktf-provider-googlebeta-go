// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleComputeResourcePolicySnapshotSchedulePolicy struct {
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_resource_policy#schedule GoogleComputeResourcePolicy#schedule}
	Schedule *GoogleComputeResourcePolicySnapshotSchedulePolicySchedule `field:"required" json:"schedule" yaml:"schedule"`
	// retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_resource_policy#retention_policy GoogleComputeResourcePolicy#retention_policy}
	RetentionPolicy *GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy `field:"optional" json:"retentionPolicy" yaml:"retentionPolicy"`
	// snapshot_properties block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_resource_policy#snapshot_properties GoogleComputeResourcePolicy#snapshot_properties}
	SnapshotProperties *GoogleComputeResourcePolicySnapshotSchedulePolicySnapshotProperties `field:"optional" json:"snapshotProperties" yaml:"snapshotProperties"`
}

