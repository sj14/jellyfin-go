/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the GetProgramsDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProgramsDto{}

// GetProgramsDto Get programs dto.
type GetProgramsDto struct {
	// Gets or sets the channels to return guide information for.
	ChannelIds []string `json:"ChannelIds,omitempty"`
	// Gets or sets optional. Filter by user id.
	UserId NullableString `json:"UserId,omitempty"`
	// Gets or sets the minimum premiere start date.
	MinStartDate NullableTime `json:"MinStartDate,omitempty"`
	// Gets or sets filter by programs that have completed airing, or not.
	HasAired NullableBool `json:"HasAired,omitempty"`
	// Gets or sets filter by programs that are currently airing, or not.
	IsAiring NullableBool `json:"IsAiring,omitempty"`
	// Gets or sets the maximum premiere start date.
	MaxStartDate NullableTime `json:"MaxStartDate,omitempty"`
	// Gets or sets the minimum premiere end date.
	MinEndDate NullableTime `json:"MinEndDate,omitempty"`
	// Gets or sets the maximum premiere end date.
	MaxEndDate NullableTime `json:"MaxEndDate,omitempty"`
	// Gets or sets filter for movies.
	IsMovie NullableBool `json:"IsMovie,omitempty"`
	// Gets or sets filter for series.
	IsSeries NullableBool `json:"IsSeries,omitempty"`
	// Gets or sets filter for news.
	IsNews NullableBool `json:"IsNews,omitempty"`
	// Gets or sets filter for kids.
	IsKids NullableBool `json:"IsKids,omitempty"`
	// Gets or sets filter for sports.
	IsSports NullableBool `json:"IsSports,omitempty"`
	// Gets or sets the record index to start at. All items with a lower index will be dropped from the results.
	StartIndex NullableInt32 `json:"StartIndex,omitempty"`
	// Gets or sets the maximum number of records to return.
	Limit NullableInt32 `json:"Limit,omitempty"`
	// Gets or sets specify one or more sort orders, comma delimited. Options: Name, StartDate.
	SortBy []ItemSortBy `json:"SortBy,omitempty"`
	// Gets or sets sort order.
	SortOrder []SortOrder `json:"SortOrder,omitempty"`
	// Gets or sets the genres to return guide information for.
	Genres []string `json:"Genres,omitempty"`
	// Gets or sets the genre ids to return guide information for.
	GenreIds []string `json:"GenreIds,omitempty"`
	// Gets or sets include image information in output.
	EnableImages NullableBool `json:"EnableImages,omitempty"`
	// Gets or sets a value indicating whether retrieve total record count.
	EnableTotalRecordCount *bool `json:"EnableTotalRecordCount,omitempty"`
	// Gets or sets the max number of images to return, per image type.
	ImageTypeLimit NullableInt32 `json:"ImageTypeLimit,omitempty"`
	// Gets or sets the image types to include in the output.
	EnableImageTypes []ImageType `json:"EnableImageTypes,omitempty"`
	// Gets or sets include user data.
	EnableUserData NullableBool `json:"EnableUserData,omitempty"`
	// Gets or sets filter by series timer id.
	SeriesTimerId NullableString `json:"SeriesTimerId,omitempty"`
	// Gets or sets filter by library series id.
	LibrarySeriesId NullableString `json:"LibrarySeriesId,omitempty"`
	// Gets or sets specify additional fields of information to return in the output.
	Fields []ItemFields `json:"Fields,omitempty"`
}

// NewGetProgramsDto instantiates a new GetProgramsDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProgramsDto() *GetProgramsDto {
	this := GetProgramsDto{}
	var enableTotalRecordCount bool = true
	this.EnableTotalRecordCount = &enableTotalRecordCount
	return &this
}

// NewGetProgramsDtoWithDefaults instantiates a new GetProgramsDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProgramsDtoWithDefaults() *GetProgramsDto {
	this := GetProgramsDto{}
	var enableTotalRecordCount bool = true
	this.EnableTotalRecordCount = &enableTotalRecordCount
	return &this
}

// GetChannelIds returns the ChannelIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetChannelIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ChannelIds
}

// GetChannelIdsOk returns a tuple with the ChannelIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetChannelIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ChannelIds) {
		return nil, false
	}
	return o.ChannelIds, true
}

