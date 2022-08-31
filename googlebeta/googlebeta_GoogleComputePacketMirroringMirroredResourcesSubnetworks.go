// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleComputePacketMirroringMirroredResourcesSubnetworks struct {
	// The URL of the subnetwork where this rule should be active.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_compute_packet_mirroring#url GoogleComputePacketMirroring#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

