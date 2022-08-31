// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type DataGoogleIamPolicyAuditConfig struct {
	// audit_log_configs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/d/google_iam_policy#audit_log_configs DataGoogleIamPolicy#audit_log_configs}
	AuditLogConfigs interface{} `field:"required" json:"auditLogConfigs" yaml:"auditLogConfigs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/d/google_iam_policy#service DataGoogleIamPolicy#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
}

