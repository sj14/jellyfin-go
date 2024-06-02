/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the LibraryChangedMessageData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LibraryChangedMessageData{}

// LibraryChangedMessageData Gets or sets the data.
type LibraryChangedMessageData struct {
	// Gets or sets the folders added to.
	FoldersAddedTo []string `json:"FoldersAddedTo,omitempty"`
	// Gets or sets the folders removed from.
	FoldersRemovedFrom []string `json:"FoldersRemovedFrom,omitempty"`
	// Gets or sets the items added.
	ItemsAdded []string `json:"ItemsAdded,omitempty"`
	// Gets or sets the items removed.
	ItemsRemoved []string `json:"ItemsRemoved,omitempty"`
	// Gets or sets the items updated.
	ItemsUpdated []string `json:"ItemsUpdated,omitempty"`
	CollectionFolders []string `json:"CollectionFolders,omitempty"`
	IsEmpty *bool `json:"IsEmpty,omitempty"`
}

// NewLibraryChangedMessageData instantiates a new LibraryChangedMessageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLibraryChangedMessageData() *LibraryChangedMessageData {
	this := LibraryChangedMessageData{}
	return &this
}

// NewLibraryChangedMessageDataWithDefaults instantiates a new LibraryChangedMessageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLibraryChangedMessageDataWithDefaults() *LibraryChangedMessageData {
	this := LibraryChangedMessageData{}
	return &this
}

// GetFoldersAddedTo returns the FoldersAddedTo field value if set, zero value otherwise.
func (o *LibraryChangedMessageData) GetFoldersAddedTo() []string {
	if o == nil || IsNil(o.FoldersAddedTo) {
		var ret []string
		return ret
	}
	return o.FoldersAddedTo
}

// GetFoldersAddedToOk returns a tuple with the FoldersAddedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryChangedMessageData) GetFoldersAddedToOk() ([]string, bool) {
	if o == nil || IsNil(o.FoldersAddedTo) {
		return nil, false
	}
	return o.FoldersAddedTo, true
}

// HasFoldersAddedTo returns a boolean if a field has been set.
func (o *LibraryChangedMessageData) HasFoldersAddedTo() bool {
	if o != nil && !IsNil(o.FoldersAddedTo) {
		return true
	}

	return false
}

// SetFoldersAddedTo gets a reference to the given []string and assigns it to the FoldersAddedTo field.
func (o *LibraryChangedMessageData) SetFoldersAddedTo(v []string) {
	o.FoldersAddedTo = v
}

// GetFoldersRemovedFrom returns the FoldersRemovedFrom field value if set, zero value otherwise.
func (o *LibraryChangedMessageData) GetFoldersRemovedFrom() []string {
	if o == nil || IsNil(o.FoldersRemovedFrom) {
		var ret []string
		return ret
	}
	return o.FoldersRemovedFrom
}

// GetFoldersRemovedFromOk returns a tuple with the FoldersRemovedFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryChangedMessageData) GetFoldersRemovedFromOk() ([]string, bool) {
	if o == nil || IsNil(o.FoldersRemovedFrom) {
		return nil, false
	}
	return o.FoldersRemovedFrom, true
}

// HasFoldersRemovedFrom returns a boolean if a field has been set.
func (o *LibraryChangedMessageData) HasFoldersRemovedFrom() bool {
	if o != nil && !IsNil(o.FoldersRemovedFrom) {
		return true
	}

	return false
}

// SetFoldersRemovedFrom gets a reference to the given []string and assigns it to the FoldersRemovedFrom field.
func (o *LibraryChangedMessageData) SetFoldersRemovedFrom(v []string) {
	o.FoldersRemovedFrom = v
}

// GetItemsAdded returns the ItemsAdded field value if set, zero value otherwise.
func (o *LibraryChangedMessageData) GetItemsAdded() []string {
	if o == nil || IsNil(o.ItemsAdded) {
		var ret []string
		return ret
	}
	return o.ItemsAdded
}

// GetItemsAddedOk returns a tuple with the ItemsAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryChangedMessageData) GetItemsAddedOk() ([]string, bool) {
	if o == nil || IsNil(o.ItemsAdded) {
		return nil, false
	}
	return o.ItemsAdded, true
}

// HasItemsAdded returns a boolean if a field has been set.
func (o *LibraryChangedMessageData) HasItemsAdded() bool {
	if o != nil && !IsNil(o.ItemsAdded) {
		return true
	}

	return false
}

// SetItemsAdded gets a reference to the given []string and assigns it to the ItemsAdded field.
func (o *LibraryChangedMessageData) SetItemsAdded(v []string) {
	o.ItemsAdded = v
}

// GetItemsRemoved returns the ItemsRemoved field value if set, zero value otherwise.
func (o *LibraryChangedMessageData) GetItemsRemoved() []string {
	if o == nil || IsNil(o.ItemsRemoved) {
		var ret []string
		return ret
	}
	return o.ItemsRemoved
}

// GetItemsRemovedOk returns a tuple with the ItemsRemoved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryChangedMessageData) GetItemsRemovedOk() ([]string, bool) {
	if o == nil || IsNil(o.ItemsRemoved) {
		return nil, false
	}
	return o.ItemsRemoved, true
}

