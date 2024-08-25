/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.10
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the ListingsProviderInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListingsProviderInfo{}

// ListingsProviderInfo struct for ListingsProviderInfo
type ListingsProviderInfo struct {
	Id NullableString `json:"Id,omitempty"`
	Type NullableString `json:"Type,omitempty"`
	Username NullableString `json:"Username,omitempty"`
	Password NullableString `json:"Password,omitempty"`
	ListingsId NullableString `json:"ListingsId,omitempty"`
	ZipCode NullableString `json:"ZipCode,omitempty"`
	Country NullableString `json:"Country,omitempty"`
	Path NullableString `json:"Path,omitempty"`
	EnabledTuners []string `json:"EnabledTuners,omitempty"`
	EnableAllTuners *bool `json:"EnableAllTuners,omitempty"`
	NewsCategories []string `json:"NewsCategories,omitempty"`
	SportsCategories []string `json:"SportsCategories,omitempty"`
	KidsCategories []string `json:"KidsCategories,omitempty"`
	MovieCategories []string `json:"MovieCategories,omitempty"`
	ChannelMappings []NameValuePair `json:"ChannelMappings,omitempty"`
	MoviePrefix NullableString `json:"MoviePrefix,omitempty"`
	PreferredLanguage NullableString `json:"PreferredLanguage,omitempty"`
	UserAgent NullableString `json:"UserAgent,omitempty"`
}

// NewListingsProviderInfo instantiates a new ListingsProviderInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListingsProviderInfo() *ListingsProviderInfo {
	this := ListingsProviderInfo{}
	return &this
}

// NewListingsProviderInfoWithDefaults instantiates a new ListingsProviderInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListingsProviderInfoWithDefaults() *ListingsProviderInfo {
	this := ListingsProviderInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingsProviderInfo) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingsProviderInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ListingsProviderInfo) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *ListingsProviderInfo) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ListingsProviderInfo) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ListingsProviderInfo) UnsetId() {
	o.Id.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingsProviderInfo) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingsProviderInfo) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *ListingsProviderInfo) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *ListingsProviderInfo) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *ListingsProviderInfo) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *ListingsProviderInfo) UnsetType() {
	o.Type.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingsProviderInfo) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingsProviderInfo) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *ListingsProviderInfo) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *ListingsProviderInfo) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *ListingsProviderInfo) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *ListingsProviderInfo) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingsProviderInfo) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingsProviderInfo) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *ListingsProviderInfo) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *ListingsProviderInfo) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *ListingsProviderInfo) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *ListingsProviderInfo) UnsetPassword() {
	o.Password.Unset()
}

// GetListingsId returns the ListingsId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingsProviderInfo) GetListingsId() string {
	if o == nil || IsNil(o.ListingsId.Get()) {
		var ret string
		return ret
	}
	return *o.ListingsId.Get()
}

// GetListingsIdOk returns a tuple with the ListingsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingsProviderInfo) GetListingsIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ListingsId.Get(), o.ListingsId.IsSet()
}

// HasListingsId returns a boolean if a field has been set.
func (o *ListingsProviderInfo) HasListingsId() bool {
	if o != nil && o.ListingsId.IsSet() {
		return true
	}

	return false
}

// SetListingsId gets a reference to the given NullableString and assigns it to the ListingsId field.
func (o *ListingsProviderInfo) SetListingsId(v string) {
	o.ListingsId.Set(&v)
}
// SetListingsIdNil sets the value for ListingsId to be an explicit nil
func (o *ListingsProviderInfo) SetListingsIdNil() {
	o.ListingsId.Set(nil)
}

// UnsetListingsId ensures that no value is present for ListingsId, not even an explicit nil
func (o *ListingsProviderInfo) UnsetListingsId() {
	o.ListingsId.Unset()
}

// GetZipCode returns the ZipCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingsProviderInfo) GetZipCode() string {
	if o == nil || IsNil(o.ZipCode.Get()) {
		var ret string
		return ret
	}
	return *o.ZipCode.Get()
}

// GetZipCodeOk returns a tuple with the ZipCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingsProviderInfo) GetZipCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ZipCode.Get(), o.ZipCode.IsSet()
}

// HasZipCode returns a boolean if a field has been set.
func (o *ListingsProviderInfo) HasZipCode() bool {
	if o != nil && o.ZipCode.IsSet() {
		return true
	}

	return false
}

