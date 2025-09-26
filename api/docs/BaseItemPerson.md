# BaseItemPerson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name. | [optional] 
**Id** | Pointer to **string** | Gets or sets the identifier. | [optional] 
**Role** | Pointer to **NullableString** | Gets or sets the role. | [optional] 
**Type** | Pointer to [**PersonKind**](PersonKind.md) | Gets or sets the type. | [optional] [default to PERSONKIND_UNKNOWN]
**PrimaryImageTag** | Pointer to **NullableString** | Gets or sets the primary image tag. | [optional] 
**ImageBlurHashes** | Pointer to [**NullableBaseItemPersonImageBlurHashes**](BaseItemPersonImageBlurHashes.md) |  | [optional] 

## Methods

### NewBaseItemPerson

`func NewBaseItemPerson() *BaseItemPerson`

NewBaseItemPerson instantiates a new BaseItemPerson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseItemPersonWithDefaults

`func NewBaseItemPersonWithDefaults() *BaseItemPerson`

NewBaseItemPersonWithDefaults instantiates a new BaseItemPerson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BaseItemPerson) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseItemPerson) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseItemPerson) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseItemPerson) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BaseItemPerson) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BaseItemPerson) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetId

`func (o *BaseItemPerson) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseItemPerson) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseItemPerson) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseItemPerson) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRole

`func (o *BaseItemPerson) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *BaseItemPerson) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *BaseItemPerson) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *BaseItemPerson) HasRole() bool`

HasRole returns a boolean if a field has been set.

### SetRoleNil

`func (o *BaseItemPerson) SetRoleNil(b bool)`

 SetRoleNil sets the value for Role to be an explicit nil

### UnsetRole
`func (o *BaseItemPerson) UnsetRole()`

UnsetRole ensures that no value is present for Role, not even an explicit nil
### GetType

`func (o *BaseItemPerson) GetType() PersonKind`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BaseItemPerson) GetTypeOk() (*PersonKind, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BaseItemPerson) SetType(v PersonKind)`

SetType sets Type field to given value.

### HasType

`func (o *BaseItemPerson) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPrimaryImageTag

`func (o *BaseItemPerson) GetPrimaryImageTag() string`

GetPrimaryImageTag returns the PrimaryImageTag field if non-nil, zero value otherwise.

### GetPrimaryImageTagOk

`func (o *BaseItemPerson) GetPrimaryImageTagOk() (*string, bool)`

GetPrimaryImageTagOk returns a tuple with the PrimaryImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryImageTag

`func (o *BaseItemPerson) SetPrimaryImageTag(v string)`

SetPrimaryImageTag sets PrimaryImageTag field to given value.

### HasPrimaryImageTag

`func (o *BaseItemPerson) HasPrimaryImageTag() bool`

HasPrimaryImageTag returns a boolean if a field has been set.

### SetPrimaryImageTagNil

`func (o *BaseItemPerson) SetPrimaryImageTagNil(b bool)`

 SetPrimaryImageTagNil sets the value for PrimaryImageTag to be an explicit nil

### UnsetPrimaryImageTag
`func (o *BaseItemPerson) UnsetPrimaryImageTag()`

UnsetPrimaryImageTag ensures that no value is present for PrimaryImageTag, not even an explicit nil
### GetImageBlurHashes

`func (o *BaseItemPerson) GetImageBlurHashes() BaseItemPersonImageBlurHashes`

GetImageBlurHashes returns the ImageBlurHashes field if non-nil, zero value otherwise.

### GetImageBlurHashesOk

`func (o *BaseItemPerson) GetImageBlurHashesOk() (*BaseItemPersonImageBlurHashes, bool)`

GetImageBlurHashesOk returns a tuple with the ImageBlurHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBlurHashes

`func (o *BaseItemPerson) SetImageBlurHashes(v BaseItemPersonImageBlurHashes)`

SetImageBlurHashes sets ImageBlurHashes field to given value.

### HasImageBlurHashes

`func (o *BaseItemPerson) HasImageBlurHashes() bool`

HasImageBlurHashes returns a boolean if a field has been set.

### SetImageBlurHashesNil

`func (o *BaseItemPerson) SetImageBlurHashesNil(b bool)`

 SetImageBlurHashesNil sets the value for ImageBlurHashes to be an explicit nil

### UnsetImageBlurHashes
`func (o *BaseItemPerson) UnsetImageBlurHashes()`

UnsetImageBlurHashes ensures that no value is present for ImageBlurHashes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


