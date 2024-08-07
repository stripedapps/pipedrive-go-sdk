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

// checks if the ResponseCallLogObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseCallLogObject{}

// ResponseCallLogObject struct for ResponseCallLogObject
type ResponseCallLogObject struct {
	// The ID of the owner of the call log. Please note that a user without account settings access cannot create call logs for other users.
	UserId *int32 `json:"user_id,omitempty"`
	// If specified, this activity will be converted into a call log, with the information provided. When this field is used, you don't need to specify `deal_id`, `person_id` or `org_id`, as they will be ignored in favor of the values already available in the activity. The `activity_id` must refer to a `call` type activity.
	ActivityId *int32 `json:"activity_id,omitempty"`
	// The name of the activity this call is attached to
	Subject *string `json:"subject,omitempty"`
	// The duration of the call in seconds
	Duration *string `json:"duration,omitempty"`
	// Describes the outcome of the call
	Outcome string `json:"outcome"`
	// The number that made the call
	FromPhoneNumber *string `json:"from_phone_number,omitempty"`
	// The number called
	ToPhoneNumber string `json:"to_phone_number"`
	// The date and time of the start of the call in UTC. Format: YYYY-MM-DD HH:MM:SS.
	StartTime time.Time `json:"start_time"`
	// The date and time of the end of the call in UTC. Format: YYYY-MM-DD HH:MM:SS.
	EndTime time.Time `json:"end_time"`
	// The ID of the person this call is associated with
	PersonId *int32 `json:"person_id,omitempty"`
	// The ID of the organization this call is associated with
	OrgId *int32 `json:"org_id,omitempty"`
	// The ID of the deal this call is associated with. A call log can be associated with either a deal or a lead, but not both at once.
	DealId *int32 `json:"deal_id,omitempty"`
	// The ID of the lead in the UUID format this call is associated with. A call log can be associated with either a deal or a lead, but not both at once.
	LeadId *string `json:"lead_id,omitempty"`
	// The note for the call log in HTML format
	Note *string `json:"note,omitempty"`
	// The call log ID, generated when the call log was created
	Id *string `json:"id,omitempty"`
	// If the call log has an audio recording attached, the value should be true
	HasRecording *bool `json:"has_recording,omitempty"`
	// The company ID of the owner of the call log
	CompanyId *int32 `json:"company_id,omitempty"`
}

type _ResponseCallLogObject ResponseCallLogObject

// NewResponseCallLogObject instantiates a new ResponseCallLogObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseCallLogObject(outcome string, toPhoneNumber string, startTime time.Time, endTime time.Time) *ResponseCallLogObject {
	this := ResponseCallLogObject{}
	this.Outcome = outcome
	this.ToPhoneNumber = toPhoneNumber
	this.StartTime = startTime
	this.EndTime = endTime
	return &this
}

// NewResponseCallLogObjectWithDefaults instantiates a new ResponseCallLogObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseCallLogObjectWithDefaults() *ResponseCallLogObject {
	this := ResponseCallLogObject{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ResponseCallLogObject) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCallLogObject) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ResponseCallLogObject) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *ResponseCallLogObject) SetUserId(v int32) {
	o.UserId = &v
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *ResponseCallLogObject) GetActivityId() int32 {
	if o == nil || IsNil(o.ActivityId) {
		var ret int32
		return ret
	}
	return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCallLogObject) GetActivityIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ActivityId) {
		return nil, false
	}
	return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *ResponseCallLogObject) HasActivityId() bool {
	if o != nil && !IsNil(o.ActivityId) {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given int32 and assigns it to the ActivityId field.
func (o *ResponseCallLogObject) SetActivityId(v int32) {
	o.ActivityId = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *ResponseCallLogObject) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCallLogObject) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *ResponseCallLogObject) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *ResponseCallLogObject) SetSubject(v string) {
	o.Subject = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *ResponseCallLogObject) GetDuration() string {
	if o == nil || IsNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCallLogObject) GetDurationOk() (*string, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *ResponseCallLogObject) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *ResponseCallLogObject) SetDuration(v string) {
	o.Duration = &v
}

