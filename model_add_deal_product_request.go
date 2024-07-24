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

// checks if the AddDealProductRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddDealProductRequest{}

// AddDealProductRequest struct for AddDealProductRequest
type AddDealProductRequest struct {
	// The ID of the product to use
	ProductId int32 `json:"product_id"`
	// The price at which this product will be added to the deal
	ItemPrice float32 `json:"item_price"`
	// How many items of this product will be added to the deal
	Quantity int32 `json:"quantity"`
	// The value of the discount. The `discount_type` field can be used to specify whether the value is an amount or a percentage
	Discount *float32 `json:"discount,omitempty"`
	// The type of the discount's value
	DiscountType *string `json:"discount_type,omitempty"`
	// The ID of the product variation to use. When omitted, no variation will be used
	ProductVariationId *int32 `json:"product_variation_id,omitempty"`
	// A textual comment associated with this product-deal attachment
	Comments *string `json:"comments,omitempty"`
	// The tax percentage
	Tax *float32 `json:"tax,omitempty"`
	// The tax option to be applied to the products. When using `inclusive`, the tax percentage will already be included in the price. When using `exclusive`, the tax will not be included in the price. When using `none`, no tax will be added. Use the `tax` field for defining the tax percentage amount
	TaxMethod *string `json:"tax_method,omitempty"`
	// Whether the product is enabled for a deal or not. This makes it possible to add products to a deal with a specific price and discount criteria, but keep them disabled, which refrains them from being included in the deal value calculation. When omitted, the product will be marked as enabled by default
	EnabledFlag *bool `json:"enabled_flag,omitempty"`
	// Only available in Advanced and above plans  How often a customer is billed for access to a service or product  A deal can have up to 20 products attached with `billing_frequency` different than `one-time` 
	BillingFrequency *string `json:"billing_frequency,omitempty"`
	// Only available in Advanced and above plans  The number of times the billing frequency repeats for a product in a deal  When `billing_frequency` is set to `one-time`, this field must be `null`  For all the other values of `billing_frequency`, `null` represents a product billed indefinitely  Must be a positive integer less or equal to 312 
	BillingFrequencyCycles NullableInt32 `json:"billing_frequency_cycles,omitempty"`
	// Only available in Advanced and above plans  The billing start date. Must be between 15 years in the past and 15 years in the future 
	BillingStartDate NullableString `json:"billing_start_date,omitempty"`
}

type _AddDealProductRequest AddDealProductRequest

// NewAddDealProductRequest instantiates a new AddDealProductRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddDealProductRequest(productId int32, itemPrice float32, quantity int32) *AddDealProductRequest {
	this := AddDealProductRequest{}
	this.ProductId = productId
	this.ItemPrice = itemPrice
	this.Quantity = quantity
	var discount float32 = 0
	this.Discount = &discount
	var discountType string = "percentage"
	this.DiscountType = &discountType
	var tax float32 = 0
	this.Tax = &tax
	var enabledFlag bool = true
	this.EnabledFlag = &enabledFlag
	var billingFrequency string = "one-time"
	this.BillingFrequency = &billingFrequency
	var billingStartDate string = "null"
	this.BillingStartDate = *NewNullableString(&billingStartDate)
	return &this
}

// NewAddDealProductRequestWithDefaults instantiates a new AddDealProductRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddDealProductRequestWithDefaults() *AddDealProductRequest {
	this := AddDealProductRequest{}
	var discount float32 = 0
	this.Discount = &discount
	var discountType string = "percentage"
	this.DiscountType = &discountType
	var tax float32 = 0
	this.Tax = &tax
	var enabledFlag bool = true
	this.EnabledFlag = &enabledFlag
	var billingFrequency string = "one-time"
	this.BillingFrequency = &billingFrequency
	var billingStartDate string = "null"
	this.BillingStartDate = *NewNullableString(&billingStartDate)
	return &this
}

