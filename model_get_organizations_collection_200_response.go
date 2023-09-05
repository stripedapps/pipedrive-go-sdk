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

// checks if the GetOrganizationsCollection200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOrganizationsCollection200Response{}

// GetOrganizationsCollection200Response struct for GetOrganizationsCollection200Response
type GetOrganizationsCollection200Response struct {
	Success *bool `json:"success,omitempty"`
	Data []OrganizationsCollectionResponseObject `json:"data,omitempty"`
	AdditionalData *GetActivitiesCollectionResponse200AdditionalData `json:"additional_data,omitempty"`
}

// NewGetOrganizationsCollection200Response instantiates a new GetOrganizationsCollection200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOrganizationsCollection200Response() *GetOrganizationsCollection200Response {
	this := GetOrganizationsCollection200Response{}
	return &this
}

// NewGetOrganizationsCollection200ResponseWithDefaults instantiates a new GetOrganizationsCollection200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOrganizationsCollection200ResponseWithDefaults() *GetOrganizationsCollection200Response {
	this := GetOrganizationsCollection200Response{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *GetOrganizationsCollection200Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationsCollection200Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *GetOrganizationsCollection200Response) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *GetOrganizationsCollection200Response) SetSuccess(v bool) {
	o.Success = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetOrganizationsCollection200Response) GetData() []OrganizationsCollectionResponseObject {
	if o == nil || IsNil(o.Data) {
		var ret []OrganizationsCollectionResponseObject
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationsCollection200Response) GetDataOk() ([]OrganizationsCollectionResponseObject, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetOrganizationsCollection200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []OrganizationsCollectionResponseObject and assigns it to the Data field.
func (o *GetOrganizationsCollection200Response) SetData(v []OrganizationsCollectionResponseObject) {
	o.Data = v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *GetOrganizationsCollection200Response) GetAdditionalData() GetActivitiesCollectionResponse200AdditionalData {
	if o == nil || IsNil(o.AdditionalData) {
		var ret GetActivitiesCollectionResponse200AdditionalData
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOrganizationsCollection200Response) GetAdditionalDataOk() (*GetActivitiesCollectionResponse200AdditionalData, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *GetOrganizationsCollection200Response) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given GetActivitiesCollectionResponse200AdditionalData and assigns it to the AdditionalData field.
func (o *GetOrganizationsCollection200Response) SetAdditionalData(v GetActivitiesCollectionResponse200AdditionalData) {
	o.AdditionalData = &v
}

func (o GetOrganizationsCollection200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetOrganizationsCollection200Response) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableGetOrganizationsCollection200Response struct {
	value *GetOrganizationsCollection200Response
	isSet bool
}

func (v NullableGetOrganizationsCollection200Response) Get() *GetOrganizationsCollection200Response {
	return v.value
}

func (v *NullableGetOrganizationsCollection200Response) Set(val *GetOrganizationsCollection200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOrganizationsCollection200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOrganizationsCollection200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOrganizationsCollection200Response(val *GetOrganizationsCollection200Response) *NullableGetOrganizationsCollection200Response {
	return &NullableGetOrganizationsCollection200Response{value: val, isSet: true}
}

func (v NullableGetOrganizationsCollection200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetOrganizationsCollection200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


