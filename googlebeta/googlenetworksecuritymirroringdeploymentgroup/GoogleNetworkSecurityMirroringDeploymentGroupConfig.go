// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlenetworksecuritymirroringdeploymentgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkSecurityMirroringDeploymentGroupConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// Resource ID segment making up resource 'name'.
	//
	// It identifies the resource within its parent collection as described in https://google.aip.dev/122. See documentation for resource type 'networksecurity.googleapis.com/MirroringDeploymentGroup'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.23.0/docs/resources/google_network_security_mirroring_deployment_group#location GoogleNetworkSecurityMirroringDeploymentGroup#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Required. Id of the requesting object If auto-generating Id server-side, remove this field and mirroring_deployment_group_id from the method_signature of Create RPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.23.0/docs/resources/google_network_security_mirroring_deployment_group#mirroring_deployment_group_id GoogleNetworkSecurityMirroringDeploymentGroup#mirroring_deployment_group_id}
	MirroringDeploymentGroupId *string `field:"required" json:"mirroringDeploymentGroupId" yaml:"mirroringDeploymentGroupId"`
	// Required. Immutable. The network that is being used for the deployment. Format is: projects/{project}/global/networks/{network}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.23.0/docs/resources/google_network_security_mirroring_deployment_group#network GoogleNetworkSecurityMirroringDeploymentGroup#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.23.0/docs/resources/google_network_security_mirroring_deployment_group#id GoogleNetworkSecurityMirroringDeploymentGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional. Labels as key value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.23.0/docs/resources/google_network_security_mirroring_deployment_group#labels GoogleNetworkSecurityMirroringDeploymentGroup#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.23.0/docs/resources/google_network_security_mirroring_deployment_group#project GoogleNetworkSecurityMirroringDeploymentGroup#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.23.0/docs/resources/google_network_security_mirroring_deployment_group#timeouts GoogleNetworkSecurityMirroringDeploymentGroup#timeouts}
	Timeouts *GoogleNetworkSecurityMirroringDeploymentGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

