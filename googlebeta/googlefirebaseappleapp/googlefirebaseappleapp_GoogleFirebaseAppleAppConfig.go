package googlefirebaseappleapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleFirebaseAppleAppConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The user-assigned display name of the App.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firebase_apple_app#display_name GoogleFirebaseAppleApp#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The automatically generated Apple ID assigned to the Apple app by Apple in the Apple App Store.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firebase_apple_app#app_store_id GoogleFirebaseAppleApp#app_store_id}
	AppStoreId *string `field:"optional" json:"appStoreId" yaml:"appStoreId"`
	// The canonical bundle ID of the Apple app as it would appear in the Apple AppStore.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firebase_apple_app#bundle_id GoogleFirebaseAppleApp#bundle_id}
	BundleId *string `field:"optional" json:"bundleId" yaml:"bundleId"`
	// (Optional) Set to 'ABANDON' to allow the AppleApp to be untracked from terraform state rather than deleted upon 'terraform destroy'.
	//
	// This is useful because the AppleApp may be
	// serving traffic. Set to 'DELETE' to delete the AppleApp. Default to 'DELETE'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firebase_apple_app#deletion_policy GoogleFirebaseAppleApp#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firebase_apple_app#id GoogleFirebaseAppleApp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firebase_apple_app#project GoogleFirebaseAppleApp#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The Apple Developer Team ID associated with the App in the App Store.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firebase_apple_app#team_id GoogleFirebaseAppleApp#team_id}
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firebase_apple_app#timeouts GoogleFirebaseAppleApp#timeouts}
	Timeouts *GoogleFirebaseAppleAppTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}