// GetOutcome returns the Outcome field value
func (o *ResponseCallLogObject) GetOutcome() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value
// and a boolean to check if the value has been set.
func (o *ResponseCallLogObject) GetOutcomeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Outcome, true
}

// SetOutcome sets field value
func (o *ResponseCallLogObject) SetOutcome(v string) {
	o.Outcome = v
}

// GetFromPhoneNumber returns the FromPhoneNumber field value if set, zero value otherwise.
func (o *ResponseCallLogObject) GetFromPhoneNumber() string {
	if o == nil || IsNil(o.FromPhoneNumber) {
		var ret string
		return ret
	}
	return *o.FromPhoneNumber
}

// GetFromPhoneNumberOk returns a tuple with the FromPhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCallLogObject) GetFromPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.FromPhoneNumber) {
		return nil, false
	}
	return o.FromPhoneNumber, true
}

// HasFromPhoneNumber returns a boolean if a field has been set.
func (o *ResponseCallLogObject) HasFromPhoneNumber() bool {
	if o != nil && !IsNil(o.FromPhoneNumber) {
		return true
	}

	return false
}

// SetFromPhoneNumber gets a reference to the given string and assigns it to the FromPhoneNumber field.
func (o *ResponseCallLogObject) SetFromPhoneNumber(v string) {
	o.FromPhoneNumber = &v
}

// GetToPhoneNumber returns the ToPhoneNumber field value
func (o *ResponseCallLogObject) GetToPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToPhoneNumber
}

// GetToPhoneNumberOk returns a tuple with the ToPhoneNumber field value
// and a boolean to check if the value has been set.
func (o *ResponseCallLogObject) GetToPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToPhoneNumber, true
}

// SetToPhoneNumber sets field value
func (o *ResponseCallLogObject) SetToPhoneNumber(v string) {
	o.ToPhoneNumber = v
}

// GetStartTime returns the StartTime field value
func (o *ResponseCallLogObject) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *ResponseCallLogObject) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *ResponseCallLogObject) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *ResponseCallLogObject) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *ResponseCallLogObject) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *ResponseCallLogObject) SetEndTime(v time.Time) {
	o.EndTime = v
}

// GetPersonId returns the PersonId field value if set, zero value otherwise.
func (o *ResponseCallLogObject) GetPersonId() int32 {
	if o == nil || IsNil(o.PersonId) {
		var ret int32
		return ret
	}
	return *o.PersonId
}

// GetPersonIdOk returns a tuple with the PersonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCallLogObject) GetPersonIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PersonId) {
		return nil, false
	}
	return o.PersonId, true
}

// HasPersonId returns a boolean if a field has been set.
func (o *ResponseCallLogObject) HasPersonId() bool {
	if o != nil && !IsNil(o.PersonId) {
		return true
	}

	return false
}

// SetPersonId gets a reference to the given int32 and assigns it to the PersonId field.
func (o *ResponseCallLogObject) SetPersonId(v int32) {
	o.PersonId = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *ResponseCallLogObject) GetOrgId() int32 {
	if o == nil || IsNil(o.OrgId) {
		var ret int32
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCallLogObject) GetOrgIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *ResponseCallLogObject) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given int32 and assigns it to the OrgId field.
func (o *ResponseCallLogObject) SetOrgId(v int32) {
	o.OrgId = &v
}

// GetDealId returns the DealId field value if set, zero value otherwise.
func (o *ResponseCallLogObject) GetDealId() int32 {
	if o == nil || IsNil(o.DealId) {
		var ret int32
		return ret
	}
	return *o.DealId
}

// GetDealIdOk returns a tuple with the DealId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCallLogObject) GetDealIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DealId) {
		return nil, false
	}
	return o.DealId, true
}

// HasDealId returns a boolean if a field has been set.
func (o *ResponseCallLogObject) HasDealId() bool {
	if o != nil && !IsNil(o.DealId) {
		return true
	}

	return false
}

// SetDealId gets a reference to the given int32 and assigns it to the DealId field.
func (o *ResponseCallLogObject) SetDealId(v int32) {
	o.DealId = &v
}

// GetLeadId returns the LeadId field value if set, zero value otherwise.
func (o *ResponseCallLogObject) GetLeadId() string {
	if o == nil || IsNil(o.LeadId) {
		var ret string
		return ret
	}
	return *o.LeadId
}

