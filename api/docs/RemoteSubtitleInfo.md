# RemoteSubtitleInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreeLetterISOLanguageName** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**ProviderName** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Format** | Pointer to **NullableString** |  | [optional] 
**Author** | Pointer to **NullableString** |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **NullableTime** |  | [optional] 
**CommunityRating** | Pointer to **NullableFloat32** |  | [optional] 
**FrameRate** | Pointer to **NullableFloat32** |  | [optional] 
**DownloadCount** | Pointer to **NullableInt32** |  | [optional] 
**IsHashMatch** | Pointer to **NullableBool** |  | [optional] 
**AiTranslated** | Pointer to **NullableBool** |  | [optional] 
**MachineTranslated** | Pointer to **NullableBool** |  | [optional] 
**Forced** | Pointer to **NullableBool** |  | [optional] 
**HearingImpaired** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewRemoteSubtitleInfo

`func NewRemoteSubtitleInfo() *RemoteSubtitleInfo`

NewRemoteSubtitleInfo instantiates a new RemoteSubtitleInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteSubtitleInfoWithDefaults

`func NewRemoteSubtitleInfoWithDefaults() *RemoteSubtitleInfo`

NewRemoteSubtitleInfoWithDefaults instantiates a new RemoteSubtitleInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreeLetterISOLanguageName

`func (o *RemoteSubtitleInfo) GetThreeLetterISOLanguageName() string`

GetThreeLetterISOLanguageName returns the ThreeLetterISOLanguageName field if non-nil, zero value otherwise.

### GetThreeLetterISOLanguageNameOk

`func (o *RemoteSubtitleInfo) GetThreeLetterISOLanguageNameOk() (*string, bool)`

GetThreeLetterISOLanguageNameOk returns a tuple with the ThreeLetterISOLanguageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeLetterISOLanguageName

`func (o *RemoteSubtitleInfo) SetThreeLetterISOLanguageName(v string)`

SetThreeLetterISOLanguageName sets ThreeLetterISOLanguageName field to given value.

### HasThreeLetterISOLanguageName

`func (o *RemoteSubtitleInfo) HasThreeLetterISOLanguageName() bool`

HasThreeLetterISOLanguageName returns a boolean if a field has been set.

### SetThreeLetterISOLanguageNameNil

`func (o *RemoteSubtitleInfo) SetThreeLetterISOLanguageNameNil(b bool)`

 SetThreeLetterISOLanguageNameNil sets the value for ThreeLetterISOLanguageName to be an explicit nil

### UnsetThreeLetterISOLanguageName
`func (o *RemoteSubtitleInfo) UnsetThreeLetterISOLanguageName()`

UnsetThreeLetterISOLanguageName ensures that no value is present for ThreeLetterISOLanguageName, not even an explicit nil
### GetId

`func (o *RemoteSubtitleInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoteSubtitleInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoteSubtitleInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RemoteSubtitleInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *RemoteSubtitleInfo) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RemoteSubtitleInfo) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetProviderName

`func (o *RemoteSubtitleInfo) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *RemoteSubtitleInfo) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *RemoteSubtitleInfo) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *RemoteSubtitleInfo) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *RemoteSubtitleInfo) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *RemoteSubtitleInfo) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
### GetName

`func (o *RemoteSubtitleInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemoteSubtitleInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemoteSubtitleInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RemoteSubtitleInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *RemoteSubtitleInfo) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *RemoteSubtitleInfo) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetFormat

`func (o *RemoteSubtitleInfo) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *RemoteSubtitleInfo) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *RemoteSubtitleInfo) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *RemoteSubtitleInfo) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### SetFormatNil

`func (o *RemoteSubtitleInfo) SetFormatNil(b bool)`

 SetFormatNil sets the value for Format to be an explicit nil

### UnsetFormat
`func (o *RemoteSubtitleInfo) UnsetFormat()`

UnsetFormat ensures that no value is present for Format, not even an explicit nil
### GetAuthor

`func (o *RemoteSubtitleInfo) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *RemoteSubtitleInfo) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *RemoteSubtitleInfo) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *RemoteSubtitleInfo) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *RemoteSubtitleInfo) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *RemoteSubtitleInfo) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetComment

`func (o *RemoteSubtitleInfo) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RemoteSubtitleInfo) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RemoteSubtitleInfo) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RemoteSubtitleInfo) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *RemoteSubtitleInfo) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *RemoteSubtitleInfo) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetDateCreated

