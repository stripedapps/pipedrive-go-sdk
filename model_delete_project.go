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

// checks if the DeleteProject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteProject{}

// DeleteProject struct for DeleteProject
type DeleteProject struct {
	// If the request was successful or not
	Success *bool `json:"success,omitempty"`
	Data *DeleteProjectData `json:"data,omitempty"`
}

// NewDeleteProject instantiates a new DeleteProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteProject() *DeleteProject {
	this := DeleteProject{}
	return &this
}

// NewDeleteProjectWithDefaults instantiates a new DeleteProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteProjectWithDefaults() *DeleteProject {
	this := DeleteProject{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *DeleteProject) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteProject) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *DeleteProject) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *DeleteProject) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DeleteProject) GetData() DeleteProjectData {
	if o == nil || IsNil(o.Data) {
		var ret DeleteProjectData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteProject) GetDataOk() (*DeleteProjectData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DeleteProject) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DeleteProjectData and assigns it to the Data field.
func (o *DeleteProject) SetData(v DeleteProjectData) {
	o.Data = &v
}

func (o DeleteProject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteProject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableDeleteProject struct {
	value *DeleteProject
	isSet bool
}

func (v NullableDeleteProject) Get() *DeleteProject {
	return v.value
}

func (v *NullableDeleteProject) Set(val *DeleteProject) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteProject) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteProject(val *DeleteProject) *NullableDeleteProject {
	return &NullableDeleteProject{value: val, isSet: true}
}

func (v NullableDeleteProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


