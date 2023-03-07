package googleworkstationsworkstationconfig


type GoogleWorkstationsWorkstationConfigHostGceInstanceShieldedInstanceConfig struct {
	// Whether the instance has integrity monitoring enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_workstations_workstation_config#enable_integrity_monitoring GoogleWorkstationsWorkstationConfig#enable_integrity_monitoring}
	EnableIntegrityMonitoring interface{} `field:"optional" json:"enableIntegrityMonitoring" yaml:"enableIntegrityMonitoring"`
	// Whether the instance has Secure Boot enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_workstations_workstation_config#enable_secure_boot GoogleWorkstationsWorkstationConfig#enable_secure_boot}
	EnableSecureBoot interface{} `field:"optional" json:"enableSecureBoot" yaml:"enableSecureBoot"`
	// Whether the instance has the vTPM enabled.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_workstations_workstation_config#enable_vtpm GoogleWorkstationsWorkstationConfig#enable_vtpm}
	EnableVtpm interface{} `field:"optional" json:"enableVtpm" yaml:"enableVtpm"`
}