// HasChannelIds returns a boolean if a field has been set.
func (o *GetProgramsDto) HasChannelIds() bool {
	if o != nil && !IsNil(o.ChannelIds) {
		return true
	}

	return false
}

// SetChannelIds gets a reference to the given []string and assigns it to the ChannelIds field.
func (o *GetProgramsDto) SetChannelIds(v []string) {
	o.ChannelIds = v
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetUserId() string {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *GetProgramsDto) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *GetProgramsDto) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *GetProgramsDto) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *GetProgramsDto) UnsetUserId() {
	o.UserId.Unset()
}

// GetMinStartDate returns the MinStartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetMinStartDate() time.Time {
	if o == nil || IsNil(o.MinStartDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.MinStartDate.Get()
}

// GetMinStartDateOk returns a tuple with the MinStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetMinStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinStartDate.Get(), o.MinStartDate.IsSet()
}

// HasMinStartDate returns a boolean if a field has been set.
func (o *GetProgramsDto) HasMinStartDate() bool {
	if o != nil && o.MinStartDate.IsSet() {
		return true
	}

	return false
}

// SetMinStartDate gets a reference to the given NullableTime and assigns it to the MinStartDate field.
func (o *GetProgramsDto) SetMinStartDate(v time.Time) {
	o.MinStartDate.Set(&v)
}
// SetMinStartDateNil sets the value for MinStartDate to be an explicit nil
func (o *GetProgramsDto) SetMinStartDateNil() {
	o.MinStartDate.Set(nil)
}

// UnsetMinStartDate ensures that no value is present for MinStartDate, not even an explicit nil
func (o *GetProgramsDto) UnsetMinStartDate() {
	o.MinStartDate.Unset()
}

// GetHasAired returns the HasAired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetHasAired() bool {
	if o == nil || IsNil(o.HasAired.Get()) {
		var ret bool
		return ret
	}
	return *o.HasAired.Get()
}

// GetHasAiredOk returns a tuple with the HasAired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetHasAiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HasAired.Get(), o.HasAired.IsSet()
}

// HasHasAired returns a boolean if a field has been set.
func (o *GetProgramsDto) HasHasAired() bool {
	if o != nil && o.HasAired.IsSet() {
		return true
	}

	return false
}

// SetHasAired gets a reference to the given NullableBool and assigns it to the HasAired field.
func (o *GetProgramsDto) SetHasAired(v bool) {
	o.HasAired.Set(&v)
}
// SetHasAiredNil sets the value for HasAired to be an explicit nil
func (o *GetProgramsDto) SetHasAiredNil() {
	o.HasAired.Set(nil)
}

// UnsetHasAired ensures that no value is present for HasAired, not even an explicit nil
func (o *GetProgramsDto) UnsetHasAired() {
	o.HasAired.Unset()
}

// GetIsAiring returns the IsAiring field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetIsAiring() bool {
	if o == nil || IsNil(o.IsAiring.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAiring.Get()
}

// GetIsAiringOk returns a tuple with the IsAiring field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetIsAiringOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAiring.Get(), o.IsAiring.IsSet()
}

// HasIsAiring returns a boolean if a field has been set.
func (o *GetProgramsDto) HasIsAiring() bool {
	if o != nil && o.IsAiring.IsSet() {
		return true
	}

	return false
}

// SetIsAiring gets a reference to the given NullableBool and assigns it to the IsAiring field.
func (o *GetProgramsDto) SetIsAiring(v bool) {
	o.IsAiring.Set(&v)
}
// SetIsAiringNil sets the value for IsAiring to be an explicit nil
func (o *GetProgramsDto) SetIsAiringNil() {
	o.IsAiring.Set(nil)
}

// UnsetIsAiring ensures that no value is present for IsAiring, not even an explicit nil
func (o *GetProgramsDto) UnsetIsAiring() {
	o.IsAiring.Unset()
}

// GetMaxStartDate returns the MaxStartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetMaxStartDate() time.Time {
	if o == nil || IsNil(o.MaxStartDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.MaxStartDate.Get()
}

// GetMaxStartDateOk returns a tuple with the MaxStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetMaxStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxStartDate.Get(), o.MaxStartDate.IsSet()
}

