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

// checks if the UpdateFilterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateFilterRequest{}

// UpdateFilterRequest struct for UpdateFilterRequest
type UpdateFilterRequest struct {
	// The name of the filter
	Name *string `json:"name,omitempty"`
	// The conditions of the filter as a JSON object. Please note that a maximum of 16 conditions is allowed per filter and `date` values must be supplied in the `YYYY-MM-DD` format. It requires a minimum structure as follows: `{\"glue\":\"and\",\"conditions\":[{\"glue\":\"and\",\"conditions\": [CONDITION_OBJECTS]},{\"glue\":\"or\",\"conditions\":[CONDITION_OBJECTS]}]}`. Replace `CONDITION_OBJECTS` with JSON objects of the following structure: `{\"object\":\"\",\"field_id\":\"\", \"operator\":\"\",\"value\":\"\", \"extra_value\":\"\"}` or leave the array empty. Depending on the object type you should use another API endpoint to get `field_id`. There are five types of objects you can choose from: `\"person\"`, `\"deal\"`, `\"organization\"`, `\"product\"`, `\"activity\"` and you can use these types of operators depending on what type of a field you have: `\"IS NOT NULL\"`, `\"IS NULL\"`, `\"<=\"`, `\">=\"`, `\"<\"`, `\">\"`, `\"!=\"`, `\"=\"`, `\"LIKE '$%'\"`, `\"LIKE '%$%'\"`, `\"NOT LIKE '$%'\"`. To get a better understanding of how filters work try creating them directly from the Pipedrive application.
	Conditions map[string]interface{} `json:"conditions"`
}

type _UpdateFilterRequest UpdateFilterRequest

// NewUpdateFilterRequest instantiates a new UpdateFilterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateFilterRequest(conditions map[string]interface{}) *UpdateFilterRequest {
	this := UpdateFilterRequest{}
	this.Conditions = conditions
	return &this
}

// NewUpdateFilterRequestWithDefaults instantiates a new UpdateFilterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateFilterRequestWithDefaults() *UpdateFilterRequest {
	this := UpdateFilterRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateFilterRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateFilterRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateFilterRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateFilterRequest) SetName(v string) {
	o.Name = &v
}

// GetConditions returns the Conditions field value
func (o *UpdateFilterRequest) GetConditions() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *UpdateFilterRequest) GetConditionsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *UpdateFilterRequest) SetConditions(v map[string]interface{}) {
	o.Conditions = v
}

func (o UpdateFilterRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateFilterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["conditions"] = o.Conditions
	return toSerialize, nil
}

func (o *UpdateFilterRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"conditions",
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

	varUpdateFilterRequest := _UpdateFilterRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateFilterRequest)

	if err != nil {
		return err
	}

	*o = UpdateFilterRequest(varUpdateFilterRequest)

	return err
}

type NullableUpdateFilterRequest struct {
	value *UpdateFilterRequest
	isSet bool
}

func (v NullableUpdateFilterRequest) Get() *UpdateFilterRequest {
	return v.value
}

func (v *NullableUpdateFilterRequest) Set(val *UpdateFilterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateFilterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateFilterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateFilterRequest(val *UpdateFilterRequest) *NullableUpdateFilterRequest {
	return &NullableUpdateFilterRequest{value: val, isSet: true}
}

func (v NullableUpdateFilterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateFilterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


