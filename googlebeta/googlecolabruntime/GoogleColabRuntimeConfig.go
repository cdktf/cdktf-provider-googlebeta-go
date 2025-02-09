// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecolabruntime

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleColabRuntimeConfig struct {
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
	// Required. The display name of the Runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.19.0/docs/resources/google_colab_runtime#display_name GoogleColabRuntime#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The location for the resource: https://cloud.google.com/colab/docs/locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.19.0/docs/resources/google_colab_runtime#location GoogleColabRuntime#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The user email of the NotebookRuntime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.19.0/docs/resources/google_colab_runtime#runtime_user GoogleColabRuntime#runtime_user}
	RuntimeUser *string `field:"required" json:"runtimeUser" yaml:"runtimeUser"`
	// The description of the Runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.19.0/docs/resources/google_colab_runtime#description GoogleColabRuntime#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.19.0/docs/resources/google_colab_runtime#id GoogleColabRuntime#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The resource name of the Runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.19.0/docs/resources/google_colab_runtime#name GoogleColabRuntime#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// notebook_runtime_template_ref block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.19.0/docs/resources/google_colab_runtime#notebook_runtime_template_ref GoogleColabRuntime#notebook_runtime_template_ref}
	NotebookRuntimeTemplateRef *GoogleColabRuntimeNotebookRuntimeTemplateRef `field:"optional" json:"notebookRuntimeTemplateRef" yaml:"notebookRuntimeTemplateRef"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.19.0/docs/resources/google_colab_runtime#project GoogleColabRuntime#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.19.0/docs/resources/google_colab_runtime#timeouts GoogleColabRuntime#timeouts}
	Timeouts *GoogleColabRuntimeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

