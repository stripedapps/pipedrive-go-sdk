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

// checks if the GetAssociatedUpdatesResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAssociatedUpdatesResponse200{}

// GetAssociatedUpdatesResponse200 struct for GetAssociatedUpdatesResponse200
type GetAssociatedUpdatesResponse200 struct {
	// If the response is successful or not
	Success *bool `json:"success,omitempty"`
	Data []GetAssociatedUpdatesResponse200AllOfDataInner `json:"data,omitempty"`
	AdditionalData *FieldsResponse200AllOfAdditionalData `json:"additional_data,omitempty"`
	RelatedObjects *GetAssociatedUpdatesResponse200AllOfRelatedObjects `json:"related_objects,omitempty"`
}

// NewGetAssociatedUpdatesResponse200 instantiates a new GetAssociatedUpdatesResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAssociatedUpdatesResponse200() *GetAssociatedUpdatesResponse200 {
	this := GetAssociatedUpdatesResponse200{}
	return &this
}

// NewGetAssociatedUpdatesResponse200WithDefaults instantiates a new GetAssociatedUpdatesResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAssociatedUpdatesResponse200WithDefaults() *GetAssociatedUpdatesResponse200 {
	this := GetAssociatedUpdatesResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *GetAssociatedUpdatesResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssociatedUpdatesResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *GetAssociatedUpdatesResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *GetAssociatedUpdatesResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetAssociatedUpdatesResponse200) GetData() []GetAssociatedUpdatesResponse200AllOfDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetAssociatedUpdatesResponse200AllOfDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssociatedUpdatesResponse200) GetDataOk() ([]GetAssociatedUpdatesResponse200AllOfDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetAssociatedUpdatesResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetAssociatedUpdatesResponse200AllOfDataInner and assigns it to the Data field.
func (o *GetAssociatedUpdatesResponse200) SetData(v []GetAssociatedUpdatesResponse200AllOfDataInner) {
	o.Data = v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *GetAssociatedUpdatesResponse200) GetAdditionalData() FieldsResponse200AllOfAdditionalData {
	if o == nil || IsNil(o.AdditionalData) {
		var ret FieldsResponse200AllOfAdditionalData
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssociatedUpdatesResponse200) GetAdditionalDataOk() (*FieldsResponse200AllOfAdditionalData, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *GetAssociatedUpdatesResponse200) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given FieldsResponse200AllOfAdditionalData and assigns it to the AdditionalData field.
func (o *GetAssociatedUpdatesResponse200) SetAdditionalData(v FieldsResponse200AllOfAdditionalData) {
	o.AdditionalData = &v
}

// GetRelatedObjects returns the RelatedObjects field value if set, zero value otherwise.
func (o *GetAssociatedUpdatesResponse200) GetRelatedObjects() GetAssociatedUpdatesResponse200AllOfRelatedObjects {
	if o == nil || IsNil(o.RelatedObjects) {
		var ret GetAssociatedUpdatesResponse200AllOfRelatedObjects
		return ret
	}
	return *o.RelatedObjects
}

// GetRelatedObjectsOk returns a tuple with the RelatedObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAssociatedUpdatesResponse200) GetRelatedObjectsOk() (*GetAssociatedUpdatesResponse200AllOfRelatedObjects, bool) {
	if o == nil || IsNil(o.RelatedObjects) {
		return nil, false
	}
	return o.RelatedObjects, true
}

// HasRelatedObjects returns a boolean if a field has been set.
func (o *GetAssociatedUpdatesResponse200) HasRelatedObjects() bool {
	if o != nil && !IsNil(o.RelatedObjects) {
		return true
	}

	return false
}

// SetRelatedObjects gets a reference to the given GetAssociatedUpdatesResponse200AllOfRelatedObjects and assigns it to the RelatedObjects field.
func (o *GetAssociatedUpdatesResponse200) SetRelatedObjects(v GetAssociatedUpdatesResponse200AllOfRelatedObjects) {
	o.RelatedObjects = &v
}

func (o GetAssociatedUpdatesResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAssociatedUpdatesResponse200) ToMap() (map[string]interface{}, error) {
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

type NullableGetAssociatedUpdatesResponse200 struct {
	value *GetAssociatedUpdatesResponse200
	isSet bool
}

func (v NullableGetAssociatedUpdatesResponse200) Get() *GetAssociatedUpdatesResponse200 {
	return v.value
}

func (v *NullableGetAssociatedUpdatesResponse200) Set(val *GetAssociatedUpdatesResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAssociatedUpdatesResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAssociatedUpdatesResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAssociatedUpdatesResponse200(val *GetAssociatedUpdatesResponse200) *NullableGetAssociatedUpdatesResponse200 {
	return &NullableGetAssociatedUpdatesResponse200{value: val, isSet: true}
}

func (v NullableGetAssociatedUpdatesResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAssociatedUpdatesResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


