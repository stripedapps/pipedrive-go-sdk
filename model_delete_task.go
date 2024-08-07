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

// checks if the DeleteTask type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteTask{}

// DeleteTask struct for DeleteTask
type DeleteTask struct {
	// If the request was successful or not
	Success *bool `json:"success,omitempty"`
	Data *DeleteTaskData `json:"data,omitempty"`
}

// NewDeleteTask instantiates a new DeleteTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteTask() *DeleteTask {
	this := DeleteTask{}
	return &this
}

// NewDeleteTaskWithDefaults instantiates a new DeleteTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteTaskWithDefaults() *DeleteTask {
	this := DeleteTask{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *DeleteTask) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTask) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *DeleteTask) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *DeleteTask) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DeleteTask) GetData() DeleteTaskData {
	if o == nil || IsNil(o.Data) {
		var ret DeleteTaskData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteTask) GetDataOk() (*DeleteTaskData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DeleteTask) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DeleteTaskData and assigns it to the Data field.
func (o *DeleteTask) SetData(v DeleteTaskData) {
	o.Data = &v
}

func (o DeleteTask) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteTask) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableDeleteTask struct {
	value *DeleteTask
	isSet bool
}

func (v NullableDeleteTask) Get() *DeleteTask {
	return v.value
}

func (v *NullableDeleteTask) Set(val *DeleteTask) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteTask) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteTask(val *DeleteTask) *NullableDeleteTask {
	return &NullableDeleteTask{value: val, isSet: true}
}

func (v NullableDeleteTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


