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

// checks if the AddProductRequest1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddProductRequest1{}

// AddProductRequest1 struct for AddProductRequest1
type AddProductRequest1 struct {
	// The name of the product
	Name string `json:"name"`
	// The product code
	Code *string `json:"code,omitempty"`
	// The unit in which this product is sold
	Unit *string `json:"unit,omitempty"`
	// The tax percentage
	Tax *float32 `json:"tax,omitempty"`
	// Whether this product will be made active or not
	ActiveFlag *bool `json:"active_flag,omitempty"`
	// Whether this product can be selected in deals or not
	Selectable *bool `json:"selectable,omitempty"`
	VisibleTo *string `json:"visible_to,omitempty"`
	// The ID of the user who will be marked as the owner of this product. When omitted, the authorized user ID will be used.
	OwnerId *int32 `json:"owner_id,omitempty"`
	// An array of objects, each containing: `currency` (string), `price` (number), `cost` (number, optional), `overhead_cost` (number, optional). Note that there can only be one price per product per currency. When `prices` is omitted altogether, a default price of 0 and a default currency based on the company's currency will be assigned.
	Prices []map[string]interface{} `json:"prices,omitempty"`
}

// NewAddProductRequest1 instantiates a new AddProductRequest1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddProductRequest1(name string) *AddProductRequest1 {
	this := AddProductRequest1{}
	this.Name = name
	var tax float32 = 0
	this.Tax = &tax
	var activeFlag bool = true
	this.ActiveFlag = &activeFlag
	var selectable bool = true
	this.Selectable = &selectable
	return &this
}

// NewAddProductRequest1WithDefaults instantiates a new AddProductRequest1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddProductRequest1WithDefaults() *AddProductRequest1 {
	this := AddProductRequest1{}
	var tax float32 = 0
	this.Tax = &tax
	var activeFlag bool = true
	this.ActiveFlag = &activeFlag
	var selectable bool = true
	this.Selectable = &selectable
	return &this
}

// GetName returns the Name field value
func (o *AddProductRequest1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddProductRequest1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddProductRequest1) SetName(v string) {
	o.Name = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AddProductRequest1) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddProductRequest1) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AddProductRequest1) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AddProductRequest1) SetCode(v string) {
	o.Code = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *AddProductRequest1) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddProductRequest1) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *AddProductRequest1) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *AddProductRequest1) SetUnit(v string) {
	o.Unit = &v
}

// GetTax returns the Tax field value if set, zero value otherwise.
func (o *AddProductRequest1) GetTax() float32 {
	if o == nil || IsNil(o.Tax) {
		var ret float32
		return ret
	}
	return *o.Tax
}

// GetTaxOk returns a tuple with the Tax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddProductRequest1) GetTaxOk() (*float32, bool) {
	if o == nil || IsNil(o.Tax) {
		return nil, false
	}
	return o.Tax, true
}

// HasTax returns a boolean if a field has been set.
func (o *AddProductRequest1) HasTax() bool {
	if o != nil && !IsNil(o.Tax) {
		return true
	}

	return false
}

// SetTax gets a reference to the given float32 and assigns it to the Tax field.
func (o *AddProductRequest1) SetTax(v float32) {
	o.Tax = &v
}

// GetActiveFlag returns the ActiveFlag field value if set, zero value otherwise.
func (o *AddProductRequest1) GetActiveFlag() bool {
	if o == nil || IsNil(o.ActiveFlag) {
		var ret bool
		return ret
	}
	return *o.ActiveFlag
}

// GetActiveFlagOk returns a tuple with the ActiveFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddProductRequest1) GetActiveFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.ActiveFlag) {
		return nil, false
	}
	return o.ActiveFlag, true
}

// HasActiveFlag returns a boolean if a field has been set.
func (o *AddProductRequest1) HasActiveFlag() bool {
	if o != nil && !IsNil(o.ActiveFlag) {
		return true
	}

	return false
}

