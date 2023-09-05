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

// checks if the PutRolePipelinesBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutRolePipelinesBody{}

// PutRolePipelinesBody struct for PutRolePipelinesBody
type PutRolePipelinesBody struct {
	// The pipeline IDs to make the pipelines visible (add) and/or hidden (remove) for the specified role. It requires the following JSON structure: `{ \"add\": \"[1]\", \"remove\": \"[3, 4]\" }`.
	VisiblePipelineIds map[string]interface{} `json:"visible_pipeline_ids"`
}

// NewPutRolePipelinesBody instantiates a new PutRolePipelinesBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutRolePipelinesBody(visiblePipelineIds map[string]interface{}) *PutRolePipelinesBody {
	this := PutRolePipelinesBody{}
	this.VisiblePipelineIds = visiblePipelineIds
	return &this
}

// NewPutRolePipelinesBodyWithDefaults instantiates a new PutRolePipelinesBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutRolePipelinesBodyWithDefaults() *PutRolePipelinesBody {
	this := PutRolePipelinesBody{}
	return &this
}

// GetVisiblePipelineIds returns the VisiblePipelineIds field value
func (o *PutRolePipelinesBody) GetVisiblePipelineIds() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.VisiblePipelineIds
}

// GetVisiblePipelineIdsOk returns a tuple with the VisiblePipelineIds field value
// and a boolean to check if the value has been set.
func (o *PutRolePipelinesBody) GetVisiblePipelineIdsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.VisiblePipelineIds, true
}

// SetVisiblePipelineIds sets field value
func (o *PutRolePipelinesBody) SetVisiblePipelineIds(v map[string]interface{}) {
	o.VisiblePipelineIds = v
}

func (o PutRolePipelinesBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutRolePipelinesBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["visible_pipeline_ids"] = o.VisiblePipelineIds
	return toSerialize, nil
}

type NullablePutRolePipelinesBody struct {
	value *PutRolePipelinesBody
	isSet bool
}

func (v NullablePutRolePipelinesBody) Get() *PutRolePipelinesBody {
	return v.value
}

func (v *NullablePutRolePipelinesBody) Set(val *PutRolePipelinesBody) {
	v.value = val
	v.isSet = true
}

func (v NullablePutRolePipelinesBody) IsSet() bool {
	return v.isSet
}

func (v *NullablePutRolePipelinesBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutRolePipelinesBody(val *PutRolePipelinesBody) *NullablePutRolePipelinesBody {
	return &NullablePutRolePipelinesBody{value: val, isSet: true}
}

func (v NullablePutRolePipelinesBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutRolePipelinesBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


