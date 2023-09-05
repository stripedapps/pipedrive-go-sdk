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

// checks if the ListPersonsResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPersonsResponse200{}

// ListPersonsResponse200 struct for ListPersonsResponse200
type ListPersonsResponse200 struct {
	// If the response is successful or not
	Success *bool `json:"success,omitempty"`
	// The array of persons
	Data []PersonItem `json:"data,omitempty"`
	AdditionalData *FieldsResponse200AllOfAdditionalData `json:"additional_data,omitempty"`
	RelatedObjects *ListPersonsResponse200AllOfRelatedObjects `json:"related_objects,omitempty"`
}

// NewListPersonsResponse200 instantiates a new ListPersonsResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPersonsResponse200() *ListPersonsResponse200 {
	this := ListPersonsResponse200{}
	return &this
}

// NewListPersonsResponse200WithDefaults instantiates a new ListPersonsResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPersonsResponse200WithDefaults() *ListPersonsResponse200 {
	this := ListPersonsResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ListPersonsResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPersonsResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ListPersonsResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ListPersonsResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListPersonsResponse200) GetData() []PersonItem {
	if o == nil || IsNil(o.Data) {
		var ret []PersonItem
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPersonsResponse200) GetDataOk() ([]PersonItem, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListPersonsResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PersonItem and assigns it to the Data field.
func (o *ListPersonsResponse200) SetData(v []PersonItem) {
	o.Data = v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *ListPersonsResponse200) GetAdditionalData() FieldsResponse200AllOfAdditionalData {
	if o == nil || IsNil(o.AdditionalData) {
		var ret FieldsResponse200AllOfAdditionalData
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPersonsResponse200) GetAdditionalDataOk() (*FieldsResponse200AllOfAdditionalData, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *ListPersonsResponse200) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given FieldsResponse200AllOfAdditionalData and assigns it to the AdditionalData field.
func (o *ListPersonsResponse200) SetAdditionalData(v FieldsResponse200AllOfAdditionalData) {
	o.AdditionalData = &v
}

// GetRelatedObjects returns the RelatedObjects field value if set, zero value otherwise.
func (o *ListPersonsResponse200) GetRelatedObjects() ListPersonsResponse200AllOfRelatedObjects {
	if o == nil || IsNil(o.RelatedObjects) {
		var ret ListPersonsResponse200AllOfRelatedObjects
		return ret
	}
	return *o.RelatedObjects
}

// GetRelatedObjectsOk returns a tuple with the RelatedObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPersonsResponse200) GetRelatedObjectsOk() (*ListPersonsResponse200AllOfRelatedObjects, bool) {
	if o == nil || IsNil(o.RelatedObjects) {
		return nil, false
	}
	return o.RelatedObjects, true
}

// HasRelatedObjects returns a boolean if a field has been set.
func (o *ListPersonsResponse200) HasRelatedObjects() bool {
	if o != nil && !IsNil(o.RelatedObjects) {
		return true
	}

	return false
}

// SetRelatedObjects gets a reference to the given ListPersonsResponse200AllOfRelatedObjects and assigns it to the RelatedObjects field.
func (o *ListPersonsResponse200) SetRelatedObjects(v ListPersonsResponse200AllOfRelatedObjects) {
	o.RelatedObjects = &v
}

func (o ListPersonsResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPersonsResponse200) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.AdditionalData) {
		toSerialize["additional_data"] = o.AdditionalData
	}
	if !IsNil(o.RelatedObjects) {
		toSerialize["related_objects"] = o.RelatedObjects
	}
	return toSerialize, nil
}

type NullableListPersonsResponse200 struct {
	value *ListPersonsResponse200
	isSet bool
}

func (v NullableListPersonsResponse200) Get() *ListPersonsResponse200 {
	return v.value
}

func (v *NullableListPersonsResponse200) Set(val *ListPersonsResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableListPersonsResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableListPersonsResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPersonsResponse200(val *ListPersonsResponse200) *NullableListPersonsResponse200 {
	return &NullableListPersonsResponse200{value: val, isSet: true}
}

func (v NullableListPersonsResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPersonsResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