// SetActiveFlag gets a reference to the given bool and assigns it to the ActiveFlag field.
func (o *AddProductRequest1) SetActiveFlag(v bool) {
	o.ActiveFlag = &v
}

// GetSelectable returns the Selectable field value if set, zero value otherwise.
func (o *AddProductRequest1) GetSelectable() bool {
	if o == nil || IsNil(o.Selectable) {
		var ret bool
		return ret
	}
	return *o.Selectable
}

// GetSelectableOk returns a tuple with the Selectable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddProductRequest1) GetSelectableOk() (*bool, bool) {
	if o == nil || IsNil(o.Selectable) {
		return nil, false
	}
	return o.Selectable, true
}

// HasSelectable returns a boolean if a field has been set.
func (o *AddProductRequest1) HasSelectable() bool {
	if o != nil && !IsNil(o.Selectable) {
		return true
	}

	return false
}

// SetSelectable gets a reference to the given bool and assigns it to the Selectable field.
func (o *AddProductRequest1) SetSelectable(v bool) {
	o.Selectable = &v
}

// GetVisibleTo returns the VisibleTo field value if set, zero value otherwise.
func (o *AddProductRequest1) GetVisibleTo() string {
	if o == nil || IsNil(o.VisibleTo) {
		var ret string
		return ret
	}
	return *o.VisibleTo
}

// GetVisibleToOk returns a tuple with the VisibleTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddProductRequest1) GetVisibleToOk() (*string, bool) {
	if o == nil || IsNil(o.VisibleTo) {
		return nil, false
	}
	return o.VisibleTo, true
}

// HasVisibleTo returns a boolean if a field has been set.
func (o *AddProductRequest1) HasVisibleTo() bool {
	if o != nil && !IsNil(o.VisibleTo) {
		return true
	}

	return false
}

// SetVisibleTo gets a reference to the given string and assigns it to the VisibleTo field.
func (o *AddProductRequest1) SetVisibleTo(v string) {
	o.VisibleTo = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *AddProductRequest1) GetOwnerId() int32 {
	if o == nil || IsNil(o.OwnerId) {
		var ret int32
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddProductRequest1) GetOwnerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *AddProductRequest1) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given int32 and assigns it to the OwnerId field.
func (o *AddProductRequest1) SetOwnerId(v int32) {
	o.OwnerId = &v
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *AddProductRequest1) GetPrices() []map[string]interface{} {
	if o == nil || IsNil(o.Prices) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddProductRequest1) GetPricesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Prices) {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *AddProductRequest1) HasPrices() bool {
	if o != nil && !IsNil(o.Prices) {
		return true
	}

	return false
}

// SetPrices gets a reference to the given []map[string]interface{} and assigns it to the Prices field.
func (o *AddProductRequest1) SetPrices(v []map[string]interface{}) {
	o.Prices = v
}

func (o AddProductRequest1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddProductRequest1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !IsNil(o.Tax) {
		toSerialize["tax"] = o.Tax
	}
	if !IsNil(o.ActiveFlag) {
		toSerialize["active_flag"] = o.ActiveFlag
	}
	if !IsNil(o.Selectable) {
		toSerialize["selectable"] = o.Selectable
	}
	if !IsNil(o.VisibleTo) {
		toSerialize["visible_to"] = o.VisibleTo
	}
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	if !IsNil(o.Prices) {
		toSerialize["prices"] = o.Prices
	}
	return toSerialize, nil
}

type NullableAddProductRequest1 struct {
	value *AddProductRequest1
	isSet bool
}

func (v NullableAddProductRequest1) Get() *AddProductRequest1 {
	return v.value
}

func (v *NullableAddProductRequest1) Set(val *AddProductRequest1) {
	v.value = val
	v.isSet = true
}

func (v NullableAddProductRequest1) IsSet() bool {
	return v.isSet
}

func (v *NullableAddProductRequest1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddProductRequest1(val *AddProductRequest1) *NullableAddProductRequest1 {
	return &NullableAddProductRequest1{value: val, isSet: true}
}

func (v NullableAddProductRequest1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddProductRequest1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


