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

// checks if the AddLeadRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddLeadRequest{}

// AddLeadRequest struct for AddLeadRequest
type AddLeadRequest struct {
	// The name of the lead
	Title string `json:"title"`
	// The ID of the user which will be the owner of the created lead. If not provided, the user making the request will be used.
	OwnerId *int32 `json:"owner_id,omitempty"`
	// The IDs of the lead labels which will be associated with the lead
	LabelIds []string `json:"label_ids,omitempty"`
	// The ID of a person which this lead will be linked to. If the person does not exist yet, it needs to be created first. This property is required unless `organization_id` is specified.
	PersonId *int32 `json:"person_id,omitempty"`
	// The ID of an organization which this lead will be linked to. If the organization does not exist yet, it needs to be created first. This property is required unless `person_id` is specified.
	OrganizationId *int32 `json:"organization_id,omitempty"`
	Value *GetLeadsResponse200DataInnerValue `json:"value,omitempty"`
	// The date of when the deal which will be created from the lead is expected to be closed. In ISO 8601 format: YYYY-MM-DD.
	ExpectedCloseDate *string `json:"expected_close_date,omitempty"`
	VisibleTo *string `json:"visible_to,omitempty"`
	// A flag indicating whether the lead was seen by someone in the Pipedrive UI
	WasSeen *bool `json:"was_seen,omitempty"`
}

// NewAddLeadRequest instantiates a new AddLeadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddLeadRequest(title string) *AddLeadRequest {
	this := AddLeadRequest{}
	this.Title = title
	return &this
}

// NewAddLeadRequestWithDefaults instantiates a new AddLeadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddLeadRequestWithDefaults() *AddLeadRequest {
	this := AddLeadRequest{}
	return &this
}

// GetTitle returns the Title field value
func (o *AddLeadRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AddLeadRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *AddLeadRequest) SetTitle(v string) {
	o.Title = v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *AddLeadRequest) GetOwnerId() int32 {
	if o == nil || IsNil(o.OwnerId) {
		var ret int32
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLeadRequest) GetOwnerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *AddLeadRequest) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given int32 and assigns it to the OwnerId field.
func (o *AddLeadRequest) SetOwnerId(v int32) {
	o.OwnerId = &v
}

// GetLabelIds returns the LabelIds field value if set, zero value otherwise.
func (o *AddLeadRequest) GetLabelIds() []string {
	if o == nil || IsNil(o.LabelIds) {
		var ret []string
		return ret
	}
	return o.LabelIds
}

// GetLabelIdsOk returns a tuple with the LabelIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLeadRequest) GetLabelIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.LabelIds) {
		return nil, false
	}
	return o.LabelIds, true
}

// HasLabelIds returns a boolean if a field has been set.
func (o *AddLeadRequest) HasLabelIds() bool {
	if o != nil && !IsNil(o.LabelIds) {
		return true
	}

	return false
}

// SetLabelIds gets a reference to the given []string and assigns it to the LabelIds field.
func (o *AddLeadRequest) SetLabelIds(v []string) {
	o.LabelIds = v
}

// GetPersonId returns the PersonId field value if set, zero value otherwise.
func (o *AddLeadRequest) GetPersonId() int32 {
	if o == nil || IsNil(o.PersonId) {
		var ret int32
		return ret
	}
	return *o.PersonId
}

// GetPersonIdOk returns a tuple with the PersonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLeadRequest) GetPersonIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PersonId) {
		return nil, false
	}
	return o.PersonId, true
}

// HasPersonId returns a boolean if a field has been set.
func (o *AddLeadRequest) HasPersonId() bool {
	if o != nil && !IsNil(o.PersonId) {
		return true
	}

	return false
}

