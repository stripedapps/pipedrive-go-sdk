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

// checks if the ListPersonsResponse200AllOfRelatedObjects type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPersonsResponse200AllOfRelatedObjects{}

// ListPersonsResponse200AllOfRelatedObjects struct for ListPersonsResponse200AllOfRelatedObjects
type ListPersonsResponse200AllOfRelatedObjects struct {
	Organization *AddActivityResponse200RelatedObjectsOrganization `json:"organization,omitempty"`
	User *GetActivitiesResponse200RelatedObjectsUser `json:"user,omitempty"`
}

// NewListPersonsResponse200AllOfRelatedObjects instantiates a new ListPersonsResponse200AllOfRelatedObjects object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPersonsResponse200AllOfRelatedObjects() *ListPersonsResponse200AllOfRelatedObjects {
	this := ListPersonsResponse200AllOfRelatedObjects{}
	return &this
}

// NewListPersonsResponse200AllOfRelatedObjectsWithDefaults instantiates a new ListPersonsResponse200AllOfRelatedObjects object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPersonsResponse200AllOfRelatedObjectsWithDefaults() *ListPersonsResponse200AllOfRelatedObjects {
	this := ListPersonsResponse200AllOfRelatedObjects{}
	return &this
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ListPersonsResponse200AllOfRelatedObjects) GetOrganization() AddActivityResponse200RelatedObjectsOrganization {
	if o == nil || IsNil(o.Organization) {
		var ret AddActivityResponse200RelatedObjectsOrganization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPersonsResponse200AllOfRelatedObjects) GetOrganizationOk() (*AddActivityResponse200RelatedObjectsOrganization, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ListPersonsResponse200AllOfRelatedObjects) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given AddActivityResponse200RelatedObjectsOrganization and assigns it to the Organization field.
func (o *ListPersonsResponse200AllOfRelatedObjects) SetOrganization(v AddActivityResponse200RelatedObjectsOrganization) {
	o.Organization = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ListPersonsResponse200AllOfRelatedObjects) GetUser() GetActivitiesResponse200RelatedObjectsUser {
	if o == nil || IsNil(o.User) {
		var ret GetActivitiesResponse200RelatedObjectsUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPersonsResponse200AllOfRelatedObjects) GetUserOk() (*GetActivitiesResponse200RelatedObjectsUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ListPersonsResponse200AllOfRelatedObjects) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given GetActivitiesResponse200RelatedObjectsUser and assigns it to the User field.
func (o *ListPersonsResponse200AllOfRelatedObjects) SetUser(v GetActivitiesResponse200RelatedObjectsUser) {
	o.User = &v
}

func (o ListPersonsResponse200AllOfRelatedObjects) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPersonsResponse200AllOfRelatedObjects) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableListPersonsResponse200AllOfRelatedObjects struct {
	value *ListPersonsResponse200AllOfRelatedObjects
	isSet bool
}

func (v NullableListPersonsResponse200AllOfRelatedObjects) Get() *ListPersonsResponse200AllOfRelatedObjects {
	return v.value
}

func (v *NullableListPersonsResponse200AllOfRelatedObjects) Set(val *ListPersonsResponse200AllOfRelatedObjects) {
	v.value = val
	v.isSet = true
}

func (v NullableListPersonsResponse200AllOfRelatedObjects) IsSet() bool {
	return v.isSet
}

func (v *NullableListPersonsResponse200AllOfRelatedObjects) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPersonsResponse200AllOfRelatedObjects(val *ListPersonsResponse200AllOfRelatedObjects) *NullableListPersonsResponse200AllOfRelatedObjects {
	return &NullableListPersonsResponse200AllOfRelatedObjects{value: val, isSet: true}
}

func (v NullableListPersonsResponse200AllOfRelatedObjects) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPersonsResponse200AllOfRelatedObjects) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


