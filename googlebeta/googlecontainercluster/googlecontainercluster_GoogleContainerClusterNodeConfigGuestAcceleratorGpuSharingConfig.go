package googlecontainercluster


type GoogleContainerClusterNodeConfigGuestAcceleratorGpuSharingConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#gpu_sharing_strategy GoogleContainerCluster#gpu_sharing_strategy}.
	GpuSharingStrategy *string `field:"optional" json:"gpuSharingStrategy" yaml:"gpuSharingStrategy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#max_shared_clients_per_gpu GoogleContainerCluster#max_shared_clients_per_gpu}.
	MaxSharedClientsPerGpu *float64 `field:"optional" json:"maxSharedClientsPerGpu" yaml:"maxSharedClientsPerGpu"`
}
