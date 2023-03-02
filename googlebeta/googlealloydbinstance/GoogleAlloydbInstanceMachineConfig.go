package googlealloydbinstance


type GoogleAlloydbInstanceMachineConfig struct {
	// The number of CPU's in the VM instance.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_alloydb_instance#cpu_count GoogleAlloydbInstance#cpu_count}
	CpuCount *float64 `field:"optional" json:"cpuCount" yaml:"cpuCount"`
}

