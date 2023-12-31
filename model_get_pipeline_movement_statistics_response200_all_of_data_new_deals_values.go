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

// checks if the GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues{}

// GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues The values of the deals
type GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues struct {
	// The value of the deals
	CURRENCY_ID *int32 `json:"CURRENCY_ID,omitempty"`
}

// NewGetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues instantiates a new GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues() *GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues {
	this := GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues{}
	return &this
}

// NewGetPipelineMovementStatisticsResponse200AllOfDataNewDealsValuesWithDefaults instantiates a new GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPipelineMovementStatisticsResponse200AllOfDataNewDealsValuesWithDefaults() *GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues {
	this := GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues{}
	return &this
}

// GetCURRENCY_ID returns the CURRENCY_ID field value if set, zero value otherwise.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues) GetCURRENCY_ID() int32 {
	if o == nil || IsNil(o.CURRENCY_ID) {
		var ret int32
		return ret
	}
	return *o.CURRENCY_ID
}

// GetCURRENCY_IDOk returns a tuple with the CURRENCY_ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues) GetCURRENCY_IDOk() (*int32, bool) {
	if o == nil || IsNil(o.CURRENCY_ID) {
		return nil, false
	}
	return o.CURRENCY_ID, true
}

// HasCURRENCY_ID returns a boolean if a field has been set.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues) HasCURRENCY_ID() bool {
	if o != nil && !IsNil(o.CURRENCY_ID) {
		return true
	}

	return false
}

// SetCURRENCY_ID gets a reference to the given int32 and assigns it to the CURRENCY_ID field.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues) SetCURRENCY_ID(v int32) {
	o.CURRENCY_ID = &v
}

func (o GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CURRENCY_ID) {
		toSerialize["CURRENCY_ID"] = o.CURRENCY_ID
	}
	return toSerialize, nil
}

type NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues struct {
	value *GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues
	isSet bool
}

func (v NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues) Get() *GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues {
	return v.value
}

func (v *NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues) Set(val *GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues(val *GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues) *NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues {
	return &NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues{value: val, isSet: true}
}

func (v NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


