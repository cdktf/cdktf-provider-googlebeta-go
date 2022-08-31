// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleOsConfigPatchDeploymentPatchConfigPostStep struct {
	// linux_exec_step_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_os_config_patch_deployment#linux_exec_step_config GoogleOsConfigPatchDeployment#linux_exec_step_config}
	LinuxExecStepConfig *GoogleOsConfigPatchDeploymentPatchConfigPostStepLinuxExecStepConfig `field:"optional" json:"linuxExecStepConfig" yaml:"linuxExecStepConfig"`
	// windows_exec_step_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_os_config_patch_deployment#windows_exec_step_config GoogleOsConfigPatchDeployment#windows_exec_step_config}
	WindowsExecStepConfig *GoogleOsConfigPatchDeploymentPatchConfigPostStepWindowsExecStepConfig `field:"optional" json:"windowsExecStepConfig" yaml:"windowsExecStepConfig"`
}

