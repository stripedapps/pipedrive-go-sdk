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

// checks if the GetOrganizationsResponse200AllOfRelatedObjectsPicture type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationsResponse200AllOfRelatedObjectsPicture{}

// GetOrganizationsResponse200AllOfRelatedObjectsPicture The picture that is associated with the item
type GetOrganizationsResponse200AllOfRelatedObjectsPicture struct {
	PICTURE_ID *GetOrganizationsResponse200AllOfRelatedObjectsPicturePICTUREID `json:"PICTURE_ID,omitempty"`
}

// NewGetOrganizationsResponse200AllOfRelatedObjectsPicture instantiates a new GetOrganizationsResponse200AllOfRelatedObjectsPicture object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationsResponse200AllOfRelatedObjectsPicture() *GetOrganizationsResponse200AllOfRelatedObjectsPicture {
	this := GetOrganizationsResponse200AllOfRelatedObjectsPicture{}
	return &this
}

// NewGetOrganizationsResponse200AllOfRelatedObjectsPictureWithDefaults instantiates a new GetOrganizationsResponse200AllOfRelatedObjectsPicture object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationsResponse200AllOfRelatedObjectsPictureWithDefaults() *GetOrganizationsResponse200AllOfRelatedObjectsPicture {
	this := GetOrganizationsResponse200AllOfRelatedObjectsPicture{}
	return &this
}

// GetPICTURE_ID returns the PICTURE_ID field value if set, zero value otherwise.
func (o *GetOrganizationsResponse200AllOfRelatedObjectsPicture) GetPICTURE_ID() GetOrganizationsResponse200AllOfRelatedObjectsPicturePICTUREID {
	if o == nil || IsNil(o.PICTURE_ID) {
		var ret GetOrganizationsResponse200AllOfRelatedObjectsPicturePICTUREID
		return ret
	}
	return *o.PICTURE_ID
}

// GetPICTURE_IDOk returns a tuple with the PICTURE_ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationsResponse200AllOfRelatedObjectsPicture) GetPICTURE_IDOk() (*GetOrganizationsResponse200AllOfRelatedObjectsPicturePICTUREID, bool) {
	if o == nil || IsNil(o.PICTURE_ID) {
		return nil, false
	}
	return o.PICTURE_ID, true
}

// HasPICTURE_ID returns a boolean if a field has been set.
func (o *GetOrganizationsResponse200AllOfRelatedObjectsPicture) HasPICTURE_ID() bool {
	if o != nil && !IsNil(o.PICTURE_ID) {
		return true
	}

	return false
}

// SetPICTURE_ID gets a reference to the given GetOrganizationsResponse200AllOfRelatedObjectsPicturePICTUREID and assigns it to the PICTURE_ID field.
func (o *GetOrganizationsResponse200AllOfRelatedObjectsPicture) SetPICTURE_ID(v GetOrganizationsResponse200AllOfRelatedObjectsPicturePICTUREID) {
	o.PICTURE_ID = &v
}

func (o GetOrganizationsResponse200AllOfRelatedObjectsPicture) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationsResponse200AllOfRelatedObjectsPicture) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PICTURE_ID) {
		toSerialize["PICTURE_ID"] = o.PICTURE_ID
	}
	return toSerialize, nil
}

type NullableGetOrganizationsResponse200AllOfRelatedObjectsPicture struct {
	value *GetOrganizationsResponse200AllOfRelatedObjectsPicture
	isSet bool
}

func (v NullableGetOrganizationsResponse200AllOfRelatedObjectsPicture) Get() *GetOrganizationsResponse200AllOfRelatedObjectsPicture {
	return v.value
}

func (v *NullableGetOrganizationsResponse200AllOfRelatedObjectsPicture) Set(val *GetOrganizationsResponse200AllOfRelatedObjectsPicture) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationsResponse200AllOfRelatedObjectsPicture) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationsResponse200AllOfRelatedObjectsPicture) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationsResponse200AllOfRelatedObjectsPicture(val *GetOrganizationsResponse200AllOfRelatedObjectsPicture) *NullableGetOrganizationsResponse200AllOfRelatedObjectsPicture {
	return &NullableGetOrganizationsResponse200AllOfRelatedObjectsPicture{value: val, isSet: true}
}

func (v NullableGetOrganizationsResponse200AllOfRelatedObjectsPicture) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationsResponse200AllOfRelatedObjectsPicture) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


