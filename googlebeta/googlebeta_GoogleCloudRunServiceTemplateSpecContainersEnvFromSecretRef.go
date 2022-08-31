// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleCloudRunServiceTemplateSpecContainersEnvFromSecretRef struct {
	// local_object_reference block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_service#local_object_reference GoogleCloudRunService#local_object_reference}
	LocalObjectReference *GoogleCloudRunServiceTemplateSpecContainersEnvFromSecretRefLocalObjectReference `field:"optional" json:"localObjectReference" yaml:"localObjectReference"`
	// Specify whether the Secret must be defined.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_service#optional GoogleCloudRunService#optional}
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

