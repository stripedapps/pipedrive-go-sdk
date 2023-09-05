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

// checks if the MailMessageItemForListAllOfFromInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MailMessageItemForListAllOfFromInner{}

// MailMessageItemForListAllOfFromInner struct for MailMessageItemForListAllOfFromInner
type MailMessageItemForListAllOfFromInner struct {
	// ID of the mail participant
	Id *int32 `json:"id,omitempty"`
	// Mail address of the mail participant
	EmailAddress *string `json:"email_address,omitempty"`
	// Name of the mail participant
	Name *string `json:"name,omitempty"`
	// ID of the linked person to the mail message
	LinkedPersonId *int32 `json:"linked_person_id,omitempty"`
	// Name of the linked person to the mail message
	LinkedPersonName *string `json:"linked_person_name,omitempty"`
	// ID of the mail message participant
	MailMessagePartyId *int32 `json:"mail_message_party_id,omitempty"`
}

// NewMailMessageItemForListAllOfFromInner instantiates a new MailMessageItemForListAllOfFromInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMailMessageItemForListAllOfFromInner() *MailMessageItemForListAllOfFromInner {
	this := MailMessageItemForListAllOfFromInner{}
	return &this
}

// NewMailMessageItemForListAllOfFromInnerWithDefaults instantiates a new MailMessageItemForListAllOfFromInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMailMessageItemForListAllOfFromInnerWithDefaults() *MailMessageItemForListAllOfFromInner {
	this := MailMessageItemForListAllOfFromInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MailMessageItemForListAllOfFromInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailMessageItemForListAllOfFromInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MailMessageItemForListAllOfFromInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MailMessageItemForListAllOfFromInner) SetId(v int32) {
	o.Id = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *MailMessageItemForListAllOfFromInner) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailMessageItemForListAllOfFromInner) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *MailMessageItemForListAllOfFromInner) HasEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *MailMessageItemForListAllOfFromInner) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MailMessageItemForListAllOfFromInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailMessageItemForListAllOfFromInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MailMessageItemForListAllOfFromInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MailMessageItemForListAllOfFromInner) SetName(v string) {
	o.Name = &v
}

// GetLinkedPersonId returns the LinkedPersonId field value if set, zero value otherwise.
func (o *MailMessageItemForListAllOfFromInner) GetLinkedPersonId() int32 {
	if o == nil || IsNil(o.LinkedPersonId) {
		var ret int32
		return ret
	}
	return *o.LinkedPersonId
}

// GetLinkedPersonIdOk returns a tuple with the LinkedPersonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailMessageItemForListAllOfFromInner) GetLinkedPersonIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LinkedPersonId) {
		return nil, false
	}
	return o.LinkedPersonId, true
}

// HasLinkedPersonId returns a boolean if a field has been set.
func (o *MailMessageItemForListAllOfFromInner) HasLinkedPersonId() bool {
	if o != nil && !IsNil(o.LinkedPersonId) {
		return true
	}

	return false
}

// SetLinkedPersonId gets a reference to the given int32 and assigns it to the LinkedPersonId field.
func (o *MailMessageItemForListAllOfFromInner) SetLinkedPersonId(v int32) {
	o.LinkedPersonId = &v
}

// GetLinkedPersonName returns the LinkedPersonName field value if set, zero value otherwise.
func (o *MailMessageItemForListAllOfFromInner) GetLinkedPersonName() string {
	if o == nil || IsNil(o.LinkedPersonName) {
		var ret string
		return ret
	}
	return *o.LinkedPersonName
}

// GetLinkedPersonNameOk returns a tuple with the LinkedPersonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailMessageItemForListAllOfFromInner) GetLinkedPersonNameOk() (*string, bool) {
	if o == nil || IsNil(o.LinkedPersonName) {
		return nil, false
	}
	return o.LinkedPersonName, true
}

// HasLinkedPersonName returns a boolean if a field has been set.
func (o *MailMessageItemForListAllOfFromInner) HasLinkedPersonName() bool {
	if o != nil && !IsNil(o.LinkedPersonName) {
		return true
	}

	return false
}

// SetLinkedPersonName gets a reference to the given string and assigns it to the LinkedPersonName field.
func (o *MailMessageItemForListAllOfFromInner) SetLinkedPersonName(v string) {
	o.LinkedPersonName = &v
}

// GetMailMessagePartyId returns the MailMessagePartyId field value if set, zero value otherwise.
func (o *MailMessageItemForListAllOfFromInner) GetMailMessagePartyId() int32 {
	if o == nil || IsNil(o.MailMessagePartyId) {
		var ret int32
		return ret
	}
	return *o.MailMessagePartyId
}

// GetMailMessagePartyIdOk returns a tuple with the MailMessagePartyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MailMessageItemForListAllOfFromInner) GetMailMessagePartyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MailMessagePartyId) {
		return nil, false
	}
	return o.MailMessagePartyId, true
}

// HasMailMessagePartyId returns a boolean if a field has been set.
func (o *MailMessageItemForListAllOfFromInner) HasMailMessagePartyId() bool {
	if o != nil && !IsNil(o.MailMessagePartyId) {
		return true
	}

	return false
}

// SetMailMessagePartyId gets a reference to the given int32 and assigns it to the MailMessagePartyId field.
func (o *MailMessageItemForListAllOfFromInner) SetMailMessagePartyId(v int32) {
	o.MailMessagePartyId = &v
}

func (o MailMessageItemForListAllOfFromInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MailMessageItemForListAllOfFromInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.EmailAddress) {
		toSerialize["email_address"] = o.EmailAddress
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.LinkedPersonId) {
		toSerialize["linked_person_id"] = o.LinkedPersonId
	}
	if !IsNil(o.LinkedPersonName) {
		toSerialize["linked_person_name"] = o.LinkedPersonName
	}
	if !IsNil(o.MailMessagePartyId) {
		toSerialize["mail_message_party_id"] = o.MailMessagePartyId
	}
	return toSerialize, nil
}

type NullableMailMessageItemForListAllOfFromInner struct {
	value *MailMessageItemForListAllOfFromInner
	isSet bool
}

func (v NullableMailMessageItemForListAllOfFromInner) Get() *MailMessageItemForListAllOfFromInner {
	return v.value
}

func (v *NullableMailMessageItemForListAllOfFromInner) Set(val *MailMessageItemForListAllOfFromInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMailMessageItemForListAllOfFromInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMailMessageItemForListAllOfFromInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMailMessageItemForListAllOfFromInner(val *MailMessageItemForListAllOfFromInner) *NullableMailMessageItemForListAllOfFromInner {
	return &NullableMailMessageItemForListAllOfFromInner{value: val, isSet: true}
}

func (v NullableMailMessageItemForListAllOfFromInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMailMessageItemForListAllOfFromInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


