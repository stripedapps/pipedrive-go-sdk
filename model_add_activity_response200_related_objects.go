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

// checks if the AddActivityResponse200RelatedObjects type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddActivityResponse200RelatedObjects{}

// AddActivityResponse200RelatedObjects struct for AddActivityResponse200RelatedObjects
type AddActivityResponse200RelatedObjects struct {
	User *GetActivitiesResponse200RelatedObjectsUser `json:"user,omitempty"`
	Deal *GetActivitiesResponse200RelatedObjectsDeal `json:"deal,omitempty"`
	Person *AddActivityResponse200RelatedObjectsPerson `json:"person,omitempty"`
	Organization *AddActivityResponse200RelatedObjectsOrganization `json:"organization,omitempty"`
}

// NewAddActivityResponse200RelatedObjects instantiates a new AddActivityResponse200RelatedObjects object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddActivityResponse200RelatedObjects() *AddActivityResponse200RelatedObjects {
	this := AddActivityResponse200RelatedObjects{}
	return &this
}

// NewAddActivityResponse200RelatedObjectsWithDefaults instantiates a new AddActivityResponse200RelatedObjects object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddActivityResponse200RelatedObjectsWithDefaults() *AddActivityResponse200RelatedObjects {
	this := AddActivityResponse200RelatedObjects{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AddActivityResponse200RelatedObjects) GetUser() GetActivitiesResponse200RelatedObjectsUser {
	if o == nil || IsNil(o.User) {
		var ret GetActivitiesResponse200RelatedObjectsUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddActivityResponse200RelatedObjects) GetUserOk() (*GetActivitiesResponse200RelatedObjectsUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AddActivityResponse200RelatedObjects) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given GetActivitiesResponse200RelatedObjectsUser and assigns it to the User field.
func (o *AddActivityResponse200RelatedObjects) SetUser(v GetActivitiesResponse200RelatedObjectsUser) {
	o.User = &v
}

// GetDeal returns the Deal field value if set, zero value otherwise.
func (o *AddActivityResponse200RelatedObjects) GetDeal() GetActivitiesResponse200RelatedObjectsDeal {
	if o == nil || IsNil(o.Deal) {
		var ret GetActivitiesResponse200RelatedObjectsDeal
		return ret
	}
	return *o.Deal
}

// GetDealOk returns a tuple with the Deal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddActivityResponse200RelatedObjects) GetDealOk() (*GetActivitiesResponse200RelatedObjectsDeal, bool) {
	if o == nil || IsNil(o.Deal) {
		return nil, false
	}
	return o.Deal, true
}

// HasDeal returns a boolean if a field has been set.
func (o *AddActivityResponse200RelatedObjects) HasDeal() bool {
	if o != nil && !IsNil(o.Deal) {
		return true
	}

	return false
}

// SetDeal gets a reference to the given GetActivitiesResponse200RelatedObjectsDeal and assigns it to the Deal field.
func (o *AddActivityResponse200RelatedObjects) SetDeal(v GetActivitiesResponse200RelatedObjectsDeal) {
	o.Deal = &v
}

// GetPerson returns the Person field value if set, zero value otherwise.
func (o *AddActivityResponse200RelatedObjects) GetPerson() AddActivityResponse200RelatedObjectsPerson {
	if o == nil || IsNil(o.Person) {
		var ret AddActivityResponse200RelatedObjectsPerson
		return ret
	}
	return *o.Person
}

// GetPersonOk returns a tuple with the Person field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddActivityResponse200RelatedObjects) GetPersonOk() (*AddActivityResponse200RelatedObjectsPerson, bool) {
	if o == nil || IsNil(o.Person) {
		return nil, false
	}
	return o.Person, true
}

// HasPerson returns a boolean if a field has been set.
func (o *AddActivityResponse200RelatedObjects) HasPerson() bool {
	if o != nil && !IsNil(o.Person) {
		return true
	}

	return false
}

// SetPerson gets a reference to the given AddActivityResponse200RelatedObjectsPerson and assigns it to the Person field.
func (o *AddActivityResponse200RelatedObjects) SetPerson(v AddActivityResponse200RelatedObjectsPerson) {
	o.Person = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *AddActivityResponse200RelatedObjects) GetOrganization() AddActivityResponse200RelatedObjectsOrganization {
	if o == nil || IsNil(o.Organization) {
		var ret AddActivityResponse200RelatedObjectsOrganization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddActivityResponse200RelatedObjects) GetOrganizationOk() (*AddActivityResponse200RelatedObjectsOrganization, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *AddActivityResponse200RelatedObjects) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given AddActivityResponse200RelatedObjectsOrganization and assigns it to the Organization field.
func (o *AddActivityResponse200RelatedObjects) SetOrganization(v AddActivityResponse200RelatedObjectsOrganization) {
	o.Organization = &v
}

func (o AddActivityResponse200RelatedObjects) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddActivityResponse200RelatedObjects) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Deal) {
		toSerialize["deal"] = o.Deal
	}
	if !IsNil(o.Person) {
		toSerialize["person"] = o.Person
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	return toSerialize, nil
}

type NullableAddActivityResponse200RelatedObjects struct {
	value *AddActivityResponse200RelatedObjects
	isSet bool
}

func (v NullableAddActivityResponse200RelatedObjects) Get() *AddActivityResponse200RelatedObjects {
	return v.value
}

func (v *NullableAddActivityResponse200RelatedObjects) Set(val *AddActivityResponse200RelatedObjects) {
	v.value = val
	v.isSet = true
}

func (v NullableAddActivityResponse200RelatedObjects) IsSet() bool {
	return v.isSet
}

func (v *NullableAddActivityResponse200RelatedObjects) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddActivityResponse200RelatedObjects(val *AddActivityResponse200RelatedObjects) *NullableAddActivityResponse200RelatedObjects {
	return &NullableAddActivityResponse200RelatedObjects{value: val, isSet: true}
}

func (v NullableAddActivityResponse200RelatedObjects) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddActivityResponse200RelatedObjects) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


