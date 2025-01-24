/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the RemoteSubtitleInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoteSubtitleInfo{}

// RemoteSubtitleInfo struct for RemoteSubtitleInfo
type RemoteSubtitleInfo struct {
	ThreeLetterISOLanguageName NullableString `json:"ThreeLetterISOLanguageName,omitempty"`
	Id NullableString `json:"Id,omitempty"`
	ProviderName NullableString `json:"ProviderName,omitempty"`
	Name NullableString `json:"Name,omitempty"`
	Format NullableString `json:"Format,omitempty"`
	Author NullableString `json:"Author,omitempty"`
	Comment NullableString `json:"Comment,omitempty"`
	DateCreated NullableTime `json:"DateCreated,omitempty"`
	CommunityRating NullableFloat32 `json:"CommunityRating,omitempty"`
	FrameRate NullableFloat32 `json:"FrameRate,omitempty"`
	DownloadCount NullableInt32 `json:"DownloadCount,omitempty"`
	IsHashMatch NullableBool `json:"IsHashMatch,omitempty"`
	AiTranslated NullableBool `json:"AiTranslated,omitempty"`
	MachineTranslated NullableBool `json:"MachineTranslated,omitempty"`
	Forced NullableBool `json:"Forced,omitempty"`
	HearingImpaired NullableBool `json:"HearingImpaired,omitempty"`
}

// NewRemoteSubtitleInfo instantiates a new RemoteSubtitleInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoteSubtitleInfo() *RemoteSubtitleInfo {
	this := RemoteSubtitleInfo{}
	return &this
}

// NewRemoteSubtitleInfoWithDefaults instantiates a new RemoteSubtitleInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoteSubtitleInfoWithDefaults() *RemoteSubtitleInfo {
	this := RemoteSubtitleInfo{}
	return &this
}

// GetThreeLetterISOLanguageName returns the ThreeLetterISOLanguageName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSubtitleInfo) GetThreeLetterISOLanguageName() string {
	if o == nil || IsNil(o.ThreeLetterISOLanguageName.Get()) {
		var ret string
		return ret
	}
	return *o.ThreeLetterISOLanguageName.Get()
}

// GetThreeLetterISOLanguageNameOk returns a tuple with the ThreeLetterISOLanguageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSubtitleInfo) GetThreeLetterISOLanguageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThreeLetterISOLanguageName.Get(), o.ThreeLetterISOLanguageName.IsSet()
}

// HasThreeLetterISOLanguageName returns a boolean if a field has been set.
func (o *RemoteSubtitleInfo) HasThreeLetterISOLanguageName() bool {
	if o != nil && o.ThreeLetterISOLanguageName.IsSet() {
		return true
	}

	return false
}

// SetThreeLetterISOLanguageName gets a reference to the given NullableString and assigns it to the ThreeLetterISOLanguageName field.
func (o *RemoteSubtitleInfo) SetThreeLetterISOLanguageName(v string) {
	o.ThreeLetterISOLanguageName.Set(&v)
}
// SetThreeLetterISOLanguageNameNil sets the value for ThreeLetterISOLanguageName to be an explicit nil
func (o *RemoteSubtitleInfo) SetThreeLetterISOLanguageNameNil() {
	o.ThreeLetterISOLanguageName.Set(nil)
}

// UnsetThreeLetterISOLanguageName ensures that no value is present for ThreeLetterISOLanguageName, not even an explicit nil
func (o *RemoteSubtitleInfo) UnsetThreeLetterISOLanguageName() {
	o.ThreeLetterISOLanguageName.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSubtitleInfo) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSubtitleInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *RemoteSubtitleInfo) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *RemoteSubtitleInfo) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *RemoteSubtitleInfo) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *RemoteSubtitleInfo) UnsetId() {
	o.Id.Unset()
}

// GetProviderName returns the ProviderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSubtitleInfo) GetProviderName() string {
	if o == nil || IsNil(o.ProviderName.Get()) {
		var ret string
		return ret
	}
	return *o.ProviderName.Get()
}

// GetProviderNameOk returns a tuple with the ProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSubtitleInfo) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderName.Get(), o.ProviderName.IsSet()
}

