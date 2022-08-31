// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleDataprocJobScheduling struct {
	// Maximum number of times per hour a driver may be restarted as a result of driver exiting with non-zero code before job is reported failed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_job#max_failures_per_hour GoogleDataprocJob#max_failures_per_hour}
	MaxFailuresPerHour *float64 `field:"required" json:"maxFailuresPerHour" yaml:"maxFailuresPerHour"`
	// Maximum number of times in total a driver may be restarted as a result of driver exiting with non-zero code before job is reported failed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_job#max_failures_total GoogleDataprocJob#max_failures_total}
	MaxFailuresTotal *float64 `field:"required" json:"maxFailuresTotal" yaml:"maxFailuresTotal"`
}

