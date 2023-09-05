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

// checks if the DeleteOrganizationsResponse200Data type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteOrganizationsResponse200Data{}

// DeleteOrganizationsResponse200Data struct for DeleteOrganizationsResponse200Data
type DeleteOrganizationsResponse200Data struct {
	// The IDs of the organizations that were deleted
	Id []float32 `json:"id,omitempty"`
}

// NewDeleteOrganizationsResponse200Data instantiates a new DeleteOrganizationsResponse200Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteOrganizationsResponse200Data() *DeleteOrganizationsResponse200Data {
	this := DeleteOrganizationsResponse200Data{}
	return &this
}

// NewDeleteOrganizationsResponse200DataWithDefaults instantiates a new DeleteOrganizationsResponse200Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteOrganizationsResponse200DataWithDefaults() *DeleteOrganizationsResponse200Data {
	this := DeleteOrganizationsResponse200Data{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeleteOrganizationsResponse200Data) GetId() []float32 {
	if o == nil || IsNil(o.Id) {
		var ret []float32
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteOrganizationsResponse200Data) GetIdOk() ([]float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeleteOrganizationsResponse200Data) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given []float32 and assigns it to the Id field.
func (o *DeleteOrganizationsResponse200Data) SetId(v []float32) {
	o.Id = v
}

func (o DeleteOrganizationsResponse200Data) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteOrganizationsResponse200Data) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableDeleteOrganizationsResponse200Data struct {
	value *DeleteOrganizationsResponse200Data
	isSet bool
}

func (v NullableDeleteOrganizationsResponse200Data) Get() *DeleteOrganizationsResponse200Data {
	return v.value
}

func (v *NullableDeleteOrganizationsResponse200Data) Set(val *DeleteOrganizationsResponse200Data) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteOrganizationsResponse200Data) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteOrganizationsResponse200Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteOrganizationsResponse200Data(val *DeleteOrganizationsResponse200Data) *NullableDeleteOrganizationsResponse200Data {
	return &NullableDeleteOrganizationsResponse200Data{value: val, isSet: true}
}

func (v NullableDeleteOrganizationsResponse200Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteOrganizationsResponse200Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


