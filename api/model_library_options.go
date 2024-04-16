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

// checks if the LibraryOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LibraryOptions{}

// LibraryOptions struct for LibraryOptions
type LibraryOptions struct {
	EnablePhotos *bool `json:"EnablePhotos,omitempty"`
	EnableRealtimeMonitor *bool `json:"EnableRealtimeMonitor,omitempty"`
	EnableChapterImageExtraction *bool `json:"EnableChapterImageExtraction,omitempty"`
	ExtractChapterImagesDuringLibraryScan *bool `json:"ExtractChapterImagesDuringLibraryScan,omitempty"`
	PathInfos []MediaPathInfo `json:"PathInfos,omitempty"`
	SaveLocalMetadata *bool `json:"SaveLocalMetadata,omitempty"`
	// Deprecated
	EnableInternetProviders *bool `json:"EnableInternetProviders,omitempty"`
	EnableAutomaticSeriesGrouping *bool `json:"EnableAutomaticSeriesGrouping,omitempty"`
	EnableEmbeddedTitles *bool `json:"EnableEmbeddedTitles,omitempty"`
	EnableEmbeddedEpisodeInfos *bool `json:"EnableEmbeddedEpisodeInfos,omitempty"`
	AutomaticRefreshIntervalDays *int32 `json:"AutomaticRefreshIntervalDays,omitempty"`
	// Gets or sets the preferred metadata language.
	PreferredMetadataLanguage NullableString `json:"PreferredMetadataLanguage,omitempty"`
	// Gets or sets the metadata country code.
	MetadataCountryCode NullableString `json:"MetadataCountryCode,omitempty"`
	SeasonZeroDisplayName *string `json:"SeasonZeroDisplayName,omitempty"`
	MetadataSavers []string `json:"MetadataSavers,omitempty"`
	DisabledLocalMetadataReaders []string `json:"DisabledLocalMetadataReaders,omitempty"`
	LocalMetadataReaderOrder []string `json:"LocalMetadataReaderOrder,omitempty"`
	DisabledSubtitleFetchers []string `json:"DisabledSubtitleFetchers,omitempty"`
	SubtitleFetcherOrder []string `json:"SubtitleFetcherOrder,omitempty"`
	SkipSubtitlesIfEmbeddedSubtitlesPresent *bool `json:"SkipSubtitlesIfEmbeddedSubtitlesPresent,omitempty"`
	SkipSubtitlesIfAudioTrackMatches *bool `json:"SkipSubtitlesIfAudioTrackMatches,omitempty"`
	SubtitleDownloadLanguages []string `json:"SubtitleDownloadLanguages,omitempty"`
	RequirePerfectSubtitleMatch *bool `json:"RequirePerfectSubtitleMatch,omitempty"`
	SaveSubtitlesWithMedia *bool `json:"SaveSubtitlesWithMedia,omitempty"`
	AutomaticallyAddToCollection *bool `json:"AutomaticallyAddToCollection,omitempty"`
	// An enum representing the options to disable embedded subs.
	AllowEmbeddedSubtitles *EmbeddedSubtitleOptions `json:"AllowEmbeddedSubtitles,omitempty"`
	TypeOptions []TypeOptions `json:"TypeOptions,omitempty"`
}

// NewLibraryOptions instantiates a new LibraryOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLibraryOptions() *LibraryOptions {
	this := LibraryOptions{}
	return &this
}

// NewLibraryOptionsWithDefaults instantiates a new LibraryOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLibraryOptionsWithDefaults() *LibraryOptions {
	this := LibraryOptions{}
	return &this
}

// GetEnablePhotos returns the EnablePhotos field value if set, zero value otherwise.
func (o *LibraryOptions) GetEnablePhotos() bool {
	if o == nil || IsNil(o.EnablePhotos) {
		var ret bool
		return ret
	}
	return *o.EnablePhotos
}

// GetEnablePhotosOk returns a tuple with the EnablePhotos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptions) GetEnablePhotosOk() (*bool, bool) {
	if o == nil || IsNil(o.EnablePhotos) {
		return nil, false
	}
	return o.EnablePhotos, true
}

// HasEnablePhotos returns a boolean if a field has been set.
func (o *LibraryOptions) HasEnablePhotos() bool {
	if o != nil && !IsNil(o.EnablePhotos) {
		return true
	}

	return false
}

// SetEnablePhotos gets a reference to the given bool and assigns it to the EnablePhotos field.
func (o *LibraryOptions) SetEnablePhotos(v bool) {
	o.EnablePhotos = &v
}

// GetEnableRealtimeMonitor returns the EnableRealtimeMonitor field value if set, zero value otherwise.
func (o *LibraryOptions) GetEnableRealtimeMonitor() bool {
	if o == nil || IsNil(o.EnableRealtimeMonitor) {
		var ret bool
		return ret
	}
	return *o.EnableRealtimeMonitor
}

// GetEnableRealtimeMonitorOk returns a tuple with the EnableRealtimeMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptions) GetEnableRealtimeMonitorOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableRealtimeMonitor) {
		return nil, false
	}
	return o.EnableRealtimeMonitor, true
}

