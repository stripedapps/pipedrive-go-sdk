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

// checks if the AddActivityResponse200RelatedObjectsOrganization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddActivityResponse200RelatedObjectsOrganization{}

// AddActivityResponse200RelatedObjectsOrganization struct for AddActivityResponse200RelatedObjectsOrganization
type AddActivityResponse200RelatedObjectsOrganization struct {
	ORGANIZATION_ID *AddActivityResponse200RelatedObjectsOrganizationORGANIZATIONID `json:"ORGANIZATION_ID,omitempty"`
}

// NewAddActivityResponse200RelatedObjectsOrganization instantiates a new AddActivityResponse200RelatedObjectsOrganization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddActivityResponse200RelatedObjectsOrganization() *AddActivityResponse200RelatedObjectsOrganization {
	this := AddActivityResponse200RelatedObjectsOrganization{}
	return &this
}

// NewAddActivityResponse200RelatedObjectsOrganizationWithDefaults instantiates a new AddActivityResponse200RelatedObjectsOrganization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddActivityResponse200RelatedObjectsOrganizationWithDefaults() *AddActivityResponse200RelatedObjectsOrganization {
	this := AddActivityResponse200RelatedObjectsOrganization{}
	return &this
}

// GetORGANIZATION_ID returns the ORGANIZATION_ID field value if set, zero value otherwise.
func (o *AddActivityResponse200RelatedObjectsOrganization) GetORGANIZATION_ID() AddActivityResponse200RelatedObjectsOrganizationORGANIZATIONID {
	if o == nil || IsNil(o.ORGANIZATION_ID) {
		var ret AddActivityResponse200RelatedObjectsOrganizationORGANIZATIONID
		return ret
	}
	return *o.ORGANIZATION_ID
}

// GetORGANIZATION_IDOk returns a tuple with the ORGANIZATION_ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddActivityResponse200RelatedObjectsOrganization) GetORGANIZATION_IDOk() (*AddActivityResponse200RelatedObjectsOrganizationORGANIZATIONID, bool) {
	if o == nil || IsNil(o.ORGANIZATION_ID) {
		return nil, false
	}
	return o.ORGANIZATION_ID, true
}

// HasORGANIZATION_ID returns a boolean if a field has been set.
func (o *AddActivityResponse200RelatedObjectsOrganization) HasORGANIZATION_ID() bool {
	if o != nil && !IsNil(o.ORGANIZATION_ID) {
		return true
	}

	return false
}

// SetORGANIZATION_ID gets a reference to the given AddActivityResponse200RelatedObjectsOrganizationORGANIZATIONID and assigns it to the ORGANIZATION_ID field.
func (o *AddActivityResponse200RelatedObjectsOrganization) SetORGANIZATION_ID(v AddActivityResponse200RelatedObjectsOrganizationORGANIZATIONID) {
	o.ORGANIZATION_ID = &v
}

func (o AddActivityResponse200RelatedObjectsOrganization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddActivityResponse200RelatedObjectsOrganization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ORGANIZATION_ID) {
		toSerialize["ORGANIZATION_ID"] = o.ORGANIZATION_ID
	}
	return toSerialize, nil
}

type NullableAddActivityResponse200RelatedObjectsOrganization struct {
	value *AddActivityResponse200RelatedObjectsOrganization
	isSet bool
}

func (v NullableAddActivityResponse200RelatedObjectsOrganization) Get() *AddActivityResponse200RelatedObjectsOrganization {
	return v.value
}

func (v *NullableAddActivityResponse200RelatedObjectsOrganization) Set(val *AddActivityResponse200RelatedObjectsOrganization) {
	v.value = val
	v.isSet = true
}

func (v NullableAddActivityResponse200RelatedObjectsOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableAddActivityResponse200RelatedObjectsOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddActivityResponse200RelatedObjectsOrganization(val *AddActivityResponse200RelatedObjectsOrganization) *NullableAddActivityResponse200RelatedObjectsOrganization {
	return &NullableAddActivityResponse200RelatedObjectsOrganization{value: val, isSet: true}
}

func (v NullableAddActivityResponse200RelatedObjectsOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddActivityResponse200RelatedObjectsOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


