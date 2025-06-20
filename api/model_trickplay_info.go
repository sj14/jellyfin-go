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

// checks if the TrickplayInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrickplayInfo{}

// TrickplayInfo An entity representing the metadata for a group of trickplay tiles.
type TrickplayInfo struct {
	// Gets or sets the id of the associated item.
	ItemId *string `json:"ItemId,omitempty"`
	// Gets or sets width of an individual thumbnail.
	Width *int32 `json:"Width,omitempty"`
	// Gets or sets height of an individual thumbnail.
	Height *int32 `json:"Height,omitempty"`
	// Gets or sets amount of thumbnails per row.
	TileWidth *int32 `json:"TileWidth,omitempty"`
	// Gets or sets amount of thumbnails per column.
	TileHeight *int32 `json:"TileHeight,omitempty"`
	// Gets or sets total amount of non-black thumbnails.
	ThumbnailCount *int32 `json:"ThumbnailCount,omitempty"`
	// Gets or sets interval in milliseconds between each trickplay thumbnail.
	Interval *int32 `json:"Interval,omitempty"`
	// Gets or sets peak bandwidth usage in bits per second.
	Bandwidth *int32 `json:"Bandwidth,omitempty"`
}

// NewTrickplayInfo instantiates a new TrickplayInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrickplayInfo() *TrickplayInfo {
	this := TrickplayInfo{}
	return &this
}

// NewTrickplayInfoWithDefaults instantiates a new TrickplayInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrickplayInfoWithDefaults() *TrickplayInfo {
	this := TrickplayInfo{}
	return &this
}

// GetItemId returns the ItemId field value if set, zero value otherwise.
func (o *TrickplayInfo) GetItemId() string {
	if o == nil || IsNil(o.ItemId) {
		var ret string
		return ret
	}
	return *o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrickplayInfo) GetItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.ItemId) {
		return nil, false
	}
	return o.ItemId, true
}

// HasItemId returns a boolean if a field has been set.
func (o *TrickplayInfo) HasItemId() bool {
	if o != nil && !IsNil(o.ItemId) {
		return true
	}

	return false
}

// SetItemId gets a reference to the given string and assigns it to the ItemId field.
func (o *TrickplayInfo) SetItemId(v string) {
	o.ItemId = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *TrickplayInfo) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrickplayInfo) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *TrickplayInfo) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *TrickplayInfo) SetWidth(v int32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *TrickplayInfo) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrickplayInfo) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *TrickplayInfo) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *TrickplayInfo) SetHeight(v int32) {
	o.Height = &v
}

// GetTileWidth returns the TileWidth field value if set, zero value otherwise.
func (o *TrickplayInfo) GetTileWidth() int32 {
	if o == nil || IsNil(o.TileWidth) {
		var ret int32
		return ret
	}
	return *o.TileWidth
}

// GetTileWidthOk returns a tuple with the TileWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrickplayInfo) GetTileWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.TileWidth) {
		return nil, false
	}
	return o.TileWidth, true
}

// HasTileWidth returns a boolean if a field has been set.
func (o *TrickplayInfo) HasTileWidth() bool {
	if o != nil && !IsNil(o.TileWidth) {
		return true
	}

	return false
}

// SetTileWidth gets a reference to the given int32 and assigns it to the TileWidth field.
func (o *TrickplayInfo) SetTileWidth(v int32) {
	o.TileWidth = &v
}

// GetTileHeight returns the TileHeight field value if set, zero value otherwise.
func (o *TrickplayInfo) GetTileHeight() int32 {
	if o == nil || IsNil(o.TileHeight) {
		var ret int32
		return ret
	}
	return *o.TileHeight
}

// GetTileHeightOk returns a tuple with the TileHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrickplayInfo) GetTileHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.TileHeight) {
		return nil, false
	}
	return o.TileHeight, true
}

// HasTileHeight returns a boolean if a field has been set.
func (o *TrickplayInfo) HasTileHeight() bool {
	if o != nil && !IsNil(o.TileHeight) {
		return true
	}

	return false
}

// SetTileHeight gets a reference to the given int32 and assigns it to the TileHeight field.
func (o *TrickplayInfo) SetTileHeight(v int32) {
	o.TileHeight = &v
}

// GetThumbnailCount returns the ThumbnailCount field value if set, zero value otherwise.
func (o *TrickplayInfo) GetThumbnailCount() int32 {
	if o == nil || IsNil(o.ThumbnailCount) {
		var ret int32
		return ret
	}
	return *o.ThumbnailCount
}

// GetThumbnailCountOk returns a tuple with the ThumbnailCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrickplayInfo) GetThumbnailCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ThumbnailCount) {
		return nil, false
	}
	return o.ThumbnailCount, true
}

// HasThumbnailCount returns a boolean if a field has been set.
func (o *TrickplayInfo) HasThumbnailCount() bool {
	if o != nil && !IsNil(o.ThumbnailCount) {
		return true
	}

	return false
}

// SetThumbnailCount gets a reference to the given int32 and assigns it to the ThumbnailCount field.
func (o *TrickplayInfo) SetThumbnailCount(v int32) {
	o.ThumbnailCount = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *TrickplayInfo) GetInterval() int32 {
	if o == nil || IsNil(o.Interval) {
		var ret int32
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrickplayInfo) GetIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *TrickplayInfo) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int32 and assigns it to the Interval field.
func (o *TrickplayInfo) SetInterval(v int32) {
	o.Interval = &v
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *TrickplayInfo) GetBandwidth() int32 {
	if o == nil || IsNil(o.Bandwidth) {
		var ret int32
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrickplayInfo) GetBandwidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Bandwidth) {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *TrickplayInfo) HasBandwidth() bool {
	if o != nil && !IsNil(o.Bandwidth) {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given int32 and assigns it to the Bandwidth field.
func (o *TrickplayInfo) SetBandwidth(v int32) {
	o.Bandwidth = &v
}

func (o TrickplayInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrickplayInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ItemId) {
		toSerialize["ItemId"] = o.ItemId
	}
	if !IsNil(o.Width) {
		toSerialize["Width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["Height"] = o.Height
	}
	if !IsNil(o.TileWidth) {
		toSerialize["TileWidth"] = o.TileWidth
	}
	if !IsNil(o.TileHeight) {
		toSerialize["TileHeight"] = o.TileHeight
	}
	if !IsNil(o.ThumbnailCount) {
		toSerialize["ThumbnailCount"] = o.ThumbnailCount
	}
	if !IsNil(o.Interval) {
		toSerialize["Interval"] = o.Interval
	}
	if !IsNil(o.Bandwidth) {
		toSerialize["Bandwidth"] = o.Bandwidth
	}
	return toSerialize, nil
}

type NullableTrickplayInfo struct {
	value *TrickplayInfo
	isSet bool
}

func (v NullableTrickplayInfo) Get() *TrickplayInfo {
	return v.value
}

func (v *NullableTrickplayInfo) Set(val *TrickplayInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTrickplayInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTrickplayInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrickplayInfo(val *TrickplayInfo) *NullableTrickplayInfo {
	return &NullableTrickplayInfo{value: val, isSet: true}
}

func (v NullableTrickplayInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrickplayInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


