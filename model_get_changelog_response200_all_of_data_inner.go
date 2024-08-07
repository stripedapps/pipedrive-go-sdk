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

// checks if the GetChangelogResponse200AllOfDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetChangelogResponse200AllOfDataInner{}

// GetChangelogResponse200AllOfDataInner struct for GetChangelogResponse200AllOfDataInner
type GetChangelogResponse200AllOfDataInner struct {
	// The key of the field that was changed
	FieldKey *string `json:"field_key,omitempty"`
	// The value of the field before the change
	OldValue NullableString `json:"old_value,omitempty"`
	// The value of the field after the change
	NewValue NullableString `json:"new_value,omitempty"`
	// The ID of the user who made the change
	ActorUserId *int32 `json:"actor_user_id,omitempty"`
	// The date and time of the change
	Time *string `json:"time,omitempty"`
	// The source of change, for example 'app', 'mobile', 'api', etc.
	ChangeSource NullableString `json:"change_source,omitempty"`
	// The user agent from which the change was made
	ChangeSourceUserAgent NullableString `json:"change_source_user_agent,omitempty"`
	// Whether the change was made as part of a bulk update
	IsBulkUpdateFlag *bool `json:"is_bulk_update_flag,omitempty"`
}

// NewGetChangelogResponse200AllOfDataInner instantiates a new GetChangelogResponse200AllOfDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetChangelogResponse200AllOfDataInner() *GetChangelogResponse200AllOfDataInner {
	this := GetChangelogResponse200AllOfDataInner{}
	return &this
}

// NewGetChangelogResponse200AllOfDataInnerWithDefaults instantiates a new GetChangelogResponse200AllOfDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetChangelogResponse200AllOfDataInnerWithDefaults() *GetChangelogResponse200AllOfDataInner {
	this := GetChangelogResponse200AllOfDataInner{}
	return &this
}

// GetFieldKey returns the FieldKey field value if set, zero value otherwise.
func (o *GetChangelogResponse200AllOfDataInner) GetFieldKey() string {
	if o == nil || IsNil(o.FieldKey) {
		var ret string
		return ret
	}
	return *o.FieldKey
}

// GetFieldKeyOk returns a tuple with the FieldKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChangelogResponse200AllOfDataInner) GetFieldKeyOk() (*string, bool) {
	if o == nil || IsNil(o.FieldKey) {
		return nil, false
	}
	return o.FieldKey, true
}

// HasFieldKey returns a boolean if a field has been set.
func (o *GetChangelogResponse200AllOfDataInner) HasFieldKey() bool {
	if o != nil && !IsNil(o.FieldKey) {
		return true
	}

	return false
}

// SetFieldKey gets a reference to the given string and assigns it to the FieldKey field.
func (o *GetChangelogResponse200AllOfDataInner) SetFieldKey(v string) {
	o.FieldKey = &v
}

// GetOldValue returns the OldValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetChangelogResponse200AllOfDataInner) GetOldValue() string {
	if o == nil || IsNil(o.OldValue.Get()) {
		var ret string
		return ret
	}
	return *o.OldValue.Get()
}

// GetOldValueOk returns a tuple with the OldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetChangelogResponse200AllOfDataInner) GetOldValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OldValue.Get(), o.OldValue.IsSet()
}

// HasOldValue returns a boolean if a field has been set.
func (o *GetChangelogResponse200AllOfDataInner) HasOldValue() bool {
	if o != nil && o.OldValue.IsSet() {
		return true
	}

	return false
}

// SetOldValue gets a reference to the given NullableString and assigns it to the OldValue field.
func (o *GetChangelogResponse200AllOfDataInner) SetOldValue(v string) {
	o.OldValue.Set(&v)
}
// SetOldValueNil sets the value for OldValue to be an explicit nil
func (o *GetChangelogResponse200AllOfDataInner) SetOldValueNil() {
	o.OldValue.Set(nil)
}

// UnsetOldValue ensures that no value is present for OldValue, not even an explicit nil
func (o *GetChangelogResponse200AllOfDataInner) UnsetOldValue() {
	o.OldValue.Unset()
}

