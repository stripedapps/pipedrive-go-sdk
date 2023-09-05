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

// checks if the WebhooksBadRequestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhooksBadRequestResponse{}

// WebhooksBadRequestResponse struct for WebhooksBadRequestResponse
type WebhooksBadRequestResponse struct {
	// If the response is successful or not
	Success *bool `json:"success,omitempty"`
	// The status of the response
	Status *string `json:"status,omitempty"`
	// List of errors
	Errors map[string]interface{} `json:"errors,omitempty"`
}

// NewWebhooksBadRequestResponse instantiates a new WebhooksBadRequestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhooksBadRequestResponse() *WebhooksBadRequestResponse {
	this := WebhooksBadRequestResponse{}
	return &this
}

// NewWebhooksBadRequestResponseWithDefaults instantiates a new WebhooksBadRequestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhooksBadRequestResponseWithDefaults() *WebhooksBadRequestResponse {
	this := WebhooksBadRequestResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *WebhooksBadRequestResponse) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksBadRequestResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *WebhooksBadRequestResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *WebhooksBadRequestResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WebhooksBadRequestResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksBadRequestResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WebhooksBadRequestResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *WebhooksBadRequestResponse) SetStatus(v string) {
	o.Status = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *WebhooksBadRequestResponse) GetErrors() map[string]interface{} {
	if o == nil || IsNil(o.Errors) {
		var ret map[string]interface{}
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksBadRequestResponse) GetErrorsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Errors) {
		return map[string]interface{}{}, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *WebhooksBadRequestResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given map[string]interface{} and assigns it to the Errors field.
func (o *WebhooksBadRequestResponse) SetErrors(v map[string]interface{}) {
	o.Errors = v
}

func (o WebhooksBadRequestResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhooksBadRequestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableWebhooksBadRequestResponse struct {
	value *WebhooksBadRequestResponse
	isSet bool
}

func (v NullableWebhooksBadRequestResponse) Get() *WebhooksBadRequestResponse {
	return v.value
}

func (v *NullableWebhooksBadRequestResponse) Set(val *WebhooksBadRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhooksBadRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhooksBadRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhooksBadRequestResponse(val *WebhooksBadRequestResponse) *NullableWebhooksBadRequestResponse {
	return &NullableWebhooksBadRequestResponse{value: val, isSet: true}
}

func (v NullableWebhooksBadRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhooksBadRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


