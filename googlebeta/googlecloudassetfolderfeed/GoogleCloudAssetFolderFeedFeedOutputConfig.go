// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecloudassetfolderfeed


type GoogleCloudAssetFolderFeedFeedOutputConfig struct {
	// pubsub_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.83.0/docs/resources/google_cloud_asset_folder_feed#pubsub_destination GoogleCloudAssetFolderFeed#pubsub_destination}
	PubsubDestination *GoogleCloudAssetFolderFeedFeedOutputConfigPubsubDestination `field:"required" json:"pubsubDestination" yaml:"pubsubDestination"`
}