// HasProviderName returns a boolean if a field has been set.
func (o *RemoteSubtitleInfo) HasProviderName() bool {
	if o != nil && o.ProviderName.IsSet() {
		return true
	}

	return false
}

// SetProviderName gets a reference to the given NullableString and assigns it to the ProviderName field.
func (o *RemoteSubtitleInfo) SetProviderName(v string) {
	o.ProviderName.Set(&v)
}
// SetProviderNameNil sets the value for ProviderName to be an explicit nil
func (o *RemoteSubtitleInfo) SetProviderNameNil() {
	o.ProviderName.Set(nil)
}

// UnsetProviderName ensures that no value is present for ProviderName, not even an explicit nil
func (o *RemoteSubtitleInfo) UnsetProviderName() {
	o.ProviderName.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSubtitleInfo) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSubtitleInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *RemoteSubtitleInfo) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *RemoteSubtitleInfo) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *RemoteSubtitleInfo) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *RemoteSubtitleInfo) UnsetName() {
	o.Name.Unset()
}

// GetFormat returns the Format field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSubtitleInfo) GetFormat() string {
	if o == nil || IsNil(o.Format.Get()) {
		var ret string
		return ret
	}
	return *o.Format.Get()
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSubtitleInfo) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Format.Get(), o.Format.IsSet()
}

// HasFormat returns a boolean if a field has been set.
func (o *RemoteSubtitleInfo) HasFormat() bool {
	if o != nil && o.Format.IsSet() {
		return true
	}

	return false
}

// SetFormat gets a reference to the given NullableString and assigns it to the Format field.
func (o *RemoteSubtitleInfo) SetFormat(v string) {
	o.Format.Set(&v)
}
// SetFormatNil sets the value for Format to be an explicit nil
func (o *RemoteSubtitleInfo) SetFormatNil() {
	o.Format.Set(nil)
}

// UnsetFormat ensures that no value is present for Format, not even an explicit nil
func (o *RemoteSubtitleInfo) UnsetFormat() {
	o.Format.Unset()
}

// GetAuthor returns the Author field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSubtitleInfo) GetAuthor() string {
	if o == nil || IsNil(o.Author.Get()) {
		var ret string
		return ret
	}
	return *o.Author.Get()
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSubtitleInfo) GetAuthorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Author.Get(), o.Author.IsSet()
}

// HasAuthor returns a boolean if a field has been set.
func (o *RemoteSubtitleInfo) HasAuthor() bool {
	if o != nil && o.Author.IsSet() {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given NullableString and assigns it to the Author field.
func (o *RemoteSubtitleInfo) SetAuthor(v string) {
	o.Author.Set(&v)
}
// SetAuthorNil sets the value for Author to be an explicit nil
func (o *RemoteSubtitleInfo) SetAuthorNil() {
	o.Author.Set(nil)
}

// UnsetAuthor ensures that no value is present for Author, not even an explicit nil
func (o *RemoteSubtitleInfo) UnsetAuthor() {
	o.Author.Unset()
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSubtitleInfo) GetComment() string {
	if o == nil || IsNil(o.Comment.Get()) {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSubtitleInfo) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *RemoteSubtitleInfo) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *RemoteSubtitleInfo) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *RemoteSubtitleInfo) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *RemoteSubtitleInfo) UnsetComment() {
	o.Comment.Unset()
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSubtitleInfo) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated.Get()
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSubtitleInfo) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateCreated.Get(), o.DateCreated.IsSet()
}

// HasDateCreated returns a boolean if a field has been set.
func (o *RemoteSubtitleInfo) HasDateCreated() bool {
	if o != nil && o.DateCreated.IsSet() {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given NullableTime and assigns it to the DateCreated field.
func (o *RemoteSubtitleInfo) SetDateCreated(v time.Time) {
	o.DateCreated.Set(&v)
}
// SetDateCreatedNil sets the value for DateCreated to be an explicit nil
func (o *RemoteSubtitleInfo) SetDateCreatedNil() {
	o.DateCreated.Set(nil)
}

// UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
func (o *RemoteSubtitleInfo) UnsetDateCreated() {
	o.DateCreated.Unset()
}

// GetCommunityRating returns the CommunityRating field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSubtitleInfo) GetCommunityRating() float32 {
	if o == nil || IsNil(o.CommunityRating.Get()) {
		var ret float32
		return ret
	}
	return *o.CommunityRating.Get()
}

// GetCommunityRatingOk returns a tuple with the CommunityRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSubtitleInfo) GetCommunityRatingOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CommunityRating.Get(), o.CommunityRating.IsSet()
}

