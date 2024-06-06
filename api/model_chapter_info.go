/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the ChapterInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChapterInfo{}

// ChapterInfo Class ChapterInfo.
type ChapterInfo struct {
	// Gets or sets the start position ticks.
	StartPositionTicks *int64 `json:"StartPositionTicks,omitempty"`
	// Gets or sets the name.
	Name NullableString `json:"Name,omitempty"`
	// Gets or sets the image path.
	ImagePath NullableString `json:"ImagePath,omitempty"`
	ImageDateModified *time.Time `json:"ImageDateModified,omitempty"`
	ImageTag NullableString `json:"ImageTag,omitempty"`
}

// NewChapterInfo instantiates a new ChapterInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChapterInfo() *ChapterInfo {
	this := ChapterInfo{}
	return &this
}

// NewChapterInfoWithDefaults instantiates a new ChapterInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChapterInfoWithDefaults() *ChapterInfo {
	this := ChapterInfo{}
	return &this
}

// GetStartPositionTicks returns the StartPositionTicks field value if set, zero value otherwise.
func (o *ChapterInfo) GetStartPositionTicks() int64 {
	if o == nil || IsNil(o.StartPositionTicks) {
		var ret int64
		return ret
	}
	return *o.StartPositionTicks
}

// GetStartPositionTicksOk returns a tuple with the StartPositionTicks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChapterInfo) GetStartPositionTicksOk() (*int64, bool) {
	if o == nil || IsNil(o.StartPositionTicks) {
		return nil, false
	}
	return o.StartPositionTicks, true
}

// HasStartPositionTicks returns a boolean if a field has been set.
func (o *ChapterInfo) HasStartPositionTicks() bool {
	if o != nil && !IsNil(o.StartPositionTicks) {
		return true
	}

	return false
}

// SetStartPositionTicks gets a reference to the given int64 and assigns it to the StartPositionTicks field.
func (o *ChapterInfo) SetStartPositionTicks(v int64) {
	o.StartPositionTicks = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChapterInfo) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChapterInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ChapterInfo) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ChapterInfo) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ChapterInfo) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ChapterInfo) UnsetName() {
	o.Name.Unset()
}

// GetImagePath returns the ImagePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChapterInfo) GetImagePath() string {
	if o == nil || IsNil(o.ImagePath.Get()) {
		var ret string
		return ret
	}
	return *o.ImagePath.Get()
}

// GetImagePathOk returns a tuple with the ImagePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChapterInfo) GetImagePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImagePath.Get(), o.ImagePath.IsSet()
}

// HasImagePath returns a boolean if a field has been set.
func (o *ChapterInfo) HasImagePath() bool {
	if o != nil && o.ImagePath.IsSet() {
		return true
	}

	return false
}

// SetImagePath gets a reference to the given NullableString and assigns it to the ImagePath field.
func (o *ChapterInfo) SetImagePath(v string) {
	o.ImagePath.Set(&v)
}
// SetImagePathNil sets the value for ImagePath to be an explicit nil
func (o *ChapterInfo) SetImagePathNil() {
	o.ImagePath.Set(nil)
}

// UnsetImagePath ensures that no value is present for ImagePath, not even an explicit nil
func (o *ChapterInfo) UnsetImagePath() {
	o.ImagePath.Unset()
}

// GetImageDateModified returns the ImageDateModified field value if set, zero value otherwise.
func (o *ChapterInfo) GetImageDateModified() time.Time {
	if o == nil || IsNil(o.ImageDateModified) {
		var ret time.Time
		return ret
	}
	return *o.ImageDateModified
}

// GetImageDateModifiedOk returns a tuple with the ImageDateModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChapterInfo) GetImageDateModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ImageDateModified) {
		return nil, false
	}
	return o.ImageDateModified, true
}

// HasImageDateModified returns a boolean if a field has been set.
func (o *ChapterInfo) HasImageDateModified() bool {
	if o != nil && !IsNil(o.ImageDateModified) {
		return true
	}

	return false
}

// SetImageDateModified gets a reference to the given time.Time and assigns it to the ImageDateModified field.
func (o *ChapterInfo) SetImageDateModified(v time.Time) {
	o.ImageDateModified = &v
}

// GetImageTag returns the ImageTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChapterInfo) GetImageTag() string {
	if o == nil || IsNil(o.ImageTag.Get()) {
		var ret string
		return ret
	}
	return *o.ImageTag.Get()
}

// GetImageTagOk returns a tuple with the ImageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChapterInfo) GetImageTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageTag.Get(), o.ImageTag.IsSet()
}

// HasImageTag returns a boolean if a field has been set.
func (o *ChapterInfo) HasImageTag() bool {
	if o != nil && o.ImageTag.IsSet() {
		return true
	}

	return false
}

// SetImageTag gets a reference to the given NullableString and assigns it to the ImageTag field.
func (o *ChapterInfo) SetImageTag(v string) {
	o.ImageTag.Set(&v)
}
// SetImageTagNil sets the value for ImageTag to be an explicit nil
func (o *ChapterInfo) SetImageTagNil() {
	o.ImageTag.Set(nil)
}

// UnsetImageTag ensures that no value is present for ImageTag, not even an explicit nil
func (o *ChapterInfo) UnsetImageTag() {
	o.ImageTag.Unset()
}

func (o ChapterInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChapterInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartPositionTicks) {
		toSerialize["StartPositionTicks"] = o.StartPositionTicks
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.ImagePath.IsSet() {
		toSerialize["ImagePath"] = o.ImagePath.Get()
	}
	if !IsNil(o.ImageDateModified) {
		toSerialize["ImageDateModified"] = o.ImageDateModified
	}
	if o.ImageTag.IsSet() {
		toSerialize["ImageTag"] = o.ImageTag.Get()
	}
	return toSerialize, nil
}

type NullableChapterInfo struct {
	value *ChapterInfo
	isSet bool
}

func (v NullableChapterInfo) Get() *ChapterInfo {
	return v.value
}

func (v *NullableChapterInfo) Set(val *ChapterInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableChapterInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableChapterInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChapterInfo(val *ChapterInfo) *NullableChapterInfo {
	return &NullableChapterInfo{value: val, isSet: true}
}

func (v NullableChapterInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChapterInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


