/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.8.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the SubtitleOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubtitleOptions{}

// SubtitleOptions struct for SubtitleOptions
type SubtitleOptions struct {
	SkipIfEmbeddedSubtitlesPresent *bool `json:"SkipIfEmbeddedSubtitlesPresent,omitempty"`
	SkipIfAudioTrackMatches *bool `json:"SkipIfAudioTrackMatches,omitempty"`
	DownloadLanguages []string `json:"DownloadLanguages,omitempty"`
	DownloadMovieSubtitles *bool `json:"DownloadMovieSubtitles,omitempty"`
	DownloadEpisodeSubtitles *bool `json:"DownloadEpisodeSubtitles,omitempty"`
	OpenSubtitlesUsername NullableString `json:"OpenSubtitlesUsername,omitempty"`
	OpenSubtitlesPasswordHash NullableString `json:"OpenSubtitlesPasswordHash,omitempty"`
	IsOpenSubtitleVipAccount *bool `json:"IsOpenSubtitleVipAccount,omitempty"`
	RequirePerfectMatch *bool `json:"RequirePerfectMatch,omitempty"`
}

// NewSubtitleOptions instantiates a new SubtitleOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubtitleOptions() *SubtitleOptions {
	this := SubtitleOptions{}
	return &this
}

// NewSubtitleOptionsWithDefaults instantiates a new SubtitleOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubtitleOptionsWithDefaults() *SubtitleOptions {
	this := SubtitleOptions{}
	return &this
}

// GetSkipIfEmbeddedSubtitlesPresent returns the SkipIfEmbeddedSubtitlesPresent field value if set, zero value otherwise.
func (o *SubtitleOptions) GetSkipIfEmbeddedSubtitlesPresent() bool {
	if o == nil || IsNil(o.SkipIfEmbeddedSubtitlesPresent) {
		var ret bool
		return ret
	}
	return *o.SkipIfEmbeddedSubtitlesPresent
}

// GetSkipIfEmbeddedSubtitlesPresentOk returns a tuple with the SkipIfEmbeddedSubtitlesPresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubtitleOptions) GetSkipIfEmbeddedSubtitlesPresentOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipIfEmbeddedSubtitlesPresent) {
		return nil, false
	}
	return o.SkipIfEmbeddedSubtitlesPresent, true
}

// HasSkipIfEmbeddedSubtitlesPresent returns a boolean if a field has been set.
func (o *SubtitleOptions) HasSkipIfEmbeddedSubtitlesPresent() bool {
	if o != nil && !IsNil(o.SkipIfEmbeddedSubtitlesPresent) {
		return true
	}

	return false
}

// SetSkipIfEmbeddedSubtitlesPresent gets a reference to the given bool and assigns it to the SkipIfEmbeddedSubtitlesPresent field.
func (o *SubtitleOptions) SetSkipIfEmbeddedSubtitlesPresent(v bool) {
	o.SkipIfEmbeddedSubtitlesPresent = &v
}

// GetSkipIfAudioTrackMatches returns the SkipIfAudioTrackMatches field value if set, zero value otherwise.
func (o *SubtitleOptions) GetSkipIfAudioTrackMatches() bool {
	if o == nil || IsNil(o.SkipIfAudioTrackMatches) {
		var ret bool
		return ret
	}
	return *o.SkipIfAudioTrackMatches
}

// GetSkipIfAudioTrackMatchesOk returns a tuple with the SkipIfAudioTrackMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubtitleOptions) GetSkipIfAudioTrackMatchesOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipIfAudioTrackMatches) {
		return nil, false
	}
	return o.SkipIfAudioTrackMatches, true
}

// HasSkipIfAudioTrackMatches returns a boolean if a field has been set.
func (o *SubtitleOptions) HasSkipIfAudioTrackMatches() bool {
	if o != nil && !IsNil(o.SkipIfAudioTrackMatches) {
		return true
	}

	return false
}

// SetSkipIfAudioTrackMatches gets a reference to the given bool and assigns it to the SkipIfAudioTrackMatches field.
func (o *SubtitleOptions) SetSkipIfAudioTrackMatches(v bool) {
	o.SkipIfAudioTrackMatches = &v
}

// GetDownloadLanguages returns the DownloadLanguages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubtitleOptions) GetDownloadLanguages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DownloadLanguages
}

// GetDownloadLanguagesOk returns a tuple with the DownloadLanguages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubtitleOptions) GetDownloadLanguagesOk() ([]string, bool) {
	if o == nil || IsNil(o.DownloadLanguages) {
		return nil, false
	}
	return o.DownloadLanguages, true
}

// HasDownloadLanguages returns a boolean if a field has been set.
func (o *SubtitleOptions) HasDownloadLanguages() bool {
	if o != nil && !IsNil(o.DownloadLanguages) {
		return true
	}

	return false
}

// SetDownloadLanguages gets a reference to the given []string and assigns it to the DownloadLanguages field.
func (o *SubtitleOptions) SetDownloadLanguages(v []string) {
	o.DownloadLanguages = v
}

