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

// checks if the AddSubscriptionInstallmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddSubscriptionInstallmentRequest{}

// AddSubscriptionInstallmentRequest struct for AddSubscriptionInstallmentRequest
type AddSubscriptionInstallmentRequest struct {
	// The ID of the deal this installment subscription is associated with
	DealId int32 `json:"deal_id"`
	// The currency of the installment subscription. Accepts a 3-character currency code.
	Currency string `json:"currency"`
	// Array of payments. It requires a minimum structure as follows: [{ amount:SUM, description:DESCRIPTION, due_at:PAYMENT_DATE }]. Replace SUM with a payment amount, DESCRIPTION with an explanation string, PAYMENT_DATE with a date (format YYYY-MM-DD).
	Payments []map[string]interface{} `json:"payments"`
	// Indicates that the deal value must be set to the installment subscription's total value
	UpdateDealValue *bool `json:"update_deal_value,omitempty"`
}

type _AddSubscriptionInstallmentRequest AddSubscriptionInstallmentRequest

// NewAddSubscriptionInstallmentRequest instantiates a new AddSubscriptionInstallmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddSubscriptionInstallmentRequest(dealId int32, currency string, payments []map[string]interface{}) *AddSubscriptionInstallmentRequest {
	this := AddSubscriptionInstallmentRequest{}
	this.DealId = dealId
	this.Currency = currency
	this.Payments = payments
	return &this
}

// NewAddSubscriptionInstallmentRequestWithDefaults instantiates a new AddSubscriptionInstallmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddSubscriptionInstallmentRequestWithDefaults() *AddSubscriptionInstallmentRequest {
	this := AddSubscriptionInstallmentRequest{}
	return &this
}

// GetDealId returns the DealId field value
func (o *AddSubscriptionInstallmentRequest) GetDealId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DealId
}

// GetDealIdOk returns a tuple with the DealId field value
// and a boolean to check if the value has been set.
func (o *AddSubscriptionInstallmentRequest) GetDealIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DealId, true
}

// SetDealId sets field value
func (o *AddSubscriptionInstallmentRequest) SetDealId(v int32) {
	o.DealId = v
}

// GetCurrency returns the Currency field value
func (o *AddSubscriptionInstallmentRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *AddSubscriptionInstallmentRequest) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *AddSubscriptionInstallmentRequest) SetCurrency(v string) {
	o.Currency = v
}

// GetPayments returns the Payments field value
func (o *AddSubscriptionInstallmentRequest) GetPayments() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Payments
}

// GetPaymentsOk returns a tuple with the Payments field value
// and a boolean to check if the value has been set.
func (o *AddSubscriptionInstallmentRequest) GetPaymentsOk() ([]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.Payments, true
}

// SetPayments sets field value
func (o *AddSubscriptionInstallmentRequest) SetPayments(v []map[string]interface{}) {
	o.Payments = v
}

// GetUpdateDealValue returns the UpdateDealValue field value if set, zero value otherwise.
func (o *AddSubscriptionInstallmentRequest) GetUpdateDealValue() bool {
	if o == nil || IsNil(o.UpdateDealValue) {
		var ret bool
		return ret
	}
	return *o.UpdateDealValue
}

// GetUpdateDealValueOk returns a tuple with the UpdateDealValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddSubscriptionInstallmentRequest) GetUpdateDealValueOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateDealValue) {
		return nil, false
	}
	return o.UpdateDealValue, true
}

// HasUpdateDealValue returns a boolean if a field has been set.
func (o *AddSubscriptionInstallmentRequest) HasUpdateDealValue() bool {
	if o != nil && !IsNil(o.UpdateDealValue) {
		return true
	}

	return false
}

// SetUpdateDealValue gets a reference to the given bool and assigns it to the UpdateDealValue field.
func (o *AddSubscriptionInstallmentRequest) SetUpdateDealValue(v bool) {
	o.UpdateDealValue = &v
}

func (o AddSubscriptionInstallmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddSubscriptionInstallmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deal_id"] = o.DealId
	toSerialize["currency"] = o.Currency
	toSerialize["payments"] = o.Payments
	if !IsNil(o.UpdateDealValue) {
		toSerialize["update_deal_value"] = o.UpdateDealValue
	}
	return toSerialize, nil
}

func (o *AddSubscriptionInstallmentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deal_id",
		"currency",
		"payments",
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

	varAddSubscriptionInstallmentRequest := _AddSubscriptionInstallmentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddSubscriptionInstallmentRequest)

	if err != nil {
		return err
	}

	*o = AddSubscriptionInstallmentRequest(varAddSubscriptionInstallmentRequest)

	return err
}

type NullableAddSubscriptionInstallmentRequest struct {
	value *AddSubscriptionInstallmentRequest
	isSet bool
}

func (v NullableAddSubscriptionInstallmentRequest) Get() *AddSubscriptionInstallmentRequest {
	return v.value
}

func (v *NullableAddSubscriptionInstallmentRequest) Set(val *AddSubscriptionInstallmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddSubscriptionInstallmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddSubscriptionInstallmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddSubscriptionInstallmentRequest(val *AddSubscriptionInstallmentRequest) *NullableAddSubscriptionInstallmentRequest {
	return &NullableAddSubscriptionInstallmentRequest{value: val, isSet: true}
}

func (v NullableAddSubscriptionInstallmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddSubscriptionInstallmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


