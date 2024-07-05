# TimerInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Gets or sets the Id of the recording. | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**ServerId** | Pointer to **NullableString** | Gets or sets the server identifier. | [optional] 
**ExternalId** | Pointer to **NullableString** | Gets or sets the external identifier. | [optional] 
**ChannelId** | Pointer to **string** | Gets or sets the channel id of the recording. | [optional] 
**ExternalChannelId** | Pointer to **NullableString** | Gets or sets the external channel identifier. | [optional] 
**ChannelName** | Pointer to **NullableString** | Gets or sets the channel name of the recording. | [optional] 
**ChannelPrimaryImageTag** | Pointer to **NullableString** |  | [optional] 
**ProgramId** | Pointer to **NullableString** | Gets or sets the program identifier. | [optional] 
**ExternalProgramId** | Pointer to **NullableString** | Gets or sets the external program identifier. | [optional] 
**Name** | Pointer to **NullableString** | Gets or sets the name of the recording. | [optional] 
**Overview** | Pointer to **NullableString** | Gets or sets the description of the recording. | [optional] 
**StartDate** | Pointer to **time.Time** | Gets or sets the start date of the recording, in UTC. | [optional] 
**EndDate** | Pointer to **time.Time** | Gets or sets the end date of the recording, in UTC. | [optional] 
**ServiceName** | Pointer to **NullableString** | Gets or sets the name of the service. | [optional] 
**Priority** | Pointer to **int32** | Gets or sets the priority. | [optional] 
**PrePaddingSeconds** | Pointer to **int32** | Gets or sets the pre padding seconds. | [optional] 
**PostPaddingSeconds** | Pointer to **int32** | Gets or sets the post padding seconds. | [optional] 
**IsPrePaddingRequired** | Pointer to **bool** | Gets or sets a value indicating whether this instance is pre padding required. | [optional] 
**ParentBackdropItemId** | Pointer to **NullableString** | Gets or sets the Id of the Parent that has a backdrop if the item does not have one. | [optional] 
**ParentBackdropImageTags** | Pointer to **[]string** | Gets or sets the parent backdrop image tags. | [optional] 
**IsPostPaddingRequired** | Pointer to **bool** | Gets or sets a value indicating whether this instance is post padding required. | [optional] 
**KeepUntil** | Pointer to [**KeepUntil**](KeepUntil.md) |  | [optional] 
**Status** | Pointer to [**RecordingStatus**](RecordingStatus.md) | Gets or sets the status. | [optional] 
**SeriesTimerId** | Pointer to **NullableString** | Gets or sets the series timer identifier. | [optional] 
**ExternalSeriesTimerId** | Pointer to **NullableString** | Gets or sets the external series timer identifier. | [optional] 
**RunTimeTicks** | Pointer to **NullableInt64** | Gets or sets the run time ticks. | [optional] 
**ProgramInfo** | Pointer to [**NullableBaseItemDto**](BaseItemDto.md) | Gets or sets the program information. | [optional] 

## Methods

### NewTimerInfoDto

`func NewTimerInfoDto() *TimerInfoDto`

NewTimerInfoDto instantiates a new TimerInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimerInfoDtoWithDefaults

`func NewTimerInfoDtoWithDefaults() *TimerInfoDto`

NewTimerInfoDtoWithDefaults instantiates a new TimerInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimerInfoDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimerInfoDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimerInfoDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TimerInfoDto) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TimerInfoDto) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TimerInfoDto) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetType

`func (o *TimerInfoDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimerInfoDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimerInfoDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TimerInfoDto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *TimerInfoDto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *TimerInfoDto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetServerId

`func (o *TimerInfoDto) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *TimerInfoDto) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *TimerInfoDto) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *TimerInfoDto) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *TimerInfoDto) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *TimerInfoDto) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
### GetExternalId

`func (o *TimerInfoDto) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *TimerInfoDto) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *TimerInfoDto) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *TimerInfoDto) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *TimerInfoDto) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *TimerInfoDto) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetChannelId

`func (o *TimerInfoDto) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *TimerInfoDto) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *TimerInfoDto) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *TimerInfoDto) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetExternalChannelId

`func (o *TimerInfoDto) GetExternalChannelId() string`

GetExternalChannelId returns the ExternalChannelId field if non-nil, zero value otherwise.

### GetExternalChannelIdOk

