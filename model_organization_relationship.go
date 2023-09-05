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

// checks if the OrganizationRelationship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationRelationship{}

// OrganizationRelationship struct for OrganizationRelationship
type OrganizationRelationship struct {
	// The ID of the base organization for the returned calculated values
	OrgId *int32 `json:"org_id,omitempty"`
	// The type of organization relationship
	Type *string `json:"type,omitempty"`
	// The owner of this relationship. If type is `parent`, then the owner is the parent and the linked organization is the daughter.
	RelOwnerOrgId *int32 `json:"rel_owner_org_id,omitempty"`
	// The linked organization in this relationship. If type is `parent`, then the linked organization is the daughter.
	RelLinkedOrgId *int32 `json:"rel_linked_org_id,omitempty"`
}

// NewOrganizationRelationship instantiates a new OrganizationRelationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationRelationship() *OrganizationRelationship {
	this := OrganizationRelationship{}
	return &this
}

// NewOrganizationRelationshipWithDefaults instantiates a new OrganizationRelationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationRelationshipWithDefaults() *OrganizationRelationship {
	this := OrganizationRelationship{}
	return &this
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *OrganizationRelationship) GetOrgId() int32 {
	if o == nil || IsNil(o.OrgId) {
		var ret int32
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationRelationship) GetOrgIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *OrganizationRelationship) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given int32 and assigns it to the OrgId field.
func (o *OrganizationRelationship) SetOrgId(v int32) {
	o.OrgId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrganizationRelationship) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationRelationship) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrganizationRelationship) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrganizationRelationship) SetType(v string) {
	o.Type = &v
}

// GetRelOwnerOrgId returns the RelOwnerOrgId field value if set, zero value otherwise.
func (o *OrganizationRelationship) GetRelOwnerOrgId() int32 {
	if o == nil || IsNil(o.RelOwnerOrgId) {
		var ret int32
		return ret
	}
	return *o.RelOwnerOrgId
}

// GetRelOwnerOrgIdOk returns a tuple with the RelOwnerOrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationRelationship) GetRelOwnerOrgIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RelOwnerOrgId) {
		return nil, false
	}
	return o.RelOwnerOrgId, true
}

// HasRelOwnerOrgId returns a boolean if a field has been set.
func (o *OrganizationRelationship) HasRelOwnerOrgId() bool {
	if o != nil && !IsNil(o.RelOwnerOrgId) {
		return true
	}

	return false
}

// SetRelOwnerOrgId gets a reference to the given int32 and assigns it to the RelOwnerOrgId field.
func (o *OrganizationRelationship) SetRelOwnerOrgId(v int32) {
	o.RelOwnerOrgId = &v
}

// GetRelLinkedOrgId returns the RelLinkedOrgId field value if set, zero value otherwise.
func (o *OrganizationRelationship) GetRelLinkedOrgId() int32 {
	if o == nil || IsNil(o.RelLinkedOrgId) {
		var ret int32
		return ret
	}
	return *o.RelLinkedOrgId
}

// GetRelLinkedOrgIdOk returns a tuple with the RelLinkedOrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationRelationship) GetRelLinkedOrgIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RelLinkedOrgId) {
		return nil, false
	}
	return o.RelLinkedOrgId, true
}

// HasRelLinkedOrgId returns a boolean if a field has been set.
func (o *OrganizationRelationship) HasRelLinkedOrgId() bool {
	if o != nil && !IsNil(o.RelLinkedOrgId) {
		return true
	}

	return false
}

// SetRelLinkedOrgId gets a reference to the given int32 and assigns it to the RelLinkedOrgId field.
func (o *OrganizationRelationship) SetRelLinkedOrgId(v int32) {
	o.RelLinkedOrgId = &v
}

func (o OrganizationRelationship) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationRelationship) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.RelOwnerOrgId) {
		toSerialize["rel_owner_org_id"] = o.RelOwnerOrgId
	}
	if !IsNil(o.RelLinkedOrgId) {
		toSerialize["rel_linked_org_id"] = o.RelLinkedOrgId
	}
	return toSerialize, nil
}

type NullableOrganizationRelationship struct {
	value *OrganizationRelationship
	isSet bool
}

func (v NullableOrganizationRelationship) Get() *OrganizationRelationship {
	return v.value
}

func (v *NullableOrganizationRelationship) Set(val *OrganizationRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationRelationship(val *OrganizationRelationship) *NullableOrganizationRelationship {
	return &NullableOrganizationRelationship{value: val, isSet: true}
}

func (v NullableOrganizationRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


