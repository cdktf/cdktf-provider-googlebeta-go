package googlecloudbuildtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-googlebeta-go/googlebeta/v3/googlecloudbuildtrigger/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudbuild_trigger google_cloudbuild_trigger}.
type GoogleCloudbuildTrigger interface {
	cdktf.TerraformResource
	ApprovalConfig() GoogleCloudbuildTriggerApprovalConfigOutputReference
	ApprovalConfigInput() *GoogleCloudbuildTriggerApprovalConfig
	BuildAttribute() GoogleCloudbuildTriggerBuildOutputReference
	BuildAttributeInput() *GoogleCloudbuildTriggerBuild
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
	Filename() *string
	SetFilename(val *string)
	FilenameInput() *string
	Filter() *string
	SetFilter(val *string)
	FilterInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GitFileSource() GoogleCloudbuildTriggerGitFileSourceOutputReference
	GitFileSourceInput() *GoogleCloudbuildTriggerGitFileSource
	Github() GoogleCloudbuildTriggerGithubOutputReference
	GithubInput() *GoogleCloudbuildTriggerGithub
	Id() *string
	SetId(val *string)
	IdInput() *string
	IgnoredFiles() *[]*string
	SetIgnoredFiles(val *[]*string)
	IgnoredFilesInput() *[]*string
	IncludeBuildLogs() *string
	SetIncludeBuildLogs(val *string)
	IncludeBuildLogsInput() *string
	IncludedFiles() *[]*string
	SetIncludedFiles(val *[]*string)
	IncludedFilesInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PubsubConfig() GoogleCloudbuildTriggerPubsubConfigOutputReference
	PubsubConfigInput() *GoogleCloudbuildTriggerPubsubConfig
	// Experimental.
	RawOverrides() interface{}
	ServiceAccount() *string
	SetServiceAccount(val *string)
	ServiceAccountInput() *string
	SourceToBuild() GoogleCloudbuildTriggerSourceToBuildOutputReference
	SourceToBuildInput() *GoogleCloudbuildTriggerSourceToBuild
	Substitutions() *map[string]*string
	SetSubstitutions(val *map[string]*string)
	SubstitutionsInput() *map[string]*string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleCloudbuildTriggerTimeoutsOutputReference
	TimeoutsInput() interface{}
	TriggerId() *string
	TriggerTemplate() GoogleCloudbuildTriggerTriggerTemplateOutputReference
	TriggerTemplateInput() *GoogleCloudbuildTriggerTriggerTemplate
	WebhookConfig() GoogleCloudbuildTriggerWebhookConfigOutputReference
	WebhookConfigInput() *GoogleCloudbuildTriggerWebhookConfig
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutApprovalConfig(value *GoogleCloudbuildTriggerApprovalConfig)
	PutBuildAttribute(value *GoogleCloudbuildTriggerBuild)
	PutGitFileSource(value *GoogleCloudbuildTriggerGitFileSource)
	PutGithub(value *GoogleCloudbuildTriggerGithub)
	PutPubsubConfig(value *GoogleCloudbuildTriggerPubsubConfig)
	PutSourceToBuild(value *GoogleCloudbuildTriggerSourceToBuild)
	PutTimeouts(value *GoogleCloudbuildTriggerTimeouts)
	PutTriggerTemplate(value *GoogleCloudbuildTriggerTriggerTemplate)
	PutWebhookConfig(value *GoogleCloudbuildTriggerWebhookConfig)
	ResetApprovalConfig()
	ResetBuildAttribute()
	ResetDescription()
	ResetDisabled()
	ResetFilename()
	ResetFilter()
	ResetGitFileSource()
	ResetGithub()
	ResetId()
	ResetIgnoredFiles()
	ResetIncludeBuildLogs()
	ResetIncludedFiles()
	ResetLocation()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetPubsubConfig()
	ResetServiceAccount()
	ResetSourceToBuild()
	ResetSubstitutions()
	ResetTags()
	ResetTimeouts()
	ResetTriggerTemplate()
	ResetWebhookConfig()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleCloudbuildTrigger
type jsiiProxy_GoogleCloudbuildTrigger struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) ApprovalConfig() GoogleCloudbuildTriggerApprovalConfigOutputReference {
	var returns GoogleCloudbuildTriggerApprovalConfigOutputReference
	_jsii_.Get(
		j,
		"approvalConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) ApprovalConfigInput() *GoogleCloudbuildTriggerApprovalConfig {
	var returns *GoogleCloudbuildTriggerApprovalConfig
	_jsii_.Get(
		j,
		"approvalConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) BuildAttribute() GoogleCloudbuildTriggerBuildOutputReference {
	var returns GoogleCloudbuildTriggerBuildOutputReference
	_jsii_.Get(
		j,
		"buildAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) BuildAttributeInput() *GoogleCloudbuildTriggerBuild {
	var returns *GoogleCloudbuildTriggerBuild
	_jsii_.Get(
		j,
		"buildAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) Filename() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) FilenameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filenameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) Filter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) FilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) GitFileSource() GoogleCloudbuildTriggerGitFileSourceOutputReference {
	var returns GoogleCloudbuildTriggerGitFileSourceOutputReference
	_jsii_.Get(
		j,
		"gitFileSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) GitFileSourceInput() *GoogleCloudbuildTriggerGitFileSource {
	var returns *GoogleCloudbuildTriggerGitFileSource
	_jsii_.Get(
		j,
		"gitFileSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) Github() GoogleCloudbuildTriggerGithubOutputReference {
	var returns GoogleCloudbuildTriggerGithubOutputReference
	_jsii_.Get(
		j,
		"github",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) GithubInput() *GoogleCloudbuildTriggerGithub {
	var returns *GoogleCloudbuildTriggerGithub
	_jsii_.Get(
		j,
		"githubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) IgnoredFiles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoredFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) IgnoredFilesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ignoredFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) IncludeBuildLogs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeBuildLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) IncludeBuildLogsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeBuildLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) IncludedFiles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) IncludedFilesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) PubsubConfig() GoogleCloudbuildTriggerPubsubConfigOutputReference {
	var returns GoogleCloudbuildTriggerPubsubConfigOutputReference
	_jsii_.Get(
		j,
		"pubsubConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) PubsubConfigInput() *GoogleCloudbuildTriggerPubsubConfig {
	var returns *GoogleCloudbuildTriggerPubsubConfig
	_jsii_.Get(
		j,
		"pubsubConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) ServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) ServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) SourceToBuild() GoogleCloudbuildTriggerSourceToBuildOutputReference {
	var returns GoogleCloudbuildTriggerSourceToBuildOutputReference
	_jsii_.Get(
		j,
		"sourceToBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) SourceToBuildInput() *GoogleCloudbuildTriggerSourceToBuild {
	var returns *GoogleCloudbuildTriggerSourceToBuild
	_jsii_.Get(
		j,
		"sourceToBuildInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) Substitutions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"substitutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) SubstitutionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"substitutionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) Timeouts() GoogleCloudbuildTriggerTimeoutsOutputReference {
	var returns GoogleCloudbuildTriggerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) TriggerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"triggerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) TriggerTemplate() GoogleCloudbuildTriggerTriggerTemplateOutputReference {
	var returns GoogleCloudbuildTriggerTriggerTemplateOutputReference
	_jsii_.Get(
		j,
		"triggerTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) TriggerTemplateInput() *GoogleCloudbuildTriggerTriggerTemplate {
	var returns *GoogleCloudbuildTriggerTriggerTemplate
	_jsii_.Get(
		j,
		"triggerTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) WebhookConfig() GoogleCloudbuildTriggerWebhookConfigOutputReference {
	var returns GoogleCloudbuildTriggerWebhookConfigOutputReference
	_jsii_.Get(
		j,
		"webhookConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTrigger) WebhookConfigInput() *GoogleCloudbuildTriggerWebhookConfig {
	var returns *GoogleCloudbuildTriggerWebhookConfig
	_jsii_.Get(
		j,
		"webhookConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudbuild_trigger google_cloudbuild_trigger} Resource.
func NewGoogleCloudbuildTrigger(scope constructs.Construct, id *string, config *GoogleCloudbuildTriggerConfig) GoogleCloudbuildTrigger {
	_init_.Initialize()

	if err := validateNewGoogleCloudbuildTriggerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudbuildTrigger{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudbuildTrigger.GoogleCloudbuildTrigger",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/google-beta/r/google_cloudbuild_trigger google_cloudbuild_trigger} Resource.
func NewGoogleCloudbuildTrigger_Override(g GoogleCloudbuildTrigger, scope constructs.Construct, id *string, config *GoogleCloudbuildTriggerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleCloudbuildTrigger.GoogleCloudbuildTrigger",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTrigger)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTrigger)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTrigger)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTrigger)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTrigger)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTrigger)SetFilename(val *string) {
	if err := j.validateSetFilenameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filename",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTrigger)SetFilter(val *string) {
	if err := j.validateSetFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filter",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTrigger)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTrigger)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTrigger)SetIgnoredFiles(val *[]*string) {
	if err := j.validateSetIgnoredFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoredFiles",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTrigger)SetIncludeBuildLogs(val *string) {
	if err := j.validateSetIncludeBuildLogsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeBuildLogs",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTrigger)SetIncludedFiles(val *[]*string) {
	if err := j.validateSetIncludedFilesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedFiles",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTrigger)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTrigger)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTrigger)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTrigger)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTrigger)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTrigger)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTrigger)SetServiceAccount(val *string) {
	if err := j.validateSetServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTrigger)SetSubstitutions(val *map[string]*string) {
	if err := j.validateSetSubstitutionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"substitutions",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTrigger)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func GoogleCloudbuildTrigger_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudbuildTrigger_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google-beta.googleCloudbuildTrigger.GoogleCloudbuildTrigger",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleCloudbuildTrigger_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google-beta.googleCloudbuildTrigger.GoogleCloudbuildTrigger",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) PutApprovalConfig(value *GoogleCloudbuildTriggerApprovalConfig) {
	if err := g.validatePutApprovalConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putApprovalConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) PutBuildAttribute(value *GoogleCloudbuildTriggerBuild) {
	if err := g.validatePutBuildAttributeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBuildAttribute",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) PutGitFileSource(value *GoogleCloudbuildTriggerGitFileSource) {
	if err := g.validatePutGitFileSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGitFileSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) PutGithub(value *GoogleCloudbuildTriggerGithub) {
	if err := g.validatePutGithubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGithub",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) PutPubsubConfig(value *GoogleCloudbuildTriggerPubsubConfig) {
	if err := g.validatePutPubsubConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPubsubConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) PutSourceToBuild(value *GoogleCloudbuildTriggerSourceToBuild) {
	if err := g.validatePutSourceToBuildParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSourceToBuild",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) PutTimeouts(value *GoogleCloudbuildTriggerTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) PutTriggerTemplate(value *GoogleCloudbuildTriggerTriggerTemplate) {
	if err := g.validatePutTriggerTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTriggerTemplate",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) PutWebhookConfig(value *GoogleCloudbuildTriggerWebhookConfig) {
	if err := g.validatePutWebhookConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putWebhookConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetApprovalConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetApprovalConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetBuildAttribute() {
	_jsii_.InvokeVoid(
		g,
		"resetBuildAttribute",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetDisabled() {
	_jsii_.InvokeVoid(
		g,
		"resetDisabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetFilename() {
	_jsii_.InvokeVoid(
		g,
		"resetFilename",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetFilter() {
	_jsii_.InvokeVoid(
		g,
		"resetFilter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetGitFileSource() {
	_jsii_.InvokeVoid(
		g,
		"resetGitFileSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetGithub() {
	_jsii_.InvokeVoid(
		g,
		"resetGithub",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetIgnoredFiles() {
	_jsii_.InvokeVoid(
		g,
		"resetIgnoredFiles",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetIncludeBuildLogs() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeBuildLogs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetIncludedFiles() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludedFiles",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetPubsubConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPubsubConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetSourceToBuild() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceToBuild",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetSubstitutions() {
	_jsii_.InvokeVoid(
		g,
		"resetSubstitutions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetTriggerTemplate() {
	_jsii_.InvokeVoid(
		g,
		"resetTriggerTemplate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ResetWebhookConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetWebhookConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTrigger) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
