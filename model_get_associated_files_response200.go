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

// checks if the GetAssociatedFilesResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAssociatedFilesResponse200{}

// GetAssociatedFilesResponse200 struct for GetAssociatedFilesResponse200
type GetAssociatedFilesResponse200 struct {
	// If the response is successful or not
	Success *bool `json:"success,omitempty"`
	// The array of files
	Data []GetAssociatedFilesResponse200AllOfDataInner `json:"data,omitempty"`
	AdditionalData *FieldsResponse200AllOfAdditionalData `json:"additional_data,omitempty"`
}

// NewGetAssociatedFilesResponse200 instantiates a new GetAssociatedFilesResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAssociatedFilesResponse200() *GetAssociatedFilesResponse200 {
	this := GetAssociatedFilesResponse200{}
	return &this
}

// NewGetAssociatedFilesResponse200WithDefaults instantiates a new GetAssociatedFilesResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAssociatedFilesResponse200WithDefaults() *GetAssociatedFilesResponse200 {
	this := GetAssociatedFilesResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *GetAssociatedFilesResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssociatedFilesResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *GetAssociatedFilesResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *GetAssociatedFilesResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetAssociatedFilesResponse200) GetData() []GetAssociatedFilesResponse200AllOfDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetAssociatedFilesResponse200AllOfDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssociatedFilesResponse200) GetDataOk() ([]GetAssociatedFilesResponse200AllOfDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetAssociatedFilesResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetAssociatedFilesResponse200AllOfDataInner and assigns it to the Data field.
func (o *GetAssociatedFilesResponse200) SetData(v []GetAssociatedFilesResponse200AllOfDataInner) {
	o.Data = v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *GetAssociatedFilesResponse200) GetAdditionalData() FieldsResponse200AllOfAdditionalData {
	if o == nil || IsNil(o.AdditionalData) {
		var ret FieldsResponse200AllOfAdditionalData
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssociatedFilesResponse200) GetAdditionalDataOk() (*FieldsResponse200AllOfAdditionalData, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *GetAssociatedFilesResponse200) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given FieldsResponse200AllOfAdditionalData and assigns it to the AdditionalData field.
func (o *GetAssociatedFilesResponse200) SetAdditionalData(v FieldsResponse200AllOfAdditionalData) {
	o.AdditionalData = &v
}

func (o GetAssociatedFilesResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAssociatedFilesResponse200) ToMap() (map[string]interface{}, error) {
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

type NullableGetAssociatedFilesResponse200 struct {
	value *GetAssociatedFilesResponse200
	isSet bool
}

func (v NullableGetAssociatedFilesResponse200) Get() *GetAssociatedFilesResponse200 {
	return v.value
}

func (v *NullableGetAssociatedFilesResponse200) Set(val *GetAssociatedFilesResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAssociatedFilesResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAssociatedFilesResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAssociatedFilesResponse200(val *GetAssociatedFilesResponse200) *NullableGetAssociatedFilesResponse200 {
	return &NullableGetAssociatedFilesResponse200{value: val, isSet: true}
}

func (v NullableGetAssociatedFilesResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAssociatedFilesResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


