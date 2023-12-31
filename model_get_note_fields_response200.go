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

// checks if the GetNoteFieldsResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNoteFieldsResponse200{}

// GetNoteFieldsResponse200 struct for GetNoteFieldsResponse200
type GetNoteFieldsResponse200 struct {
	// If the response is successful or not
	Success *bool `json:"success,omitempty"`
	Data []GetNoteFieldsResponse200AllOfDataInner `json:"data,omitempty"`
	AdditionalData *FieldsResponse200AllOfAdditionalData `json:"additional_data,omitempty"`
}

// NewGetNoteFieldsResponse200 instantiates a new GetNoteFieldsResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNoteFieldsResponse200() *GetNoteFieldsResponse200 {
	this := GetNoteFieldsResponse200{}
	return &this
}

// NewGetNoteFieldsResponse200WithDefaults instantiates a new GetNoteFieldsResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNoteFieldsResponse200WithDefaults() *GetNoteFieldsResponse200 {
	this := GetNoteFieldsResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *GetNoteFieldsResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNoteFieldsResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *GetNoteFieldsResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *GetNoteFieldsResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetNoteFieldsResponse200) GetData() []GetNoteFieldsResponse200AllOfDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetNoteFieldsResponse200AllOfDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNoteFieldsResponse200) GetDataOk() ([]GetNoteFieldsResponse200AllOfDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetNoteFieldsResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetNoteFieldsResponse200AllOfDataInner and assigns it to the Data field.
func (o *GetNoteFieldsResponse200) SetData(v []GetNoteFieldsResponse200AllOfDataInner) {
	o.Data = v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *GetNoteFieldsResponse200) GetAdditionalData() FieldsResponse200AllOfAdditionalData {
	if o == nil || IsNil(o.AdditionalData) {
		var ret FieldsResponse200AllOfAdditionalData
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNoteFieldsResponse200) GetAdditionalDataOk() (*FieldsResponse200AllOfAdditionalData, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *GetNoteFieldsResponse200) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given FieldsResponse200AllOfAdditionalData and assigns it to the AdditionalData field.
func (o *GetNoteFieldsResponse200) SetAdditionalData(v FieldsResponse200AllOfAdditionalData) {
	o.AdditionalData = &v
}

func (o GetNoteFieldsResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNoteFieldsResponse200) ToMap() (map[string]interface{}, error) {
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

type NullableGetNoteFieldsResponse200 struct {
	value *GetNoteFieldsResponse200
	isSet bool
}

func (v NullableGetNoteFieldsResponse200) Get() *GetNoteFieldsResponse200 {
	return v.value
}

func (v *NullableGetNoteFieldsResponse200) Set(val *GetNoteFieldsResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNoteFieldsResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNoteFieldsResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNoteFieldsResponse200(val *GetNoteFieldsResponse200) *NullableGetNoteFieldsResponse200 {
	return &NullableGetNoteFieldsResponse200{value: val, isSet: true}
}

func (v NullableGetNoteFieldsResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNoteFieldsResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


