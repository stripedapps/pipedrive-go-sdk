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
	"bytes"
	"fmt"
)

// checks if the ReceiveMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReceiveMessageRequest{}

// ReceiveMessageRequest struct for ReceiveMessageRequest
type ReceiveMessageRequest struct {
	// The ID of the message
	Id string `json:"id"`
	// The channel ID as in the provider
	ChannelId string `json:"channel_id"`
	// The ID of the provider's user that sent the message
	SenderId string `json:"sender_id"`
	// The ID of the conversation
	ConversationId string `json:"conversation_id"`
	// The body of the message
	Message string `json:"message"`
	// The status of the message
	Status string `json:"status"`
	// The date and time when the message was created in the provider, in UTC. Format: YYYY-MM-DD HH:MM
	CreatedAt time.Time `json:"created_at"`
	// The date and time when the message can no longer receive a reply, in UTC. Format: YYYY-MM-DD HH:MM
	ReplyBy *time.Time `json:"reply_by,omitempty"`
	// A URL that can open the conversation in the provider's side
	ConversationLink *string `json:"conversation_link,omitempty"`
	// The list of attachments available in the message
	Attachments []ReceiveMessageRequestAttachmentsInner `json:"attachments,omitempty"`
}

type _ReceiveMessageRequest ReceiveMessageRequest

// NewReceiveMessageRequest instantiates a new ReceiveMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReceiveMessageRequest(id string, channelId string, senderId string, conversationId string, message string, status string, createdAt time.Time) *ReceiveMessageRequest {
	this := ReceiveMessageRequest{}
	this.Id = id
	this.ChannelId = channelId
	this.SenderId = senderId
	this.ConversationId = conversationId
	this.Message = message
	this.Status = status
	this.CreatedAt = createdAt
	return &this
}

// NewReceiveMessageRequestWithDefaults instantiates a new ReceiveMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReceiveMessageRequestWithDefaults() *ReceiveMessageRequest {
	this := ReceiveMessageRequest{}
	return &this
}

// GetId returns the Id field value
func (o *ReceiveMessageRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReceiveMessageRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReceiveMessageRequest) SetId(v string) {
	o.Id = v
}

// GetChannelId returns the ChannelId field value
func (o *ReceiveMessageRequest) GetChannelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value
// and a boolean to check if the value has been set.
func (o *ReceiveMessageRequest) GetChannelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChannelId, true
}

// SetChannelId sets field value
func (o *ReceiveMessageRequest) SetChannelId(v string) {
	o.ChannelId = v
}

// GetSenderId returns the SenderId field value
func (o *ReceiveMessageRequest) GetSenderId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderId
}

// GetSenderIdOk returns a tuple with the SenderId field value
// and a boolean to check if the value has been set.
func (o *ReceiveMessageRequest) GetSenderIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderId, true
}

// SetSenderId sets field value
func (o *ReceiveMessageRequest) SetSenderId(v string) {
	o.SenderId = v
}

// GetConversationId returns the ConversationId field value
func (o *ReceiveMessageRequest) GetConversationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value
// and a boolean to check if the value has been set.
func (o *ReceiveMessageRequest) GetConversationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversationId, true
}

// SetConversationId sets field value
func (o *ReceiveMessageRequest) SetConversationId(v string) {
	o.ConversationId = v
}

// GetMessage returns the Message field value
func (o *ReceiveMessageRequest) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ReceiveMessageRequest) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ReceiveMessageRequest) SetMessage(v string) {
	o.Message = v
}

// GetStatus returns the Status field value
func (o *ReceiveMessageRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ReceiveMessageRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ReceiveMessageRequest) SetStatus(v string) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ReceiveMessageRequest) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ReceiveMessageRequest) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ReceiveMessageRequest) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetReplyBy returns the ReplyBy field value if set, zero value otherwise.
func (o *ReceiveMessageRequest) GetReplyBy() time.Time {
	if o == nil || IsNil(o.ReplyBy) {
		var ret time.Time
		return ret
	}
	return *o.ReplyBy
}

// GetReplyByOk returns a tuple with the ReplyBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiveMessageRequest) GetReplyByOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReplyBy) {
		return nil, false
	}
	return o.ReplyBy, true
}

// HasReplyBy returns a boolean if a field has been set.
func (o *ReceiveMessageRequest) HasReplyBy() bool {
	if o != nil && !IsNil(o.ReplyBy) {
		return true
	}

	return false
}

// SetReplyBy gets a reference to the given time.Time and assigns it to the ReplyBy field.
func (o *ReceiveMessageRequest) SetReplyBy(v time.Time) {
	o.ReplyBy = &v
}

// GetConversationLink returns the ConversationLink field value if set, zero value otherwise.
func (o *ReceiveMessageRequest) GetConversationLink() string {
	if o == nil || IsNil(o.ConversationLink) {
		var ret string
		return ret
	}
	return *o.ConversationLink
}

// GetConversationLinkOk returns a tuple with the ConversationLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiveMessageRequest) GetConversationLinkOk() (*string, bool) {
	if o == nil || IsNil(o.ConversationLink) {
		return nil, false
	}
	return o.ConversationLink, true
}

// HasConversationLink returns a boolean if a field has been set.
func (o *ReceiveMessageRequest) HasConversationLink() bool {
	if o != nil && !IsNil(o.ConversationLink) {
		return true
	}

	return false
}

// SetConversationLink gets a reference to the given string and assigns it to the ConversationLink field.
func (o *ReceiveMessageRequest) SetConversationLink(v string) {
	o.ConversationLink = &v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *ReceiveMessageRequest) GetAttachments() []ReceiveMessageRequestAttachmentsInner {
	if o == nil || IsNil(o.Attachments) {
		var ret []ReceiveMessageRequestAttachmentsInner
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReceiveMessageRequest) GetAttachmentsOk() ([]ReceiveMessageRequestAttachmentsInner, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *ReceiveMessageRequest) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []ReceiveMessageRequestAttachmentsInner and assigns it to the Attachments field.
func (o *ReceiveMessageRequest) SetAttachments(v []ReceiveMessageRequestAttachmentsInner) {
	o.Attachments = v
}

func (o ReceiveMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReceiveMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["channel_id"] = o.ChannelId
	toSerialize["sender_id"] = o.SenderId
	toSerialize["conversation_id"] = o.ConversationId
	toSerialize["message"] = o.Message
	toSerialize["status"] = o.Status
	toSerialize["created_at"] = o.CreatedAt
	if !IsNil(o.ReplyBy) {
		toSerialize["reply_by"] = o.ReplyBy
	}
	if !IsNil(o.ConversationLink) {
		toSerialize["conversation_link"] = o.ConversationLink
	}
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	return toSerialize, nil
}

func (o *ReceiveMessageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"channel_id",
		"sender_id",
		"conversation_id",
		"message",
		"status",
		"created_at",
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

	varReceiveMessageRequest := _ReceiveMessageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReceiveMessageRequest)

	if err != nil {
		return err
	}

	*o = ReceiveMessageRequest(varReceiveMessageRequest)

	return err
}

type NullableReceiveMessageRequest struct {
	value *ReceiveMessageRequest
	isSet bool
}

func (v NullableReceiveMessageRequest) Get() *ReceiveMessageRequest {
	return v.value
}

func (v *NullableReceiveMessageRequest) Set(val *ReceiveMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReceiveMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReceiveMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReceiveMessageRequest(val *ReceiveMessageRequest) *NullableReceiveMessageRequest {
	return &NullableReceiveMessageRequest{value: val, isSet: true}
}

func (v NullableReceiveMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReceiveMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


