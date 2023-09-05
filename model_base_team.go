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

// checks if the BaseTeam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseTeam{}

// BaseTeam struct for BaseTeam
type BaseTeam struct {
	// The team ID
	Id *int32 `json:"id,omitempty"`
	// The team name
	Name *string `json:"name,omitempty"`
	// The team description
	Description *string `json:"description,omitempty"`
	// The team manager ID
	ManagerId *int32 `json:"manager_id,omitempty"`
	// The list of user IDs
	Users []int32 `json:"users,omitempty"`
	ActiveFlag *float32 `json:"active_flag,omitempty"`
	DeletedFlag *float32 `json:"deleted_flag,omitempty"`
	// The team creation time. Format: YYYY-MM-DD HH:MM:SS
	AddTime *string `json:"add_time,omitempty"`
	// The ID of the user who created the team
	CreatedByUserId *int32 `json:"created_by_user_id,omitempty"`
}

// NewBaseTeam instantiates a new BaseTeam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseTeam() *BaseTeam {
	this := BaseTeam{}
	return &this
}

// NewBaseTeamWithDefaults instantiates a new BaseTeam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseTeamWithDefaults() *BaseTeam {
	this := BaseTeam{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseTeam) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTeam) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseTeam) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BaseTeam) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BaseTeam) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTeam) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BaseTeam) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BaseTeam) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BaseTeam) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTeam) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BaseTeam) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BaseTeam) SetDescription(v string) {
	o.Description = &v
}

// GetManagerId returns the ManagerId field value if set, zero value otherwise.
func (o *BaseTeam) GetManagerId() int32 {
	if o == nil || IsNil(o.ManagerId) {
		var ret int32
		return ret
	}
	return *o.ManagerId
}

// GetManagerIdOk returns a tuple with the ManagerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTeam) GetManagerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ManagerId) {
		return nil, false
	}
	return o.ManagerId, true
}

// HasManagerId returns a boolean if a field has been set.
func (o *BaseTeam) HasManagerId() bool {
	if o != nil && !IsNil(o.ManagerId) {
		return true
	}

	return false
}

// SetManagerId gets a reference to the given int32 and assigns it to the ManagerId field.
func (o *BaseTeam) SetManagerId(v int32) {
	o.ManagerId = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *BaseTeam) GetUsers() []int32 {
	if o == nil || IsNil(o.Users) {
		var ret []int32
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTeam) GetUsersOk() ([]int32, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *BaseTeam) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []int32 and assigns it to the Users field.
func (o *BaseTeam) SetUsers(v []int32) {
	o.Users = v
}

// GetActiveFlag returns the ActiveFlag field value if set, zero value otherwise.
func (o *BaseTeam) GetActiveFlag() float32 {
	if o == nil || IsNil(o.ActiveFlag) {
		var ret float32
		return ret
	}
	return *o.ActiveFlag
}

// GetActiveFlagOk returns a tuple with the ActiveFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTeam) GetActiveFlagOk() (*float32, bool) {
	if o == nil || IsNil(o.ActiveFlag) {
		return nil, false
	}
	return o.ActiveFlag, true
}

// HasActiveFlag returns a boolean if a field has been set.
func (o *BaseTeam) HasActiveFlag() bool {
	if o != nil && !IsNil(o.ActiveFlag) {
		return true
	}

	return false
}

// SetActiveFlag gets a reference to the given float32 and assigns it to the ActiveFlag field.
func (o *BaseTeam) SetActiveFlag(v float32) {
	o.ActiveFlag = &v
}

// GetDeletedFlag returns the DeletedFlag field value if set, zero value otherwise.
func (o *BaseTeam) GetDeletedFlag() float32 {
	if o == nil || IsNil(o.DeletedFlag) {
		var ret float32
		return ret
	}
	return *o.DeletedFlag
}

// GetDeletedFlagOk returns a tuple with the DeletedFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTeam) GetDeletedFlagOk() (*float32, bool) {
	if o == nil || IsNil(o.DeletedFlag) {
		return nil, false
	}
	return o.DeletedFlag, true
}

// HasDeletedFlag returns a boolean if a field has been set.
func (o *BaseTeam) HasDeletedFlag() bool {
	if o != nil && !IsNil(o.DeletedFlag) {
		return true
	}

	return false
}

// SetDeletedFlag gets a reference to the given float32 and assigns it to the DeletedFlag field.
func (o *BaseTeam) SetDeletedFlag(v float32) {
	o.DeletedFlag = &v
}

// GetAddTime returns the AddTime field value if set, zero value otherwise.
func (o *BaseTeam) GetAddTime() string {
	if o == nil || IsNil(o.AddTime) {
		var ret string
		return ret
	}
	return *o.AddTime
}

// GetAddTimeOk returns a tuple with the AddTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTeam) GetAddTimeOk() (*string, bool) {
	if o == nil || IsNil(o.AddTime) {
		return nil, false
	}
	return o.AddTime, true
}

// HasAddTime returns a boolean if a field has been set.
func (o *BaseTeam) HasAddTime() bool {
	if o != nil && !IsNil(o.AddTime) {
		return true
	}

	return false
}

// SetAddTime gets a reference to the given string and assigns it to the AddTime field.
func (o *BaseTeam) SetAddTime(v string) {
	o.AddTime = &v
}

// GetCreatedByUserId returns the CreatedByUserId field value if set, zero value otherwise.
func (o *BaseTeam) GetCreatedByUserId() int32 {
	if o == nil || IsNil(o.CreatedByUserId) {
		var ret int32
		return ret
	}
	return *o.CreatedByUserId
}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseTeam) GetCreatedByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedByUserId) {
		return nil, false
	}
	return o.CreatedByUserId, true
}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *BaseTeam) HasCreatedByUserId() bool {
	if o != nil && !IsNil(o.CreatedByUserId) {
		return true
	}

	return false
}

// SetCreatedByUserId gets a reference to the given int32 and assigns it to the CreatedByUserId field.
func (o *BaseTeam) SetCreatedByUserId(v int32) {
	o.CreatedByUserId = &v
}

func (o BaseTeam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseTeam) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.ManagerId) {
		toSerialize["manager_id"] = o.ManagerId
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !IsNil(o.ActiveFlag) {
		toSerialize["active_flag"] = o.ActiveFlag
	}
	if !IsNil(o.DeletedFlag) {
		toSerialize["deleted_flag"] = o.DeletedFlag
	}
	if !IsNil(o.AddTime) {
		toSerialize["add_time"] = o.AddTime
	}
	if !IsNil(o.CreatedByUserId) {
		toSerialize["created_by_user_id"] = o.CreatedByUserId
	}
	return toSerialize, nil
}

type NullableBaseTeam struct {
	value *BaseTeam
	isSet bool
}

func (v NullableBaseTeam) Get() *BaseTeam {
	return v.value
}

func (v *NullableBaseTeam) Set(val *BaseTeam) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseTeam) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseTeam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseTeam(val *BaseTeam) *NullableBaseTeam {
	return &NullableBaseTeam{value: val, isSet: true}
}

func (v NullableBaseTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseTeam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


