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

// checks if the SearchProductsResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchProductsResponse200{}

// SearchProductsResponse200 struct for SearchProductsResponse200
type SearchProductsResponse200 struct {
	// If the response is successful or not
	Success *bool `json:"success,omitempty"`
	Data *SearchProductsResponse200AllOfData `json:"data,omitempty"`
	AdditionalData *GetActivitiesResponse200AdditionalData `json:"additional_data,omitempty"`
}

// NewSearchProductsResponse200 instantiates a new SearchProductsResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchProductsResponse200() *SearchProductsResponse200 {
	this := SearchProductsResponse200{}
	return &this
}

// NewSearchProductsResponse200WithDefaults instantiates a new SearchProductsResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchProductsResponse200WithDefaults() *SearchProductsResponse200 {
	this := SearchProductsResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SearchProductsResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchProductsResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SearchProductsResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *SearchProductsResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SearchProductsResponse200) GetData() SearchProductsResponse200AllOfData {
	if o == nil || IsNil(o.Data) {
		var ret SearchProductsResponse200AllOfData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchProductsResponse200) GetDataOk() (*SearchProductsResponse200AllOfData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SearchProductsResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given SearchProductsResponse200AllOfData and assigns it to the Data field.
func (o *SearchProductsResponse200) SetData(v SearchProductsResponse200AllOfData) {
	o.Data = &v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *SearchProductsResponse200) GetAdditionalData() GetActivitiesResponse200AdditionalData {
	if o == nil || IsNil(o.AdditionalData) {
		var ret GetActivitiesResponse200AdditionalData
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchProductsResponse200) GetAdditionalDataOk() (*GetActivitiesResponse200AdditionalData, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *SearchProductsResponse200) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given GetActivitiesResponse200AdditionalData and assigns it to the AdditionalData field.
func (o *SearchProductsResponse200) SetAdditionalData(v GetActivitiesResponse200AdditionalData) {
	o.AdditionalData = &v
}

func (o SearchProductsResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchProductsResponse200) ToMap() (map[string]interface{}, error) {
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

type NullableSearchProductsResponse200 struct {
	value *SearchProductsResponse200
	isSet bool
}

func (v NullableSearchProductsResponse200) Get() *SearchProductsResponse200 {
	return v.value
}

func (v *NullableSearchProductsResponse200) Set(val *SearchProductsResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchProductsResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchProductsResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchProductsResponse200(val *SearchProductsResponse200) *NullableSearchProductsResponse200 {
	return &NullableSearchProductsResponse200{value: val, isSet: true}
}

func (v NullableSearchProductsResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchProductsResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


