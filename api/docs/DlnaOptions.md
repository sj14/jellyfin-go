# DlnaOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnablePlayTo** | Pointer to **bool** | Gets or sets a value indicating whether gets or sets a value to indicate the status of the dlna playTo subsystem. | [optional] 
**EnableServer** | Pointer to **bool** | Gets or sets a value indicating whether gets or sets a value to indicate the status of the dlna server subsystem. | [optional] 
**EnableDebugLog** | Pointer to **bool** | Gets or sets a value indicating whether detailed dlna server logs are sent to the console/log.  If the setting \&quot;Emby.Dlna\&quot;: \&quot;Debug\&quot; msut be set in logging.default.json for this property to work. | [optional] 
**EnablePlayToTracing** | Pointer to **bool** | Gets or sets a value indicating whether whether detailed playTo debug logs are sent to the console/log.  If the setting \&quot;Emby.Dlna.PlayTo\&quot;: \&quot;Debug\&quot; msut be set in logging.default.json for this property to work. | [optional] 
**ClientDiscoveryIntervalSeconds** | Pointer to **int32** | Gets or sets the ssdp client discovery interval time (in seconds).  This is the time after which the server will send a ssdp search request. | [optional] 
**AliveMessageIntervalSeconds** | Pointer to **int32** | Gets or sets the frequency at which ssdp alive notifications are transmitted. | [optional] 
**BlastAliveMessageIntervalSeconds** | Pointer to **int32** | Gets or sets the frequency at which ssdp alive notifications are transmitted. MIGRATING - TO BE REMOVED ONCE WEB HAS BEEN ALTERED. | [optional] 
**DefaultUserId** | Pointer to **NullableString** | Gets or sets the default user account that the dlna server uses. | [optional] 
**AutoCreatePlayToProfiles** | Pointer to **bool** | Gets or sets a value indicating whether playTo device profiles should be created. | [optional] 
**BlastAliveMessages** | Pointer to **bool** | Gets or sets a value indicating whether to blast alive messages. | [optional] 
**SendOnlyMatchedHost** | Pointer to **bool** | gets or sets a value indicating whether to send only matched host. | [optional] 

## Methods

### NewDlnaOptions

`func NewDlnaOptions() *DlnaOptions`

NewDlnaOptions instantiates a new DlnaOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDlnaOptionsWithDefaults

`func NewDlnaOptionsWithDefaults() *DlnaOptions`

NewDlnaOptionsWithDefaults instantiates a new DlnaOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnablePlayTo

`func (o *DlnaOptions) GetEnablePlayTo() bool`

GetEnablePlayTo returns the EnablePlayTo field if non-nil, zero value otherwise.

### GetEnablePlayToOk

`func (o *DlnaOptions) GetEnablePlayToOk() (*bool, bool)`

GetEnablePlayToOk returns a tuple with the EnablePlayTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePlayTo

`func (o *DlnaOptions) SetEnablePlayTo(v bool)`

SetEnablePlayTo sets EnablePlayTo field to given value.

### HasEnablePlayTo

`func (o *DlnaOptions) HasEnablePlayTo() bool`

HasEnablePlayTo returns a boolean if a field has been set.

### GetEnableServer

`func (o *DlnaOptions) GetEnableServer() bool`

GetEnableServer returns the EnableServer field if non-nil, zero value otherwise.

### GetEnableServerOk

`func (o *DlnaOptions) GetEnableServerOk() (*bool, bool)`

GetEnableServerOk returns a tuple with the EnableServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableServer

`func (o *DlnaOptions) SetEnableServer(v bool)`

SetEnableServer sets EnableServer field to given value.

### HasEnableServer

`func (o *DlnaOptions) HasEnableServer() bool`

HasEnableServer returns a boolean if a field has been set.

### GetEnableDebugLog

`func (o *DlnaOptions) GetEnableDebugLog() bool`

GetEnableDebugLog returns the EnableDebugLog field if non-nil, zero value otherwise.

