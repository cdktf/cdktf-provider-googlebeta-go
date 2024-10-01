// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googlememorystoreinstance


type GoogleMemorystoreInstancePersistenceConfigAofConfig struct {
	// Optional. The fsync mode.   Possible values:  NEVER EVERY_SEC ALWAYS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.5.0/docs/resources/google_memorystore_instance#append_fsync GoogleMemorystoreInstance#append_fsync}
	AppendFsync *string `field:"optional" json:"appendFsync" yaml:"appendFsync"`
}