// HasMaxStartDate returns a boolean if a field has been set.
func (o *GetProgramsDto) HasMaxStartDate() bool {
	if o != nil && o.MaxStartDate.IsSet() {
		return true
	}

	return false
}

// SetMaxStartDate gets a reference to the given NullableTime and assigns it to the MaxStartDate field.
func (o *GetProgramsDto) SetMaxStartDate(v time.Time) {
	o.MaxStartDate.Set(&v)
}
// SetMaxStartDateNil sets the value for MaxStartDate to be an explicit nil
func (o *GetProgramsDto) SetMaxStartDateNil() {
	o.MaxStartDate.Set(nil)
}

// UnsetMaxStartDate ensures that no value is present for MaxStartDate, not even an explicit nil
func (o *GetProgramsDto) UnsetMaxStartDate() {
	o.MaxStartDate.Unset()
}

// GetMinEndDate returns the MinEndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetMinEndDate() time.Time {
	if o == nil || IsNil(o.MinEndDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.MinEndDate.Get()
}

// GetMinEndDateOk returns a tuple with the MinEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetMinEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinEndDate.Get(), o.MinEndDate.IsSet()
}

// HasMinEndDate returns a boolean if a field has been set.
func (o *GetProgramsDto) HasMinEndDate() bool {
	if o != nil && o.MinEndDate.IsSet() {
		return true
	}

	return false
}

// SetMinEndDate gets a reference to the given NullableTime and assigns it to the MinEndDate field.
func (o *GetProgramsDto) SetMinEndDate(v time.Time) {
	o.MinEndDate.Set(&v)
}
// SetMinEndDateNil sets the value for MinEndDate to be an explicit nil
func (o *GetProgramsDto) SetMinEndDateNil() {
	o.MinEndDate.Set(nil)
}

// UnsetMinEndDate ensures that no value is present for MinEndDate, not even an explicit nil
func (o *GetProgramsDto) UnsetMinEndDate() {
	o.MinEndDate.Unset()
}

// GetMaxEndDate returns the MaxEndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetMaxEndDate() time.Time {
	if o == nil || IsNil(o.MaxEndDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.MaxEndDate.Get()
}

// GetMaxEndDateOk returns a tuple with the MaxEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetMaxEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxEndDate.Get(), o.MaxEndDate.IsSet()
}

// HasMaxEndDate returns a boolean if a field has been set.
func (o *GetProgramsDto) HasMaxEndDate() bool {
	if o != nil && o.MaxEndDate.IsSet() {
		return true
	}

	return false
}

// SetMaxEndDate gets a reference to the given NullableTime and assigns it to the MaxEndDate field.
func (o *GetProgramsDto) SetMaxEndDate(v time.Time) {
	o.MaxEndDate.Set(&v)
}
// SetMaxEndDateNil sets the value for MaxEndDate to be an explicit nil
func (o *GetProgramsDto) SetMaxEndDateNil() {
	o.MaxEndDate.Set(nil)
}

// UnsetMaxEndDate ensures that no value is present for MaxEndDate, not even an explicit nil
func (o *GetProgramsDto) UnsetMaxEndDate() {
	o.MaxEndDate.Unset()
}

// GetIsMovie returns the IsMovie field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetIsMovie() bool {
	if o == nil || IsNil(o.IsMovie.Get()) {
		var ret bool
		return ret
	}
	return *o.IsMovie.Get()
}

// GetIsMovieOk returns a tuple with the IsMovie field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetIsMovieOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsMovie.Get(), o.IsMovie.IsSet()
}

// HasIsMovie returns a boolean if a field has been set.
func (o *GetProgramsDto) HasIsMovie() bool {
	if o != nil && o.IsMovie.IsSet() {
		return true
	}

	return false
}

// SetIsMovie gets a reference to the given NullableBool and assigns it to the IsMovie field.
func (o *GetProgramsDto) SetIsMovie(v bool) {
	o.IsMovie.Set(&v)
}
// SetIsMovieNil sets the value for IsMovie to be an explicit nil
func (o *GetProgramsDto) SetIsMovieNil() {
	o.IsMovie.Set(nil)
}

// UnsetIsMovie ensures that no value is present for IsMovie, not even an explicit nil
func (o *GetProgramsDto) UnsetIsMovie() {
	o.IsMovie.Unset()
}