// HasItemsRemoved returns a boolean if a field has been set.
func (o *LibraryChangedMessageData) HasItemsRemoved() bool {
	if o != nil && !IsNil(o.ItemsRemoved) {
		return true
	}

	return false
}

// SetItemsRemoved gets a reference to the given []string and assigns it to the ItemsRemoved field.
func (o *LibraryChangedMessageData) SetItemsRemoved(v []string) {
	o.ItemsRemoved = v
}

// GetItemsUpdated returns the ItemsUpdated field value if set, zero value otherwise.
func (o *LibraryChangedMessageData) GetItemsUpdated() []string {
	if o == nil || IsNil(o.ItemsUpdated) {
		var ret []string
		return ret
	}
	return o.ItemsUpdated
}

// GetItemsUpdatedOk returns a tuple with the ItemsUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryChangedMessageData) GetItemsUpdatedOk() ([]string, bool) {
	if o == nil || IsNil(o.ItemsUpdated) {
		return nil, false
	}
	return o.ItemsUpdated, true
}

// HasItemsUpdated returns a boolean if a field has been set.
func (o *LibraryChangedMessageData) HasItemsUpdated() bool {
	if o != nil && !IsNil(o.ItemsUpdated) {
		return true
	}

	return false
}

// SetItemsUpdated gets a reference to the given []string and assigns it to the ItemsUpdated field.
func (o *LibraryChangedMessageData) SetItemsUpdated(v []string) {
	o.ItemsUpdated = v
}

// GetCollectionFolders returns the CollectionFolders field value if set, zero value otherwise.
func (o *LibraryChangedMessageData) GetCollectionFolders() []string {
	if o == nil || IsNil(o.CollectionFolders) {
		var ret []string
		return ret
	}
	return o.CollectionFolders
}

// GetCollectionFoldersOk returns a tuple with the CollectionFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryChangedMessageData) GetCollectionFoldersOk() ([]string, bool) {
	if o == nil || IsNil(o.CollectionFolders) {
		return nil, false
	}
	return o.CollectionFolders, true
}

// HasCollectionFolders returns a boolean if a field has been set.
func (o *LibraryChangedMessageData) HasCollectionFolders() bool {
	if o != nil && !IsNil(o.CollectionFolders) {
		return true
	}

	return false
}

// SetCollectionFolders gets a reference to the given []string and assigns it to the CollectionFolders field.
func (o *LibraryChangedMessageData) SetCollectionFolders(v []string) {
	o.CollectionFolders = v
}

// GetIsEmpty returns the IsEmpty field value if set, zero value otherwise.
func (o *LibraryChangedMessageData) GetIsEmpty() bool {
	if o == nil || IsNil(o.IsEmpty) {
		var ret bool
		return ret
	}
	return *o.IsEmpty
}

// GetIsEmptyOk returns a tuple with the IsEmpty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryChangedMessageData) GetIsEmptyOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEmpty) {
		return nil, false
	}
	return o.IsEmpty, true
}

// HasIsEmpty returns a boolean if a field has been set.
func (o *LibraryChangedMessageData) HasIsEmpty() bool {
	if o != nil && !IsNil(o.IsEmpty) {
		return true
	}

	return false
}

// SetIsEmpty gets a reference to the given bool and assigns it to the IsEmpty field.
func (o *LibraryChangedMessageData) SetIsEmpty(v bool) {
	o.IsEmpty = &v
}

func (o LibraryChangedMessageData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LibraryChangedMessageData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FoldersAddedTo) {
		toSerialize["FoldersAddedTo"] = o.FoldersAddedTo
	}
	if !IsNil(o.FoldersRemovedFrom) {
		toSerialize["FoldersRemovedFrom"] = o.FoldersRemovedFrom
	}
	if !IsNil(o.ItemsAdded) {
		toSerialize["ItemsAdded"] = o.ItemsAdded
	}
	if !IsNil(o.ItemsRemoved) {
		toSerialize["ItemsRemoved"] = o.ItemsRemoved
	}
	if !IsNil(o.ItemsUpdated) {
		toSerialize["ItemsUpdated"] = o.ItemsUpdated
	}
	if !IsNil(o.CollectionFolders) {
		toSerialize["CollectionFolders"] = o.CollectionFolders
	}
	if !IsNil(o.IsEmpty) {
		toSerialize["IsEmpty"] = o.IsEmpty
	}
	return toSerialize, nil
}

type NullableLibraryChangedMessageData struct {
	value *LibraryChangedMessageData
	isSet bool
}

func (v NullableLibraryChangedMessageData) Get() *LibraryChangedMessageData {
	return v.value
}

func (v *NullableLibraryChangedMessageData) Set(val *LibraryChangedMessageData) {
	v.value = val
	v.isSet = true
}

func (v NullableLibraryChangedMessageData) IsSet() bool {
	return v.isSet
}

func (v *NullableLibraryChangedMessageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLibraryChangedMessageData(val *LibraryChangedMessageData) *NullableLibraryChangedMessageData {
	return &NullableLibraryChangedMessageData{value: val, isSet: true}
}

func (v NullableLibraryChangedMessageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLibraryChangedMessageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


