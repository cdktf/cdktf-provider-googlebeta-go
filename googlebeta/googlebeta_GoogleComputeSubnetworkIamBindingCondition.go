// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleComputeSubnetworkIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_subnetwork_iam_binding#expression GoogleComputeSubnetworkIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_subnetwork_iam_binding#title GoogleComputeSubnetworkIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_subnetwork_iam_binding#description GoogleComputeSubnetworkIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

