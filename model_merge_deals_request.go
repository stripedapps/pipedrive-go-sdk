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

// checks if the MergeDealsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MergeDealsRequest{}

// MergeDealsRequest struct for MergeDealsRequest
type MergeDealsRequest struct {
	// The ID of the deal that the deal will be merged with
	MergeWithId int32 `json:"merge_with_id"`
}

// NewMergeDealsRequest instantiates a new MergeDealsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMergeDealsRequest(mergeWithId int32) *MergeDealsRequest {
	this := MergeDealsRequest{}
	this.MergeWithId = mergeWithId
	return &this
}

// NewMergeDealsRequestWithDefaults instantiates a new MergeDealsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMergeDealsRequestWithDefaults() *MergeDealsRequest {
	this := MergeDealsRequest{}
	return &this
}

// GetMergeWithId returns the MergeWithId field value
func (o *MergeDealsRequest) GetMergeWithId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MergeWithId
}

// GetMergeWithIdOk returns a tuple with the MergeWithId field value
// and a boolean to check if the value has been set.
func (o *MergeDealsRequest) GetMergeWithIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MergeWithId, true
}

// SetMergeWithId sets field value
func (o *MergeDealsRequest) SetMergeWithId(v int32) {
	o.MergeWithId = v
}

func (o MergeDealsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MergeDealsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merge_with_id"] = o.MergeWithId
	return toSerialize, nil
}

type NullableMergeDealsRequest struct {
	value *MergeDealsRequest
	isSet bool
}

func (v NullableMergeDealsRequest) Get() *MergeDealsRequest {
	return v.value
}

func (v *NullableMergeDealsRequest) Set(val *MergeDealsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMergeDealsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMergeDealsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMergeDealsRequest(val *MergeDealsRequest) *NullableMergeDealsRequest {
	return &NullableMergeDealsRequest{value: val, isSet: true}
}

func (v NullableMergeDealsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMergeDealsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


