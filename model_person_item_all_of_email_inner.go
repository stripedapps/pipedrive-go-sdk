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

// checks if the PersonItemAllOfEmailInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PersonItemAllOfEmailInner{}

// PersonItemAllOfEmailInner struct for PersonItemAllOfEmailInner
type PersonItemAllOfEmailInner struct {
	// Email
	Value *string `json:"value,omitempty"`
	// Boolean that indicates if email is primary for the person or not
	Primary *bool `json:"primary,omitempty"`
	// The label that indicates the type of the email. (Possible values - work, home or other)
	Label *string `json:"label,omitempty"`
}

// NewPersonItemAllOfEmailInner instantiates a new PersonItemAllOfEmailInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonItemAllOfEmailInner() *PersonItemAllOfEmailInner {
	this := PersonItemAllOfEmailInner{}
	return &this
}

// NewPersonItemAllOfEmailInnerWithDefaults instantiates a new PersonItemAllOfEmailInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonItemAllOfEmailInnerWithDefaults() *PersonItemAllOfEmailInner {
	this := PersonItemAllOfEmailInner{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PersonItemAllOfEmailInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonItemAllOfEmailInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PersonItemAllOfEmailInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PersonItemAllOfEmailInner) SetValue(v string) {
	o.Value = &v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *PersonItemAllOfEmailInner) GetPrimary() bool {
	if o == nil || IsNil(o.Primary) {
		var ret bool
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonItemAllOfEmailInner) GetPrimaryOk() (*bool, bool) {
	if o == nil || IsNil(o.Primary) {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *PersonItemAllOfEmailInner) HasPrimary() bool {
	if o != nil && !IsNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given bool and assigns it to the Primary field.
func (o *PersonItemAllOfEmailInner) SetPrimary(v bool) {
	o.Primary = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PersonItemAllOfEmailInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonItemAllOfEmailInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PersonItemAllOfEmailInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PersonItemAllOfEmailInner) SetLabel(v string) {
	o.Label = &v
}

func (o PersonItemAllOfEmailInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PersonItemAllOfEmailInner) ToMap() (map[string]interface{}, error) {
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

type NullablePersonItemAllOfEmailInner struct {
	value *PersonItemAllOfEmailInner
	isSet bool
}

func (v NullablePersonItemAllOfEmailInner) Get() *PersonItemAllOfEmailInner {
	return v.value
}

func (v *NullablePersonItemAllOfEmailInner) Set(val *PersonItemAllOfEmailInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonItemAllOfEmailInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonItemAllOfEmailInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonItemAllOfEmailInner(val *PersonItemAllOfEmailInner) *NullablePersonItemAllOfEmailInner {
	return &NullablePersonItemAllOfEmailInner{value: val, isSet: true}
}

func (v NullablePersonItemAllOfEmailInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonItemAllOfEmailInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


