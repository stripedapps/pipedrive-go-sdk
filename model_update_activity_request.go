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

// checks if the UpdateActivityRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateActivityRequest{}

// UpdateActivityRequest struct for UpdateActivityRequest
type UpdateActivityRequest struct {
	// The due date of the activity. Format: YYYY-MM-DD
	DueDate *string `json:"due_date,omitempty"`
	// The due time of the activity in UTC. Format: HH:MM
	DueTime *string `json:"due_time,omitempty"`
	// The duration of the activity. Format: HH:MM
	Duration *string `json:"duration,omitempty"`
	// The ID of the deal this activity is associated with
	DealId *int32 `json:"deal_id,omitempty"`
	// The ID of the lead in the UUID format this activity is associated with
	LeadId NullableString `json:"lead_id,omitempty"`
	// The ID of the person this activity is associated with
	PersonId *int32 `json:"person_id,omitempty"`
	// The ID of the project this activity is associated with
	ProjectId NullableInt32 `json:"project_id,omitempty"`
	// The ID of the organization this activity is associated with
	OrgId *int32 `json:"org_id,omitempty"`
	// The address of the activity.
	Location *string `json:"location,omitempty"`
	// Additional details about the activity that is synced to your external calendar. Unlike the note added to the activity, the description is publicly visible to any guests added to the activity.
	PublicDescription *string `json:"public_description,omitempty"`
	// The note of the activity (HTML format)
	Note *string `json:"note,omitempty"`
	// The subject of the activity
	Subject *string `json:"subject,omitempty"`
	// The type of the activity. This is in correlation with the `key_string` parameter of ActivityTypes.
	Type *string `json:"type,omitempty"`
	// The ID of the user whom the activity is assigned to
	UserId *int32 `json:"user_id,omitempty"`
	// List of multiple persons (participants) this activity is associated with. It requires a structure as follows: `[{\"person_id\":1,\"primary_flag\":true}]`
	Participants []map[string]interface{} `json:"participants,omitempty"`
	// Set the activity as 'Busy' or 'Free'. If the flag is set to `true`, your customers will not be able to book that time slot through any Scheduler links. The flag can also be unset by never setting it or overriding it with `null`. When the value of the flag is unset (`null`), the flag defaults to 'Busy' if it has a time set, and 'Free' if it is an all-day event without specified time.
	BusyFlag *bool `json:"busy_flag,omitempty"`
	// The attendees of the activity. This can be either your existing Pipedrive contacts or an external email address. It requires a structure as follows: `[{\"email_address\":\"mail@example.org\"}]` or `[{\"person_id\":1, \"email_address\":\"mail@example.org\"}]`
	Attendees []map[string]interface{} `json:"attendees,omitempty"`
	Done *float32 `json:"done,omitempty"`
}

// NewUpdateActivityRequest instantiates a new UpdateActivityRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateActivityRequest() *UpdateActivityRequest {
	this := UpdateActivityRequest{}
	return &this
}

// NewUpdateActivityRequestWithDefaults instantiates a new UpdateActivityRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateActivityRequestWithDefaults() *UpdateActivityRequest {
	this := UpdateActivityRequest{}
	return &this
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *UpdateActivityRequest) GetDueDate() string {
	if o == nil || IsNil(o.DueDate) {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActivityRequest) GetDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *UpdateActivityRequest) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *UpdateActivityRequest) SetDueDate(v string) {
	o.DueDate = &v
}

// GetDueTime returns the DueTime field value if set, zero value otherwise.
func (o *UpdateActivityRequest) GetDueTime() string {
	if o == nil || IsNil(o.DueTime) {
		var ret string
		return ret
	}
	return *o.DueTime
}

// GetDueTimeOk returns a tuple with the DueTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActivityRequest) GetDueTimeOk() (*string, bool) {
	if o == nil || IsNil(o.DueTime) {
		return nil, false
	}
	return o.DueTime, true
}

// HasDueTime returns a boolean if a field has been set.
func (o *UpdateActivityRequest) HasDueTime() bool {
	if o != nil && !IsNil(o.DueTime) {
		return true
	}

	return false
}

// SetDueTime gets a reference to the given string and assigns it to the DueTime field.
func (o *UpdateActivityRequest) SetDueTime(v string) {
	o.DueTime = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *UpdateActivityRequest) GetDuration() string {
	if o == nil || IsNil(o.Duration) {
		var ret string
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActivityRequest) GetDurationOk() (*string, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *UpdateActivityRequest) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the Duration field.
func (o *UpdateActivityRequest) SetDuration(v string) {
	o.Duration = &v
}

// GetDealId returns the DealId field value if set, zero value otherwise.
func (o *UpdateActivityRequest) GetDealId() int32 {
	if o == nil || IsNil(o.DealId) {
		var ret int32
		return ret
	}
	return *o.DealId
}

// GetDealIdOk returns a tuple with the DealId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActivityRequest) GetDealIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DealId) {
		return nil, false
	}
	return o.DealId, true
}

