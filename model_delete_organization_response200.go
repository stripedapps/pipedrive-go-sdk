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

// checks if the DeleteOrganizationResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteOrganizationResponse200{}

// DeleteOrganizationResponse200 struct for DeleteOrganizationResponse200
type DeleteOrganizationResponse200 struct {
	// If the request was successful or not
	Success *bool `json:"success,omitempty"`
	Data *DeleteOrganizationResponse200Data `json:"data,omitempty"`
}

// NewDeleteOrganizationResponse200 instantiates a new DeleteOrganizationResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteOrganizationResponse200() *DeleteOrganizationResponse200 {
	this := DeleteOrganizationResponse200{}
	return &this
}

// NewDeleteOrganizationResponse200WithDefaults instantiates a new DeleteOrganizationResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteOrganizationResponse200WithDefaults() *DeleteOrganizationResponse200 {
	this := DeleteOrganizationResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *DeleteOrganizationResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteOrganizationResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *DeleteOrganizationResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *DeleteOrganizationResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DeleteOrganizationResponse200) GetData() DeleteOrganizationResponse200Data {
	if o == nil || IsNil(o.Data) {
		var ret DeleteOrganizationResponse200Data
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteOrganizationResponse200) GetDataOk() (*DeleteOrganizationResponse200Data, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DeleteOrganizationResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DeleteOrganizationResponse200Data and assigns it to the Data field.
func (o *DeleteOrganizationResponse200) SetData(v DeleteOrganizationResponse200Data) {
	o.Data = &v
}

func (o DeleteOrganizationResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteOrganizationResponse200) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableDeleteOrganizationResponse200 struct {
	value *DeleteOrganizationResponse200
	isSet bool
}

func (v NullableDeleteOrganizationResponse200) Get() *DeleteOrganizationResponse200 {
	return v.value
}

func (v *NullableDeleteOrganizationResponse200) Set(val *DeleteOrganizationResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteOrganizationResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteOrganizationResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteOrganizationResponse200(val *DeleteOrganizationResponse200) *NullableDeleteOrganizationResponse200 {
	return &NullableDeleteOrganizationResponse200{value: val, isSet: true}
}

func (v NullableDeleteOrganizationResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteOrganizationResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