`func (o *TimerInfoDto) GetExternalChannelIdOk() (*string, bool)`

GetExternalChannelIdOk returns a tuple with the ExternalChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalChannelId

`func (o *TimerInfoDto) SetExternalChannelId(v string)`

SetExternalChannelId sets ExternalChannelId field to given value.

### HasExternalChannelId

`func (o *TimerInfoDto) HasExternalChannelId() bool`

HasExternalChannelId returns a boolean if a field has been set.

### SetExternalChannelIdNil

`func (o *TimerInfoDto) SetExternalChannelIdNil(b bool)`

 SetExternalChannelIdNil sets the value for ExternalChannelId to be an explicit nil

### UnsetExternalChannelId
`func (o *TimerInfoDto) UnsetExternalChannelId()`

UnsetExternalChannelId ensures that no value is present for ExternalChannelId, not even an explicit nil
### GetChannelName

`func (o *TimerInfoDto) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *TimerInfoDto) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *TimerInfoDto) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *TimerInfoDto) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### SetChannelNameNil

`func (o *TimerInfoDto) SetChannelNameNil(b bool)`

 SetChannelNameNil sets the value for ChannelName to be an explicit nil

### UnsetChannelName
`func (o *TimerInfoDto) UnsetChannelName()`

UnsetChannelName ensures that no value is present for ChannelName, not even an explicit nil
### GetChannelPrimaryImageTag

`func (o *TimerInfoDto) GetChannelPrimaryImageTag() string`

GetChannelPrimaryImageTag returns the ChannelPrimaryImageTag field if non-nil, zero value otherwise.

### GetChannelPrimaryImageTagOk

`func (o *TimerInfoDto) GetChannelPrimaryImageTagOk() (*string, bool)`

GetChannelPrimaryImageTagOk returns a tuple with the ChannelPrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelPrimaryImageTag

`func (o *TimerInfoDto) SetChannelPrimaryImageTag(v string)`

SetChannelPrimaryImageTag sets ChannelPrimaryImageTag field to given value.

### HasChannelPrimaryImageTag

`func (o *TimerInfoDto) HasChannelPrimaryImageTag() bool`

HasChannelPrimaryImageTag returns a boolean if a field has been set.

### SetChannelPrimaryImageTagNil

`func (o *TimerInfoDto) SetChannelPrimaryImageTagNil(b bool)`

 SetChannelPrimaryImageTagNil sets the value for ChannelPrimaryImageTag to be an explicit nil

### UnsetChannelPrimaryImageTag
`func (o *TimerInfoDto) UnsetChannelPrimaryImageTag()`

UnsetChannelPrimaryImageTag ensures that no value is present for ChannelPrimaryImageTag, not even an explicit nil
### GetProgramId

`func (o *TimerInfoDto) GetProgramId() string`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *TimerInfoDto) GetProgramIdOk() (*string, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *TimerInfoDto) SetProgramId(v string)`

SetProgramId sets ProgramId field to given value.

### HasProgramId

`func (o *TimerInfoDto) HasProgramId() bool`

HasProgramId returns a boolean if a field has been set.

### SetProgramIdNil

`func (o *TimerInfoDto) SetProgramIdNil(b bool)`

 SetProgramIdNil sets the value for ProgramId to be an explicit nil

### UnsetProgramId
`func (o *TimerInfoDto) UnsetProgramId()`

UnsetProgramId ensures that no value is present for ProgramId, not even an explicit nil
### GetExternalProgramId

`func (o *TimerInfoDto) GetExternalProgramId() string`

GetExternalProgramId returns the ExternalProgramId field if non-nil, zero value otherwise.

### GetExternalProgramIdOk

`func (o *TimerInfoDto) GetExternalProgramIdOk() (*string, bool)`

GetExternalProgramIdOk returns a tuple with the ExternalProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalProgramId

`func (o *TimerInfoDto) SetExternalProgramId(v string)`

SetExternalProgramId sets ExternalProgramId field to given value.

### HasExternalProgramId

`func (o *TimerInfoDto) HasExternalProgramId() bool`

HasExternalProgramId returns a boolean if a field has been set.

### SetExternalProgramIdNil

`func (o *TimerInfoDto) SetExternalProgramIdNil(b bool)`

 SetExternalProgramIdNil sets the value for ExternalProgramId to be an explicit nil

