// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleContainerClusterCostManagementConfig struct {
	// Whether to enable GKE cost allocation.
	//
	// When you enable GKE cost allocation, the cluster name and namespace of your GKE workloads appear in the labels field of the billing export to BigQuery. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_cluster#enabled GoogleContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

