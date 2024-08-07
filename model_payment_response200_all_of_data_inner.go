/*
Pipedrive API v1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the PaymentResponse200AllOfDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentResponse200AllOfDataInner{}

// PaymentResponse200AllOfDataInner struct for PaymentResponse200AllOfDataInner
type PaymentResponse200AllOfDataInner struct {
	// The ID of the payment
	Id *int32 `json:"id,omitempty"`
	// The ID of the subscription this payment is associated with
	SubscriptionId *int32 `json:"subscription_id,omitempty"`
	// The ID of the deal this payment is associated with
	DealId *int32 `json:"deal_id,omitempty"`
	// The payment status
	IsActive *bool `json:"is_active,omitempty"`
	// The payment amount
	Amount *float64 `json:"amount,omitempty"`
	// The currency of the payment
	Currency *string `json:"currency,omitempty"`
	// The difference between the amount of the current payment and the previous payment. The value can be either positive or negative.
	ChangeAmount *float64 `json:"change_amount,omitempty"`
	// The date when payment occurs
	DueAt *string `json:"due_at,omitempty"`
	// Represents the movement of revenue in comparison with the previous payment. Possible values are: `New` - first payment of the subscription. `Recurring` - no movement. `Expansion` - current payment amount > previous payment amount. `Contraction` - current payment amount < previous payment amount. `Churn` - last payment of the subscription.
	RevenueMovementType *string `json:"revenue_movement_type,omitempty"`
	// The type of the payment. Possible values are: `Recurring` - payments occur over fixed intervals of time, `Additional` - extra payment not the recurring payment of the recurring subscription, `Installment` - payment of the installment subscription.
	PaymentType *string `json:"payment_type,omitempty"`
	// The description of the payment
	Description *string `json:"description,omitempty"`
	// The creation time of the payment
	AddTime *time.Time `json:"add_time,omitempty"`
	// The update time of the payment
	UpdateTime *time.Time `json:"update_time,omitempty"`
}

// NewPaymentResponse200AllOfDataInner instantiates a new PaymentResponse200AllOfDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentResponse200AllOfDataInner() *PaymentResponse200AllOfDataInner {
	this := PaymentResponse200AllOfDataInner{}
	return &this
}

// NewPaymentResponse200AllOfDataInnerWithDefaults instantiates a new PaymentResponse200AllOfDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentResponse200AllOfDataInnerWithDefaults() *PaymentResponse200AllOfDataInner {
	this := PaymentResponse200AllOfDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PaymentResponse200AllOfDataInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponse200AllOfDataInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PaymentResponse200AllOfDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PaymentResponse200AllOfDataInner) SetId(v int32) {
	o.Id = &v
}

// GetSubscriptionId returns the SubscriptionId field value if set, zero value otherwise.
func (o *PaymentResponse200AllOfDataInner) GetSubscriptionId() int32 {
	if o == nil || IsNil(o.SubscriptionId) {
		var ret int32
		return ret
	}
	return *o.SubscriptionId
}

// GetSubscriptionIdOk returns a tuple with the SubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponse200AllOfDataInner) GetSubscriptionIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SubscriptionId) {
		return nil, false
	}
	return o.SubscriptionId, true
}

// HasSubscriptionId returns a boolean if a field has been set.
func (o *PaymentResponse200AllOfDataInner) HasSubscriptionId() bool {
	if o != nil && !IsNil(o.SubscriptionId) {
		return true
	}

	return false
}

// SetSubscriptionId gets a reference to the given int32 and assigns it to the SubscriptionId field.
func (o *PaymentResponse200AllOfDataInner) SetSubscriptionId(v int32) {
	o.SubscriptionId = &v
}

// GetDealId returns the DealId field value if set, zero value otherwise.
func (o *PaymentResponse200AllOfDataInner) GetDealId() int32 {
	if o == nil || IsNil(o.DealId) {
		var ret int32
		return ret
	}
	return *o.DealId
}

// GetDealIdOk returns a tuple with the DealId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponse200AllOfDataInner) GetDealIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DealId) {
		return nil, false
	}
	return o.DealId, true
}

// HasDealId returns a boolean if a field has been set.
func (o *PaymentResponse200AllOfDataInner) HasDealId() bool {
	if o != nil && !IsNil(o.DealId) {
		return true
	}

	return false
}

// SetDealId gets a reference to the given int32 and assigns it to the DealId field.
func (o *PaymentResponse200AllOfDataInner) SetDealId(v int32) {
	o.DealId = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *PaymentResponse200AllOfDataInner) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponse200AllOfDataInner) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *PaymentResponse200AllOfDataInner) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *PaymentResponse200AllOfDataInner) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PaymentResponse200AllOfDataInner) GetAmount() float64 {
	if o == nil || IsNil(o.Amount) {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponse200AllOfDataInner) GetAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PaymentResponse200AllOfDataInner) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *PaymentResponse200AllOfDataInner) SetAmount(v float64) {
	o.Amount = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *PaymentResponse200AllOfDataInner) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponse200AllOfDataInner) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *PaymentResponse200AllOfDataInner) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *PaymentResponse200AllOfDataInner) SetCurrency(v string) {
	o.Currency = &v
}

// GetChangeAmount returns the ChangeAmount field value if set, zero value otherwise.
func (o *PaymentResponse200AllOfDataInner) GetChangeAmount() float64 {
	if o == nil || IsNil(o.ChangeAmount) {
		var ret float64
		return ret
	}
	return *o.ChangeAmount
}

// GetChangeAmountOk returns a tuple with the ChangeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponse200AllOfDataInner) GetChangeAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.ChangeAmount) {
		return nil, false
	}
	return o.ChangeAmount, true
}

// HasChangeAmount returns a boolean if a field has been set.
func (o *PaymentResponse200AllOfDataInner) HasChangeAmount() bool {
	if o != nil && !IsNil(o.ChangeAmount) {
		return true
	}

	return false
}

// SetChangeAmount gets a reference to the given float64 and assigns it to the ChangeAmount field.
func (o *PaymentResponse200AllOfDataInner) SetChangeAmount(v float64) {
	o.ChangeAmount = &v
}

// GetDueAt returns the DueAt field value if set, zero value otherwise.
func (o *PaymentResponse200AllOfDataInner) GetDueAt() string {
	if o == nil || IsNil(o.DueAt) {
		var ret string
		return ret
	}
	return *o.DueAt
}

// GetDueAtOk returns a tuple with the DueAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponse200AllOfDataInner) GetDueAtOk() (*string, bool) {
	if o == nil || IsNil(o.DueAt) {
		return nil, false
	}
	return o.DueAt, true
}

// HasDueAt returns a boolean if a field has been set.
func (o *PaymentResponse200AllOfDataInner) HasDueAt() bool {
	if o != nil && !IsNil(o.DueAt) {
		return true
	}

	return false
}

// SetDueAt gets a reference to the given string and assigns it to the DueAt field.
func (o *PaymentResponse200AllOfDataInner) SetDueAt(v string) {
	o.DueAt = &v
}

// GetRevenueMovementType returns the RevenueMovementType field value if set, zero value otherwise.
func (o *PaymentResponse200AllOfDataInner) GetRevenueMovementType() string {
	if o == nil || IsNil(o.RevenueMovementType) {
		var ret string
		return ret
	}
	return *o.RevenueMovementType
}

// GetRevenueMovementTypeOk returns a tuple with the RevenueMovementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponse200AllOfDataInner) GetRevenueMovementTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RevenueMovementType) {
		return nil, false
	}
	return o.RevenueMovementType, true
}

// HasRevenueMovementType returns a boolean if a field has been set.
func (o *PaymentResponse200AllOfDataInner) HasRevenueMovementType() bool {
	if o != nil && !IsNil(o.RevenueMovementType) {
		return true
	}

	return false
}

// SetRevenueMovementType gets a reference to the given string and assigns it to the RevenueMovementType field.
func (o *PaymentResponse200AllOfDataInner) SetRevenueMovementType(v string) {
	o.RevenueMovementType = &v
}

// GetPaymentType returns the PaymentType field value if set, zero value otherwise.
func (o *PaymentResponse200AllOfDataInner) GetPaymentType() string {
	if o == nil || IsNil(o.PaymentType) {
		var ret string
		return ret
	}
	return *o.PaymentType
}

// GetPaymentTypeOk returns a tuple with the PaymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponse200AllOfDataInner) GetPaymentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentType) {
		return nil, false
	}
	return o.PaymentType, true
}

// HasPaymentType returns a boolean if a field has been set.
func (o *PaymentResponse200AllOfDataInner) HasPaymentType() bool {
	if o != nil && !IsNil(o.PaymentType) {
		return true
	}

	return false
}

// SetPaymentType gets a reference to the given string and assigns it to the PaymentType field.
func (o *PaymentResponse200AllOfDataInner) SetPaymentType(v string) {
	o.PaymentType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PaymentResponse200AllOfDataInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponse200AllOfDataInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PaymentResponse200AllOfDataInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PaymentResponse200AllOfDataInner) SetDescription(v string) {
	o.Description = &v
}

// GetAddTime returns the AddTime field value if set, zero value otherwise.
func (o *PaymentResponse200AllOfDataInner) GetAddTime() time.Time {
	if o == nil || IsNil(o.AddTime) {
		var ret time.Time
		return ret
	}
	return *o.AddTime
}

// GetAddTimeOk returns a tuple with the AddTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponse200AllOfDataInner) GetAddTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AddTime) {
		return nil, false
	}
	return o.AddTime, true
}

// HasAddTime returns a boolean if a field has been set.
func (o *PaymentResponse200AllOfDataInner) HasAddTime() bool {
	if o != nil && !IsNil(o.AddTime) {
		return true
	}

	return false
}

// SetAddTime gets a reference to the given time.Time and assigns it to the AddTime field.
func (o *PaymentResponse200AllOfDataInner) SetAddTime(v time.Time) {
	o.AddTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *PaymentResponse200AllOfDataInner) GetUpdateTime() time.Time {
	if o == nil || IsNil(o.UpdateTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentResponse200AllOfDataInner) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *PaymentResponse200AllOfDataInner) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given time.Time and assigns it to the UpdateTime field.
func (o *PaymentResponse200AllOfDataInner) SetUpdateTime(v time.Time) {
	o.UpdateTime = &v
}

func (o PaymentResponse200AllOfDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentResponse200AllOfDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SubscriptionId) {
		toSerialize["subscription_id"] = o.SubscriptionId
	}
	if !IsNil(o.DealId) {
		toSerialize["deal_id"] = o.DealId
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.ChangeAmount) {
		toSerialize["change_amount"] = o.ChangeAmount
	}
	if !IsNil(o.DueAt) {
		toSerialize["due_at"] = o.DueAt
	}
	if !IsNil(o.RevenueMovementType) {
		toSerialize["revenue_movement_type"] = o.RevenueMovementType
	}
	if !IsNil(o.PaymentType) {
		toSerialize["payment_type"] = o.PaymentType
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.AddTime) {
		toSerialize["add_time"] = o.AddTime
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["update_time"] = o.UpdateTime
	}
	return toSerialize, nil
}

type NullablePaymentResponse200AllOfDataInner struct {
	value *PaymentResponse200AllOfDataInner
	isSet bool
}

func (v NullablePaymentResponse200AllOfDataInner) Get() *PaymentResponse200AllOfDataInner {
	return v.value
}

func (v *NullablePaymentResponse200AllOfDataInner) Set(val *PaymentResponse200AllOfDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentResponse200AllOfDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentResponse200AllOfDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentResponse200AllOfDataInner(val *PaymentResponse200AllOfDataInner) *NullablePaymentResponse200AllOfDataInner {
	return &NullablePaymentResponse200AllOfDataInner{value: val, isSet: true}
}

func (v NullablePaymentResponse200AllOfDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentResponse200AllOfDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


