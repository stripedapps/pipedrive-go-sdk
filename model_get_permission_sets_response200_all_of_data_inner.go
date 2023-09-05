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

// checks if the GetPermissionSetsResponse200AllOfDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPermissionSetsResponse200AllOfDataInner{}

// GetPermissionSetsResponse200AllOfDataInner struct for GetPermissionSetsResponse200AllOfDataInner
type GetPermissionSetsResponse200AllOfDataInner struct {
	// The ID of user permission set
	Id *string `json:"id,omitempty"`
	// The name of the permission set
	Name *string `json:"name,omitempty"`
	// The description of the permission set
	Description *string `json:"description,omitempty"`
	// The app that permission set belongs to
	App *string `json:"app,omitempty"`
	// The type of permission set
	Type *string `json:"type,omitempty"`
	// The number of users assigned to this permission set
	AssignmentCount *int32 `json:"assignment_count,omitempty"`
}

// NewGetPermissionSetsResponse200AllOfDataInner instantiates a new GetPermissionSetsResponse200AllOfDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPermissionSetsResponse200AllOfDataInner() *GetPermissionSetsResponse200AllOfDataInner {
	this := GetPermissionSetsResponse200AllOfDataInner{}
	return &this
}

// NewGetPermissionSetsResponse200AllOfDataInnerWithDefaults instantiates a new GetPermissionSetsResponse200AllOfDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPermissionSetsResponse200AllOfDataInnerWithDefaults() *GetPermissionSetsResponse200AllOfDataInner {
	this := GetPermissionSetsResponse200AllOfDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetPermissionSetsResponse200AllOfDataInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPermissionSetsResponse200AllOfDataInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetPermissionSetsResponse200AllOfDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetPermissionSetsResponse200AllOfDataInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetPermissionSetsResponse200AllOfDataInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPermissionSetsResponse200AllOfDataInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetPermissionSetsResponse200AllOfDataInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetPermissionSetsResponse200AllOfDataInner) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetPermissionSetsResponse200AllOfDataInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPermissionSetsResponse200AllOfDataInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetPermissionSetsResponse200AllOfDataInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetPermissionSetsResponse200AllOfDataInner) SetDescription(v string) {
	o.Description = &v
}

// GetApp returns the App field value if set, zero value otherwise.
func (o *GetPermissionSetsResponse200AllOfDataInner) GetApp() string {
	if o == nil || IsNil(o.App) {
		var ret string
		return ret
	}
	return *o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPermissionSetsResponse200AllOfDataInner) GetAppOk() (*string, bool) {
	if o == nil || IsNil(o.App) {
		return nil, false
	}
	return o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *GetPermissionSetsResponse200AllOfDataInner) HasApp() bool {
	if o != nil && !IsNil(o.App) {
		return true
	}

	return false
}

// SetApp gets a reference to the given string and assigns it to the App field.
func (o *GetPermissionSetsResponse200AllOfDataInner) SetApp(v string) {
	o.App = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetPermissionSetsResponse200AllOfDataInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPermissionSetsResponse200AllOfDataInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetPermissionSetsResponse200AllOfDataInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetPermissionSetsResponse200AllOfDataInner) SetType(v string) {
	o.Type = &v
}

// GetAssignmentCount returns the AssignmentCount field value if set, zero value otherwise.
func (o *GetPermissionSetsResponse200AllOfDataInner) GetAssignmentCount() int32 {
	if o == nil || IsNil(o.AssignmentCount) {
		var ret int32
		return ret
	}
	return *o.AssignmentCount
}

// GetAssignmentCountOk returns a tuple with the AssignmentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPermissionSetsResponse200AllOfDataInner) GetAssignmentCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AssignmentCount) {
		return nil, false
	}
	return o.AssignmentCount, true
}

// HasAssignmentCount returns a boolean if a field has been set.
func (o *GetPermissionSetsResponse200AllOfDataInner) HasAssignmentCount() bool {
	if o != nil && !IsNil(o.AssignmentCount) {
		return true
	}

	return false
}

// SetAssignmentCount gets a reference to the given int32 and assigns it to the AssignmentCount field.
func (o *GetPermissionSetsResponse200AllOfDataInner) SetAssignmentCount(v int32) {
	o.AssignmentCount = &v
}

func (o GetPermissionSetsResponse200AllOfDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPermissionSetsResponse200AllOfDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.App) {
		toSerialize["app"] = o.App
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.AssignmentCount) {
		toSerialize["assignment_count"] = o.AssignmentCount
	}
	return toSerialize, nil
}

type NullableGetPermissionSetsResponse200AllOfDataInner struct {
	value *GetPermissionSetsResponse200AllOfDataInner
	isSet bool
}

func (v NullableGetPermissionSetsResponse200AllOfDataInner) Get() *GetPermissionSetsResponse200AllOfDataInner {
	return v.value
}

func (v *NullableGetPermissionSetsResponse200AllOfDataInner) Set(val *GetPermissionSetsResponse200AllOfDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPermissionSetsResponse200AllOfDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPermissionSetsResponse200AllOfDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPermissionSetsResponse200AllOfDataInner(val *GetPermissionSetsResponse200AllOfDataInner) *NullableGetPermissionSetsResponse200AllOfDataInner {
	return &NullableGetPermissionSetsResponse200AllOfDataInner{value: val, isSet: true}
}

func (v NullableGetPermissionSetsResponse200AllOfDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPermissionSetsResponse200AllOfDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


