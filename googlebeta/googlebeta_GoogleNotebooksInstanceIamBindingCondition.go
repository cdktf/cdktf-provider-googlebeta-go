// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleNotebooksInstanceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_notebooks_instance_iam_binding#expression GoogleNotebooksInstanceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_notebooks_instance_iam_binding#title GoogleNotebooksInstanceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_notebooks_instance_iam_binding#description GoogleNotebooksInstanceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

