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

// checks if the DeleteConversation403ResponseAdditionalData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteConversation403ResponseAdditionalData{}

// DeleteConversation403ResponseAdditionalData struct for DeleteConversation403ResponseAdditionalData
type DeleteConversation403ResponseAdditionalData struct {
	// An error code sent by the API
	Code *string `json:"code,omitempty"`
}

// NewDeleteConversation403ResponseAdditionalData instantiates a new DeleteConversation403ResponseAdditionalData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteConversation403ResponseAdditionalData() *DeleteConversation403ResponseAdditionalData {
	this := DeleteConversation403ResponseAdditionalData{}
	return &this
}

// NewDeleteConversation403ResponseAdditionalDataWithDefaults instantiates a new DeleteConversation403ResponseAdditionalData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteConversation403ResponseAdditionalDataWithDefaults() *DeleteConversation403ResponseAdditionalData {
	this := DeleteConversation403ResponseAdditionalData{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *DeleteConversation403ResponseAdditionalData) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteConversation403ResponseAdditionalData) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *DeleteConversation403ResponseAdditionalData) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *DeleteConversation403ResponseAdditionalData) SetCode(v string) {
	o.Code = &v
}

func (o DeleteConversation403ResponseAdditionalData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteConversation403ResponseAdditionalData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return toSerialize, nil
}

type NullableDeleteConversation403ResponseAdditionalData struct {
	value *DeleteConversation403ResponseAdditionalData
	isSet bool
}

func (v NullableDeleteConversation403ResponseAdditionalData) Get() *DeleteConversation403ResponseAdditionalData {
	return v.value
}

func (v *NullableDeleteConversation403ResponseAdditionalData) Set(val *DeleteConversation403ResponseAdditionalData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteConversation403ResponseAdditionalData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteConversation403ResponseAdditionalData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteConversation403ResponseAdditionalData(val *DeleteConversation403ResponseAdditionalData) *NullableDeleteConversation403ResponseAdditionalData {
	return &NullableDeleteConversation403ResponseAdditionalData{value: val, isSet: true}
}

func (v NullableDeleteConversation403ResponseAdditionalData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteConversation403ResponseAdditionalData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


