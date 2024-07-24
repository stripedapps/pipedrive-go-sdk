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

// checks if the CreateFieldRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFieldRequest{}

// CreateFieldRequest struct for CreateFieldRequest
type CreateFieldRequest struct {
	// The name of the field
	Name string `json:"name"`
	// When `field_type` is either set or enum, possible options must be supplied as a JSON-encoded sequential array of objects. Example: `[{\"label\":\"New Item\"}]`
	Options []map[string]interface{} `json:"options,omitempty"`
	// Whether the field is available in the 'add new' modal or not (both in the web and mobile app)
	AddVisibleFlag *bool `json:"add_visible_flag,omitempty"`
	// The type of the field<table><tr><th>Value</th><th>Description</th></tr><tr><td>`address`</td><td>Address field</td></tr><tr><td>`date`</td><td>Date (format YYYY-MM-DD)</td></tr><tr><td>`daterange`</td><td>Date-range field (has a start date and end date value, both YYYY-MM-DD)</td></tr><tr><td>`double`</td><td>Numeric value</td></tr><tr><td>`enum`</td><td>Options field with a single possible chosen option</td></tr><tr></tr><tr><td>`monetary`</td><td>Monetary field (has a numeric value and a currency value)</td></tr><tr><td>`org`</td><td>Organization field (contains an organization ID which is stored on the same account)</td></tr><tr><td>`people`</td><td>Person field (contains a person ID which is stored on the same account)</td></tr><tr><td>`phone`</td><td>Phone field (up to 255 numbers and/or characters)</td></tr><tr><td>`set`</td><td>Options field with a possibility of having multiple chosen options</td></tr><tr><td>`text`</td><td>Long text (up to 65k characters)</td></tr><tr><td>`time`</td><td>Time field (format HH:MM:SS)</td></tr><tr><td>`timerange`</td><td>Time-range field (has a start time and end time value, both HH:MM:SS)</td></tr><tr><td>`user`</td><td>User field (contains a user ID of another Pipedrive user)</td></tr><tr><td>`varchar`</td><td>Text (up to 255 characters)</td></tr><tr><td>`varchar_auto`</td><td>Autocomplete text (up to 255 characters)</td></tr><tr><td>`visible_to`</td><td>System field that keeps item's visibility setting</td></tr></table>
	FieldType string `json:"field_type"`
}

type _CreateFieldRequest CreateFieldRequest

// NewCreateFieldRequest instantiates a new CreateFieldRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFieldRequest(name string, fieldType string) *CreateFieldRequest {
	this := CreateFieldRequest{}
	this.Name = name
	var addVisibleFlag bool = true
	this.AddVisibleFlag = &addVisibleFlag
	this.FieldType = fieldType
	return &this
}

// NewCreateFieldRequestWithDefaults instantiates a new CreateFieldRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFieldRequestWithDefaults() *CreateFieldRequest {
	this := CreateFieldRequest{}
	var addVisibleFlag bool = true
	this.AddVisibleFlag = &addVisibleFlag
	return &this
}

// GetName returns the Name field value
func (o *CreateFieldRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateFieldRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateFieldRequest) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *CreateFieldRequest) GetOptions() []map[string]interface{} {
	if o == nil || IsNil(o.Options) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFieldRequest) GetOptionsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *CreateFieldRequest) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []map[string]interface{} and assigns it to the Options field.
func (o *CreateFieldRequest) SetOptions(v []map[string]interface{}) {
	o.Options = v
}

// GetAddVisibleFlag returns the AddVisibleFlag field value if set, zero value otherwise.
func (o *CreateFieldRequest) GetAddVisibleFlag() bool {
	if o == nil || IsNil(o.AddVisibleFlag) {
		var ret bool
		return ret
	}
	return *o.AddVisibleFlag
}

// GetAddVisibleFlagOk returns a tuple with the AddVisibleFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFieldRequest) GetAddVisibleFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.AddVisibleFlag) {
		return nil, false
	}
	return o.AddVisibleFlag, true
}

// HasAddVisibleFlag returns a boolean if a field has been set.
func (o *CreateFieldRequest) HasAddVisibleFlag() bool {
	if o != nil && !IsNil(o.AddVisibleFlag) {
		return true
	}

	return false
}

// SetAddVisibleFlag gets a reference to the given bool and assigns it to the AddVisibleFlag field.
func (o *CreateFieldRequest) SetAddVisibleFlag(v bool) {
	o.AddVisibleFlag = &v
}

// GetFieldType returns the FieldType field value
func (o *CreateFieldRequest) GetFieldType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldType
}

// GetFieldTypeOk returns a tuple with the FieldType field value
// and a boolean to check if the value has been set.
func (o *CreateFieldRequest) GetFieldTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldType, true
}

// SetFieldType sets field value
func (o *CreateFieldRequest) SetFieldType(v string) {
	o.FieldType = v
}

func (o CreateFieldRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFieldRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.AddVisibleFlag) {
		toSerialize["add_visible_flag"] = o.AddVisibleFlag
	}
	toSerialize["field_type"] = o.FieldType
	return toSerialize, nil
}

func (o *CreateFieldRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"field_type",
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

	varCreateFieldRequest := _CreateFieldRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateFieldRequest)

	if err != nil {
		return err
	}

	*o = CreateFieldRequest(varCreateFieldRequest)

	return err
}

type NullableCreateFieldRequest struct {
	value *CreateFieldRequest
	isSet bool
}

func (v NullableCreateFieldRequest) Get() *CreateFieldRequest {
	return v.value
}

func (v *NullableCreateFieldRequest) Set(val *CreateFieldRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFieldRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFieldRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFieldRequest(val *CreateFieldRequest) *NullableCreateFieldRequest {
	return &NullableCreateFieldRequest{value: val, isSet: true}
}

func (v NullableCreateFieldRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFieldRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


