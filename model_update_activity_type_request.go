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

// checks if the UpdateActivityTypeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateActivityTypeRequest{}

// UpdateActivityTypeRequest struct for UpdateActivityTypeRequest
type UpdateActivityTypeRequest struct {
	// The name of the activity type
	Name *string `json:"name,omitempty"`
	// Icon graphic to use for representing this activity type
	IconKey *string `json:"icon_key,omitempty"`
	// A designated color for the activity type in 6-character HEX format (e.g. `FFFFFF` for white, `000000` for black)
	Color *string `json:"color,omitempty"`
	// An order number for this activity type. Order numbers should be used to order the types in the activity type selections.
	OrderNr *int32 `json:"order_nr,omitempty"`
}

// NewUpdateActivityTypeRequest instantiates a new UpdateActivityTypeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateActivityTypeRequest() *UpdateActivityTypeRequest {
	this := UpdateActivityTypeRequest{}
	return &this
}

// NewUpdateActivityTypeRequestWithDefaults instantiates a new UpdateActivityTypeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateActivityTypeRequestWithDefaults() *UpdateActivityTypeRequest {
	this := UpdateActivityTypeRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateActivityTypeRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActivityTypeRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateActivityTypeRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateActivityTypeRequest) SetName(v string) {
	o.Name = &v
}

// GetIconKey returns the IconKey field value if set, zero value otherwise.
func (o *UpdateActivityTypeRequest) GetIconKey() string {
	if o == nil || IsNil(o.IconKey) {
		var ret string
		return ret
	}
	return *o.IconKey
}

// GetIconKeyOk returns a tuple with the IconKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActivityTypeRequest) GetIconKeyOk() (*string, bool) {
	if o == nil || IsNil(o.IconKey) {
		return nil, false
	}
	return o.IconKey, true
}

// HasIconKey returns a boolean if a field has been set.
func (o *UpdateActivityTypeRequest) HasIconKey() bool {
	if o != nil && !IsNil(o.IconKey) {
		return true
	}

	return false
}

// SetIconKey gets a reference to the given string and assigns it to the IconKey field.
func (o *UpdateActivityTypeRequest) SetIconKey(v string) {
	o.IconKey = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *UpdateActivityTypeRequest) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActivityTypeRequest) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *UpdateActivityTypeRequest) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *UpdateActivityTypeRequest) SetColor(v string) {
	o.Color = &v
}

// GetOrderNr returns the OrderNr field value if set, zero value otherwise.
func (o *UpdateActivityTypeRequest) GetOrderNr() int32 {
	if o == nil || IsNil(o.OrderNr) {
		var ret int32
		return ret
	}
	return *o.OrderNr
}

// GetOrderNrOk returns a tuple with the OrderNr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActivityTypeRequest) GetOrderNrOk() (*int32, bool) {
	if o == nil || IsNil(o.OrderNr) {
		return nil, false
	}
	return o.OrderNr, true
}

// HasOrderNr returns a boolean if a field has been set.
func (o *UpdateActivityTypeRequest) HasOrderNr() bool {
	if o != nil && !IsNil(o.OrderNr) {
		return true
	}

	return false
}

// SetOrderNr gets a reference to the given int32 and assigns it to the OrderNr field.
func (o *UpdateActivityTypeRequest) SetOrderNr(v int32) {
	o.OrderNr = &v
}

func (o UpdateActivityTypeRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateActivityTypeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.IconKey) {
		toSerialize["icon_key"] = o.IconKey
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.OrderNr) {
		toSerialize["order_nr"] = o.OrderNr
	}
	return toSerialize, nil
}

type NullableUpdateActivityTypeRequest struct {
	value *UpdateActivityTypeRequest
	isSet bool
}

func (v NullableUpdateActivityTypeRequest) Get() *UpdateActivityTypeRequest {
	return v.value
}

func (v *NullableUpdateActivityTypeRequest) Set(val *UpdateActivityTypeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateActivityTypeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateActivityTypeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateActivityTypeRequest(val *UpdateActivityTypeRequest) *NullableUpdateActivityTypeRequest {
	return &NullableUpdateActivityTypeRequest{value: val, isSet: true}
}

func (v NullableUpdateActivityTypeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateActivityTypeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


