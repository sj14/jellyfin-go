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

// checks if the ItemCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemCounts{}

// ItemCounts Class LibrarySummary.
type ItemCounts struct {
	// Gets or sets the movie count.
	MovieCount *int32 `json:"MovieCount,omitempty"`
	// Gets or sets the series count.
	SeriesCount *int32 `json:"SeriesCount,omitempty"`
	// Gets or sets the episode count.
	EpisodeCount *int32 `json:"EpisodeCount,omitempty"`
	// Gets or sets the artist count.
	ArtistCount *int32 `json:"ArtistCount,omitempty"`
	// Gets or sets the program count.
	ProgramCount *int32 `json:"ProgramCount,omitempty"`
	// Gets or sets the trailer count.
	TrailerCount *int32 `json:"TrailerCount,omitempty"`
	// Gets or sets the song count.
	SongCount *int32 `json:"SongCount,omitempty"`
	// Gets or sets the album count.
	AlbumCount *int32 `json:"AlbumCount,omitempty"`
	// Gets or sets the music video count.
	MusicVideoCount *int32 `json:"MusicVideoCount,omitempty"`
	// Gets or sets the box set count.
	BoxSetCount *int32 `json:"BoxSetCount,omitempty"`
	// Gets or sets the book count.
	BookCount *int32 `json:"BookCount,omitempty"`
	// Gets or sets the item count.
	ItemCount *int32 `json:"ItemCount,omitempty"`
}

// NewItemCounts instantiates a new ItemCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemCounts() *ItemCounts {
	this := ItemCounts{}
	return &this
}

// NewItemCountsWithDefaults instantiates a new ItemCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemCountsWithDefaults() *ItemCounts {
	this := ItemCounts{}
	return &this
}

// GetMovieCount returns the MovieCount field value if set, zero value otherwise.
func (o *ItemCounts) GetMovieCount() int32 {
	if o == nil || IsNil(o.MovieCount) {
		var ret int32
		return ret
	}
	return *o.MovieCount
}

// GetMovieCountOk returns a tuple with the MovieCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemCounts) GetMovieCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MovieCount) {
		return nil, false
	}
	return o.MovieCount, true
}

// HasMovieCount returns a boolean if a field has been set.
func (o *ItemCounts) HasMovieCount() bool {
	if o != nil && !IsNil(o.MovieCount) {
		return true
	}

	return false
}

// SetMovieCount gets a reference to the given int32 and assigns it to the MovieCount field.
func (o *ItemCounts) SetMovieCount(v int32) {
	o.MovieCount = &v
}

// GetSeriesCount returns the SeriesCount field value if set, zero value otherwise.
func (o *ItemCounts) GetSeriesCount() int32 {
	if o == nil || IsNil(o.SeriesCount) {
		var ret int32
		return ret
	}
	return *o.SeriesCount
}

// GetSeriesCountOk returns a tuple with the SeriesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemCounts) GetSeriesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SeriesCount) {
		return nil, false
	}
	return o.SeriesCount, true
}

// HasSeriesCount returns a boolean if a field has been set.
func (o *ItemCounts) HasSeriesCount() bool {
	if o != nil && !IsNil(o.SeriesCount) {
		return true
	}

	return false
}

// SetSeriesCount gets a reference to the given int32 and assigns it to the SeriesCount field.
func (o *ItemCounts) SetSeriesCount(v int32) {
	o.SeriesCount = &v
}

// GetEpisodeCount returns the EpisodeCount field value if set, zero value otherwise.
func (o *ItemCounts) GetEpisodeCount() int32 {
	if o == nil || IsNil(o.EpisodeCount) {
		var ret int32
		return ret
	}
	return *o.EpisodeCount
}

// GetEpisodeCountOk returns a tuple with the EpisodeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemCounts) GetEpisodeCountOk() (*int32, bool) {
	if o == nil || IsNil(o.EpisodeCount) {
		return nil, false
	}
	return o.EpisodeCount, true
}

// HasEpisodeCount returns a boolean if a field has been set.
func (o *ItemCounts) HasEpisodeCount() bool {
	if o != nil && !IsNil(o.EpisodeCount) {
		return true
	}

	return false
}

// SetEpisodeCount gets a reference to the given int32 and assigns it to the EpisodeCount field.
func (o *ItemCounts) SetEpisodeCount(v int32) {
	o.EpisodeCount = &v
}

// GetArtistCount returns the ArtistCount field value if set, zero value otherwise.
func (o *ItemCounts) GetArtistCount() int32 {
	if o == nil || IsNil(o.ArtistCount) {
		var ret int32
		return ret
	}
	return *o.ArtistCount
}