// HasCommunityRating returns a boolean if a field has been set.
func (o *RemoteSubtitleInfo) HasCommunityRating() bool {
	if o != nil && o.CommunityRating.IsSet() {
		return true
	}

	return false
}

// SetCommunityRating gets a reference to the given NullableFloat32 and assigns it to the CommunityRating field.
func (o *RemoteSubtitleInfo) SetCommunityRating(v float32) {
	o.CommunityRating.Set(&v)
}
// SetCommunityRatingNil sets the value for CommunityRating to be an explicit nil
func (o *RemoteSubtitleInfo) SetCommunityRatingNil() {
	o.CommunityRating.Set(nil)
}

// UnsetCommunityRating ensures that no value is present for CommunityRating, not even an explicit nil
func (o *RemoteSubtitleInfo) UnsetCommunityRating() {
	o.CommunityRating.Unset()
}

// GetFrameRate returns the FrameRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSubtitleInfo) GetFrameRate() float32 {
	if o == nil || IsNil(o.FrameRate.Get()) {
		var ret float32
		return ret
	}
	return *o.FrameRate.Get()
}

// GetFrameRateOk returns a tuple with the FrameRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSubtitleInfo) GetFrameRateOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FrameRate.Get(), o.FrameRate.IsSet()
}

// HasFrameRate returns a boolean if a field has been set.
func (o *RemoteSubtitleInfo) HasFrameRate() bool {
	if o != nil && o.FrameRate.IsSet() {
		return true
	}

	return false
}

// SetFrameRate gets a reference to the given NullableFloat32 and assigns it to the FrameRate field.
func (o *RemoteSubtitleInfo) SetFrameRate(v float32) {
	o.FrameRate.Set(&v)
}
// SetFrameRateNil sets the value for FrameRate to be an explicit nil
func (o *RemoteSubtitleInfo) SetFrameRateNil() {
	o.FrameRate.Set(nil)
}

// UnsetFrameRate ensures that no value is present for FrameRate, not even an explicit nil
func (o *RemoteSubtitleInfo) UnsetFrameRate() {
	o.FrameRate.Unset()
}

// GetDownloadCount returns the DownloadCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSubtitleInfo) GetDownloadCount() int32 {
	if o == nil || IsNil(o.DownloadCount.Get()) {
		var ret int32
		return ret
	}
	return *o.DownloadCount.Get()
}

// GetDownloadCountOk returns a tuple with the DownloadCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSubtitleInfo) GetDownloadCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownloadCount.Get(), o.DownloadCount.IsSet()
}

// HasDownloadCount returns a boolean if a field has been set.
func (o *RemoteSubtitleInfo) HasDownloadCount() bool {
	if o != nil && o.DownloadCount.IsSet() {
		return true
	}

	return false
}

// SetDownloadCount gets a reference to the given NullableInt32 and assigns it to the DownloadCount field.
func (o *RemoteSubtitleInfo) SetDownloadCount(v int32) {
	o.DownloadCount.Set(&v)
}
// SetDownloadCountNil sets the value for DownloadCount to be an explicit nil
func (o *RemoteSubtitleInfo) SetDownloadCountNil() {
	o.DownloadCount.Set(nil)
}

// UnsetDownloadCount ensures that no value is present for DownloadCount, not even an explicit nil
func (o *RemoteSubtitleInfo) UnsetDownloadCount() {
	o.DownloadCount.Unset()
}

// GetIsHashMatch returns the IsHashMatch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSubtitleInfo) GetIsHashMatch() bool {
	if o == nil || IsNil(o.IsHashMatch.Get()) {
		var ret bool
		return ret
	}
	return *o.IsHashMatch.Get()
}

// GetIsHashMatchOk returns a tuple with the IsHashMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSubtitleInfo) GetIsHashMatchOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsHashMatch.Get(), o.IsHashMatch.IsSet()
}

// HasIsHashMatch returns a boolean if a field has been set.
func (o *RemoteSubtitleInfo) HasIsHashMatch() bool {
	if o != nil && o.IsHashMatch.IsSet() {
		return true
	}

	return false
}

