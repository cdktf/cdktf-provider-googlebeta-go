package googlecloudrunservice


type GoogleCloudRunServiceTemplateSpecContainersStartupProbeGrpc struct {
	// Port number to access on the container. Number must be in the range 1 to 65535.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_service#port GoogleCloudRunService#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloud_run_service#service GoogleCloudRunService#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}
