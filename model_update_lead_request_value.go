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

// checks if the UpdateLeadRequestValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateLeadRequestValue{}

// UpdateLeadRequestValue The potential value of the lead
type UpdateLeadRequestValue struct {
	Amount float32 `json:"amount"`
	Currency string `json:"currency"`
}

// NewUpdateLeadRequestValue instantiates a new UpdateLeadRequestValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateLeadRequestValue(amount float32, currency string) *UpdateLeadRequestValue {
	this := UpdateLeadRequestValue{}
	this.Amount = amount
	this.Currency = currency
	return &this
}

// NewUpdateLeadRequestValueWithDefaults instantiates a new UpdateLeadRequestValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateLeadRequestValueWithDefaults() *UpdateLeadRequestValue {
	this := UpdateLeadRequestValue{}
	return &this
}

// GetAmount returns the Amount field value
func (o *UpdateLeadRequestValue) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *UpdateLeadRequestValue) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *UpdateLeadRequestValue) SetAmount(v float32) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *UpdateLeadRequestValue) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *UpdateLeadRequestValue) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *UpdateLeadRequestValue) SetCurrency(v string) {
	o.Currency = v
}

func (o UpdateLeadRequestValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateLeadRequestValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["currency"] = o.Currency
	return toSerialize, nil
}

type NullableUpdateLeadRequestValue struct {
	value *UpdateLeadRequestValue
	isSet bool
}

func (v NullableUpdateLeadRequestValue) Get() *UpdateLeadRequestValue {
	return v.value
}

func (v *NullableUpdateLeadRequestValue) Set(val *UpdateLeadRequestValue) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateLeadRequestValue) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateLeadRequestValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateLeadRequestValue(val *UpdateLeadRequestValue) *NullableUpdateLeadRequestValue {
	return &NullableUpdateLeadRequestValue{value: val, isSet: true}
}

func (v NullableUpdateLeadRequestValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateLeadRequestValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


