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

// checks if the GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues{}

// GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues The formatted values of the deals
type GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues struct {
	// The formatted values of the deals
	CURRENCY_ID *string `json:"CURRENCY_ID,omitempty"`
}

// NewGetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues instantiates a new GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues() *GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues {
	this := GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues{}
	return &this
}

// NewGetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValuesWithDefaults instantiates a new GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValuesWithDefaults() *GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues {
	this := GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues{}
	return &this
}

// GetCURRENCY_ID returns the CURRENCY_ID field value if set, zero value otherwise.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues) GetCURRENCY_ID() string {
	if o == nil || IsNil(o.CURRENCY_ID) {
		var ret string
		return ret
	}
	return *o.CURRENCY_ID
}

// GetCURRENCY_IDOk returns a tuple with the CURRENCY_ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues) GetCURRENCY_IDOk() (*string, bool) {
	if o == nil || IsNil(o.CURRENCY_ID) {
		return nil, false
	}
	return o.CURRENCY_ID, true
}

// HasCURRENCY_ID returns a boolean if a field has been set.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues) HasCURRENCY_ID() bool {
	if o != nil && !IsNil(o.CURRENCY_ID) {
		return true
	}

	return false
}

// SetCURRENCY_ID gets a reference to the given string and assigns it to the CURRENCY_ID field.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues) SetCURRENCY_ID(v string) {
	o.CURRENCY_ID = &v
}

func (o GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CURRENCY_ID) {
		toSerialize["CURRENCY_ID"] = o.CURRENCY_ID
	}
	return toSerialize, nil
}

type NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues struct {
	value *GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues
	isSet bool
}

func (v NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues) Get() *GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues {
	return v.value
}

func (v *NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues) Set(val *GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues(val *GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues) *NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues {
	return &NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues{value: val, isSet: true}
}

func (v NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


