/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the LyricMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LyricMetadata{}

// LyricMetadata LyricMetadata model.
type LyricMetadata struct {
	// Gets or sets the song artist.
	Artist NullableString `json:"Artist,omitempty"`
	// Gets or sets the album this song is on.
	Album NullableString `json:"Album,omitempty"`
	// Gets or sets the title of the song.
	Title NullableString `json:"Title,omitempty"`
	// Gets or sets the author of the lyric data.
	Author NullableString `json:"Author,omitempty"`
	// Gets or sets the length of the song in ticks.
	Length NullableInt64 `json:"Length,omitempty"`
	// Gets or sets who the LRC file was created by.
	By NullableString `json:"By,omitempty"`
	// Gets or sets the lyric offset compared to audio in ticks.
	Offset NullableInt64 `json:"Offset,omitempty"`
	// Gets or sets the software used to create the LRC file.
	Creator NullableString `json:"Creator,omitempty"`
	// Gets or sets the version of the creator used.
	Version NullableString `json:"Version,omitempty"`
	// Gets or sets a value indicating whether this lyric is synced.
	IsSynced NullableBool `json:"IsSynced,omitempty"`
}

// NewLyricMetadata instantiates a new LyricMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLyricMetadata() *LyricMetadata {
	this := LyricMetadata{}
	return &this
}

// NewLyricMetadataWithDefaults instantiates a new LyricMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLyricMetadataWithDefaults() *LyricMetadata {
	this := LyricMetadata{}
	return &this
}

// GetArtist returns the Artist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LyricMetadata) GetArtist() string {
	if o == nil || IsNil(o.Artist.Get()) {
		var ret string
		return ret
	}
	return *o.Artist.Get()
}

// GetArtistOk returns a tuple with the Artist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LyricMetadata) GetArtistOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Artist.Get(), o.Artist.IsSet()
}

// HasArtist returns a boolean if a field has been set.
func (o *LyricMetadata) HasArtist() bool {
	if o != nil && o.Artist.IsSet() {
		return true
	}

	return false
}

// SetArtist gets a reference to the given NullableString and assigns it to the Artist field.
func (o *LyricMetadata) SetArtist(v string) {
	o.Artist.Set(&v)
}
// SetArtistNil sets the value for Artist to be an explicit nil
func (o *LyricMetadata) SetArtistNil() {
	o.Artist.Set(nil)
}

// UnsetArtist ensures that no value is present for Artist, not even an explicit nil
func (o *LyricMetadata) UnsetArtist() {
	o.Artist.Unset()
}

// GetAlbum returns the Album field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LyricMetadata) GetAlbum() string {
	if o == nil || IsNil(o.Album.Get()) {
		var ret string
		return ret
	}
	return *o.Album.Get()
}

// GetAlbumOk returns a tuple with the Album field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LyricMetadata) GetAlbumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Album.Get(), o.Album.IsSet()
}

// HasAlbum returns a boolean if a field has been set.
func (o *LyricMetadata) HasAlbum() bool {
	if o != nil && o.Album.IsSet() {
		return true
	}

	return false
}

// SetAlbum gets a reference to the given NullableString and assigns it to the Album field.
func (o *LyricMetadata) SetAlbum(v string) {
	o.Album.Set(&v)
}
// SetAlbumNil sets the value for Album to be an explicit nil
func (o *LyricMetadata) SetAlbumNil() {
	o.Album.Set(nil)
}

// UnsetAlbum ensures that no value is present for Album, not even an explicit nil
func (o *LyricMetadata) UnsetAlbum() {
	o.Album.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LyricMetadata) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LyricMetadata) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *LyricMetadata) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *LyricMetadata) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *LyricMetadata) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *LyricMetadata) UnsetTitle() {
	o.Title.Unset()
}

// GetAuthor returns the Author field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LyricMetadata) GetAuthor() string {
	if o == nil || IsNil(o.Author.Get()) {
		var ret string
		return ret
	}
	return *o.Author.Get()
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LyricMetadata) GetAuthorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Author.Get(), o.Author.IsSet()
}

// HasAuthor returns a boolean if a field has been set.
func (o *LyricMetadata) HasAuthor() bool {
	if o != nil && o.Author.IsSet() {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given NullableString and assigns it to the Author field.
func (o *LyricMetadata) SetAuthor(v string) {
	o.Author.Set(&v)
}
// SetAuthorNil sets the value for Author to be an explicit nil
func (o *LyricMetadata) SetAuthorNil() {
	o.Author.Set(nil)
}

// UnsetAuthor ensures that no value is present for Author, not even an explicit nil
func (o *LyricMetadata) UnsetAuthor() {
	o.Author.Unset()
}

// GetLength returns the Length field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LyricMetadata) GetLength() int64 {
	if o == nil || IsNil(o.Length.Get()) {
		var ret int64
		return ret
	}
	return *o.Length.Get()
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LyricMetadata) GetLengthOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Length.Get(), o.Length.IsSet()
}

// HasLength returns a boolean if a field has been set.
func (o *LyricMetadata) HasLength() bool {
	if o != nil && o.Length.IsSet() {
		return true
	}

	return false
}

// SetLength gets a reference to the given NullableInt64 and assigns it to the Length field.
func (o *LyricMetadata) SetLength(v int64) {
	o.Length.Set(&v)
}
// SetLengthNil sets the value for Length to be an explicit nil
func (o *LyricMetadata) SetLengthNil() {
	o.Length.Set(nil)
}

// UnsetLength ensures that no value is present for Length, not even an explicit nil
func (o *LyricMetadata) UnsetLength() {
	o.Length.Unset()
}

// GetBy returns the By field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LyricMetadata) GetBy() string {
	if o == nil || IsNil(o.By.Get()) {
		var ret string
		return ret
	}
	return *o.By.Get()
}