// GetProductId returns the ProductId field value
func (o *AddDealProductRequest) GetProductId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value
// and a boolean to check if the value has been set.
func (o *AddDealProductRequest) GetProductIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductId, true
}

// SetProductId sets field value
func (o *AddDealProductRequest) SetProductId(v int32) {
	o.ProductId = v
}

// GetItemPrice returns the ItemPrice field value
func (o *AddDealProductRequest) GetItemPrice() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ItemPrice
}

// GetItemPriceOk returns a tuple with the ItemPrice field value
// and a boolean to check if the value has been set.
func (o *AddDealProductRequest) GetItemPriceOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ItemPrice, true
}

// SetItemPrice sets field value
func (o *AddDealProductRequest) SetItemPrice(v float32) {
	o.ItemPrice = v
}

// GetQuantity returns the Quantity field value
func (o *AddDealProductRequest) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *AddDealProductRequest) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *AddDealProductRequest) SetQuantity(v int32) {
	o.Quantity = v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *AddDealProductRequest) GetDiscount() float32 {
	if o == nil || IsNil(o.Discount) {
		var ret float32
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDealProductRequest) GetDiscountOk() (*float32, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *AddDealProductRequest) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given float32 and assigns it to the Discount field.
func (o *AddDealProductRequest) SetDiscount(v float32) {
	o.Discount = &v
}

// GetDiscountType returns the DiscountType field value if set, zero value otherwise.
func (o *AddDealProductRequest) GetDiscountType() string {
	if o == nil || IsNil(o.DiscountType) {
		var ret string
		return ret
	}
	return *o.DiscountType
}

// GetDiscountTypeOk returns a tuple with the DiscountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDealProductRequest) GetDiscountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountType) {
		return nil, false
	}
	return o.DiscountType, true
}

// HasDiscountType returns a boolean if a field has been set.
func (o *AddDealProductRequest) HasDiscountType() bool {
	if o != nil && !IsNil(o.DiscountType) {
		return true
	}

	return false
}

// SetDiscountType gets a reference to the given string and assigns it to the DiscountType field.
func (o *AddDealProductRequest) SetDiscountType(v string) {
	o.DiscountType = &v
}

// GetProductVariationId returns the ProductVariationId field value if set, zero value otherwise.
func (o *AddDealProductRequest) GetProductVariationId() int32 {
	if o == nil || IsNil(o.ProductVariationId) {
		var ret int32
		return ret
	}
	return *o.ProductVariationId
}

// GetProductVariationIdOk returns a tuple with the ProductVariationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDealProductRequest) GetProductVariationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProductVariationId) {
		return nil, false
	}
	return o.ProductVariationId, true
}

// HasProductVariationId returns a boolean if a field has been set.
func (o *AddDealProductRequest) HasProductVariationId() bool {
	if o != nil && !IsNil(o.ProductVariationId) {
		return true
	}

	return false
}

// SetProductVariationId gets a reference to the given int32 and assigns it to the ProductVariationId field.
func (o *AddDealProductRequest) SetProductVariationId(v int32) {
	o.ProductVariationId = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *AddDealProductRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDealProductRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *AddDealProductRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *AddDealProductRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTax returns the Tax field value if set, zero value otherwise.
func (o *AddDealProductRequest) GetTax() float32 {
	if o == nil || IsNil(o.Tax) {
		var ret float32
		return ret
	}
	return *o.Tax
}

// GetTaxOk returns a tuple with the Tax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDealProductRequest) GetTaxOk() (*float32, bool) {
	if o == nil || IsNil(o.Tax) {
		return nil, false
	}
	return o.Tax, true
}

// HasTax returns a boolean if a field has been set.
func (o *AddDealProductRequest) HasTax() bool {
	if o != nil && !IsNil(o.Tax) {
		return true
	}

	return false
}

// SetTax gets a reference to the given float32 and assigns it to the Tax field.
func (o *AddDealProductRequest) SetTax(v float32) {
	o.Tax = &v
}

