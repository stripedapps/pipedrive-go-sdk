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

// checks if the UserIds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserIds{}

// UserIds struct for UserIds
type UserIds struct {
	// If the response is successful or not
	Success *bool `json:"success,omitempty"`
	// The list of user IDs
	Data []int32 `json:"data,omitempty"`
}

// NewUserIds instantiates a new UserIds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserIds() *UserIds {
	this := UserIds{}
	return &this
}

// NewUserIdsWithDefaults instantiates a new UserIds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIdsWithDefaults() *UserIds {
	this := UserIds{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *UserIds) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIds) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *UserIds) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *UserIds) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UserIds) GetData() []int32 {
	if o == nil || IsNil(o.Data) {
		var ret []int32
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIds) GetDataOk() ([]int32, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UserIds) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []int32 and assigns it to the Data field.
func (o *UserIds) SetData(v []int32) {
	o.Data = v
}

func (o UserIds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserIds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableUserIds struct {
	value *UserIds
	isSet bool
}

func (v NullableUserIds) Get() *UserIds {
	return v.value
}

func (v *NullableUserIds) Set(val *UserIds) {
	v.value = val
	v.isSet = true
}

func (v NullableUserIds) IsSet() bool {
	return v.isSet
}

func (v *NullableUserIds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserIds(val *UserIds) *NullableUserIds {
	return &NullableUserIds{value: val, isSet: true}
}

func (v NullableUserIds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserIds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


