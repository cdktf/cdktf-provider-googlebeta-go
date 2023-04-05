package datagooglefirebaseappleappconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGoogleFirebaseAppleAppConfigAConfig struct {
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
	// The id of the Firebase iOS App.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/d/google_firebase_apple_app_config#app_id DataGoogleFirebaseAppleAppConfigA#app_id}
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// The project id of the Firebase iOS App.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/d/google_firebase_apple_app_config#project DataGoogleFirebaseAppleAppConfigA#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