// GetLeadIdOk returns a tuple with the LeadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCallLogObject) GetLeadIdOk() (*string, bool) {
	if o == nil || IsNil(o.LeadId) {
		return nil, false
	}
	return o.LeadId, true
}

// HasLeadId returns a boolean if a field has been set.
func (o *ResponseCallLogObject) HasLeadId() bool {
	if o != nil && !IsNil(o.LeadId) {
		return true
	}

	return false
}

// SetLeadId gets a reference to the given string and assigns it to the LeadId field.
func (o *ResponseCallLogObject) SetLeadId(v string) {
	o.LeadId = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *ResponseCallLogObject) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCallLogObject) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *ResponseCallLogObject) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *ResponseCallLogObject) SetNote(v string) {
	o.Note = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResponseCallLogObject) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCallLogObject) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResponseCallLogObject) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResponseCallLogObject) SetId(v string) {
	o.Id = &v
}

// GetHasRecording returns the HasRecording field value if set, zero value otherwise.
func (o *ResponseCallLogObject) GetHasRecording() bool {
	if o == nil || IsNil(o.HasRecording) {
		var ret bool
		return ret
	}
	return *o.HasRecording
}

// GetHasRecordingOk returns a tuple with the HasRecording field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCallLogObject) GetHasRecordingOk() (*bool, bool) {
	if o == nil || IsNil(o.HasRecording) {
		return nil, false
	}
	return o.HasRecording, true
}

// HasHasRecording returns a boolean if a field has been set.
func (o *ResponseCallLogObject) HasHasRecording() bool {
	if o != nil && !IsNil(o.HasRecording) {
		return true
	}

	return false
}

// SetHasRecording gets a reference to the given bool and assigns it to the HasRecording field.
func (o *ResponseCallLogObject) SetHasRecording(v bool) {
	o.HasRecording = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *ResponseCallLogObject) GetCompanyId() int32 {
	if o == nil || IsNil(o.CompanyId) {
		var ret int32
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseCallLogObject) GetCompanyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CompanyId) {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *ResponseCallLogObject) HasCompanyId() bool {
	if o != nil && !IsNil(o.CompanyId) {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given int32 and assigns it to the CompanyId field.
func (o *ResponseCallLogObject) SetCompanyId(v int32) {
	o.CompanyId = &v
}

func (o ResponseCallLogObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseCallLogObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.ActivityId) {
		toSerialize["activity_id"] = o.ActivityId
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	toSerialize["outcome"] = o.Outcome
	if !IsNil(o.FromPhoneNumber) {
		toSerialize["from_phone_number"] = o.FromPhoneNumber
	}
	toSerialize["to_phone_number"] = o.ToPhoneNumber
	toSerialize["start_time"] = o.StartTime
	toSerialize["end_time"] = o.EndTime
	if !IsNil(o.PersonId) {
		toSerialize["person_id"] = o.PersonId
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.DealId) {
		toSerialize["deal_id"] = o.DealId
	}
	if !IsNil(o.LeadId) {
		toSerialize["lead_id"] = o.LeadId
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.HasRecording) {
		toSerialize["has_recording"] = o.HasRecording
	}
	if !IsNil(o.CompanyId) {
		toSerialize["company_id"] = o.CompanyId
	}
	return toSerialize, nil
}

func (o *ResponseCallLogObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"outcome",
		"to_phone_number",
		"start_time",
		"end_time",
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

	varResponseCallLogObject := _ResponseCallLogObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResponseCallLogObject)

	if err != nil {
		return err
	}

	*o = ResponseCallLogObject(varResponseCallLogObject)

	return err
}

type NullableResponseCallLogObject struct {
	value *ResponseCallLogObject
	isSet bool
}

func (v NullableResponseCallLogObject) Get() *ResponseCallLogObject {
	return v.value
}

func (v *NullableResponseCallLogObject) Set(val *ResponseCallLogObject) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseCallLogObject) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseCallLogObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseCallLogObject(val *ResponseCallLogObject) *NullableResponseCallLogObject {
	return &NullableResponseCallLogObject{value: val, isSet: true}
}

func (v NullableResponseCallLogObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseCallLogObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


