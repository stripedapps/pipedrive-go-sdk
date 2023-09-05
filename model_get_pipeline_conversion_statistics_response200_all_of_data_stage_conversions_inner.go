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

// checks if the GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner{}

// GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner struct for GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner
type GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner struct {
	// The stage ID from where conversion starts
	FromStageId *int32 `json:"from_stage_id,omitempty"`
	// The stage ID to where conversion ends
	ToStageId *int32 `json:"to_stage_id,omitempty"`
	// The conversion rate
	ConversionRate *int32 `json:"conversion_rate,omitempty"`
}

// NewGetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner instantiates a new GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner() *GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner {
	this := GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner{}
	return &this
}

// NewGetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInnerWithDefaults instantiates a new GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInnerWithDefaults() *GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner {
	this := GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner{}
	return &this
}

// GetFromStageId returns the FromStageId field value if set, zero value otherwise.
func (o *GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner) GetFromStageId() int32 {
	if o == nil || IsNil(o.FromStageId) {
		var ret int32
		return ret
	}
	return *o.FromStageId
}

// GetFromStageIdOk returns a tuple with the FromStageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner) GetFromStageIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FromStageId) {
		return nil, false
	}
	return o.FromStageId, true
}

// HasFromStageId returns a boolean if a field has been set.
func (o *GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner) HasFromStageId() bool {
	if o != nil && !IsNil(o.FromStageId) {
		return true
	}

	return false
}

// SetFromStageId gets a reference to the given int32 and assigns it to the FromStageId field.
func (o *GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner) SetFromStageId(v int32) {
	o.FromStageId = &v
}

// GetToStageId returns the ToStageId field value if set, zero value otherwise.
func (o *GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner) GetToStageId() int32 {
	if o == nil || IsNil(o.ToStageId) {
		var ret int32
		return ret
	}
	return *o.ToStageId
}

// GetToStageIdOk returns a tuple with the ToStageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner) GetToStageIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ToStageId) {
		return nil, false
	}
	return o.ToStageId, true
}

// HasToStageId returns a boolean if a field has been set.
func (o *GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner) HasToStageId() bool {
	if o != nil && !IsNil(o.ToStageId) {
		return true
	}

	return false
}

// SetToStageId gets a reference to the given int32 and assigns it to the ToStageId field.
func (o *GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner) SetToStageId(v int32) {
	o.ToStageId = &v
}

// GetConversionRate returns the ConversionRate field value if set, zero value otherwise.
func (o *GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner) GetConversionRate() int32 {
	if o == nil || IsNil(o.ConversionRate) {
		var ret int32
		return ret
	}
	return *o.ConversionRate
}

// GetConversionRateOk returns a tuple with the ConversionRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner) GetConversionRateOk() (*int32, bool) {
	if o == nil || IsNil(o.ConversionRate) {
		return nil, false
	}
	return o.ConversionRate, true
}

// HasConversionRate returns a boolean if a field has been set.
func (o *GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner) HasConversionRate() bool {
	if o != nil && !IsNil(o.ConversionRate) {
		return true
	}

	return false
}

// SetConversionRate gets a reference to the given int32 and assigns it to the ConversionRate field.
func (o *GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner) SetConversionRate(v int32) {
	o.ConversionRate = &v
}

func (o GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FromStageId) {
		toSerialize["from_stage_id"] = o.FromStageId
	}
	if !IsNil(o.ToStageId) {
		toSerialize["to_stage_id"] = o.ToStageId
	}
	if !IsNil(o.ConversionRate) {
		toSerialize["conversion_rate"] = o.ConversionRate
	}
	return toSerialize, nil
}

type NullableGetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner struct {
	value *GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner
	isSet bool
}

func (v NullableGetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner) Get() *GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner {
	return v.value
}

func (v *NullableGetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner) Set(val *GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner(val *GetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner) *NullableGetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner {
	return &NullableGetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner{value: val, isSet: true}
}

func (v NullableGetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPipelineConversionStatisticsResponse200AllOfDataStageConversionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


