package googlevertexaiindex


type GoogleVertexAiIndexMetadataConfigAlgorithmConfig struct {
	// brute_force_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_vertex_ai_index#brute_force_config GoogleVertexAiIndex#brute_force_config}
	BruteForceConfig *GoogleVertexAiIndexMetadataConfigAlgorithmConfigBruteForceConfig `field:"optional" json:"bruteForceConfig" yaml:"bruteForceConfig"`
	// tree_ah_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_vertex_ai_index#tree_ah_config GoogleVertexAiIndex#tree_ah_config}
	TreeAhConfig *GoogleVertexAiIndexMetadataConfigAlgorithmConfigTreeAhConfig `field:"optional" json:"treeAhConfig" yaml:"treeAhConfig"`
}
