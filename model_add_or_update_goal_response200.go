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

// checks if the AddOrUpdateGoalResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddOrUpdateGoalResponse200{}

// AddOrUpdateGoalResponse200 struct for AddOrUpdateGoalResponse200
type AddOrUpdateGoalResponse200 struct {
	// If the request was successful or not
	Success *bool `json:"success,omitempty"`
	Data *AddOrUpdateGoalResponse200Data `json:"data,omitempty"`
}

// NewAddOrUpdateGoalResponse200 instantiates a new AddOrUpdateGoalResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddOrUpdateGoalResponse200() *AddOrUpdateGoalResponse200 {
	this := AddOrUpdateGoalResponse200{}
	return &this
}

// NewAddOrUpdateGoalResponse200WithDefaults instantiates a new AddOrUpdateGoalResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddOrUpdateGoalResponse200WithDefaults() *AddOrUpdateGoalResponse200 {
	this := AddOrUpdateGoalResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *AddOrUpdateGoalResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOrUpdateGoalResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *AddOrUpdateGoalResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *AddOrUpdateGoalResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AddOrUpdateGoalResponse200) GetData() AddOrUpdateGoalResponse200Data {
	if o == nil || IsNil(o.Data) {
		var ret AddOrUpdateGoalResponse200Data
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOrUpdateGoalResponse200) GetDataOk() (*AddOrUpdateGoalResponse200Data, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AddOrUpdateGoalResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AddOrUpdateGoalResponse200Data and assigns it to the Data field.
func (o *AddOrUpdateGoalResponse200) SetData(v AddOrUpdateGoalResponse200Data) {
	o.Data = &v
}

func (o AddOrUpdateGoalResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddOrUpdateGoalResponse200) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAddOrUpdateGoalResponse200 struct {
	value *AddOrUpdateGoalResponse200
	isSet bool
}

func (v NullableAddOrUpdateGoalResponse200) Get() *AddOrUpdateGoalResponse200 {
	return v.value
}

func (v *NullableAddOrUpdateGoalResponse200) Set(val *AddOrUpdateGoalResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableAddOrUpdateGoalResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableAddOrUpdateGoalResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddOrUpdateGoalResponse200(val *AddOrUpdateGoalResponse200) *NullableAddOrUpdateGoalResponse200 {
	return &NullableAddOrUpdateGoalResponse200{value: val, isSet: true}
}

func (v NullableAddOrUpdateGoalResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddOrUpdateGoalResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


