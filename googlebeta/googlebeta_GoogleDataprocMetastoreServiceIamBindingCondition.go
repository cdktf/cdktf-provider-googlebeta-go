// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleDataprocMetastoreServiceIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_metastore_service_iam_binding#expression GoogleDataprocMetastoreServiceIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_metastore_service_iam_binding#title GoogleDataprocMetastoreServiceIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_metastore_service_iam_binding#description GoogleDataprocMetastoreServiceIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

