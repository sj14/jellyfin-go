/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the BaseItemDtoImageBlurHashes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseItemDtoImageBlurHashes{}

// BaseItemDtoImageBlurHashes Gets or sets the blurhashes for the image tags.  Maps image type to dictionary mapping image tag to blurhash value.
type BaseItemDtoImageBlurHashes struct {
	Primary *map[string]string `json:"Primary,omitempty"`
	Art *map[string]string `json:"Art,omitempty"`
	Backdrop *map[string]string `json:"Backdrop,omitempty"`
	Banner *map[string]string `json:"Banner,omitempty"`
	Logo *map[string]string `json:"Logo,omitempty"`
	Thumb *map[string]string `json:"Thumb,omitempty"`
	Disc *map[string]string `json:"Disc,omitempty"`
	Box *map[string]string `json:"Box,omitempty"`
	Screenshot *map[string]string `json:"Screenshot,omitempty"`
	Menu *map[string]string `json:"Menu,omitempty"`
	Chapter *map[string]string `json:"Chapter,omitempty"`
	BoxRear *map[string]string `json:"BoxRear,omitempty"`
	Profile *map[string]string `json:"Profile,omitempty"`
}

// NewBaseItemDtoImageBlurHashes instantiates a new BaseItemDtoImageBlurHashes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseItemDtoImageBlurHashes() *BaseItemDtoImageBlurHashes {
	this := BaseItemDtoImageBlurHashes{}
	return &this
}

// NewBaseItemDtoImageBlurHashesWithDefaults instantiates a new BaseItemDtoImageBlurHashes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseItemDtoImageBlurHashesWithDefaults() *BaseItemDtoImageBlurHashes {
	this := BaseItemDtoImageBlurHashes{}
	return &this
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *BaseItemDtoImageBlurHashes) GetPrimary() map[string]string {
	if o == nil || IsNil(o.Primary) {
		var ret map[string]string
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItemDtoImageBlurHashes) GetPrimaryOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Primary) {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *BaseItemDtoImageBlurHashes) HasPrimary() bool {
	if o != nil && !IsNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given map[string]string and assigns it to the Primary field.
func (o *BaseItemDtoImageBlurHashes) SetPrimary(v map[string]string) {
	o.Primary = &v
}

// GetArt returns the Art field value if set, zero value otherwise.
func (o *BaseItemDtoImageBlurHashes) GetArt() map[string]string {
	if o == nil || IsNil(o.Art) {
		var ret map[string]string
		return ret
	}
	return *o.Art
}

// GetArtOk returns a tuple with the Art field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItemDtoImageBlurHashes) GetArtOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Art) {
		return nil, false
	}
	return o.Art, true
}

// HasArt returns a boolean if a field has been set.
func (o *BaseItemDtoImageBlurHashes) HasArt() bool {
	if o != nil && !IsNil(o.Art) {
		return true
	}

	return false
}

// SetArt gets a reference to the given map[string]string and assigns it to the Art field.
func (o *BaseItemDtoImageBlurHashes) SetArt(v map[string]string) {
	o.Art = &v
}

// GetBackdrop returns the Backdrop field value if set, zero value otherwise.
func (o *BaseItemDtoImageBlurHashes) GetBackdrop() map[string]string {
	if o == nil || IsNil(o.Backdrop) {
		var ret map[string]string
		return ret
	}
	return *o.Backdrop
}

// GetBackdropOk returns a tuple with the Backdrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItemDtoImageBlurHashes) GetBackdropOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Backdrop) {
		return nil, false
	}
	return o.Backdrop, true
}

// HasBackdrop returns a boolean if a field has been set.
func (o *BaseItemDtoImageBlurHashes) HasBackdrop() bool {
	if o != nil && !IsNil(o.Backdrop) {
		return true
	}

	return false
}

// SetBackdrop gets a reference to the given map[string]string and assigns it to the Backdrop field.
func (o *BaseItemDtoImageBlurHashes) SetBackdrop(v map[string]string) {
	o.Backdrop = &v
}

// GetBanner returns the Banner field value if set, zero value otherwise.
func (o *BaseItemDtoImageBlurHashes) GetBanner() map[string]string {
	if o == nil || IsNil(o.Banner) {
		var ret map[string]string
		return ret
	}
	return *o.Banner
}

// GetBannerOk returns a tuple with the Banner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItemDtoImageBlurHashes) GetBannerOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Banner) {
		return nil, false
	}
	return o.Banner, true
}

// HasBanner returns a boolean if a field has been set.
func (o *BaseItemDtoImageBlurHashes) HasBanner() bool {
	if o != nil && !IsNil(o.Banner) {
		return true
	}

	return false
}