// SetIsHashMatch gets a reference to the given NullableBool and assigns it to the IsHashMatch field.
func (o *RemoteSubtitleInfo) SetIsHashMatch(v bool) {
	o.IsHashMatch.Set(&v)
}
// SetIsHashMatchNil sets the value for IsHashMatch to be an explicit nil
func (o *RemoteSubtitleInfo) SetIsHashMatchNil() {
	o.IsHashMatch.Set(nil)
}

// UnsetIsHashMatch ensures that no value is present for IsHashMatch, not even an explicit nil
func (o *RemoteSubtitleInfo) UnsetIsHashMatch() {
	o.IsHashMatch.Unset()
}

// GetAiTranslated returns the AiTranslated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSubtitleInfo) GetAiTranslated() bool {
	if o == nil || IsNil(o.AiTranslated.Get()) {
		var ret bool
		return ret
	}
	return *o.AiTranslated.Get()
}

// GetAiTranslatedOk returns a tuple with the AiTranslated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSubtitleInfo) GetAiTranslatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AiTranslated.Get(), o.AiTranslated.IsSet()
}

// HasAiTranslated returns a boolean if a field has been set.
func (o *RemoteSubtitleInfo) HasAiTranslated() bool {
	if o != nil && o.AiTranslated.IsSet() {
		return true
	}

	return false
}

// SetAiTranslated gets a reference to the given NullableBool and assigns it to the AiTranslated field.
func (o *RemoteSubtitleInfo) SetAiTranslated(v bool) {
	o.AiTranslated.Set(&v)
}
// SetAiTranslatedNil sets the value for AiTranslated to be an explicit nil
func (o *RemoteSubtitleInfo) SetAiTranslatedNil() {
	o.AiTranslated.Set(nil)
}

// UnsetAiTranslated ensures that no value is present for AiTranslated, not even an explicit nil
func (o *RemoteSubtitleInfo) UnsetAiTranslated() {
	o.AiTranslated.Unset()
}

// GetMachineTranslated returns the MachineTranslated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSubtitleInfo) GetMachineTranslated() bool {
	if o == nil || IsNil(o.MachineTranslated.Get()) {
		var ret bool
		return ret
	}
	return *o.MachineTranslated.Get()
}

// GetMachineTranslatedOk returns a tuple with the MachineTranslated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSubtitleInfo) GetMachineTranslatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.MachineTranslated.Get(), o.MachineTranslated.IsSet()
}

// HasMachineTranslated returns a boolean if a field has been set.
func (o *RemoteSubtitleInfo) HasMachineTranslated() bool {
	if o != nil && o.MachineTranslated.IsSet() {
		return true
	}

	return false
}

// SetMachineTranslated gets a reference to the given NullableBool and assigns it to the MachineTranslated field.
func (o *RemoteSubtitleInfo) SetMachineTranslated(v bool) {
	o.MachineTranslated.Set(&v)
}
// SetMachineTranslatedNil sets the value for MachineTranslated to be an explicit nil
func (o *RemoteSubtitleInfo) SetMachineTranslatedNil() {
	o.MachineTranslated.Set(nil)
}

// UnsetMachineTranslated ensures that no value is present for MachineTranslated, not even an explicit nil
func (o *RemoteSubtitleInfo) UnsetMachineTranslated() {
	o.MachineTranslated.Unset()
}

// GetForced returns the Forced field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSubtitleInfo) GetForced() bool {
	if o == nil || IsNil(o.Forced.Get()) {
		var ret bool
		return ret
	}
	return *o.Forced.Get()
}

// GetForcedOk returns a tuple with the Forced field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSubtitleInfo) GetForcedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Forced.Get(), o.Forced.IsSet()
}

// HasForced returns a boolean if a field has been set.
func (o *RemoteSubtitleInfo) HasForced() bool {
	if o != nil && o.Forced.IsSet() {
		return true
	}

	return false
}

// SetForced gets a reference to the given NullableBool and assigns it to the Forced field.
func (o *RemoteSubtitleInfo) SetForced(v bool) {
	o.Forced.Set(&v)
}
// SetForcedNil sets the value for Forced to be an explicit nil
func (o *RemoteSubtitleInfo) SetForcedNil() {
	o.Forced.Set(nil)
}

