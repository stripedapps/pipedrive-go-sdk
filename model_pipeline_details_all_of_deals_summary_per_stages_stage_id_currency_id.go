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

// checks if the PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID{}

// PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID The currency summary. This parameter is dynamic and changes according to `currency_id` value.
type PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID struct {
	// Deals count per currency
	Count *int32 `json:"count,omitempty"`
	// Deals value per currency
	Value *int32 `json:"value,omitempty"`
	// Deals value formatted per currency
	ValueFormatted *string `json:"value_formatted,omitempty"`
	// Deals weighted value per currency
	WeightedValue *int32 `json:"weighted_value,omitempty"`
	// Deals weighted value formatted per currency
	WeightedValueFormatted *string `json:"weighted_value_formatted,omitempty"`
}

// NewPipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID instantiates a new PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID() *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID {
	this := PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID{}
	return &this
}

// NewPipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYIDWithDefaults instantiates a new PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYIDWithDefaults() *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID {
	this := PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) SetCount(v int32) {
	o.Count = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) SetValue(v int32) {
	o.Value = &v
}

// GetValueFormatted returns the ValueFormatted field value if set, zero value otherwise.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) GetValueFormatted() string {
	if o == nil || IsNil(o.ValueFormatted) {
		var ret string
		return ret
	}
	return *o.ValueFormatted
}

// GetValueFormattedOk returns a tuple with the ValueFormatted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) GetValueFormattedOk() (*string, bool) {
	if o == nil || IsNil(o.ValueFormatted) {
		return nil, false
	}
	return o.ValueFormatted, true
}

// HasValueFormatted returns a boolean if a field has been set.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) HasValueFormatted() bool {
	if o != nil && !IsNil(o.ValueFormatted) {
		return true
	}

	return false
}

// SetValueFormatted gets a reference to the given string and assigns it to the ValueFormatted field.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) SetValueFormatted(v string) {
	o.ValueFormatted = &v
}

// GetWeightedValue returns the WeightedValue field value if set, zero value otherwise.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) GetWeightedValue() int32 {
	if o == nil || IsNil(o.WeightedValue) {
		var ret int32
		return ret
	}
	return *o.WeightedValue
}

// GetWeightedValueOk returns a tuple with the WeightedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) GetWeightedValueOk() (*int32, bool) {
	if o == nil || IsNil(o.WeightedValue) {
		return nil, false
	}
	return o.WeightedValue, true
}

// HasWeightedValue returns a boolean if a field has been set.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) HasWeightedValue() bool {
	if o != nil && !IsNil(o.WeightedValue) {
		return true
	}

	return false
}

// SetWeightedValue gets a reference to the given int32 and assigns it to the WeightedValue field.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) SetWeightedValue(v int32) {
	o.WeightedValue = &v
}

// GetWeightedValueFormatted returns the WeightedValueFormatted field value if set, zero value otherwise.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) GetWeightedValueFormatted() string {
	if o == nil || IsNil(o.WeightedValueFormatted) {
		var ret string
		return ret
	}
	return *o.WeightedValueFormatted
}

// GetWeightedValueFormattedOk returns a tuple with the WeightedValueFormatted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) GetWeightedValueFormattedOk() (*string, bool) {
	if o == nil || IsNil(o.WeightedValueFormatted) {
		return nil, false
	}
	return o.WeightedValueFormatted, true
}

// HasWeightedValueFormatted returns a boolean if a field has been set.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) HasWeightedValueFormatted() bool {
	if o != nil && !IsNil(o.WeightedValueFormatted) {
		return true
	}

	return false
}

// SetWeightedValueFormatted gets a reference to the given string and assigns it to the WeightedValueFormatted field.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) SetWeightedValueFormatted(v string) {
	o.WeightedValueFormatted = &v
}

func (o PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.ValueFormatted) {
		toSerialize["value_formatted"] = o.ValueFormatted
	}
	if !IsNil(o.WeightedValue) {
		toSerialize["weighted_value"] = o.WeightedValue
	}
	if !IsNil(o.WeightedValueFormatted) {
		toSerialize["weighted_value_formatted"] = o.WeightedValueFormatted
	}
	return toSerialize, nil
}

type NullablePipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID struct {
	value *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID
	isSet bool
}

func (v NullablePipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) Get() *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID {
	return v.value
}

func (v *NullablePipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) Set(val *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID(val *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) *NullablePipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID {
	return &NullablePipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID{value: val, isSet: true}
}

func (v NullablePipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


