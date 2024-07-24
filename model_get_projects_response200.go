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

// checks if the GetProjectsResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProjectsResponse200{}

// GetProjectsResponse200 struct for GetProjectsResponse200
type GetProjectsResponse200 struct {
	Success *bool `json:"success,omitempty"`
	Data []ProjectResponseObject `json:"data,omitempty"`
	AdditionalData *GetActivitiesCollectionResponse200AdditionalData `json:"additional_data,omitempty"`
}

// NewGetProjectsResponse200 instantiates a new GetProjectsResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProjectsResponse200() *GetProjectsResponse200 {
	this := GetProjectsResponse200{}
	return &this
}

// NewGetProjectsResponse200WithDefaults instantiates a new GetProjectsResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProjectsResponse200WithDefaults() *GetProjectsResponse200 {
	this := GetProjectsResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *GetProjectsResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectsResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *GetProjectsResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *GetProjectsResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetProjectsResponse200) GetData() []ProjectResponseObject {
	if o == nil || IsNil(o.Data) {
		var ret []ProjectResponseObject
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectsResponse200) GetDataOk() ([]ProjectResponseObject, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetProjectsResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ProjectResponseObject and assigns it to the Data field.
func (o *GetProjectsResponse200) SetData(v []ProjectResponseObject) {
	o.Data = v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *GetProjectsResponse200) GetAdditionalData() GetActivitiesCollectionResponse200AdditionalData {
	if o == nil || IsNil(o.AdditionalData) {
		var ret GetActivitiesCollectionResponse200AdditionalData
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectsResponse200) GetAdditionalDataOk() (*GetActivitiesCollectionResponse200AdditionalData, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *GetProjectsResponse200) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given GetActivitiesCollectionResponse200AdditionalData and assigns it to the AdditionalData field.
func (o *GetProjectsResponse200) SetAdditionalData(v GetActivitiesCollectionResponse200AdditionalData) {
	o.AdditionalData = &v
}

func (o GetProjectsResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProjectsResponse200) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.AdditionalData) {
		toSerialize["additional_data"] = o.AdditionalData
	}
	return toSerialize, nil
}

type NullableGetProjectsResponse200 struct {
	value *GetProjectsResponse200
	isSet bool
}

func (v NullableGetProjectsResponse200) Get() *GetProjectsResponse200 {
	return v.value
}

func (v *NullableGetProjectsResponse200) Set(val *GetProjectsResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProjectsResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProjectsResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProjectsResponse200(val *GetProjectsResponse200) *NullableGetProjectsResponse200 {
	return &NullableGetProjectsResponse200{value: val, isSet: true}
}

func (v NullableGetProjectsResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProjectsResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

