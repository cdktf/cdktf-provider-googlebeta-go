// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleCloudSchedulerJobAppEngineHttpTargetAppEngineRouting struct {
	// App instance. By default, the job is sent to an instance which is available when the job is attempted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_scheduler_job#instance GoogleCloudSchedulerJob#instance}
	Instance *string `field:"optional" json:"instance" yaml:"instance"`
	// App service.
	//
	// By default, the job is sent to the service which is the default service when the job is attempted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_scheduler_job#service GoogleCloudSchedulerJob#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// App version.
	//
	// By default, the job is sent to the version which is the default version when the job is attempted.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_scheduler_job#version GoogleCloudSchedulerJob#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

