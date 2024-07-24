/*
Pipedrive API v1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UserProviderLinkSuccessResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserProviderLinkSuccessResponseData{}

// UserProviderLinkSuccessResponseData struct for UserProviderLinkSuccessResponseData
type UserProviderLinkSuccessResponseData struct {
	// The success message of the request
	Message *string `json:"message,omitempty"`
}

// NewUserProviderLinkSuccessResponseData instantiates a new UserProviderLinkSuccessResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProviderLinkSuccessResponseData() *UserProviderLinkSuccessResponseData {
	this := UserProviderLinkSuccessResponseData{}
	return &this
}

// NewUserProviderLinkSuccessResponseDataWithDefaults instantiates a new UserProviderLinkSuccessResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserProviderLinkSuccessResponseDataWithDefaults() *UserProviderLinkSuccessResponseData {
	this := UserProviderLinkSuccessResponseData{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UserProviderLinkSuccessResponseData) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProviderLinkSuccessResponseData) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UserProviderLinkSuccessResponseData) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UserProviderLinkSuccessResponseData) SetMessage(v string) {
	o.Message = &v
}

func (o UserProviderLinkSuccessResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserProviderLinkSuccessResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableUserProviderLinkSuccessResponseData struct {
	value *UserProviderLinkSuccessResponseData
	isSet bool
}

func (v NullableUserProviderLinkSuccessResponseData) Get() *UserProviderLinkSuccessResponseData {
	return v.value
}

func (v *NullableUserProviderLinkSuccessResponseData) Set(val *UserProviderLinkSuccessResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProviderLinkSuccessResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProviderLinkSuccessResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProviderLinkSuccessResponseData(val *UserProviderLinkSuccessResponseData) *NullableUserProviderLinkSuccessResponseData {
	return &NullableUserProviderLinkSuccessResponseData{value: val, isSet: true}
}

func (v NullableUserProviderLinkSuccessResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProviderLinkSuccessResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