// HasEnableRealtimeMonitor returns a boolean if a field has been set.
func (o *LibraryOptions) HasEnableRealtimeMonitor() bool {
	if o != nil && !IsNil(o.EnableRealtimeMonitor) {
		return true
	}

	return false
}

// SetEnableRealtimeMonitor gets a reference to the given bool and assigns it to the EnableRealtimeMonitor field.
func (o *LibraryOptions) SetEnableRealtimeMonitor(v bool) {
	o.EnableRealtimeMonitor = &v
}

// GetEnableChapterImageExtraction returns the EnableChapterImageExtraction field value if set, zero value otherwise.
func (o *LibraryOptions) GetEnableChapterImageExtraction() bool {
	if o == nil || IsNil(o.EnableChapterImageExtraction) {
		var ret bool
		return ret
	}
	return *o.EnableChapterImageExtraction
}

// GetEnableChapterImageExtractionOk returns a tuple with the EnableChapterImageExtraction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptions) GetEnableChapterImageExtractionOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableChapterImageExtraction) {
		return nil, false
	}
	return o.EnableChapterImageExtraction, true
}

// HasEnableChapterImageExtraction returns a boolean if a field has been set.
func (o *LibraryOptions) HasEnableChapterImageExtraction() bool {
	if o != nil && !IsNil(o.EnableChapterImageExtraction) {
		return true
	}

	return false
}

// SetEnableChapterImageExtraction gets a reference to the given bool and assigns it to the EnableChapterImageExtraction field.
func (o *LibraryOptions) SetEnableChapterImageExtraction(v bool) {
	o.EnableChapterImageExtraction = &v
}

// GetExtractChapterImagesDuringLibraryScan returns the ExtractChapterImagesDuringLibraryScan field value if set, zero value otherwise.
func (o *LibraryOptions) GetExtractChapterImagesDuringLibraryScan() bool {
	if o == nil || IsNil(o.ExtractChapterImagesDuringLibraryScan) {
		var ret bool
		return ret
	}
	return *o.ExtractChapterImagesDuringLibraryScan
}

// GetExtractChapterImagesDuringLibraryScanOk returns a tuple with the ExtractChapterImagesDuringLibraryScan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptions) GetExtractChapterImagesDuringLibraryScanOk() (*bool, bool) {
	if o == nil || IsNil(o.ExtractChapterImagesDuringLibraryScan) {
		return nil, false
	}
	return o.ExtractChapterImagesDuringLibraryScan, true
}

// HasExtractChapterImagesDuringLibraryScan returns a boolean if a field has been set.
func (o *LibraryOptions) HasExtractChapterImagesDuringLibraryScan() bool {
	if o != nil && !IsNil(o.ExtractChapterImagesDuringLibraryScan) {
		return true
	}

	return false
}

// SetExtractChapterImagesDuringLibraryScan gets a reference to the given bool and assigns it to the ExtractChapterImagesDuringLibraryScan field.
func (o *LibraryOptions) SetExtractChapterImagesDuringLibraryScan(v bool) {
	o.ExtractChapterImagesDuringLibraryScan = &v
}

// GetPathInfos returns the PathInfos field value if set, zero value otherwise.
func (o *LibraryOptions) GetPathInfos() []MediaPathInfo {
	if o == nil || IsNil(o.PathInfos) {
		var ret []MediaPathInfo
		return ret
	}
	return o.PathInfos
}

// GetPathInfosOk returns a tuple with the PathInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptions) GetPathInfosOk() ([]MediaPathInfo, bool) {
	if o == nil || IsNil(o.PathInfos) {
		return nil, false
	}
	return o.PathInfos, true
}

// HasPathInfos returns a boolean if a field has been set.
func (o *LibraryOptions) HasPathInfos() bool {
	if o != nil && !IsNil(o.PathInfos) {
		return true
	}

	return false
}

// SetPathInfos gets a reference to the given []MediaPathInfo and assigns it to the PathInfos field.
func (o *LibraryOptions) SetPathInfos(v []MediaPathInfo) {
	o.PathInfos = v
}

// GetSaveLocalMetadata returns the SaveLocalMetadata field value if set, zero value otherwise.
func (o *LibraryOptions) GetSaveLocalMetadata() bool {
	if o == nil || IsNil(o.SaveLocalMetadata) {
		var ret bool
		return ret
	}
	return *o.SaveLocalMetadata
}

// GetSaveLocalMetadataOk returns a tuple with the SaveLocalMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptions) GetSaveLocalMetadataOk() (*bool, bool) {
	if o == nil || IsNil(o.SaveLocalMetadata) {
		return nil, false
	}
	return o.SaveLocalMetadata, true
}

// HasSaveLocalMetadata returns a boolean if a field has been set.
func (o *LibraryOptions) HasSaveLocalMetadata() bool {
	if o != nil && !IsNil(o.SaveLocalMetadata) {
		return true
	}

	return false
}

// SetSaveLocalMetadata gets a reference to the given bool and assigns it to the SaveLocalMetadata field.
func (o *LibraryOptions) SetSaveLocalMetadata(v bool) {
	o.SaveLocalMetadata = &v
}