// GetIsSeries returns the IsSeries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetIsSeries() bool {
	if o == nil || IsNil(o.IsSeries.Get()) {
		var ret bool
		return ret
	}
	return *o.IsSeries.Get()
}

// GetIsSeriesOk returns a tuple with the IsSeries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetIsSeriesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsSeries.Get(), o.IsSeries.IsSet()
}

// HasIsSeries returns a boolean if a field has been set.
func (o *GetProgramsDto) HasIsSeries() bool {
	if o != nil && o.IsSeries.IsSet() {
		return true
	}

	return false
}

// SetIsSeries gets a reference to the given NullableBool and assigns it to the IsSeries field.
func (o *GetProgramsDto) SetIsSeries(v bool) {
	o.IsSeries.Set(&v)
}
// SetIsSeriesNil sets the value for IsSeries to be an explicit nil
func (o *GetProgramsDto) SetIsSeriesNil() {
	o.IsSeries.Set(nil)
}

// UnsetIsSeries ensures that no value is present for IsSeries, not even an explicit nil
func (o *GetProgramsDto) UnsetIsSeries() {
	o.IsSeries.Unset()
}

// GetIsNews returns the IsNews field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetIsNews() bool {
	if o == nil || IsNil(o.IsNews.Get()) {
		var ret bool
		return ret
	}
	return *o.IsNews.Get()
}

// GetIsNewsOk returns a tuple with the IsNews field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetIsNewsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsNews.Get(), o.IsNews.IsSet()
}

// HasIsNews returns a boolean if a field has been set.
func (o *GetProgramsDto) HasIsNews() bool {
	if o != nil && o.IsNews.IsSet() {
		return true
	}

	return false
}

// SetIsNews gets a reference to the given NullableBool and assigns it to the IsNews field.
func (o *GetProgramsDto) SetIsNews(v bool) {
	o.IsNews.Set(&v)
}
// SetIsNewsNil sets the value for IsNews to be an explicit nil
func (o *GetProgramsDto) SetIsNewsNil() {
	o.IsNews.Set(nil)
}

// UnsetIsNews ensures that no value is present for IsNews, not even an explicit nil
func (o *GetProgramsDto) UnsetIsNews() {
	o.IsNews.Unset()
}

// GetIsKids returns the IsKids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetIsKids() bool {
	if o == nil || IsNil(o.IsKids.Get()) {
		var ret bool
		return ret
	}
	return *o.IsKids.Get()
}

// GetIsKidsOk returns a tuple with the IsKids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetIsKidsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsKids.Get(), o.IsKids.IsSet()
}

// HasIsKids returns a boolean if a field has been set.
func (o *GetProgramsDto) HasIsKids() bool {
	if o != nil && o.IsKids.IsSet() {
		return true
	}

	return false
}

// SetIsKids gets a reference to the given NullableBool and assigns it to the IsKids field.
func (o *GetProgramsDto) SetIsKids(v bool) {
	o.IsKids.Set(&v)
}
// SetIsKidsNil sets the value for IsKids to be an explicit nil
func (o *GetProgramsDto) SetIsKidsNil() {
	o.IsKids.Set(nil)
}

// UnsetIsKids ensures that no value is present for IsKids, not even an explicit nil
func (o *GetProgramsDto) UnsetIsKids() {
	o.IsKids.Unset()
}

// GetIsSports returns the IsSports field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetIsSports() bool {
	if o == nil || IsNil(o.IsSports.Get()) {
		var ret bool
		return ret
	}
	return *o.IsSports.Get()
}

// GetIsSportsOk returns a tuple with the IsSports field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetIsSportsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsSports.Get(), o.IsSports.IsSet()
}

// HasIsSports returns a boolean if a field has been set.
func (o *GetProgramsDto) HasIsSports() bool {
	if o != nil && o.IsSports.IsSet() {
		return true
	}

	return false
}

// SetIsSports gets a reference to the given NullableBool and assigns it to the IsSports field.
func (o *GetProgramsDto) SetIsSports(v bool) {
	o.IsSports.Set(&v)
}
// SetIsSportsNil sets the value for IsSports to be an explicit nil
func (o *GetProgramsDto) SetIsSportsNil() {
	o.IsSports.Set(nil)
}