// GetByOk returns a tuple with the By field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LyricMetadata) GetByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.By.Get(), o.By.IsSet()
}

// HasBy returns a boolean if a field has been set.
func (o *LyricMetadata) HasBy() bool {
	if o != nil && o.By.IsSet() {
		return true
	}

	return false
}

// SetBy gets a reference to the given NullableString and assigns it to the By field.
func (o *LyricMetadata) SetBy(v string) {
	o.By.Set(&v)
}
// SetByNil sets the value for By to be an explicit nil
func (o *LyricMetadata) SetByNil() {
	o.By.Set(nil)
}

// UnsetBy ensures that no value is present for By, not even an explicit nil
func (o *LyricMetadata) UnsetBy() {
	o.By.Unset()
}

// GetOffset returns the Offset field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LyricMetadata) GetOffset() int64 {
	if o == nil || IsNil(o.Offset.Get()) {
		var ret int64
		return ret
	}
	return *o.Offset.Get()
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LyricMetadata) GetOffsetOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Offset.Get(), o.Offset.IsSet()
}

// HasOffset returns a boolean if a field has been set.
func (o *LyricMetadata) HasOffset() bool {
	if o != nil && o.Offset.IsSet() {
		return true
	}

	return false
}

// SetOffset gets a reference to the given NullableInt64 and assigns it to the Offset field.
func (o *LyricMetadata) SetOffset(v int64) {
	o.Offset.Set(&v)
}
// SetOffsetNil sets the value for Offset to be an explicit nil
func (o *LyricMetadata) SetOffsetNil() {
	o.Offset.Set(nil)
}

// UnsetOffset ensures that no value is present for Offset, not even an explicit nil
func (o *LyricMetadata) UnsetOffset() {
	o.Offset.Unset()
}

// GetCreator returns the Creator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LyricMetadata) GetCreator() string {
	if o == nil || IsNil(o.Creator.Get()) {
		var ret string
		return ret
	}
	return *o.Creator.Get()
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LyricMetadata) GetCreatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Creator.Get(), o.Creator.IsSet()
}

// HasCreator returns a boolean if a field has been set.
func (o *LyricMetadata) HasCreator() bool {
	if o != nil && o.Creator.IsSet() {
		return true
	}

	return false
}

// SetCreator gets a reference to the given NullableString and assigns it to the Creator field.
func (o *LyricMetadata) SetCreator(v string) {
	o.Creator.Set(&v)
}
// SetCreatorNil sets the value for Creator to be an explicit nil
func (o *LyricMetadata) SetCreatorNil() {
	o.Creator.Set(nil)
}

// UnsetCreator ensures that no value is present for Creator, not even an explicit nil
func (o *LyricMetadata) UnsetCreator() {
	o.Creator.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LyricMetadata) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LyricMetadata) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *LyricMetadata) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *LyricMetadata) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *LyricMetadata) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *LyricMetadata) UnsetVersion() {
	o.Version.Unset()
}

// GetIsSynced returns the IsSynced field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LyricMetadata) GetIsSynced() bool {
	if o == nil || IsNil(o.IsSynced.Get()) {
		var ret bool
		return ret
	}
	return *o.IsSynced.Get()
}

// GetIsSyncedOk returns a tuple with the IsSynced field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LyricMetadata) GetIsSyncedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsSynced.Get(), o.IsSynced.IsSet()
}

// HasIsSynced returns a boolean if a field has been set.
func (o *LyricMetadata) HasIsSynced() bool {
	if o != nil && o.IsSynced.IsSet() {
		return true
	}

	return false
}

// SetIsSynced gets a reference to the given NullableBool and assigns it to the IsSynced field.
func (o *LyricMetadata) SetIsSynced(v bool) {
	o.IsSynced.Set(&v)
}
// SetIsSyncedNil sets the value for IsSynced to be an explicit nil
func (o *LyricMetadata) SetIsSyncedNil() {
	o.IsSynced.Set(nil)
}

// UnsetIsSynced ensures that no value is present for IsSynced, not even an explicit nil
func (o *LyricMetadata) UnsetIsSynced() {
	o.IsSynced.Unset()
}

func (o LyricMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LyricMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Artist.IsSet() {
		toSerialize["Artist"] = o.Artist.Get()
	}
	if o.Album.IsSet() {
		toSerialize["Album"] = o.Album.Get()
	}
	if o.Title.IsSet() {
		toSerialize["Title"] = o.Title.Get()
	}
	if o.Author.IsSet() {
		toSerialize["Author"] = o.Author.Get()
	}
	if o.Length.IsSet() {
		toSerialize["Length"] = o.Length.Get()
	}
	if o.By.IsSet() {
		toSerialize["By"] = o.By.Get()
	}
	if o.Offset.IsSet() {
		toSerialize["Offset"] = o.Offset.Get()
	}
	if o.Creator.IsSet() {
		toSerialize["Creator"] = o.Creator.Get()
	}
	if o.Version.IsSet() {
		toSerialize["Version"] = o.Version.Get()
	}
	if o.IsSynced.IsSet() {
		toSerialize["IsSynced"] = o.IsSynced.Get()
	}
	return toSerialize, nil
}

type NullableLyricMetadata struct {
	value *LyricMetadata
	isSet bool
}

func (v NullableLyricMetadata) Get() *LyricMetadata {
	return v.value
}

func (v *NullableLyricMetadata) Set(val *LyricMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableLyricMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableLyricMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLyricMetadata(val *LyricMetadata) *NullableLyricMetadata {
	return &NullableLyricMetadata{value: val, isSet: true}
}

func (v NullableLyricMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLyricMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


