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

// checks if the GetAddProductAttachementResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAddProductAttachementResponse200{}

// GetAddProductAttachementResponse200 struct for GetAddProductAttachementResponse200
type GetAddProductAttachementResponse200 struct {
	// If the response is successful or not
	Success *bool `json:"success,omitempty"`
	Data *GetAddProductAttachementResponse200Data `json:"data,omitempty"`
}

// NewGetAddProductAttachementResponse200 instantiates a new GetAddProductAttachementResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAddProductAttachementResponse200() *GetAddProductAttachementResponse200 {
	this := GetAddProductAttachementResponse200{}
	return &this
}

// NewGetAddProductAttachementResponse200WithDefaults instantiates a new GetAddProductAttachementResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAddProductAttachementResponse200WithDefaults() *GetAddProductAttachementResponse200 {
	this := GetAddProductAttachementResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *GetAddProductAttachementResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAddProductAttachementResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *GetAddProductAttachementResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *GetAddProductAttachementResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetAddProductAttachementResponse200) GetData() GetAddProductAttachementResponse200Data {
	if o == nil || IsNil(o.Data) {
		var ret GetAddProductAttachementResponse200Data
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAddProductAttachementResponse200) GetDataOk() (*GetAddProductAttachementResponse200Data, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetAddProductAttachementResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetAddProductAttachementResponse200Data and assigns it to the Data field.
func (o *GetAddProductAttachementResponse200) SetData(v GetAddProductAttachementResponse200Data) {
	o.Data = &v
}

func (o GetAddProductAttachementResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAddProductAttachementResponse200) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetAddProductAttachementResponse200 struct {
	value *GetAddProductAttachementResponse200
	isSet bool
}

func (v NullableGetAddProductAttachementResponse200) Get() *GetAddProductAttachementResponse200 {
	return v.value
}

func (v *NullableGetAddProductAttachementResponse200) Set(val *GetAddProductAttachementResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAddProductAttachementResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAddProductAttachementResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAddProductAttachementResponse200(val *GetAddProductAttachementResponse200) *NullableGetAddProductAttachementResponse200 {
	return &NullableGetAddProductAttachementResponse200{value: val, isSet: true}
}

func (v NullableGetAddProductAttachementResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAddProductAttachementResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


