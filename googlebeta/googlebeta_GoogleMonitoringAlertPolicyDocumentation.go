// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleMonitoringAlertPolicyDocumentation struct {
	// The text of the documentation, interpreted according to mimeType.
	//
	// The content may not exceed 8,192 Unicode characters and may not
	// exceed more than 10,240 bytes when encoded in UTF-8 format,
	// whichever is smaller.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_monitoring_alert_policy#content GoogleMonitoringAlertPolicy#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
	// The format of the content field. Presently, only the value "text/markdown" is supported.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_monitoring_alert_policy#mime_type GoogleMonitoringAlertPolicy#mime_type}
	MimeType *string `field:"optional" json:"mimeType" yaml:"mimeType"`
}

