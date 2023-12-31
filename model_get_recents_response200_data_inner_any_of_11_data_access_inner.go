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

// checks if the GetRecentsResponse200DataInnerAnyOf11DataAccessInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRecentsResponse200DataInnerAnyOf11DataAccessInner{}

// GetRecentsResponse200DataInnerAnyOf11DataAccessInner struct for GetRecentsResponse200DataInnerAnyOf11DataAccessInner
type GetRecentsResponse200DataInnerAnyOf11DataAccessInner struct {
	App *string `json:"app,omitempty"`
	Admin *bool `json:"admin,omitempty"`
	PermissionSetId *string `json:"permission_set_id,omitempty"`
}

// NewGetRecentsResponse200DataInnerAnyOf11DataAccessInner instantiates a new GetRecentsResponse200DataInnerAnyOf11DataAccessInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRecentsResponse200DataInnerAnyOf11DataAccessInner() *GetRecentsResponse200DataInnerAnyOf11DataAccessInner {
	this := GetRecentsResponse200DataInnerAnyOf11DataAccessInner{}
	return &this
}

// NewGetRecentsResponse200DataInnerAnyOf11DataAccessInnerWithDefaults instantiates a new GetRecentsResponse200DataInnerAnyOf11DataAccessInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRecentsResponse200DataInnerAnyOf11DataAccessInnerWithDefaults() *GetRecentsResponse200DataInnerAnyOf11DataAccessInner {
	this := GetRecentsResponse200DataInnerAnyOf11DataAccessInner{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *GetRecentsResponse200DataInnerAnyOf11DataAccessInner) GetApp() string {
	if o == nil || IsNil(o.App) {
		var ret string
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecentsResponse200DataInnerAnyOf11DataAccessInner) GetAppOk() (*string, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *GetRecentsResponse200DataInnerAnyOf11DataAccessInner) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given string and assigns it to the App field.
func (o *GetRecentsResponse200DataInnerAnyOf11DataAccessInner) SetApp(v string) {
	o.App = &v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *GetRecentsResponse200DataInnerAnyOf11DataAccessInner) GetAdmin() bool {
	if o == nil || IsNil(o.Admin) {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecentsResponse200DataInnerAnyOf11DataAccessInner) GetAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.Admin) {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *GetRecentsResponse200DataInnerAnyOf11DataAccessInner) HasAdmin() bool {
	if o != nil && !IsNil(o.Admin) {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *GetRecentsResponse200DataInnerAnyOf11DataAccessInner) SetAdmin(v bool) {
	o.Admin = &v
}

// GetPermissionSetId returns the PermissionSetId field value if set, zero value otherwise.
func (o *GetRecentsResponse200DataInnerAnyOf11DataAccessInner) GetPermissionSetId() string {
	if o == nil || IsNil(o.PermissionSetId) {
		var ret string
		return ret
	}
	return *o.PermissionSetId
}

// GetPermissionSetIdOk returns a tuple with the PermissionSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecentsResponse200DataInnerAnyOf11DataAccessInner) GetPermissionSetIdOk() (*string, bool) {
	if o == nil || IsNil(o.PermissionSetId) {
		return nil, false
	}
	return o.PermissionSetId, true
}

// HasPermissionSetId returns a boolean if a field has been set.
func (o *GetRecentsResponse200DataInnerAnyOf11DataAccessInner) HasPermissionSetId() bool {
	if o != nil && !IsNil(o.PermissionSetId) {
		return true
	}

	return false
}

// SetPermissionSetId gets a reference to the given string and assigns it to the PermissionSetId field.
func (o *GetRecentsResponse200DataInnerAnyOf11DataAccessInner) SetPermissionSetId(v string) {
	o.PermissionSetId = &v
}

func (o GetRecentsResponse200DataInnerAnyOf11DataAccessInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRecentsResponse200DataInnerAnyOf11DataAccessInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.Admin) {
		toSerialize["admin"] = o.Admin
	}
	if !IsNil(o.PermissionSetId) {
		toSerialize["permission_set_id"] = o.PermissionSetId
	}
	return toSerialize, nil
}

type NullableGetRecentsResponse200DataInnerAnyOf11DataAccessInner struct {
	value *GetRecentsResponse200DataInnerAnyOf11DataAccessInner
	isSet bool
}

func (v NullableGetRecentsResponse200DataInnerAnyOf11DataAccessInner) Get() *GetRecentsResponse200DataInnerAnyOf11DataAccessInner {
	return v.value
}

func (v *NullableGetRecentsResponse200DataInnerAnyOf11DataAccessInner) Set(val *GetRecentsResponse200DataInnerAnyOf11DataAccessInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRecentsResponse200DataInnerAnyOf11DataAccessInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRecentsResponse200DataInnerAnyOf11DataAccessInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRecentsResponse200DataInnerAnyOf11DataAccessInner(val *GetRecentsResponse200DataInnerAnyOf11DataAccessInner) *NullableGetRecentsResponse200DataInnerAnyOf11DataAccessInner {
	return &NullableGetRecentsResponse200DataInnerAnyOf11DataAccessInner{value: val, isSet: true}
}

func (v NullableGetRecentsResponse200DataInnerAnyOf11DataAccessInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRecentsResponse200DataInnerAnyOf11DataAccessInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