// SetPersonId gets a reference to the given int32 and assigns it to the PersonId field.
func (o *AddLeadRequest) SetPersonId(v int32) {
	o.PersonId = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *AddLeadRequest) GetOrganizationId() int32 {
	if o == nil || IsNil(o.OrganizationId) {
		var ret int32
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLeadRequest) GetOrganizationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *AddLeadRequest) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given int32 and assigns it to the OrganizationId field.
func (o *AddLeadRequest) SetOrganizationId(v int32) {
	o.OrganizationId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AddLeadRequest) GetValue() GetLeadsResponse200DataInnerValue {
	if o == nil || IsNil(o.Value) {
		var ret GetLeadsResponse200DataInnerValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLeadRequest) GetValueOk() (*GetLeadsResponse200DataInnerValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AddLeadRequest) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given GetLeadsResponse200DataInnerValue and assigns it to the Value field.
func (o *AddLeadRequest) SetValue(v GetLeadsResponse200DataInnerValue) {
	o.Value = &v
}

// GetExpectedCloseDate returns the ExpectedCloseDate field value if set, zero value otherwise.
func (o *AddLeadRequest) GetExpectedCloseDate() string {
	if o == nil || IsNil(o.ExpectedCloseDate) {
		var ret string
		return ret
	}
	return *o.ExpectedCloseDate
}

// GetExpectedCloseDateOk returns a tuple with the ExpectedCloseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLeadRequest) GetExpectedCloseDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpectedCloseDate) {
		return nil, false
	}
	return o.ExpectedCloseDate, true
}

// HasExpectedCloseDate returns a boolean if a field has been set.
func (o *AddLeadRequest) HasExpectedCloseDate() bool {
	if o != nil && !IsNil(o.ExpectedCloseDate) {
		return true
	}

	return false
}

// SetExpectedCloseDate gets a reference to the given string and assigns it to the ExpectedCloseDate field.
func (o *AddLeadRequest) SetExpectedCloseDate(v string) {
	o.ExpectedCloseDate = &v
}

// GetVisibleTo returns the VisibleTo field value if set, zero value otherwise.
func (o *AddLeadRequest) GetVisibleTo() string {
	if o == nil || IsNil(o.VisibleTo) {
		var ret string
		return ret
	}
	return *o.VisibleTo
}

// GetVisibleToOk returns a tuple with the VisibleTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLeadRequest) GetVisibleToOk() (*string, bool) {
	if o == nil || IsNil(o.VisibleTo) {
		return nil, false
	}
	return o.VisibleTo, true
}

// HasVisibleTo returns a boolean if a field has been set.
func (o *AddLeadRequest) HasVisibleTo() bool {
	if o != nil && !IsNil(o.VisibleTo) {
		return true
	}

	return false
}

// SetVisibleTo gets a reference to the given string and assigns it to the VisibleTo field.
func (o *AddLeadRequest) SetVisibleTo(v string) {
	o.VisibleTo = &v
}

// GetWasSeen returns the WasSeen field value if set, zero value otherwise.
func (o *AddLeadRequest) GetWasSeen() bool {
	if o == nil || IsNil(o.WasSeen) {
		var ret bool
		return ret
	}
	return *o.WasSeen
}

// GetWasSeenOk returns a tuple with the WasSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLeadRequest) GetWasSeenOk() (*bool, bool) {
	if o == nil || IsNil(o.WasSeen) {
		return nil, false
	}
	return o.WasSeen, true
}

// HasWasSeen returns a boolean if a field has been set.
func (o *AddLeadRequest) HasWasSeen() bool {
	if o != nil && !IsNil(o.WasSeen) {
		return true
	}

	return false
}

// SetWasSeen gets a reference to the given bool and assigns it to the WasSeen field.
func (o *AddLeadRequest) SetWasSeen(v bool) {
	o.WasSeen = &v
}

func (o AddLeadRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddLeadRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	if !IsNil(o.LabelIds) {
		toSerialize["label_ids"] = o.LabelIds
	}
	if !IsNil(o.PersonId) {
		toSerialize["person_id"] = o.PersonId
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.ExpectedCloseDate) {
		toSerialize["expected_close_date"] = o.ExpectedCloseDate
	}
	if !IsNil(o.VisibleTo) {
		toSerialize["visible_to"] = o.VisibleTo
	}
	if !IsNil(o.WasSeen) {
		toSerialize["was_seen"] = o.WasSeen
	}
	return toSerialize, nil
}

type NullableAddLeadRequest struct {
	value *AddLeadRequest
	isSet bool
}

func (v NullableAddLeadRequest) Get() *AddLeadRequest {
	return v.value
}

func (v *NullableAddLeadRequest) Set(val *AddLeadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddLeadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddLeadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddLeadRequest(val *AddLeadRequest) *NullableAddLeadRequest {
	return &NullableAddLeadRequest{value: val, isSet: true}
}

func (v NullableAddLeadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddLeadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


