package googlestoragetransferagentpool


type GoogleStorageTransferAgentPoolTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_transfer_agent_pool#create GoogleStorageTransferAgentPool#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_transfer_agent_pool#delete GoogleStorageTransferAgentPool#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_transfer_agent_pool#update GoogleStorageTransferAgentPool#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
