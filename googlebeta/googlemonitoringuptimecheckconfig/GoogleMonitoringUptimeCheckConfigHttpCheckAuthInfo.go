// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlemonitoringuptimecheckconfig


type GoogleMonitoringUptimeCheckConfigHttpCheckAuthInfo struct {
	// The username to authenticate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.0/docs/resources/google_monitoring_uptime_check_config#username GoogleMonitoringUptimeCheckConfig#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// The password to authenticate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.0/docs/resources/google_monitoring_uptime_check_config#password GoogleMonitoringUptimeCheckConfig#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The password to authenticate.
	//
	// Note: This property is write-only and will not be read from the API. For more info see [updating write-only attributes](/docs/providers/google/guides/using_write_only_attributes.html#updating-write-only-attributes)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.0/docs/resources/google_monitoring_uptime_check_config#password_wo GoogleMonitoringUptimeCheckConfig#password_wo}
	PasswordWo *string `field:"optional" json:"passwordWo" yaml:"passwordWo"`
	// Triggers update of password_wo write-only. For more info see [updating write-only attributes](/docs/providers/google/guides/using_write_only_attributes.html#updating-write-only-attributes).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.0/docs/resources/google_monitoring_uptime_check_config#password_wo_version GoogleMonitoringUptimeCheckConfig#password_wo_version}
	PasswordWoVersion *string `field:"optional" json:"passwordWoVersion" yaml:"passwordWoVersion"`
}

