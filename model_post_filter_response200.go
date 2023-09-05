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

// checks if the PostFilterResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostFilterResponse200{}

// PostFilterResponse200 struct for PostFilterResponse200
type PostFilterResponse200 struct {
	// If the response is successful or not
	Success *bool `json:"success,omitempty"`
	Data *PostFilterResponse200AllOfData `json:"data,omitempty"`
}

// NewPostFilterResponse200 instantiates a new PostFilterResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostFilterResponse200() *PostFilterResponse200 {
	this := PostFilterResponse200{}
	return &this
}

// NewPostFilterResponse200WithDefaults instantiates a new PostFilterResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostFilterResponse200WithDefaults() *PostFilterResponse200 {
	this := PostFilterResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *PostFilterResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostFilterResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *PostFilterResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *PostFilterResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PostFilterResponse200) GetData() PostFilterResponse200AllOfData {
	if o == nil || IsNil(o.Data) {
		var ret PostFilterResponse200AllOfData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostFilterResponse200) GetDataOk() (*PostFilterResponse200AllOfData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PostFilterResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PostFilterResponse200AllOfData and assigns it to the Data field.
func (o *PostFilterResponse200) SetData(v PostFilterResponse200AllOfData) {
	o.Data = &v
}

func (o PostFilterResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostFilterResponse200) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePostFilterResponse200 struct {
	value *PostFilterResponse200
	isSet bool
}

func (v NullablePostFilterResponse200) Get() *PostFilterResponse200 {
	return v.value
}

func (v *NullablePostFilterResponse200) Set(val *PostFilterResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullablePostFilterResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullablePostFilterResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostFilterResponse200(val *PostFilterResponse200) *NullablePostFilterResponse200 {
	return &NullablePostFilterResponse200{value: val, isSet: true}
}

func (v NullablePostFilterResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostFilterResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


