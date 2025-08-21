// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledialogflowconversationprofile


type GoogleDialogflowConversationProfileTtsConfigVoice struct {
	// The name of the voice.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.1/docs/resources/google_dialogflow_conversation_profile#name GoogleDialogflowConversationProfile#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The preferred gender of the voice. Possible values: ["SSML_VOICE_GENDER_UNSPECIFIED", "SSML_VOICE_GENDER_MALE", "SSML_VOICE_GENDER_FEMALE", "SSML_VOICE_GENDER_NEUTRAL"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.1/docs/resources/google_dialogflow_conversation_profile#ssml_gender GoogleDialogflowConversationProfile#ssml_gender}
	SsmlGender *string `field:"optional" json:"ssmlGender" yaml:"ssmlGender"`
}

