/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.10.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"fmt"
)

// MediaProtocol the model 'MediaProtocol'
type MediaProtocol string

// List of MediaProtocol
const (
	MEDIAPROTOCOL_FILE MediaProtocol = "File"
	MEDIAPROTOCOL_HTTP MediaProtocol = "Http"
	MEDIAPROTOCOL_RTMP MediaProtocol = "Rtmp"
	MEDIAPROTOCOL_RTSP MediaProtocol = "Rtsp"
	MEDIAPROTOCOL_UDP MediaProtocol = "Udp"
	MEDIAPROTOCOL_RTP MediaProtocol = "Rtp"
	MEDIAPROTOCOL_FTP MediaProtocol = "Ftp"
)

// All allowed values of MediaProtocol enum
var AllowedMediaProtocolEnumValues = []MediaProtocol{
	"File",
	"Http",
	"Rtmp",
	"Rtsp",
	"Udp",
	"Rtp",
	"Ftp",
}

func (v *MediaProtocol) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MediaProtocol(value)
	for _, existing := range AllowedMediaProtocolEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MediaProtocol", value)
}

// NewMediaProtocolFromValue returns a pointer to a valid MediaProtocol
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMediaProtocolFromValue(v string) (*MediaProtocol, error) {
	ev := MediaProtocol(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MediaProtocol: valid values are %v", v, AllowedMediaProtocolEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MediaProtocol) IsValid() bool {
	for _, existing := range AllowedMediaProtocolEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MediaProtocol value
func (v MediaProtocol) Ptr() *MediaProtocol {
	return &v
}

type NullableMediaProtocol struct {
	value *MediaProtocol
	isSet bool
}

func (v NullableMediaProtocol) Get() *MediaProtocol {
	return v.value
}

func (v *NullableMediaProtocol) Set(val *MediaProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaProtocol(val *MediaProtocol) *NullableMediaProtocol {
	return &NullableMediaProtocol{value: val, isSet: true}
}

func (v NullableMediaProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

