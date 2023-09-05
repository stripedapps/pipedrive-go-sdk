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

// checks if the GetNotesResponse200AdditionalDataPagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNotesResponse200AdditionalDataPagination{}

// GetNotesResponse200AdditionalDataPagination The pagination details of the list
type GetNotesResponse200AdditionalDataPagination struct {
	// Pagination start
	Start *int32 `json:"start,omitempty"`
	// Items shown per page
	Limit *int32 `json:"limit,omitempty"`
	// If there are more list items in the collection than displayed or not
	MoreItemsInCollection *bool `json:"more_items_in_collection,omitempty"`
	// Next pagination start
	NextStart *int32 `json:"next_start,omitempty"`
}

// NewGetNotesResponse200AdditionalDataPagination instantiates a new GetNotesResponse200AdditionalDataPagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNotesResponse200AdditionalDataPagination() *GetNotesResponse200AdditionalDataPagination {
	this := GetNotesResponse200AdditionalDataPagination{}
	return &this
}

// NewGetNotesResponse200AdditionalDataPaginationWithDefaults instantiates a new GetNotesResponse200AdditionalDataPagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNotesResponse200AdditionalDataPaginationWithDefaults() *GetNotesResponse200AdditionalDataPagination {
	this := GetNotesResponse200AdditionalDataPagination{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *GetNotesResponse200AdditionalDataPagination) GetStart() int32 {
	if o == nil || IsNil(o.Start) {
		var ret int32
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNotesResponse200AdditionalDataPagination) GetStartOk() (*int32, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *GetNotesResponse200AdditionalDataPagination) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given int32 and assigns it to the Start field.
func (o *GetNotesResponse200AdditionalDataPagination) SetStart(v int32) {
	o.Start = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GetNotesResponse200AdditionalDataPagination) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNotesResponse200AdditionalDataPagination) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GetNotesResponse200AdditionalDataPagination) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *GetNotesResponse200AdditionalDataPagination) SetLimit(v int32) {
	o.Limit = &v
}

// GetMoreItemsInCollection returns the MoreItemsInCollection field value if set, zero value otherwise.
func (o *GetNotesResponse200AdditionalDataPagination) GetMoreItemsInCollection() bool {
	if o == nil || IsNil(o.MoreItemsInCollection) {
		var ret bool
		return ret
	}
	return *o.MoreItemsInCollection
}

// GetMoreItemsInCollectionOk returns a tuple with the MoreItemsInCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNotesResponse200AdditionalDataPagination) GetMoreItemsInCollectionOk() (*bool, bool) {
	if o == nil || IsNil(o.MoreItemsInCollection) {
		return nil, false
	}
	return o.MoreItemsInCollection, true
}

// HasMoreItemsInCollection returns a boolean if a field has been set.
func (o *GetNotesResponse200AdditionalDataPagination) HasMoreItemsInCollection() bool {
	if o != nil && !IsNil(o.MoreItemsInCollection) {
		return true
	}

	return false
}

// SetMoreItemsInCollection gets a reference to the given bool and assigns it to the MoreItemsInCollection field.
func (o *GetNotesResponse200AdditionalDataPagination) SetMoreItemsInCollection(v bool) {
	o.MoreItemsInCollection = &v
}

// GetNextStart returns the NextStart field value if set, zero value otherwise.
func (o *GetNotesResponse200AdditionalDataPagination) GetNextStart() int32 {
	if o == nil || IsNil(o.NextStart) {
		var ret int32
		return ret
	}
	return *o.NextStart
}

// GetNextStartOk returns a tuple with the NextStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNotesResponse200AdditionalDataPagination) GetNextStartOk() (*int32, bool) {
	if o == nil || IsNil(o.NextStart) {
		return nil, false
	}
	return o.NextStart, true
}

// HasNextStart returns a boolean if a field has been set.
func (o *GetNotesResponse200AdditionalDataPagination) HasNextStart() bool {
	if o != nil && !IsNil(o.NextStart) {
		return true
	}

	return false
}

// SetNextStart gets a reference to the given int32 and assigns it to the NextStart field.
func (o *GetNotesResponse200AdditionalDataPagination) SetNextStart(v int32) {
	o.NextStart = &v
}

func (o GetNotesResponse200AdditionalDataPagination) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNotesResponse200AdditionalDataPagination) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.NextStart) {
		toSerialize["next_start"] = o.NextStart
	}
	return toSerialize, nil
}

type NullableGetNotesResponse200AdditionalDataPagination struct {
	value *GetNotesResponse200AdditionalDataPagination
	isSet bool
}

func (v NullableGetNotesResponse200AdditionalDataPagination) Get() *GetNotesResponse200AdditionalDataPagination {
	return v.value
}

func (v *NullableGetNotesResponse200AdditionalDataPagination) Set(val *GetNotesResponse200AdditionalDataPagination) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNotesResponse200AdditionalDataPagination) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNotesResponse200AdditionalDataPagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNotesResponse200AdditionalDataPagination(val *GetNotesResponse200AdditionalDataPagination) *NullableGetNotesResponse200AdditionalDataPagination {
	return &NullableGetNotesResponse200AdditionalDataPagination{value: val, isSet: true}
}

func (v NullableGetNotesResponse200AdditionalDataPagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNotesResponse200AdditionalDataPagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