// HasDealId returns a boolean if a field has been set.
func (o *UpdateActivityRequest) HasDealId() bool {
	if o != nil && !IsNil(o.DealId) {
		return true
	}

	return false
}

// SetDealId gets a reference to the given int32 and assigns it to the DealId field.
func (o *UpdateActivityRequest) SetDealId(v int32) {
	o.DealId = &v
}

// GetLeadId returns the LeadId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateActivityRequest) GetLeadId() string {
	if o == nil || IsNil(o.LeadId.Get()) {
		var ret string
		return ret
	}
	return *o.LeadId.Get()
}

// GetLeadIdOk returns a tuple with the LeadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateActivityRequest) GetLeadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LeadId.Get(), o.LeadId.IsSet()
}

// HasLeadId returns a boolean if a field has been set.
func (o *UpdateActivityRequest) HasLeadId() bool {
	if o != nil && o.LeadId.IsSet() {
		return true
	}

	return false
}

// SetLeadId gets a reference to the given NullableString and assigns it to the LeadId field.
func (o *UpdateActivityRequest) SetLeadId(v string) {
	o.LeadId.Set(&v)
}
// SetLeadIdNil sets the value for LeadId to be an explicit nil
func (o *UpdateActivityRequest) SetLeadIdNil() {
	o.LeadId.Set(nil)
}

// UnsetLeadId ensures that no value is present for LeadId, not even an explicit nil
func (o *UpdateActivityRequest) UnsetLeadId() {
	o.LeadId.Unset()
}

// GetPersonId returns the PersonId field value if set, zero value otherwise.
func (o *UpdateActivityRequest) GetPersonId() int32 {
	if o == nil || IsNil(o.PersonId) {
		var ret int32
		return ret
	}
	return *o.PersonId
}

// GetPersonIdOk returns a tuple with the PersonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActivityRequest) GetPersonIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PersonId) {
		return nil, false
	}
	return o.PersonId, true
}

// HasPersonId returns a boolean if a field has been set.
func (o *UpdateActivityRequest) HasPersonId() bool {
	if o != nil && !IsNil(o.PersonId) {
		return true
	}

	return false
}

// SetPersonId gets a reference to the given int32 and assigns it to the PersonId field.
func (o *UpdateActivityRequest) SetPersonId(v int32) {
	o.PersonId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateActivityRequest) GetProjectId() int32 {
	if o == nil || IsNil(o.ProjectId.Get()) {
		var ret int32
		return ret
	}
	return *o.ProjectId.Get()
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateActivityRequest) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProjectId.Get(), o.ProjectId.IsSet()
}

// HasProjectId returns a boolean if a field has been set.
func (o *UpdateActivityRequest) HasProjectId() bool {
	if o != nil && o.ProjectId.IsSet() {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given NullableInt32 and assigns it to the ProjectId field.
func (o *UpdateActivityRequest) SetProjectId(v int32) {
	o.ProjectId.Set(&v)
}
// SetProjectIdNil sets the value for ProjectId to be an explicit nil
func (o *UpdateActivityRequest) SetProjectIdNil() {
	o.ProjectId.Set(nil)
}

// UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
func (o *UpdateActivityRequest) UnsetProjectId() {
	o.ProjectId.Unset()
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *UpdateActivityRequest) GetOrgId() int32 {
	if o == nil || IsNil(o.OrgId) {
		var ret int32
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActivityRequest) GetOrgIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *UpdateActivityRequest) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given int32 and assigns it to the OrgId field.
func (o *UpdateActivityRequest) SetOrgId(v int32) {
	o.OrgId = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *UpdateActivityRequest) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActivityRequest) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *UpdateActivityRequest) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *UpdateActivityRequest) SetLocation(v string) {
	o.Location = &v
}

// GetPublicDescription returns the PublicDescription field value if set, zero value otherwise.
func (o *UpdateActivityRequest) GetPublicDescription() string {
	if o == nil || IsNil(o.PublicDescription) {
		var ret string
		return ret
	}
	return *o.PublicDescription
}

// GetPublicDescriptionOk returns a tuple with the PublicDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActivityRequest) GetPublicDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.PublicDescription) {
		return nil, false
	}
	return o.PublicDescription, true
}

// HasPublicDescription returns a boolean if a field has been set.
func (o *UpdateActivityRequest) HasPublicDescription() bool {
	if o != nil && !IsNil(o.PublicDescription) {
		return true
	}

	return false
}

// SetPublicDescription gets a reference to the given string and assigns it to the PublicDescription field.
func (o *UpdateActivityRequest) SetPublicDescription(v string) {
	o.PublicDescription = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *UpdateActivityRequest) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActivityRequest) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *UpdateActivityRequest) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *UpdateActivityRequest) SetNote(v string) {
	o.Note = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *UpdateActivityRequest) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActivityRequest) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *UpdateActivityRequest) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *UpdateActivityRequest) SetSubject(v string) {
	o.Subject = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateActivityRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActivityRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateActivityRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UpdateActivityRequest) SetType(v string) {
	o.Type = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UpdateActivityRequest) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActivityRequest) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UpdateActivityRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *UpdateActivityRequest) SetUserId(v int32) {
	o.UserId = &v
}

