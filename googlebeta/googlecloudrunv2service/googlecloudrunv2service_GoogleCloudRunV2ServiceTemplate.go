package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplate struct {
	// containers block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#containers GoogleCloudRunV2Service#containers}
	Containers interface{} `field:"optional" json:"containers" yaml:"containers"`
	// A reference to a customer managed encryption key (CMEK) to use to encrypt this container image.
	//
	// For more information, go to https://cloud.google.com/run/docs/securing/using-cmek
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#encryption_key GoogleCloudRunV2Service#encryption_key}
	EncryptionKey *string `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// The sandbox environment to host this Revision. Possible values: ["EXECUTION_ENVIRONMENT_GEN1", "EXECUTION_ENVIRONMENT_GEN2"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#execution_environment GoogleCloudRunV2Service#execution_environment}
	ExecutionEnvironment *string `field:"optional" json:"executionEnvironment" yaml:"executionEnvironment"`
	// KRM-style labels for the resource.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#labels GoogleCloudRunV2Service#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Sets the maximum number of requests that each serving instance can receive.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#max_instance_request_concurrency GoogleCloudRunV2Service#max_instance_request_concurrency}
	MaxInstanceRequestConcurrency *float64 `field:"optional" json:"maxInstanceRequestConcurrency" yaml:"maxInstanceRequestConcurrency"`
	// The unique name for the revision.
	//
	// If this field is omitted, it will be automatically generated based on the Service name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#revision GoogleCloudRunV2Service#revision}
	Revision *string `field:"optional" json:"revision" yaml:"revision"`
	// scaling block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#scaling GoogleCloudRunV2Service#scaling}
	Scaling *GoogleCloudRunV2ServiceTemplateScaling `field:"optional" json:"scaling" yaml:"scaling"`
	// Email address of the IAM service account associated with the revision of the service.
	//
	// The service account represents the identity of the running revision, and determines what permissions the revision has. If not provided, the revision will use the project's default service account.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#service_account GoogleCloudRunV2Service#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// Max allowed time for an instance to respond to a request.
	//
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#timeout GoogleCloudRunV2Service#timeout}
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
	// volumes block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#volumes GoogleCloudRunV2Service#volumes}
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
	// vpc_access block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_v2_service#vpc_access GoogleCloudRunV2Service#vpc_access}
	VpcAccess *GoogleCloudRunV2ServiceTemplateVpcAccess `field:"optional" json:"vpcAccess" yaml:"vpcAccess"`
}
