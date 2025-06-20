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

// checks if the LibraryOptionsResultDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LibraryOptionsResultDto{}

// LibraryOptionsResultDto Library options result dto.
type LibraryOptionsResultDto struct {
	// Gets or sets the metadata savers.
	MetadataSavers []LibraryOptionInfoDto `json:"MetadataSavers,omitempty"`
	// Gets or sets the metadata readers.
	MetadataReaders []LibraryOptionInfoDto `json:"MetadataReaders,omitempty"`
	// Gets or sets the subtitle fetchers.
	SubtitleFetchers []LibraryOptionInfoDto `json:"SubtitleFetchers,omitempty"`
	// Gets or sets the list of lyric fetchers.
	LyricFetchers []LibraryOptionInfoDto `json:"LyricFetchers,omitempty"`
	// Gets or sets the list of MediaSegment Providers.
	MediaSegmentProviders []LibraryOptionInfoDto `json:"MediaSegmentProviders,omitempty"`
	// Gets or sets the type options.
	TypeOptions []LibraryTypeOptionsDto `json:"TypeOptions,omitempty"`
}

// NewLibraryOptionsResultDto instantiates a new LibraryOptionsResultDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLibraryOptionsResultDto() *LibraryOptionsResultDto {
	this := LibraryOptionsResultDto{}
	return &this
}

// NewLibraryOptionsResultDtoWithDefaults instantiates a new LibraryOptionsResultDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLibraryOptionsResultDtoWithDefaults() *LibraryOptionsResultDto {
	this := LibraryOptionsResultDto{}
	return &this
}

// GetMetadataSavers returns the MetadataSavers field value if set, zero value otherwise.
func (o *LibraryOptionsResultDto) GetMetadataSavers() []LibraryOptionInfoDto {
	if o == nil || IsNil(o.MetadataSavers) {
		var ret []LibraryOptionInfoDto
		return ret
	}
	return o.MetadataSavers
}

// GetMetadataSaversOk returns a tuple with the MetadataSavers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptionsResultDto) GetMetadataSaversOk() ([]LibraryOptionInfoDto, bool) {
	if o == nil || IsNil(o.MetadataSavers) {
		return nil, false
	}
	return o.MetadataSavers, true
}

// HasMetadataSavers returns a boolean if a field has been set.
func (o *LibraryOptionsResultDto) HasMetadataSavers() bool {
	if o != nil && !IsNil(o.MetadataSavers) {
		return true
	}

	return false
}

// SetMetadataSavers gets a reference to the given []LibraryOptionInfoDto and assigns it to the MetadataSavers field.
func (o *LibraryOptionsResultDto) SetMetadataSavers(v []LibraryOptionInfoDto) {
	o.MetadataSavers = v
}

// GetMetadataReaders returns the MetadataReaders field value if set, zero value otherwise.
func (o *LibraryOptionsResultDto) GetMetadataReaders() []LibraryOptionInfoDto {
	if o == nil || IsNil(o.MetadataReaders) {
		var ret []LibraryOptionInfoDto
		return ret
	}
	return o.MetadataReaders
}

// GetMetadataReadersOk returns a tuple with the MetadataReaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptionsResultDto) GetMetadataReadersOk() ([]LibraryOptionInfoDto, bool) {
	if o == nil || IsNil(o.MetadataReaders) {
		return nil, false
	}
	return o.MetadataReaders, true
}

// HasMetadataReaders returns a boolean if a field has been set.
func (o *LibraryOptionsResultDto) HasMetadataReaders() bool {
	if o != nil && !IsNil(o.MetadataReaders) {
		return true
	}

	return false
}

// SetMetadataReaders gets a reference to the given []LibraryOptionInfoDto and assigns it to the MetadataReaders field.
func (o *LibraryOptionsResultDto) SetMetadataReaders(v []LibraryOptionInfoDto) {
	o.MetadataReaders = v
}

// GetSubtitleFetchers returns the SubtitleFetchers field value if set, zero value otherwise.
func (o *LibraryOptionsResultDto) GetSubtitleFetchers() []LibraryOptionInfoDto {
	if o == nil || IsNil(o.SubtitleFetchers) {
		var ret []LibraryOptionInfoDto
		return ret
	}
	return o.SubtitleFetchers
}

// GetSubtitleFetchersOk returns a tuple with the SubtitleFetchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptionsResultDto) GetSubtitleFetchersOk() ([]LibraryOptionInfoDto, bool) {
	if o == nil || IsNil(o.SubtitleFetchers) {
		return nil, false
	}
	return o.SubtitleFetchers, true
}

// HasSubtitleFetchers returns a boolean if a field has been set.
func (o *LibraryOptionsResultDto) HasSubtitleFetchers() bool {
	if o != nil && !IsNil(o.SubtitleFetchers) {
		return true
	}

	return false
}

// SetSubtitleFetchers gets a reference to the given []LibraryOptionInfoDto and assigns it to the SubtitleFetchers field.
func (o *LibraryOptionsResultDto) SetSubtitleFetchers(v []LibraryOptionInfoDto) {
	o.SubtitleFetchers = v
}

