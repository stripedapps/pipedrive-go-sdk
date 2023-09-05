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

// checks if the AddFilterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddFilterRequest{}

// AddFilterRequest struct for AddFilterRequest
type AddFilterRequest struct {
	// The name of the filter
	Name string `json:"name"`
	// The conditions of the filter as a JSON object. Please note that a maximum of 16 conditions is allowed per filter and `date` values must be supplied in the `YYYY-MM-DD` format. It requires a minimum structure as follows: `{\"glue\":\"and\",\"conditions\":[{\"glue\":\"and\",\"conditions\": [CONDITION_OBJECTS]},{\"glue\":\"or\",\"conditions\":[CONDITION_OBJECTS]}]}`. Replace `CONDITION_OBJECTS` with JSON objects of the following structure: `{\"object\":\"\",\"field_id\":\"\", \"operator\":\"\",\"value\":\"\", \"extra_value\":\"\"}` or leave the array empty. Depending on the object type you should use another API endpoint to get `field_id`. There are five types of objects you can choose from: `\"person\"`, `\"deal\"`, `\"organization\"`, `\"product\"`, `\"activity\"` and you can use these types of operators depending on what type of a field you have: `\"IS NOT NULL\"`, `\"IS NULL\"`, `\"<=\"`, `\">=\"`, `\"<\"`, `\">\"`, `\"!=\"`, `\"=\"`, `\"LIKE '$%'\"`, `\"LIKE '%$%'\"`, `\"NOT LIKE '$%'\"`. To get a better understanding of how filters work try creating them directly from the Pipedrive application.
	Conditions map[string]interface{} `json:"conditions"`
	Type string `json:"type"`
}

// NewAddFilterRequest instantiates a new AddFilterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddFilterRequest(name string, conditions map[string]interface{}, type_ string) *AddFilterRequest {
	this := AddFilterRequest{}
	this.Name = name
	this.Conditions = conditions
	this.Type = type_
	return &this
}

// NewAddFilterRequestWithDefaults instantiates a new AddFilterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddFilterRequestWithDefaults() *AddFilterRequest {
	this := AddFilterRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AddFilterRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddFilterRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddFilterRequest) SetName(v string) {
	o.Name = v
}

// GetConditions returns the Conditions field value
func (o *AddFilterRequest) GetConditions() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *AddFilterRequest) GetConditionsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *AddFilterRequest) SetConditions(v map[string]interface{}) {
	o.Conditions = v
}

// GetType returns the Type field value
func (o *AddFilterRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AddFilterRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AddFilterRequest) SetType(v string) {
	o.Type = v
}

func (o AddFilterRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddFilterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["conditions"] = o.Conditions
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAddFilterRequest struct {
	value *AddFilterRequest
	isSet bool
}

func (v NullableAddFilterRequest) Get() *AddFilterRequest {
	return v.value
}

func (v *NullableAddFilterRequest) Set(val *AddFilterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddFilterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddFilterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddFilterRequest(val *AddFilterRequest) *NullableAddFilterRequest {
	return &NullableAddFilterRequest{value: val, isSet: true}
}

func (v NullableAddFilterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddFilterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