// UnsetIsSports ensures that no value is present for IsSports, not even an explicit nil
func (o *GetProgramsDto) UnsetIsSports() {
	o.IsSports.Unset()
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetStartIndex() int32 {
	if o == nil || IsNil(o.StartIndex.Get()) {
		var ret int32
		return ret
	}
	return *o.StartIndex.Get()
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetStartIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartIndex.Get(), o.StartIndex.IsSet()
}

// HasStartIndex returns a boolean if a field has been set.
func (o *GetProgramsDto) HasStartIndex() bool {
	if o != nil && o.StartIndex.IsSet() {
		return true
	}

	return false
}

// SetStartIndex gets a reference to the given NullableInt32 and assigns it to the StartIndex field.
func (o *GetProgramsDto) SetStartIndex(v int32) {
	o.StartIndex.Set(&v)
}
// SetStartIndexNil sets the value for StartIndex to be an explicit nil
func (o *GetProgramsDto) SetStartIndexNil() {
	o.StartIndex.Set(nil)
}

// UnsetStartIndex ensures that no value is present for StartIndex, not even an explicit nil
func (o *GetProgramsDto) UnsetStartIndex() {
	o.StartIndex.Unset()
}

// GetLimit returns the Limit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetLimit() int32 {
	if o == nil || IsNil(o.Limit.Get()) {
		var ret int32
		return ret
	}
	return *o.Limit.Get()
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Limit.Get(), o.Limit.IsSet()
}

// HasLimit returns a boolean if a field has been set.
func (o *GetProgramsDto) HasLimit() bool {
	if o != nil && o.Limit.IsSet() {
		return true
	}

	return false
}

// SetLimit gets a reference to the given NullableInt32 and assigns it to the Limit field.
func (o *GetProgramsDto) SetLimit(v int32) {
	o.Limit.Set(&v)
}
// SetLimitNil sets the value for Limit to be an explicit nil
func (o *GetProgramsDto) SetLimitNil() {
	o.Limit.Set(nil)
}

// UnsetLimit ensures that no value is present for Limit, not even an explicit nil
func (o *GetProgramsDto) UnsetLimit() {
	o.Limit.Unset()
}

// GetSortBy returns the SortBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetSortBy() []ItemSortBy {
	if o == nil {
		var ret []ItemSortBy
		return ret
	}
	return o.SortBy
}

// GetSortByOk returns a tuple with the SortBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetSortByOk() ([]ItemSortBy, bool) {
	if o == nil || IsNil(o.SortBy) {
		return nil, false
	}
	return o.SortBy, true
}

// HasSortBy returns a boolean if a field has been set.
func (o *GetProgramsDto) HasSortBy() bool {
	if o != nil && !IsNil(o.SortBy) {
		return true
	}

	return false
}

// SetSortBy gets a reference to the given []ItemSortBy and assigns it to the SortBy field.
func (o *GetProgramsDto) SetSortBy(v []ItemSortBy) {
	o.SortBy = v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetSortOrder() []SortOrder {
	if o == nil {
		var ret []SortOrder
		return ret
	}
	return o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetSortOrderOk() ([]SortOrder, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *GetProgramsDto) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given []SortOrder and assigns it to the SortOrder field.
func (o *GetProgramsDto) SetSortOrder(v []SortOrder) {
	o.SortOrder = v
}

// GetGenres returns the Genres field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetGenres() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetGenresOk() ([]string, bool) {
	if o == nil || IsNil(o.Genres) {
		return nil, false
	}
	return o.Genres, true
}

// HasGenres returns a boolean if a field has been set.
func (o *GetProgramsDto) HasGenres() bool {
	if o != nil && !IsNil(o.Genres) {
		return true
	}

	return false
}

// SetGenres gets a reference to the given []string and assigns it to the Genres field.
func (o *GetProgramsDto) SetGenres(v []string) {
	o.Genres = v
}

// GetGenreIds returns the GenreIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetGenreIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.GenreIds
}

// GetGenreIdsOk returns a tuple with the GenreIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetGenreIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.GenreIds) {
		return nil, false
	}
	return o.GenreIds, true
}

