package googlednsrecordset


type GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeo struct {
	// The location name defined in Google Cloud.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dns_record_set#location GoogleDnsRecordSet#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// health_checked_targets block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dns_record_set#health_checked_targets GoogleDnsRecordSet#health_checked_targets}
	HealthCheckedTargets *GoogleDnsRecordSetRoutingPolicyPrimaryBackupBackupGeoHealthCheckedTargets `field:"optional" json:"healthCheckedTargets" yaml:"healthCheckedTargets"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dns_record_set#rrdatas GoogleDnsRecordSet#rrdatas}.
	Rrdatas *[]*string `field:"optional" json:"rrdatas" yaml:"rrdatas"`
}