// GetParticipants returns the Participants field value if set, zero value otherwise.
func (o *UpdateActivityRequest) GetParticipants() []map[string]interface{} {
	if o == nil || IsNil(o.Participants) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Participants
}

// GetParticipantsOk returns a tuple with the Participants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActivityRequest) GetParticipantsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Participants) {
		return nil, false
	}
	return o.Participants, true
}

// HasParticipants returns a boolean if a field has been set.
func (o *UpdateActivityRequest) HasParticipants() bool {
	if o != nil && !IsNil(o.Participants) {
		return true
	}

	return false
}

// SetParticipants gets a reference to the given []map[string]interface{} and assigns it to the Participants field.
func (o *UpdateActivityRequest) SetParticipants(v []map[string]interface{}) {
	o.Participants = v
}

// GetBusyFlag returns the BusyFlag field value if set, zero value otherwise.
func (o *UpdateActivityRequest) GetBusyFlag() bool {
	if o == nil || IsNil(o.BusyFlag) {
		var ret bool
		return ret
	}
	return *o.BusyFlag
}

// GetBusyFlagOk returns a tuple with the BusyFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActivityRequest) GetBusyFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.BusyFlag) {
		return nil, false
	}
	return o.BusyFlag, true
}

// HasBusyFlag returns a boolean if a field has been set.
func (o *UpdateActivityRequest) HasBusyFlag() bool {
	if o != nil && !IsNil(o.BusyFlag) {
		return true
	}

	return false
}

// SetBusyFlag gets a reference to the given bool and assigns it to the BusyFlag field.
func (o *UpdateActivityRequest) SetBusyFlag(v bool) {
	o.BusyFlag = &v
}

// GetAttendees returns the Attendees field value if set, zero value otherwise.
func (o *UpdateActivityRequest) GetAttendees() []map[string]interface{} {
	if o == nil || IsNil(o.Attendees) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Attendees
}

// GetAttendeesOk returns a tuple with the Attendees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActivityRequest) GetAttendeesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Attendees) {
		return nil, false
	}
	return o.Attendees, true
}

// HasAttendees returns a boolean if a field has been set.
func (o *UpdateActivityRequest) HasAttendees() bool {
	if o != nil && !IsNil(o.Attendees) {
		return true
	}

	return false
}

// SetAttendees gets a reference to the given []map[string]interface{} and assigns it to the Attendees field.
func (o *UpdateActivityRequest) SetAttendees(v []map[string]interface{}) {
	o.Attendees = v
}

// GetDone returns the Done field value if set, zero value otherwise.
func (o *UpdateActivityRequest) GetDone() float32 {
	if o == nil || IsNil(o.Done) {
		var ret float32
		return ret
	}
	return *o.Done
}

// GetDoneOk returns a tuple with the Done field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActivityRequest) GetDoneOk() (*float32, bool) {
	if o == nil || IsNil(o.Done) {
		return nil, false
	}
	return o.Done, true
}

// HasDone returns a boolean if a field has been set.
func (o *UpdateActivityRequest) HasDone() bool {
	if o != nil && !IsNil(o.Done) {
		return true
	}

	return false
}

// SetDone gets a reference to the given float32 and assigns it to the Done field.
func (o *UpdateActivityRequest) SetDone(v float32) {
	o.Done = &v
}

func (o UpdateActivityRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateActivityRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DueDate) {
		toSerialize["due_date"] = o.DueDate
	}
	if !IsNil(o.DueTime) {
		toSerialize["due_time"] = o.DueTime
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.DealId) {
		toSerialize["deal_id"] = o.DealId
	}
	if o.LeadId.IsSet() {
		toSerialize["lead_id"] = o.LeadId.Get()
	}
	if !IsNil(o.PersonId) {
		toSerialize["person_id"] = o.PersonId
	}
	if o.ProjectId.IsSet() {
		toSerialize["project_id"] = o.ProjectId.Get()
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.PublicDescription) {
		toSerialize["public_description"] = o.PublicDescription
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.Participants) {
		toSerialize["participants"] = o.Participants
	}
	if !IsNil(o.BusyFlag) {
		toSerialize["busy_flag"] = o.BusyFlag
	}
	if !IsNil(o.Attendees) {
		toSerialize["attendees"] = o.Attendees
	}
	if !IsNil(o.Done) {
		toSerialize["done"] = o.Done
	}
	return toSerialize, nil
}

type NullableUpdateActivityRequest struct {
	value *UpdateActivityRequest
	isSet bool
}

func (v NullableUpdateActivityRequest) Get() *UpdateActivityRequest {
	return v.value
}

func (v *NullableUpdateActivityRequest) Set(val *UpdateActivityRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateActivityRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateActivityRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateActivityRequest(val *UpdateActivityRequest) *NullableUpdateActivityRequest {
	return &NullableUpdateActivityRequest{value: val, isSet: true}
}

func (v NullableUpdateActivityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateActivityRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