// GetTaxMethod returns the TaxMethod field value if set, zero value otherwise.
func (o *AddDealProductRequest) GetTaxMethod() string {
	if o == nil || IsNil(o.TaxMethod) {
		var ret string
		return ret
	}
	return *o.TaxMethod
}

// GetTaxMethodOk returns a tuple with the TaxMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDealProductRequest) GetTaxMethodOk() (*string, bool) {
	if o == nil || IsNil(o.TaxMethod) {
		return nil, false
	}
	return o.TaxMethod, true
}

// HasTaxMethod returns a boolean if a field has been set.
func (o *AddDealProductRequest) HasTaxMethod() bool {
	if o != nil && !IsNil(o.TaxMethod) {
		return true
	}

	return false
}

// SetTaxMethod gets a reference to the given string and assigns it to the TaxMethod field.
func (o *AddDealProductRequest) SetTaxMethod(v string) {
	o.TaxMethod = &v
}

// GetEnabledFlag returns the EnabledFlag field value if set, zero value otherwise.
func (o *AddDealProductRequest) GetEnabledFlag() bool {
	if o == nil || IsNil(o.EnabledFlag) {
		var ret bool
		return ret
	}
	return *o.EnabledFlag
}

// GetEnabledFlagOk returns a tuple with the EnabledFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDealProductRequest) GetEnabledFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.EnabledFlag) {
		return nil, false
	}
	return o.EnabledFlag, true
}

// HasEnabledFlag returns a boolean if a field has been set.
func (o *AddDealProductRequest) HasEnabledFlag() bool {
	if o != nil && !IsNil(o.EnabledFlag) {
		return true
	}

	return false
}

// SetEnabledFlag gets a reference to the given bool and assigns it to the EnabledFlag field.
func (o *AddDealProductRequest) SetEnabledFlag(v bool) {
	o.EnabledFlag = &v
}

// GetBillingFrequency returns the BillingFrequency field value if set, zero value otherwise.
func (o *AddDealProductRequest) GetBillingFrequency() string {
	if o == nil || IsNil(o.BillingFrequency) {
		var ret string
		return ret
	}
	return *o.BillingFrequency
}

// GetBillingFrequencyOk returns a tuple with the BillingFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDealProductRequest) GetBillingFrequencyOk() (*string, bool) {
	if o == nil || IsNil(o.BillingFrequency) {
		return nil, false
	}
	return o.BillingFrequency, true
}

// HasBillingFrequency returns a boolean if a field has been set.
func (o *AddDealProductRequest) HasBillingFrequency() bool {
	if o != nil && !IsNil(o.BillingFrequency) {
		return true
	}

	return false
}

// SetBillingFrequency gets a reference to the given string and assigns it to the BillingFrequency field.
func (o *AddDealProductRequest) SetBillingFrequency(v string) {
	o.BillingFrequency = &v
}

// GetBillingFrequencyCycles returns the BillingFrequencyCycles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddDealProductRequest) GetBillingFrequencyCycles() int32 {
	if o == nil || IsNil(o.BillingFrequencyCycles.Get()) {
		var ret int32
		return ret
	}
	return *o.BillingFrequencyCycles.Get()
}

// GetBillingFrequencyCyclesOk returns a tuple with the BillingFrequencyCycles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddDealProductRequest) GetBillingFrequencyCyclesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingFrequencyCycles.Get(), o.BillingFrequencyCycles.IsSet()
}

// HasBillingFrequencyCycles returns a boolean if a field has been set.
func (o *AddDealProductRequest) HasBillingFrequencyCycles() bool {
	if o != nil && o.BillingFrequencyCycles.IsSet() {
		return true
	}

	return false
}

