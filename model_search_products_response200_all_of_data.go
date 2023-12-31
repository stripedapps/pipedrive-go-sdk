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

// checks if the SearchProductsResponse200AllOfData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchProductsResponse200AllOfData{}

// SearchProductsResponse200AllOfData struct for SearchProductsResponse200AllOfData
type SearchProductsResponse200AllOfData struct {
	// The array of found items
	Items []SearchProductsResponse200AllOfDataItemsInner `json:"items,omitempty"`
}

// NewSearchProductsResponse200AllOfData instantiates a new SearchProductsResponse200AllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchProductsResponse200AllOfData() *SearchProductsResponse200AllOfData {
	this := SearchProductsResponse200AllOfData{}
	return &this
}

// NewSearchProductsResponse200AllOfDataWithDefaults instantiates a new SearchProductsResponse200AllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchProductsResponse200AllOfDataWithDefaults() *SearchProductsResponse200AllOfData {
	this := SearchProductsResponse200AllOfData{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchProductsResponse200AllOfData) GetItems() []SearchProductsResponse200AllOfDataItemsInner {
	if o == nil || IsNil(o.Items) {
		var ret []SearchProductsResponse200AllOfDataItemsInner
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchProductsResponse200AllOfData) GetItemsOk() ([]SearchProductsResponse200AllOfDataItemsInner, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchProductsResponse200AllOfData) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SearchProductsResponse200AllOfDataItemsInner and assigns it to the Items field.
func (o *SearchProductsResponse200AllOfData) SetItems(v []SearchProductsResponse200AllOfDataItemsInner) {
	o.Items = v
}

func (o SearchProductsResponse200AllOfData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchProductsResponse200AllOfData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableSearchProductsResponse200AllOfData struct {
	value *SearchProductsResponse200AllOfData
	isSet bool
}

func (v NullableSearchProductsResponse200AllOfData) Get() *SearchProductsResponse200AllOfData {
	return v.value
}

func (v *NullableSearchProductsResponse200AllOfData) Set(val *SearchProductsResponse200AllOfData) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchProductsResponse200AllOfData) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchProductsResponse200AllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchProductsResponse200AllOfData(val *SearchProductsResponse200AllOfData) *NullableSearchProductsResponse200AllOfData {
	return &NullableSearchProductsResponse200AllOfData{value: val, isSet: true}
}

func (v NullableSearchProductsResponse200AllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchProductsResponse200AllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


