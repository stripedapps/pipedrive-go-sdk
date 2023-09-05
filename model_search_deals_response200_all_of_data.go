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

// checks if the SearchDealsResponse200AllOfData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchDealsResponse200AllOfData{}

// SearchDealsResponse200AllOfData struct for SearchDealsResponse200AllOfData
type SearchDealsResponse200AllOfData struct {
	// The array of deals
	Items []SearchDealsResponse200AllOfDataItemsInner `json:"items,omitempty"`
}

// NewSearchDealsResponse200AllOfData instantiates a new SearchDealsResponse200AllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchDealsResponse200AllOfData() *SearchDealsResponse200AllOfData {
	this := SearchDealsResponse200AllOfData{}
	return &this
}

// NewSearchDealsResponse200AllOfDataWithDefaults instantiates a new SearchDealsResponse200AllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchDealsResponse200AllOfDataWithDefaults() *SearchDealsResponse200AllOfData {
	this := SearchDealsResponse200AllOfData{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchDealsResponse200AllOfData) GetItems() []SearchDealsResponse200AllOfDataItemsInner {
	if o == nil || IsNil(o.Items) {
		var ret []SearchDealsResponse200AllOfDataItemsInner
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchDealsResponse200AllOfData) GetItemsOk() ([]SearchDealsResponse200AllOfDataItemsInner, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchDealsResponse200AllOfData) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SearchDealsResponse200AllOfDataItemsInner and assigns it to the Items field.
func (o *SearchDealsResponse200AllOfData) SetItems(v []SearchDealsResponse200AllOfDataItemsInner) {
	o.Items = v
}

func (o SearchDealsResponse200AllOfData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchDealsResponse200AllOfData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableSearchDealsResponse200AllOfData struct {
	value *SearchDealsResponse200AllOfData
	isSet bool
}

func (v NullableSearchDealsResponse200AllOfData) Get() *SearchDealsResponse200AllOfData {
	return v.value
}

func (v *NullableSearchDealsResponse200AllOfData) Set(val *SearchDealsResponse200AllOfData) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchDealsResponse200AllOfData) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchDealsResponse200AllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchDealsResponse200AllOfData(val *SearchDealsResponse200AllOfData) *NullableSearchDealsResponse200AllOfData {
	return &NullableSearchDealsResponse200AllOfData{value: val, isSet: true}
}

func (v NullableSearchDealsResponse200AllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchDealsResponse200AllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


