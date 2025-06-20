/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the LyricDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LyricDto{}

// LyricDto LyricResponse model.
type LyricDto struct {
	// Gets or sets Metadata for the lyrics.
	Metadata *LyricMetadata `json:"Metadata,omitempty"`
	// Gets or sets a collection of individual lyric lines.
	Lyrics []LyricLine `json:"Lyrics,omitempty"`
}

// NewLyricDto instantiates a new LyricDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLyricDto() *LyricDto {
	this := LyricDto{}
	return &this
}

// NewLyricDtoWithDefaults instantiates a new LyricDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLyricDtoWithDefaults() *LyricDto {
	this := LyricDto{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *LyricDto) GetMetadata() LyricMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret LyricMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LyricDto) GetMetadataOk() (*LyricMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *LyricDto) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given LyricMetadata and assigns it to the Metadata field.
func (o *LyricDto) SetMetadata(v LyricMetadata) {
	o.Metadata = &v
}

// GetLyrics returns the Lyrics field value if set, zero value otherwise.
func (o *LyricDto) GetLyrics() []LyricLine {
	if o == nil || IsNil(o.Lyrics) {
		var ret []LyricLine
		return ret
	}
	return o.Lyrics
}

// GetLyricsOk returns a tuple with the Lyrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LyricDto) GetLyricsOk() ([]LyricLine, bool) {
	if o == nil || IsNil(o.Lyrics) {
		return nil, false
	}
	return o.Lyrics, true
}

// HasLyrics returns a boolean if a field has been set.
func (o *LyricDto) HasLyrics() bool {
	if o != nil && !IsNil(o.Lyrics) {
		return true
	}

	return false
}

// SetLyrics gets a reference to the given []LyricLine and assigns it to the Lyrics field.
func (o *LyricDto) SetLyrics(v []LyricLine) {
	o.Lyrics = v
}

func (o LyricDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LyricDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["Metadata"] = o.Metadata
	}
	if !IsNil(o.Lyrics) {
		toSerialize["Lyrics"] = o.Lyrics
	}
	return toSerialize, nil
}

type NullableLyricDto struct {
	value *LyricDto
	isSet bool
}

func (v NullableLyricDto) Get() *LyricDto {
	return v.value
}

func (v *NullableLyricDto) Set(val *LyricDto) {
	v.value = val
	v.isSet = true
}

func (v NullableLyricDto) IsSet() bool {
	return v.isSet
}

func (v *NullableLyricDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLyricDto(val *LyricDto) *NullableLyricDto {
	return &NullableLyricDto{value: val, isSet: true}
}

func (v NullableLyricDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLyricDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


