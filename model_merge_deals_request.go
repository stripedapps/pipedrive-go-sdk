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

// checks if the MergeDealsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MergeDealsRequest{}

// MergeDealsRequest struct for MergeDealsRequest
type MergeDealsRequest struct {
	// The ID of the deal that the deal will be merged with
	MergeWithId int32 `json:"merge_with_id"`
}

type _MergeDealsRequest MergeDealsRequest

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

func (o *MergeDealsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"merge_with_id",
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

	varMergeDealsRequest := _MergeDealsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMergeDealsRequest)

	if err != nil {
		return err
	}

	*o = MergeDealsRequest(varMergeDealsRequest)

	return err
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


