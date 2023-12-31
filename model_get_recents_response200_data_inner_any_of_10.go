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

// checks if the GetRecentsResponse200DataInnerAnyOf10 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRecentsResponse200DataInnerAnyOf10{}

// GetRecentsResponse200DataInnerAnyOf10 struct for GetRecentsResponse200DataInnerAnyOf10
type GetRecentsResponse200DataInnerAnyOf10 struct {
	Item *string `json:"item,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Data *GetAssociatedDealsResponse200AllOfRelatedObjectsStage `json:"data,omitempty"`
}

// NewGetRecentsResponse200DataInnerAnyOf10 instantiates a new GetRecentsResponse200DataInnerAnyOf10 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRecentsResponse200DataInnerAnyOf10() *GetRecentsResponse200DataInnerAnyOf10 {
	this := GetRecentsResponse200DataInnerAnyOf10{}
	return &this
}

// NewGetRecentsResponse200DataInnerAnyOf10WithDefaults instantiates a new GetRecentsResponse200DataInnerAnyOf10 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRecentsResponse200DataInnerAnyOf10WithDefaults() *GetRecentsResponse200DataInnerAnyOf10 {
	this := GetRecentsResponse200DataInnerAnyOf10{}
	return &this
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *GetRecentsResponse200DataInnerAnyOf10) GetItem() string {
	if o == nil || IsNil(o.Item) {
		var ret string
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecentsResponse200DataInnerAnyOf10) GetItemOk() (*string, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *GetRecentsResponse200DataInnerAnyOf10) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given string and assigns it to the Item field.
func (o *GetRecentsResponse200DataInnerAnyOf10) SetItem(v string) {
	o.Item = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetRecentsResponse200DataInnerAnyOf10) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecentsResponse200DataInnerAnyOf10) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetRecentsResponse200DataInnerAnyOf10) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetRecentsResponse200DataInnerAnyOf10) SetId(v int32) {
	o.Id = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetRecentsResponse200DataInnerAnyOf10) GetData() GetAssociatedDealsResponse200AllOfRelatedObjectsStage {
	if o == nil || IsNil(o.Data) {
		var ret GetAssociatedDealsResponse200AllOfRelatedObjectsStage
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecentsResponse200DataInnerAnyOf10) GetDataOk() (*GetAssociatedDealsResponse200AllOfRelatedObjectsStage, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetRecentsResponse200DataInnerAnyOf10) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GetAssociatedDealsResponse200AllOfRelatedObjectsStage and assigns it to the Data field.
func (o *GetRecentsResponse200DataInnerAnyOf10) SetData(v GetAssociatedDealsResponse200AllOfRelatedObjectsStage) {
	o.Data = &v
}

func (o GetRecentsResponse200DataInnerAnyOf10) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRecentsResponse200DataInnerAnyOf10) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Item) {
		toSerialize["item"] = o.Item
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGetRecentsResponse200DataInnerAnyOf10 struct {
	value *GetRecentsResponse200DataInnerAnyOf10
	isSet bool
}

func (v NullableGetRecentsResponse200DataInnerAnyOf10) Get() *GetRecentsResponse200DataInnerAnyOf10 {
	return v.value
}

func (v *NullableGetRecentsResponse200DataInnerAnyOf10) Set(val *GetRecentsResponse200DataInnerAnyOf10) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRecentsResponse200DataInnerAnyOf10) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRecentsResponse200DataInnerAnyOf10) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRecentsResponse200DataInnerAnyOf10(val *GetRecentsResponse200DataInnerAnyOf10) *NullableGetRecentsResponse200DataInnerAnyOf10 {
	return &NullableGetRecentsResponse200DataInnerAnyOf10{value: val, isSet: true}
}

func (v NullableGetRecentsResponse200DataInnerAnyOf10) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRecentsResponse200DataInnerAnyOf10) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


