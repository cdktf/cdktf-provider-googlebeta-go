package googlestoragebucket


type GoogleStorageBucketAutoclass struct {
	// While set to true, autoclass automatically transitions objects in your bucket to appropriate storage classes based on each object's access pattern.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket#enabled GoogleStorageBucket#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

