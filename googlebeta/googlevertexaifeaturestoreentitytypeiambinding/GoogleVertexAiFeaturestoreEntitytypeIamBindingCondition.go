package googlevertexaifeaturestoreentitytypeiambinding


type GoogleVertexAiFeaturestoreEntitytypeIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_vertex_ai_featurestore_entitytype_iam_binding#expression GoogleVertexAiFeaturestoreEntitytypeIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_vertex_ai_featurestore_entitytype_iam_binding#title GoogleVertexAiFeaturestoreEntitytypeIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_vertex_ai_featurestore_entitytype_iam_binding#description GoogleVertexAiFeaturestoreEntitytypeIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