// GetEnableInternetProviders returns the EnableInternetProviders field value if set, zero value otherwise.
// Deprecated
func (o *LibraryOptions) GetEnableInternetProviders() bool {
	if o == nil || IsNil(o.EnableInternetProviders) {
		var ret bool
		return ret
	}
	return *o.EnableInternetProviders
}

// GetEnableInternetProvidersOk returns a tuple with the EnableInternetProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *LibraryOptions) GetEnableInternetProvidersOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableInternetProviders) {
		return nil, false
	}
	return o.EnableInternetProviders, true
}

// HasEnableInternetProviders returns a boolean if a field has been set.
func (o *LibraryOptions) HasEnableInternetProviders() bool {
	if o != nil && !IsNil(o.EnableInternetProviders) {
		return true
	}

	return false
}

// SetEnableInternetProviders gets a reference to the given bool and assigns it to the EnableInternetProviders field.
// Deprecated
func (o *LibraryOptions) SetEnableInternetProviders(v bool) {
	o.EnableInternetProviders = &v
}

// GetEnableAutomaticSeriesGrouping returns the EnableAutomaticSeriesGrouping field value if set, zero value otherwise.
func (o *LibraryOptions) GetEnableAutomaticSeriesGrouping() bool {
	if o == nil || IsNil(o.EnableAutomaticSeriesGrouping) {
		var ret bool
		return ret
	}
	return *o.EnableAutomaticSeriesGrouping
}

// GetEnableAutomaticSeriesGroupingOk returns a tuple with the EnableAutomaticSeriesGrouping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptions) GetEnableAutomaticSeriesGroupingOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableAutomaticSeriesGrouping) {
		return nil, false
	}
	return o.EnableAutomaticSeriesGrouping, true
}

// HasEnableAutomaticSeriesGrouping returns a boolean if a field has been set.
func (o *LibraryOptions) HasEnableAutomaticSeriesGrouping() bool {
	if o != nil && !IsNil(o.EnableAutomaticSeriesGrouping) {
		return true
	}

	return false
}

// SetEnableAutomaticSeriesGrouping gets a reference to the given bool and assigns it to the EnableAutomaticSeriesGrouping field.
func (o *LibraryOptions) SetEnableAutomaticSeriesGrouping(v bool) {
	o.EnableAutomaticSeriesGrouping = &v
}

// GetEnableEmbeddedTitles returns the EnableEmbeddedTitles field value if set, zero value otherwise.
func (o *LibraryOptions) GetEnableEmbeddedTitles() bool {
	if o == nil || IsNil(o.EnableEmbeddedTitles) {
		var ret bool
		return ret
	}
	return *o.EnableEmbeddedTitles
}

// GetEnableEmbeddedTitlesOk returns a tuple with the EnableEmbeddedTitles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptions) GetEnableEmbeddedTitlesOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableEmbeddedTitles) {
		return nil, false
	}
	return o.EnableEmbeddedTitles, true
}

// HasEnableEmbeddedTitles returns a boolean if a field has been set.
func (o *LibraryOptions) HasEnableEmbeddedTitles() bool {
	if o != nil && !IsNil(o.EnableEmbeddedTitles) {
		return true
	}

	return false
}

// SetEnableEmbeddedTitles gets a reference to the given bool and assigns it to the EnableEmbeddedTitles field.
func (o *LibraryOptions) SetEnableEmbeddedTitles(v bool) {
	o.EnableEmbeddedTitles = &v
}

// GetEnableEmbeddedEpisodeInfos returns the EnableEmbeddedEpisodeInfos field value if set, zero value otherwise.
func (o *LibraryOptions) GetEnableEmbeddedEpisodeInfos() bool {
	if o == nil || IsNil(o.EnableEmbeddedEpisodeInfos) {
		var ret bool
		return ret
	}
	return *o.EnableEmbeddedEpisodeInfos
}

// GetEnableEmbeddedEpisodeInfosOk returns a tuple with the EnableEmbeddedEpisodeInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptions) GetEnableEmbeddedEpisodeInfosOk() (*bool, bool) {
	if o == nil || IsNil(o.EnableEmbeddedEpisodeInfos) {
		return nil, false
	}
	return o.EnableEmbeddedEpisodeInfos, true
}

// HasEnableEmbeddedEpisodeInfos returns a boolean if a field has been set.
func (o *LibraryOptions) HasEnableEmbeddedEpisodeInfos() bool {
	if o != nil && !IsNil(o.EnableEmbeddedEpisodeInfos) {
		return true
	}

	return false
}

// SetEnableEmbeddedEpisodeInfos gets a reference to the given bool and assigns it to the EnableEmbeddedEpisodeInfos field.
func (o *LibraryOptions) SetEnableEmbeddedEpisodeInfos(v bool) {
	o.EnableEmbeddedEpisodeInfos = &v
}

// GetAutomaticRefreshIntervalDays returns the AutomaticRefreshIntervalDays field value if set, zero value otherwise.
func (o *LibraryOptions) GetAutomaticRefreshIntervalDays() int32 {
	if o == nil || IsNil(o.AutomaticRefreshIntervalDays) {
		var ret int32
		return ret
	}
	return *o.AutomaticRefreshIntervalDays
}