// GetLyricFetchers returns the LyricFetchers field value if set, zero value otherwise.
func (o *LibraryOptionsResultDto) GetLyricFetchers() []LibraryOptionInfoDto {
	if o == nil || IsNil(o.LyricFetchers) {
		var ret []LibraryOptionInfoDto
		return ret
	}
	return o.LyricFetchers
}

// GetLyricFetchersOk returns a tuple with the LyricFetchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptionsResultDto) GetLyricFetchersOk() ([]LibraryOptionInfoDto, bool) {
	if o == nil || IsNil(o.LyricFetchers) {
		return nil, false
	}
	return o.LyricFetchers, true
}

// HasLyricFetchers returns a boolean if a field has been set.
func (o *LibraryOptionsResultDto) HasLyricFetchers() bool {
	if o != nil && !IsNil(o.LyricFetchers) {
		return true
	}

	return false
}

// SetLyricFetchers gets a reference to the given []LibraryOptionInfoDto and assigns it to the LyricFetchers field.
func (o *LibraryOptionsResultDto) SetLyricFetchers(v []LibraryOptionInfoDto) {
	o.LyricFetchers = v
}

// GetMediaSegmentProviders returns the MediaSegmentProviders field value if set, zero value otherwise.
func (o *LibraryOptionsResultDto) GetMediaSegmentProviders() []LibraryOptionInfoDto {
	if o == nil || IsNil(o.MediaSegmentProviders) {
		var ret []LibraryOptionInfoDto
		return ret
	}
	return o.MediaSegmentProviders
}

// GetMediaSegmentProvidersOk returns a tuple with the MediaSegmentProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptionsResultDto) GetMediaSegmentProvidersOk() ([]LibraryOptionInfoDto, bool) {
	if o == nil || IsNil(o.MediaSegmentProviders) {
		return nil, false
	}
	return o.MediaSegmentProviders, true
}

// HasMediaSegmentProviders returns a boolean if a field has been set.
func (o *LibraryOptionsResultDto) HasMediaSegmentProviders() bool {
	if o != nil && !IsNil(o.MediaSegmentProviders) {
		return true
	}

	return false
}

// SetMediaSegmentProviders gets a reference to the given []LibraryOptionInfoDto and assigns it to the MediaSegmentProviders field.
func (o *LibraryOptionsResultDto) SetMediaSegmentProviders(v []LibraryOptionInfoDto) {
	o.MediaSegmentProviders = v
}

// GetTypeOptions returns the TypeOptions field value if set, zero value otherwise.
func (o *LibraryOptionsResultDto) GetTypeOptions() []LibraryTypeOptionsDto {
	if o == nil || IsNil(o.TypeOptions) {
		var ret []LibraryTypeOptionsDto
		return ret
	}
	return o.TypeOptions
}

// GetTypeOptionsOk returns a tuple with the TypeOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptionsResultDto) GetTypeOptionsOk() ([]LibraryTypeOptionsDto, bool) {
	if o == nil || IsNil(o.TypeOptions) {
		return nil, false
	}
	return o.TypeOptions, true
}

// HasTypeOptions returns a boolean if a field has been set.
func (o *LibraryOptionsResultDto) HasTypeOptions() bool {
	if o != nil && !IsNil(o.TypeOptions) {
		return true
	}

	return false
}

// SetTypeOptions gets a reference to the given []LibraryTypeOptionsDto and assigns it to the TypeOptions field.
func (o *LibraryOptionsResultDto) SetTypeOptions(v []LibraryTypeOptionsDto) {
	o.TypeOptions = v
}

func (o LibraryOptionsResultDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LibraryOptionsResultDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MetadataSavers) {
		toSerialize["MetadataSavers"] = o.MetadataSavers
	}
	if !IsNil(o.MetadataReaders) {
		toSerialize["MetadataReaders"] = o.MetadataReaders
	}
	if !IsNil(o.SubtitleFetchers) {
		toSerialize["SubtitleFetchers"] = o.SubtitleFetchers
	}
	if !IsNil(o.LyricFetchers) {
		toSerialize["LyricFetchers"] = o.LyricFetchers
	}
	if !IsNil(o.MediaSegmentProviders) {
		toSerialize["MediaSegmentProviders"] = o.MediaSegmentProviders
	}
	if !IsNil(o.TypeOptions) {
		toSerialize["TypeOptions"] = o.TypeOptions
	}
	return toSerialize, nil
}

type NullableLibraryOptionsResultDto struct {
	value *LibraryOptionsResultDto
	isSet bool
}

func (v NullableLibraryOptionsResultDto) Get() *LibraryOptionsResultDto {
	return v.value
}

func (v *NullableLibraryOptionsResultDto) Set(val *LibraryOptionsResultDto) {
	v.value = val
	v.isSet = true
}

func (v NullableLibraryOptionsResultDto) IsSet() bool {
	return v.isSet
}

func (v *NullableLibraryOptionsResultDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLibraryOptionsResultDto(val *LibraryOptionsResultDto) *NullableLibraryOptionsResultDto {
	return &NullableLibraryOptionsResultDto{value: val, isSet: true}
}

func (v NullableLibraryOptionsResultDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLibraryOptionsResultDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


