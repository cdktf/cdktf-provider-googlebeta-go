package googlecloudrunv2service

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleCloudRunV2ServiceConfig struct {
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
	// Name of the Service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#name GoogleCloudRunV2Service#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// template block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#template GoogleCloudRunV2Service#template}
	Template *GoogleCloudRunV2ServiceTemplate `field:"required" json:"template" yaml:"template"`
	// binary_authorization block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#binary_authorization GoogleCloudRunV2Service#binary_authorization}
	BinaryAuthorization *GoogleCloudRunV2ServiceBinaryAuthorization `field:"optional" json:"binaryAuthorization" yaml:"binaryAuthorization"`
	// Arbitrary identifier for the API client.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#client GoogleCloudRunV2Service#client}
	Client *string `field:"optional" json:"client" yaml:"client"`
	// Arbitrary version identifier for the API client.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#client_version GoogleCloudRunV2Service#client_version}
	ClientVersion *string `field:"optional" json:"clientVersion" yaml:"clientVersion"`
	// User-provided description of the Service. This field currently has a 512-character limit.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#description GoogleCloudRunV2Service#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#id GoogleCloudRunV2Service#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Provides the ingress settings for this Service.
	//
	// On output, returns the currently observed ingress settings, or INGRESS_TRAFFIC_UNSPECIFIED if no revision is active. Possible values: ["INGRESS_TRAFFIC_ALL", "INGRESS_TRAFFIC_INTERNAL_ONLY", "INGRESS_TRAFFIC_INTERNAL_LOAD_BALANCER"]
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#ingress GoogleCloudRunV2Service#ingress}
	Ingress *string `field:"optional" json:"ingress" yaml:"ingress"`
	// Map of string keys and values that can be used to organize and categorize objects.
	//
	// User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels Cloud Run will populate some labels with 'run.googleapis.com' or 'serving.knative.dev' namespaces. Those labels are read-only, and user changes will not be preserved.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#labels GoogleCloudRunV2Service#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The launch stage as defined by Google Cloud Platform Launch Stages.
	//
	// Cloud Run supports ALPHA, BETA, and GA. If no value is specified, GA is assumed. Possible values: ["UNIMPLEMENTED", "PRELAUNCH", "EARLY_ACCESS", "ALPHA", "BETA", "GA", "DEPRECATED"]
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#launch_stage GoogleCloudRunV2Service#launch_stage}
	LaunchStage *string `field:"optional" json:"launchStage" yaml:"launchStage"`
	// The location of the cloud run service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#location GoogleCloudRunV2Service#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#project GoogleCloudRunV2Service#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#timeouts GoogleCloudRunV2Service#timeouts}
	Timeouts *GoogleCloudRunV2ServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// traffic block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#traffic GoogleCloudRunV2Service#traffic}
	Traffic interface{} `field:"optional" json:"traffic" yaml:"traffic"`
}
