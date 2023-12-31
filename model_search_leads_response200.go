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

// checks if the SearchLeadsResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchLeadsResponse200{}

// SearchLeadsResponse200 struct for SearchLeadsResponse200
type SearchLeadsResponse200 struct {
	// If the response is successful or not
	Success *bool `json:"success,omitempty"`
	Data *SearchLeadsResponse200AllOfData `json:"data,omitempty"`
	AdditionalData *GetActivitiesResponse200AdditionalData `json:"additional_data,omitempty"`
}

// NewSearchLeadsResponse200 instantiates a new SearchLeadsResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchLeadsResponse200() *SearchLeadsResponse200 {
	this := SearchLeadsResponse200{}
	return &this
}

// NewSearchLeadsResponse200WithDefaults instantiates a new SearchLeadsResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchLeadsResponse200WithDefaults() *SearchLeadsResponse200 {
	this := SearchLeadsResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SearchLeadsResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchLeadsResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SearchLeadsResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *SearchLeadsResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SearchLeadsResponse200) GetData() SearchLeadsResponse200AllOfData {
	if o == nil || IsNil(o.Data) {
		var ret SearchLeadsResponse200AllOfData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchLeadsResponse200) GetDataOk() (*SearchLeadsResponse200AllOfData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SearchLeadsResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given SearchLeadsResponse200AllOfData and assigns it to the Data field.
func (o *SearchLeadsResponse200) SetData(v SearchLeadsResponse200AllOfData) {
	o.Data = &v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *SearchLeadsResponse200) GetAdditionalData() GetActivitiesResponse200AdditionalData {
	if o == nil || IsNil(o.AdditionalData) {
		var ret GetActivitiesResponse200AdditionalData
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchLeadsResponse200) GetAdditionalDataOk() (*GetActivitiesResponse200AdditionalData, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *SearchLeadsResponse200) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given GetActivitiesResponse200AdditionalData and assigns it to the AdditionalData field.
func (o *SearchLeadsResponse200) SetAdditionalData(v GetActivitiesResponse200AdditionalData) {
	o.AdditionalData = &v
}

func (o SearchLeadsResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchLeadsResponse200) ToMap() (map[string]interface{}, error) {
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

type NullableSearchLeadsResponse200 struct {
	value *SearchLeadsResponse200
	isSet bool
}

func (v NullableSearchLeadsResponse200) Get() *SearchLeadsResponse200 {
	return v.value
}

func (v *NullableSearchLeadsResponse200) Set(val *SearchLeadsResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchLeadsResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchLeadsResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchLeadsResponse200(val *SearchLeadsResponse200) *NullableSearchLeadsResponse200 {
	return &NullableSearchLeadsResponse200{value: val, isSet: true}
}

func (v NullableSearchLeadsResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchLeadsResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