### UnsetExternalProgramId
`func (o *TimerInfoDto) UnsetExternalProgramId()`

UnsetExternalProgramId ensures that no value is present for ExternalProgramId, not even an explicit nil
### GetName

`func (o *TimerInfoDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TimerInfoDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TimerInfoDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TimerInfoDto) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *TimerInfoDto) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *TimerInfoDto) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOverview

`func (o *TimerInfoDto) GetOverview() string`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *TimerInfoDto) GetOverviewOk() (*string, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *TimerInfoDto) SetOverview(v string)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *TimerInfoDto) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### SetOverviewNil

`func (o *TimerInfoDto) SetOverviewNil(b bool)`

 SetOverviewNil sets the value for Overview to be an explicit nil

### UnsetOverview
`func (o *TimerInfoDto) UnsetOverview()`

UnsetOverview ensures that no value is present for Overview, not even an explicit nil
### GetStartDate

`func (o *TimerInfoDto) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TimerInfoDto) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TimerInfoDto) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TimerInfoDto) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *TimerInfoDto) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TimerInfoDto) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TimerInfoDto) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *TimerInfoDto) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetServiceName

`func (o *TimerInfoDto) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *TimerInfoDto) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *TimerInfoDto) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *TimerInfoDto) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceNameNil

`func (o *TimerInfoDto) SetServiceNameNil(b bool)`

 SetServiceNameNil sets the value for ServiceName to be an explicit nil

### UnsetServiceName
`func (o *TimerInfoDto) UnsetServiceName()`

UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil
### GetPriority

`func (o *TimerInfoDto) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *TimerInfoDto) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *TimerInfoDto) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *TimerInfoDto) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPrePaddingSeconds

`func (o *TimerInfoDto) GetPrePaddingSeconds() int32`

GetPrePaddingSeconds returns the PrePaddingSeconds field if non-nil, zero value otherwise.

### GetPrePaddingSecondsOk

`func (o *TimerInfoDto) GetPrePaddingSecondsOk() (*int32, bool)`

GetPrePaddingSecondsOk returns a tuple with the PrePaddingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePaddingSeconds

`func (o *TimerInfoDto) SetPrePaddingSeconds(v int32)`

SetPrePaddingSeconds sets PrePaddingSeconds field to given value.

### HasPrePaddingSeconds

`func (o *TimerInfoDto) HasPrePaddingSeconds() bool`

HasPrePaddingSeconds returns a boolean if a field has been set.

### GetPostPaddingSeconds

`func (o *TimerInfoDto) GetPostPaddingSeconds() int32`

GetPostPaddingSeconds returns the PostPaddingSeconds field if non-nil, zero value otherwise.

### GetPostPaddingSecondsOk

`func (o *TimerInfoDto) GetPostPaddingSecondsOk() (*int32, bool)`

GetPostPaddingSecondsOk returns a tuple with the PostPaddingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostPaddingSeconds

`func (o *TimerInfoDto) SetPostPaddingSeconds(v int32)`

SetPostPaddingSeconds sets PostPaddingSeconds field to given value.

### HasPostPaddingSeconds

`func (o *TimerInfoDto) HasPostPaddingSeconds() bool`

HasPostPaddingSeconds returns a boolean if a field has been set.

### GetIsPrePaddingRequired

`func (o *TimerInfoDto) GetIsPrePaddingRequired() bool`

GetIsPrePaddingRequired returns the IsPrePaddingRequired field if non-nil, zero value otherwise.

### GetIsPrePaddingRequiredOk

`func (o *TimerInfoDto) GetIsPrePaddingRequiredOk() (*bool, bool)`

GetIsPrePaddingRequiredOk returns a tuple with the IsPrePaddingRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrePaddingRequired

`func (o *TimerInfoDto) SetIsPrePaddingRequired(v bool)`

SetIsPrePaddingRequired sets IsPrePaddingRequired field to given value.

### HasIsPrePaddingRequired

`func (o *TimerInfoDto) HasIsPrePaddingRequired() bool`

HasIsPrePaddingRequired returns a boolean if a field has been set.

### GetParentBackdropItemId

`func (o *TimerInfoDto) GetParentBackdropItemId() string`

GetParentBackdropItemId returns the ParentBackdropItemId field if non-nil, zero value otherwise.

