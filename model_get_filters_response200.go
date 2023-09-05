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

// checks if the GetFiltersResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFiltersResponse200{}

// GetFiltersResponse200 struct for GetFiltersResponse200
type GetFiltersResponse200 struct {
	// If the response is successful or not
	Success *bool `json:"success,omitempty"`
	// The array of filters
	Data []GetFiltersResponse200AllOfDataInner `json:"data,omitempty"`
}

// NewGetFiltersResponse200 instantiates a new GetFiltersResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFiltersResponse200() *GetFiltersResponse200 {
	this := GetFiltersResponse200{}
	return &this
}

// NewGetFiltersResponse200WithDefaults instantiates a new GetFiltersResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFiltersResponse200WithDefaults() *GetFiltersResponse200 {
	this := GetFiltersResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *GetFiltersResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiltersResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *GetFiltersResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *GetFiltersResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetFiltersResponse200) GetData() []GetFiltersResponse200AllOfDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetFiltersResponse200AllOfDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFiltersResponse200) GetDataOk() ([]GetFiltersResponse200AllOfDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetFiltersResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetFiltersResponse200AllOfDataInner and assigns it to the Data field.
func (o *GetFiltersResponse200) SetData(v []GetFiltersResponse200AllOfDataInner) {
	o.Data = v
}

func (o GetFiltersResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFiltersResponse200) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetFiltersResponse200 struct {
	value *GetFiltersResponse200
	isSet bool
}

func (v NullableGetFiltersResponse200) Get() *GetFiltersResponse200 {
	return v.value
}

func (v *NullableGetFiltersResponse200) Set(val *GetFiltersResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFiltersResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFiltersResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFiltersResponse200(val *GetFiltersResponse200) *NullableGetFiltersResponse200 {
	return &NullableGetFiltersResponse200{value: val, isSet: true}
}

func (v NullableGetFiltersResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFiltersResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


