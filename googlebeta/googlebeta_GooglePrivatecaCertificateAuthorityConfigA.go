// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GooglePrivatecaCertificateAuthorityConfigA struct {
	// subject_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_certificate_authority#subject_config GooglePrivatecaCertificateAuthority#subject_config}
	SubjectConfig *GooglePrivatecaCertificateAuthorityConfigSubjectConfig `field:"required" json:"subjectConfig" yaml:"subjectConfig"`
	// x509_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_privateca_certificate_authority#x509_config GooglePrivatecaCertificateAuthority#x509_config}
	X509Config *GooglePrivatecaCertificateAuthorityConfigX509Config `field:"required" json:"x509Config" yaml:"x509Config"`
}

