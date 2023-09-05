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

// checks if the PipelineDetailsAllOfDealsSummaryPerStagesSTAGEID type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PipelineDetailsAllOfDealsSummaryPerStagesSTAGEID{}

// PipelineDetailsAllOfDealsSummaryPerStagesSTAGEID The currency summaries per stage. This parameter is dynamic and changes according to `stage_id` value.
type PipelineDetailsAllOfDealsSummaryPerStagesSTAGEID struct {
	CURRENCY_ID *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID `json:"CURRENCY_ID,omitempty"`
}

// NewPipelineDetailsAllOfDealsSummaryPerStagesSTAGEID instantiates a new PipelineDetailsAllOfDealsSummaryPerStagesSTAGEID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineDetailsAllOfDealsSummaryPerStagesSTAGEID() *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEID {
	this := PipelineDetailsAllOfDealsSummaryPerStagesSTAGEID{}
	return &this
}

// NewPipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDWithDefaults instantiates a new PipelineDetailsAllOfDealsSummaryPerStagesSTAGEID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDWithDefaults() *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEID {
	this := PipelineDetailsAllOfDealsSummaryPerStagesSTAGEID{}
	return &this
}

// GetCURRENCY_ID returns the CURRENCY_ID field value if set, zero value otherwise.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEID) GetCURRENCY_ID() PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID {
	if o == nil || IsNil(o.CURRENCY_ID) {
		var ret PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID
		return ret
	}
	return *o.CURRENCY_ID
}

// GetCURRENCY_IDOk returns a tuple with the CURRENCY_ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEID) GetCURRENCY_IDOk() (*PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID, bool) {
	if o == nil || IsNil(o.CURRENCY_ID) {
		return nil, false
	}
	return o.CURRENCY_ID, true
}

// HasCURRENCY_ID returns a boolean if a field has been set.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEID) HasCURRENCY_ID() bool {
	if o != nil && !IsNil(o.CURRENCY_ID) {
		return true
	}

	return false
}

// SetCURRENCY_ID gets a reference to the given PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID and assigns it to the CURRENCY_ID field.
func (o *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEID) SetCURRENCY_ID(v PipelineDetailsAllOfDealsSummaryPerStagesSTAGEIDCURRENCYID) {
	o.CURRENCY_ID = &v
}

func (o PipelineDetailsAllOfDealsSummaryPerStagesSTAGEID) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PipelineDetailsAllOfDealsSummaryPerStagesSTAGEID) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CURRENCY_ID) {
		toSerialize["CURRENCY_ID"] = o.CURRENCY_ID
	}
	return toSerialize, nil
}

type NullablePipelineDetailsAllOfDealsSummaryPerStagesSTAGEID struct {
	value *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEID
	isSet bool
}

func (v NullablePipelineDetailsAllOfDealsSummaryPerStagesSTAGEID) Get() *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEID {
	return v.value
}

func (v *NullablePipelineDetailsAllOfDealsSummaryPerStagesSTAGEID) Set(val *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEID) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineDetailsAllOfDealsSummaryPerStagesSTAGEID) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineDetailsAllOfDealsSummaryPerStagesSTAGEID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineDetailsAllOfDealsSummaryPerStagesSTAGEID(val *PipelineDetailsAllOfDealsSummaryPerStagesSTAGEID) *NullablePipelineDetailsAllOfDealsSummaryPerStagesSTAGEID {
	return &NullablePipelineDetailsAllOfDealsSummaryPerStagesSTAGEID{value: val, isSet: true}
}

func (v NullablePipelineDetailsAllOfDealsSummaryPerStagesSTAGEID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineDetailsAllOfDealsSummaryPerStagesSTAGEID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


