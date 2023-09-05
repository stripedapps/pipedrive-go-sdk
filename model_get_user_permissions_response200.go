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

// checks if the GetUserPermissionsResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserPermissionsResponse200{}

// GetUserPermissionsResponse200 struct for GetUserPermissionsResponse200
type GetUserPermissionsResponse200 struct {
	// If the response is successful or not
	Success *bool `json:"success,omitempty"`
	Data *GetUserPermissionsResponse200AllOfData `json:"data,omitempty"`
}

// NewGetUserPermissionsResponse200 instantiates a new GetUserPermissionsResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserPermissionsResponse200() *GetUserPermissionsResponse200 {
	this := GetUserPermissionsResponse200{}
	return &this
}

// NewGetUserPermissionsResponse200WithDefaults instantiates a new GetUserPermissionsResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserPermissionsResponse200WithDefaults() *GetUserPermissionsResponse200 {
	this := GetUserPermissionsResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *GetUserPermissionsResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserPermissionsResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *GetUserPermissionsResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *GetUserPermissionsResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetUserPermissionsResponse200) GetData() GetUserPermissionsResponse200AllOfData {
	if o == nil || IsNil(o.Data) {
		var ret GetUserPermissionsResponse200AllOfData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserPermissionsResponse200) GetDataOk() (*GetUserPermissionsResponse200AllOfData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetUserPermissionsResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetUserPermissionsResponse200AllOfData and assigns it to the Data field.
func (o *GetUserPermissionsResponse200) SetData(v GetUserPermissionsResponse200AllOfData) {
	o.Data = &v
}

func (o GetUserPermissionsResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserPermissionsResponse200) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetUserPermissionsResponse200 struct {
	value *GetUserPermissionsResponse200
	isSet bool
}

func (v NullableGetUserPermissionsResponse200) Get() *GetUserPermissionsResponse200 {
	return v.value
}

func (v *NullableGetUserPermissionsResponse200) Set(val *GetUserPermissionsResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserPermissionsResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserPermissionsResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserPermissionsResponse200(val *GetUserPermissionsResponse200) *NullableGetUserPermissionsResponse200 {
	return &NullableGetUserPermissionsResponse200{value: val, isSet: true}
}

func (v NullableGetUserPermissionsResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserPermissionsResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