// HasGenreIds returns a boolean if a field has been set.
func (o *GetProgramsDto) HasGenreIds() bool {
	if o != nil && !IsNil(o.GenreIds) {
		return true
	}

	return false
}

// SetGenreIds gets a reference to the given []string and assigns it to the GenreIds field.
func (o *GetProgramsDto) SetGenreIds(v []string) {
	o.GenreIds = v
}

// GetEnableImages returns the EnableImages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetEnableImages() bool {
	if o == nil || IsNil(o.EnableImages.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableImages.Get()
}

// GetEnableImagesOk returns a tuple with the EnableImages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetEnableImagesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableImages.Get(), o.EnableImages.IsSet()
}

// HasEnableImages returns a boolean if a field has been set.
func (o *GetProgramsDto) HasEnableImages() bool {
	if o != nil && o.EnableImages.IsSet() {
		return true
	}

	return false
}

// SetEnableImages gets a reference to the given NullableBool and assigns it to the EnableImages field.
func (o *GetProgramsDto) SetEnableImages(v bool) {
	o.EnableImages.Set(&v)
}
// SetEnableImagesNil sets the value for EnableImages to be an explicit nil
func (o *GetProgramsDto) SetEnableImagesNil() {
	o.EnableImages.Set(nil)
}

// UnsetEnableImages ensures that no value is present for EnableImages, not even an explicit nil
func (o *GetProgramsDto) UnsetEnableImages() {
	o.EnableImages.Unset()
}

// GetEnableTotalRecordCount returns the EnableTotalRecordCount field value if set, zero value otherwise.
func (o *GetProgramsDto) GetEnableTotalRecordCount() bool {
	if o == nil || IsNil(o.EnableTotalRecordCount) {
		var ret bool
		return ret
	}
	return *o.EnableTotalRecordCount
}

// GetEnableTotalRecordCountOk returns a tuple with the EnableTotalRecordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProgramsDto) GetEnableTotalRecordCountOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableTotalRecordCount) {
		return nil, false
	}
	return o.EnableTotalRecordCount, true
}

// HasEnableTotalRecordCount returns a boolean if a field has been set.
func (o *GetProgramsDto) HasEnableTotalRecordCount() bool {
	if o != nil && !IsNil(o.EnableTotalRecordCount) {
		return true
	}

	return false
}

// SetEnableTotalRecordCount gets a reference to the given bool and assigns it to the EnableTotalRecordCount field.
func (o *GetProgramsDto) SetEnableTotalRecordCount(v bool) {
	o.EnableTotalRecordCount = &v
}

// GetImageTypeLimit returns the ImageTypeLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetImageTypeLimit() int32 {
	if o == nil || IsNil(o.ImageTypeLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.ImageTypeLimit.Get()
}

// GetImageTypeLimitOk returns a tuple with the ImageTypeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetImageTypeLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageTypeLimit.Get(), o.ImageTypeLimit.IsSet()
}

// HasImageTypeLimit returns a boolean if a field has been set.
func (o *GetProgramsDto) HasImageTypeLimit() bool {
	if o != nil && o.ImageTypeLimit.IsSet() {
		return true
	}

	return false
}

// SetImageTypeLimit gets a reference to the given NullableInt32 and assigns it to the ImageTypeLimit field.
func (o *GetProgramsDto) SetImageTypeLimit(v int32) {
	o.ImageTypeLimit.Set(&v)
}
// SetImageTypeLimitNil sets the value for ImageTypeLimit to be an explicit nil
func (o *GetProgramsDto) SetImageTypeLimitNil() {
	o.ImageTypeLimit.Set(nil)
}

// UnsetImageTypeLimit ensures that no value is present for ImageTypeLimit, not even an explicit nil
func (o *GetProgramsDto) UnsetImageTypeLimit() {
	o.ImageTypeLimit.Unset()
}

// GetEnableImageTypes returns the EnableImageTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetEnableImageTypes() []ImageType {
	if o == nil {
		var ret []ImageType
		return ret
	}
	return o.EnableImageTypes
}

// GetEnableImageTypesOk returns a tuple with the EnableImageTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetEnableImageTypesOk() ([]ImageType, bool) {
	if o == nil || IsNil(o.EnableImageTypes) {
		return nil, false
	}
	return o.EnableImageTypes, true
}

