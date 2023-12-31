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

// checks if the DeleteOrganizationFollowerResponse200Data type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteOrganizationFollowerResponse200Data{}

// DeleteOrganizationFollowerResponse200Data struct for DeleteOrganizationFollowerResponse200Data
type DeleteOrganizationFollowerResponse200Data struct {
	// The ID of the follower that was deleted from the organization
	Id *int32 `json:"id,omitempty"`
}

// NewDeleteOrganizationFollowerResponse200Data instantiates a new DeleteOrganizationFollowerResponse200Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteOrganizationFollowerResponse200Data() *DeleteOrganizationFollowerResponse200Data {
	this := DeleteOrganizationFollowerResponse200Data{}
	return &this
}

// NewDeleteOrganizationFollowerResponse200DataWithDefaults instantiates a new DeleteOrganizationFollowerResponse200Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteOrganizationFollowerResponse200DataWithDefaults() *DeleteOrganizationFollowerResponse200Data {
	this := DeleteOrganizationFollowerResponse200Data{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeleteOrganizationFollowerResponse200Data) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteOrganizationFollowerResponse200Data) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeleteOrganizationFollowerResponse200Data) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *DeleteOrganizationFollowerResponse200Data) SetId(v int32) {
	o.Id = &v
}

func (o DeleteOrganizationFollowerResponse200Data) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteOrganizationFollowerResponse200Data) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableDeleteOrganizationFollowerResponse200Data struct {
	value *DeleteOrganizationFollowerResponse200Data
	isSet bool
}

func (v NullableDeleteOrganizationFollowerResponse200Data) Get() *DeleteOrganizationFollowerResponse200Data {
	return v.value
}

func (v *NullableDeleteOrganizationFollowerResponse200Data) Set(val *DeleteOrganizationFollowerResponse200Data) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteOrganizationFollowerResponse200Data) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteOrganizationFollowerResponse200Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteOrganizationFollowerResponse200Data(val *DeleteOrganizationFollowerResponse200Data) *NullableDeleteOrganizationFollowerResponse200Data {
	return &NullableDeleteOrganizationFollowerResponse200Data{value: val, isSet: true}
}

func (v NullableDeleteOrganizationFollowerResponse200Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteOrganizationFollowerResponse200Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


