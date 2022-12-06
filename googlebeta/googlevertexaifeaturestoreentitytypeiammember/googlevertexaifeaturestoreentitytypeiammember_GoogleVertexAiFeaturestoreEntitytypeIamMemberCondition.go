package googlevertexaifeaturestoreentitytypeiammember


type GoogleVertexAiFeaturestoreEntitytypeIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_vertex_ai_featurestore_entitytype_iam_member#expression GoogleVertexAiFeaturestoreEntitytypeIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_vertex_ai_featurestore_entitytype_iam_member#title GoogleVertexAiFeaturestoreEntitytypeIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_vertex_ai_featurestore_entitytype_iam_member#description GoogleVertexAiFeaturestoreEntitytypeIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

