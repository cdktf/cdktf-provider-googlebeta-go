package googledatastoreindex


type GoogleDatastoreIndexProperties struct {
	// The direction the index should optimize for sorting. Possible values: ["ASCENDING", "DESCENDING"].
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastore_index#direction GoogleDatastoreIndex#direction}
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// The property name to index.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_datastore_index#name GoogleDatastoreIndex#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}