// GetArtistCountOk returns a tuple with the ArtistCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemCounts) GetArtistCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ArtistCount) {
		return nil, false
	}
	return o.ArtistCount, true
}

// HasArtistCount returns a boolean if a field has been set.
func (o *ItemCounts) HasArtistCount() bool {
	if o != nil && !IsNil(o.ArtistCount) {
		return true
	}

	return false
}

// SetArtistCount gets a reference to the given int32 and assigns it to the ArtistCount field.
func (o *ItemCounts) SetArtistCount(v int32) {
	o.ArtistCount = &v
}

// GetProgramCount returns the ProgramCount field value if set, zero value otherwise.
func (o *ItemCounts) GetProgramCount() int32 {
	if o == nil || IsNil(o.ProgramCount) {
		var ret int32
		return ret
	}
	return *o.ProgramCount
}

// GetProgramCountOk returns a tuple with the ProgramCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemCounts) GetProgramCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ProgramCount) {
		return nil, false
	}
	return o.ProgramCount, true
}

// HasProgramCount returns a boolean if a field has been set.
func (o *ItemCounts) HasProgramCount() bool {
	if o != nil && !IsNil(o.ProgramCount) {
		return true
	}

	return false
}

// SetProgramCount gets a reference to the given int32 and assigns it to the ProgramCount field.
func (o *ItemCounts) SetProgramCount(v int32) {
	o.ProgramCount = &v
}

// GetTrailerCount returns the TrailerCount field value if set, zero value otherwise.
func (o *ItemCounts) GetTrailerCount() int32 {
	if o == nil || IsNil(o.TrailerCount) {
		var ret int32
		return ret
	}
	return *o.TrailerCount
}

// GetTrailerCountOk returns a tuple with the TrailerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemCounts) GetTrailerCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TrailerCount) {
		return nil, false
	}
	return o.TrailerCount, true
}

// HasTrailerCount returns a boolean if a field has been set.
func (o *ItemCounts) HasTrailerCount() bool {
	if o != nil && !IsNil(o.TrailerCount) {
		return true
	}

	return false
}

// SetTrailerCount gets a reference to the given int32 and assigns it to the TrailerCount field.
func (o *ItemCounts) SetTrailerCount(v int32) {
	o.TrailerCount = &v
}

// GetSongCount returns the SongCount field value if set, zero value otherwise.
func (o *ItemCounts) GetSongCount() int32 {
	if o == nil || IsNil(o.SongCount) {
		var ret int32
		return ret
	}
	return *o.SongCount
}

// GetSongCountOk returns a tuple with the SongCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemCounts) GetSongCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SongCount) {
		return nil, false
	}
	return o.SongCount, true
}

// HasSongCount returns a boolean if a field has been set.
func (o *ItemCounts) HasSongCount() bool {
	if o != nil && !IsNil(o.SongCount) {
		return true
	}

	return false
}

// SetSongCount gets a reference to the given int32 and assigns it to the SongCount field.
func (o *ItemCounts) SetSongCount(v int32) {
	o.SongCount = &v
}

// GetAlbumCount returns the AlbumCount field value if set, zero value otherwise.
func (o *ItemCounts) GetAlbumCount() int32 {
	if o == nil || IsNil(o.AlbumCount) {
		var ret int32
		return ret
	}
	return *o.AlbumCount
}

// GetAlbumCountOk returns a tuple with the AlbumCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemCounts) GetAlbumCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AlbumCount) {
		return nil, false
	}
	return o.AlbumCount, true
}

// HasAlbumCount returns a boolean if a field has been set.
func (o *ItemCounts) HasAlbumCount() bool {
	if o != nil && !IsNil(o.AlbumCount) {
		return true
	}

	return false
}

// SetAlbumCount gets a reference to the given int32 and assigns it to the AlbumCount field.
func (o *ItemCounts) SetAlbumCount(v int32) {
	o.AlbumCount = &v
}

// GetMusicVideoCount returns the MusicVideoCount field value if set, zero value otherwise.
func (o *ItemCounts) GetMusicVideoCount() int32 {
	if o == nil || IsNil(o.MusicVideoCount) {
		var ret int32
		return ret
	}
	return *o.MusicVideoCount
}

// GetMusicVideoCountOk returns a tuple with the MusicVideoCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemCounts) GetMusicVideoCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MusicVideoCount) {
		return nil, false
	}
	return o.MusicVideoCount, true
}

// HasMusicVideoCount returns a boolean if a field has been set.
func (o *ItemCounts) HasMusicVideoCount() bool {
	if o != nil && !IsNil(o.MusicVideoCount) {
		return true
	}

	return false
}

