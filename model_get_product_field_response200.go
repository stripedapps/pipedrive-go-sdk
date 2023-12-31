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

// checks if the GetProductFieldResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProductFieldResponse200{}

// GetProductFieldResponse200 struct for GetProductFieldResponse200
type GetProductFieldResponse200 struct {
	// If the response is successful or not
	Success *bool `json:"success,omitempty"`
	Data *GetProductFieldResponse200Data `json:"data,omitempty"`
}

// NewGetProductFieldResponse200 instantiates a new GetProductFieldResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductFieldResponse200() *GetProductFieldResponse200 {
	this := GetProductFieldResponse200{}
	return &this
}

// NewGetProductFieldResponse200WithDefaults instantiates a new GetProductFieldResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductFieldResponse200WithDefaults() *GetProductFieldResponse200 {
	this := GetProductFieldResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *GetProductFieldResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductFieldResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *GetProductFieldResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *GetProductFieldResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetProductFieldResponse200) GetData() GetProductFieldResponse200Data {
	if o == nil || IsNil(o.Data) {
		var ret GetProductFieldResponse200Data
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductFieldResponse200) GetDataOk() (*GetProductFieldResponse200Data, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetProductFieldResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetProductFieldResponse200Data and assigns it to the Data field.
func (o *GetProductFieldResponse200) SetData(v GetProductFieldResponse200Data) {
	o.Data = &v
}

func (o GetProductFieldResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProductFieldResponse200) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetProductFieldResponse200 struct {
	value *GetProductFieldResponse200
	isSet bool
}

func (v NullableGetProductFieldResponse200) Get() *GetProductFieldResponse200 {
	return v.value
}

func (v *NullableGetProductFieldResponse200) Set(val *GetProductFieldResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductFieldResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductFieldResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductFieldResponse200(val *GetProductFieldResponse200) *NullableGetProductFieldResponse200 {
	return &NullableGetProductFieldResponse200{value: val, isSet: true}
}

func (v NullableGetProductFieldResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProductFieldResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


