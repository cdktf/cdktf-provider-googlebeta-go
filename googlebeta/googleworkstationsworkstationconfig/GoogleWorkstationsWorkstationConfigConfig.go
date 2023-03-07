package googleworkstationsworkstationconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleWorkstationsWorkstationConfigConfig struct {
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
	// The location where the workstation cluster config should reside.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_workstations_workstation_config#location GoogleWorkstationsWorkstationConfig#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the workstation cluster.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_workstations_workstation_config#workstation_cluster_id GoogleWorkstationsWorkstationConfig#workstation_cluster_id}
	WorkstationClusterId *string `field:"required" json:"workstationClusterId" yaml:"workstationClusterId"`
	// The ID of the workstation cluster config.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_workstations_workstation_config#workstation_config_id GoogleWorkstationsWorkstationConfig#workstation_config_id}
	WorkstationConfigId *string `field:"required" json:"workstationConfigId" yaml:"workstationConfigId"`
	// Client-specified annotations. This is distinct from labels.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_workstations_workstation_config#annotations GoogleWorkstationsWorkstationConfig#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// container block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_workstations_workstation_config#container GoogleWorkstationsWorkstationConfig#container}
	Container *GoogleWorkstationsWorkstationConfigContainer `field:"optional" json:"container" yaml:"container"`
	// Human-readable name for this resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_workstations_workstation_config#display_name GoogleWorkstationsWorkstationConfig#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_workstations_workstation_config#encryption_key GoogleWorkstationsWorkstationConfig#encryption_key}
	EncryptionKey *GoogleWorkstationsWorkstationConfigEncryptionKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// host block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_workstations_workstation_config#host GoogleWorkstationsWorkstationConfig#host}
	Host *GoogleWorkstationsWorkstationConfigHost `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_workstations_workstation_config#id GoogleWorkstationsWorkstationConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_workstations_workstation_config#labels GoogleWorkstationsWorkstationConfig#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// persistent_directories block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_workstations_workstation_config#persistent_directories GoogleWorkstationsWorkstationConfig#persistent_directories}
	PersistentDirectories interface{} `field:"optional" json:"persistentDirectories" yaml:"persistentDirectories"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_workstations_workstation_config#project GoogleWorkstationsWorkstationConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_workstations_workstation_config#timeouts GoogleWorkstationsWorkstationConfig#timeouts}
	Timeouts *GoogleWorkstationsWorkstationConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

