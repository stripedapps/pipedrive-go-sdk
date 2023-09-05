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

// checks if the GetActivitiesResponse200RelatedObjectsPersonPERSONID type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetActivitiesResponse200RelatedObjectsPersonPERSONID{}

// GetActivitiesResponse200RelatedObjectsPersonPERSONID The ID of the person associated with the item
type GetActivitiesResponse200RelatedObjectsPersonPERSONID struct {
	// The ID of the person associated with the item
	Id *int32 `json:"id,omitempty"`
	// The name of the person associated with the item
	Name *string `json:"name,omitempty"`
	// The emails of the person associated with the item
	Email []GetActivitiesResponse200RelatedObjectsPersonPERSONIDAllOfEmailInner `json:"email,omitempty"`
	// The phone numbers of the person associated with the item
	Phone []GetActivitiesResponse200RelatedObjectsPersonPERSONIDAllOfPhoneInner `json:"phone,omitempty"`
	// The ID of the owner of the person that is associated with the item
	OwnerId *int32 `json:"owner_id,omitempty"`
}

// NewGetActivitiesResponse200RelatedObjectsPersonPERSONID instantiates a new GetActivitiesResponse200RelatedObjectsPersonPERSONID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetActivitiesResponse200RelatedObjectsPersonPERSONID() *GetActivitiesResponse200RelatedObjectsPersonPERSONID {
	this := GetActivitiesResponse200RelatedObjectsPersonPERSONID{}
	return &this
}

// NewGetActivitiesResponse200RelatedObjectsPersonPERSONIDWithDefaults instantiates a new GetActivitiesResponse200RelatedObjectsPersonPERSONID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetActivitiesResponse200RelatedObjectsPersonPERSONIDWithDefaults() *GetActivitiesResponse200RelatedObjectsPersonPERSONID {
	this := GetActivitiesResponse200RelatedObjectsPersonPERSONID{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetActivitiesResponse200RelatedObjectsPersonPERSONID) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetActivitiesResponse200RelatedObjectsPersonPERSONID) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetActivitiesResponse200RelatedObjectsPersonPERSONID) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetActivitiesResponse200RelatedObjectsPersonPERSONID) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetActivitiesResponse200RelatedObjectsPersonPERSONID) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetActivitiesResponse200RelatedObjectsPersonPERSONID) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetActivitiesResponse200RelatedObjectsPersonPERSONID) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetActivitiesResponse200RelatedObjectsPersonPERSONID) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetActivitiesResponse200RelatedObjectsPersonPERSONID) GetEmail() []GetActivitiesResponse200RelatedObjectsPersonPERSONIDAllOfEmailInner {
	if o == nil || IsNil(o.Email) {
		var ret []GetActivitiesResponse200RelatedObjectsPersonPERSONIDAllOfEmailInner
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetActivitiesResponse200RelatedObjectsPersonPERSONID) GetEmailOk() ([]GetActivitiesResponse200RelatedObjectsPersonPERSONIDAllOfEmailInner, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetActivitiesResponse200RelatedObjectsPersonPERSONID) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given []GetActivitiesResponse200RelatedObjectsPersonPERSONIDAllOfEmailInner and assigns it to the Email field.
func (o *GetActivitiesResponse200RelatedObjectsPersonPERSONID) SetEmail(v []GetActivitiesResponse200RelatedObjectsPersonPERSONIDAllOfEmailInner) {
	o.Email = v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *GetActivitiesResponse200RelatedObjectsPersonPERSONID) GetPhone() []GetActivitiesResponse200RelatedObjectsPersonPERSONIDAllOfPhoneInner {
	if o == nil || IsNil(o.Phone) {
		var ret []GetActivitiesResponse200RelatedObjectsPersonPERSONIDAllOfPhoneInner
		return ret
	}
	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetActivitiesResponse200RelatedObjectsPersonPERSONID) GetPhoneOk() ([]GetActivitiesResponse200RelatedObjectsPersonPERSONIDAllOfPhoneInner, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *GetActivitiesResponse200RelatedObjectsPersonPERSONID) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given []GetActivitiesResponse200RelatedObjectsPersonPERSONIDAllOfPhoneInner and assigns it to the Phone field.
func (o *GetActivitiesResponse200RelatedObjectsPersonPERSONID) SetPhone(v []GetActivitiesResponse200RelatedObjectsPersonPERSONIDAllOfPhoneInner) {
	o.Phone = v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *GetActivitiesResponse200RelatedObjectsPersonPERSONID) GetOwnerId() int32 {
	if o == nil || IsNil(o.OwnerId) {
		var ret int32
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetActivitiesResponse200RelatedObjectsPersonPERSONID) GetOwnerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *GetActivitiesResponse200RelatedObjectsPersonPERSONID) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given int32 and assigns it to the OwnerId field.
func (o *GetActivitiesResponse200RelatedObjectsPersonPERSONID) SetOwnerId(v int32) {
	o.OwnerId = &v
}

func (o GetActivitiesResponse200RelatedObjectsPersonPERSONID) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetActivitiesResponse200RelatedObjectsPersonPERSONID) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	return toSerialize, nil
}

type NullableGetActivitiesResponse200RelatedObjectsPersonPERSONID struct {
	value *GetActivitiesResponse200RelatedObjectsPersonPERSONID
	isSet bool
}

func (v NullableGetActivitiesResponse200RelatedObjectsPersonPERSONID) Get() *GetActivitiesResponse200RelatedObjectsPersonPERSONID {
	return v.value
}

func (v *NullableGetActivitiesResponse200RelatedObjectsPersonPERSONID) Set(val *GetActivitiesResponse200RelatedObjectsPersonPERSONID) {
	v.value = val
	v.isSet = true
}

func (v NullableGetActivitiesResponse200RelatedObjectsPersonPERSONID) IsSet() bool {
	return v.isSet
}

func (v *NullableGetActivitiesResponse200RelatedObjectsPersonPERSONID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetActivitiesResponse200RelatedObjectsPersonPERSONID(val *GetActivitiesResponse200RelatedObjectsPersonPERSONID) *NullableGetActivitiesResponse200RelatedObjectsPersonPERSONID {
	return &NullableGetActivitiesResponse200RelatedObjectsPersonPERSONID{value: val, isSet: true}
}

func (v NullableGetActivitiesResponse200RelatedObjectsPersonPERSONID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetActivitiesResponse200RelatedObjectsPersonPERSONID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


