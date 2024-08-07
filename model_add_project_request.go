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

// checks if the AddProjectRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddProjectRequest{}

// AddProjectRequest struct for AddProjectRequest
type AddProjectRequest struct {
	// The title of the project
	Title string `json:"title"`
	// The ID of a project board
	BoardId float32 `json:"board_id"`
	// The ID of a phase on a project board
	PhaseId float32 `json:"phase_id"`
	// The description of the project
	Description *string `json:"description,omitempty"`
	// The status of the project
	Status *string `json:"status,omitempty"`
	// The ID of a project owner
	OwnerId *float32 `json:"owner_id,omitempty"`
	// The start date of the project. Format: YYYY-MM-DD.
	StartDate *string `json:"start_date,omitempty"`
	// The end date of the project. Format: YYYY-MM-DD.
	EndDate *string `json:"end_date,omitempty"`
	// An array of IDs of the deals this project is associated with
	DealIds []int32 `json:"deal_ids,omitempty"`
	// The ID of the organization this project is associated with
	OrgId *float32 `json:"org_id,omitempty"`
	// The ID of the person this project is associated with
	PersonId *float32 `json:"person_id,omitempty"`
	// An array of IDs of the labels this project has
	Labels []int32 `json:"labels,omitempty"`
	// The ID of the template the project will be based on
	TemplateId *float32 `json:"template_id,omitempty"`
}

type _AddProjectRequest AddProjectRequest

// NewAddProjectRequest instantiates a new AddProjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddProjectRequest(title string, boardId float32, phaseId float32) *AddProjectRequest {
	this := AddProjectRequest{}
	this.Title = title
	this.BoardId = boardId
	this.PhaseId = phaseId
	return &this
}

// NewAddProjectRequestWithDefaults instantiates a new AddProjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddProjectRequestWithDefaults() *AddProjectRequest {
	this := AddProjectRequest{}
	return &this
}

// GetTitle returns the Title field value
func (o *AddProjectRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AddProjectRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *AddProjectRequest) SetTitle(v string) {
	o.Title = v
}

// GetBoardId returns the BoardId field value
func (o *AddProjectRequest) GetBoardId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BoardId
}

// GetBoardIdOk returns a tuple with the BoardId field value
// and a boolean to check if the value has been set.
func (o *AddProjectRequest) GetBoardIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoardId, true
}

// SetBoardId sets field value
func (o *AddProjectRequest) SetBoardId(v float32) {
	o.BoardId = v
}

// GetPhaseId returns the PhaseId field value
func (o *AddProjectRequest) GetPhaseId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PhaseId
}

// GetPhaseIdOk returns a tuple with the PhaseId field value
// and a boolean to check if the value has been set.
func (o *AddProjectRequest) GetPhaseIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhaseId, true
}

// SetPhaseId sets field value
func (o *AddProjectRequest) SetPhaseId(v float32) {
	o.PhaseId = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddProjectRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddProjectRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddProjectRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddProjectRequest) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AddProjectRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddProjectRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AddProjectRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AddProjectRequest) SetStatus(v string) {
	o.Status = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *AddProjectRequest) GetOwnerId() float32 {
	if o == nil || IsNil(o.OwnerId) {
		var ret float32
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddProjectRequest) GetOwnerIdOk() (*float32, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *AddProjectRequest) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given float32 and assigns it to the OwnerId field.
func (o *AddProjectRequest) SetOwnerId(v float32) {
	o.OwnerId = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *AddProjectRequest) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddProjectRequest) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *AddProjectRequest) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *AddProjectRequest) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *AddProjectRequest) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddProjectRequest) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *AddProjectRequest) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *AddProjectRequest) SetEndDate(v string) {
	o.EndDate = &v
}

// GetDealIds returns the DealIds field value if set, zero value otherwise.
func (o *AddProjectRequest) GetDealIds() []int32 {
	if o == nil || IsNil(o.DealIds) {
		var ret []int32
		return ret
	}
	return o.DealIds
}

// GetDealIdsOk returns a tuple with the DealIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddProjectRequest) GetDealIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.DealIds) {
		return nil, false
	}
	return o.DealIds, true
}

// HasDealIds returns a boolean if a field has been set.
func (o *AddProjectRequest) HasDealIds() bool {
	if o != nil && !IsNil(o.DealIds) {
		return true
	}

	return false
}

// SetDealIds gets a reference to the given []int32 and assigns it to the DealIds field.
func (o *AddProjectRequest) SetDealIds(v []int32) {
	o.DealIds = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *AddProjectRequest) GetOrgId() float32 {
	if o == nil || IsNil(o.OrgId) {
		var ret float32
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddProjectRequest) GetOrgIdOk() (*float32, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *AddProjectRequest) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given float32 and assigns it to the OrgId field.
func (o *AddProjectRequest) SetOrgId(v float32) {
	o.OrgId = &v
}

// GetPersonId returns the PersonId field value if set, zero value otherwise.
func (o *AddProjectRequest) GetPersonId() float32 {
	if o == nil || IsNil(o.PersonId) {
		var ret float32
		return ret
	}
	return *o.PersonId
}

// GetPersonIdOk returns a tuple with the PersonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddProjectRequest) GetPersonIdOk() (*float32, bool) {
	if o == nil || IsNil(o.PersonId) {
		return nil, false
	}
	return o.PersonId, true
}

// HasPersonId returns a boolean if a field has been set.
func (o *AddProjectRequest) HasPersonId() bool {
	if o != nil && !IsNil(o.PersonId) {
		return true
	}

	return false
}

// SetPersonId gets a reference to the given float32 and assigns it to the PersonId field.
func (o *AddProjectRequest) SetPersonId(v float32) {
	o.PersonId = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *AddProjectRequest) GetLabels() []int32 {
	if o == nil || IsNil(o.Labels) {
		var ret []int32
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddProjectRequest) GetLabelsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *AddProjectRequest) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []int32 and assigns it to the Labels field.
func (o *AddProjectRequest) SetLabels(v []int32) {
	o.Labels = v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *AddProjectRequest) GetTemplateId() float32 {
	if o == nil || IsNil(o.TemplateId) {
		var ret float32
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddProjectRequest) GetTemplateIdOk() (*float32, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *AddProjectRequest) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given float32 and assigns it to the TemplateId field.
func (o *AddProjectRequest) SetTemplateId(v float32) {
	o.TemplateId = &v
}

func (o AddProjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddProjectRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	toSerialize["board_id"] = o.BoardId
	toSerialize["phase_id"] = o.PhaseId
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	if !IsNil(o.DealIds) {
		toSerialize["deal_ids"] = o.DealIds
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.PersonId) {
		toSerialize["person_id"] = o.PersonId
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	return toSerialize, nil
}

func (o *AddProjectRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"board_id",
		"phase_id",
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

	varAddProjectRequest := _AddProjectRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddProjectRequest)

	if err != nil {
		return err
	}

	*o = AddProjectRequest(varAddProjectRequest)

	return err
}

type NullableAddProjectRequest struct {
	value *AddProjectRequest
	isSet bool
}

func (v NullableAddProjectRequest) Get() *AddProjectRequest {
	return v.value
}

func (v *NullableAddProjectRequest) Set(val *AddProjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddProjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddProjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddProjectRequest(val *AddProjectRequest) *NullableAddProjectRequest {
	return &NullableAddProjectRequest{value: val, isSet: true}
}

func (v NullableAddProjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddProjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


