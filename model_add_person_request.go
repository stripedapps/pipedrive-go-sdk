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

// checks if the AddPersonRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddPersonRequest{}

// AddPersonRequest struct for AddPersonRequest
type AddPersonRequest struct {
	// The name of the person
	Name string `json:"name"`
	// The ID of the user who will be marked as the owner of this person. When omitted, the authorized user ID will be used.
	OwnerId *int32 `json:"owner_id,omitempty"`
	// The ID of the organization this person will belong to
	OrgId *int32 `json:"org_id,omitempty"`
	// An email address as a string or an array of email objects related to the person. The structure of the array is as follows: `[{ \"value\": \"mail@example.com\", \"primary\": \"true\", \"label\": \"main\" }]`. Please note that only `value` is required.
	Email []BasicPersonRequestEmailInner `json:"email,omitempty"`
	// A phone number supplied as a string or an array of phone objects related to the person. The structure of the array is as follows: `[{ \"value\": \"12345\", \"primary\": \"true\", \"label\": \"mobile\" }]`. Please note that only `value` is required.
	Phone []PersonItemAllOfPhoneInner `json:"phone,omitempty"`
	// The ID of the label.
	Label *int32 `json:"label,omitempty"`
	VisibleTo *string `json:"visible_to,omitempty"`
	MarketingStatus *string `json:"marketing_status,omitempty"`
	// The optional creation date & time of the person in UTC. Format: YYYY-MM-DD HH:MM:SS
	AddTime *string `json:"add_time,omitempty"`
}

type _AddPersonRequest AddPersonRequest

// NewAddPersonRequest instantiates a new AddPersonRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPersonRequest(name string) *AddPersonRequest {
	this := AddPersonRequest{}
	this.Name = name
	return &this
}

// NewAddPersonRequestWithDefaults instantiates a new AddPersonRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPersonRequestWithDefaults() *AddPersonRequest {
	this := AddPersonRequest{}
	return &this
}

