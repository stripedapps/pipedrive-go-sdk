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

// checks if the GetNotesResponse200DataInnerDeal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNotesResponse200DataInnerDeal{}

// GetNotesResponse200DataInnerDeal The deal this note is attached to
type GetNotesResponse200DataInnerDeal struct {
	// The title of the deal this note is attached to
	Title *string `json:"title,omitempty"`
}

// NewGetNotesResponse200DataInnerDeal instantiates a new GetNotesResponse200DataInnerDeal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNotesResponse200DataInnerDeal() *GetNotesResponse200DataInnerDeal {
	this := GetNotesResponse200DataInnerDeal{}
	return &this
}

// NewGetNotesResponse200DataInnerDealWithDefaults instantiates a new GetNotesResponse200DataInnerDeal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNotesResponse200DataInnerDealWithDefaults() *GetNotesResponse200DataInnerDeal {
	this := GetNotesResponse200DataInnerDeal{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *GetNotesResponse200DataInnerDeal) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNotesResponse200DataInnerDeal) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *GetNotesResponse200DataInnerDeal) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *GetNotesResponse200DataInnerDeal) SetTitle(v string) {
	o.Title = &v
}

func (o GetNotesResponse200DataInnerDeal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNotesResponse200DataInnerDeal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableGetNotesResponse200DataInnerDeal struct {
	value *GetNotesResponse200DataInnerDeal
	isSet bool
}

func (v NullableGetNotesResponse200DataInnerDeal) Get() *GetNotesResponse200DataInnerDeal {
	return v.value
}

func (v *NullableGetNotesResponse200DataInnerDeal) Set(val *GetNotesResponse200DataInnerDeal) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNotesResponse200DataInnerDeal) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNotesResponse200DataInnerDeal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNotesResponse200DataInnerDeal(val *GetNotesResponse200DataInnerDeal) *NullableGetNotesResponse200DataInnerDeal {
	return &NullableGetNotesResponse200DataInnerDeal{value: val, isSet: true}
}

func (v NullableGetNotesResponse200DataInnerDeal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNotesResponse200DataInnerDeal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