// GetAutomaticRefreshIntervalDaysOk returns a tuple with the AutomaticRefreshIntervalDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptions) GetAutomaticRefreshIntervalDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.AutomaticRefreshIntervalDays) {
		return nil, false
	}
	return o.AutomaticRefreshIntervalDays, true
}

// HasAutomaticRefreshIntervalDays returns a boolean if a field has been set.
func (o *LibraryOptions) HasAutomaticRefreshIntervalDays() bool {
	if o != nil && !IsNil(o.AutomaticRefreshIntervalDays) {
		return true
	}

	return false
}

// SetAutomaticRefreshIntervalDays gets a reference to the given int32 and assigns it to the AutomaticRefreshIntervalDays field.
func (o *LibraryOptions) SetAutomaticRefreshIntervalDays(v int32) {
	o.AutomaticRefreshIntervalDays = &v
}

// GetPreferredMetadataLanguage returns the PreferredMetadataLanguage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LibraryOptions) GetPreferredMetadataLanguage() string {
	if o == nil || IsNil(o.PreferredMetadataLanguage.Get()) {
		var ret string
		return ret
	}
	return *o.PreferredMetadataLanguage.Get()
}

// GetPreferredMetadataLanguageOk returns a tuple with the PreferredMetadataLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LibraryOptions) GetPreferredMetadataLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreferredMetadataLanguage.Get(), o.PreferredMetadataLanguage.IsSet()
}

// HasPreferredMetadataLanguage returns a boolean if a field has been set.
func (o *LibraryOptions) HasPreferredMetadataLanguage() bool {
	if o != nil && o.PreferredMetadataLanguage.IsSet() {
		return true
	}

	return false
}

// SetPreferredMetadataLanguage gets a reference to the given NullableString and assigns it to the PreferredMetadataLanguage field.
func (o *LibraryOptions) SetPreferredMetadataLanguage(v string) {
	o.PreferredMetadataLanguage.Set(&v)
}
// SetPreferredMetadataLanguageNil sets the value for PreferredMetadataLanguage to be an explicit nil
func (o *LibraryOptions) SetPreferredMetadataLanguageNil() {
	o.PreferredMetadataLanguage.Set(nil)
}

// UnsetPreferredMetadataLanguage ensures that no value is present for PreferredMetadataLanguage, not even an explicit nil
func (o *LibraryOptions) UnsetPreferredMetadataLanguage() {
	o.PreferredMetadataLanguage.Unset()
}

// GetMetadataCountryCode returns the MetadataCountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LibraryOptions) GetMetadataCountryCode() string {
	if o == nil || IsNil(o.MetadataCountryCode.Get()) {
		var ret string
		return ret
	}
	return *o.MetadataCountryCode.Get()
}

// GetMetadataCountryCodeOk returns a tuple with the MetadataCountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LibraryOptions) GetMetadataCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetadataCountryCode.Get(), o.MetadataCountryCode.IsSet()
}

// HasMetadataCountryCode returns a boolean if a field has been set.
func (o *LibraryOptions) HasMetadataCountryCode() bool {
	if o != nil && o.MetadataCountryCode.IsSet() {
		return true
	}

	return false
}

// SetMetadataCountryCode gets a reference to the given NullableString and assigns it to the MetadataCountryCode field.
func (o *LibraryOptions) SetMetadataCountryCode(v string) {
	o.MetadataCountryCode.Set(&v)
}
// SetMetadataCountryCodeNil sets the value for MetadataCountryCode to be an explicit nil
func (o *LibraryOptions) SetMetadataCountryCodeNil() {
	o.MetadataCountryCode.Set(nil)
}

// UnsetMetadataCountryCode ensures that no value is present for MetadataCountryCode, not even an explicit nil
func (o *LibraryOptions) UnsetMetadataCountryCode() {
	o.MetadataCountryCode.Unset()
}

// GetSeasonZeroDisplayName returns the SeasonZeroDisplayName field value if set, zero value otherwise.
func (o *LibraryOptions) GetSeasonZeroDisplayName() string {
	if o == nil || IsNil(o.SeasonZeroDisplayName) {
		var ret string
		return ret
	}
	return *o.SeasonZeroDisplayName
}

// GetSeasonZeroDisplayNameOk returns a tuple with the SeasonZeroDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptions) GetSeasonZeroDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.SeasonZeroDisplayName) {
		return nil, false
	}
	return o.SeasonZeroDisplayName, true
}

// HasSeasonZeroDisplayName returns a boolean if a field has been set.
func (o *LibraryOptions) HasSeasonZeroDisplayName() bool {
	if o != nil && !IsNil(o.SeasonZeroDisplayName) {
		return true
	}

	return false
}

// SetSeasonZeroDisplayName gets a reference to the given string and assigns it to the SeasonZeroDisplayName field.
func (o *LibraryOptions) SetSeasonZeroDisplayName(v string) {
	o.SeasonZeroDisplayName = &v
}

// GetMetadataSavers returns the MetadataSavers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LibraryOptions) GetMetadataSavers() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.MetadataSavers
}

// GetMetadataSaversOk returns a tuple with the MetadataSavers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LibraryOptions) GetMetadataSaversOk() ([]string, bool) {
	if o == nil || IsNil(o.MetadataSavers) {
		return nil, false
	}
	return o.MetadataSavers, true
}