// SetZipCode gets a reference to the given NullableString and assigns it to the ZipCode field.
func (o *ListingsProviderInfo) SetZipCode(v string) {
	o.ZipCode.Set(&v)
}
// SetZipCodeNil sets the value for ZipCode to be an explicit nil
func (o *ListingsProviderInfo) SetZipCodeNil() {
	o.ZipCode.Set(nil)
}

// UnsetZipCode ensures that no value is present for ZipCode, not even an explicit nil
func (o *ListingsProviderInfo) UnsetZipCode() {
	o.ZipCode.Unset()
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingsProviderInfo) GetCountry() string {
	if o == nil || IsNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingsProviderInfo) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *ListingsProviderInfo) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *ListingsProviderInfo) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *ListingsProviderInfo) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *ListingsProviderInfo) UnsetCountry() {
	o.Country.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingsProviderInfo) GetPath() string {
	if o == nil || IsNil(o.Path.Get()) {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingsProviderInfo) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *ListingsProviderInfo) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *ListingsProviderInfo) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *ListingsProviderInfo) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *ListingsProviderInfo) UnsetPath() {
	o.Path.Unset()
}

// GetEnabledTuners returns the EnabledTuners field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingsProviderInfo) GetEnabledTuners() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.EnabledTuners
}

// GetEnabledTunersOk returns a tuple with the EnabledTuners field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingsProviderInfo) GetEnabledTunersOk() ([]string, bool) {
	if o == nil || IsNil(o.EnabledTuners) {
		return nil, false
	}
	return o.EnabledTuners, true
}

// HasEnabledTuners returns a boolean if a field has been set.
func (o *ListingsProviderInfo) HasEnabledTuners() bool {
	if o != nil && !IsNil(o.EnabledTuners) {
		return true
	}

	return false
}

// SetEnabledTuners gets a reference to the given []string and assigns it to the EnabledTuners field.
func (o *ListingsProviderInfo) SetEnabledTuners(v []string) {
	o.EnabledTuners = v
}

// GetEnableAllTuners returns the EnableAllTuners field value if set, zero value otherwise.
func (o *ListingsProviderInfo) GetEnableAllTuners() bool {
	if o == nil || IsNil(o.EnableAllTuners) {
		var ret bool
		return ret
	}
	return *o.EnableAllTuners
}

// GetEnableAllTunersOk returns a tuple with the EnableAllTuners field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListingsProviderInfo) GetEnableAllTunersOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAllTuners) {
		return nil, false
	}
	return o.EnableAllTuners, true
}

// HasEnableAllTuners returns a boolean if a field has been set.
func (o *ListingsProviderInfo) HasEnableAllTuners() bool {
	if o != nil && !IsNil(o.EnableAllTuners) {
		return true
	}

	return false
}

// SetEnableAllTuners gets a reference to the given bool and assigns it to the EnableAllTuners field.
func (o *ListingsProviderInfo) SetEnableAllTuners(v bool) {
	o.EnableAllTuners = &v
}

// GetNewsCategories returns the NewsCategories field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingsProviderInfo) GetNewsCategories() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.NewsCategories
}

// GetNewsCategoriesOk returns a tuple with the NewsCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingsProviderInfo) GetNewsCategoriesOk() ([]string, bool) {
	if o == nil || IsNil(o.NewsCategories) {
		return nil, false
	}
	return o.NewsCategories, true
}

// HasNewsCategories returns a boolean if a field has been set.
func (o *ListingsProviderInfo) HasNewsCategories() bool {
	if o != nil && !IsNil(o.NewsCategories) {
		return true
	}

	return false
}

// SetNewsCategories gets a reference to the given []string and assigns it to the NewsCategories field.
func (o *ListingsProviderInfo) SetNewsCategories(v []string) {
	o.NewsCategories = v
}

// GetSportsCategories returns the SportsCategories field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingsProviderInfo) GetSportsCategories() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SportsCategories
}

// GetSportsCategoriesOk returns a tuple with the SportsCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingsProviderInfo) GetSportsCategoriesOk() ([]string, bool) {
	if o == nil || IsNil(o.SportsCategories) {
		return nil, false
	}
	return o.SportsCategories, true
}

// HasSportsCategories returns a boolean if a field has been set.
func (o *ListingsProviderInfo) HasSportsCategories() bool {
	if o != nil && !IsNil(o.SportsCategories) {
		return true
	}

	return false
}

