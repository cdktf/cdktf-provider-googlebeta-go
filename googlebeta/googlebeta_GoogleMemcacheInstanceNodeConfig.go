// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleMemcacheInstanceNodeConfig struct {
	// Number of CPUs per node.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_memcache_instance#cpu_count GoogleMemcacheInstance#cpu_count}
	CpuCount *float64 `field:"required" json:"cpuCount" yaml:"cpuCount"`
	// Memory size in Mebibytes for each memcache node.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_memcache_instance#memory_size_mb GoogleMemcacheInstance#memory_size_mb}
	MemorySizeMb *float64 `field:"required" json:"memorySizeMb" yaml:"memorySizeMb"`
}