// HasMetadataSavers returns a boolean if a field has been set.
func (o *LibraryOptions) HasMetadataSavers() bool {
	if o != nil && !IsNil(o.MetadataSavers) {
		return true
	}

	return false
}

// SetMetadataSavers gets a reference to the given []string and assigns it to the MetadataSavers field.
func (o *LibraryOptions) SetMetadataSavers(v []string) {
	o.MetadataSavers = v
}

// GetDisabledLocalMetadataReaders returns the DisabledLocalMetadataReaders field value if set, zero value otherwise.
func (o *LibraryOptions) GetDisabledLocalMetadataReaders() []string {
	if o == nil || IsNil(o.DisabledLocalMetadataReaders) {
		var ret []string
		return ret
	}
	return o.DisabledLocalMetadataReaders
}

// GetDisabledLocalMetadataReadersOk returns a tuple with the DisabledLocalMetadataReaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptions) GetDisabledLocalMetadataReadersOk() ([]string, bool) {
	if o == nil || IsNil(o.DisabledLocalMetadataReaders) {
		return nil, false
	}
	return o.DisabledLocalMetadataReaders, true
}

// HasDisabledLocalMetadataReaders returns a boolean if a field has been set.
func (o *LibraryOptions) HasDisabledLocalMetadataReaders() bool {
	if o != nil && !IsNil(o.DisabledLocalMetadataReaders) {
		return true
	}

	return false
}

// SetDisabledLocalMetadataReaders gets a reference to the given []string and assigns it to the DisabledLocalMetadataReaders field.
func (o *LibraryOptions) SetDisabledLocalMetadataReaders(v []string) {
	o.DisabledLocalMetadataReaders = v
}

// GetLocalMetadataReaderOrder returns the LocalMetadataReaderOrder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LibraryOptions) GetLocalMetadataReaderOrder() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.LocalMetadataReaderOrder
}

// GetLocalMetadataReaderOrderOk returns a tuple with the LocalMetadataReaderOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LibraryOptions) GetLocalMetadataReaderOrderOk() ([]string, bool) {
	if o == nil || IsNil(o.LocalMetadataReaderOrder) {
		return nil, false
	}
	return o.LocalMetadataReaderOrder, true
}

// HasLocalMetadataReaderOrder returns a boolean if a field has been set.
func (o *LibraryOptions) HasLocalMetadataReaderOrder() bool {
	if o != nil && !IsNil(o.LocalMetadataReaderOrder) {
		return true
	}

	return false
}

// SetLocalMetadataReaderOrder gets a reference to the given []string and assigns it to the LocalMetadataReaderOrder field.
func (o *LibraryOptions) SetLocalMetadataReaderOrder(v []string) {
	o.LocalMetadataReaderOrder = v
}

// GetDisabledSubtitleFetchers returns the DisabledSubtitleFetchers field value if set, zero value otherwise.
func (o *LibraryOptions) GetDisabledSubtitleFetchers() []string {
	if o == nil || IsNil(o.DisabledSubtitleFetchers) {
		var ret []string
		return ret
	}
	return o.DisabledSubtitleFetchers
}

// GetDisabledSubtitleFetchersOk returns a tuple with the DisabledSubtitleFetchers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptions) GetDisabledSubtitleFetchersOk() ([]string, bool) {
	if o == nil || IsNil(o.DisabledSubtitleFetchers) {
		return nil, false
	}
	return o.DisabledSubtitleFetchers, true
}

// HasDisabledSubtitleFetchers returns a boolean if a field has been set.
func (o *LibraryOptions) HasDisabledSubtitleFetchers() bool {
	if o != nil && !IsNil(o.DisabledSubtitleFetchers) {
		return true
	}

	return false
}

// SetDisabledSubtitleFetchers gets a reference to the given []string and assigns it to the DisabledSubtitleFetchers field.
func (o *LibraryOptions) SetDisabledSubtitleFetchers(v []string) {
	o.DisabledSubtitleFetchers = v
}

// GetSubtitleFetcherOrder returns the SubtitleFetcherOrder field value if set, zero value otherwise.
func (o *LibraryOptions) GetSubtitleFetcherOrder() []string {
	if o == nil || IsNil(o.SubtitleFetcherOrder) {
		var ret []string
		return ret
	}
	return o.SubtitleFetcherOrder
}

// GetSubtitleFetcherOrderOk returns a tuple with the SubtitleFetcherOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptions) GetSubtitleFetcherOrderOk() ([]string, bool) {
	if o == nil || IsNil(o.SubtitleFetcherOrder) {
		return nil, false
	}
	return o.SubtitleFetcherOrder, true
}

// HasSubtitleFetcherOrder returns a boolean if a field has been set.
func (o *LibraryOptions) HasSubtitleFetcherOrder() bool {
	if o != nil && !IsNil(o.SubtitleFetcherOrder) {
		return true
	}

	return false
}

// SetSubtitleFetcherOrder gets a reference to the given []string and assigns it to the SubtitleFetcherOrder field.
func (o *LibraryOptions) SetSubtitleFetcherOrder(v []string) {
	o.SubtitleFetcherOrder = v
}