// SetSportsCategories gets a reference to the given []string and assigns it to the SportsCategories field.
func (o *ListingsProviderInfo) SetSportsCategories(v []string) {
	o.SportsCategories = v
}

// GetKidsCategories returns the KidsCategories field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingsProviderInfo) GetKidsCategories() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.KidsCategories
}

// GetKidsCategoriesOk returns a tuple with the KidsCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingsProviderInfo) GetKidsCategoriesOk() ([]string, bool) {
	if o == nil || IsNil(o.KidsCategories) {
		return nil, false
	}
	return o.KidsCategories, true
}

// HasKidsCategories returns a boolean if a field has been set.
func (o *ListingsProviderInfo) HasKidsCategories() bool {
	if o != nil && !IsNil(o.KidsCategories) {
		return true
	}

	return false
}

// SetKidsCategories gets a reference to the given []string and assigns it to the KidsCategories field.
func (o *ListingsProviderInfo) SetKidsCategories(v []string) {
	o.KidsCategories = v
}

// GetMovieCategories returns the MovieCategories field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingsProviderInfo) GetMovieCategories() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MovieCategories
}

// GetMovieCategoriesOk returns a tuple with the MovieCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingsProviderInfo) GetMovieCategoriesOk() ([]string, bool) {
	if o == nil || IsNil(o.MovieCategories) {
		return nil, false
	}
	return o.MovieCategories, true
}

// HasMovieCategories returns a boolean if a field has been set.
func (o *ListingsProviderInfo) HasMovieCategories() bool {
	if o != nil && !IsNil(o.MovieCategories) {
		return true
	}

	return false
}

// SetMovieCategories gets a reference to the given []string and assigns it to the MovieCategories field.
func (o *ListingsProviderInfo) SetMovieCategories(v []string) {
	o.MovieCategories = v
}

// GetChannelMappings returns the ChannelMappings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingsProviderInfo) GetChannelMappings() []NameValuePair {
	if o == nil {
		var ret []NameValuePair
		return ret
	}
	return o.ChannelMappings
}

// GetChannelMappingsOk returns a tuple with the ChannelMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingsProviderInfo) GetChannelMappingsOk() ([]NameValuePair, bool) {
	if o == nil || IsNil(o.ChannelMappings) {
		return nil, false
	}
	return o.ChannelMappings, true
}

// HasChannelMappings returns a boolean if a field has been set.
func (o *ListingsProviderInfo) HasChannelMappings() bool {
	if o != nil && !IsNil(o.ChannelMappings) {
		return true
	}

	return false
}

// SetChannelMappings gets a reference to the given []NameValuePair and assigns it to the ChannelMappings field.
func (o *ListingsProviderInfo) SetChannelMappings(v []NameValuePair) {
	o.ChannelMappings = v
}

// GetMoviePrefix returns the MoviePrefix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingsProviderInfo) GetMoviePrefix() string {
	if o == nil || IsNil(o.MoviePrefix.Get()) {
		var ret string
		return ret
	}
	return *o.MoviePrefix.Get()
}

// GetMoviePrefixOk returns a tuple with the MoviePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingsProviderInfo) GetMoviePrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MoviePrefix.Get(), o.MoviePrefix.IsSet()
}

// HasMoviePrefix returns a boolean if a field has been set.
func (o *ListingsProviderInfo) HasMoviePrefix() bool {
	if o != nil && o.MoviePrefix.IsSet() {
		return true
	}

	return false
}

// SetMoviePrefix gets a reference to the given NullableString and assigns it to the MoviePrefix field.
func (o *ListingsProviderInfo) SetMoviePrefix(v string) {
	o.MoviePrefix.Set(&v)
}
// SetMoviePrefixNil sets the value for MoviePrefix to be an explicit nil
func (o *ListingsProviderInfo) SetMoviePrefixNil() {
	o.MoviePrefix.Set(nil)
}

// UnsetMoviePrefix ensures that no value is present for MoviePrefix, not even an explicit nil
func (o *ListingsProviderInfo) UnsetMoviePrefix() {
	o.MoviePrefix.Unset()
}

// GetPreferredLanguage returns the PreferredLanguage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingsProviderInfo) GetPreferredLanguage() string {
	if o == nil || IsNil(o.PreferredLanguage.Get()) {
		var ret string
		return ret
	}
	return *o.PreferredLanguage.Get()
}

