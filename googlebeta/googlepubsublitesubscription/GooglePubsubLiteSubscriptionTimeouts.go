// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlepubsublitesubscription


type GooglePubsubLiteSubscriptionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.17.0/docs/resources/google_pubsub_lite_subscription#create GooglePubsubLiteSubscription#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.17.0/docs/resources/google_pubsub_lite_subscription#delete GooglePubsubLiteSubscription#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.17.0/docs/resources/google_pubsub_lite_subscription#update GooglePubsubLiteSubscription#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