// HasEnableImageTypes returns a boolean if a field has been set.
func (o *GetProgramsDto) HasEnableImageTypes() bool {
	if o != nil && !IsNil(o.EnableImageTypes) {
		return true
	}

	return false
}

// SetEnableImageTypes gets a reference to the given []ImageType and assigns it to the EnableImageTypes field.
func (o *GetProgramsDto) SetEnableImageTypes(v []ImageType) {
	o.EnableImageTypes = v
}

// GetEnableUserData returns the EnableUserData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetEnableUserData() bool {
	if o == nil || IsNil(o.EnableUserData.Get()) {
		var ret bool
		return ret
	}
	return *o.EnableUserData.Get()
}

// GetEnableUserDataOk returns a tuple with the EnableUserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetEnableUserDataOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnableUserData.Get(), o.EnableUserData.IsSet()
}

// HasEnableUserData returns a boolean if a field has been set.
func (o *GetProgramsDto) HasEnableUserData() bool {
	if o != nil && o.EnableUserData.IsSet() {
		return true
	}

	return false
}

// SetEnableUserData gets a reference to the given NullableBool and assigns it to the EnableUserData field.
func (o *GetProgramsDto) SetEnableUserData(v bool) {
	o.EnableUserData.Set(&v)
}
// SetEnableUserDataNil sets the value for EnableUserData to be an explicit nil
func (o *GetProgramsDto) SetEnableUserDataNil() {
	o.EnableUserData.Set(nil)
}

// UnsetEnableUserData ensures that no value is present for EnableUserData, not even an explicit nil
func (o *GetProgramsDto) UnsetEnableUserData() {
	o.EnableUserData.Unset()
}

// GetSeriesTimerId returns the SeriesTimerId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetSeriesTimerId() string {
	if o == nil || IsNil(o.SeriesTimerId.Get()) {
		var ret string
		return ret
	}
	return *o.SeriesTimerId.Get()
}

// GetSeriesTimerIdOk returns a tuple with the SeriesTimerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetSeriesTimerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SeriesTimerId.Get(), o.SeriesTimerId.IsSet()
}

// HasSeriesTimerId returns a boolean if a field has been set.
func (o *GetProgramsDto) HasSeriesTimerId() bool {
	if o != nil && o.SeriesTimerId.IsSet() {
		return true
	}

	return false
}

// SetSeriesTimerId gets a reference to the given NullableString and assigns it to the SeriesTimerId field.
func (o *GetProgramsDto) SetSeriesTimerId(v string) {
	o.SeriesTimerId.Set(&v)
}
// SetSeriesTimerIdNil sets the value for SeriesTimerId to be an explicit nil
func (o *GetProgramsDto) SetSeriesTimerIdNil() {
	o.SeriesTimerId.Set(nil)
}

// UnsetSeriesTimerId ensures that no value is present for SeriesTimerId, not even an explicit nil
func (o *GetProgramsDto) UnsetSeriesTimerId() {
	o.SeriesTimerId.Unset()
}

// GetLibrarySeriesId returns the LibrarySeriesId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetLibrarySeriesId() string {
	if o == nil || IsNil(o.LibrarySeriesId.Get()) {
		var ret string
		return ret
	}
	return *o.LibrarySeriesId.Get()
}

// GetLibrarySeriesIdOk returns a tuple with the LibrarySeriesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetLibrarySeriesIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LibrarySeriesId.Get(), o.LibrarySeriesId.IsSet()
}

// HasLibrarySeriesId returns a boolean if a field has been set.
func (o *GetProgramsDto) HasLibrarySeriesId() bool {
	if o != nil && o.LibrarySeriesId.IsSet() {
		return true
	}

	return false
}

// SetLibrarySeriesId gets a reference to the given NullableString and assigns it to the LibrarySeriesId field.
func (o *GetProgramsDto) SetLibrarySeriesId(v string) {
	o.LibrarySeriesId.Set(&v)
}
// SetLibrarySeriesIdNil sets the value for LibrarySeriesId to be an explicit nil
func (o *GetProgramsDto) SetLibrarySeriesIdNil() {
	o.LibrarySeriesId.Set(nil)
}

// UnsetLibrarySeriesId ensures that no value is present for LibrarySeriesId, not even an explicit nil
func (o *GetProgramsDto) UnsetLibrarySeriesId() {
	o.LibrarySeriesId.Unset()
}

