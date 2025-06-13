# LyricLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** | Gets the text of this lyric line. | [optional] 
**Start** | Pointer to **NullableInt64** | Gets the start time in ticks. | [optional] 
**Cues** | Pointer to [**[]LyricLineCue**](LyricLineCue.md) | Gets the time-aligned cues for the song&#39;s lyrics. | [optional] 

## Methods

### NewLyricLine

`func NewLyricLine() *LyricLine`

NewLyricLine instantiates a new LyricLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLyricLineWithDefaults

`func NewLyricLineWithDefaults() *LyricLine`

NewLyricLineWithDefaults instantiates a new LyricLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *LyricLine) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *LyricLine) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *LyricLine) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *LyricLine) HasText() bool`

HasText returns a boolean if a field has been set.

### GetStart

`func (o *LyricLine) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *LyricLine) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *LyricLine) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *LyricLine) HasStart() bool`

HasStart returns a boolean if a field has been set.

### SetStartNil

`func (o *LyricLine) SetStartNil(b bool)`

 SetStartNil sets the value for Start to be an explicit nil

### UnsetStart
`func (o *LyricLine) UnsetStart()`

UnsetStart ensures that no value is present for Start, not even an explicit nil
### GetCues

`func (o *LyricLine) GetCues() []LyricLineCue`

GetCues returns the Cues field if non-nil, zero value otherwise.

### GetCuesOk

`func (o *LyricLine) GetCuesOk() (*[]LyricLineCue, bool)`

GetCuesOk returns a tuple with the Cues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCues

`func (o *LyricLine) SetCues(v []LyricLineCue)`

SetCues sets Cues field to given value.

### HasCues

`func (o *LyricLine) HasCues() bool`

HasCues returns a boolean if a field has been set.

### SetCuesNil

`func (o *LyricLine) SetCuesNil(b bool)`

 SetCuesNil sets the value for Cues to be an explicit nil

### UnsetCues
`func (o *LyricLine) UnsetCues()`

UnsetCues ensures that no value is present for Cues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


