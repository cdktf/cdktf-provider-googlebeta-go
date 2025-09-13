// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigKubeletConfigEvictionMinimumReclaim struct {
	// Defines percentage of minimum reclaim for imagefs.available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.3/docs/resources/google_container_cluster#imagefs_available GoogleContainerCluster#imagefs_available}
	ImagefsAvailable *string `field:"optional" json:"imagefsAvailable" yaml:"imagefsAvailable"`
	// Defines percentage of minimum reclaim for imagefs.inodesFree.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.3/docs/resources/google_container_cluster#imagefs_inodes_free GoogleContainerCluster#imagefs_inodes_free}
	ImagefsInodesFree *string `field:"optional" json:"imagefsInodesFree" yaml:"imagefsInodesFree"`
	// Defines percentage of minimum reclaim for memory.available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.3/docs/resources/google_container_cluster#memory_available GoogleContainerCluster#memory_available}
	MemoryAvailable *string `field:"optional" json:"memoryAvailable" yaml:"memoryAvailable"`
	// Defines percentage of minimum reclaim for nodefs.available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.3/docs/resources/google_container_cluster#nodefs_available GoogleContainerCluster#nodefs_available}
	NodefsAvailable *string `field:"optional" json:"nodefsAvailable" yaml:"nodefsAvailable"`
	// Defines percentage of minimum reclaim for nodefs.inodesFree.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.3/docs/resources/google_container_cluster#nodefs_inodes_free GoogleContainerCluster#nodefs_inodes_free}
	NodefsInodesFree *string `field:"optional" json:"nodefsInodesFree" yaml:"nodefsInodesFree"`
	// Defines percentage of minimum reclaim for pid.available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.3/docs/resources/google_container_cluster#pid_available GoogleContainerCluster#pid_available}
	PidAvailable *string `field:"optional" json:"pidAvailable" yaml:"pidAvailable"`
}

