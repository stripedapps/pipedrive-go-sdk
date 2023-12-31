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

// checks if the AddOrUpdateGoalResponse200DataGoalDuration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddOrUpdateGoalResponse200DataGoalDuration{}

// AddOrUpdateGoalResponse200DataGoalDuration The duration of the goal
type AddOrUpdateGoalResponse200DataGoalDuration struct {
	// The start date of the goal
	Start *string `json:"start,omitempty"`
	// The end date of the goal
	End *string `json:"end,omitempty"`
}

// NewAddOrUpdateGoalResponse200DataGoalDuration instantiates a new AddOrUpdateGoalResponse200DataGoalDuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddOrUpdateGoalResponse200DataGoalDuration() *AddOrUpdateGoalResponse200DataGoalDuration {
	this := AddOrUpdateGoalResponse200DataGoalDuration{}
	return &this
}

// NewAddOrUpdateGoalResponse200DataGoalDurationWithDefaults instantiates a new AddOrUpdateGoalResponse200DataGoalDuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddOrUpdateGoalResponse200DataGoalDurationWithDefaults() *AddOrUpdateGoalResponse200DataGoalDuration {
	this := AddOrUpdateGoalResponse200DataGoalDuration{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *AddOrUpdateGoalResponse200DataGoalDuration) GetStart() string {
	if o == nil || IsNil(o.Start) {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOrUpdateGoalResponse200DataGoalDuration) GetStartOk() (*string, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *AddOrUpdateGoalResponse200DataGoalDuration) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *AddOrUpdateGoalResponse200DataGoalDuration) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *AddOrUpdateGoalResponse200DataGoalDuration) GetEnd() string {
	if o == nil || IsNil(o.End) {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddOrUpdateGoalResponse200DataGoalDuration) GetEndOk() (*string, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *AddOrUpdateGoalResponse200DataGoalDuration) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *AddOrUpdateGoalResponse200DataGoalDuration) SetEnd(v string) {
	o.End = &v
}

func (o AddOrUpdateGoalResponse200DataGoalDuration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddOrUpdateGoalResponse200DataGoalDuration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	return toSerialize, nil
}

type NullableAddOrUpdateGoalResponse200DataGoalDuration struct {
	value *AddOrUpdateGoalResponse200DataGoalDuration
	isSet bool
}

func (v NullableAddOrUpdateGoalResponse200DataGoalDuration) Get() *AddOrUpdateGoalResponse200DataGoalDuration {
	return v.value
}

func (v *NullableAddOrUpdateGoalResponse200DataGoalDuration) Set(val *AddOrUpdateGoalResponse200DataGoalDuration) {
	v.value = val
	v.isSet = true
}

func (v NullableAddOrUpdateGoalResponse200DataGoalDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableAddOrUpdateGoalResponse200DataGoalDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddOrUpdateGoalResponse200DataGoalDuration(val *AddOrUpdateGoalResponse200DataGoalDuration) *NullableAddOrUpdateGoalResponse200DataGoalDuration {
	return &NullableAddOrUpdateGoalResponse200DataGoalDuration{value: val, isSet: true}
}

func (v NullableAddOrUpdateGoalResponse200DataGoalDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddOrUpdateGoalResponse200DataGoalDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