// GetSkipSubtitlesIfEmbeddedSubtitlesPresent returns the SkipSubtitlesIfEmbeddedSubtitlesPresent field value if set, zero value otherwise.
func (o *LibraryOptions) GetSkipSubtitlesIfEmbeddedSubtitlesPresent() bool {
	if o == nil || IsNil(o.SkipSubtitlesIfEmbeddedSubtitlesPresent) {
		var ret bool
		return ret
	}
	return *o.SkipSubtitlesIfEmbeddedSubtitlesPresent
}

// GetSkipSubtitlesIfEmbeddedSubtitlesPresentOk returns a tuple with the SkipSubtitlesIfEmbeddedSubtitlesPresent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptions) GetSkipSubtitlesIfEmbeddedSubtitlesPresentOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipSubtitlesIfEmbeddedSubtitlesPresent) {
		return nil, false
	}
	return o.SkipSubtitlesIfEmbeddedSubtitlesPresent, true
}

// HasSkipSubtitlesIfEmbeddedSubtitlesPresent returns a boolean if a field has been set.
func (o *LibraryOptions) HasSkipSubtitlesIfEmbeddedSubtitlesPresent() bool {
	if o != nil && !IsNil(o.SkipSubtitlesIfEmbeddedSubtitlesPresent) {
		return true
	}

	return false
}

// SetSkipSubtitlesIfEmbeddedSubtitlesPresent gets a reference to the given bool and assigns it to the SkipSubtitlesIfEmbeddedSubtitlesPresent field.
func (o *LibraryOptions) SetSkipSubtitlesIfEmbeddedSubtitlesPresent(v bool) {
	o.SkipSubtitlesIfEmbeddedSubtitlesPresent = &v
}

// GetSkipSubtitlesIfAudioTrackMatches returns the SkipSubtitlesIfAudioTrackMatches field value if set, zero value otherwise.
func (o *LibraryOptions) GetSkipSubtitlesIfAudioTrackMatches() bool {
	if o == nil || IsNil(o.SkipSubtitlesIfAudioTrackMatches) {
		var ret bool
		return ret
	}
	return *o.SkipSubtitlesIfAudioTrackMatches
}

// GetSkipSubtitlesIfAudioTrackMatchesOk returns a tuple with the SkipSubtitlesIfAudioTrackMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptions) GetSkipSubtitlesIfAudioTrackMatchesOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipSubtitlesIfAudioTrackMatches) {
		return nil, false
	}
	return o.SkipSubtitlesIfAudioTrackMatches, true
}

// HasSkipSubtitlesIfAudioTrackMatches returns a boolean if a field has been set.
func (o *LibraryOptions) HasSkipSubtitlesIfAudioTrackMatches() bool {
	if o != nil && !IsNil(o.SkipSubtitlesIfAudioTrackMatches) {
		return true
	}

	return false
}

// SetSkipSubtitlesIfAudioTrackMatches gets a reference to the given bool and assigns it to the SkipSubtitlesIfAudioTrackMatches field.
func (o *LibraryOptions) SetSkipSubtitlesIfAudioTrackMatches(v bool) {
	o.SkipSubtitlesIfAudioTrackMatches = &v
}

// GetSubtitleDownloadLanguages returns the SubtitleDownloadLanguages field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LibraryOptions) GetSubtitleDownloadLanguages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SubtitleDownloadLanguages
}

// GetSubtitleDownloadLanguagesOk returns a tuple with the SubtitleDownloadLanguages field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LibraryOptions) GetSubtitleDownloadLanguagesOk() ([]string, bool) {
	if o == nil || IsNil(o.SubtitleDownloadLanguages) {
		return nil, false
	}
	return o.SubtitleDownloadLanguages, true
}

// HasSubtitleDownloadLanguages returns a boolean if a field has been set.
func (o *LibraryOptions) HasSubtitleDownloadLanguages() bool {
	if o != nil && !IsNil(o.SubtitleDownloadLanguages) {
		return true
	}

	return false
}

// SetSubtitleDownloadLanguages gets a reference to the given []string and assigns it to the SubtitleDownloadLanguages field.
func (o *LibraryOptions) SetSubtitleDownloadLanguages(v []string) {
	o.SubtitleDownloadLanguages = v
}

// GetRequirePerfectSubtitleMatch returns the RequirePerfectSubtitleMatch field value if set, zero value otherwise.
func (o *LibraryOptions) GetRequirePerfectSubtitleMatch() bool {
	if o == nil || IsNil(o.RequirePerfectSubtitleMatch) {
		var ret bool
		return ret
	}
	return *o.RequirePerfectSubtitleMatch
}

// GetRequirePerfectSubtitleMatchOk returns a tuple with the RequirePerfectSubtitleMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptions) GetRequirePerfectSubtitleMatchOk() (*bool, bool) {
	if o == nil || IsNil(o.RequirePerfectSubtitleMatch) {
		return nil, false
	}
	return o.RequirePerfectSubtitleMatch, true
}

// HasRequirePerfectSubtitleMatch returns a boolean if a field has been set.
func (o *LibraryOptions) HasRequirePerfectSubtitleMatch() bool {
	if o != nil && !IsNil(o.RequirePerfectSubtitleMatch) {
		return true
	}

	return false
}

