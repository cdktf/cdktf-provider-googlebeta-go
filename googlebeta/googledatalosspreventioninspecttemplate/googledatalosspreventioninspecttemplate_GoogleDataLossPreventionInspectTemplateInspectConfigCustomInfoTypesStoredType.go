package googledatalosspreventioninspecttemplate


type GoogleDataLossPreventionInspectTemplateInspectConfigCustomInfoTypesStoredType struct {
	// Resource name of the requested StoredInfoType, for example 'organizations/433245324/storedInfoTypes/432452342' or 'projects/project-id/storedInfoTypes/432452342'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_data_loss_prevention_inspect_template#name GoogleDataLossPreventionInspectTemplate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}