// GetNewValue returns the NewValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetChangelogResponse200AllOfDataInner) GetNewValue() string {
	if o == nil || IsNil(o.NewValue.Get()) {
		var ret string
		return ret
	}
	return *o.NewValue.Get()
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetChangelogResponse200AllOfDataInner) GetNewValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewValue.Get(), o.NewValue.IsSet()
}

// HasNewValue returns a boolean if a field has been set.
func (o *GetChangelogResponse200AllOfDataInner) HasNewValue() bool {
	if o != nil && o.NewValue.IsSet() {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given NullableString and assigns it to the NewValue field.
func (o *GetChangelogResponse200AllOfDataInner) SetNewValue(v string) {
	o.NewValue.Set(&v)
}
// SetNewValueNil sets the value for NewValue to be an explicit nil
func (o *GetChangelogResponse200AllOfDataInner) SetNewValueNil() {
	o.NewValue.Set(nil)
}

// UnsetNewValue ensures that no value is present for NewValue, not even an explicit nil
func (o *GetChangelogResponse200AllOfDataInner) UnsetNewValue() {
	o.NewValue.Unset()
}

// GetActorUserId returns the ActorUserId field value if set, zero value otherwise.
func (o *GetChangelogResponse200AllOfDataInner) GetActorUserId() int32 {
	if o == nil || IsNil(o.ActorUserId) {
		var ret int32
		return ret
	}
	return *o.ActorUserId
}

// GetActorUserIdOk returns a tuple with the ActorUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChangelogResponse200AllOfDataInner) GetActorUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ActorUserId) {
		return nil, false
	}
	return o.ActorUserId, true
}

// HasActorUserId returns a boolean if a field has been set.
func (o *GetChangelogResponse200AllOfDataInner) HasActorUserId() bool {
	if o != nil && !IsNil(o.ActorUserId) {
		return true
	}

	return false
}

// SetActorUserId gets a reference to the given int32 and assigns it to the ActorUserId field.
func (o *GetChangelogResponse200AllOfDataInner) SetActorUserId(v int32) {
	o.ActorUserId = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GetChangelogResponse200AllOfDataInner) GetTime() string {
	if o == nil || IsNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChangelogResponse200AllOfDataInner) GetTimeOk() (*string, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GetChangelogResponse200AllOfDataInner) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *GetChangelogResponse200AllOfDataInner) SetTime(v string) {
	o.Time = &v
}

// GetChangeSource returns the ChangeSource field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetChangelogResponse200AllOfDataInner) GetChangeSource() string {
	if o == nil || IsNil(o.ChangeSource.Get()) {
		var ret string
		return ret
	}
	return *o.ChangeSource.Get()
}

// GetChangeSourceOk returns a tuple with the ChangeSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetChangelogResponse200AllOfDataInner) GetChangeSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChangeSource.Get(), o.ChangeSource.IsSet()
}

// HasChangeSource returns a boolean if a field has been set.
func (o *GetChangelogResponse200AllOfDataInner) HasChangeSource() bool {
	if o != nil && o.ChangeSource.IsSet() {
		return true
	}

	return false
}

// SetChangeSource gets a reference to the given NullableString and assigns it to the ChangeSource field.
func (o *GetChangelogResponse200AllOfDataInner) SetChangeSource(v string) {
	o.ChangeSource.Set(&v)
}
// SetChangeSourceNil sets the value for ChangeSource to be an explicit nil
func (o *GetChangelogResponse200AllOfDataInner) SetChangeSourceNil() {
	o.ChangeSource.Set(nil)
}

// UnsetChangeSource ensures that no value is present for ChangeSource, not even an explicit nil
func (o *GetChangelogResponse200AllOfDataInner) UnsetChangeSource() {
	o.ChangeSource.Unset()
}

// GetChangeSourceUserAgent returns the ChangeSourceUserAgent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetChangelogResponse200AllOfDataInner) GetChangeSourceUserAgent() string {
	if o == nil || IsNil(o.ChangeSourceUserAgent.Get()) {
		var ret string
		return ret
	}
	return *o.ChangeSourceUserAgent.Get()
}

// GetChangeSourceUserAgentOk returns a tuple with the ChangeSourceUserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetChangelogResponse200AllOfDataInner) GetChangeSourceUserAgentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChangeSourceUserAgent.Get(), o.ChangeSourceUserAgent.IsSet()
}

