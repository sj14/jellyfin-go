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

// checks if the LiveTvOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LiveTvOptions{}

// LiveTvOptions struct for LiveTvOptions
type LiveTvOptions struct {
	GuideDays NullableInt32 `json:"GuideDays,omitempty"`
	RecordingPath NullableString `json:"RecordingPath,omitempty"`
	MovieRecordingPath NullableString `json:"MovieRecordingPath,omitempty"`
	SeriesRecordingPath NullableString `json:"SeriesRecordingPath,omitempty"`
	EnableRecordingSubfolders *bool `json:"EnableRecordingSubfolders,omitempty"`
	EnableOriginalAudioWithEncodedRecordings *bool `json:"EnableOriginalAudioWithEncodedRecordings,omitempty"`
	TunerHosts []TunerHostInfo `json:"TunerHosts,omitempty"`
	ListingProviders []ListingsProviderInfo `json:"ListingProviders,omitempty"`
	PrePaddingSeconds *int32 `json:"PrePaddingSeconds,omitempty"`
	PostPaddingSeconds *int32 `json:"PostPaddingSeconds,omitempty"`
	MediaLocationsCreated []string `json:"MediaLocationsCreated,omitempty"`
	RecordingPostProcessor NullableString `json:"RecordingPostProcessor,omitempty"`
	RecordingPostProcessorArguments NullableString `json:"RecordingPostProcessorArguments,omitempty"`
	SaveRecordingNFO *bool `json:"SaveRecordingNFO,omitempty"`
	SaveRecordingImages *bool `json:"SaveRecordingImages,omitempty"`
}

// NewLiveTvOptions instantiates a new LiveTvOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveTvOptions() *LiveTvOptions {
	this := LiveTvOptions{}
	return &this
}

// NewLiveTvOptionsWithDefaults instantiates a new LiveTvOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveTvOptionsWithDefaults() *LiveTvOptions {
	this := LiveTvOptions{}
	return &this
}

// GetGuideDays returns the GuideDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LiveTvOptions) GetGuideDays() int32 {
	if o == nil || IsNil(o.GuideDays.Get()) {
		var ret int32
		return ret
	}
	return *o.GuideDays.Get()
}

// GetGuideDaysOk returns a tuple with the GuideDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveTvOptions) GetGuideDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GuideDays.Get(), o.GuideDays.IsSet()
}

// HasGuideDays returns a boolean if a field has been set.
func (o *LiveTvOptions) HasGuideDays() bool {
	if o != nil && o.GuideDays.IsSet() {
		return true
	}

	return false
}

// SetGuideDays gets a reference to the given NullableInt32 and assigns it to the GuideDays field.
func (o *LiveTvOptions) SetGuideDays(v int32) {
	o.GuideDays.Set(&v)
}
// SetGuideDaysNil sets the value for GuideDays to be an explicit nil
func (o *LiveTvOptions) SetGuideDaysNil() {
	o.GuideDays.Set(nil)
}

// UnsetGuideDays ensures that no value is present for GuideDays, not even an explicit nil
func (o *LiveTvOptions) UnsetGuideDays() {
	o.GuideDays.Unset()
}

// GetRecordingPath returns the RecordingPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LiveTvOptions) GetRecordingPath() string {
	if o == nil || IsNil(o.RecordingPath.Get()) {
		var ret string
		return ret
	}
	return *o.RecordingPath.Get()
}

// GetRecordingPathOk returns a tuple with the RecordingPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveTvOptions) GetRecordingPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecordingPath.Get(), o.RecordingPath.IsSet()
}

// HasRecordingPath returns a boolean if a field has been set.
func (o *LiveTvOptions) HasRecordingPath() bool {
	if o != nil && o.RecordingPath.IsSet() {
		return true
	}

	return false
}

// SetRecordingPath gets a reference to the given NullableString and assigns it to the RecordingPath field.
func (o *LiveTvOptions) SetRecordingPath(v string) {
	o.RecordingPath.Set(&v)
}
// SetRecordingPathNil sets the value for RecordingPath to be an explicit nil
func (o *LiveTvOptions) SetRecordingPathNil() {
	o.RecordingPath.Set(nil)
}

// UnsetRecordingPath ensures that no value is present for RecordingPath, not even an explicit nil
func (o *LiveTvOptions) UnsetRecordingPath() {
	o.RecordingPath.Unset()
}

// GetMovieRecordingPath returns the MovieRecordingPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LiveTvOptions) GetMovieRecordingPath() string {
	if o == nil || IsNil(o.MovieRecordingPath.Get()) {
		var ret string
		return ret
	}
	return *o.MovieRecordingPath.Get()
}

