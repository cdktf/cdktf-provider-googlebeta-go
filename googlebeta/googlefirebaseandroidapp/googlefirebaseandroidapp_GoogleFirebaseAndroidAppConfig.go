package googlefirebaseandroidapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleFirebaseAndroidAppConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firebase_android_app#display_name GoogleFirebaseAndroidApp#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// (Optional) Set to 'ABANDON' to allow the AndroidApp to be untracked from terraform state rather than deleted upon 'terraform destroy'.
	//
	// This is useful because the AndroidApp may be
	// serving traffic. Set to 'DELETE' to delete the AndroidApp. Default to 'DELETE'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firebase_android_app#deletion_policy GoogleFirebaseAndroidApp#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firebase_android_app#id GoogleFirebaseAndroidApp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Immutable. The canonical package name of the Android app as would appear in the Google Play Developer Console.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firebase_android_app#package_name GoogleFirebaseAndroidApp#package_name}
	PackageName *string `field:"optional" json:"packageName" yaml:"packageName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firebase_android_app#project GoogleFirebaseAndroidApp#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_firebase_android_app#timeouts GoogleFirebaseAndroidApp#timeouts}
	Timeouts *GoogleFirebaseAndroidAppTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}