// HasChangeSourceUserAgent returns a boolean if a field has been set.
func (o *GetChangelogResponse200AllOfDataInner) HasChangeSourceUserAgent() bool {
	if o != nil && o.ChangeSourceUserAgent.IsSet() {
		return true
	}

	return false
}

// SetChangeSourceUserAgent gets a reference to the given NullableString and assigns it to the ChangeSourceUserAgent field.
func (o *GetChangelogResponse200AllOfDataInner) SetChangeSourceUserAgent(v string) {
	o.ChangeSourceUserAgent.Set(&v)
}
// SetChangeSourceUserAgentNil sets the value for ChangeSourceUserAgent to be an explicit nil
func (o *GetChangelogResponse200AllOfDataInner) SetChangeSourceUserAgentNil() {
	o.ChangeSourceUserAgent.Set(nil)
}

// UnsetChangeSourceUserAgent ensures that no value is present for ChangeSourceUserAgent, not even an explicit nil
func (o *GetChangelogResponse200AllOfDataInner) UnsetChangeSourceUserAgent() {
	o.ChangeSourceUserAgent.Unset()
}

// GetIsBulkUpdateFlag returns the IsBulkUpdateFlag field value if set, zero value otherwise.
func (o *GetChangelogResponse200AllOfDataInner) GetIsBulkUpdateFlag() bool {
	if o == nil || IsNil(o.IsBulkUpdateFlag) {
		var ret bool
		return ret
	}
	return *o.IsBulkUpdateFlag
}

// GetIsBulkUpdateFlagOk returns a tuple with the IsBulkUpdateFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChangelogResponse200AllOfDataInner) GetIsBulkUpdateFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBulkUpdateFlag) {
		return nil, false
	}
	return o.IsBulkUpdateFlag, true
}

// HasIsBulkUpdateFlag returns a boolean if a field has been set.
func (o *GetChangelogResponse200AllOfDataInner) HasIsBulkUpdateFlag() bool {
	if o != nil && !IsNil(o.IsBulkUpdateFlag) {
		return true
	}

	return false
}

// SetIsBulkUpdateFlag gets a reference to the given bool and assigns it to the IsBulkUpdateFlag field.
func (o *GetChangelogResponse200AllOfDataInner) SetIsBulkUpdateFlag(v bool) {
	o.IsBulkUpdateFlag = &v
}

func (o GetChangelogResponse200AllOfDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetChangelogResponse200AllOfDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FieldKey) {
		toSerialize["field_key"] = o.FieldKey
	}
	if o.OldValue.IsSet() {
		toSerialize["old_value"] = o.OldValue.Get()
	}
	if o.NewValue.IsSet() {
		toSerialize["new_value"] = o.NewValue.Get()
	}
	if !IsNil(o.ActorUserId) {
		toSerialize["actor_user_id"] = o.ActorUserId
	}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if o.ChangeSource.IsSet() {
		toSerialize["change_source"] = o.ChangeSource.Get()
	}
	if o.ChangeSourceUserAgent.IsSet() {
		toSerialize["change_source_user_agent"] = o.ChangeSourceUserAgent.Get()
	}
	if !IsNil(o.IsBulkUpdateFlag) {
		toSerialize["is_bulk_update_flag"] = o.IsBulkUpdateFlag
	}
	return toSerialize, nil
}

type NullableGetChangelogResponse200AllOfDataInner struct {
	value *GetChangelogResponse200AllOfDataInner
	isSet bool
}

func (v NullableGetChangelogResponse200AllOfDataInner) Get() *GetChangelogResponse200AllOfDataInner {
	return v.value
}

func (v *NullableGetChangelogResponse200AllOfDataInner) Set(val *GetChangelogResponse200AllOfDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetChangelogResponse200AllOfDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetChangelogResponse200AllOfDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetChangelogResponse200AllOfDataInner(val *GetChangelogResponse200AllOfDataInner) *NullableGetChangelogResponse200AllOfDataInner {
	return &NullableGetChangelogResponse200AllOfDataInner{value: val, isSet: true}
}

func (v NullableGetChangelogResponse200AllOfDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetChangelogResponse200AllOfDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


