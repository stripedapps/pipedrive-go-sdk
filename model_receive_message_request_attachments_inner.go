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

// checks if the ReceiveMessageRequestAttachmentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReceiveMessageRequestAttachmentsInner{}

// ReceiveMessageRequestAttachmentsInner struct for ReceiveMessageRequestAttachmentsInner
type ReceiveMessageRequestAttachmentsInner struct {
	// The ID of the attachment
	Id string `json:"id"`
	// The mime-type of the attachment
	Type string `json:"type"`
	// The name of the attachment
	Name *string `json:"name,omitempty"`
	// The size of the attachment
	Size *float32 `json:"size,omitempty"`
	// A URL to the file
	Url string `json:"url"`
	// A URL to a preview picture of the file
	PreviewUrl *string `json:"preview_url,omitempty"`
	// If true, it will use the getMessageById endpoint for fetching updated attachment's urls. Find out more [here](https://pipedrive.readme.io/docs/implementing-messaging-app-extension)
	LinkExpires *bool `json:"link_expires,omitempty"`
}

type _ReceiveMessageRequestAttachmentsInner ReceiveMessageRequestAttachmentsInner

// NewReceiveMessageRequestAttachmentsInner instantiates a new ReceiveMessageRequestAttachmentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceiveMessageRequestAttachmentsInner(id string, type_ string, url string) *ReceiveMessageRequestAttachmentsInner {
	this := ReceiveMessageRequestAttachmentsInner{}
	this.Id = id
	this.Type = type_
	this.Url = url
	var linkExpires bool = false
	this.LinkExpires = &linkExpires
	return &this
}

// NewReceiveMessageRequestAttachmentsInnerWithDefaults instantiates a new ReceiveMessageRequestAttachmentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceiveMessageRequestAttachmentsInnerWithDefaults() *ReceiveMessageRequestAttachmentsInner {
	this := ReceiveMessageRequestAttachmentsInner{}
	var linkExpires bool = false
	this.LinkExpires = &linkExpires
	return &this
}

// GetId returns the Id field value
func (o *ReceiveMessageRequestAttachmentsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReceiveMessageRequestAttachmentsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReceiveMessageRequestAttachmentsInner) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ReceiveMessageRequestAttachmentsInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReceiveMessageRequestAttachmentsInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ReceiveMessageRequestAttachmentsInner) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ReceiveMessageRequestAttachmentsInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiveMessageRequestAttachmentsInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ReceiveMessageRequestAttachmentsInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ReceiveMessageRequestAttachmentsInner) SetName(v string) {
	o.Name = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ReceiveMessageRequestAttachmentsInner) GetSize() float32 {
	if o == nil || IsNil(o.Size) {
		var ret float32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiveMessageRequestAttachmentsInner) GetSizeOk() (*float32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ReceiveMessageRequestAttachmentsInner) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given float32 and assigns it to the Size field.
func (o *ReceiveMessageRequestAttachmentsInner) SetSize(v float32) {
	o.Size = &v
}

// GetUrl returns the Url field value
func (o *ReceiveMessageRequestAttachmentsInner) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ReceiveMessageRequestAttachmentsInner) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ReceiveMessageRequestAttachmentsInner) SetUrl(v string) {
	o.Url = v
}

// GetPreviewUrl returns the PreviewUrl field value if set, zero value otherwise.
func (o *ReceiveMessageRequestAttachmentsInner) GetPreviewUrl() string {
	if o == nil || IsNil(o.PreviewUrl) {
		var ret string
		return ret
	}
	return *o.PreviewUrl
}

// GetPreviewUrlOk returns a tuple with the PreviewUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiveMessageRequestAttachmentsInner) GetPreviewUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PreviewUrl) {
		return nil, false
	}
	return o.PreviewUrl, true
}

// HasPreviewUrl returns a boolean if a field has been set.
func (o *ReceiveMessageRequestAttachmentsInner) HasPreviewUrl() bool {
	if o != nil && !IsNil(o.PreviewUrl) {
		return true
	}

	return false
}

// SetPreviewUrl gets a reference to the given string and assigns it to the PreviewUrl field.
func (o *ReceiveMessageRequestAttachmentsInner) SetPreviewUrl(v string) {
	o.PreviewUrl = &v
}

// GetLinkExpires returns the LinkExpires field value if set, zero value otherwise.
func (o *ReceiveMessageRequestAttachmentsInner) GetLinkExpires() bool {
	if o == nil || IsNil(o.LinkExpires) {
		var ret bool
		return ret
	}
	return *o.LinkExpires
}

// GetLinkExpiresOk returns a tuple with the LinkExpires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiveMessageRequestAttachmentsInner) GetLinkExpiresOk() (*bool, bool) {
	if o == nil || IsNil(o.LinkExpires) {
		return nil, false
	}
	return o.LinkExpires, true
}

// HasLinkExpires returns a boolean if a field has been set.
func (o *ReceiveMessageRequestAttachmentsInner) HasLinkExpires() bool {
	if o != nil && !IsNil(o.LinkExpires) {
		return true
	}

	return false
}

// SetLinkExpires gets a reference to the given bool and assigns it to the LinkExpires field.
func (o *ReceiveMessageRequestAttachmentsInner) SetLinkExpires(v bool) {
	o.LinkExpires = &v
}

func (o ReceiveMessageRequestAttachmentsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReceiveMessageRequestAttachmentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	toSerialize["url"] = o.Url
	if !IsNil(o.PreviewUrl) {
		toSerialize["preview_url"] = o.PreviewUrl
	}
	if !IsNil(o.LinkExpires) {
		toSerialize["link_expires"] = o.LinkExpires
	}
	return toSerialize, nil
}

func (o *ReceiveMessageRequestAttachmentsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"url",
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

	varReceiveMessageRequestAttachmentsInner := _ReceiveMessageRequestAttachmentsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReceiveMessageRequestAttachmentsInner)

	if err != nil {
		return err
	}

	*o = ReceiveMessageRequestAttachmentsInner(varReceiveMessageRequestAttachmentsInner)

	return err
}

type NullableReceiveMessageRequestAttachmentsInner struct {
	value *ReceiveMessageRequestAttachmentsInner
	isSet bool
}

func (v NullableReceiveMessageRequestAttachmentsInner) Get() *ReceiveMessageRequestAttachmentsInner {
	return v.value
}

func (v *NullableReceiveMessageRequestAttachmentsInner) Set(val *ReceiveMessageRequestAttachmentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableReceiveMessageRequestAttachmentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableReceiveMessageRequestAttachmentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceiveMessageRequestAttachmentsInner(val *ReceiveMessageRequestAttachmentsInner) *NullableReceiveMessageRequestAttachmentsInner {
	return &NullableReceiveMessageRequestAttachmentsInner{value: val, isSet: true}
}

func (v NullableReceiveMessageRequestAttachmentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceiveMessageRequestAttachmentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


