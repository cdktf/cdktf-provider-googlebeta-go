package googledatastreamstream

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDatastreamStreamConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// destination_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#destination_config GoogleDatastreamStream#destination_config}
	DestinationConfig *GoogleDatastreamStreamDestinationConfig `field:"required" json:"destinationConfig" yaml:"destinationConfig"`
	// Display name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#display_name GoogleDatastreamStream#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The name of the location this stream is located in.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#location GoogleDatastreamStream#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// source_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#source_config GoogleDatastreamStream#source_config}
	SourceConfig *GoogleDatastreamStreamSourceConfig `field:"required" json:"sourceConfig" yaml:"sourceConfig"`
	// The stream identifier.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#stream_id GoogleDatastreamStream#stream_id}
	StreamId *string `field:"required" json:"streamId" yaml:"streamId"`
	// backfill_all block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#backfill_all GoogleDatastreamStream#backfill_all}
	BackfillAll *GoogleDatastreamStreamBackfillAll `field:"optional" json:"backfillAll" yaml:"backfillAll"`
	// backfill_none block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#backfill_none GoogleDatastreamStream#backfill_none}
	BackfillNone *GoogleDatastreamStreamBackfillNone `field:"optional" json:"backfillNone" yaml:"backfillNone"`
	// Desired state of the Stream.
	//
	// Set this field to 'RUNNING' to start the stream, and 'PAUSED' to pause the stream.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#desired_state GoogleDatastreamStream#desired_state}
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#id GoogleDatastreamStream#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#labels GoogleDatastreamStream#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#project GoogleDatastreamStream#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastream_stream#timeouts GoogleDatastreamStream#timeouts}
	Timeouts *GoogleDatastreamStreamTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

