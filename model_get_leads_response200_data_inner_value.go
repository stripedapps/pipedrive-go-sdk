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

// checks if the GetLeadsResponse200DataInnerValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLeadsResponse200DataInnerValue{}

// GetLeadsResponse200DataInnerValue The potential value of the lead represented by a JSON object: `{ \"amount\": 200, \"currency\": \"EUR\" }`. Both amount and currency are required.
type GetLeadsResponse200DataInnerValue struct {
	Amount float32 `json:"amount"`
	Currency string `json:"currency"`
}

type _GetLeadsResponse200DataInnerValue GetLeadsResponse200DataInnerValue

// NewGetLeadsResponse200DataInnerValue instantiates a new GetLeadsResponse200DataInnerValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLeadsResponse200DataInnerValue(amount float32, currency string) *GetLeadsResponse200DataInnerValue {
	this := GetLeadsResponse200DataInnerValue{}
	this.Amount = amount
	this.Currency = currency
	return &this
}

// NewGetLeadsResponse200DataInnerValueWithDefaults instantiates a new GetLeadsResponse200DataInnerValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLeadsResponse200DataInnerValueWithDefaults() *GetLeadsResponse200DataInnerValue {
	this := GetLeadsResponse200DataInnerValue{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetLeadsResponse200DataInnerValue) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetLeadsResponse200DataInnerValue) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetLeadsResponse200DataInnerValue) SetAmount(v float32) {
	o.Amount = v
}

// GetCurrency returns the Currency field value
func (o *GetLeadsResponse200DataInnerValue) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *GetLeadsResponse200DataInnerValue) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *GetLeadsResponse200DataInnerValue) SetCurrency(v string) {
	o.Currency = v
}

func (o GetLeadsResponse200DataInnerValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLeadsResponse200DataInnerValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["currency"] = o.Currency
	return toSerialize, nil
}

func (o *GetLeadsResponse200DataInnerValue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"currency",
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

	varGetLeadsResponse200DataInnerValue := _GetLeadsResponse200DataInnerValue{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetLeadsResponse200DataInnerValue)

	if err != nil {
		return err
	}

	*o = GetLeadsResponse200DataInnerValue(varGetLeadsResponse200DataInnerValue)

	return err
}

type NullableGetLeadsResponse200DataInnerValue struct {
	value *GetLeadsResponse200DataInnerValue
	isSet bool
}

func (v NullableGetLeadsResponse200DataInnerValue) Get() *GetLeadsResponse200DataInnerValue {
	return v.value
}

func (v *NullableGetLeadsResponse200DataInnerValue) Set(val *GetLeadsResponse200DataInnerValue) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLeadsResponse200DataInnerValue) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLeadsResponse200DataInnerValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLeadsResponse200DataInnerValue(val *GetLeadsResponse200DataInnerValue) *NullableGetLeadsResponse200DataInnerValue {
	return &NullableGetLeadsResponse200DataInnerValue{value: val, isSet: true}
}

func (v NullableGetLeadsResponse200DataInnerValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLeadsResponse200DataInnerValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