// UnsetForced ensures that no value is present for Forced, not even an explicit nil
func (o *RemoteSubtitleInfo) UnsetForced() {
	o.Forced.Unset()
}

// GetHearingImpaired returns the HearingImpaired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RemoteSubtitleInfo) GetHearingImpaired() bool {
	if o == nil || IsNil(o.HearingImpaired.Get()) {
		var ret bool
		return ret
	}
	return *o.HearingImpaired.Get()
}

// GetHearingImpairedOk returns a tuple with the HearingImpaired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RemoteSubtitleInfo) GetHearingImpairedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HearingImpaired.Get(), o.HearingImpaired.IsSet()
}

// HasHearingImpaired returns a boolean if a field has been set.
func (o *RemoteSubtitleInfo) HasHearingImpaired() bool {
	if o != nil && o.HearingImpaired.IsSet() {
		return true
	}

	return false
}

// SetHearingImpaired gets a reference to the given NullableBool and assigns it to the HearingImpaired field.
func (o *RemoteSubtitleInfo) SetHearingImpaired(v bool) {
	o.HearingImpaired.Set(&v)
}
// SetHearingImpairedNil sets the value for HearingImpaired to be an explicit nil
func (o *RemoteSubtitleInfo) SetHearingImpairedNil() {
	o.HearingImpaired.Set(nil)
}

// UnsetHearingImpaired ensures that no value is present for HearingImpaired, not even an explicit nil
func (o *RemoteSubtitleInfo) UnsetHearingImpaired() {
	o.HearingImpaired.Unset()
}

func (o RemoteSubtitleInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoteSubtitleInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ThreeLetterISOLanguageName.IsSet() {
		toSerialize["ThreeLetterISOLanguageName"] = o.ThreeLetterISOLanguageName.Get()
	}
	if o.Id.IsSet() {
		toSerialize["Id"] = o.Id.Get()
	}
	if o.ProviderName.IsSet() {
		toSerialize["ProviderName"] = o.ProviderName.Get()
	}
	if o.Name.IsSet() {
		toSerialize["Name"] = o.Name.Get()
	}
	if o.Format.IsSet() {
		toSerialize["Format"] = o.Format.Get()
	}
	if o.Author.IsSet() {
		toSerialize["Author"] = o.Author.Get()
	}
	if o.Comment.IsSet() {
		toSerialize["Comment"] = o.Comment.Get()
	}
	if o.DateCreated.IsSet() {
		toSerialize["DateCreated"] = o.DateCreated.Get()
	}
	if o.CommunityRating.IsSet() {
		toSerialize["CommunityRating"] = o.CommunityRating.Get()
	}
	if o.FrameRate.IsSet() {
		toSerialize["FrameRate"] = o.FrameRate.Get()
	}
	if o.DownloadCount.IsSet() {
		toSerialize["DownloadCount"] = o.DownloadCount.Get()
	}
	if o.IsHashMatch.IsSet() {
		toSerialize["IsHashMatch"] = o.IsHashMatch.Get()
	}
	if o.AiTranslated.IsSet() {
		toSerialize["AiTranslated"] = o.AiTranslated.Get()
	}
	if o.MachineTranslated.IsSet() {
		toSerialize["MachineTranslated"] = o.MachineTranslated.Get()
	}
	if o.Forced.IsSet() {
		toSerialize["Forced"] = o.Forced.Get()
	}
	if o.HearingImpaired.IsSet() {
		toSerialize["HearingImpaired"] = o.HearingImpaired.Get()
	}
	return toSerialize, nil
}

type NullableRemoteSubtitleInfo struct {
	value *RemoteSubtitleInfo
	isSet bool
}

func (v NullableRemoteSubtitleInfo) Get() *RemoteSubtitleInfo {
	return v.value
}

func (v *NullableRemoteSubtitleInfo) Set(val *RemoteSubtitleInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoteSubtitleInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoteSubtitleInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoteSubtitleInfo(val *RemoteSubtitleInfo) *NullableRemoteSubtitleInfo {
	return &NullableRemoteSubtitleInfo{value: val, isSet: true}
}

func (v NullableRemoteSubtitleInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoteSubtitleInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