// GetMovieRecordingPathOk returns a tuple with the MovieRecordingPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveTvOptions) GetMovieRecordingPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MovieRecordingPath.Get(), o.MovieRecordingPath.IsSet()
}

// HasMovieRecordingPath returns a boolean if a field has been set.
func (o *LiveTvOptions) HasMovieRecordingPath() bool {
	if o != nil && o.MovieRecordingPath.IsSet() {
		return true
	}

	return false
}

// SetMovieRecordingPath gets a reference to the given NullableString and assigns it to the MovieRecordingPath field.
func (o *LiveTvOptions) SetMovieRecordingPath(v string) {
	o.MovieRecordingPath.Set(&v)
}
// SetMovieRecordingPathNil sets the value for MovieRecordingPath to be an explicit nil
func (o *LiveTvOptions) SetMovieRecordingPathNil() {
	o.MovieRecordingPath.Set(nil)
}

// UnsetMovieRecordingPath ensures that no value is present for MovieRecordingPath, not even an explicit nil
func (o *LiveTvOptions) UnsetMovieRecordingPath() {
	o.MovieRecordingPath.Unset()
}

// GetSeriesRecordingPath returns the SeriesRecordingPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LiveTvOptions) GetSeriesRecordingPath() string {
	if o == nil || IsNil(o.SeriesRecordingPath.Get()) {
		var ret string
		return ret
	}
	return *o.SeriesRecordingPath.Get()
}

// GetSeriesRecordingPathOk returns a tuple with the SeriesRecordingPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveTvOptions) GetSeriesRecordingPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SeriesRecordingPath.Get(), o.SeriesRecordingPath.IsSet()
}

// HasSeriesRecordingPath returns a boolean if a field has been set.
func (o *LiveTvOptions) HasSeriesRecordingPath() bool {
	if o != nil && o.SeriesRecordingPath.IsSet() {
		return true
	}

	return false
}

// SetSeriesRecordingPath gets a reference to the given NullableString and assigns it to the SeriesRecordingPath field.
func (o *LiveTvOptions) SetSeriesRecordingPath(v string) {
	o.SeriesRecordingPath.Set(&v)
}
// SetSeriesRecordingPathNil sets the value for SeriesRecordingPath to be an explicit nil
func (o *LiveTvOptions) SetSeriesRecordingPathNil() {
	o.SeriesRecordingPath.Set(nil)
}

// UnsetSeriesRecordingPath ensures that no value is present for SeriesRecordingPath, not even an explicit nil
func (o *LiveTvOptions) UnsetSeriesRecordingPath() {
	o.SeriesRecordingPath.Unset()
}

// GetEnableRecordingSubfolders returns the EnableRecordingSubfolders field value if set, zero value otherwise.
func (o *LiveTvOptions) GetEnableRecordingSubfolders() bool {
	if o == nil || IsNil(o.EnableRecordingSubfolders) {
		var ret bool
		return ret
	}
	return *o.EnableRecordingSubfolders
}

// GetEnableRecordingSubfoldersOk returns a tuple with the EnableRecordingSubfolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvOptions) GetEnableRecordingSubfoldersOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableRecordingSubfolders) {
		return nil, false
	}
	return o.EnableRecordingSubfolders, true
}

// HasEnableRecordingSubfolders returns a boolean if a field has been set.
func (o *LiveTvOptions) HasEnableRecordingSubfolders() bool {
	if o != nil && !IsNil(o.EnableRecordingSubfolders) {
		return true
	}

	return false
}

// SetEnableRecordingSubfolders gets a reference to the given bool and assigns it to the EnableRecordingSubfolders field.
func (o *LiveTvOptions) SetEnableRecordingSubfolders(v bool) {
	o.EnableRecordingSubfolders = &v
}

// GetEnableOriginalAudioWithEncodedRecordings returns the EnableOriginalAudioWithEncodedRecordings field value if set, zero value otherwise.
func (o *LiveTvOptions) GetEnableOriginalAudioWithEncodedRecordings() bool {
	if o == nil || IsNil(o.EnableOriginalAudioWithEncodedRecordings) {
		var ret bool
		return ret
	}
	return *o.EnableOriginalAudioWithEncodedRecordings
}

// GetEnableOriginalAudioWithEncodedRecordingsOk returns a tuple with the EnableOriginalAudioWithEncodedRecordings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvOptions) GetEnableOriginalAudioWithEncodedRecordingsOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableOriginalAudioWithEncodedRecordings) {
		return nil, false
	}
	return o.EnableOriginalAudioWithEncodedRecordings, true
}

