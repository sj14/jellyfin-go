# LyricDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**LyricMetadata**](LyricMetadata.md) | Gets or sets Metadata for the lyrics. | [optional] 
**Lyrics** | Pointer to [**[]LyricLine**](LyricLine.md) | Gets or sets a collection of individual lyric lines. | [optional] 

## Methods

### NewLyricDto

`func NewLyricDto() *LyricDto`

NewLyricDto instantiates a new LyricDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLyricDtoWithDefaults

`func NewLyricDtoWithDefaults() *LyricDto`

NewLyricDtoWithDefaults instantiates a new LyricDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *LyricDto) GetMetadata() LyricMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *LyricDto) GetMetadataOk() (*LyricMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *LyricDto) SetMetadata(v LyricMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *LyricDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetLyrics

`func (o *LyricDto) GetLyrics() []LyricLine`

GetLyrics returns the Lyrics field if non-nil, zero value otherwise.

### GetLyricsOk

`func (o *LyricDto) GetLyricsOk() (*[]LyricLine, bool)`

GetLyricsOk returns a tuple with the Lyrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLyrics

`func (o *LyricDto) SetLyrics(v []LyricLine)`

SetLyrics sets Lyrics field to given value.

### HasLyrics

`func (o *LyricDto) HasLyrics() bool`

HasLyrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


