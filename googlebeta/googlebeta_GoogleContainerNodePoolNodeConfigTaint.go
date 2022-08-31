// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleContainerNodePoolNodeConfigTaint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_node_pool#effect GoogleContainerNodePool#effect}.
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_node_pool#key GoogleContainerNodePool#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_node_pool#value GoogleContainerNodePool#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