// HasEnableOriginalAudioWithEncodedRecordings returns a boolean if a field has been set.
func (o *LiveTvOptions) HasEnableOriginalAudioWithEncodedRecordings() bool {
	if o != nil && !IsNil(o.EnableOriginalAudioWithEncodedRecordings) {
		return true
	}

	return false
}

// SetEnableOriginalAudioWithEncodedRecordings gets a reference to the given bool and assigns it to the EnableOriginalAudioWithEncodedRecordings field.
func (o *LiveTvOptions) SetEnableOriginalAudioWithEncodedRecordings(v bool) {
	o.EnableOriginalAudioWithEncodedRecordings = &v
}

// GetTunerHosts returns the TunerHosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LiveTvOptions) GetTunerHosts() []TunerHostInfo {
	if o == nil {
		var ret []TunerHostInfo
		return ret
	}
	return o.TunerHosts
}

// GetTunerHostsOk returns a tuple with the TunerHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveTvOptions) GetTunerHostsOk() ([]TunerHostInfo, bool) {
	if o == nil || IsNil(o.TunerHosts) {
		return nil, false
	}
	return o.TunerHosts, true
}

// HasTunerHosts returns a boolean if a field has been set.
func (o *LiveTvOptions) HasTunerHosts() bool {
	if o != nil && !IsNil(o.TunerHosts) {
		return true
	}

	return false
}

// SetTunerHosts gets a reference to the given []TunerHostInfo and assigns it to the TunerHosts field.
func (o *LiveTvOptions) SetTunerHosts(v []TunerHostInfo) {
	o.TunerHosts = v
}

// GetListingProviders returns the ListingProviders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LiveTvOptions) GetListingProviders() []ListingsProviderInfo {
	if o == nil {
		var ret []ListingsProviderInfo
		return ret
	}
	return o.ListingProviders
}

// GetListingProvidersOk returns a tuple with the ListingProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveTvOptions) GetListingProvidersOk() ([]ListingsProviderInfo, bool) {
	if o == nil || IsNil(o.ListingProviders) {
		return nil, false
	}
	return o.ListingProviders, true
}

// HasListingProviders returns a boolean if a field has been set.
func (o *LiveTvOptions) HasListingProviders() bool {
	if o != nil && !IsNil(o.ListingProviders) {
		return true
	}

	return false
}

// SetListingProviders gets a reference to the given []ListingsProviderInfo and assigns it to the ListingProviders field.
func (o *LiveTvOptions) SetListingProviders(v []ListingsProviderInfo) {
	o.ListingProviders = v
}

// GetPrePaddingSeconds returns the PrePaddingSeconds field value if set, zero value otherwise.
func (o *LiveTvOptions) GetPrePaddingSeconds() int32 {
	if o == nil || IsNil(o.PrePaddingSeconds) {
		var ret int32
		return ret
	}
	return *o.PrePaddingSeconds
}

// GetPrePaddingSecondsOk returns a tuple with the PrePaddingSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvOptions) GetPrePaddingSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.PrePaddingSeconds) {
		return nil, false
	}
	return o.PrePaddingSeconds, true
}

// HasPrePaddingSeconds returns a boolean if a field has been set.
func (o *LiveTvOptions) HasPrePaddingSeconds() bool {
	if o != nil && !IsNil(o.PrePaddingSeconds) {
		return true
	}

	return false
}

// SetPrePaddingSeconds gets a reference to the given int32 and assigns it to the PrePaddingSeconds field.
func (o *LiveTvOptions) SetPrePaddingSeconds(v int32) {
	o.PrePaddingSeconds = &v
}

// GetPostPaddingSeconds returns the PostPaddingSeconds field value if set, zero value otherwise.
func (o *LiveTvOptions) GetPostPaddingSeconds() int32 {
	if o == nil || IsNil(o.PostPaddingSeconds) {
		var ret int32
		return ret
	}
	return *o.PostPaddingSeconds
}

// GetPostPaddingSecondsOk returns a tuple with the PostPaddingSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvOptions) GetPostPaddingSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.PostPaddingSeconds) {
		return nil, false
	}
	return o.PostPaddingSeconds, true
}

// HasPostPaddingSeconds returns a boolean if a field has been set.
func (o *LiveTvOptions) HasPostPaddingSeconds() bool {
	if o != nil && !IsNil(o.PostPaddingSeconds) {
		return true
	}

	return false
}

// SetPostPaddingSeconds gets a reference to the given int32 and assigns it to the PostPaddingSeconds field.
func (o *LiveTvOptions) SetPostPaddingSeconds(v int32) {
	o.PostPaddingSeconds = &v
}

