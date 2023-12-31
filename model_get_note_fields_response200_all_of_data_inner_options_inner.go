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

// checks if the GetNoteFieldsResponse200AllOfDataInnerOptionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNoteFieldsResponse200AllOfDataInnerOptionsInner{}

// GetNoteFieldsResponse200AllOfDataInnerOptionsInner struct for GetNoteFieldsResponse200AllOfDataInnerOptionsInner
type GetNoteFieldsResponse200AllOfDataInnerOptionsInner struct {
	Id *int32 `json:"id,omitempty"`
	Label *string `json:"label,omitempty"`
}

// NewGetNoteFieldsResponse200AllOfDataInnerOptionsInner instantiates a new GetNoteFieldsResponse200AllOfDataInnerOptionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNoteFieldsResponse200AllOfDataInnerOptionsInner() *GetNoteFieldsResponse200AllOfDataInnerOptionsInner {
	this := GetNoteFieldsResponse200AllOfDataInnerOptionsInner{}
	return &this
}

// NewGetNoteFieldsResponse200AllOfDataInnerOptionsInnerWithDefaults instantiates a new GetNoteFieldsResponse200AllOfDataInnerOptionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNoteFieldsResponse200AllOfDataInnerOptionsInnerWithDefaults() *GetNoteFieldsResponse200AllOfDataInnerOptionsInner {
	this := GetNoteFieldsResponse200AllOfDataInnerOptionsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNoteFieldsResponse200AllOfDataInnerOptionsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNoteFieldsResponse200AllOfDataInnerOptionsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNoteFieldsResponse200AllOfDataInnerOptionsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetNoteFieldsResponse200AllOfDataInnerOptionsInner) SetId(v int32) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *GetNoteFieldsResponse200AllOfDataInnerOptionsInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNoteFieldsResponse200AllOfDataInnerOptionsInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *GetNoteFieldsResponse200AllOfDataInnerOptionsInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *GetNoteFieldsResponse200AllOfDataInnerOptionsInner) SetLabel(v string) {
	o.Label = &v
}

func (o GetNoteFieldsResponse200AllOfDataInnerOptionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNoteFieldsResponse200AllOfDataInnerOptionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullableGetNoteFieldsResponse200AllOfDataInnerOptionsInner struct {
	value *GetNoteFieldsResponse200AllOfDataInnerOptionsInner
	isSet bool
}

func (v NullableGetNoteFieldsResponse200AllOfDataInnerOptionsInner) Get() *GetNoteFieldsResponse200AllOfDataInnerOptionsInner {
	return v.value
}

func (v *NullableGetNoteFieldsResponse200AllOfDataInnerOptionsInner) Set(val *GetNoteFieldsResponse200AllOfDataInnerOptionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNoteFieldsResponse200AllOfDataInnerOptionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNoteFieldsResponse200AllOfDataInnerOptionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNoteFieldsResponse200AllOfDataInnerOptionsInner(val *GetNoteFieldsResponse200AllOfDataInnerOptionsInner) *NullableGetNoteFieldsResponse200AllOfDataInnerOptionsInner {
	return &NullableGetNoteFieldsResponse200AllOfDataInnerOptionsInner{value: val, isSet: true}
}

func (v NullableGetNoteFieldsResponse200AllOfDataInnerOptionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNoteFieldsResponse200AllOfDataInnerOptionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


