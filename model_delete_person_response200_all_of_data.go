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

// checks if the DeletePersonResponse200AllOfData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeletePersonResponse200AllOfData{}

// DeletePersonResponse200AllOfData struct for DeletePersonResponse200AllOfData
type DeletePersonResponse200AllOfData struct {
	// The ID of the deleted person
	Id *int32 `json:"id,omitempty"`
}

// NewDeletePersonResponse200AllOfData instantiates a new DeletePersonResponse200AllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeletePersonResponse200AllOfData() *DeletePersonResponse200AllOfData {
	this := DeletePersonResponse200AllOfData{}
	return &this
}

// NewDeletePersonResponse200AllOfDataWithDefaults instantiates a new DeletePersonResponse200AllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeletePersonResponse200AllOfDataWithDefaults() *DeletePersonResponse200AllOfData {
	this := DeletePersonResponse200AllOfData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeletePersonResponse200AllOfData) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeletePersonResponse200AllOfData) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeletePersonResponse200AllOfData) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DeletePersonResponse200AllOfData) SetId(v int32) {
	o.Id = &v
}

func (o DeletePersonResponse200AllOfData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeletePersonResponse200AllOfData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableDeletePersonResponse200AllOfData struct {
	value *DeletePersonResponse200AllOfData
	isSet bool
}

func (v NullableDeletePersonResponse200AllOfData) Get() *DeletePersonResponse200AllOfData {
	return v.value
}

func (v *NullableDeletePersonResponse200AllOfData) Set(val *DeletePersonResponse200AllOfData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeletePersonResponse200AllOfData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeletePersonResponse200AllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeletePersonResponse200AllOfData(val *DeletePersonResponse200AllOfData) *NullableDeletePersonResponse200AllOfData {
	return &NullableDeletePersonResponse200AllOfData{value: val, isSet: true}
}

func (v NullableDeletePersonResponse200AllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeletePersonResponse200AllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