// SetMusicVideoCount gets a reference to the given int32 and assigns it to the MusicVideoCount field.
func (o *ItemCounts) SetMusicVideoCount(v int32) {
	o.MusicVideoCount = &v
}

// GetBoxSetCount returns the BoxSetCount field value if set, zero value otherwise.
func (o *ItemCounts) GetBoxSetCount() int32 {
	if o == nil || IsNil(o.BoxSetCount) {
		var ret int32
		return ret
	}
	return *o.BoxSetCount
}

// GetBoxSetCountOk returns a tuple with the BoxSetCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemCounts) GetBoxSetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.BoxSetCount) {
		return nil, false
	}
	return o.BoxSetCount, true
}

// HasBoxSetCount returns a boolean if a field has been set.
func (o *ItemCounts) HasBoxSetCount() bool {
	if o != nil && !IsNil(o.BoxSetCount) {
		return true
	}

	return false
}

// SetBoxSetCount gets a reference to the given int32 and assigns it to the BoxSetCount field.
func (o *ItemCounts) SetBoxSetCount(v int32) {
	o.BoxSetCount = &v
}

// GetBookCount returns the BookCount field value if set, zero value otherwise.
func (o *ItemCounts) GetBookCount() int32 {
	if o == nil || IsNil(o.BookCount) {
		var ret int32
		return ret
	}
	return *o.BookCount
}

// GetBookCountOk returns a tuple with the BookCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemCounts) GetBookCountOk() (*int32, bool) {
	if o == nil || IsNil(o.BookCount) {
		return nil, false
	}
	return o.BookCount, true
}

// HasBookCount returns a boolean if a field has been set.
func (o *ItemCounts) HasBookCount() bool {
	if o != nil && !IsNil(o.BookCount) {
		return true
	}

	return false
}

// SetBookCount gets a reference to the given int32 and assigns it to the BookCount field.
func (o *ItemCounts) SetBookCount(v int32) {
	o.BookCount = &v
}

// GetItemCount returns the ItemCount field value if set, zero value otherwise.
func (o *ItemCounts) GetItemCount() int32 {
	if o == nil || IsNil(o.ItemCount) {
		var ret int32
		return ret
	}
	return *o.ItemCount
}

// GetItemCountOk returns a tuple with the ItemCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemCounts) GetItemCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ItemCount) {
		return nil, false
	}
	return o.ItemCount, true
}

// HasItemCount returns a boolean if a field has been set.
func (o *ItemCounts) HasItemCount() bool {
	if o != nil && !IsNil(o.ItemCount) {
		return true
	}

	return false
}

// SetItemCount gets a reference to the given int32 and assigns it to the ItemCount field.
func (o *ItemCounts) SetItemCount(v int32) {
	o.ItemCount = &v
}

func (o ItemCounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MovieCount) {
		toSerialize["MovieCount"] = o.MovieCount
	}
	if !IsNil(o.SeriesCount) {
		toSerialize["SeriesCount"] = o.SeriesCount
	}
	if !IsNil(o.EpisodeCount) {
		toSerialize["EpisodeCount"] = o.EpisodeCount
	}
	if !IsNil(o.ArtistCount) {
		toSerialize["ArtistCount"] = o.ArtistCount
	}
	if !IsNil(o.ProgramCount) {
		toSerialize["ProgramCount"] = o.ProgramCount
	}
	if !IsNil(o.TrailerCount) {
		toSerialize["TrailerCount"] = o.TrailerCount
	}
	if !IsNil(o.SongCount) {
		toSerialize["SongCount"] = o.SongCount
	}
	if !IsNil(o.AlbumCount) {
		toSerialize["AlbumCount"] = o.AlbumCount
	}
	if !IsNil(o.MusicVideoCount) {
		toSerialize["MusicVideoCount"] = o.MusicVideoCount
	}
	if !IsNil(o.BoxSetCount) {
		toSerialize["BoxSetCount"] = o.BoxSetCount
	}
	if !IsNil(o.BookCount) {
		toSerialize["BookCount"] = o.BookCount
	}
	if !IsNil(o.ItemCount) {
		toSerialize["ItemCount"] = o.ItemCount
	}
	return toSerialize, nil
}

type NullableItemCounts struct {
	value *ItemCounts
	isSet bool
}

func (v NullableItemCounts) Get() *ItemCounts {
	return v.value
}

func (v *NullableItemCounts) Set(val *ItemCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableItemCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableItemCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemCounts(val *ItemCounts) *NullableItemCounts {
	return &NullableItemCounts{value: val, isSet: true}
}

func (v NullableItemCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


