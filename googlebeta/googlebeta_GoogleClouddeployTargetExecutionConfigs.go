// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleClouddeployTargetExecutionConfigs struct {
	// Required. Usages when this configuration should be applied.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_clouddeploy_target#usages GoogleClouddeployTarget#usages}
	Usages *[]*string `field:"required" json:"usages" yaml:"usages"`
	// Optional.
	//
	// Cloud Storage location in which to store execution outputs. This can either be a bucket ("gs://my-bucket") or a path within a bucket ("gs://my-bucket/my-dir"). If unspecified, a default bucket located in the same region will be used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_clouddeploy_target#artifact_storage GoogleClouddeployTarget#artifact_storage}
	ArtifactStorage *string `field:"optional" json:"artifactStorage" yaml:"artifactStorage"`
	// Optional. Google service account to use for execution. If unspecified, the project execution service account (-compute@developer.gserviceaccount.com) is used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_clouddeploy_target#service_account GoogleClouddeployTarget#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// Optional.
	//
	// The resource name of the `WorkerPool`, with the format `projects/{project}/locations/{location}/workerPools/{worker_pool}`. If this optional field is unspecified, the default Cloud Build pool will be used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_clouddeploy_target#worker_pool GoogleClouddeployTarget#worker_pool}
	WorkerPool *string `field:"optional" json:"workerPool" yaml:"workerPool"`
}

