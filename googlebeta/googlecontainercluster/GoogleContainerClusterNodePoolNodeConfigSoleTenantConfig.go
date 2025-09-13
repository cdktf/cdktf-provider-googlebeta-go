// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlecontainercluster


type GoogleContainerClusterNodePoolNodeConfigSoleTenantConfig struct {
	// node_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.3/docs/resources/google_container_cluster#node_affinity GoogleContainerCluster#node_affinity}
	NodeAffinity interface{} `field:"required" json:"nodeAffinity" yaml:"nodeAffinity"`
	// Specifies the minimum number of vCPUs that each sole tenant node must have to use CPU overcommit.
	//
	// If not specified, the CPU overcommit feature is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.3/docs/resources/google_container_cluster#min_node_cpus GoogleContainerCluster#min_node_cpus}
	MinNodeCpus *float64 `field:"optional" json:"minNodeCpus" yaml:"minNodeCpus"`
}

