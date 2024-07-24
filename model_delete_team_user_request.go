/*
Pipedrive API v1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DeleteTeamUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteTeamUserRequest{}

// DeleteTeamUserRequest struct for DeleteTeamUserRequest
type DeleteTeamUserRequest struct {
	// The list of user IDs
	Users []int32 `json:"users"`
}

type _DeleteTeamUserRequest DeleteTeamUserRequest

// NewDeleteTeamUserRequest instantiates a new DeleteTeamUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteTeamUserRequest(users []int32) *DeleteTeamUserRequest {
	this := DeleteTeamUserRequest{}
	this.Users = users
	return &this
}

// NewDeleteTeamUserRequestWithDefaults instantiates a new DeleteTeamUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteTeamUserRequestWithDefaults() *DeleteTeamUserRequest {
	this := DeleteTeamUserRequest{}
	return &this
}

// GetUsers returns the Users field value
func (o *DeleteTeamUserRequest) GetUsers() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *DeleteTeamUserRequest) GetUsersOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *DeleteTeamUserRequest) SetUsers(v []int32) {
	o.Users = v
}

func (o DeleteTeamUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteTeamUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	return toSerialize, nil
}

func (o *DeleteTeamUserRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"users",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDeleteTeamUserRequest := _DeleteTeamUserRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteTeamUserRequest)

	if err != nil {
		return err
	}

	*o = DeleteTeamUserRequest(varDeleteTeamUserRequest)

	return err
}

type NullableDeleteTeamUserRequest struct {
	value *DeleteTeamUserRequest
	isSet bool
}

func (v NullableDeleteTeamUserRequest) Get() *DeleteTeamUserRequest {
	return v.value
}

func (v *NullableDeleteTeamUserRequest) Set(val *DeleteTeamUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteTeamUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteTeamUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteTeamUserRequest(val *DeleteTeamUserRequest) *NullableDeleteTeamUserRequest {
	return &NullableDeleteTeamUserRequest{value: val, isSet: true}
}

func (v NullableDeleteTeamUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteTeamUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


