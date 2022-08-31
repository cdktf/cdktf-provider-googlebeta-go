// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleRedisInstanceMaintenancePolicy struct {
	// Optional. Description of what this policy is for. Create/Update methods return INVALID_ARGUMENT if the length is greater than 512.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_redis_instance#description GoogleRedisInstance#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// weekly_maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_redis_instance#weekly_maintenance_window GoogleRedisInstance#weekly_maintenance_window}
	WeeklyMaintenanceWindow interface{} `field:"optional" json:"weeklyMaintenanceWindow" yaml:"weeklyMaintenanceWindow"`
}

