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

// checks if the GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner{}

// GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner The moved deals average age by the stage
type GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner struct {
	// The stage ID
	StageId *int32 `json:"stage_id,omitempty"`
	// The average deals age in specific stage
	Value *int32 `json:"value,omitempty"`
}

// NewGetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner instantiates a new GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner() *GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner {
	this := GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner{}
	return &this
}

// NewGetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInnerWithDefaults instantiates a new GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInnerWithDefaults() *GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner {
	this := GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner{}
	return &this
}

// GetStageId returns the StageId field value if set, zero value otherwise.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner) GetStageId() int32 {
	if o == nil || IsNil(o.StageId) {
		var ret int32
		return ret
	}
	return *o.StageId
}

// GetStageIdOk returns a tuple with the StageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner) GetStageIdOk() (*int32, bool) {
	if o == nil || IsNil(o.StageId) {
		return nil, false
	}
	return o.StageId, true
}

// HasStageId returns a boolean if a field has been set.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner) HasStageId() bool {
	if o != nil && !IsNil(o.StageId) {
		return true
	}

	return false
}

// SetStageId gets a reference to the given int32 and assigns it to the StageId field.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner) SetStageId(v int32) {
	o.StageId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner) SetValue(v int32) {
	o.Value = &v
}

func (o GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StageId) {
		toSerialize["stage_id"] = o.StageId
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableGetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner struct {
	value *GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner
	isSet bool
}

func (v NullableGetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner) Get() *GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner {
	return v.value
}

func (v *NullableGetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner) Set(val *GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner(val *GetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner) *NullableGetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner {
	return &NullableGetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner{value: val, isSet: true}
}

func (v NullableGetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPipelineMovementStatisticsResponse200AllOfDataAverageAgeInDaysByStagesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


