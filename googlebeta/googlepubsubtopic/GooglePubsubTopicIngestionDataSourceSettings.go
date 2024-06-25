// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlepubsubtopic


type GooglePubsubTopicIngestionDataSourceSettings struct {
	// aws_kinesis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.35.0/docs/resources/google_pubsub_topic#aws_kinesis GooglePubsubTopic#aws_kinesis}
	AwsKinesis *GooglePubsubTopicIngestionDataSourceSettingsAwsKinesis `field:"optional" json:"awsKinesis" yaml:"awsKinesis"`
}

