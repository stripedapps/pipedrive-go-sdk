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

// checks if the UpdateProjectResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProjectResponse200{}

// UpdateProjectResponse200 struct for UpdateProjectResponse200
type UpdateProjectResponse200 struct {
	Success *bool `json:"success,omitempty"`
	Data *ProjectResponseObject `json:"data,omitempty"`
	AdditionalData map[string]interface{} `json:"additional_data,omitempty"`
}

// NewUpdateProjectResponse200 instantiates a new UpdateProjectResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProjectResponse200() *UpdateProjectResponse200 {
	this := UpdateProjectResponse200{}
	return &this
}

// NewUpdateProjectResponse200WithDefaults instantiates a new UpdateProjectResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProjectResponse200WithDefaults() *UpdateProjectResponse200 {
	this := UpdateProjectResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *UpdateProjectResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProjectResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *UpdateProjectResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *UpdateProjectResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UpdateProjectResponse200) GetData() ProjectResponseObject {
	if o == nil || IsNil(o.Data) {
		var ret ProjectResponseObject
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProjectResponse200) GetDataOk() (*ProjectResponseObject, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UpdateProjectResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ProjectResponseObject and assigns it to the Data field.
func (o *UpdateProjectResponse200) SetData(v ProjectResponseObject) {
	o.Data = &v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateProjectResponse200) GetAdditionalData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateProjectResponse200) GetAdditionalDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *UpdateProjectResponse200) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]interface{} and assigns it to the AdditionalData field.
func (o *UpdateProjectResponse200) SetAdditionalData(v map[string]interface{}) {
	o.AdditionalData = v
}

func (o UpdateProjectResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProjectResponse200) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.AdditionalData != nil {
		toSerialize["additional_data"] = o.AdditionalData
	}
	return toSerialize, nil
}

type NullableUpdateProjectResponse200 struct {
	value *UpdateProjectResponse200
	isSet bool
}

func (v NullableUpdateProjectResponse200) Get() *UpdateProjectResponse200 {
	return v.value
}

func (v *NullableUpdateProjectResponse200) Set(val *UpdateProjectResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProjectResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProjectResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProjectResponse200(val *UpdateProjectResponse200) *NullableUpdateProjectResponse200 {
	return &NullableUpdateProjectResponse200{value: val, isSet: true}
}

func (v NullableUpdateProjectResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProjectResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