// SetBanner gets a reference to the given map[string]string and assigns it to the Banner field.
func (o *BaseItemDtoImageBlurHashes) SetBanner(v map[string]string) {
	o.Banner = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *BaseItemDtoImageBlurHashes) GetLogo() map[string]string {
	if o == nil || IsNil(o.Logo) {
		var ret map[string]string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItemDtoImageBlurHashes) GetLogoOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *BaseItemDtoImageBlurHashes) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given map[string]string and assigns it to the Logo field.
func (o *BaseItemDtoImageBlurHashes) SetLogo(v map[string]string) {
	o.Logo = &v
}

// GetThumb returns the Thumb field value if set, zero value otherwise.
func (o *BaseItemDtoImageBlurHashes) GetThumb() map[string]string {
	if o == nil || IsNil(o.Thumb) {
		var ret map[string]string
		return ret
	}
	return *o.Thumb
}

// GetThumbOk returns a tuple with the Thumb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItemDtoImageBlurHashes) GetThumbOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Thumb) {
		return nil, false
	}
	return o.Thumb, true
}

// HasThumb returns a boolean if a field has been set.
func (o *BaseItemDtoImageBlurHashes) HasThumb() bool {
	if o != nil && !IsNil(o.Thumb) {
		return true
	}

	return false
}

// SetThumb gets a reference to the given map[string]string and assigns it to the Thumb field.
func (o *BaseItemDtoImageBlurHashes) SetThumb(v map[string]string) {
	o.Thumb = &v
}

// GetDisc returns the Disc field value if set, zero value otherwise.
func (o *BaseItemDtoImageBlurHashes) GetDisc() map[string]string {
	if o == nil || IsNil(o.Disc) {
		var ret map[string]string
		return ret
	}
	return *o.Disc
}

// GetDiscOk returns a tuple with the Disc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItemDtoImageBlurHashes) GetDiscOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Disc) {
		return nil, false
	}
	return o.Disc, true
}

// HasDisc returns a boolean if a field has been set.
func (o *BaseItemDtoImageBlurHashes) HasDisc() bool {
	if o != nil && !IsNil(o.Disc) {
		return true
	}

	return false
}

// SetDisc gets a reference to the given map[string]string and assigns it to the Disc field.
func (o *BaseItemDtoImageBlurHashes) SetDisc(v map[string]string) {
	o.Disc = &v
}

// GetBox returns the Box field value if set, zero value otherwise.
func (o *BaseItemDtoImageBlurHashes) GetBox() map[string]string {
	if o == nil || IsNil(o.Box) {
		var ret map[string]string
		return ret
	}
	return *o.Box
}

// GetBoxOk returns a tuple with the Box field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItemDtoImageBlurHashes) GetBoxOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Box) {
		return nil, false
	}
	return o.Box, true
}

// HasBox returns a boolean if a field has been set.
func (o *BaseItemDtoImageBlurHashes) HasBox() bool {
	if o != nil && !IsNil(o.Box) {
		return true
	}

	return false
}

// SetBox gets a reference to the given map[string]string and assigns it to the Box field.
func (o *BaseItemDtoImageBlurHashes) SetBox(v map[string]string) {
	o.Box = &v
}

// GetScreenshot returns the Screenshot field value if set, zero value otherwise.
func (o *BaseItemDtoImageBlurHashes) GetScreenshot() map[string]string {
	if o == nil || IsNil(o.Screenshot) {
		var ret map[string]string
		return ret
	}
	return *o.Screenshot
}

// GetScreenshotOk returns a tuple with the Screenshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItemDtoImageBlurHashes) GetScreenshotOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Screenshot) {
		return nil, false
	}
	return o.Screenshot, true
}

// HasScreenshot returns a boolean if a field has been set.
func (o *BaseItemDtoImageBlurHashes) HasScreenshot() bool {
	if o != nil && !IsNil(o.Screenshot) {
		return true
	}

	return false
}

// SetScreenshot gets a reference to the given map[string]string and assigns it to the Screenshot field.
func (o *BaseItemDtoImageBlurHashes) SetScreenshot(v map[string]string) {
	o.Screenshot = &v
}

// GetMenu returns the Menu field value if set, zero value otherwise.
func (o *BaseItemDtoImageBlurHashes) GetMenu() map[string]string {
	if o == nil || IsNil(o.Menu) {
		var ret map[string]string
		return ret
	}
	return *o.Menu
}

// GetMenuOk returns a tuple with the Menu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItemDtoImageBlurHashes) GetMenuOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Menu) {
		return nil, false
	}
	return o.Menu, true
}

// HasMenu returns a boolean if a field has been set.
func (o *BaseItemDtoImageBlurHashes) HasMenu() bool {
	if o != nil && !IsNil(o.Menu) {
		return true
	}

	return false
}

