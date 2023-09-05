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

// checks if the GetLeadSourcesResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLeadSourcesResponse200{}

// GetLeadSourcesResponse200 struct for GetLeadSourcesResponse200
type GetLeadSourcesResponse200 struct {
	Success *bool `json:"success,omitempty"`
	Data []GetLeadSourcesResponse200DataInner `json:"data,omitempty"`
}

// NewGetLeadSourcesResponse200 instantiates a new GetLeadSourcesResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLeadSourcesResponse200() *GetLeadSourcesResponse200 {
	this := GetLeadSourcesResponse200{}
	return &this
}

// NewGetLeadSourcesResponse200WithDefaults instantiates a new GetLeadSourcesResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLeadSourcesResponse200WithDefaults() *GetLeadSourcesResponse200 {
	this := GetLeadSourcesResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *GetLeadSourcesResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLeadSourcesResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *GetLeadSourcesResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *GetLeadSourcesResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetLeadSourcesResponse200) GetData() []GetLeadSourcesResponse200DataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetLeadSourcesResponse200DataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLeadSourcesResponse200) GetDataOk() ([]GetLeadSourcesResponse200DataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetLeadSourcesResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetLeadSourcesResponse200DataInner and assigns it to the Data field.
func (o *GetLeadSourcesResponse200) SetData(v []GetLeadSourcesResponse200DataInner) {
	o.Data = v
}

func (o GetLeadSourcesResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLeadSourcesResponse200) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetLeadSourcesResponse200 struct {
	value *GetLeadSourcesResponse200
	isSet bool
}

func (v NullableGetLeadSourcesResponse200) Get() *GetLeadSourcesResponse200 {
	return v.value
}

func (v *NullableGetLeadSourcesResponse200) Set(val *GetLeadSourcesResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLeadSourcesResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLeadSourcesResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLeadSourcesResponse200(val *GetLeadSourcesResponse200) *NullableGetLeadSourcesResponse200 {
	return &NullableGetLeadSourcesResponse200{value: val, isSet: true}
}

func (v NullableGetLeadSourcesResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLeadSourcesResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


