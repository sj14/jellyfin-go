# VersionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | Gets or sets the version. | [optional] 
**VersionNumber** | Pointer to **string** | Gets the version as a System.Version. | [optional] [readonly] 
**Changelog** | Pointer to **NullableString** | Gets or sets the changelog for this version. | [optional] 
**TargetAbi** | Pointer to **NullableString** | Gets or sets the ABI that this version was built against. | [optional] 
**SourceUrl** | Pointer to **NullableString** | Gets or sets the source URL. | [optional] 
**Checksum** | Pointer to **NullableString** | Gets or sets a checksum for the binary. | [optional] 
**Timestamp** | Pointer to **NullableString** | Gets or sets a timestamp of when the binary was built. | [optional] 
**RepositoryName** | Pointer to **string** | Gets or sets the repository name. | [optional] 
**RepositoryUrl** | Pointer to **string** | Gets or sets the repository url. | [optional] 

## Methods

### NewVersionInfo

`func NewVersionInfo() *VersionInfo`

NewVersionInfo instantiates a new VersionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionInfoWithDefaults

`func NewVersionInfoWithDefaults() *VersionInfo`

NewVersionInfoWithDefaults instantiates a new VersionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *VersionInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VersionInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VersionInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VersionInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionNumber

`func (o *VersionInfo) GetVersionNumber() string`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *VersionInfo) GetVersionNumberOk() (*string, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *VersionInfo) SetVersionNumber(v string)`

SetVersionNumber sets VersionNumber field to given value.

### HasVersionNumber

`func (o *VersionInfo) HasVersionNumber() bool`

HasVersionNumber returns a boolean if a field has been set.

### GetChangelog

`func (o *VersionInfo) GetChangelog() string`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *VersionInfo) GetChangelogOk() (*string, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *VersionInfo) SetChangelog(v string)`

SetChangelog sets Changelog field to given value.

### HasChangelog

`func (o *VersionInfo) HasChangelog() bool`

HasChangelog returns a boolean if a field has been set.

### SetChangelogNil

`func (o *VersionInfo) SetChangelogNil(b bool)`

 SetChangelogNil sets the value for Changelog to be an explicit nil

### UnsetChangelog
`func (o *VersionInfo) UnsetChangelog()`

UnsetChangelog ensures that no value is present for Changelog, not even an explicit nil
### GetTargetAbi

`func (o *VersionInfo) GetTargetAbi() string`

GetTargetAbi returns the TargetAbi field if non-nil, zero value otherwise.

### GetTargetAbiOk

`func (o *VersionInfo) GetTargetAbiOk() (*string, bool)`

GetTargetAbiOk returns a tuple with the TargetAbi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAbi

`func (o *VersionInfo) SetTargetAbi(v string)`

SetTargetAbi sets TargetAbi field to given value.

### HasTargetAbi

`func (o *VersionInfo) HasTargetAbi() bool`

HasTargetAbi returns a boolean if a field has been set.

### SetTargetAbiNil

`func (o *VersionInfo) SetTargetAbiNil(b bool)`

 SetTargetAbiNil sets the value for TargetAbi to be an explicit nil

### UnsetTargetAbi
`func (o *VersionInfo) UnsetTargetAbi()`

UnsetTargetAbi ensures that no value is present for TargetAbi, not even an explicit nil
### GetSourceUrl

`func (o *VersionInfo) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *VersionInfo) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *VersionInfo) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *VersionInfo) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### SetSourceUrlNil

`func (o *VersionInfo) SetSourceUrlNil(b bool)`

 SetSourceUrlNil sets the value for SourceUrl to be an explicit nil

### UnsetSourceUrl
`func (o *VersionInfo) UnsetSourceUrl()`

UnsetSourceUrl ensures that no value is present for SourceUrl, not even an explicit nil
### GetChecksum

`func (o *VersionInfo) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *VersionInfo) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *VersionInfo) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *VersionInfo) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### SetChecksumNil

`func (o *VersionInfo) SetChecksumNil(b bool)`

 SetChecksumNil sets the value for Checksum to be an explicit nil

### UnsetChecksum
`func (o *VersionInfo) UnsetChecksum()`

UnsetChecksum ensures that no value is present for Checksum, not even an explicit nil
### GetTimestamp

`func (o *VersionInfo) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *VersionInfo) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *VersionInfo) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *VersionInfo) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *VersionInfo) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *VersionInfo) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetRepositoryName

`func (o *VersionInfo) GetRepositoryName() string`

GetRepositoryName returns the RepositoryName field if non-nil, zero value otherwise.

### GetRepositoryNameOk

`func (o *VersionInfo) GetRepositoryNameOk() (*string, bool)`

GetRepositoryNameOk returns a tuple with the RepositoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryName

`func (o *VersionInfo) SetRepositoryName(v string)`

SetRepositoryName sets RepositoryName field to given value.

### HasRepositoryName

`func (o *VersionInfo) HasRepositoryName() bool`

HasRepositoryName returns a boolean if a field has been set.

### GetRepositoryUrl

`func (o *VersionInfo) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *VersionInfo) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *VersionInfo) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.

### HasRepositoryUrl

`func (o *VersionInfo) HasRepositoryUrl() bool`

HasRepositoryUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


