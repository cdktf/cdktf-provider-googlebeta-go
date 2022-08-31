// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleAppEngineStandardAppVersionHandlersScript struct {
	// Path to the script from the application root directory.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_app_engine_standard_app_version#script_path GoogleAppEngineStandardAppVersion#script_path}
	ScriptPath *string `field:"required" json:"scriptPath" yaml:"scriptPath"`
}