`func (o *RemoteSubtitleInfo) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *RemoteSubtitleInfo) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *RemoteSubtitleInfo) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *RemoteSubtitleInfo) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *RemoteSubtitleInfo) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *RemoteSubtitleInfo) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetCommunityRating

`func (o *RemoteSubtitleInfo) GetCommunityRating() float32`

GetCommunityRating returns the CommunityRating field if non-nil, zero value otherwise.

### GetCommunityRatingOk

`func (o *RemoteSubtitleInfo) GetCommunityRatingOk() (*float32, bool)`

GetCommunityRatingOk returns a tuple with the CommunityRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityRating

`func (o *RemoteSubtitleInfo) SetCommunityRating(v float32)`

SetCommunityRating sets CommunityRating field to given value.

### HasCommunityRating

`func (o *RemoteSubtitleInfo) HasCommunityRating() bool`

HasCommunityRating returns a boolean if a field has been set.

### SetCommunityRatingNil

`func (o *RemoteSubtitleInfo) SetCommunityRatingNil(b bool)`

 SetCommunityRatingNil sets the value for CommunityRating to be an explicit nil

### UnsetCommunityRating
`func (o *RemoteSubtitleInfo) UnsetCommunityRating()`

UnsetCommunityRating ensures that no value is present for CommunityRating, not even an explicit nil
### GetFrameRate

`func (o *RemoteSubtitleInfo) GetFrameRate() float32`

GetFrameRate returns the FrameRate field if non-nil, zero value otherwise.

### GetFrameRateOk

`func (o *RemoteSubtitleInfo) GetFrameRateOk() (*float32, bool)`

GetFrameRateOk returns a tuple with the FrameRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameRate

`func (o *RemoteSubtitleInfo) SetFrameRate(v float32)`

SetFrameRate sets FrameRate field to given value.

### HasFrameRate

`func (o *RemoteSubtitleInfo) HasFrameRate() bool`

HasFrameRate returns a boolean if a field has been set.

### SetFrameRateNil

`func (o *RemoteSubtitleInfo) SetFrameRateNil(b bool)`

 SetFrameRateNil sets the value for FrameRate to be an explicit nil

### UnsetFrameRate
`func (o *RemoteSubtitleInfo) UnsetFrameRate()`

UnsetFrameRate ensures that no value is present for FrameRate, not even an explicit nil
### GetDownloadCount

`func (o *RemoteSubtitleInfo) GetDownloadCount() int32`

GetDownloadCount returns the DownloadCount field if non-nil, zero value otherwise.

### GetDownloadCountOk

`func (o *RemoteSubtitleInfo) GetDownloadCountOk() (*int32, bool)`

GetDownloadCountOk returns a tuple with the DownloadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadCount

`func (o *RemoteSubtitleInfo) SetDownloadCount(v int32)`

SetDownloadCount sets DownloadCount field to given value.

### HasDownloadCount

`func (o *RemoteSubtitleInfo) HasDownloadCount() bool`

HasDownloadCount returns a boolean if a field has been set.

### SetDownloadCountNil

`func (o *RemoteSubtitleInfo) SetDownloadCountNil(b bool)`

 SetDownloadCountNil sets the value for DownloadCount to be an explicit nil

### UnsetDownloadCount
`func (o *RemoteSubtitleInfo) UnsetDownloadCount()`

UnsetDownloadCount ensures that no value is present for DownloadCount, not even an explicit nil
### GetIsHashMatch

`func (o *RemoteSubtitleInfo) GetIsHashMatch() bool`

GetIsHashMatch returns the IsHashMatch field if non-nil, zero value otherwise.

### GetIsHashMatchOk

`func (o *RemoteSubtitleInfo) GetIsHashMatchOk() (*bool, bool)`

GetIsHashMatchOk returns a tuple with the IsHashMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHashMatch

`func (o *RemoteSubtitleInfo) SetIsHashMatch(v bool)`

SetIsHashMatch sets IsHashMatch field to given value.

### HasIsHashMatch

`func (o *RemoteSubtitleInfo) HasIsHashMatch() bool`

HasIsHashMatch returns a boolean if a field has been set.

### SetIsHashMatchNil

`func (o *RemoteSubtitleInfo) SetIsHashMatchNil(b bool)`

 SetIsHashMatchNil sets the value for IsHashMatch to be an explicit nil

### UnsetIsHashMatch
`func (o *RemoteSubtitleInfo) UnsetIsHashMatch()`

UnsetIsHashMatch ensures that no value is present for IsHashMatch, not even an explicit nil
### GetAiTranslated

`func (o *RemoteSubtitleInfo) GetAiTranslated() bool`

GetAiTranslated returns the AiTranslated field if non-nil, zero value otherwise.

### GetAiTranslatedOk

`func (o *RemoteSubtitleInfo) GetAiTranslatedOk() (*bool, bool)`

GetAiTranslatedOk returns a tuple with the AiTranslated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiTranslated

`func (o *RemoteSubtitleInfo) SetAiTranslated(v bool)`

SetAiTranslated sets AiTranslated field to given value.

### HasAiTranslated

`func (o *RemoteSubtitleInfo) HasAiTranslated() bool`

HasAiTranslated returns a boolean if a field has been set.

### SetAiTranslatedNil

`func (o *RemoteSubtitleInfo) SetAiTranslatedNil(b bool)`

 SetAiTranslatedNil sets the value for AiTranslated to be an explicit nil

### UnsetAiTranslated
`func (o *RemoteSubtitleInfo) UnsetAiTranslated()`

UnsetAiTranslated ensures that no value is present for AiTranslated, not even an explicit nil
### GetMachineTranslated

`func (o *RemoteSubtitleInfo) GetMachineTranslated() bool`

GetMachineTranslated returns the MachineTranslated field if non-nil, zero value otherwise.

### GetMachineTranslatedOk

`func (o *RemoteSubtitleInfo) GetMachineTranslatedOk() (*bool, bool)`

GetMachineTranslatedOk returns a tuple with the MachineTranslated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineTranslated

`func (o *RemoteSubtitleInfo) SetMachineTranslated(v bool)`

SetMachineTranslated sets MachineTranslated field to given value.

### HasMachineTranslated

`func (o *RemoteSubtitleInfo) HasMachineTranslated() bool`

HasMachineTranslated returns a boolean if a field has been set.

### SetMachineTranslatedNil

`func (o *RemoteSubtitleInfo) SetMachineTranslatedNil(b bool)`

 SetMachineTranslatedNil sets the value for MachineTranslated to be an explicit nil

### UnsetMachineTranslated
`func (o *RemoteSubtitleInfo) UnsetMachineTranslated()`

UnsetMachineTranslated ensures that no value is present for MachineTranslated, not even an explicit nil
### GetForced

`func (o *RemoteSubtitleInfo) GetForced() bool`

GetForced returns the Forced field if non-nil, zero value otherwise.

### GetForcedOk

`func (o *RemoteSubtitleInfo) GetForcedOk() (*bool, bool)`

GetForcedOk returns a tuple with the Forced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForced

`func (o *RemoteSubtitleInfo) SetForced(v bool)`

SetForced sets Forced field to given value.

### HasForced

`func (o *RemoteSubtitleInfo) HasForced() bool`

HasForced returns a boolean if a field has been set.

### SetForcedNil

`func (o *RemoteSubtitleInfo) SetForcedNil(b bool)`

 SetForcedNil sets the value for Forced to be an explicit nil

### UnsetForced
`func (o *RemoteSubtitleInfo) UnsetForced()`

UnsetForced ensures that no value is present for Forced, not even an explicit nil
### GetHearingImpaired

`func (o *RemoteSubtitleInfo) GetHearingImpaired() bool`

GetHearingImpaired returns the HearingImpaired field if non-nil, zero value otherwise.

### GetHearingImpairedOk

`func (o *RemoteSubtitleInfo) GetHearingImpairedOk() (*bool, bool)`

GetHearingImpairedOk returns a tuple with the HearingImpaired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHearingImpaired

`func (o *RemoteSubtitleInfo) SetHearingImpaired(v bool)`

SetHearingImpaired sets HearingImpaired field to given value.

### HasHearingImpaired

`func (o *RemoteSubtitleInfo) HasHearingImpaired() bool`

HasHearingImpaired returns a boolean if a field has been set.

### SetHearingImpairedNil

`func (o *RemoteSubtitleInfo) SetHearingImpairedNil(b bool)`

 SetHearingImpairedNil sets the value for HearingImpaired to be an explicit nil

### UnsetHearingImpaired
`func (o *RemoteSubtitleInfo) UnsetHearingImpaired()`

UnsetHearingImpaired ensures that no value is present for HearingImpaired, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


