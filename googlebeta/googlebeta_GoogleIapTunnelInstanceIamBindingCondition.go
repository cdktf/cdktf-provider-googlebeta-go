// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleIapTunnelInstanceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_tunnel_instance_iam_binding#expression GoogleIapTunnelInstanceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_tunnel_instance_iam_binding#title GoogleIapTunnelInstanceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_iap_tunnel_instance_iam_binding#description GoogleIapTunnelInstanceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

