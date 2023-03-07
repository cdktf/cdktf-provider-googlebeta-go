package googleworkstationsworkstationconfig


type GoogleWorkstationsWorkstationConfigPersistentDirectoriesGcePd struct {
	// Type of the disk to use.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_workstations_workstation_config#disk_type GoogleWorkstationsWorkstationConfig#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// Type of file system that the disk should be formatted with.
	//
	// The workstation image must support this file system type. Must be empty if sourceSnapshot is set.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_workstations_workstation_config#fs_type GoogleWorkstationsWorkstationConfig#fs_type}
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// What should happen to the disk after the workstation is deleted. Defaults to DELETE. Possible values: ["RECLAIM_POLICY_UNSPECIFIED", "DELETE", "RETAIN"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_workstations_workstation_config#reclaim_policy GoogleWorkstationsWorkstationConfig#reclaim_policy}
	ReclaimPolicy *string `field:"optional" json:"reclaimPolicy" yaml:"reclaimPolicy"`
	// Size of the disk in GB. Must be empty if sourceSnapshot is set.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_workstations_workstation_config#size_gb GoogleWorkstationsWorkstationConfig#size_gb}
	SizeGb *float64 `field:"optional" json:"sizeGb" yaml:"sizeGb"`
}

