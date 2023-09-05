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

// checks if the GetPersonProductsResponse200AllOfDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPersonProductsResponse200AllOfDataInner{}

// GetPersonProductsResponse200AllOfDataInner struct for GetPersonProductsResponse200AllOfDataInner
type GetPersonProductsResponse200AllOfDataInner struct {
	DEAL_ID *GetPersonProductsResponse200AllOfDataInnerDEALID `json:"DEAL_ID,omitempty"`
}

// NewGetPersonProductsResponse200AllOfDataInner instantiates a new GetPersonProductsResponse200AllOfDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPersonProductsResponse200AllOfDataInner() *GetPersonProductsResponse200AllOfDataInner {
	this := GetPersonProductsResponse200AllOfDataInner{}
	return &this
}

// NewGetPersonProductsResponse200AllOfDataInnerWithDefaults instantiates a new GetPersonProductsResponse200AllOfDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPersonProductsResponse200AllOfDataInnerWithDefaults() *GetPersonProductsResponse200AllOfDataInner {
	this := GetPersonProductsResponse200AllOfDataInner{}
	return &this
}

// GetDEAL_ID returns the DEAL_ID field value if set, zero value otherwise.
func (o *GetPersonProductsResponse200AllOfDataInner) GetDEAL_ID() GetPersonProductsResponse200AllOfDataInnerDEALID {
	if o == nil || IsNil(o.DEAL_ID) {
		var ret GetPersonProductsResponse200AllOfDataInnerDEALID
		return ret
	}
	return *o.DEAL_ID
}

// GetDEAL_IDOk returns a tuple with the DEAL_ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonProductsResponse200AllOfDataInner) GetDEAL_IDOk() (*GetPersonProductsResponse200AllOfDataInnerDEALID, bool) {
	if o == nil || IsNil(o.DEAL_ID) {
		return nil, false
	}
	return o.DEAL_ID, true
}

// HasDEAL_ID returns a boolean if a field has been set.
func (o *GetPersonProductsResponse200AllOfDataInner) HasDEAL_ID() bool {
	if o != nil && !IsNil(o.DEAL_ID) {
		return true
	}

	return false
}

// SetDEAL_ID gets a reference to the given GetPersonProductsResponse200AllOfDataInnerDEALID and assigns it to the DEAL_ID field.
func (o *GetPersonProductsResponse200AllOfDataInner) SetDEAL_ID(v GetPersonProductsResponse200AllOfDataInnerDEALID) {
	o.DEAL_ID = &v
}

func (o GetPersonProductsResponse200AllOfDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPersonProductsResponse200AllOfDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DEAL_ID) {
		toSerialize["DEAL_ID"] = o.DEAL_ID
	}
	return toSerialize, nil
}

type NullableGetPersonProductsResponse200AllOfDataInner struct {
	value *GetPersonProductsResponse200AllOfDataInner
	isSet bool
}

func (v NullableGetPersonProductsResponse200AllOfDataInner) Get() *GetPersonProductsResponse200AllOfDataInner {
	return v.value
}

func (v *NullableGetPersonProductsResponse200AllOfDataInner) Set(val *GetPersonProductsResponse200AllOfDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPersonProductsResponse200AllOfDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPersonProductsResponse200AllOfDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPersonProductsResponse200AllOfDataInner(val *GetPersonProductsResponse200AllOfDataInner) *NullableGetPersonProductsResponse200AllOfDataInner {
	return &NullableGetPersonProductsResponse200AllOfDataInner{value: val, isSet: true}
}

func (v NullableGetPersonProductsResponse200AllOfDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPersonProductsResponse200AllOfDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


