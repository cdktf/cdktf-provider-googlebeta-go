package googlecertificatemanagercertificate


type GoogleCertificateManagerCertificateSelfManaged struct {
	// **Deprecated** The certificate chain in PEM-encoded form.
	//
	// Leaf certificate comes first, followed by intermediate ones if any.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_certificate_manager_certificate#certificate_pem GoogleCertificateManagerCertificate#certificate_pem}
	CertificatePem *string `field:"optional" json:"certificatePem" yaml:"certificatePem"`
	// The certificate chain in PEM-encoded form.
	//
	// Leaf certificate comes first, followed by intermediate ones if any.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_certificate_manager_certificate#pem_certificate GoogleCertificateManagerCertificate#pem_certificate}
	PemCertificate *string `field:"optional" json:"pemCertificate" yaml:"pemCertificate"`
	// The private key of the leaf certificate in PEM-encoded form.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_certificate_manager_certificate#pem_private_key GoogleCertificateManagerCertificate#pem_private_key}
	PemPrivateKey *string `field:"optional" json:"pemPrivateKey" yaml:"pemPrivateKey"`
	// **Deprecated** The private key of the leaf certificate in PEM-encoded form.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_certificate_manager_certificate#private_key_pem GoogleCertificateManagerCertificate#private_key_pem}
	PrivateKeyPem *string `field:"optional" json:"privateKeyPem" yaml:"privateKeyPem"`
}