// GetFields returns the Fields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProgramsDto) GetFields() []ItemFields {
	if o == nil {
		var ret []ItemFields
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProgramsDto) GetFieldsOk() ([]ItemFields, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *GetProgramsDto) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given []ItemFields and assigns it to the Fields field.
func (o *GetProgramsDto) SetFields(v []ItemFields) {
	o.Fields = v
}

func (o GetProgramsDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProgramsDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ChannelIds != nil {
		toSerialize["ChannelIds"] = o.ChannelIds
	}
	if o.UserId.IsSet() {
		toSerialize["UserId"] = o.UserId.Get()
	}
	if o.MinStartDate.IsSet() {
		toSerialize["MinStartDate"] = o.MinStartDate.Get()
	}
	if o.HasAired.IsSet() {
		toSerialize["HasAired"] = o.HasAired.Get()
	}
	if o.IsAiring.IsSet() {
		toSerialize["IsAiring"] = o.IsAiring.Get()
	}
	if o.MaxStartDate.IsSet() {
		toSerialize["MaxStartDate"] = o.MaxStartDate.Get()
	}
	if o.MinEndDate.IsSet() {
		toSerialize["MinEndDate"] = o.MinEndDate.Get()
	}
	if o.MaxEndDate.IsSet() {
		toSerialize["MaxEndDate"] = o.MaxEndDate.Get()
	}
	if o.IsMovie.IsSet() {
		toSerialize["IsMovie"] = o.IsMovie.Get()
	}
	if o.IsSeries.IsSet() {
		toSerialize["IsSeries"] = o.IsSeries.Get()
	}
	if o.IsNews.IsSet() {
		toSerialize["IsNews"] = o.IsNews.Get()
	}
	if o.IsKids.IsSet() {
		toSerialize["IsKids"] = o.IsKids.Get()
	}
	if o.IsSports.IsSet() {
		toSerialize["IsSports"] = o.IsSports.Get()
	}
	if o.StartIndex.IsSet() {
		toSerialize["StartIndex"] = o.StartIndex.Get()
	}
	if o.Limit.IsSet() {
		toSerialize["Limit"] = o.Limit.Get()
	}
	if o.SortBy != nil {
		toSerialize["SortBy"] = o.SortBy
	}
	if o.SortOrder != nil {
		toSerialize["SortOrder"] = o.SortOrder
	}
	if o.Genres != nil {
		toSerialize["Genres"] = o.Genres
	}
	if o.GenreIds != nil {
		toSerialize["GenreIds"] = o.GenreIds
	}
	if o.EnableImages.IsSet() {
		toSerialize["EnableImages"] = o.EnableImages.Get()
	}
	if !IsNil(o.EnableTotalRecordCount) {
		toSerialize["EnableTotalRecordCount"] = o.EnableTotalRecordCount
	}
	if o.ImageTypeLimit.IsSet() {
		toSerialize["ImageTypeLimit"] = o.ImageTypeLimit.Get()
	}
	if o.EnableImageTypes != nil {
		toSerialize["EnableImageTypes"] = o.EnableImageTypes
	}
	if o.EnableUserData.IsSet() {
		toSerialize["EnableUserData"] = o.EnableUserData.Get()
	}
	if o.SeriesTimerId.IsSet() {
		toSerialize["SeriesTimerId"] = o.SeriesTimerId.Get()
	}
	if o.LibrarySeriesId.IsSet() {
		toSerialize["LibrarySeriesId"] = o.LibrarySeriesId.Get()
	}
	if o.Fields != nil {
		toSerialize["Fields"] = o.Fields
	}
	return toSerialize, nil
}

type NullableGetProgramsDto struct {
	value *GetProgramsDto
	isSet bool
}

func (v NullableGetProgramsDto) Get() *GetProgramsDto {
	return v.value
}

func (v *NullableGetProgramsDto) Set(val *GetProgramsDto) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProgramsDto) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProgramsDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProgramsDto(val *GetProgramsDto) *NullableGetProgramsDto {
	return &NullableGetProgramsDto{value: val, isSet: true}
}

func (v NullableGetProgramsDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProgramsDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


