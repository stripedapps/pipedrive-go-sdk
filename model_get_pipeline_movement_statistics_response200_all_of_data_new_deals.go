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

// checks if the GetPipelineMovementStatisticsResponse200AllOfDataNewDeals type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPipelineMovementStatisticsResponse200AllOfDataNewDeals{}

// GetPipelineMovementStatisticsResponse200AllOfDataNewDeals Deals summary
type GetPipelineMovementStatisticsResponse200AllOfDataNewDeals struct {
	// The count of the deals
	Count *int32 `json:"count,omitempty"`
	// The IDs of the deals that have been moved
	DealsIds []int32 `json:"deals_ids,omitempty"`
	Values *GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues `json:"values,omitempty"`
	FormattedValues *GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues `json:"formatted_values,omitempty"`
}

// NewGetPipelineMovementStatisticsResponse200AllOfDataNewDeals instantiates a new GetPipelineMovementStatisticsResponse200AllOfDataNewDeals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPipelineMovementStatisticsResponse200AllOfDataNewDeals() *GetPipelineMovementStatisticsResponse200AllOfDataNewDeals {
	this := GetPipelineMovementStatisticsResponse200AllOfDataNewDeals{}
	return &this
}

// NewGetPipelineMovementStatisticsResponse200AllOfDataNewDealsWithDefaults instantiates a new GetPipelineMovementStatisticsResponse200AllOfDataNewDeals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPipelineMovementStatisticsResponse200AllOfDataNewDealsWithDefaults() *GetPipelineMovementStatisticsResponse200AllOfDataNewDeals {
	this := GetPipelineMovementStatisticsResponse200AllOfDataNewDeals{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDeals) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDeals) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDeals) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDeals) SetCount(v int32) {
	o.Count = &v
}

// GetDealsIds returns the DealsIds field value if set, zero value otherwise.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDeals) GetDealsIds() []int32 {
	if o == nil || IsNil(o.DealsIds) {
		var ret []int32
		return ret
	}
	return o.DealsIds
}

// GetDealsIdsOk returns a tuple with the DealsIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDeals) GetDealsIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.DealsIds) {
		return nil, false
	}
	return o.DealsIds, true
}

// HasDealsIds returns a boolean if a field has been set.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDeals) HasDealsIds() bool {
	if o != nil && !IsNil(o.DealsIds) {
		return true
	}

	return false
}

// SetDealsIds gets a reference to the given []int32 and assigns it to the DealsIds field.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDeals) SetDealsIds(v []int32) {
	o.DealsIds = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDeals) GetValues() GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues {
	if o == nil || IsNil(o.Values) {
		var ret GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDeals) GetValuesOk() (*GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDeals) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues and assigns it to the Values field.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDeals) SetValues(v GetPipelineMovementStatisticsResponse200AllOfDataNewDealsValues) {
	o.Values = &v
}

// GetFormattedValues returns the FormattedValues field value if set, zero value otherwise.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDeals) GetFormattedValues() GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues {
	if o == nil || IsNil(o.FormattedValues) {
		var ret GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues
		return ret
	}
	return *o.FormattedValues
}

// GetFormattedValuesOk returns a tuple with the FormattedValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDeals) GetFormattedValuesOk() (*GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues, bool) {
	if o == nil || IsNil(o.FormattedValues) {
		return nil, false
	}
	return o.FormattedValues, true
}

// HasFormattedValues returns a boolean if a field has been set.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDeals) HasFormattedValues() bool {
	if o != nil && !IsNil(o.FormattedValues) {
		return true
	}

	return false
}

// SetFormattedValues gets a reference to the given GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues and assigns it to the FormattedValues field.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataNewDeals) SetFormattedValues(v GetPipelineMovementStatisticsResponse200AllOfDataNewDealsFormattedValues) {
	o.FormattedValues = &v
}

func (o GetPipelineMovementStatisticsResponse200AllOfDataNewDeals) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPipelineMovementStatisticsResponse200AllOfDataNewDeals) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.DealsIds) {
		toSerialize["deals_ids"] = o.DealsIds
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	if !IsNil(o.FormattedValues) {
		toSerialize["formatted_values"] = o.FormattedValues
	}
	return toSerialize, nil
}

type NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDeals struct {
	value *GetPipelineMovementStatisticsResponse200AllOfDataNewDeals
	isSet bool
}

func (v NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDeals) Get() *GetPipelineMovementStatisticsResponse200AllOfDataNewDeals {
	return v.value
}

func (v *NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDeals) Set(val *GetPipelineMovementStatisticsResponse200AllOfDataNewDeals) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDeals) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDeals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPipelineMovementStatisticsResponse200AllOfDataNewDeals(val *GetPipelineMovementStatisticsResponse200AllOfDataNewDeals) *NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDeals {
	return &NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDeals{value: val, isSet: true}
}

func (v NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDeals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPipelineMovementStatisticsResponse200AllOfDataNewDeals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