// SetRequirePerfectSubtitleMatch gets a reference to the given bool and assigns it to the RequirePerfectSubtitleMatch field.
func (o *LibraryOptions) SetRequirePerfectSubtitleMatch(v bool) {
	o.RequirePerfectSubtitleMatch = &v
}

// GetSaveSubtitlesWithMedia returns the SaveSubtitlesWithMedia field value if set, zero value otherwise.
func (o *LibraryOptions) GetSaveSubtitlesWithMedia() bool {
	if o == nil || IsNil(o.SaveSubtitlesWithMedia) {
		var ret bool
		return ret
	}
	return *o.SaveSubtitlesWithMedia
}

// GetSaveSubtitlesWithMediaOk returns a tuple with the SaveSubtitlesWithMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptions) GetSaveSubtitlesWithMediaOk() (*bool, bool) {
	if o == nil || IsNil(o.SaveSubtitlesWithMedia) {
		return nil, false
	}
	return o.SaveSubtitlesWithMedia, true
}

// HasSaveSubtitlesWithMedia returns a boolean if a field has been set.
func (o *LibraryOptions) HasSaveSubtitlesWithMedia() bool {
	if o != nil && !IsNil(o.SaveSubtitlesWithMedia) {
		return true
	}

	return false
}

// SetSaveSubtitlesWithMedia gets a reference to the given bool and assigns it to the SaveSubtitlesWithMedia field.
func (o *LibraryOptions) SetSaveSubtitlesWithMedia(v bool) {
	o.SaveSubtitlesWithMedia = &v
}

// GetAutomaticallyAddToCollection returns the AutomaticallyAddToCollection field value if set, zero value otherwise.
func (o *LibraryOptions) GetAutomaticallyAddToCollection() bool {
	if o == nil || IsNil(o.AutomaticallyAddToCollection) {
		var ret bool
		return ret
	}
	return *o.AutomaticallyAddToCollection
}

// GetAutomaticallyAddToCollectionOk returns a tuple with the AutomaticallyAddToCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptions) GetAutomaticallyAddToCollectionOk() (*bool, bool) {
	if o == nil || IsNil(o.AutomaticallyAddToCollection) {
		return nil, false
	}
	return o.AutomaticallyAddToCollection, true
}

// HasAutomaticallyAddToCollection returns a boolean if a field has been set.
func (o *LibraryOptions) HasAutomaticallyAddToCollection() bool {
	if o != nil && !IsNil(o.AutomaticallyAddToCollection) {
		return true
	}

	return false
}

// SetAutomaticallyAddToCollection gets a reference to the given bool and assigns it to the AutomaticallyAddToCollection field.
func (o *LibraryOptions) SetAutomaticallyAddToCollection(v bool) {
	o.AutomaticallyAddToCollection = &v
}

// GetAllowEmbeddedSubtitles returns the AllowEmbeddedSubtitles field value if set, zero value otherwise.
func (o *LibraryOptions) GetAllowEmbeddedSubtitles() EmbeddedSubtitleOptions {
	if o == nil || IsNil(o.AllowEmbeddedSubtitles) {
		var ret EmbeddedSubtitleOptions
		return ret
	}
	return *o.AllowEmbeddedSubtitles
}

// GetAllowEmbeddedSubtitlesOk returns a tuple with the AllowEmbeddedSubtitles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptions) GetAllowEmbeddedSubtitlesOk() (*EmbeddedSubtitleOptions, bool) {
	if o == nil || IsNil(o.AllowEmbeddedSubtitles) {
		return nil, false
	}
	return o.AllowEmbeddedSubtitles, true
}

// HasAllowEmbeddedSubtitles returns a boolean if a field has been set.
func (o *LibraryOptions) HasAllowEmbeddedSubtitles() bool {
	if o != nil && !IsNil(o.AllowEmbeddedSubtitles) {
		return true
	}

	return false
}

// SetAllowEmbeddedSubtitles gets a reference to the given EmbeddedSubtitleOptions and assigns it to the AllowEmbeddedSubtitles field.
func (o *LibraryOptions) SetAllowEmbeddedSubtitles(v EmbeddedSubtitleOptions) {
	o.AllowEmbeddedSubtitles = &v
}

// GetTypeOptions returns the TypeOptions field value if set, zero value otherwise.
func (o *LibraryOptions) GetTypeOptions() []TypeOptions {
	if o == nil || IsNil(o.TypeOptions) {
		var ret []TypeOptions
		return ret
	}
	return o.TypeOptions
}

// GetTypeOptionsOk returns a tuple with the TypeOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryOptions) GetTypeOptionsOk() ([]TypeOptions, bool) {
	if o == nil || IsNil(o.TypeOptions) {
		return nil, false
	}
	return o.TypeOptions, true
}

// HasTypeOptions returns a boolean if a field has been set.
func (o *LibraryOptions) HasTypeOptions() bool {
	if o != nil && !IsNil(o.TypeOptions) {
		return true
	}

	return false
}