// GetMediaLocationsCreated returns the MediaLocationsCreated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LiveTvOptions) GetMediaLocationsCreated() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MediaLocationsCreated
}

// GetMediaLocationsCreatedOk returns a tuple with the MediaLocationsCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveTvOptions) GetMediaLocationsCreatedOk() ([]string, bool) {
	if o == nil || IsNil(o.MediaLocationsCreated) {
		return nil, false
	}
	return o.MediaLocationsCreated, true
}

// HasMediaLocationsCreated returns a boolean if a field has been set.
func (o *LiveTvOptions) HasMediaLocationsCreated() bool {
	if o != nil && !IsNil(o.MediaLocationsCreated) {
		return true
	}

	return false
}

// SetMediaLocationsCreated gets a reference to the given []string and assigns it to the MediaLocationsCreated field.
func (o *LiveTvOptions) SetMediaLocationsCreated(v []string) {
	o.MediaLocationsCreated = v
}

// GetRecordingPostProcessor returns the RecordingPostProcessor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LiveTvOptions) GetRecordingPostProcessor() string {
	if o == nil || IsNil(o.RecordingPostProcessor.Get()) {
		var ret string
		return ret
	}
	return *o.RecordingPostProcessor.Get()
}

// GetRecordingPostProcessorOk returns a tuple with the RecordingPostProcessor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveTvOptions) GetRecordingPostProcessorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecordingPostProcessor.Get(), o.RecordingPostProcessor.IsSet()
}

// HasRecordingPostProcessor returns a boolean if a field has been set.
func (o *LiveTvOptions) HasRecordingPostProcessor() bool {
	if o != nil && o.RecordingPostProcessor.IsSet() {
		return true
	}

	return false
}

// SetRecordingPostProcessor gets a reference to the given NullableString and assigns it to the RecordingPostProcessor field.
func (o *LiveTvOptions) SetRecordingPostProcessor(v string) {
	o.RecordingPostProcessor.Set(&v)
}
// SetRecordingPostProcessorNil sets the value for RecordingPostProcessor to be an explicit nil
func (o *LiveTvOptions) SetRecordingPostProcessorNil() {
	o.RecordingPostProcessor.Set(nil)
}

// UnsetRecordingPostProcessor ensures that no value is present for RecordingPostProcessor, not even an explicit nil
func (o *LiveTvOptions) UnsetRecordingPostProcessor() {
	o.RecordingPostProcessor.Unset()
}

// GetRecordingPostProcessorArguments returns the RecordingPostProcessorArguments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LiveTvOptions) GetRecordingPostProcessorArguments() string {
	if o == nil || IsNil(o.RecordingPostProcessorArguments.Get()) {
		var ret string
		return ret
	}
	return *o.RecordingPostProcessorArguments.Get()
}

// GetRecordingPostProcessorArgumentsOk returns a tuple with the RecordingPostProcessorArguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveTvOptions) GetRecordingPostProcessorArgumentsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RecordingPostProcessorArguments.Get(), o.RecordingPostProcessorArguments.IsSet()
}

// HasRecordingPostProcessorArguments returns a boolean if a field has been set.
func (o *LiveTvOptions) HasRecordingPostProcessorArguments() bool {
	if o != nil && o.RecordingPostProcessorArguments.IsSet() {
		return true
	}

	return false
}

// SetRecordingPostProcessorArguments gets a reference to the given NullableString and assigns it to the RecordingPostProcessorArguments field.
func (o *LiveTvOptions) SetRecordingPostProcessorArguments(v string) {
	o.RecordingPostProcessorArguments.Set(&v)
}
// SetRecordingPostProcessorArgumentsNil sets the value for RecordingPostProcessorArguments to be an explicit nil
func (o *LiveTvOptions) SetRecordingPostProcessorArgumentsNil() {
	o.RecordingPostProcessorArguments.Set(nil)
}

// UnsetRecordingPostProcessorArguments ensures that no value is present for RecordingPostProcessorArguments, not even an explicit nil
func (o *LiveTvOptions) UnsetRecordingPostProcessorArguments() {
	o.RecordingPostProcessorArguments.Unset()
}

// GetSaveRecordingNFO returns the SaveRecordingNFO field value if set, zero value otherwise.
func (o *LiveTvOptions) GetSaveRecordingNFO() bool {
	if o == nil || IsNil(o.SaveRecordingNFO) {
		var ret bool
		return ret
	}
	return *o.SaveRecordingNFO
}

