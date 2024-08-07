/*
Pipedrive API v1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AddRoleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddRoleRequest{}

// AddRoleRequest The details of the role
type AddRoleRequest struct {
	// The name of the role
	Name string `json:"name"`
	// The ID of the parent role
	ParentRoleId *int32 `json:"parent_role_id,omitempty"`
}

type _AddRoleRequest AddRoleRequest

// NewAddRoleRequest instantiates a new AddRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddRoleRequest(name string) *AddRoleRequest {
	this := AddRoleRequest{}
	this.Name = name
	return &this
}

// NewAddRoleRequestWithDefaults instantiates a new AddRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddRoleRequestWithDefaults() *AddRoleRequest {
	this := AddRoleRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AddRoleRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddRoleRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddRoleRequest) SetName(v string) {
	o.Name = v
}

// GetParentRoleId returns the ParentRoleId field value if set, zero value otherwise.
func (o *AddRoleRequest) GetParentRoleId() int32 {
	if o == nil || IsNil(o.ParentRoleId) {
		var ret int32
		return ret
	}
	return *o.ParentRoleId
}

// GetParentRoleIdOk returns a tuple with the ParentRoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddRoleRequest) GetParentRoleIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ParentRoleId) {
		return nil, false
	}
	return o.ParentRoleId, true
}

// HasParentRoleId returns a boolean if a field has been set.
func (o *AddRoleRequest) HasParentRoleId() bool {
	if o != nil && !IsNil(o.ParentRoleId) {
		return true
	}

	return false
}

// SetParentRoleId gets a reference to the given int32 and assigns it to the ParentRoleId field.
func (o *AddRoleRequest) SetParentRoleId(v int32) {
	o.ParentRoleId = &v
}

func (o AddRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddRoleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.ParentRoleId) {
		toSerialize["parent_role_id"] = o.ParentRoleId
	}
	return toSerialize, nil
}

func (o *AddRoleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varAddRoleRequest := _AddRoleRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddRoleRequest)

	if err != nil {
		return err
	}

	*o = AddRoleRequest(varAddRoleRequest)

	return err
}

type NullableAddRoleRequest struct {
	value *AddRoleRequest
	isSet bool
}

func (v NullableAddRoleRequest) Get() *AddRoleRequest {
	return v.value
}

func (v *NullableAddRoleRequest) Set(val *AddRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddRoleRequest(val *AddRoleRequest) *NullableAddRoleRequest {
	return &NullableAddRoleRequest{value: val, isSet: true}
}

func (v NullableAddRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


