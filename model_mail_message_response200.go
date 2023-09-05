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

// checks if the MailMessageResponse200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MailMessageResponse200{}

// MailMessageResponse200 struct for MailMessageResponse200
type MailMessageResponse200 struct {
	// If the response is successful or not
	Success *bool `json:"success,omitempty"`
	// The email service specific status code and it is returned through the response body.
	StatusCode *int32 `json:"statusCode,omitempty"`
	// The status text of the response.
	StatusText *string `json:"statusText,omitempty"`
	// The service name of the response.
	Service *string `json:"service,omitempty"`
	Data *MailMessageResponse200AllOfData `json:"data,omitempty"`
}

// NewMailMessageResponse200 instantiates a new MailMessageResponse200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMailMessageResponse200() *MailMessageResponse200 {
	this := MailMessageResponse200{}
	return &this
}

// NewMailMessageResponse200WithDefaults instantiates a new MailMessageResponse200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMailMessageResponse200WithDefaults() *MailMessageResponse200 {
	this := MailMessageResponse200{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *MailMessageResponse200) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailMessageResponse200) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *MailMessageResponse200) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *MailMessageResponse200) SetSuccess(v bool) {
	o.Success = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *MailMessageResponse200) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailMessageResponse200) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *MailMessageResponse200) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *MailMessageResponse200) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetStatusText returns the StatusText field value if set, zero value otherwise.
func (o *MailMessageResponse200) GetStatusText() string {
	if o == nil || IsNil(o.StatusText) {
		var ret string
		return ret
	}
	return *o.StatusText
}

// GetStatusTextOk returns a tuple with the StatusText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailMessageResponse200) GetStatusTextOk() (*string, bool) {
	if o == nil || IsNil(o.StatusText) {
		return nil, false
	}
	return o.StatusText, true
}

// HasStatusText returns a boolean if a field has been set.
func (o *MailMessageResponse200) HasStatusText() bool {
	if o != nil && !IsNil(o.StatusText) {
		return true
	}

	return false
}

// SetStatusText gets a reference to the given string and assigns it to the StatusText field.
func (o *MailMessageResponse200) SetStatusText(v string) {
	o.StatusText = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *MailMessageResponse200) GetService() string {
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailMessageResponse200) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *MailMessageResponse200) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *MailMessageResponse200) SetService(v string) {
	o.Service = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *MailMessageResponse200) GetData() MailMessageResponse200AllOfData {
	if o == nil || IsNil(o.Data) {
		var ret MailMessageResponse200AllOfData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailMessageResponse200) GetDataOk() (*MailMessageResponse200AllOfData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *MailMessageResponse200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given MailMessageResponse200AllOfData and assigns it to the Data field.
func (o *MailMessageResponse200) SetData(v MailMessageResponse200AllOfData) {
	o.Data = &v
}

func (o MailMessageResponse200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MailMessageResponse200) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.StatusCode) {
		toSerialize["statusCode"] = o.StatusCode
	}
	if !IsNil(o.StatusText) {
		toSerialize["statusText"] = o.StatusText
	}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableMailMessageResponse200 struct {
	value *MailMessageResponse200
	isSet bool
}

func (v NullableMailMessageResponse200) Get() *MailMessageResponse200 {
	return v.value
}

func (v *NullableMailMessageResponse200) Set(val *MailMessageResponse200) {
	v.value = val
	v.isSet = true
}

func (v NullableMailMessageResponse200) IsSet() bool {
	return v.isSet
}

func (v *NullableMailMessageResponse200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMailMessageResponse200(val *MailMessageResponse200) *NullableMailMessageResponse200 {
	return &NullableMailMessageResponse200{value: val, isSet: true}
}

func (v NullableMailMessageResponse200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMailMessageResponse200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


