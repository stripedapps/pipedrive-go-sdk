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

// checks if the ReceiveMessage400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReceiveMessage400Response{}

// ReceiveMessage400Response struct for ReceiveMessage400Response
type ReceiveMessage400Response struct {
	Success *bool `json:"success,omitempty"`
	// The error description
	Error *string `json:"error,omitempty"`
	ErrorInfo *string `json:"error_info,omitempty"`
	AdditionalData *ReceiveMessage400ResponseAdditionalData `json:"additional_data,omitempty"`
}

// NewReceiveMessage400Response instantiates a new ReceiveMessage400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceiveMessage400Response() *ReceiveMessage400Response {
	this := ReceiveMessage400Response{}
	return &this
}

// NewReceiveMessage400ResponseWithDefaults instantiates a new ReceiveMessage400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceiveMessage400ResponseWithDefaults() *ReceiveMessage400Response {
	this := ReceiveMessage400Response{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ReceiveMessage400Response) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiveMessage400Response) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *ReceiveMessage400Response) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ReceiveMessage400Response) SetSuccess(v bool) {
	o.Success = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ReceiveMessage400Response) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiveMessage400Response) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ReceiveMessage400Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ReceiveMessage400Response) SetError(v string) {
	o.Error = &v
}

// GetErrorInfo returns the ErrorInfo field value if set, zero value otherwise.
func (o *ReceiveMessage400Response) GetErrorInfo() string {
	if o == nil || IsNil(o.ErrorInfo) {
		var ret string
		return ret
	}
	return *o.ErrorInfo
}

// GetErrorInfoOk returns a tuple with the ErrorInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiveMessage400Response) GetErrorInfoOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorInfo) {
		return nil, false
	}
	return o.ErrorInfo, true
}

// HasErrorInfo returns a boolean if a field has been set.
func (o *ReceiveMessage400Response) HasErrorInfo() bool {
	if o != nil && !IsNil(o.ErrorInfo) {
		return true
	}

	return false
}

// SetErrorInfo gets a reference to the given string and assigns it to the ErrorInfo field.
func (o *ReceiveMessage400Response) SetErrorInfo(v string) {
	o.ErrorInfo = &v
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *ReceiveMessage400Response) GetAdditionalData() ReceiveMessage400ResponseAdditionalData {
	if o == nil || IsNil(o.AdditionalData) {
		var ret ReceiveMessage400ResponseAdditionalData
		return ret
	}
	return *o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiveMessage400Response) GetAdditionalDataOk() (*ReceiveMessage400ResponseAdditionalData, bool) {
	if o == nil || IsNil(o.AdditionalData) {
		return nil, false
	}
	return o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *ReceiveMessage400Response) HasAdditionalData() bool {
	if o != nil && !IsNil(o.AdditionalData) {
		return true
	}

	return false
}

// SetAdditionalData gets a reference to the given ReceiveMessage400ResponseAdditionalData and assigns it to the AdditionalData field.
func (o *ReceiveMessage400Response) SetAdditionalData(v ReceiveMessage400ResponseAdditionalData) {
	o.AdditionalData = &v
}

func (o ReceiveMessage400Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReceiveMessage400Response) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AdditionalData) {
		toSerialize["additional_data"] = o.AdditionalData
	}
	return toSerialize, nil
}

type NullableReceiveMessage400Response struct {
	value *ReceiveMessage400Response
	isSet bool
}

func (v NullableReceiveMessage400Response) Get() *ReceiveMessage400Response {
	return v.value
}

func (v *NullableReceiveMessage400Response) Set(val *ReceiveMessage400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableReceiveMessage400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableReceiveMessage400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceiveMessage400Response(val *ReceiveMessage400Response) *NullableReceiveMessage400Response {
	return &NullableReceiveMessage400Response{value: val, isSet: true}
}

func (v NullableReceiveMessage400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceiveMessage400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


