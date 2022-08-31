// Prebuilt google-beta Provider for Terraform CDK (cdktf)
package googlebeta


type GoogleContainerAwsClusterControlPlaneSshConfig struct {
	// The name of the EC2 key pair used to login into cluster machines.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_container_aws_cluster#ec2_key_pair GoogleContainerAwsCluster#ec2_key_pair}
	Ec2KeyPair *string `field:"required" json:"ec2KeyPair" yaml:"ec2KeyPair"`
}

