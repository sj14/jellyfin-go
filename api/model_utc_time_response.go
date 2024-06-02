/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.9.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// checks if the UtcTimeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UtcTimeResponse{}

// UtcTimeResponse Class UtcTimeResponse.
type UtcTimeResponse struct {
	// Gets the UTC time when request has been received.
	RequestReceptionTime *time.Time `json:"RequestReceptionTime,omitempty"`
	// Gets the UTC time when response has been sent.
	ResponseTransmissionTime *time.Time `json:"ResponseTransmissionTime,omitempty"`
}

// NewUtcTimeResponse instantiates a new UtcTimeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUtcTimeResponse() *UtcTimeResponse {
	this := UtcTimeResponse{}
	return &this
}

// NewUtcTimeResponseWithDefaults instantiates a new UtcTimeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUtcTimeResponseWithDefaults() *UtcTimeResponse {
	this := UtcTimeResponse{}
	return &this
}

// GetRequestReceptionTime returns the RequestReceptionTime field value if set, zero value otherwise.
func (o *UtcTimeResponse) GetRequestReceptionTime() time.Time {
	if o == nil || IsNil(o.RequestReceptionTime) {
		var ret time.Time
		return ret
	}
	return *o.RequestReceptionTime
}

// GetRequestReceptionTimeOk returns a tuple with the RequestReceptionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtcTimeResponse) GetRequestReceptionTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RequestReceptionTime) {
		return nil, false
	}
	return o.RequestReceptionTime, true
}

// HasRequestReceptionTime returns a boolean if a field has been set.
func (o *UtcTimeResponse) HasRequestReceptionTime() bool {
	if o != nil && !IsNil(o.RequestReceptionTime) {
		return true
	}

	return false
}

// SetRequestReceptionTime gets a reference to the given time.Time and assigns it to the RequestReceptionTime field.
func (o *UtcTimeResponse) SetRequestReceptionTime(v time.Time) {
	o.RequestReceptionTime = &v
}

// GetResponseTransmissionTime returns the ResponseTransmissionTime field value if set, zero value otherwise.
func (o *UtcTimeResponse) GetResponseTransmissionTime() time.Time {
	if o == nil || IsNil(o.ResponseTransmissionTime) {
		var ret time.Time
		return ret
	}
	return *o.ResponseTransmissionTime
}

// GetResponseTransmissionTimeOk returns a tuple with the ResponseTransmissionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UtcTimeResponse) GetResponseTransmissionTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ResponseTransmissionTime) {
		return nil, false
	}
	return o.ResponseTransmissionTime, true
}

// HasResponseTransmissionTime returns a boolean if a field has been set.
func (o *UtcTimeResponse) HasResponseTransmissionTime() bool {
	if o != nil && !IsNil(o.ResponseTransmissionTime) {
		return true
	}

	return false
}

// SetResponseTransmissionTime gets a reference to the given time.Time and assigns it to the ResponseTransmissionTime field.
func (o *UtcTimeResponse) SetResponseTransmissionTime(v time.Time) {
	o.ResponseTransmissionTime = &v
}

func (o UtcTimeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UtcTimeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestReceptionTime) {
		toSerialize["RequestReceptionTime"] = o.RequestReceptionTime
	}
	if !IsNil(o.ResponseTransmissionTime) {
		toSerialize["ResponseTransmissionTime"] = o.ResponseTransmissionTime
	}
	return toSerialize, nil
}

type NullableUtcTimeResponse struct {
	value *UtcTimeResponse
	isSet bool
}

func (v NullableUtcTimeResponse) Get() *UtcTimeResponse {
	return v.value
}

func (v *NullableUtcTimeResponse) Set(val *UtcTimeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUtcTimeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUtcTimeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtcTimeResponse(val *UtcTimeResponse) *NullableUtcTimeResponse {
	return &NullableUtcTimeResponse{value: val, isSet: true}
}

func (v NullableUtcTimeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtcTimeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


