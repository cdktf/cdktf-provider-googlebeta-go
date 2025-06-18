// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledatabasemigrationservicemigrationjob


type GoogleDatabaseMigrationServiceMigrationJobVpcPeeringConnectivity struct {
	// The name of the VPC network to peer with the Cloud SQL private network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.40.0/docs/resources/google_database_migration_service_migration_job#vpc GoogleDatabaseMigrationServiceMigrationJob#vpc}
	Vpc *string `field:"optional" json:"vpc" yaml:"vpc"`
}