// GetSaveRecordingNFOOk returns a tuple with the SaveRecordingNFO field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvOptions) GetSaveRecordingNFOOk() (*bool, bool) {
	if o == nil || IsNil(o.SaveRecordingNFO) {
		return nil, false
	}
	return o.SaveRecordingNFO, true
}

// HasSaveRecordingNFO returns a boolean if a field has been set.
func (o *LiveTvOptions) HasSaveRecordingNFO() bool {
	if o != nil && !IsNil(o.SaveRecordingNFO) {
		return true
	}

	return false
}

// SetSaveRecordingNFO gets a reference to the given bool and assigns it to the SaveRecordingNFO field.
func (o *LiveTvOptions) SetSaveRecordingNFO(v bool) {
	o.SaveRecordingNFO = &v
}

// GetSaveRecordingImages returns the SaveRecordingImages field value if set, zero value otherwise.
func (o *LiveTvOptions) GetSaveRecordingImages() bool {
	if o == nil || IsNil(o.SaveRecordingImages) {
		var ret bool
		return ret
	}
	return *o.SaveRecordingImages
}

// GetSaveRecordingImagesOk returns a tuple with the SaveRecordingImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveTvOptions) GetSaveRecordingImagesOk() (*bool, bool) {
	if o == nil || IsNil(o.SaveRecordingImages) {
		return nil, false
	}
	return o.SaveRecordingImages, true
}

// HasSaveRecordingImages returns a boolean if a field has been set.
func (o *LiveTvOptions) HasSaveRecordingImages() bool {
	if o != nil && !IsNil(o.SaveRecordingImages) {
		return true
	}

	return false
}

// SetSaveRecordingImages gets a reference to the given bool and assigns it to the SaveRecordingImages field.
func (o *LiveTvOptions) SetSaveRecordingImages(v bool) {
	o.SaveRecordingImages = &v
}

func (o LiveTvOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LiveTvOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GuideDays.IsSet() {
		toSerialize["GuideDays"] = o.GuideDays.Get()
	}
	if o.RecordingPath.IsSet() {
		toSerialize["RecordingPath"] = o.RecordingPath.Get()
	}
	if o.MovieRecordingPath.IsSet() {
		toSerialize["MovieRecordingPath"] = o.MovieRecordingPath.Get()
	}
	if o.SeriesRecordingPath.IsSet() {
		toSerialize["SeriesRecordingPath"] = o.SeriesRecordingPath.Get()
	}
	if !IsNil(o.EnableRecordingSubfolders) {
		toSerialize["EnableRecordingSubfolders"] = o.EnableRecordingSubfolders
	}
	if !IsNil(o.EnableOriginalAudioWithEncodedRecordings) {
		toSerialize["EnableOriginalAudioWithEncodedRecordings"] = o.EnableOriginalAudioWithEncodedRecordings
	}
	if o.TunerHosts != nil {
		toSerialize["TunerHosts"] = o.TunerHosts
	}
	if o.ListingProviders != nil {
		toSerialize["ListingProviders"] = o.ListingProviders
	}
	if !IsNil(o.PrePaddingSeconds) {
		toSerialize["PrePaddingSeconds"] = o.PrePaddingSeconds
	}
	if !IsNil(o.PostPaddingSeconds) {
		toSerialize["PostPaddingSeconds"] = o.PostPaddingSeconds
	}
	if o.MediaLocationsCreated != nil {
		toSerialize["MediaLocationsCreated"] = o.MediaLocationsCreated
	}
	if o.RecordingPostProcessor.IsSet() {
		toSerialize["RecordingPostProcessor"] = o.RecordingPostProcessor.Get()
	}
	if o.RecordingPostProcessorArguments.IsSet() {
		toSerialize["RecordingPostProcessorArguments"] = o.RecordingPostProcessorArguments.Get()
	}
	if !IsNil(o.SaveRecordingNFO) {
		toSerialize["SaveRecordingNFO"] = o.SaveRecordingNFO
	}
	if !IsNil(o.SaveRecordingImages) {
		toSerialize["SaveRecordingImages"] = o.SaveRecordingImages
	}
	return toSerialize, nil
}

type NullableLiveTvOptions struct {
	value *LiveTvOptions
	isSet bool
}

func (v NullableLiveTvOptions) Get() *LiveTvOptions {
	return v.value
}

func (v *NullableLiveTvOptions) Set(val *LiveTvOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveTvOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveTvOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveTvOptions(val *LiveTvOptions) *NullableLiveTvOptions {
	return &NullableLiveTvOptions{value: val, isSet: true}
}

func (v NullableLiveTvOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveTvOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


