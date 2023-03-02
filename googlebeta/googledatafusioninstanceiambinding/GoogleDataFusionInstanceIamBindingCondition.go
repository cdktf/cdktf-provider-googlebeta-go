package googledatafusioninstanceiambinding


type GoogleDataFusionInstanceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_fusion_instance_iam_binding#expression GoogleDataFusionInstanceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_fusion_instance_iam_binding#title GoogleDataFusionInstanceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_fusion_instance_iam_binding#description GoogleDataFusionInstanceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

