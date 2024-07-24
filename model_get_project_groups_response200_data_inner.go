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

// checks if the GetProjectGroupsResponse200DataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProjectGroupsResponse200DataInner{}

// GetProjectGroupsResponse200DataInner struct for GetProjectGroupsResponse200DataInner
type GetProjectGroupsResponse200DataInner struct {
	// ID of the group
	Id *float32 `json:"id,omitempty"`
	// Name of the group
	Name *string `json:"name,omitempty"`
	// Order number of the group
	OrderNr *float32 `json:"order_nr,omitempty"`
}

// NewGetProjectGroupsResponse200DataInner instantiates a new GetProjectGroupsResponse200DataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProjectGroupsResponse200DataInner() *GetProjectGroupsResponse200DataInner {
	this := GetProjectGroupsResponse200DataInner{}
	return &this
}

// NewGetProjectGroupsResponse200DataInnerWithDefaults instantiates a new GetProjectGroupsResponse200DataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProjectGroupsResponse200DataInnerWithDefaults() *GetProjectGroupsResponse200DataInner {
	this := GetProjectGroupsResponse200DataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetProjectGroupsResponse200DataInner) GetId() float32 {
	if o == nil || IsNil(o.Id) {
		var ret float32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectGroupsResponse200DataInner) GetIdOk() (*float32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetProjectGroupsResponse200DataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given float32 and assigns it to the Id field.
func (o *GetProjectGroupsResponse200DataInner) SetId(v float32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetProjectGroupsResponse200DataInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectGroupsResponse200DataInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetProjectGroupsResponse200DataInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetProjectGroupsResponse200DataInner) SetName(v string) {
	o.Name = &v
}

// GetOrderNr returns the OrderNr field value if set, zero value otherwise.
func (o *GetProjectGroupsResponse200DataInner) GetOrderNr() float32 {
	if o == nil || IsNil(o.OrderNr) {
		var ret float32
		return ret
	}
	return *o.OrderNr
}

// GetOrderNrOk returns a tuple with the OrderNr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProjectGroupsResponse200DataInner) GetOrderNrOk() (*float32, bool) {
	if o == nil || IsNil(o.OrderNr) {
		return nil, false
	}
	return o.OrderNr, true
}

// HasOrderNr returns a boolean if a field has been set.
func (o *GetProjectGroupsResponse200DataInner) HasOrderNr() bool {
	if o != nil && !IsNil(o.OrderNr) {
		return true
	}

	return false
}

// SetOrderNr gets a reference to the given float32 and assigns it to the OrderNr field.
func (o *GetProjectGroupsResponse200DataInner) SetOrderNr(v float32) {
	o.OrderNr = &v
}

func (o GetProjectGroupsResponse200DataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProjectGroupsResponse200DataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OrderNr) {
		toSerialize["order_nr"] = o.OrderNr
	}
	return toSerialize, nil
}

type NullableGetProjectGroupsResponse200DataInner struct {
	value *GetProjectGroupsResponse200DataInner
	isSet bool
}

func (v NullableGetProjectGroupsResponse200DataInner) Get() *GetProjectGroupsResponse200DataInner {
	return v.value
}

func (v *NullableGetProjectGroupsResponse200DataInner) Set(val *GetProjectGroupsResponse200DataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProjectGroupsResponse200DataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProjectGroupsResponse200DataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProjectGroupsResponse200DataInner(val *GetProjectGroupsResponse200DataInner) *NullableGetProjectGroupsResponse200DataInner {
	return &NullableGetProjectGroupsResponse200DataInner{value: val, isSet: true}
}

func (v NullableGetProjectGroupsResponse200DataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProjectGroupsResponse200DataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