### GetParentBackdropItemIdOk

`func (o *TimerInfoDto) GetParentBackdropItemIdOk() (*string, bool)`

GetParentBackdropItemIdOk returns a tuple with the ParentBackdropItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBackdropItemId

`func (o *TimerInfoDto) SetParentBackdropItemId(v string)`

SetParentBackdropItemId sets ParentBackdropItemId field to given value.

### HasParentBackdropItemId

`func (o *TimerInfoDto) HasParentBackdropItemId() bool`

HasParentBackdropItemId returns a boolean if a field has been set.

### SetParentBackdropItemIdNil

`func (o *TimerInfoDto) SetParentBackdropItemIdNil(b bool)`

 SetParentBackdropItemIdNil sets the value for ParentBackdropItemId to be an explicit nil

### UnsetParentBackdropItemId
`func (o *TimerInfoDto) UnsetParentBackdropItemId()`

UnsetParentBackdropItemId ensures that no value is present for ParentBackdropItemId, not even an explicit nil
### GetParentBackdropImageTags

`func (o *TimerInfoDto) GetParentBackdropImageTags() []string`

GetParentBackdropImageTags returns the ParentBackdropImageTags field if non-nil, zero value otherwise.

### GetParentBackdropImageTagsOk

`func (o *TimerInfoDto) GetParentBackdropImageTagsOk() (*[]string, bool)`

GetParentBackdropImageTagsOk returns a tuple with the ParentBackdropImageTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBackdropImageTags

`func (o *TimerInfoDto) SetParentBackdropImageTags(v []string)`

SetParentBackdropImageTags sets ParentBackdropImageTags field to given value.

### HasParentBackdropImageTags

`func (o *TimerInfoDto) HasParentBackdropImageTags() bool`

HasParentBackdropImageTags returns a boolean if a field has been set.

### SetParentBackdropImageTagsNil

`func (o *TimerInfoDto) SetParentBackdropImageTagsNil(b bool)`

 SetParentBackdropImageTagsNil sets the value for ParentBackdropImageTags to be an explicit nil

### UnsetParentBackdropImageTags
`func (o *TimerInfoDto) UnsetParentBackdropImageTags()`

UnsetParentBackdropImageTags ensures that no value is present for ParentBackdropImageTags, not even an explicit nil
### GetIsPostPaddingRequired

`func (o *TimerInfoDto) GetIsPostPaddingRequired() bool`

GetIsPostPaddingRequired returns the IsPostPaddingRequired field if non-nil, zero value otherwise.

### GetIsPostPaddingRequiredOk

`func (o *TimerInfoDto) GetIsPostPaddingRequiredOk() (*bool, bool)`

GetIsPostPaddingRequiredOk returns a tuple with the IsPostPaddingRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPostPaddingRequired

`func (o *TimerInfoDto) SetIsPostPaddingRequired(v bool)`

SetIsPostPaddingRequired sets IsPostPaddingRequired field to given value.

### HasIsPostPaddingRequired

`func (o *TimerInfoDto) HasIsPostPaddingRequired() bool`

HasIsPostPaddingRequired returns a boolean if a field has been set.

### GetKeepUntil

`func (o *TimerInfoDto) GetKeepUntil() KeepUntil`

GetKeepUntil returns the KeepUntil field if non-nil, zero value otherwise.

### GetKeepUntilOk

`func (o *TimerInfoDto) GetKeepUntilOk() (*KeepUntil, bool)`

GetKeepUntilOk returns a tuple with the KeepUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepUntil

`func (o *TimerInfoDto) SetKeepUntil(v KeepUntil)`

SetKeepUntil sets KeepUntil field to given value.

### HasKeepUntil

`func (o *TimerInfoDto) HasKeepUntil() bool`

HasKeepUntil returns a boolean if a field has been set.

### GetStatus