// GetDownloadMovieSubtitles returns the DownloadMovieSubtitles field value if set, zero value otherwise.
func (o *SubtitleOptions) GetDownloadMovieSubtitles() bool {
	if o == nil || IsNil(o.DownloadMovieSubtitles) {
		var ret bool
		return ret
	}
	return *o.DownloadMovieSubtitles
}

// GetDownloadMovieSubtitlesOk returns a tuple with the DownloadMovieSubtitles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubtitleOptions) GetDownloadMovieSubtitlesOk() (*bool, bool) {
	if o == nil || IsNil(o.DownloadMovieSubtitles) {
		return nil, false
	}
	return o.DownloadMovieSubtitles, true
}

// HasDownloadMovieSubtitles returns a boolean if a field has been set.
func (o *SubtitleOptions) HasDownloadMovieSubtitles() bool {
	if o != nil && !IsNil(o.DownloadMovieSubtitles) {
		return true
	}

	return false
}

// SetDownloadMovieSubtitles gets a reference to the given bool and assigns it to the DownloadMovieSubtitles field.
func (o *SubtitleOptions) SetDownloadMovieSubtitles(v bool) {
	o.DownloadMovieSubtitles = &v
}

// GetDownloadEpisodeSubtitles returns the DownloadEpisodeSubtitles field value if set, zero value otherwise.
func (o *SubtitleOptions) GetDownloadEpisodeSubtitles() bool {
	if o == nil || IsNil(o.DownloadEpisodeSubtitles) {
		var ret bool
		return ret
	}
	return *o.DownloadEpisodeSubtitles
}

// GetDownloadEpisodeSubtitlesOk returns a tuple with the DownloadEpisodeSubtitles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubtitleOptions) GetDownloadEpisodeSubtitlesOk() (*bool, bool) {
	if o == nil || IsNil(o.DownloadEpisodeSubtitles) {
		return nil, false
	}
	return o.DownloadEpisodeSubtitles, true
}

// HasDownloadEpisodeSubtitles returns a boolean if a field has been set.
func (o *SubtitleOptions) HasDownloadEpisodeSubtitles() bool {
	if o != nil && !IsNil(o.DownloadEpisodeSubtitles) {
		return true
	}

	return false
}

// SetDownloadEpisodeSubtitles gets a reference to the given bool and assigns it to the DownloadEpisodeSubtitles field.
func (o *SubtitleOptions) SetDownloadEpisodeSubtitles(v bool) {
	o.DownloadEpisodeSubtitles = &v
}

// GetOpenSubtitlesUsername returns the OpenSubtitlesUsername field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubtitleOptions) GetOpenSubtitlesUsername() string {
	if o == nil || IsNil(o.OpenSubtitlesUsername.Get()) {
		var ret string
		return ret
	}
	return *o.OpenSubtitlesUsername.Get()
}

// GetOpenSubtitlesUsernameOk returns a tuple with the OpenSubtitlesUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubtitleOptions) GetOpenSubtitlesUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenSubtitlesUsername.Get(), o.OpenSubtitlesUsername.IsSet()
}

// HasOpenSubtitlesUsername returns a boolean if a field has been set.
func (o *SubtitleOptions) HasOpenSubtitlesUsername() bool {
	if o != nil && o.OpenSubtitlesUsername.IsSet() {
		return true
	}

	return false
}

// SetOpenSubtitlesUsername gets a reference to the given NullableString and assigns it to the OpenSubtitlesUsername field.
func (o *SubtitleOptions) SetOpenSubtitlesUsername(v string) {
	o.OpenSubtitlesUsername.Set(&v)
}
// SetOpenSubtitlesUsernameNil sets the value for OpenSubtitlesUsername to be an explicit nil
func (o *SubtitleOptions) SetOpenSubtitlesUsernameNil() {
	o.OpenSubtitlesUsername.Set(nil)
}

// UnsetOpenSubtitlesUsername ensures that no value is present for OpenSubtitlesUsername, not even an explicit nil
func (o *SubtitleOptions) UnsetOpenSubtitlesUsername() {
	o.OpenSubtitlesUsername.Unset()
}

// GetOpenSubtitlesPasswordHash returns the OpenSubtitlesPasswordHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SubtitleOptions) GetOpenSubtitlesPasswordHash() string {
	if o == nil || IsNil(o.OpenSubtitlesPasswordHash.Get()) {
		var ret string
		return ret
	}
	return *o.OpenSubtitlesPasswordHash.Get()
}

// GetOpenSubtitlesPasswordHashOk returns a tuple with the OpenSubtitlesPasswordHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SubtitleOptions) GetOpenSubtitlesPasswordHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OpenSubtitlesPasswordHash.Get(), o.OpenSubtitlesPasswordHash.IsSet()
}

// HasOpenSubtitlesPasswordHash returns a boolean if a field has been set.
func (o *SubtitleOptions) HasOpenSubtitlesPasswordHash() bool {
	if o != nil && o.OpenSubtitlesPasswordHash.IsSet() {
		return true
	}

	return false
}

