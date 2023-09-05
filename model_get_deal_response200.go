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

// checks if the GetDealResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDealResponse200{}

// GetDealResponse200 struct for GetDealResponse200
type GetDealResponse200 struct {
	// If the response is successful or not
	Success *bool `json:"success,omitempty"`
	Data *DealNonStrictWithDetails `json:"data,omitempty"`
	AdditionalData *GetDealResponse200AdditionalData `json:"additional_data,omitempty"`
	RelatedObjects *GetDealsResponse200RelatedObjects `json:"related_objects,omitempty"`
}

// NewGetDealResponse200 instantiates a new GetDealResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDealResponse200() *GetDealResponse200 {
	this := GetDealResponse200{}
	return &this
}

// NewGetDealResponse200WithDefaults instantiates a new GetDealResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDealResponse200WithDefaults() *GetDealResponse200 {
	this := GetDealResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *GetDealResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDealResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *GetDealResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *GetDealResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetDealResponse200) GetData() DealNonStrictWithDetails {
	if o == nil || IsNil(o.Data) {
		var ret DealNonStrictWithDetails
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDealResponse200) GetDataOk() (*DealNonStrictWithDetails, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetDealResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DealNonStrictWithDetails and assigns it to the Data field.
func (o *GetDealResponse200) SetData(v DealNonStrictWithDetails) {
	o.Data = &v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *GetDealResponse200) GetAdditionalData() GetDealResponse200AdditionalData {
	if o == nil || IsNil(o.AdditionalData) {
		var ret GetDealResponse200AdditionalData
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDealResponse200) GetAdditionalDataOk() (*GetDealResponse200AdditionalData, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *GetDealResponse200) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given GetDealResponse200AdditionalData and assigns it to the AdditionalData field.
func (o *GetDealResponse200) SetAdditionalData(v GetDealResponse200AdditionalData) {
	o.AdditionalData = &v
}

// GetRelatedObjects returns the RelatedObjects field value if set, zero value otherwise.
func (o *GetDealResponse200) GetRelatedObjects() GetDealsResponse200RelatedObjects {
	if o == nil || IsNil(o.RelatedObjects) {
		var ret GetDealsResponse200RelatedObjects
		return ret
	}
	return *o.RelatedObjects
}

// GetRelatedObjectsOk returns a tuple with the RelatedObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDealResponse200) GetRelatedObjectsOk() (*GetDealsResponse200RelatedObjects, bool) {
	if o == nil || IsNil(o.RelatedObjects) {
		return nil, false
	}
	return o.RelatedObjects, true
}

// HasRelatedObjects returns a boolean if a field has been set.
func (o *GetDealResponse200) HasRelatedObjects() bool {
	if o != nil && !IsNil(o.RelatedObjects) {
		return true
	}

	return false
}

// SetRelatedObjects gets a reference to the given GetDealsResponse200RelatedObjects and assigns it to the RelatedObjects field.
func (o *GetDealResponse200) SetRelatedObjects(v GetDealsResponse200RelatedObjects) {
	o.RelatedObjects = &v
}

func (o GetDealResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDealResponse200) ToMap() (map[string]interface{}, error) {
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

type NullableGetDealResponse200 struct {
	value *GetDealResponse200
	isSet bool
}

func (v NullableGetDealResponse200) Get() *GetDealResponse200 {
	return v.value
}

func (v *NullableGetDealResponse200) Set(val *GetDealResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDealResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDealResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDealResponse200(val *GetDealResponse200) *NullableGetDealResponse200 {
	return &NullableGetDealResponse200{value: val, isSet: true}
}

func (v NullableGetDealResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDealResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


