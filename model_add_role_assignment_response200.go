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

// checks if the AddRoleAssignmentResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddRoleAssignmentResponse200{}

// AddRoleAssignmentResponse200 struct for AddRoleAssignmentResponse200
type AddRoleAssignmentResponse200 struct {
	// If the response is successful or not
	Success *bool `json:"success,omitempty"`
	Data *AddRoleAssignmentResponse200AllOfData `json:"data,omitempty"`
}

// NewAddRoleAssignmentResponse200 instantiates a new AddRoleAssignmentResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddRoleAssignmentResponse200() *AddRoleAssignmentResponse200 {
	this := AddRoleAssignmentResponse200{}
	return &this
}

// NewAddRoleAssignmentResponse200WithDefaults instantiates a new AddRoleAssignmentResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddRoleAssignmentResponse200WithDefaults() *AddRoleAssignmentResponse200 {
	this := AddRoleAssignmentResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *AddRoleAssignmentResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddRoleAssignmentResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *AddRoleAssignmentResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *AddRoleAssignmentResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AddRoleAssignmentResponse200) GetData() AddRoleAssignmentResponse200AllOfData {
	if o == nil || IsNil(o.Data) {
		var ret AddRoleAssignmentResponse200AllOfData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddRoleAssignmentResponse200) GetDataOk() (*AddRoleAssignmentResponse200AllOfData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AddRoleAssignmentResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AddRoleAssignmentResponse200AllOfData and assigns it to the Data field.
func (o *AddRoleAssignmentResponse200) SetData(v AddRoleAssignmentResponse200AllOfData) {
	o.Data = &v
}

func (o AddRoleAssignmentResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddRoleAssignmentResponse200) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAddRoleAssignmentResponse200 struct {
	value *AddRoleAssignmentResponse200
	isSet bool
}

func (v NullableAddRoleAssignmentResponse200) Get() *AddRoleAssignmentResponse200 {
	return v.value
}

func (v *NullableAddRoleAssignmentResponse200) Set(val *AddRoleAssignmentResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableAddRoleAssignmentResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableAddRoleAssignmentResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddRoleAssignmentResponse200(val *AddRoleAssignmentResponse200) *NullableAddRoleAssignmentResponse200 {
	return &NullableAddRoleAssignmentResponse200{value: val, isSet: true}
}

func (v NullableAddRoleAssignmentResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddRoleAssignmentResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


