// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleSecretManagerSecretTopics struct {
	// The resource name of the Pub/Sub topic that will be published to, in the following format: projects/*\/topics/*.
	//
	// For publication to succeed, the Secret Manager Service Agent service account must have pubsub.publisher permissions on the topic.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_secret_manager_secret#name GoogleSecretManagerSecret#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

