// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package googledialogflowconversationprofile


type GoogleDialogflowConversationProfileHumanAgentAssistantConfigEndUserSuggestionConfigFeatureConfigsQueryConfigKnowledgeBaseQuerySource struct {
	// Knowledge bases to query. Format: projects/<Project ID>/locations/<Location ID>/knowledgeBases/<Knowledge Base ID>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.49.1/docs/resources/google_dialogflow_conversation_profile#knowledge_bases GoogleDialogflowConversationProfile#knowledge_bases}
	KnowledgeBases *[]*string `field:"required" json:"knowledgeBases" yaml:"knowledgeBases"`
}

