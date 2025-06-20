/*
Jellyfin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 10.11.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateUserByName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUserByName{}

// CreateUserByName The create user by name request body.
type CreateUserByName struct {
	// Gets or sets the username.
	Name string `json:"Name"`
	// Gets or sets the password.
	Password NullableString `json:"Password,omitempty"`
}

type _CreateUserByName CreateUserByName

// NewCreateUserByName instantiates a new CreateUserByName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUserByName(name string) *CreateUserByName {
	this := CreateUserByName{}
	this.Name = name
	return &this
}

// NewCreateUserByNameWithDefaults instantiates a new CreateUserByName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUserByNameWithDefaults() *CreateUserByName {
	this := CreateUserByName{}
	return &this
}

// GetName returns the Name field value
func (o *CreateUserByName) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateUserByName) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateUserByName) SetName(v string) {
	o.Name = v
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateUserByName) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateUserByName) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *CreateUserByName) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *CreateUserByName) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *CreateUserByName) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *CreateUserByName) UnsetPassword() {
	o.Password.Unset()
}

func (o CreateUserByName) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUserByName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Name"] = o.Name
	if o.Password.IsSet() {
		toSerialize["Password"] = o.Password.Get()
	}
	return toSerialize, nil
}

func (o *CreateUserByName) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateUserByName := _CreateUserByName{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateUserByName)

	if err != nil {
		return err
	}

	*o = CreateUserByName(varCreateUserByName)

	return err
}

type NullableCreateUserByName struct {
	value *CreateUserByName
	isSet bool
}

func (v NullableCreateUserByName) Get() *CreateUserByName {
	return v.value
}

func (v *NullableCreateUserByName) Set(val *CreateUserByName) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUserByName) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUserByName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUserByName(val *CreateUserByName) *NullableCreateUserByName {
	return &NullableCreateUserByName{value: val, isSet: true}
}

func (v NullableCreateUserByName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUserByName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