// SetOpenSubtitlesPasswordHash gets a reference to the given NullableString and assigns it to the OpenSubtitlesPasswordHash field.
func (o *SubtitleOptions) SetOpenSubtitlesPasswordHash(v string) {
	o.OpenSubtitlesPasswordHash.Set(&v)
}
// SetOpenSubtitlesPasswordHashNil sets the value for OpenSubtitlesPasswordHash to be an explicit nil
func (o *SubtitleOptions) SetOpenSubtitlesPasswordHashNil() {
	o.OpenSubtitlesPasswordHash.Set(nil)
}

// UnsetOpenSubtitlesPasswordHash ensures that no value is present for OpenSubtitlesPasswordHash, not even an explicit nil
func (o *SubtitleOptions) UnsetOpenSubtitlesPasswordHash() {
	o.OpenSubtitlesPasswordHash.Unset()
}

// GetIsOpenSubtitleVipAccount returns the IsOpenSubtitleVipAccount field value if set, zero value otherwise.
func (o *SubtitleOptions) GetIsOpenSubtitleVipAccount() bool {
	if o == nil || IsNil(o.IsOpenSubtitleVipAccount) {
		var ret bool
		return ret
	}
	return *o.IsOpenSubtitleVipAccount
}

// GetIsOpenSubtitleVipAccountOk returns a tuple with the IsOpenSubtitleVipAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubtitleOptions) GetIsOpenSubtitleVipAccountOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOpenSubtitleVipAccount) {
		return nil, false
	}
	return o.IsOpenSubtitleVipAccount, true
}

// HasIsOpenSubtitleVipAccount returns a boolean if a field has been set.
func (o *SubtitleOptions) HasIsOpenSubtitleVipAccount() bool {
	if o != nil && !IsNil(o.IsOpenSubtitleVipAccount) {
		return true
	}

	return false
}

// SetIsOpenSubtitleVipAccount gets a reference to the given bool and assigns it to the IsOpenSubtitleVipAccount field.
func (o *SubtitleOptions) SetIsOpenSubtitleVipAccount(v bool) {
	o.IsOpenSubtitleVipAccount = &v
}

// GetRequirePerfectMatch returns the RequirePerfectMatch field value if set, zero value otherwise.
func (o *SubtitleOptions) GetRequirePerfectMatch() bool {
	if o == nil || IsNil(o.RequirePerfectMatch) {
		var ret bool
		return ret
	}
	return *o.RequirePerfectMatch
}

// GetRequirePerfectMatchOk returns a tuple with the RequirePerfectMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubtitleOptions) GetRequirePerfectMatchOk() (*bool, bool) {
	if o == nil || IsNil(o.RequirePerfectMatch) {
		return nil, false
	}
	return o.RequirePerfectMatch, true
}

// HasRequirePerfectMatch returns a boolean if a field has been set.
func (o *SubtitleOptions) HasRequirePerfectMatch() bool {
	if o != nil && !IsNil(o.RequirePerfectMatch) {
		return true
	}

	return false
}

// SetRequirePerfectMatch gets a reference to the given bool and assigns it to the RequirePerfectMatch field.
func (o *SubtitleOptions) SetRequirePerfectMatch(v bool) {
	o.RequirePerfectMatch = &v
}

func (o SubtitleOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubtitleOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SkipIfEmbeddedSubtitlesPresent) {
		toSerialize["SkipIfEmbeddedSubtitlesPresent"] = o.SkipIfEmbeddedSubtitlesPresent
	}
	if !IsNil(o.SkipIfAudioTrackMatches) {
		toSerialize["SkipIfAudioTrackMatches"] = o.SkipIfAudioTrackMatches
	}
	if o.DownloadLanguages != nil {
		toSerialize["DownloadLanguages"] = o.DownloadLanguages
	}
	if !IsNil(o.DownloadMovieSubtitles) {
		toSerialize["DownloadMovieSubtitles"] = o.DownloadMovieSubtitles
	}
	if !IsNil(o.DownloadEpisodeSubtitles) {
		toSerialize["DownloadEpisodeSubtitles"] = o.DownloadEpisodeSubtitles
	}
	if o.OpenSubtitlesUsername.IsSet() {
		toSerialize["OpenSubtitlesUsername"] = o.OpenSubtitlesUsername.Get()
	}
	if o.OpenSubtitlesPasswordHash.IsSet() {
		toSerialize["OpenSubtitlesPasswordHash"] = o.OpenSubtitlesPasswordHash.Get()
	}
	if !IsNil(o.IsOpenSubtitleVipAccount) {
		toSerialize["IsOpenSubtitleVipAccount"] = o.IsOpenSubtitleVipAccount
	}
	if !IsNil(o.RequirePerfectMatch) {
		toSerialize["RequirePerfectMatch"] = o.RequirePerfectMatch
	}
	return toSerialize, nil
}

type NullableSubtitleOptions struct {
	value *SubtitleOptions
	isSet bool
}

func (v NullableSubtitleOptions) Get() *SubtitleOptions {
	return v.value
}

func (v *NullableSubtitleOptions) Set(val *SubtitleOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSubtitleOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSubtitleOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubtitleOptions(val *SubtitleOptions) *NullableSubtitleOptions {
	return &NullableSubtitleOptions{value: val, isSet: true}
}

func (v NullableSubtitleOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubtitleOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


