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

// checks if the UpdateMailThreadDetailsResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMailThreadDetailsResponse200{}

// UpdateMailThreadDetailsResponse200 struct for UpdateMailThreadDetailsResponse200
type UpdateMailThreadDetailsResponse200 struct {
	// If the response is successful or not
	Success *bool `json:"success,omitempty"`
	Data *BaseMailThread1 `json:"data,omitempty"`
}

// NewUpdateMailThreadDetailsResponse200 instantiates a new UpdateMailThreadDetailsResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMailThreadDetailsResponse200() *UpdateMailThreadDetailsResponse200 {
	this := UpdateMailThreadDetailsResponse200{}
	return &this
}

// NewUpdateMailThreadDetailsResponse200WithDefaults instantiates a new UpdateMailThreadDetailsResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMailThreadDetailsResponse200WithDefaults() *UpdateMailThreadDetailsResponse200 {
	this := UpdateMailThreadDetailsResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *UpdateMailThreadDetailsResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMailThreadDetailsResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *UpdateMailThreadDetailsResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *UpdateMailThreadDetailsResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UpdateMailThreadDetailsResponse200) GetData() BaseMailThread1 {
	if o == nil || IsNil(o.Data) {
		var ret BaseMailThread1
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMailThreadDetailsResponse200) GetDataOk() (*BaseMailThread1, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UpdateMailThreadDetailsResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given BaseMailThread1 and assigns it to the Data field.
func (o *UpdateMailThreadDetailsResponse200) SetData(v BaseMailThread1) {
	o.Data = &v
}

func (o UpdateMailThreadDetailsResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMailThreadDetailsResponse200) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableUpdateMailThreadDetailsResponse200 struct {
	value *UpdateMailThreadDetailsResponse200
	isSet bool
}

func (v NullableUpdateMailThreadDetailsResponse200) Get() *UpdateMailThreadDetailsResponse200 {
	return v.value
}

func (v *NullableUpdateMailThreadDetailsResponse200) Set(val *UpdateMailThreadDetailsResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMailThreadDetailsResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMailThreadDetailsResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMailThreadDetailsResponse200(val *UpdateMailThreadDetailsResponse200) *NullableUpdateMailThreadDetailsResponse200 {
	return &NullableUpdateMailThreadDetailsResponse200{value: val, isSet: true}
}

func (v NullableUpdateMailThreadDetailsResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMailThreadDetailsResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


