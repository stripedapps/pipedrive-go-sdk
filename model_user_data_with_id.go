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

// checks if the UserDataWithId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDataWithId{}

// UserDataWithId struct for UserDataWithId
type UserDataWithId struct {
	// The ID of the user
	Id *int32 `json:"id,omitempty"`
	// The name of the user
	Name *string `json:"name,omitempty"`
	// The email of the user
	Email *string `json:"email,omitempty"`
	// Whether the user has picture or not. 0 = No picture, 1 = Has picture.
	HasPic *int32 `json:"has_pic,omitempty"`
	// The user picture hash
	PicHash NullableString `json:"pic_hash,omitempty"`
	// Whether the user is active or not
	ActiveFlag *bool `json:"active_flag,omitempty"`
}

// NewUserDataWithId instantiates a new UserDataWithId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDataWithId() *UserDataWithId {
	this := UserDataWithId{}
	return &this
}

// NewUserDataWithIdWithDefaults instantiates a new UserDataWithId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDataWithIdWithDefaults() *UserDataWithId {
	this := UserDataWithId{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserDataWithId) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataWithId) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserDataWithId) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UserDataWithId) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserDataWithId) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataWithId) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserDataWithId) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserDataWithId) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserDataWithId) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataWithId) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserDataWithId) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserDataWithId) SetEmail(v string) {
	o.Email = &v
}

// GetHasPic returns the HasPic field value if set, zero value otherwise.
func (o *UserDataWithId) GetHasPic() int32 {
	if o == nil || IsNil(o.HasPic) {
		var ret int32
		return ret
	}
	return *o.HasPic
}

// GetHasPicOk returns a tuple with the HasPic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataWithId) GetHasPicOk() (*int32, bool) {
	if o == nil || IsNil(o.HasPic) {
		return nil, false
	}
	return o.HasPic, true
}

// HasHasPic returns a boolean if a field has been set.
func (o *UserDataWithId) HasHasPic() bool {
	if o != nil && !IsNil(o.HasPic) {
		return true
	}

	return false
}

// SetHasPic gets a reference to the given int32 and assigns it to the HasPic field.
func (o *UserDataWithId) SetHasPic(v int32) {
	o.HasPic = &v
}

// GetPicHash returns the PicHash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserDataWithId) GetPicHash() string {
	if o == nil || IsNil(o.PicHash.Get()) {
		var ret string
		return ret
	}
	return *o.PicHash.Get()
}

// GetPicHashOk returns a tuple with the PicHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserDataWithId) GetPicHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PicHash.Get(), o.PicHash.IsSet()
}

// HasPicHash returns a boolean if a field has been set.
func (o *UserDataWithId) HasPicHash() bool {
	if o != nil && o.PicHash.IsSet() {
		return true
	}

	return false
}

// SetPicHash gets a reference to the given NullableString and assigns it to the PicHash field.
func (o *UserDataWithId) SetPicHash(v string) {
	o.PicHash.Set(&v)
}
// SetPicHashNil sets the value for PicHash to be an explicit nil
func (o *UserDataWithId) SetPicHashNil() {
	o.PicHash.Set(nil)
}

// UnsetPicHash ensures that no value is present for PicHash, not even an explicit nil
func (o *UserDataWithId) UnsetPicHash() {
	o.PicHash.Unset()
}

// GetActiveFlag returns the ActiveFlag field value if set, zero value otherwise.
func (o *UserDataWithId) GetActiveFlag() bool {
	if o == nil || IsNil(o.ActiveFlag) {
		var ret bool
		return ret
	}
	return *o.ActiveFlag
}

// GetActiveFlagOk returns a tuple with the ActiveFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataWithId) GetActiveFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.ActiveFlag) {
		return nil, false
	}
	return o.ActiveFlag, true
}

// HasActiveFlag returns a boolean if a field has been set.
func (o *UserDataWithId) HasActiveFlag() bool {
	if o != nil && !IsNil(o.ActiveFlag) {
		return true
	}

	return false
}

// SetActiveFlag gets a reference to the given bool and assigns it to the ActiveFlag field.
func (o *UserDataWithId) SetActiveFlag(v bool) {
	o.ActiveFlag = &v
}

func (o UserDataWithId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDataWithId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.HasPic) {
		toSerialize["has_pic"] = o.HasPic
	}
	if o.PicHash.IsSet() {
		toSerialize["pic_hash"] = o.PicHash.Get()
	}
	if !IsNil(o.ActiveFlag) {
		toSerialize["active_flag"] = o.ActiveFlag
	}
	return toSerialize, nil
}

type NullableUserDataWithId struct {
	value *UserDataWithId
	isSet bool
}

func (v NullableUserDataWithId) Get() *UserDataWithId {
	return v.value
}

func (v *NullableUserDataWithId) Set(val *UserDataWithId) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDataWithId) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDataWithId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDataWithId(val *UserDataWithId) *NullableUserDataWithId {
	return &NullableUserDataWithId{value: val, isSet: true}
}

func (v NullableUserDataWithId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDataWithId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


