package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJobActionsPubSub struct {
	// Cloud Pub/Sub topic to send notifications to.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_job_trigger#topic GoogleDataLossPreventionJobTrigger#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
}

