// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleiamworkloadidentitypool


type GoogleIamWorkloadIdentityPoolInlineTrustConfigAdditionalTrustBundlesTrustAnchors struct {
	// PEM certificate of the PKI used for validation. Must only contain one ca certificate(either root or intermediate cert).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_iam_workload_identity_pool#pem_certificate GoogleIamWorkloadIdentityPool#pem_certificate}
	PemCertificate *string `field:"required" json:"pemCertificate" yaml:"pemCertificate"`
}

