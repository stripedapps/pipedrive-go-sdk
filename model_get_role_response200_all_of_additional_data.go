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

// checks if the GetRoleResponse200AllOfAdditionalData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRoleResponse200AllOfAdditionalData{}

// GetRoleResponse200AllOfAdditionalData The additional data in the role
type GetRoleResponse200AllOfAdditionalData struct {
	Settings *GetRoleResponse200AllOfAdditionalDataSettings `json:"settings,omitempty"`
}

// NewGetRoleResponse200AllOfAdditionalData instantiates a new GetRoleResponse200AllOfAdditionalData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRoleResponse200AllOfAdditionalData() *GetRoleResponse200AllOfAdditionalData {
	this := GetRoleResponse200AllOfAdditionalData{}
	return &this
}

// NewGetRoleResponse200AllOfAdditionalDataWithDefaults instantiates a new GetRoleResponse200AllOfAdditionalData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRoleResponse200AllOfAdditionalDataWithDefaults() *GetRoleResponse200AllOfAdditionalData {
	this := GetRoleResponse200AllOfAdditionalData{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *GetRoleResponse200AllOfAdditionalData) GetSettings() GetRoleResponse200AllOfAdditionalDataSettings {
	if o == nil || IsNil(o.Settings) {
		var ret GetRoleResponse200AllOfAdditionalDataSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoleResponse200AllOfAdditionalData) GetSettingsOk() (*GetRoleResponse200AllOfAdditionalDataSettings, bool) {
	if o == nil || IsNil(o.Settings) {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *GetRoleResponse200AllOfAdditionalData) HasSettings() bool {
	if o != nil && !IsNil(o.Settings) {
		return true
	}

	return false
}

// SetSettings gets a reference to the given GetRoleResponse200AllOfAdditionalDataSettings and assigns it to the Settings field.
func (o *GetRoleResponse200AllOfAdditionalData) SetSettings(v GetRoleResponse200AllOfAdditionalDataSettings) {
	o.Settings = &v
}

func (o GetRoleResponse200AllOfAdditionalData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRoleResponse200AllOfAdditionalData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Settings) {
		toSerialize["settings"] = o.Settings
	}
	return toSerialize, nil
}

type NullableGetRoleResponse200AllOfAdditionalData struct {
	value *GetRoleResponse200AllOfAdditionalData
	isSet bool
}

func (v NullableGetRoleResponse200AllOfAdditionalData) Get() *GetRoleResponse200AllOfAdditionalData {
	return v.value
}

func (v *NullableGetRoleResponse200AllOfAdditionalData) Set(val *GetRoleResponse200AllOfAdditionalData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRoleResponse200AllOfAdditionalData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRoleResponse200AllOfAdditionalData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRoleResponse200AllOfAdditionalData(val *GetRoleResponse200AllOfAdditionalData) *NullableGetRoleResponse200AllOfAdditionalData {
	return &NullableGetRoleResponse200AllOfAdditionalData{value: val, isSet: true}
}

func (v NullableGetRoleResponse200AllOfAdditionalData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRoleResponse200AllOfAdditionalData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


