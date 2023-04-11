package googlebigqueryreservation


type GoogleBigqueryReservationAutoscale struct {
	// Number of slots to be scaled when needed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_bigquery_reservation#max_slots GoogleBigqueryReservation#max_slots}
	MaxSlots *float64 `field:"optional" json:"maxSlots" yaml:"maxSlots"`
}