### GetEnableDebugLogOk

`func (o *DlnaOptions) GetEnableDebugLogOk() (*bool, bool)`

GetEnableDebugLogOk returns a tuple with the EnableDebugLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDebugLog

`func (o *DlnaOptions) SetEnableDebugLog(v bool)`

SetEnableDebugLog sets EnableDebugLog field to given value.

### HasEnableDebugLog

`func (o *DlnaOptions) HasEnableDebugLog() bool`

HasEnableDebugLog returns a boolean if a field has been set.

### GetEnablePlayToTracing

`func (o *DlnaOptions) GetEnablePlayToTracing() bool`

GetEnablePlayToTracing returns the EnablePlayToTracing field if non-nil, zero value otherwise.

### GetEnablePlayToTracingOk

`func (o *DlnaOptions) GetEnablePlayToTracingOk() (*bool, bool)`

GetEnablePlayToTracingOk returns a tuple with the EnablePlayToTracing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePlayToTracing

`func (o *DlnaOptions) SetEnablePlayToTracing(v bool)`

SetEnablePlayToTracing sets EnablePlayToTracing field to given value.

### HasEnablePlayToTracing

`func (o *DlnaOptions) HasEnablePlayToTracing() bool`

HasEnablePlayToTracing returns a boolean if a field has been set.

### GetClientDiscoveryIntervalSeconds

`func (o *DlnaOptions) GetClientDiscoveryIntervalSeconds() int32`

GetClientDiscoveryIntervalSeconds returns the ClientDiscoveryIntervalSeconds field if non-nil, zero value otherwise.

### GetClientDiscoveryIntervalSecondsOk

`func (o *DlnaOptions) GetClientDiscoveryIntervalSecondsOk() (*int32, bool)`

GetClientDiscoveryIntervalSecondsOk returns a tuple with the ClientDiscoveryIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDiscoveryIntervalSeconds

`func (o *DlnaOptions) SetClientDiscoveryIntervalSeconds(v int32)`

SetClientDiscoveryIntervalSeconds sets ClientDiscoveryIntervalSeconds field to given value.

### HasClientDiscoveryIntervalSeconds

`func (o *DlnaOptions) HasClientDiscoveryIntervalSeconds() bool`

HasClientDiscoveryIntervalSeconds returns a boolean if a field has been set.

### GetAliveMessageIntervalSeconds

`func (o *DlnaOptions) GetAliveMessageIntervalSeconds() int32`

GetAliveMessageIntervalSeconds returns the AliveMessageIntervalSeconds field if non-nil, zero value otherwise.

### GetAliveMessageIntervalSecondsOk

`func (o *DlnaOptions) GetAliveMessageIntervalSecondsOk() (*int32, bool)`

GetAliveMessageIntervalSecondsOk returns a tuple with the AliveMessageIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliveMessageIntervalSeconds

`func (o *DlnaOptions) SetAliveMessageIntervalSeconds(v int32)`

SetAliveMessageIntervalSeconds sets AliveMessageIntervalSeconds field to given value.

### HasAliveMessageIntervalSeconds

`func (o *DlnaOptions) HasAliveMessageIntervalSeconds() bool`

HasAliveMessageIntervalSeconds returns a boolean if a field has been set.

### GetBlastAliveMessageIntervalSeconds

`func (o *DlnaOptions) GetBlastAliveMessageIntervalSeconds() int32`

GetBlastAliveMessageIntervalSeconds returns the BlastAliveMessageIntervalSeconds field if non-nil, zero value otherwise.

### GetBlastAliveMessageIntervalSecondsOk

`func (o *DlnaOptions) GetBlastAliveMessageIntervalSecondsOk() (*int32, bool)`

GetBlastAliveMessageIntervalSecondsOk returns a tuple with the BlastAliveMessageIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlastAliveMessageIntervalSeconds

`func (o *DlnaOptions) SetBlastAliveMessageIntervalSeconds(v int32)`

