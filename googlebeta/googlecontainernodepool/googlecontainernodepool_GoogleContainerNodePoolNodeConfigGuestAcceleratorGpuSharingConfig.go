package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigGuestAcceleratorGpuSharingConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_node_pool#gpu_sharing_strategy GoogleContainerNodePool#gpu_sharing_strategy}.
	GpuSharingStrategy *string `field:"optional" json:"gpuSharingStrategy" yaml:"gpuSharingStrategy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_node_pool#max_shared_clients_per_gpu GoogleContainerNodePool#max_shared_clients_per_gpu}.
	MaxSharedClientsPerGpu *float64 `field:"optional" json:"maxSharedClientsPerGpu" yaml:"maxSharedClientsPerGpu"`
}
