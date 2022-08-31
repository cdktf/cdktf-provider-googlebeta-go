// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleHealthcareFhirStoreStreamConfigs struct {
	// bigquery_destination block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_healthcare_fhir_store#bigquery_destination GoogleHealthcareFhirStore#bigquery_destination}
	BigqueryDestination *GoogleHealthcareFhirStoreStreamConfigsBigqueryDestination `field:"required" json:"bigqueryDestination" yaml:"bigqueryDestination"`
	// Supply a FHIR resource type (such as "Patient" or "Observation").
	//
	// See
	// https://www.hl7.org/fhir/valueset-resource-types.html for a list of all FHIR resource types. The server treats
	// an empty list as an intent to stream all the supported resource types in this FHIR store.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_healthcare_fhir_store#resource_types GoogleHealthcareFhirStore#resource_types}
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
}

