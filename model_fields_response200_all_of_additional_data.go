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

// checks if the FieldsResponse200AllOfAdditionalData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FieldsResponse200AllOfAdditionalData{}

// FieldsResponse200AllOfAdditionalData The additional data of the list
type FieldsResponse200AllOfAdditionalData struct {
	// Pagination start
	Start *int32 `json:"start,omitempty"`
	// Items shown per page
	Limit *int32 `json:"limit,omitempty"`
	// If there are more list items in the collection than displayed or not
	MoreItemsInCollection *bool `json:"more_items_in_collection,omitempty"`
}

// NewFieldsResponse200AllOfAdditionalData instantiates a new FieldsResponse200AllOfAdditionalData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldsResponse200AllOfAdditionalData() *FieldsResponse200AllOfAdditionalData {
	this := FieldsResponse200AllOfAdditionalData{}
	return &this
}

// NewFieldsResponse200AllOfAdditionalDataWithDefaults instantiates a new FieldsResponse200AllOfAdditionalData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldsResponse200AllOfAdditionalDataWithDefaults() *FieldsResponse200AllOfAdditionalData {
	this := FieldsResponse200AllOfAdditionalData{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfAdditionalData) GetStart() int32 {
	if o == nil || IsNil(o.Start) {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfAdditionalData) GetStartOk() (*int32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfAdditionalData) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *FieldsResponse200AllOfAdditionalData) SetStart(v int32) {
	o.Start = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfAdditionalData) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfAdditionalData) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfAdditionalData) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *FieldsResponse200AllOfAdditionalData) SetLimit(v int32) {
	o.Limit = &v
}

// GetMoreItemsInCollection returns the MoreItemsInCollection field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfAdditionalData) GetMoreItemsInCollection() bool {
	if o == nil || IsNil(o.MoreItemsInCollection) {
		var ret bool
		return ret
	}
	return *o.MoreItemsInCollection
}

// GetMoreItemsInCollectionOk returns a tuple with the MoreItemsInCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfAdditionalData) GetMoreItemsInCollectionOk() (*bool, bool) {
	if o == nil || IsNil(o.MoreItemsInCollection) {
		return nil, false
	}
	return o.MoreItemsInCollection, true
}

// HasMoreItemsInCollection returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfAdditionalData) HasMoreItemsInCollection() bool {
	if o != nil && !IsNil(o.MoreItemsInCollection) {
		return true
	}

	return false
}

// SetMoreItemsInCollection gets a reference to the given bool and assigns it to the MoreItemsInCollection field.
func (o *FieldsResponse200AllOfAdditionalData) SetMoreItemsInCollection(v bool) {
	o.MoreItemsInCollection = &v
}

func (o FieldsResponse200AllOfAdditionalData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FieldsResponse200AllOfAdditionalData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.MoreItemsInCollection) {
		toSerialize["more_items_in_collection"] = o.MoreItemsInCollection
	}
	return toSerialize, nil
}

type NullableFieldsResponse200AllOfAdditionalData struct {
	value *FieldsResponse200AllOfAdditionalData
	isSet bool
}

func (v NullableFieldsResponse200AllOfAdditionalData) Get() *FieldsResponse200AllOfAdditionalData {
	return v.value
}

func (v *NullableFieldsResponse200AllOfAdditionalData) Set(val *FieldsResponse200AllOfAdditionalData) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldsResponse200AllOfAdditionalData) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldsResponse200AllOfAdditionalData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldsResponse200AllOfAdditionalData(val *FieldsResponse200AllOfAdditionalData) *NullableFieldsResponse200AllOfAdditionalData {
	return &NullableFieldsResponse200AllOfAdditionalData{value: val, isSet: true}
}

func (v NullableFieldsResponse200AllOfAdditionalData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldsResponse200AllOfAdditionalData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


