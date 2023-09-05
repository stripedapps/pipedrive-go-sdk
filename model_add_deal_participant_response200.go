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

// checks if the AddDealParticipantResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddDealParticipantResponse200{}

// AddDealParticipantResponse200 struct for AddDealParticipantResponse200
type AddDealParticipantResponse200 struct {
	// If the request was successful or not
	Success *bool `json:"success,omitempty"`
	Data *AddDealParticipantResponse200Data `json:"data,omitempty"`
	RelatedObjects *GetDealsResponse200RelatedObjects `json:"related_objects,omitempty"`
}

// NewAddDealParticipantResponse200 instantiates a new AddDealParticipantResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddDealParticipantResponse200() *AddDealParticipantResponse200 {
	this := AddDealParticipantResponse200{}
	return &this
}

// NewAddDealParticipantResponse200WithDefaults instantiates a new AddDealParticipantResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddDealParticipantResponse200WithDefaults() *AddDealParticipantResponse200 {
	this := AddDealParticipantResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *AddDealParticipantResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDealParticipantResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *AddDealParticipantResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *AddDealParticipantResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AddDealParticipantResponse200) GetData() AddDealParticipantResponse200Data {
	if o == nil || IsNil(o.Data) {
		var ret AddDealParticipantResponse200Data
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDealParticipantResponse200) GetDataOk() (*AddDealParticipantResponse200Data, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AddDealParticipantResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AddDealParticipantResponse200Data and assigns it to the Data field.
func (o *AddDealParticipantResponse200) SetData(v AddDealParticipantResponse200Data) {
	o.Data = &v
}

// GetRelatedObjects returns the RelatedObjects field value if set, zero value otherwise.
func (o *AddDealParticipantResponse200) GetRelatedObjects() GetDealsResponse200RelatedObjects {
	if o == nil || IsNil(o.RelatedObjects) {
		var ret GetDealsResponse200RelatedObjects
		return ret
	}
	return *o.RelatedObjects
}

// GetRelatedObjectsOk returns a tuple with the RelatedObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDealParticipantResponse200) GetRelatedObjectsOk() (*GetDealsResponse200RelatedObjects, bool) {
	if o == nil || IsNil(o.RelatedObjects) {
		return nil, false
	}
	return o.RelatedObjects, true
}

// HasRelatedObjects returns a boolean if a field has been set.
func (o *AddDealParticipantResponse200) HasRelatedObjects() bool {
	if o != nil && !IsNil(o.RelatedObjects) {
		return true
	}

	return false
}

// SetRelatedObjects gets a reference to the given GetDealsResponse200RelatedObjects and assigns it to the RelatedObjects field.
func (o *AddDealParticipantResponse200) SetRelatedObjects(v GetDealsResponse200RelatedObjects) {
	o.RelatedObjects = &v
}

func (o AddDealParticipantResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddDealParticipantResponse200) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.RelatedObjects) {
		toSerialize["related_objects"] = o.RelatedObjects
	}
	return toSerialize, nil
}

type NullableAddDealParticipantResponse200 struct {
	value *AddDealParticipantResponse200
	isSet bool
}

func (v NullableAddDealParticipantResponse200) Get() *AddDealParticipantResponse200 {
	return v.value
}

func (v *NullableAddDealParticipantResponse200) Set(val *AddDealParticipantResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableAddDealParticipantResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableAddDealParticipantResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddDealParticipantResponse200(val *AddDealParticipantResponse200) *NullableAddDealParticipantResponse200 {
	return &NullableAddDealParticipantResponse200{value: val, isSet: true}
}

func (v NullableAddDealParticipantResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddDealParticipantResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


