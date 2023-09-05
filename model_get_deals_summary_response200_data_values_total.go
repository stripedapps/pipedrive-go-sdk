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

// checks if the GetDealsSummaryResponse200DataValuesTotal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDealsSummaryResponse200DataValuesTotal{}

// GetDealsSummaryResponse200DataValuesTotal The total values of the deals grouped by deal currency
type GetDealsSummaryResponse200DataValuesTotal struct {
	// The total value of deals in the deal currency group
	Value *float32 `json:"value,omitempty"`
	// The number of deals in the deal currency group
	Count *int32 `json:"count,omitempty"`
	// The total value of deals converted into the company default currency
	ValueConverted *float32 `json:"value_converted,omitempty"`
	// The total value of deals formatted with deal currency. E.g. €50
	ValueFormatted *string `json:"value_formatted,omitempty"`
	// The value_converted formatted with deal currency. E.g. US$50.10
	ValueConvertedFormatted *string `json:"value_converted_formatted,omitempty"`
}

// NewGetDealsSummaryResponse200DataValuesTotal instantiates a new GetDealsSummaryResponse200DataValuesTotal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDealsSummaryResponse200DataValuesTotal() *GetDealsSummaryResponse200DataValuesTotal {
	this := GetDealsSummaryResponse200DataValuesTotal{}
	return &this
}

// NewGetDealsSummaryResponse200DataValuesTotalWithDefaults instantiates a new GetDealsSummaryResponse200DataValuesTotal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDealsSummaryResponse200DataValuesTotalWithDefaults() *GetDealsSummaryResponse200DataValuesTotal {
	this := GetDealsSummaryResponse200DataValuesTotal{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetDealsSummaryResponse200DataValuesTotal) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDealsSummaryResponse200DataValuesTotal) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetDealsSummaryResponse200DataValuesTotal) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *GetDealsSummaryResponse200DataValuesTotal) SetValue(v float32) {
	o.Value = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GetDealsSummaryResponse200DataValuesTotal) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDealsSummaryResponse200DataValuesTotal) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GetDealsSummaryResponse200DataValuesTotal) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *GetDealsSummaryResponse200DataValuesTotal) SetCount(v int32) {
	o.Count = &v
}

// GetValueConverted returns the ValueConverted field value if set, zero value otherwise.
func (o *GetDealsSummaryResponse200DataValuesTotal) GetValueConverted() float32 {
	if o == nil || IsNil(o.ValueConverted) {
		var ret float32
		return ret
	}
	return *o.ValueConverted
}

// GetValueConvertedOk returns a tuple with the ValueConverted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDealsSummaryResponse200DataValuesTotal) GetValueConvertedOk() (*float32, bool) {
	if o == nil || IsNil(o.ValueConverted) {
		return nil, false
	}
	return o.ValueConverted, true
}

// HasValueConverted returns a boolean if a field has been set.
func (o *GetDealsSummaryResponse200DataValuesTotal) HasValueConverted() bool {
	if o != nil && !IsNil(o.ValueConverted) {
		return true
	}

	return false
}

// SetValueConverted gets a reference to the given float32 and assigns it to the ValueConverted field.
func (o *GetDealsSummaryResponse200DataValuesTotal) SetValueConverted(v float32) {
	o.ValueConverted = &v
}

// GetValueFormatted returns the ValueFormatted field value if set, zero value otherwise.
func (o *GetDealsSummaryResponse200DataValuesTotal) GetValueFormatted() string {
	if o == nil || IsNil(o.ValueFormatted) {
		var ret string
		return ret
	}
	return *o.ValueFormatted
}

// GetValueFormattedOk returns a tuple with the ValueFormatted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDealsSummaryResponse200DataValuesTotal) GetValueFormattedOk() (*string, bool) {
	if o == nil || IsNil(o.ValueFormatted) {
		return nil, false
	}
	return o.ValueFormatted, true
}

// HasValueFormatted returns a boolean if a field has been set.
func (o *GetDealsSummaryResponse200DataValuesTotal) HasValueFormatted() bool {
	if o != nil && !IsNil(o.ValueFormatted) {
		return true
	}

	return false
}

// SetValueFormatted gets a reference to the given string and assigns it to the ValueFormatted field.
func (o *GetDealsSummaryResponse200DataValuesTotal) SetValueFormatted(v string) {
	o.ValueFormatted = &v
}

// GetValueConvertedFormatted returns the ValueConvertedFormatted field value if set, zero value otherwise.
func (o *GetDealsSummaryResponse200DataValuesTotal) GetValueConvertedFormatted() string {
	if o == nil || IsNil(o.ValueConvertedFormatted) {
		var ret string
		return ret
	}
	return *o.ValueConvertedFormatted
}

// GetValueConvertedFormattedOk returns a tuple with the ValueConvertedFormatted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDealsSummaryResponse200DataValuesTotal) GetValueConvertedFormattedOk() (*string, bool) {
	if o == nil || IsNil(o.ValueConvertedFormatted) {
		return nil, false
	}
	return o.ValueConvertedFormatted, true
}

// HasValueConvertedFormatted returns a boolean if a field has been set.
func (o *GetDealsSummaryResponse200DataValuesTotal) HasValueConvertedFormatted() bool {
	if o != nil && !IsNil(o.ValueConvertedFormatted) {
		return true
	}

	return false
}

// SetValueConvertedFormatted gets a reference to the given string and assigns it to the ValueConvertedFormatted field.
func (o *GetDealsSummaryResponse200DataValuesTotal) SetValueConvertedFormatted(v string) {
	o.ValueConvertedFormatted = &v
}

func (o GetDealsSummaryResponse200DataValuesTotal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDealsSummaryResponse200DataValuesTotal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.ValueConverted) {
		toSerialize["value_converted"] = o.ValueConverted
	}
	if !IsNil(o.ValueFormatted) {
		toSerialize["value_formatted"] = o.ValueFormatted
	}
	if !IsNil(o.ValueConvertedFormatted) {
		toSerialize["value_converted_formatted"] = o.ValueConvertedFormatted
	}
	return toSerialize, nil
}

type NullableGetDealsSummaryResponse200DataValuesTotal struct {
	value *GetDealsSummaryResponse200DataValuesTotal
	isSet bool
}

func (v NullableGetDealsSummaryResponse200DataValuesTotal) Get() *GetDealsSummaryResponse200DataValuesTotal {
	return v.value
}

func (v *NullableGetDealsSummaryResponse200DataValuesTotal) Set(val *GetDealsSummaryResponse200DataValuesTotal) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDealsSummaryResponse200DataValuesTotal) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDealsSummaryResponse200DataValuesTotal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDealsSummaryResponse200DataValuesTotal(val *GetDealsSummaryResponse200DataValuesTotal) *NullableGetDealsSummaryResponse200DataValuesTotal {
	return &NullableGetDealsSummaryResponse200DataValuesTotal{value: val, isSet: true}
}

func (v NullableGetDealsSummaryResponse200DataValuesTotal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDealsSummaryResponse200DataValuesTotal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


