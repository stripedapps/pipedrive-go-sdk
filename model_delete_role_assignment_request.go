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

// checks if the DeleteRoleAssignmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteRoleAssignmentRequest{}

// DeleteRoleAssignmentRequest struct for DeleteRoleAssignmentRequest
type DeleteRoleAssignmentRequest struct {
	// The ID of the user
	UserId int32 `json:"user_id"`
}

type _DeleteRoleAssignmentRequest DeleteRoleAssignmentRequest

// NewDeleteRoleAssignmentRequest instantiates a new DeleteRoleAssignmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteRoleAssignmentRequest(userId int32) *DeleteRoleAssignmentRequest {
	this := DeleteRoleAssignmentRequest{}
	this.UserId = userId
	return &this
}

// NewDeleteRoleAssignmentRequestWithDefaults instantiates a new DeleteRoleAssignmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteRoleAssignmentRequestWithDefaults() *DeleteRoleAssignmentRequest {
	this := DeleteRoleAssignmentRequest{}
	return &this
}

// GetUserId returns the UserId field value
func (o *DeleteRoleAssignmentRequest) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *DeleteRoleAssignmentRequest) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *DeleteRoleAssignmentRequest) SetUserId(v int32) {
	o.UserId = v
}

func (o DeleteRoleAssignmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteRoleAssignmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_id"] = o.UserId
	return toSerialize, nil
}

func (o *DeleteRoleAssignmentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_id",
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

	varDeleteRoleAssignmentRequest := _DeleteRoleAssignmentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteRoleAssignmentRequest)

	if err != nil {
		return err
	}

	*o = DeleteRoleAssignmentRequest(varDeleteRoleAssignmentRequest)

	return err
}

type NullableDeleteRoleAssignmentRequest struct {
	value *DeleteRoleAssignmentRequest
	isSet bool
}

func (v NullableDeleteRoleAssignmentRequest) Get() *DeleteRoleAssignmentRequest {
	return v.value
}

func (v *NullableDeleteRoleAssignmentRequest) Set(val *DeleteRoleAssignmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteRoleAssignmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteRoleAssignmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteRoleAssignmentRequest(val *DeleteRoleAssignmentRequest) *NullableDeleteRoleAssignmentRequest {
	return &NullableDeleteRoleAssignmentRequest{value: val, isSet: true}
}

func (v NullableDeleteRoleAssignmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteRoleAssignmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


