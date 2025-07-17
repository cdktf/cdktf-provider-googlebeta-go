// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlebeyondcorpapplication


type GoogleBeyondcorpApplicationUpstreams struct {
	// egress_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.44.0/docs/resources/google_beyondcorp_application#egress_policy GoogleBeyondcorpApplication#egress_policy}
	EgressPolicy *GoogleBeyondcorpApplicationUpstreamsEgressPolicy `field:"optional" json:"egressPolicy" yaml:"egressPolicy"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.44.0/docs/resources/google_beyondcorp_application#network GoogleBeyondcorpApplication#network}
	Network *GoogleBeyondcorpApplicationUpstreamsNetwork `field:"optional" json:"network" yaml:"network"`
}

