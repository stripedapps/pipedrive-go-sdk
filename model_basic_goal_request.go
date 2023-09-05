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

// checks if the BasicGoalRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BasicGoalRequest{}

// BasicGoalRequest struct for BasicGoalRequest
type BasicGoalRequest struct {
	// The title of the goal
	Title *string `json:"title,omitempty"`
	// Who this goal is assigned to. It requires the following JSON structure: `{ \"id\": \"1\", \"type\": \"person\" }`. `type` can be either `person`, `company` or `team`. ID of the assignee person, company or team.
	Assignee map[string]interface{} `json:"assignee,omitempty"`
	// The type of the goal. It requires the following JSON structure: `{ \"name\": \"deals_started\", \"params\": { \"pipeline_id\": [1, 2], \"activity_type_id\": [9] } }`. Type can be one of: `deals_won`, `deals_progressed`, `activities_completed`, `activities_added`, `deals_started` or `revenue_forecast`. `params` can include `pipeline_id`, `stage_id` or `activity_type_id`. `stage_id` is related to only `deals_progressed` type of goals and `activity_type_id` to `activities_completed` or `activities_added` types of goals. The `pipeline_id` and `activity_type_id` need to be given as an array of integers. To track the goal in all pipelines, set `pipeline_id` as `null` and similarly, to track the goal for all activities, set `activity_type_id` as `null`.”
	Type map[string]interface{} `json:"type,omitempty"`
	// The expected outcome of the goal. Expected outcome can be tracked either by `quantity` or by `sum`. It requires the following JSON structure: `{ \"target\": \"50\", \"tracking_metric\": \"quantity\" }` or `{ \"target\": \"50\", \"tracking_metric\": \"sum\", \"currency_id\": 1 }`. `currency_id` should only be added to `sum` type of goals.
	ExpectedOutcome map[string]interface{} `json:"expected_outcome,omitempty"`
	// The date when the goal starts and ends. It requires the following JSON structure: `{ \"start\": \"2019-01-01\", \"end\": \"2022-12-31\" }`. Date in format of YYYY-MM-DD. \"end\" can be set to `null` for an infinite, open-ended goal.
	Duration map[string]interface{} `json:"duration,omitempty"`
	// The interval of the goal
	Interval *string `json:"interval,omitempty"`
}

// NewBasicGoalRequest instantiates a new BasicGoalRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicGoalRequest() *BasicGoalRequest {
	this := BasicGoalRequest{}
	return &this
}

// NewBasicGoalRequestWithDefaults instantiates a new BasicGoalRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicGoalRequestWithDefaults() *BasicGoalRequest {
	this := BasicGoalRequest{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *BasicGoalRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicGoalRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *BasicGoalRequest) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *BasicGoalRequest) SetTitle(v string) {
	o.Title = &v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *BasicGoalRequest) GetAssignee() map[string]interface{} {
	if o == nil || IsNil(o.Assignee) {
		var ret map[string]interface{}
		return ret
	}
	return o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicGoalRequest) GetAssigneeOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Assignee) {
		return map[string]interface{}{}, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *BasicGoalRequest) HasAssignee() bool {
	if o != nil && !IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given map[string]interface{} and assigns it to the Assignee field.
func (o *BasicGoalRequest) SetAssignee(v map[string]interface{}) {
	o.Assignee = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BasicGoalRequest) GetType() map[string]interface{} {
	if o == nil || IsNil(o.Type) {
		var ret map[string]interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicGoalRequest) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BasicGoalRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given map[string]interface{} and assigns it to the Type field.
func (o *BasicGoalRequest) SetType(v map[string]interface{}) {
	o.Type = v
}

// GetExpectedOutcome returns the ExpectedOutcome field value if set, zero value otherwise.
func (o *BasicGoalRequest) GetExpectedOutcome() map[string]interface{} {
	if o == nil || IsNil(o.ExpectedOutcome) {
		var ret map[string]interface{}
		return ret
	}
	return o.ExpectedOutcome
}

// GetExpectedOutcomeOk returns a tuple with the ExpectedOutcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicGoalRequest) GetExpectedOutcomeOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ExpectedOutcome) {
		return map[string]interface{}{}, false
	}
	return o.ExpectedOutcome, true
}

// HasExpectedOutcome returns a boolean if a field has been set.
func (o *BasicGoalRequest) HasExpectedOutcome() bool {
	if o != nil && !IsNil(o.ExpectedOutcome) {
		return true
	}

	return false
}

// SetExpectedOutcome gets a reference to the given map[string]interface{} and assigns it to the ExpectedOutcome field.
func (o *BasicGoalRequest) SetExpectedOutcome(v map[string]interface{}) {
	o.ExpectedOutcome = v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *BasicGoalRequest) GetDuration() map[string]interface{} {
	if o == nil || IsNil(o.Duration) {
		var ret map[string]interface{}
		return ret
	}
	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicGoalRequest) GetDurationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Duration) {
		return map[string]interface{}{}, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *BasicGoalRequest) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given map[string]interface{} and assigns it to the Duration field.
func (o *BasicGoalRequest) SetDuration(v map[string]interface{}) {
	o.Duration = v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *BasicGoalRequest) GetInterval() string {
	if o == nil || IsNil(o.Interval) {
		var ret string
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicGoalRequest) GetIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *BasicGoalRequest) HasInterval() bool {
	if o != nil && !IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given string and assigns it to the Interval field.
func (o *BasicGoalRequest) SetInterval(v string) {
	o.Interval = &v
}

func (o BasicGoalRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BasicGoalRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Assignee) {
		toSerialize["assignee"] = o.Assignee
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ExpectedOutcome) {
		toSerialize["expected_outcome"] = o.ExpectedOutcome
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	return toSerialize, nil
}

type NullableBasicGoalRequest struct {
	value *BasicGoalRequest
	isSet bool
}

func (v NullableBasicGoalRequest) Get() *BasicGoalRequest {
	return v.value
}

func (v *NullableBasicGoalRequest) Set(val *BasicGoalRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicGoalRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicGoalRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicGoalRequest(val *BasicGoalRequest) *NullableBasicGoalRequest {
	return &NullableBasicGoalRequest{value: val, isSet: true}
}

func (v NullableBasicGoalRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicGoalRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


