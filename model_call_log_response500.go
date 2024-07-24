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

// checks if the CallLogResponse500 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallLogResponse500{}

// CallLogResponse500 struct for CallLogResponse500
type CallLogResponse500 struct {
	Success *bool `json:"success,omitempty"`
	// The description of the error
	Error *string `json:"error,omitempty"`
	// A message describing how to solve the problem
	ErrorInfo *string `json:"error_info,omitempty"`
	Data map[string]interface{} `json:"data,omitempty"`
	AdditionalData map[string]interface{} `json:"additional_data,omitempty"`
}

// NewCallLogResponse500 instantiates a new CallLogResponse500 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallLogResponse500() *CallLogResponse500 {
	this := CallLogResponse500{}
	return &this
}

// NewCallLogResponse500WithDefaults instantiates a new CallLogResponse500 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallLogResponse500WithDefaults() *CallLogResponse500 {
	this := CallLogResponse500{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *CallLogResponse500) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLogResponse500) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *CallLogResponse500) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *CallLogResponse500) SetSuccess(v bool) {
	o.Success = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *CallLogResponse500) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLogResponse500) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *CallLogResponse500) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *CallLogResponse500) SetError(v string) {
	o.Error = &v
}

// GetErrorInfo returns the ErrorInfo field value if set, zero value otherwise.
func (o *CallLogResponse500) GetErrorInfo() string {
	if o == nil || IsNil(o.ErrorInfo) {
		var ret string
		return ret
	}
	return *o.ErrorInfo
}

// GetErrorInfoOk returns a tuple with the ErrorInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLogResponse500) GetErrorInfoOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorInfo) {
		return nil, false
	}
	return o.ErrorInfo, true
}

// HasErrorInfo returns a boolean if a field has been set.
func (o *CallLogResponse500) HasErrorInfo() bool {
	if o != nil && !IsNil(o.ErrorInfo) {
		return true
	}

	return false
}

// SetErrorInfo gets a reference to the given string and assigns it to the ErrorInfo field.
func (o *CallLogResponse500) SetErrorInfo(v string) {
	o.ErrorInfo = &v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CallLogResponse500) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CallLogResponse500) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CallLogResponse500) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *CallLogResponse500) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CallLogResponse500) GetAdditionalData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CallLogResponse500) GetAdditionalDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *CallLogResponse500) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given map[string]interface{} and assigns it to the AdditionalData field.
func (o *CallLogResponse500) SetAdditionalData(v map[string]interface{}) {
	o.AdditionalData = v
}

func (o CallLogResponse500) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallLogResponse500) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.ErrorInfo) {
		toSerialize["error_info"] = o.ErrorInfo
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.AdditionalData != nil {
		toSerialize["additional_data"] = o.AdditionalData
	}
	return toSerialize, nil
}

type NullableCallLogResponse500 struct {
	value *CallLogResponse500
	isSet bool
}

func (v NullableCallLogResponse500) Get() *CallLogResponse500 {
	return v.value
}

func (v *NullableCallLogResponse500) Set(val *CallLogResponse500) {
	v.value = val
	v.isSet = true
}

func (v NullableCallLogResponse500) IsSet() bool {
	return v.isSet
}

func (v *NullableCallLogResponse500) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallLogResponse500(val *CallLogResponse500) *NullableCallLogResponse500 {
	return &NullableCallLogResponse500{value: val, isSet: true}
}

func (v NullableCallLogResponse500) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallLogResponse500) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


