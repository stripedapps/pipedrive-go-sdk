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

// checks if the GetActivitiesResponse200AdditionalData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetActivitiesResponse200AdditionalData{}

// GetActivitiesResponse200AdditionalData struct for GetActivitiesResponse200AdditionalData
type GetActivitiesResponse200AdditionalData struct {
	Pagination *GetActivitiesResponse200AdditionalDataPagination `json:"pagination,omitempty"`
}

// NewGetActivitiesResponse200AdditionalData instantiates a new GetActivitiesResponse200AdditionalData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetActivitiesResponse200AdditionalData() *GetActivitiesResponse200AdditionalData {
	this := GetActivitiesResponse200AdditionalData{}
	return &this
}

// NewGetActivitiesResponse200AdditionalDataWithDefaults instantiates a new GetActivitiesResponse200AdditionalData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetActivitiesResponse200AdditionalDataWithDefaults() *GetActivitiesResponse200AdditionalData {
	this := GetActivitiesResponse200AdditionalData{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *GetActivitiesResponse200AdditionalData) GetPagination() GetActivitiesResponse200AdditionalDataPagination {
	if o == nil || IsNil(o.Pagination) {
		var ret GetActivitiesResponse200AdditionalDataPagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetActivitiesResponse200AdditionalData) GetPaginationOk() (*GetActivitiesResponse200AdditionalDataPagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *GetActivitiesResponse200AdditionalData) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given GetActivitiesResponse200AdditionalDataPagination and assigns it to the Pagination field.
func (o *GetActivitiesResponse200AdditionalData) SetPagination(v GetActivitiesResponse200AdditionalDataPagination) {
	o.Pagination = &v
}

func (o GetActivitiesResponse200AdditionalData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetActivitiesResponse200AdditionalData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

type NullableGetActivitiesResponse200AdditionalData struct {
	value *GetActivitiesResponse200AdditionalData
	isSet bool
}

func (v NullableGetActivitiesResponse200AdditionalData) Get() *GetActivitiesResponse200AdditionalData {
	return v.value
}

func (v *NullableGetActivitiesResponse200AdditionalData) Set(val *GetActivitiesResponse200AdditionalData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetActivitiesResponse200AdditionalData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetActivitiesResponse200AdditionalData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetActivitiesResponse200AdditionalData(val *GetActivitiesResponse200AdditionalData) *NullableGetActivitiesResponse200AdditionalData {
	return &NullableGetActivitiesResponse200AdditionalData{value: val, isSet: true}
}

func (v NullableGetActivitiesResponse200AdditionalData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetActivitiesResponse200AdditionalData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


