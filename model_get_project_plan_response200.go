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

// checks if the GetProjectPlanResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProjectPlanResponse200{}

// GetProjectPlanResponse200 struct for GetProjectPlanResponse200
type GetProjectPlanResponse200 struct {
	Success *bool `json:"success,omitempty"`
	Data []GetProjectPlanResponse200DataInner `json:"data,omitempty"`
	AdditionalData map[string]interface{} `json:"additional_data,omitempty"`
}

// NewGetProjectPlanResponse200 instantiates a new GetProjectPlanResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProjectPlanResponse200() *GetProjectPlanResponse200 {
	this := GetProjectPlanResponse200{}
	return &this
}

// NewGetProjectPlanResponse200WithDefaults instantiates a new GetProjectPlanResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProjectPlanResponse200WithDefaults() *GetProjectPlanResponse200 {
	this := GetProjectPlanResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *GetProjectPlanResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectPlanResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *GetProjectPlanResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *GetProjectPlanResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetProjectPlanResponse200) GetData() []GetProjectPlanResponse200DataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetProjectPlanResponse200DataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectPlanResponse200) GetDataOk() ([]GetProjectPlanResponse200DataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetProjectPlanResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetProjectPlanResponse200DataInner and assigns it to the Data field.
func (o *GetProjectPlanResponse200) SetData(v []GetProjectPlanResponse200DataInner) {
	o.Data = v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetProjectPlanResponse200) GetAdditionalData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetProjectPlanResponse200) GetAdditionalDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *GetProjectPlanResponse200) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]interface{} and assigns it to the AdditionalData field.
func (o *GetProjectPlanResponse200) SetAdditionalData(v map[string]interface{}) {
	o.AdditionalData = v
}

func (o GetProjectPlanResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProjectPlanResponse200) ToMap() (map[string]interface{}, error) {
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

type NullableGetProjectPlanResponse200 struct {
	value *GetProjectPlanResponse200
	isSet bool
}

func (v NullableGetProjectPlanResponse200) Get() *GetProjectPlanResponse200 {
	return v.value
}

func (v *NullableGetProjectPlanResponse200) Set(val *GetProjectPlanResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProjectPlanResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProjectPlanResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProjectPlanResponse200(val *GetProjectPlanResponse200) *NullableGetProjectPlanResponse200 {
	return &NullableGetProjectPlanResponse200{value: val, isSet: true}
}

func (v NullableGetProjectPlanResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProjectPlanResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