SetBlastAliveMessageIntervalSeconds sets BlastAliveMessageIntervalSeconds field to given value.

### HasBlastAliveMessageIntervalSeconds

`func (o *DlnaOptions) HasBlastAliveMessageIntervalSeconds() bool`

HasBlastAliveMessageIntervalSeconds returns a boolean if a field has been set.

### GetDefaultUserId

`func (o *DlnaOptions) GetDefaultUserId() string`

GetDefaultUserId returns the DefaultUserId field if non-nil, zero value otherwise.

### GetDefaultUserIdOk

`func (o *DlnaOptions) GetDefaultUserIdOk() (*string, bool)`

GetDefaultUserIdOk returns a tuple with the DefaultUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultUserId

`func (o *DlnaOptions) SetDefaultUserId(v string)`

SetDefaultUserId sets DefaultUserId field to given value.

### HasDefaultUserId

`func (o *DlnaOptions) HasDefaultUserId() bool`

HasDefaultUserId returns a boolean if a field has been set.

### SetDefaultUserIdNil

`func (o *DlnaOptions) SetDefaultUserIdNil(b bool)`

 SetDefaultUserIdNil sets the value for DefaultUserId to be an explicit nil

### UnsetDefaultUserId
`func (o *DlnaOptions) UnsetDefaultUserId()`

UnsetDefaultUserId ensures that no value is present for DefaultUserId, not even an explicit nil
### GetAutoCreatePlayToProfiles

`func (o *DlnaOptions) GetAutoCreatePlayToProfiles() bool`

GetAutoCreatePlayToProfiles returns the AutoCreatePlayToProfiles field if non-nil, zero value otherwise.

### GetAutoCreatePlayToProfilesOk

`func (o *DlnaOptions) GetAutoCreatePlayToProfilesOk() (*bool, bool)`

GetAutoCreatePlayToProfilesOk returns a tuple with the AutoCreatePlayToProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreatePlayToProfiles

`func (o *DlnaOptions) SetAutoCreatePlayToProfiles(v bool)`

SetAutoCreatePlayToProfiles sets AutoCreatePlayToProfiles field to given value.

### HasAutoCreatePlayToProfiles

`func (o *DlnaOptions) HasAutoCreatePlayToProfiles() bool`

HasAutoCreatePlayToProfiles returns a boolean if a field has been set.

### GetBlastAliveMessages

`func (o *DlnaOptions) GetBlastAliveMessages() bool`

GetBlastAliveMessages returns the BlastAliveMessages field if non-nil, zero value otherwise.

### GetBlastAliveMessagesOk

`func (o *DlnaOptions) GetBlastAliveMessagesOk() (*bool, bool)`

GetBlastAliveMessagesOk returns a tuple with the BlastAliveMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlastAliveMessages

`func (o *DlnaOptions) SetBlastAliveMessages(v bool)`

SetBlastAliveMessages sets BlastAliveMessages field to given value.

### HasBlastAliveMessages

`func (o *DlnaOptions) HasBlastAliveMessages() bool`

HasBlastAliveMessages returns a boolean if a field has been set.

### GetSendOnlyMatchedHost

`func (o *DlnaOptions) GetSendOnlyMatchedHost() bool`

GetSendOnlyMatchedHost returns the SendOnlyMatchedHost field if non-nil, zero value otherwise.

### GetSendOnlyMatchedHostOk

`func (o *DlnaOptions) GetSendOnlyMatchedHostOk() (*bool, bool)`

GetSendOnlyMatchedHostOk returns a tuple with the SendOnlyMatchedHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendOnlyMatchedHost

`func (o *DlnaOptions) SetSendOnlyMatchedHost(v bool)`

SetSendOnlyMatchedHost sets SendOnlyMatchedHost field to given value.

### HasSendOnlyMatchedHost

`func (o *DlnaOptions) HasSendOnlyMatchedHost() bool`

HasSendOnlyMatchedHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


