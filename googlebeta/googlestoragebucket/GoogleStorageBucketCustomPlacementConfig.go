package googlestoragebucket


type GoogleStorageBucketCustomPlacementConfig struct {
	// The list of individual regions that comprise a dual-region bucket.
	//
	// See the docs for a list of acceptable regions. Note: If any of the data_locations changes, it will recreate the bucket.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_bucket#data_locations GoogleStorageBucket#data_locations}
	DataLocations *[]*string `field:"required" json:"dataLocations" yaml:"dataLocations"`
}

