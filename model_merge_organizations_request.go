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

// checks if the MergeOrganizationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MergeOrganizationsRequest{}

// MergeOrganizationsRequest struct for MergeOrganizationsRequest
type MergeOrganizationsRequest struct {
	// The ID of the organization that the organization will be merged with
	MergeWithId int32 `json:"merge_with_id"`
}

type _MergeOrganizationsRequest MergeOrganizationsRequest

// NewMergeOrganizationsRequest instantiates a new MergeOrganizationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMergeOrganizationsRequest(mergeWithId int32) *MergeOrganizationsRequest {
	this := MergeOrganizationsRequest{}
	this.MergeWithId = mergeWithId
	return &this
}

// NewMergeOrganizationsRequestWithDefaults instantiates a new MergeOrganizationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMergeOrganizationsRequestWithDefaults() *MergeOrganizationsRequest {
	this := MergeOrganizationsRequest{}
	return &this
}

// GetMergeWithId returns the MergeWithId field value
func (o *MergeOrganizationsRequest) GetMergeWithId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MergeWithId
}

// GetMergeWithIdOk returns a tuple with the MergeWithId field value
// and a boolean to check if the value has been set.
func (o *MergeOrganizationsRequest) GetMergeWithIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MergeWithId, true
}

// SetMergeWithId sets field value
func (o *MergeOrganizationsRequest) SetMergeWithId(v int32) {
	o.MergeWithId = v
}

func (o MergeOrganizationsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MergeOrganizationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["merge_with_id"] = o.MergeWithId
	return toSerialize, nil
}

func (o *MergeOrganizationsRequest) UnmarshalJSON(data []byte) (err error) {
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

	varMergeOrganizationsRequest := _MergeOrganizationsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMergeOrganizationsRequest)

	if err != nil {
		return err
	}

	*o = MergeOrganizationsRequest(varMergeOrganizationsRequest)

	return err
}

type NullableMergeOrganizationsRequest struct {
	value *MergeOrganizationsRequest
	isSet bool
}

func (v NullableMergeOrganizationsRequest) Get() *MergeOrganizationsRequest {
	return v.value
}

func (v *NullableMergeOrganizationsRequest) Set(val *MergeOrganizationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMergeOrganizationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMergeOrganizationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMergeOrganizationsRequest(val *MergeOrganizationsRequest) *NullableMergeOrganizationsRequest {
	return &NullableMergeOrganizationsRequest{value: val, isSet: true}
}

func (v NullableMergeOrganizationsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMergeOrganizationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