// GetName returns the Name field value
func (o *AddPersonRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddPersonRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddPersonRequest) SetName(v string) {
	o.Name = v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *AddPersonRequest) GetOwnerId() int32 {
	if o == nil || IsNil(o.OwnerId) {
		var ret int32
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPersonRequest) GetOwnerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *AddPersonRequest) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given int32 and assigns it to the OwnerId field.
func (o *AddPersonRequest) SetOwnerId(v int32) {
	o.OwnerId = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *AddPersonRequest) GetOrgId() int32 {
	if o == nil || IsNil(o.OrgId) {
		var ret int32
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPersonRequest) GetOrgIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *AddPersonRequest) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given int32 and assigns it to the OrgId field.
func (o *AddPersonRequest) SetOrgId(v int32) {
	o.OrgId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AddPersonRequest) GetEmail() []BasicPersonRequestEmailInner {
	if o == nil || IsNil(o.Email) {
		var ret []BasicPersonRequestEmailInner
		return ret
	}
	return o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPersonRequest) GetEmailOk() ([]BasicPersonRequestEmailInner, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AddPersonRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given []BasicPersonRequestEmailInner and assigns it to the Email field.
func (o *AddPersonRequest) SetEmail(v []BasicPersonRequestEmailInner) {
	o.Email = v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *AddPersonRequest) GetPhone() []PersonItemAllOfPhoneInner {
	if o == nil || IsNil(o.Phone) {
		var ret []PersonItemAllOfPhoneInner
		return ret
	}
	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPersonRequest) GetPhoneOk() ([]PersonItemAllOfPhoneInner, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *AddPersonRequest) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given []PersonItemAllOfPhoneInner and assigns it to the Phone field.
func (o *AddPersonRequest) SetPhone(v []PersonItemAllOfPhoneInner) {
	o.Phone = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *AddPersonRequest) GetLabel() int32 {
	if o == nil || IsNil(o.Label) {
		var ret int32
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPersonRequest) GetLabelOk() (*int32, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *AddPersonRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given int32 and assigns it to the Label field.
func (o *AddPersonRequest) SetLabel(v int32) {
	o.Label = &v
}

// GetVisibleTo returns the VisibleTo field value if set, zero value otherwise.
func (o *AddPersonRequest) GetVisibleTo() string {
	if o == nil || IsNil(o.VisibleTo) {
		var ret string
		return ret
	}
	return *o.VisibleTo
}

// GetVisibleToOk returns a tuple with the VisibleTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPersonRequest) GetVisibleToOk() (*string, bool) {
	if o == nil || IsNil(o.VisibleTo) {
		return nil, false
	}
	return o.VisibleTo, true
}

// HasVisibleTo returns a boolean if a field has been set.
func (o *AddPersonRequest) HasVisibleTo() bool {
	if o != nil && !IsNil(o.VisibleTo) {
		return true
	}

	return false
}

// SetVisibleTo gets a reference to the given string and assigns it to the VisibleTo field.
func (o *AddPersonRequest) SetVisibleTo(v string) {
	o.VisibleTo = &v
}

// GetMarketingStatus returns the MarketingStatus field value if set, zero value otherwise.
func (o *AddPersonRequest) GetMarketingStatus() string {
	if o == nil || IsNil(o.MarketingStatus) {
		var ret string
		return ret
	}
	return *o.MarketingStatus
}

// GetMarketingStatusOk returns a tuple with the MarketingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPersonRequest) GetMarketingStatusOk() (*string, bool) {
	if o == nil || IsNil(o.MarketingStatus) {
		return nil, false
	}
	return o.MarketingStatus, true
}

// HasMarketingStatus returns a boolean if a field has been set.
func (o *AddPersonRequest) HasMarketingStatus() bool {
	if o != nil && !IsNil(o.MarketingStatus) {
		return true
	}

	return false
}

// SetMarketingStatus gets a reference to the given string and assigns it to the MarketingStatus field.
func (o *AddPersonRequest) SetMarketingStatus(v string) {
	o.MarketingStatus = &v
}

// GetAddTime returns the AddTime field value if set, zero value otherwise.
func (o *AddPersonRequest) GetAddTime() string {
	if o == nil || IsNil(o.AddTime) {
		var ret string
		return ret
	}
	return *o.AddTime
}

// GetAddTimeOk returns a tuple with the AddTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddPersonRequest) GetAddTimeOk() (*string, bool) {
	if o == nil || IsNil(o.AddTime) {
		return nil, false
	}
	return o.AddTime, true
}

// HasAddTime returns a boolean if a field has been set.
func (o *AddPersonRequest) HasAddTime() bool {
	if o != nil && !IsNil(o.AddTime) {
		return true
	}

	return false
}

// SetAddTime gets a reference to the given string and assigns it to the AddTime field.
func (o *AddPersonRequest) SetAddTime(v string) {
	o.AddTime = &v
}

func (o AddPersonRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddPersonRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.VisibleTo) {
		toSerialize["visible_to"] = o.VisibleTo
	}
	if !IsNil(o.MarketingStatus) {
		toSerialize["marketing_status"] = o.MarketingStatus
	}
	if !IsNil(o.AddTime) {
		toSerialize["add_time"] = o.AddTime
	}
	return toSerialize, nil
}

func (o *AddPersonRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varAddPersonRequest := _AddPersonRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddPersonRequest)

	if err != nil {
		return err
	}

	*o = AddPersonRequest(varAddPersonRequest)

	return err
}

type NullableAddPersonRequest struct {
	value *AddPersonRequest
	isSet bool
}

func (v NullableAddPersonRequest) Get() *AddPersonRequest {
	return v.value
}

func (v *NullableAddPersonRequest) Set(val *AddPersonRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPersonRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPersonRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPersonRequest(val *AddPersonRequest) *NullableAddPersonRequest {
	return &NullableAddPersonRequest{value: val, isSet: true}
}

func (v NullableAddPersonRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPersonRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


