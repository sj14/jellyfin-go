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

// checks if the RemoteLyricInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteLyricInfoDto{}

// RemoteLyricInfoDto The remote lyric info dto.
type RemoteLyricInfoDto struct {
	// Gets or sets the id for the lyric.
	Id *string `json:"Id,omitempty"`
	// Gets the provider name.
	ProviderName *string `json:"ProviderName,omitempty"`
	// Gets the lyrics.
	Lyrics *LyricDto `json:"Lyrics,omitempty"`
}

// NewRemoteLyricInfoDto instantiates a new RemoteLyricInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteLyricInfoDto() *RemoteLyricInfoDto {
	this := RemoteLyricInfoDto{}
	return &this
}

// NewRemoteLyricInfoDtoWithDefaults instantiates a new RemoteLyricInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteLyricInfoDtoWithDefaults() *RemoteLyricInfoDto {
	this := RemoteLyricInfoDto{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RemoteLyricInfoDto) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteLyricInfoDto) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RemoteLyricInfoDto) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RemoteLyricInfoDto) SetId(v string) {
	o.Id = &v
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise.
func (o *RemoteLyricInfoDto) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName) {
		var ret string
		return ret
	}
	return *o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteLyricInfoDto) GetProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderName) {
		return nil, false
	}
	return o.ProviderName, true
}

// HasProviderName returns a boolean if a field has been set.
func (o *RemoteLyricInfoDto) HasProviderName() bool {
	if o != nil && !IsNil(o.ProviderName) {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given string and assigns it to the ProviderName field.
func (o *RemoteLyricInfoDto) SetProviderName(v string) {
	o.ProviderName = &v
}

// GetLyrics returns the Lyrics field value if set, zero value otherwise.
func (o *RemoteLyricInfoDto) GetLyrics() LyricDto {
	if o == nil || IsNil(o.Lyrics) {
		var ret LyricDto
		return ret
	}
	return *o.Lyrics
}

// GetLyricsOk returns a tuple with the Lyrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoteLyricInfoDto) GetLyricsOk() (*LyricDto, bool) {
	if o == nil || IsNil(o.Lyrics) {
		return nil, false
	}
	return o.Lyrics, true
}

// HasLyrics returns a boolean if a field has been set.
func (o *RemoteLyricInfoDto) HasLyrics() bool {
	if o != nil && !IsNil(o.Lyrics) {
		return true
	}

	return false
}

// SetLyrics gets a reference to the given LyricDto and assigns it to the Lyrics field.
func (o *RemoteLyricInfoDto) SetLyrics(v LyricDto) {
	o.Lyrics = &v
}

func (o RemoteLyricInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteLyricInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.ProviderName) {
		toSerialize["ProviderName"] = o.ProviderName
	}
	if !IsNil(o.Lyrics) {
		toSerialize["Lyrics"] = o.Lyrics
	}
	return toSerialize, nil
}

type NullableRemoteLyricInfoDto struct {
	value *RemoteLyricInfoDto
	isSet bool
}

func (v NullableRemoteLyricInfoDto) Get() *RemoteLyricInfoDto {
	return v.value
}

func (v *NullableRemoteLyricInfoDto) Set(val *RemoteLyricInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteLyricInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteLyricInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteLyricInfoDto(val *RemoteLyricInfoDto) *NullableRemoteLyricInfoDto {
	return &NullableRemoteLyricInfoDto{value: val, isSet: true}
}

func (v NullableRemoteLyricInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteLyricInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


