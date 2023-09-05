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

// checks if the SearchItemByFieldResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchItemByFieldResponse200{}

// SearchItemByFieldResponse200 struct for SearchItemByFieldResponse200
type SearchItemByFieldResponse200 struct {
	// If the response is successful or not
	Success *bool `json:"success,omitempty"`
	// The array of results
	Data []SearchItemByFieldResponse200AllOfDataInner `json:"data,omitempty"`
	AdditionalData *GetActivitiesResponse200AdditionalData `json:"additional_data,omitempty"`
}

// NewSearchItemByFieldResponse200 instantiates a new SearchItemByFieldResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchItemByFieldResponse200() *SearchItemByFieldResponse200 {
	this := SearchItemByFieldResponse200{}
	return &this
}

// NewSearchItemByFieldResponse200WithDefaults instantiates a new SearchItemByFieldResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchItemByFieldResponse200WithDefaults() *SearchItemByFieldResponse200 {
	this := SearchItemByFieldResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *SearchItemByFieldResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchItemByFieldResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *SearchItemByFieldResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *SearchItemByFieldResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SearchItemByFieldResponse200) GetData() []SearchItemByFieldResponse200AllOfDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []SearchItemByFieldResponse200AllOfDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchItemByFieldResponse200) GetDataOk() ([]SearchItemByFieldResponse200AllOfDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SearchItemByFieldResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []SearchItemByFieldResponse200AllOfDataInner and assigns it to the Data field.
func (o *SearchItemByFieldResponse200) SetData(v []SearchItemByFieldResponse200AllOfDataInner) {
	o.Data = v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *SearchItemByFieldResponse200) GetAdditionalData() GetActivitiesResponse200AdditionalData {
	if o == nil || IsNil(o.AdditionalData) {
		var ret GetActivitiesResponse200AdditionalData
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchItemByFieldResponse200) GetAdditionalDataOk() (*GetActivitiesResponse200AdditionalData, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *SearchItemByFieldResponse200) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given GetActivitiesResponse200AdditionalData and assigns it to the AdditionalData field.
func (o *SearchItemByFieldResponse200) SetAdditionalData(v GetActivitiesResponse200AdditionalData) {
	o.AdditionalData = &v
}

func (o SearchItemByFieldResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchItemByFieldResponse200) ToMap() (map[string]interface{}, error) {
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

type NullableSearchItemByFieldResponse200 struct {
	value *SearchItemByFieldResponse200
	isSet bool
}

func (v NullableSearchItemByFieldResponse200) Get() *SearchItemByFieldResponse200 {
	return v.value
}

func (v *NullableSearchItemByFieldResponse200) Set(val *SearchItemByFieldResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchItemByFieldResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchItemByFieldResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchItemByFieldResponse200(val *SearchItemByFieldResponse200) *NullableSearchItemByFieldResponse200 {
	return &NullableSearchItemByFieldResponse200{value: val, isSet: true}
}

func (v NullableSearchItemByFieldResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchItemByFieldResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


