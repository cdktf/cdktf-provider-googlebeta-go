package googlednsrecordset


type GoogleDnsRecordSetRoutingPolicyWrrHealthCheckedTargets struct {
	// internal_load_balancers block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dns_record_set#internal_load_balancers GoogleDnsRecordSet#internal_load_balancers}
	InternalLoadBalancers interface{} `field:"required" json:"internalLoadBalancers" yaml:"internalLoadBalancers"`
}
