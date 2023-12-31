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

// checks if the CreateUpdateDeleteActivityTypeResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUpdateDeleteActivityTypeResponse200{}

// CreateUpdateDeleteActivityTypeResponse200 struct for CreateUpdateDeleteActivityTypeResponse200
type CreateUpdateDeleteActivityTypeResponse200 struct {
	// If the response is successful or not
	Success *bool `json:"success,omitempty"`
	Data *GetActivityTypesResponse200AllOfDataInner `json:"data,omitempty"`
}

// NewCreateUpdateDeleteActivityTypeResponse200 instantiates a new CreateUpdateDeleteActivityTypeResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUpdateDeleteActivityTypeResponse200() *CreateUpdateDeleteActivityTypeResponse200 {
	this := CreateUpdateDeleteActivityTypeResponse200{}
	return &this
}

// NewCreateUpdateDeleteActivityTypeResponse200WithDefaults instantiates a new CreateUpdateDeleteActivityTypeResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUpdateDeleteActivityTypeResponse200WithDefaults() *CreateUpdateDeleteActivityTypeResponse200 {
	this := CreateUpdateDeleteActivityTypeResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *CreateUpdateDeleteActivityTypeResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateDeleteActivityTypeResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *CreateUpdateDeleteActivityTypeResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *CreateUpdateDeleteActivityTypeResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateUpdateDeleteActivityTypeResponse200) GetData() GetActivityTypesResponse200AllOfDataInner {
	if o == nil || IsNil(o.Data) {
		var ret GetActivityTypesResponse200AllOfDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUpdateDeleteActivityTypeResponse200) GetDataOk() (*GetActivityTypesResponse200AllOfDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateUpdateDeleteActivityTypeResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetActivityTypesResponse200AllOfDataInner and assigns it to the Data field.
func (o *CreateUpdateDeleteActivityTypeResponse200) SetData(v GetActivityTypesResponse200AllOfDataInner) {
	o.Data = &v
}

func (o CreateUpdateDeleteActivityTypeResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUpdateDeleteActivityTypeResponse200) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCreateUpdateDeleteActivityTypeResponse200 struct {
	value *CreateUpdateDeleteActivityTypeResponse200
	isSet bool
}

func (v NullableCreateUpdateDeleteActivityTypeResponse200) Get() *CreateUpdateDeleteActivityTypeResponse200 {
	return v.value
}

func (v *NullableCreateUpdateDeleteActivityTypeResponse200) Set(val *CreateUpdateDeleteActivityTypeResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUpdateDeleteActivityTypeResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUpdateDeleteActivityTypeResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUpdateDeleteActivityTypeResponse200(val *CreateUpdateDeleteActivityTypeResponse200) *NullableCreateUpdateDeleteActivityTypeResponse200 {
	return &NullableCreateUpdateDeleteActivityTypeResponse200{value: val, isSet: true}
}

func (v NullableCreateUpdateDeleteActivityTypeResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUpdateDeleteActivityTypeResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


