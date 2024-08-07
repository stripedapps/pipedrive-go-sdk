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

// checks if the AddOrUpdateRoleSettingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddOrUpdateRoleSettingRequest{}

// AddOrUpdateRoleSettingRequest struct for AddOrUpdateRoleSettingRequest
type AddOrUpdateRoleSettingRequest struct {
	SettingKey string `json:"setting_key"`
	// Possible values for the `default_visibility` setting depending on the subscription plan:<br> <table class='role-setting'> <caption><b>Essential / Advanced plan</b></caption> <tr><th><b>Value</b></th><th><b>Description</b></th></tr> <tr><td>`1`</td><td>Owner & Followers</td></tr> <tr><td>`3`</td><td>Entire company</td></tr> </table> <br> <table class='role-setting'> <caption><b>Professional / Enterprise plan</b></caption> <tr><th><b>Value</b></th><th><b>Description</b></th></tr> <tr><td>`1`</td><td>Owner only</td></tr> <tr><td>`3`</td><td>Owner&#39;s visibility group</td></tr> <tr><td>`5`</td><td>Owner&#39;s visibility group and sub-groups</td></tr> <tr><td>`7`</td><td>Entire company</td></tr> </table> <br> Read more about visibility groups <a href='https://support.pipedrive.com/en/article/visibility-groups'>here</a>.
	Value int32 `json:"value"`
}

type _AddOrUpdateRoleSettingRequest AddOrUpdateRoleSettingRequest

// NewAddOrUpdateRoleSettingRequest instantiates a new AddOrUpdateRoleSettingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddOrUpdateRoleSettingRequest(settingKey string, value int32) *AddOrUpdateRoleSettingRequest {
	this := AddOrUpdateRoleSettingRequest{}
	this.SettingKey = settingKey
	this.Value = value
	return &this
}

// NewAddOrUpdateRoleSettingRequestWithDefaults instantiates a new AddOrUpdateRoleSettingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddOrUpdateRoleSettingRequestWithDefaults() *AddOrUpdateRoleSettingRequest {
	this := AddOrUpdateRoleSettingRequest{}
	return &this
}

// GetSettingKey returns the SettingKey field value
func (o *AddOrUpdateRoleSettingRequest) GetSettingKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SettingKey
}

// GetSettingKeyOk returns a tuple with the SettingKey field value
// and a boolean to check if the value has been set.
func (o *AddOrUpdateRoleSettingRequest) GetSettingKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SettingKey, true
}

// SetSettingKey sets field value
func (o *AddOrUpdateRoleSettingRequest) SetSettingKey(v string) {
	o.SettingKey = v
}

// GetValue returns the Value field value
func (o *AddOrUpdateRoleSettingRequest) GetValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AddOrUpdateRoleSettingRequest) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AddOrUpdateRoleSettingRequest) SetValue(v int32) {
	o.Value = v
}

func (o AddOrUpdateRoleSettingRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddOrUpdateRoleSettingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["setting_key"] = o.SettingKey
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *AddOrUpdateRoleSettingRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"setting_key",
		"value",
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

	varAddOrUpdateRoleSettingRequest := _AddOrUpdateRoleSettingRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddOrUpdateRoleSettingRequest)

	if err != nil {
		return err
	}

	*o = AddOrUpdateRoleSettingRequest(varAddOrUpdateRoleSettingRequest)

	return err
}

type NullableAddOrUpdateRoleSettingRequest struct {
	value *AddOrUpdateRoleSettingRequest
	isSet bool
}

func (v NullableAddOrUpdateRoleSettingRequest) Get() *AddOrUpdateRoleSettingRequest {
	return v.value
}

func (v *NullableAddOrUpdateRoleSettingRequest) Set(val *AddOrUpdateRoleSettingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddOrUpdateRoleSettingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddOrUpdateRoleSettingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddOrUpdateRoleSettingRequest(val *AddOrUpdateRoleSettingRequest) *NullableAddOrUpdateRoleSettingRequest {
	return &NullableAddOrUpdateRoleSettingRequest{value: val, isSet: true}
}

func (v NullableAddOrUpdateRoleSettingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddOrUpdateRoleSettingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


