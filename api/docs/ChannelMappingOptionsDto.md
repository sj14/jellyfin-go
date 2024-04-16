# ChannelMappingOptionsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TunerChannels** | Pointer to [**[]TunerChannelMapping**](TunerChannelMapping.md) | Gets or sets list of tuner channels. | [optional] 
**ProviderChannels** | Pointer to [**[]NameIdPair**](NameIdPair.md) | Gets or sets list of provider channels. | [optional] 
**Mappings** | Pointer to [**[]NameValuePair**](NameValuePair.md) | Gets or sets list of mappings. | [optional] 
**ProviderName** | Pointer to **NullableString** | Gets or sets provider name. | [optional] 

## Methods

### NewChannelMappingOptionsDto

`func NewChannelMappingOptionsDto() *ChannelMappingOptionsDto`

NewChannelMappingOptionsDto instantiates a new ChannelMappingOptionsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelMappingOptionsDtoWithDefaults

`func NewChannelMappingOptionsDtoWithDefaults() *ChannelMappingOptionsDto`

NewChannelMappingOptionsDtoWithDefaults instantiates a new ChannelMappingOptionsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTunerChannels

`func (o *ChannelMappingOptionsDto) GetTunerChannels() []TunerChannelMapping`

GetTunerChannels returns the TunerChannels field if non-nil, zero value otherwise.

### GetTunerChannelsOk

`func (o *ChannelMappingOptionsDto) GetTunerChannelsOk() (*[]TunerChannelMapping, bool)`

GetTunerChannelsOk returns a tuple with the TunerChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunerChannels

`func (o *ChannelMappingOptionsDto) SetTunerChannels(v []TunerChannelMapping)`

SetTunerChannels sets TunerChannels field to given value.

### HasTunerChannels

`func (o *ChannelMappingOptionsDto) HasTunerChannels() bool`

HasTunerChannels returns a boolean if a field has been set.

### GetProviderChannels

`func (o *ChannelMappingOptionsDto) GetProviderChannels() []NameIdPair`

GetProviderChannels returns the ProviderChannels field if non-nil, zero value otherwise.

### GetProviderChannelsOk

`func (o *ChannelMappingOptionsDto) GetProviderChannelsOk() (*[]NameIdPair, bool)`

GetProviderChannelsOk returns a tuple with the ProviderChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderChannels

`func (o *ChannelMappingOptionsDto) SetProviderChannels(v []NameIdPair)`

SetProviderChannels sets ProviderChannels field to given value.

### HasProviderChannels

`func (o *ChannelMappingOptionsDto) HasProviderChannels() bool`

HasProviderChannels returns a boolean if a field has been set.

### GetMappings

`func (o *ChannelMappingOptionsDto) GetMappings() []NameValuePair`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *ChannelMappingOptionsDto) GetMappingsOk() (*[]NameValuePair, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *ChannelMappingOptionsDto) SetMappings(v []NameValuePair)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *ChannelMappingOptionsDto) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### GetProviderName

`func (o *ChannelMappingOptionsDto) GetProviderName() string`

GetProviderName returns the ProviderName field if non-nil, zero value otherwise.

### GetProviderNameOk

`func (o *ChannelMappingOptionsDto) GetProviderNameOk() (*string, bool)`

GetProviderNameOk returns a tuple with the ProviderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderName

`func (o *ChannelMappingOptionsDto) SetProviderName(v string)`

SetProviderName sets ProviderName field to given value.

### HasProviderName

`func (o *ChannelMappingOptionsDto) HasProviderName() bool`

HasProviderName returns a boolean if a field has been set.

### SetProviderNameNil

`func (o *ChannelMappingOptionsDto) SetProviderNameNil(b bool)`

 SetProviderNameNil sets the value for ProviderName to be an explicit nil

### UnsetProviderName
`func (o *ChannelMappingOptionsDto) UnsetProviderName()`

UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


