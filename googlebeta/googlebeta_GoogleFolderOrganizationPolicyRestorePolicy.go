// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleFolderOrganizationPolicyRestorePolicy struct {
	// May only be set to true. If set, then the default Policy is restored.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_folder_organization_policy#default GoogleFolderOrganizationPolicy#default}
	Default interface{} `field:"required" json:"default" yaml:"default"`
}