// SetBillingFrequencyCycles gets a reference to the given NullableInt32 and assigns it to the BillingFrequencyCycles field.
func (o *AddDealProductRequest) SetBillingFrequencyCycles(v int32) {
	o.BillingFrequencyCycles.Set(&v)
}
// SetBillingFrequencyCyclesNil sets the value for BillingFrequencyCycles to be an explicit nil
func (o *AddDealProductRequest) SetBillingFrequencyCyclesNil() {
	o.BillingFrequencyCycles.Set(nil)
}

// UnsetBillingFrequencyCycles ensures that no value is present for BillingFrequencyCycles, not even an explicit nil
func (o *AddDealProductRequest) UnsetBillingFrequencyCycles() {
	o.BillingFrequencyCycles.Unset()
}

// GetBillingStartDate returns the BillingStartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AddDealProductRequest) GetBillingStartDate() string {
	if o == nil || IsNil(o.BillingStartDate.Get()) {
		var ret string
		return ret
	}
	return *o.BillingStartDate.Get()
}

// GetBillingStartDateOk returns a tuple with the BillingStartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AddDealProductRequest) GetBillingStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingStartDate.Get(), o.BillingStartDate.IsSet()
}

// HasBillingStartDate returns a boolean if a field has been set.
func (o *AddDealProductRequest) HasBillingStartDate() bool {
	if o != nil && o.BillingStartDate.IsSet() {
		return true
	}

	return false
}

// SetBillingStartDate gets a reference to the given NullableString and assigns it to the BillingStartDate field.
func (o *AddDealProductRequest) SetBillingStartDate(v string) {
	o.BillingStartDate.Set(&v)
}
// SetBillingStartDateNil sets the value for BillingStartDate to be an explicit nil
func (o *AddDealProductRequest) SetBillingStartDateNil() {
	o.BillingStartDate.Set(nil)
}

// UnsetBillingStartDate ensures that no value is present for BillingStartDate, not even an explicit nil
func (o *AddDealProductRequest) UnsetBillingStartDate() {
	o.BillingStartDate.Unset()
}

func (o AddDealProductRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddDealProductRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["product_id"] = o.ProductId
	toSerialize["item_price"] = o.ItemPrice
	toSerialize["quantity"] = o.Quantity
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !IsNil(o.DiscountType) {
		toSerialize["discount_type"] = o.DiscountType
	}
	if !IsNil(o.ProductVariationId) {
		toSerialize["product_variation_id"] = o.ProductVariationId
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Tax) {
		toSerialize["tax"] = o.Tax
	}
	if !IsNil(o.TaxMethod) {
		toSerialize["tax_method"] = o.TaxMethod
	}
	if !IsNil(o.EnabledFlag) {
		toSerialize["enabled_flag"] = o.EnabledFlag
	}
	if !IsNil(o.BillingFrequency) {
		toSerialize["billing_frequency"] = o.BillingFrequency
	}
	if o.BillingFrequencyCycles.IsSet() {
		toSerialize["billing_frequency_cycles"] = o.BillingFrequencyCycles.Get()
	}
	if o.BillingStartDate.IsSet() {
		toSerialize["billing_start_date"] = o.BillingStartDate.Get()
	}
	return toSerialize, nil
}

func (o *AddDealProductRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"product_id",
		"item_price",
		"quantity",
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

	varAddDealProductRequest := _AddDealProductRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddDealProductRequest)

	if err != nil {
		return err
	}

	*o = AddDealProductRequest(varAddDealProductRequest)

	return err
}

type NullableAddDealProductRequest struct {
	value *AddDealProductRequest
	isSet bool
}

func (v NullableAddDealProductRequest) Get() *AddDealProductRequest {
	return v.value
}

func (v *NullableAddDealProductRequest) Set(val *AddDealProductRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddDealProductRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddDealProductRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddDealProductRequest(val *AddDealProductRequest) *NullableAddDealProductRequest {
	return &NullableAddDealProductRequest{value: val, isSet: true}
}

func (v NullableAddDealProductRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddDealProductRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


