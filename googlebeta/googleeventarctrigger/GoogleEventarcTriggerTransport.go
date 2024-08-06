// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googleeventarctrigger


type GoogleEventarcTriggerTransport struct {
	// pubsub block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.40.0/docs/resources/google_eventarc_trigger#pubsub GoogleEventarcTrigger#pubsub}
	Pubsub *GoogleEventarcTriggerTransportPubsub `field:"optional" json:"pubsub" yaml:"pubsub"`
}