// SetTypeOptions gets a reference to the given []TypeOptions and assigns it to the TypeOptions field.
func (o *LibraryOptions) SetTypeOptions(v []TypeOptions) {
	o.TypeOptions = v
}

func (o LibraryOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LibraryOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnablePhotos) {
		toSerialize["EnablePhotos"] = o.EnablePhotos
	}
	if !IsNil(o.EnableRealtimeMonitor) {
		toSerialize["EnableRealtimeMonitor"] = o.EnableRealtimeMonitor
	}
	if !IsNil(o.EnableChapterImageExtraction) {
		toSerialize["EnableChapterImageExtraction"] = o.EnableChapterImageExtraction
	}
	if !IsNil(o.ExtractChapterImagesDuringLibraryScan) {
		toSerialize["ExtractChapterImagesDuringLibraryScan"] = o.ExtractChapterImagesDuringLibraryScan
	}
	if !IsNil(o.PathInfos) {
		toSerialize["PathInfos"] = o.PathInfos
	}
	if !IsNil(o.SaveLocalMetadata) {
		toSerialize["SaveLocalMetadata"] = o.SaveLocalMetadata
	}
	if !IsNil(o.EnableInternetProviders) {
		toSerialize["EnableInternetProviders"] = o.EnableInternetProviders
	}
	if !IsNil(o.EnableAutomaticSeriesGrouping) {
		toSerialize["EnableAutomaticSeriesGrouping"] = o.EnableAutomaticSeriesGrouping
	}
	if !IsNil(o.EnableEmbeddedTitles) {
		toSerialize["EnableEmbeddedTitles"] = o.EnableEmbeddedTitles
	}
	if !IsNil(o.EnableEmbeddedEpisodeInfos) {
		toSerialize["EnableEmbeddedEpisodeInfos"] = o.EnableEmbeddedEpisodeInfos
	}
	if !IsNil(o.AutomaticRefreshIntervalDays) {
		toSerialize["AutomaticRefreshIntervalDays"] = o.AutomaticRefreshIntervalDays
	}
	if o.PreferredMetadataLanguage.IsSet() {
		toSerialize["PreferredMetadataLanguage"] = o.PreferredMetadataLanguage.Get()
	}
	if o.MetadataCountryCode.IsSet() {
		toSerialize["MetadataCountryCode"] = o.MetadataCountryCode.Get()
	}
	if !IsNil(o.SeasonZeroDisplayName) {
		toSerialize["SeasonZeroDisplayName"] = o.SeasonZeroDisplayName
	}
	if o.MetadataSavers != nil {
		toSerialize["MetadataSavers"] = o.MetadataSavers
	}
	if !IsNil(o.DisabledLocalMetadataReaders) {
		toSerialize["DisabledLocalMetadataReaders"] = o.DisabledLocalMetadataReaders
	}
	if o.LocalMetadataReaderOrder != nil {
		toSerialize["LocalMetadataReaderOrder"] = o.LocalMetadataReaderOrder
	}
	if !IsNil(o.DisabledSubtitleFetchers) {
		toSerialize["DisabledSubtitleFetchers"] = o.DisabledSubtitleFetchers
	}
	if !IsNil(o.SubtitleFetcherOrder) {
		toSerialize["SubtitleFetcherOrder"] = o.SubtitleFetcherOrder
	}
	if !IsNil(o.SkipSubtitlesIfEmbeddedSubtitlesPresent) {
		toSerialize["SkipSubtitlesIfEmbeddedSubtitlesPresent"] = o.SkipSubtitlesIfEmbeddedSubtitlesPresent
	}
	if !IsNil(o.SkipSubtitlesIfAudioTrackMatches) {
		toSerialize["SkipSubtitlesIfAudioTrackMatches"] = o.SkipSubtitlesIfAudioTrackMatches
	}
	if o.SubtitleDownloadLanguages != nil {
		toSerialize["SubtitleDownloadLanguages"] = o.SubtitleDownloadLanguages
	}
	if !IsNil(o.RequirePerfectSubtitleMatch) {
		toSerialize["RequirePerfectSubtitleMatch"] = o.RequirePerfectSubtitleMatch
	}
	if !IsNil(o.SaveSubtitlesWithMedia) {
		toSerialize["SaveSubtitlesWithMedia"] = o.SaveSubtitlesWithMedia
	}
	if !IsNil(o.AutomaticallyAddToCollection) {
		toSerialize["AutomaticallyAddToCollection"] = o.AutomaticallyAddToCollection
	}
	if !IsNil(o.AllowEmbeddedSubtitles) {
		toSerialize["AllowEmbeddedSubtitles"] = o.AllowEmbeddedSubtitles
	}
	if !IsNil(o.TypeOptions) {
		toSerialize["TypeOptions"] = o.TypeOptions
	}
	return toSerialize, nil
}

type NullableLibraryOptions struct {
	value *LibraryOptions
	isSet bool
}

func (v NullableLibraryOptions) Get() *LibraryOptions {
	return v.value
}

func (v *NullableLibraryOptions) Set(val *LibraryOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableLibraryOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableLibraryOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLibraryOptions(val *LibraryOptions) *NullableLibraryOptions {
	return &NullableLibraryOptions{value: val, isSet: true}
}

func (v NullableLibraryOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLibraryOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


