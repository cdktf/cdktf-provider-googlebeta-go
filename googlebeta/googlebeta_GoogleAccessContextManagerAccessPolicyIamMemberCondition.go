// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleAccessContextManagerAccessPolicyIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_access_context_manager_access_policy_iam_member#expression GoogleAccessContextManagerAccessPolicyIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_access_context_manager_access_policy_iam_member#title GoogleAccessContextManagerAccessPolicyIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_access_context_manager_access_policy_iam_member#description GoogleAccessContextManagerAccessPolicyIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

