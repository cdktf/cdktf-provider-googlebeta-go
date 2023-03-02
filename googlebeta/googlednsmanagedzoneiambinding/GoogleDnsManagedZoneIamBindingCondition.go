package googlednsmanagedzoneiambinding


type GoogleDnsManagedZoneIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dns_managed_zone_iam_binding#expression GoogleDnsManagedZoneIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dns_managed_zone_iam_binding#title GoogleDnsManagedZoneIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dns_managed_zone_iam_binding#description GoogleDnsManagedZoneIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

