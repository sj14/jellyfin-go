# FontFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Gets or sets the name. | [optional] 
**Size** | Pointer to **int64** | Gets or sets the size. | [optional] 
**DateCreated** | Pointer to **time.Time** | Gets or sets the date created. | [optional] 
**DateModified** | Pointer to **time.Time** | Gets or sets the date modified. | [optional] 

## Methods

### NewFontFile

`func NewFontFile() *FontFile`

NewFontFile instantiates a new FontFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFontFileWithDefaults

`func NewFontFileWithDefaults() *FontFile`

NewFontFileWithDefaults instantiates a new FontFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FontFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FontFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FontFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FontFile) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *FontFile) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *FontFile) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSize

`func (o *FontFile) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FontFile) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FontFile) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *FontFile) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetDateCreated

`func (o *FontFile) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *FontFile) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *FontFile) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *FontFile) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateModified

`func (o *FontFile) GetDateModified() time.Time`

GetDateModified returns the DateModified field if non-nil, zero value otherwise.

### GetDateModifiedOk

`func (o *FontFile) GetDateModifiedOk() (*time.Time, bool)`

GetDateModifiedOk returns a tuple with the DateModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateModified

`func (o *FontFile) SetDateModified(v time.Time)`

SetDateModified sets DateModified field to given value.

### HasDateModified

`func (o *FontFile) HasDateModified() bool`

HasDateModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


