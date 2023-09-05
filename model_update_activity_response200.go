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

// checks if the UpdateActivityResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateActivityResponse200{}

// UpdateActivityResponse200 struct for UpdateActivityResponse200
type UpdateActivityResponse200 struct {
	Success *bool `json:"success,omitempty"`
	Data *ActivityResponseObject `json:"data,omitempty"`
	RelatedObjects *AddActivityResponse200RelatedObjects `json:"related_objects,omitempty"`
}

// NewUpdateActivityResponse200 instantiates a new UpdateActivityResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateActivityResponse200() *UpdateActivityResponse200 {
	this := UpdateActivityResponse200{}
	return &this
}

// NewUpdateActivityResponse200WithDefaults instantiates a new UpdateActivityResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateActivityResponse200WithDefaults() *UpdateActivityResponse200 {
	this := UpdateActivityResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *UpdateActivityResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActivityResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *UpdateActivityResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *UpdateActivityResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UpdateActivityResponse200) GetData() ActivityResponseObject {
	if o == nil || IsNil(o.Data) {
		var ret ActivityResponseObject
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActivityResponse200) GetDataOk() (*ActivityResponseObject, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UpdateActivityResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ActivityResponseObject and assigns it to the Data field.
func (o *UpdateActivityResponse200) SetData(v ActivityResponseObject) {
	o.Data = &v
}

// GetRelatedObjects returns the RelatedObjects field value if set, zero value otherwise.
func (o *UpdateActivityResponse200) GetRelatedObjects() AddActivityResponse200RelatedObjects {
	if o == nil || IsNil(o.RelatedObjects) {
		var ret AddActivityResponse200RelatedObjects
		return ret
	}
	return *o.RelatedObjects
}

// GetRelatedObjectsOk returns a tuple with the RelatedObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActivityResponse200) GetRelatedObjectsOk() (*AddActivityResponse200RelatedObjects, bool) {
	if o == nil || IsNil(o.RelatedObjects) {
		return nil, false
	}
	return o.RelatedObjects, true
}

// HasRelatedObjects returns a boolean if a field has been set.
func (o *UpdateActivityResponse200) HasRelatedObjects() bool {
	if o != nil && !IsNil(o.RelatedObjects) {
		return true
	}

	return false
}

// SetRelatedObjects gets a reference to the given AddActivityResponse200RelatedObjects and assigns it to the RelatedObjects field.
func (o *UpdateActivityResponse200) SetRelatedObjects(v AddActivityResponse200RelatedObjects) {
	o.RelatedObjects = &v
}

func (o UpdateActivityResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateActivityResponse200) ToMap() (map[string]interface{}, error) {
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

type NullableUpdateActivityResponse200 struct {
	value *UpdateActivityResponse200
	isSet bool
}

func (v NullableUpdateActivityResponse200) Get() *UpdateActivityResponse200 {
	return v.value
}

func (v *NullableUpdateActivityResponse200) Set(val *UpdateActivityResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateActivityResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateActivityResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateActivityResponse200(val *UpdateActivityResponse200) *NullableUpdateActivityResponse200 {
	return &NullableUpdateActivityResponse200{value: val, isSet: true}
}

func (v NullableUpdateActivityResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateActivityResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