`func (o *TimerInfoDto) GetStatus() RecordingStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TimerInfoDto) GetStatusOk() (*RecordingStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TimerInfoDto) SetStatus(v RecordingStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TimerInfoDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSeriesTimerId

`func (o *TimerInfoDto) GetSeriesTimerId() string`

GetSeriesTimerId returns the SeriesTimerId field if non-nil, zero value otherwise.

### GetSeriesTimerIdOk

`func (o *TimerInfoDto) GetSeriesTimerIdOk() (*string, bool)`

GetSeriesTimerIdOk returns a tuple with the SeriesTimerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesTimerId

`func (o *TimerInfoDto) SetSeriesTimerId(v string)`

SetSeriesTimerId sets SeriesTimerId field to given value.

### HasSeriesTimerId

`func (o *TimerInfoDto) HasSeriesTimerId() bool`

HasSeriesTimerId returns a boolean if a field has been set.

### SetSeriesTimerIdNil

`func (o *TimerInfoDto) SetSeriesTimerIdNil(b bool)`

 SetSeriesTimerIdNil sets the value for SeriesTimerId to be an explicit nil

### UnsetSeriesTimerId
`func (o *TimerInfoDto) UnsetSeriesTimerId()`

UnsetSeriesTimerId ensures that no value is present for SeriesTimerId, not even an explicit nil
### GetExternalSeriesTimerId

`func (o *TimerInfoDto) GetExternalSeriesTimerId() string`

GetExternalSeriesTimerId returns the ExternalSeriesTimerId field if non-nil, zero value otherwise.

### GetExternalSeriesTimerIdOk

`func (o *TimerInfoDto) GetExternalSeriesTimerIdOk() (*string, bool)`

GetExternalSeriesTimerIdOk returns a tuple with the ExternalSeriesTimerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSeriesTimerId

`func (o *TimerInfoDto) SetExternalSeriesTimerId(v string)`

SetExternalSeriesTimerId sets ExternalSeriesTimerId field to given value.

### HasExternalSeriesTimerId

`func (o *TimerInfoDto) HasExternalSeriesTimerId() bool`

HasExternalSeriesTimerId returns a boolean if a field has been set.

### SetExternalSeriesTimerIdNil

`func (o *TimerInfoDto) SetExternalSeriesTimerIdNil(b bool)`

 SetExternalSeriesTimerIdNil sets the value for ExternalSeriesTimerId to be an explicit nil

### UnsetExternalSeriesTimerId
`func (o *TimerInfoDto) UnsetExternalSeriesTimerId()`

UnsetExternalSeriesTimerId ensures that no value is present for ExternalSeriesTimerId, not even an explicit nil
### GetRunTimeTicks

`func (o *TimerInfoDto) GetRunTimeTicks() int64`

GetRunTimeTicks returns the RunTimeTicks field if non-nil, zero value otherwise.

### GetRunTimeTicksOk

`func (o *TimerInfoDto) GetRunTimeTicksOk() (*int64, bool)`

GetRunTimeTicksOk returns a tuple with the RunTimeTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTimeTicks

`func (o *TimerInfoDto) SetRunTimeTicks(v int64)`

SetRunTimeTicks sets RunTimeTicks field to given value.

### HasRunTimeTicks

`func (o *TimerInfoDto) HasRunTimeTicks() bool`

HasRunTimeTicks returns a boolean if a field has been set.

### SetRunTimeTicksNil

`func (o *TimerInfoDto) SetRunTimeTicksNil(b bool)`

 SetRunTimeTicksNil sets the value for RunTimeTicks to be an explicit nil

### UnsetRunTimeTicks
`func (o *TimerInfoDto) UnsetRunTimeTicks()`

UnsetRunTimeTicks ensures that no value is present for RunTimeTicks, not even an explicit nil
### GetProgramInfo

`func (o *TimerInfoDto) GetProgramInfo() BaseItemDto`

GetProgramInfo returns the ProgramInfo field if non-nil, zero value otherwise.

### GetProgramInfoOk

`func (o *TimerInfoDto) GetProgramInfoOk() (*BaseItemDto, bool)`

GetProgramInfoOk returns a tuple with the ProgramInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramInfo

`func (o *TimerInfoDto) SetProgramInfo(v BaseItemDto)`

SetProgramInfo sets ProgramInfo field to given value.

### HasProgramInfo

`func (o *TimerInfoDto) HasProgramInfo() bool`

HasProgramInfo returns a boolean if a field has been set.

### SetProgramInfoNil

`func (o *TimerInfoDto) SetProgramInfoNil(b bool)`

 SetProgramInfoNil sets the value for ProgramInfo to be an explicit nil

### UnsetProgramInfo
`func (o *TimerInfoDto) UnsetProgramInfo()`

UnsetProgramInfo ensures that no value is present for ProgramInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