// SetMenu gets a reference to the given map[string]string and assigns it to the Menu field.
func (o *BaseItemDtoImageBlurHashes) SetMenu(v map[string]string) {
	o.Menu = &v
}

// GetChapter returns the Chapter field value if set, zero value otherwise.
func (o *BaseItemDtoImageBlurHashes) GetChapter() map[string]string {
	if o == nil || IsNil(o.Chapter) {
		var ret map[string]string
		return ret
	}
	return *o.Chapter
}

// GetChapterOk returns a tuple with the Chapter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItemDtoImageBlurHashes) GetChapterOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Chapter) {
		return nil, false
	}
	return o.Chapter, true
}

// HasChapter returns a boolean if a field has been set.
func (o *BaseItemDtoImageBlurHashes) HasChapter() bool {
	if o != nil && !IsNil(o.Chapter) {
		return true
	}

	return false
}

// SetChapter gets a reference to the given map[string]string and assigns it to the Chapter field.
func (o *BaseItemDtoImageBlurHashes) SetChapter(v map[string]string) {
	o.Chapter = &v
}

// GetBoxRear returns the BoxRear field value if set, zero value otherwise.
func (o *BaseItemDtoImageBlurHashes) GetBoxRear() map[string]string {
	if o == nil || IsNil(o.BoxRear) {
		var ret map[string]string
		return ret
	}
	return *o.BoxRear
}

// GetBoxRearOk returns a tuple with the BoxRear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItemDtoImageBlurHashes) GetBoxRearOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.BoxRear) {
		return nil, false
	}
	return o.BoxRear, true
}

// HasBoxRear returns a boolean if a field has been set.
func (o *BaseItemDtoImageBlurHashes) HasBoxRear() bool {
	if o != nil && !IsNil(o.BoxRear) {
		return true
	}

	return false
}

// SetBoxRear gets a reference to the given map[string]string and assigns it to the BoxRear field.
func (o *BaseItemDtoImageBlurHashes) SetBoxRear(v map[string]string) {
	o.BoxRear = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *BaseItemDtoImageBlurHashes) GetProfile() map[string]string {
	if o == nil || IsNil(o.Profile) {
		var ret map[string]string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseItemDtoImageBlurHashes) GetProfileOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *BaseItemDtoImageBlurHashes) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given map[string]string and assigns it to the Profile field.
func (o *BaseItemDtoImageBlurHashes) SetProfile(v map[string]string) {
	o.Profile = &v
}

func (o BaseItemDtoImageBlurHashes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseItemDtoImageBlurHashes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Primary) {
		toSerialize["Primary"] = o.Primary
	}
	if !IsNil(o.Art) {
		toSerialize["Art"] = o.Art
	}
	if !IsNil(o.Backdrop) {
		toSerialize["Backdrop"] = o.Backdrop
	}
	if !IsNil(o.Banner) {
		toSerialize["Banner"] = o.Banner
	}
	if !IsNil(o.Logo) {
		toSerialize["Logo"] = o.Logo
	}
	if !IsNil(o.Thumb) {
		toSerialize["Thumb"] = o.Thumb
	}
	if !IsNil(o.Disc) {
		toSerialize["Disc"] = o.Disc
	}
	if !IsNil(o.Box) {
		toSerialize["Box"] = o.Box
	}
	if !IsNil(o.Screenshot) {
		toSerialize["Screenshot"] = o.Screenshot
	}
	if !IsNil(o.Menu) {
		toSerialize["Menu"] = o.Menu
	}
	if !IsNil(o.Chapter) {
		toSerialize["Chapter"] = o.Chapter
	}
	if !IsNil(o.BoxRear) {
		toSerialize["BoxRear"] = o.BoxRear
	}
	if !IsNil(o.Profile) {
		toSerialize["Profile"] = o.Profile
	}
	return toSerialize, nil
}

type NullableBaseItemDtoImageBlurHashes struct {
	value *BaseItemDtoImageBlurHashes
	isSet bool
}

func (v NullableBaseItemDtoImageBlurHashes) Get() *BaseItemDtoImageBlurHashes {
	return v.value
}

func (v *NullableBaseItemDtoImageBlurHashes) Set(val *BaseItemDtoImageBlurHashes) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseItemDtoImageBlurHashes) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseItemDtoImageBlurHashes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseItemDtoImageBlurHashes(val *BaseItemDtoImageBlurHashes) *NullableBaseItemDtoImageBlurHashes {
	return &NullableBaseItemDtoImageBlurHashes{value: val, isSet: true}
}

func (v NullableBaseItemDtoImageBlurHashes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseItemDtoImageBlurHashes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


