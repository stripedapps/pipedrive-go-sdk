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

// checks if the CancelRecurringSubscriptionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CancelRecurringSubscriptionRequest{}

// CancelRecurringSubscriptionRequest struct for CancelRecurringSubscriptionRequest
type CancelRecurringSubscriptionRequest struct {
	// The subscription termination date. All payments after specified date will be deleted. Default value is the current date.
	EndDate *string `json:"end_date,omitempty"`
}

// NewCancelRecurringSubscriptionRequest instantiates a new CancelRecurringSubscriptionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelRecurringSubscriptionRequest() *CancelRecurringSubscriptionRequest {
	this := CancelRecurringSubscriptionRequest{}
	return &this
}

// NewCancelRecurringSubscriptionRequestWithDefaults instantiates a new CancelRecurringSubscriptionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelRecurringSubscriptionRequestWithDefaults() *CancelRecurringSubscriptionRequest {
	this := CancelRecurringSubscriptionRequest{}
	return &this
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *CancelRecurringSubscriptionRequest) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelRecurringSubscriptionRequest) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CancelRecurringSubscriptionRequest) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *CancelRecurringSubscriptionRequest) SetEndDate(v string) {
	o.EndDate = &v
}

func (o CancelRecurringSubscriptionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelRecurringSubscriptionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	return toSerialize, nil
}

type NullableCancelRecurringSubscriptionRequest struct {
	value *CancelRecurringSubscriptionRequest
	isSet bool
}

func (v NullableCancelRecurringSubscriptionRequest) Get() *CancelRecurringSubscriptionRequest {
	return v.value
}

func (v *NullableCancelRecurringSubscriptionRequest) Set(val *CancelRecurringSubscriptionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelRecurringSubscriptionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelRecurringSubscriptionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelRecurringSubscriptionRequest(val *CancelRecurringSubscriptionRequest) *NullableCancelRecurringSubscriptionRequest {
	return &NullableCancelRecurringSubscriptionRequest{value: val, isSet: true}
}

func (v NullableCancelRecurringSubscriptionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelRecurringSubscriptionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


