# MediaSegmentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Gets or sets the id of the media segment. | [optional] 
**ItemId** | Pointer to **string** | Gets or sets the id of the associated item. | [optional] 
**Type** | Pointer to [**MediaSegmentType**](MediaSegmentType.md) | Defines the types of content an individual Jellyfin.Data.Entities.MediaSegment represents. | [optional] 
**StartTicks** | Pointer to **int64** | Gets or sets the start of the segment. | [optional] 
**EndTicks** | Pointer to **int64** | Gets or sets the end of the segment. | [optional] 

## Methods

### NewMediaSegmentDto

`func NewMediaSegmentDto() *MediaSegmentDto`

NewMediaSegmentDto instantiates a new MediaSegmentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaSegmentDtoWithDefaults

`func NewMediaSegmentDtoWithDefaults() *MediaSegmentDto`

NewMediaSegmentDtoWithDefaults instantiates a new MediaSegmentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MediaSegmentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MediaSegmentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MediaSegmentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MediaSegmentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItemId

`func (o *MediaSegmentDto) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *MediaSegmentDto) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *MediaSegmentDto) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *MediaSegmentDto) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetType

`func (o *MediaSegmentDto) GetType() MediaSegmentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MediaSegmentDto) GetTypeOk() (*MediaSegmentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MediaSegmentDto) SetType(v MediaSegmentType)`

SetType sets Type field to given value.

### HasType

`func (o *MediaSegmentDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStartTicks

`func (o *MediaSegmentDto) GetStartTicks() int64`

GetStartTicks returns the StartTicks field if non-nil, zero value otherwise.

### GetStartTicksOk

`func (o *MediaSegmentDto) GetStartTicksOk() (*int64, bool)`

GetStartTicksOk returns a tuple with the StartTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTicks

`func (o *MediaSegmentDto) SetStartTicks(v int64)`

SetStartTicks sets StartTicks field to given value.

### HasStartTicks

`func (o *MediaSegmentDto) HasStartTicks() bool`

HasStartTicks returns a boolean if a field has been set.

### GetEndTicks

`func (o *MediaSegmentDto) GetEndTicks() int64`

GetEndTicks returns the EndTicks field if non-nil, zero value otherwise.

### GetEndTicksOk

`func (o *MediaSegmentDto) GetEndTicksOk() (*int64, bool)`

GetEndTicksOk returns a tuple with the EndTicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTicks

`func (o *MediaSegmentDto) SetEndTicks(v int64)`

SetEndTicks sets EndTicks field to given value.

### HasEndTicks

`func (o *MediaSegmentDto) HasEndTicks() bool`

HasEndTicks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


