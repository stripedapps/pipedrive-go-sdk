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

// checks if the UserConnectionsResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserConnectionsResponse200{}

// UserConnectionsResponse200 struct for UserConnectionsResponse200
type UserConnectionsResponse200 struct {
	// If the response is successful or not
	Success *bool `json:"success,omitempty"`
	Data *UserConnectionsResponse200AllOfData `json:"data,omitempty"`
}

// NewUserConnectionsResponse200 instantiates a new UserConnectionsResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserConnectionsResponse200() *UserConnectionsResponse200 {
	this := UserConnectionsResponse200{}
	return &this
}

// NewUserConnectionsResponse200WithDefaults instantiates a new UserConnectionsResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserConnectionsResponse200WithDefaults() *UserConnectionsResponse200 {
	this := UserConnectionsResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *UserConnectionsResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConnectionsResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *UserConnectionsResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *UserConnectionsResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UserConnectionsResponse200) GetData() UserConnectionsResponse200AllOfData {
	if o == nil || IsNil(o.Data) {
		var ret UserConnectionsResponse200AllOfData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserConnectionsResponse200) GetDataOk() (*UserConnectionsResponse200AllOfData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UserConnectionsResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given UserConnectionsResponse200AllOfData and assigns it to the Data field.
func (o *UserConnectionsResponse200) SetData(v UserConnectionsResponse200AllOfData) {
	o.Data = &v
}

func (o UserConnectionsResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserConnectionsResponse200) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableUserConnectionsResponse200 struct {
	value *UserConnectionsResponse200
	isSet bool
}

func (v NullableUserConnectionsResponse200) Get() *UserConnectionsResponse200 {
	return v.value
}

func (v *NullableUserConnectionsResponse200) Set(val *UserConnectionsResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableUserConnectionsResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableUserConnectionsResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserConnectionsResponse200(val *UserConnectionsResponse200) *NullableUserConnectionsResponse200 {
	return &NullableUserConnectionsResponse200{value: val, isSet: true}
}

func (v NullableUserConnectionsResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserConnectionsResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


