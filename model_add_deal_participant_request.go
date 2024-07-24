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

// checks if the AddDealParticipantRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddDealParticipantRequest{}

// AddDealParticipantRequest struct for AddDealParticipantRequest
type AddDealParticipantRequest struct {
	// The ID of the person
	PersonId int32 `json:"person_id"`
}

type _AddDealParticipantRequest AddDealParticipantRequest

// NewAddDealParticipantRequest instantiates a new AddDealParticipantRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddDealParticipantRequest(personId int32) *AddDealParticipantRequest {
	this := AddDealParticipantRequest{}
	this.PersonId = personId
	return &this
}

// NewAddDealParticipantRequestWithDefaults instantiates a new AddDealParticipantRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddDealParticipantRequestWithDefaults() *AddDealParticipantRequest {
	this := AddDealParticipantRequest{}
	return &this
}

// GetPersonId returns the PersonId field value
func (o *AddDealParticipantRequest) GetPersonId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PersonId
}

// GetPersonIdOk returns a tuple with the PersonId field value
// and a boolean to check if the value has been set.
func (o *AddDealParticipantRequest) GetPersonIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PersonId, true
}

// SetPersonId sets field value
func (o *AddDealParticipantRequest) SetPersonId(v int32) {
	o.PersonId = v
}

func (o AddDealParticipantRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddDealParticipantRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["person_id"] = o.PersonId
	return toSerialize, nil
}

func (o *AddDealParticipantRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"person_id",
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

	varAddDealParticipantRequest := _AddDealParticipantRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddDealParticipantRequest)

	if err != nil {
		return err
	}

	*o = AddDealParticipantRequest(varAddDealParticipantRequest)

	return err
}

type NullableAddDealParticipantRequest struct {
	value *AddDealParticipantRequest
	isSet bool
}

func (v NullableAddDealParticipantRequest) Get() *AddDealParticipantRequest {
	return v.value
}

func (v *NullableAddDealParticipantRequest) Set(val *AddDealParticipantRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddDealParticipantRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddDealParticipantRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddDealParticipantRequest(val *AddDealParticipantRequest) *NullableAddDealParticipantRequest {
	return &NullableAddDealParticipantRequest{value: val, isSet: true}
}

func (v NullableAddDealParticipantRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddDealParticipantRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


