package googlestoragetransferagentpool


type GoogleStorageTransferAgentPoolBandwidthLimit struct {
	// Bandwidth rate in megabytes per second, distributed across all the agents in the pool.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_storage_transfer_agent_pool#limit_mbps GoogleStorageTransferAgentPool#limit_mbps}
	LimitMbps *string `field:"required" json:"limitMbps" yaml:"limitMbps"`
}

