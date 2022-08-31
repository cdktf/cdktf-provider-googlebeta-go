// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleSecurityScannerScanConfigAuthenticationCustomAccount struct {
	// The login form URL of the website.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_security_scanner_scan_config#login_url GoogleSecurityScannerScanConfig#login_url}
	LoginUrl *string `field:"required" json:"loginUrl" yaml:"loginUrl"`
	// The password of the custom account. The credential is stored encrypted in GCP.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_security_scanner_scan_config#password GoogleSecurityScannerScanConfig#password}
	Password *string `field:"required" json:"password" yaml:"password"`
	// The user name of the custom account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_security_scanner_scan_config#username GoogleSecurityScannerScanConfig#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

