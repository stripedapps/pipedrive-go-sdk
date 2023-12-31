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

// checks if the UpdatePipelineResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePipelineResponse200{}

// UpdatePipelineResponse200 struct for UpdatePipelineResponse200
type UpdatePipelineResponse200 struct {
	// If the response is successful or not
	Success *bool `json:"success,omitempty"`
	Data *UpdatePipelineResponse200AllOfData `json:"data,omitempty"`
}

// NewUpdatePipelineResponse200 instantiates a new UpdatePipelineResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePipelineResponse200() *UpdatePipelineResponse200 {
	this := UpdatePipelineResponse200{}
	return &this
}

// NewUpdatePipelineResponse200WithDefaults instantiates a new UpdatePipelineResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePipelineResponse200WithDefaults() *UpdatePipelineResponse200 {
	this := UpdatePipelineResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *UpdatePipelineResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePipelineResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *UpdatePipelineResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *UpdatePipelineResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UpdatePipelineResponse200) GetData() UpdatePipelineResponse200AllOfData {
	if o == nil || IsNil(o.Data) {
		var ret UpdatePipelineResponse200AllOfData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePipelineResponse200) GetDataOk() (*UpdatePipelineResponse200AllOfData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UpdatePipelineResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given UpdatePipelineResponse200AllOfData and assigns it to the Data field.
func (o *UpdatePipelineResponse200) SetData(v UpdatePipelineResponse200AllOfData) {
	o.Data = &v
}

func (o UpdatePipelineResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePipelineResponse200) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableUpdatePipelineResponse200 struct {
	value *UpdatePipelineResponse200
	isSet bool
}

func (v NullableUpdatePipelineResponse200) Get() *UpdatePipelineResponse200 {
	return v.value
}

func (v *NullableUpdatePipelineResponse200) Set(val *UpdatePipelineResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePipelineResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePipelineResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePipelineResponse200(val *UpdatePipelineResponse200) *NullableUpdatePipelineResponse200 {
	return &NullableUpdatePipelineResponse200{value: val, isSet: true}
}

func (v NullableUpdatePipelineResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePipelineResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