// GetPreferredLanguageOk returns a tuple with the PreferredLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingsProviderInfo) GetPreferredLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreferredLanguage.Get(), o.PreferredLanguage.IsSet()
}

// HasPreferredLanguage returns a boolean if a field has been set.
func (o *ListingsProviderInfo) HasPreferredLanguage() bool {
	if o != nil && o.PreferredLanguage.IsSet() {
		return true
	}

	return false
}

// SetPreferredLanguage gets a reference to the given NullableString and assigns it to the PreferredLanguage field.
func (o *ListingsProviderInfo) SetPreferredLanguage(v string) {
	o.PreferredLanguage.Set(&v)
}
// SetPreferredLanguageNil sets the value for PreferredLanguage to be an explicit nil
func (o *ListingsProviderInfo) SetPreferredLanguageNil() {
	o.PreferredLanguage.Set(nil)
}

// UnsetPreferredLanguage ensures that no value is present for PreferredLanguage, not even an explicit nil
func (o *ListingsProviderInfo) UnsetPreferredLanguage() {
	o.PreferredLanguage.Unset()
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListingsProviderInfo) GetUserAgent() string {
	if o == nil || IsNil(o.UserAgent.Get()) {
		var ret string
		return ret
	}
	return *o.UserAgent.Get()
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListingsProviderInfo) GetUserAgentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserAgent.Get(), o.UserAgent.IsSet()
}

// HasUserAgent returns a boolean if a field has been set.
func (o *ListingsProviderInfo) HasUserAgent() bool {
	if o != nil && o.UserAgent.IsSet() {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given NullableString and assigns it to the UserAgent field.
func (o *ListingsProviderInfo) SetUserAgent(v string) {
	o.UserAgent.Set(&v)
}
// SetUserAgentNil sets the value for UserAgent to be an explicit nil
func (o *ListingsProviderInfo) SetUserAgentNil() {
	o.UserAgent.Set(nil)
}

// UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil
func (o *ListingsProviderInfo) UnsetUserAgent() {
	o.UserAgent.Unset()
}

func (o ListingsProviderInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListingsProviderInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if o.Type.IsSet() {
		toSerialize["Type"] = o.Type.Get()
	}
	if o.Username.IsSet() {
		toSerialize["Username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["Password"] = o.Password.Get()
	}
	if o.ListingsId.IsSet() {
		toSerialize["ListingsId"] = o.ListingsId.Get()
	}
	if o.ZipCode.IsSet() {
		toSerialize["ZipCode"] = o.ZipCode.Get()
	}
	if o.Country.IsSet() {
		toSerialize["Country"] = o.Country.Get()
	}
	if o.Path.IsSet() {
		toSerialize["Path"] = o.Path.Get()
	}
	if o.EnabledTuners != nil {
		toSerialize["EnabledTuners"] = o.EnabledTuners
	}
	if !IsNil(o.EnableAllTuners) {
		toSerialize["EnableAllTuners"] = o.EnableAllTuners
	}
	if o.NewsCategories != nil {
		toSerialize["NewsCategories"] = o.NewsCategories
	}
	if o.SportsCategories != nil {
		toSerialize["SportsCategories"] = o.SportsCategories
	}
	if o.KidsCategories != nil {
		toSerialize["KidsCategories"] = o.KidsCategories
	}
	if o.MovieCategories != nil {
		toSerialize["MovieCategories"] = o.MovieCategories
	}
	if o.ChannelMappings != nil {
		toSerialize["ChannelMappings"] = o.ChannelMappings
	}
	if o.MoviePrefix.IsSet() {
		toSerialize["MoviePrefix"] = o.MoviePrefix.Get()
	}
	if o.PreferredLanguage.IsSet() {
		toSerialize["PreferredLanguage"] = o.PreferredLanguage.Get()
	}
	if o.UserAgent.IsSet() {
		toSerialize["UserAgent"] = o.UserAgent.Get()
	}
	return toSerialize, nil
}

type NullableListingsProviderInfo struct {
	value *ListingsProviderInfo
	isSet bool
}

func (v NullableListingsProviderInfo) Get() *ListingsProviderInfo {
	return v.value
}

func (v *NullableListingsProviderInfo) Set(val *ListingsProviderInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableListingsProviderInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableListingsProviderInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListingsProviderInfo(val *ListingsProviderInfo) *NullableListingsProviderInfo {
	return &NullableListingsProviderInfo{value: val, isSet: true}
}

func (v NullableListingsProviderInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListingsProviderInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


