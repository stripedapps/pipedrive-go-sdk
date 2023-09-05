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

// checks if the GetUserSettingsResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserSettingsResponse200{}

// GetUserSettingsResponse200 struct for GetUserSettingsResponse200
type GetUserSettingsResponse200 struct {
	// If the response is successful or not
	Success *bool `json:"success,omitempty"`
	Data *GetUserSettingsResponse200AllOfData `json:"data,omitempty"`
}

// NewGetUserSettingsResponse200 instantiates a new GetUserSettingsResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserSettingsResponse200() *GetUserSettingsResponse200 {
	this := GetUserSettingsResponse200{}
	return &this
}

// NewGetUserSettingsResponse200WithDefaults instantiates a new GetUserSettingsResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserSettingsResponse200WithDefaults() *GetUserSettingsResponse200 {
	this := GetUserSettingsResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *GetUserSettingsResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserSettingsResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *GetUserSettingsResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *GetUserSettingsResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetUserSettingsResponse200) GetData() GetUserSettingsResponse200AllOfData {
	if o == nil || IsNil(o.Data) {
		var ret GetUserSettingsResponse200AllOfData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserSettingsResponse200) GetDataOk() (*GetUserSettingsResponse200AllOfData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetUserSettingsResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetUserSettingsResponse200AllOfData and assigns it to the Data field.
func (o *GetUserSettingsResponse200) SetData(v GetUserSettingsResponse200AllOfData) {
	o.Data = &v
}

func (o GetUserSettingsResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserSettingsResponse200) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetUserSettingsResponse200 struct {
	value *GetUserSettingsResponse200
	isSet bool
}

func (v NullableGetUserSettingsResponse200) Get() *GetUserSettingsResponse200 {
	return v.value
}

func (v *NullableGetUserSettingsResponse200) Set(val *GetUserSettingsResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserSettingsResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserSettingsResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserSettingsResponse200(val *GetUserSettingsResponse200) *NullableGetUserSettingsResponse200 {
	return &NullableGetUserSettingsResponse200{value: val, isSet: true}
}

func (v NullableGetUserSettingsResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserSettingsResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


