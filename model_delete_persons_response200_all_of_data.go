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

// checks if the DeletePersonsResponse200AllOfData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeletePersonsResponse200AllOfData{}

// DeletePersonsResponse200AllOfData struct for DeletePersonsResponse200AllOfData
type DeletePersonsResponse200AllOfData struct {
	// The list of deleted persons IDs
	Id []int32 `json:"id,omitempty"`
}

// NewDeletePersonsResponse200AllOfData instantiates a new DeletePersonsResponse200AllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeletePersonsResponse200AllOfData() *DeletePersonsResponse200AllOfData {
	this := DeletePersonsResponse200AllOfData{}
	return &this
}

// NewDeletePersonsResponse200AllOfDataWithDefaults instantiates a new DeletePersonsResponse200AllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeletePersonsResponse200AllOfDataWithDefaults() *DeletePersonsResponse200AllOfData {
	this := DeletePersonsResponse200AllOfData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeletePersonsResponse200AllOfData) GetId() []int32 {
	if o == nil || IsNil(o.Id) {
		var ret []int32
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletePersonsResponse200AllOfData) GetIdOk() ([]int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeletePersonsResponse200AllOfData) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given []int32 and assigns it to the Id field.
func (o *DeletePersonsResponse200AllOfData) SetId(v []int32) {
	o.Id = v
}

func (o DeletePersonsResponse200AllOfData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeletePersonsResponse200AllOfData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableDeletePersonsResponse200AllOfData struct {
	value *DeletePersonsResponse200AllOfData
	isSet bool
}

func (v NullableDeletePersonsResponse200AllOfData) Get() *DeletePersonsResponse200AllOfData {
	return v.value
}

func (v *NullableDeletePersonsResponse200AllOfData) Set(val *DeletePersonsResponse200AllOfData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeletePersonsResponse200AllOfData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeletePersonsResponse200AllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeletePersonsResponse200AllOfData(val *DeletePersonsResponse200AllOfData) *NullableDeletePersonsResponse200AllOfData {
	return &NullableDeletePersonsResponse200AllOfData{value: val, isSet: true}
}

func (v NullableDeletePersonsResponse200AllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeletePersonsResponse200AllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


