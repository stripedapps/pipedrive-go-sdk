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

// checks if the PersonItemAllOfPhoneInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonItemAllOfPhoneInner{}

// PersonItemAllOfPhoneInner struct for PersonItemAllOfPhoneInner
type PersonItemAllOfPhoneInner struct {
	// The phone number
	Value *string `json:"value,omitempty"`
	// Boolean that indicates if phone number is primary for the person or not
	Primary *bool `json:"primary,omitempty"`
	// The label that indicates the type of the phone number. (Possible values - work, home, mobile or other)
	Label *string `json:"label,omitempty"`
}

// NewPersonItemAllOfPhoneInner instantiates a new PersonItemAllOfPhoneInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonItemAllOfPhoneInner() *PersonItemAllOfPhoneInner {
	this := PersonItemAllOfPhoneInner{}
	return &this
}

// NewPersonItemAllOfPhoneInnerWithDefaults instantiates a new PersonItemAllOfPhoneInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonItemAllOfPhoneInnerWithDefaults() *PersonItemAllOfPhoneInner {
	this := PersonItemAllOfPhoneInner{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PersonItemAllOfPhoneInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonItemAllOfPhoneInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PersonItemAllOfPhoneInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PersonItemAllOfPhoneInner) SetValue(v string) {
	o.Value = &v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *PersonItemAllOfPhoneInner) GetPrimary() bool {
	if o == nil || IsNil(o.Primary) {
		var ret bool
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonItemAllOfPhoneInner) GetPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.Primary) {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *PersonItemAllOfPhoneInner) HasPrimary() bool {
	if o != nil && !IsNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given bool and assigns it to the Primary field.
func (o *PersonItemAllOfPhoneInner) SetPrimary(v bool) {
	o.Primary = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PersonItemAllOfPhoneInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonItemAllOfPhoneInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PersonItemAllOfPhoneInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PersonItemAllOfPhoneInner) SetLabel(v string) {
	o.Label = &v
}

func (o PersonItemAllOfPhoneInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonItemAllOfPhoneInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Primary) {
		toSerialize["primary"] = o.Primary
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullablePersonItemAllOfPhoneInner struct {
	value *PersonItemAllOfPhoneInner
	isSet bool
}

func (v NullablePersonItemAllOfPhoneInner) Get() *PersonItemAllOfPhoneInner {
	return v.value
}

func (v *NullablePersonItemAllOfPhoneInner) Set(val *PersonItemAllOfPhoneInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonItemAllOfPhoneInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonItemAllOfPhoneInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonItemAllOfPhoneInner(val *PersonItemAllOfPhoneInner) *NullablePersonItemAllOfPhoneInner {
	return &NullablePersonItemAllOfPhoneInner{value: val, isSet: true}
}

func (v NullablePersonItemAllOfPhoneInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonItemAllOfPhoneInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


