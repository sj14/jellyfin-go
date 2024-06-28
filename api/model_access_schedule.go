/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// checks if the AccessSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessSchedule{}

// AccessSchedule An entity representing a user's access schedule.
type AccessSchedule struct {
	// Gets the id of this instance.
	Id *int32 `json:"Id,omitempty"`
	// Gets the id of the associated user.
	UserId *string `json:"UserId,omitempty"`
	// Gets or sets the day of week.
	DayOfWeek *DynamicDayOfWeek `json:"DayOfWeek,omitempty"`
	// Gets or sets the start hour.
	StartHour *float64 `json:"StartHour,omitempty"`
	// Gets or sets the end hour.
	EndHour *float64 `json:"EndHour,omitempty"`
}

// NewAccessSchedule instantiates a new AccessSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessSchedule() *AccessSchedule {
	this := AccessSchedule{}
	return &this
}

// NewAccessScheduleWithDefaults instantiates a new AccessSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessScheduleWithDefaults() *AccessSchedule {
	this := AccessSchedule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccessSchedule) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessSchedule) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccessSchedule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AccessSchedule) SetId(v int32) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AccessSchedule) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessSchedule) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AccessSchedule) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *AccessSchedule) SetUserId(v string) {
	o.UserId = &v
}

// GetDayOfWeek returns the DayOfWeek field value if set, zero value otherwise.
func (o *AccessSchedule) GetDayOfWeek() DynamicDayOfWeek {
	if o == nil || IsNil(o.DayOfWeek) {
		var ret DynamicDayOfWeek
		return ret
	}
	return *o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessSchedule) GetDayOfWeekOk() (*DynamicDayOfWeek, bool) {
	if o == nil || IsNil(o.DayOfWeek) {
		return nil, false
	}
	return o.DayOfWeek, true
}

// HasDayOfWeek returns a boolean if a field has been set.
func (o *AccessSchedule) HasDayOfWeek() bool {
	if o != nil && !IsNil(o.DayOfWeek) {
		return true
	}

	return false
}

// SetDayOfWeek gets a reference to the given DynamicDayOfWeek and assigns it to the DayOfWeek field.
func (o *AccessSchedule) SetDayOfWeek(v DynamicDayOfWeek) {
	o.DayOfWeek = &v
}

// GetStartHour returns the StartHour field value if set, zero value otherwise.
func (o *AccessSchedule) GetStartHour() float64 {
	if o == nil || IsNil(o.StartHour) {
		var ret float64
		return ret
	}
	return *o.StartHour
}

// GetStartHourOk returns a tuple with the StartHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessSchedule) GetStartHourOk() (*float64, bool) {
	if o == nil || IsNil(o.StartHour) {
		return nil, false
	}
	return o.StartHour, true
}

// HasStartHour returns a boolean if a field has been set.
func (o *AccessSchedule) HasStartHour() bool {
	if o != nil && !IsNil(o.StartHour) {
		return true
	}

	return false
}

// SetStartHour gets a reference to the given float64 and assigns it to the StartHour field.
func (o *AccessSchedule) SetStartHour(v float64) {
	o.StartHour = &v
}

// GetEndHour returns the EndHour field value if set, zero value otherwise.
func (o *AccessSchedule) GetEndHour() float64 {
	if o == nil || IsNil(o.EndHour) {
		var ret float64
		return ret
	}
	return *o.EndHour
}

// GetEndHourOk returns a tuple with the EndHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessSchedule) GetEndHourOk() (*float64, bool) {
	if o == nil || IsNil(o.EndHour) {
		return nil, false
	}
	return o.EndHour, true
}

// HasEndHour returns a boolean if a field has been set.
func (o *AccessSchedule) HasEndHour() bool {
	if o != nil && !IsNil(o.EndHour) {
		return true
	}

	return false
}

// SetEndHour gets a reference to the given float64 and assigns it to the EndHour field.
func (o *AccessSchedule) SetEndHour(v float64) {
	o.EndHour = &v
}

func (o AccessSchedule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["UserId"] = o.UserId
	}
	if !IsNil(o.DayOfWeek) {
		toSerialize["DayOfWeek"] = o.DayOfWeek
	}
	if !IsNil(o.StartHour) {
		toSerialize["StartHour"] = o.StartHour
	}
	if !IsNil(o.EndHour) {
		toSerialize["EndHour"] = o.EndHour
	}
	return toSerialize, nil
}

type NullableAccessSchedule struct {
	value *AccessSchedule
	isSet bool
}

func (v NullableAccessSchedule) Get() *AccessSchedule {
	return v.value
}

func (v *NullableAccessSchedule) Set(val *AccessSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessSchedule(val *AccessSchedule) *NullableAccessSchedule {
	return &NullableAccessSchedule{value: val, isSet: true}
}

func (v NullableAccessSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


