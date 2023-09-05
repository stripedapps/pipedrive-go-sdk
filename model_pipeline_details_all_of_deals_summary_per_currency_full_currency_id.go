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

// checks if the PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID{}

// PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID The currency summary. This parameter is dynamic and changes according to `currency_id` value.
type PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID struct {
	// Deals count per currency
	Count *int32 `json:"count,omitempty"`
	// Deals value per currency
	Value *int32 `json:"value,omitempty"`
}

// NewPipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID instantiates a new PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID() *PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID {
	this := PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID{}
	return &this
}

// NewPipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYIDWithDefaults instantiates a new PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYIDWithDefaults() *PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID {
	this := PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID) SetCount(v int32) {
	o.Count = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID) SetValue(v int32) {
	o.Value = &v
}

func (o PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullablePipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID struct {
	value *PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID
	isSet bool
}

func (v NullablePipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID) Get() *PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID {
	return v.value
}

func (v *NullablePipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID) Set(val *PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID(val *PipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID) *NullablePipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID {
	return &NullablePipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID{value: val, isSet: true}
}

func (v NullablePipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineDetailsAllOfDealsSummaryPerCurrencyFullCURRENCYID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


