// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlechroniclewatchlist


type GoogleChronicleWatchlistEntityPopulationMechanism struct {
	// manual block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.42.0/docs/resources/google_chronicle_watchlist#manual GoogleChronicleWatchlist#manual}
	Manual *GoogleChronicleWatchlistEntityPopulationMechanismManual `field:"optional" json:"manual" yaml:"manual"`
}

