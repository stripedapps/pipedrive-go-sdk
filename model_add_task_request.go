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

// checks if the AddTaskRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddTaskRequest{}

// AddTaskRequest struct for AddTaskRequest
type AddTaskRequest struct {
	// The title of the task
	Title string `json:"title"`
	// The ID of a project
	ProjectId float32 `json:"project_id"`
	// The description of the task
	Description *string `json:"description,omitempty"`
	// The ID of a parent task. Can not be ID of a task which is already a subtask.
	ParentTaskId *float32 `json:"parent_task_id,omitempty"`
	// The ID of the user who will be the assignee of the task
	AssigneeId *float32 `json:"assignee_id,omitempty"`
	Done *float32 `json:"done,omitempty"`
	// The due date of the task. Format: YYYY-MM-DD.
	DueDate *string `json:"due_date,omitempty"`
}

type _AddTaskRequest AddTaskRequest

// NewAddTaskRequest instantiates a new AddTaskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddTaskRequest(title string, projectId float32) *AddTaskRequest {
	this := AddTaskRequest{}
	this.Title = title
	this.ProjectId = projectId
	return &this
}

// NewAddTaskRequestWithDefaults instantiates a new AddTaskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddTaskRequestWithDefaults() *AddTaskRequest {
	this := AddTaskRequest{}
	return &this
}

// GetTitle returns the Title field value
func (o *AddTaskRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AddTaskRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *AddTaskRequest) SetTitle(v string) {
	o.Title = v
}

// GetProjectId returns the ProjectId field value
func (o *AddTaskRequest) GetProjectId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *AddTaskRequest) GetProjectIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *AddTaskRequest) SetProjectId(v float32) {
	o.ProjectId = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddTaskRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTaskRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddTaskRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddTaskRequest) SetDescription(v string) {
	o.Description = &v
}

// GetParentTaskId returns the ParentTaskId field value if set, zero value otherwise.
func (o *AddTaskRequest) GetParentTaskId() float32 {
	if o == nil || IsNil(o.ParentTaskId) {
		var ret float32
		return ret
	}
	return *o.ParentTaskId
}

// GetParentTaskIdOk returns a tuple with the ParentTaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTaskRequest) GetParentTaskIdOk() (*float32, bool) {
	if o == nil || IsNil(o.ParentTaskId) {
		return nil, false
	}
	return o.ParentTaskId, true
}

// HasParentTaskId returns a boolean if a field has been set.
func (o *AddTaskRequest) HasParentTaskId() bool {
	if o != nil && !IsNil(o.ParentTaskId) {
		return true
	}

	return false
}

// SetParentTaskId gets a reference to the given float32 and assigns it to the ParentTaskId field.
func (o *AddTaskRequest) SetParentTaskId(v float32) {
	o.ParentTaskId = &v
}

// GetAssigneeId returns the AssigneeId field value if set, zero value otherwise.
func (o *AddTaskRequest) GetAssigneeId() float32 {
	if o == nil || IsNil(o.AssigneeId) {
		var ret float32
		return ret
	}
	return *o.AssigneeId
}

// GetAssigneeIdOk returns a tuple with the AssigneeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTaskRequest) GetAssigneeIdOk() (*float32, bool) {
	if o == nil || IsNil(o.AssigneeId) {
		return nil, false
	}
	return o.AssigneeId, true
}

// HasAssigneeId returns a boolean if a field has been set.
func (o *AddTaskRequest) HasAssigneeId() bool {
	if o != nil && !IsNil(o.AssigneeId) {
		return true
	}

	return false
}

// SetAssigneeId gets a reference to the given float32 and assigns it to the AssigneeId field.
func (o *AddTaskRequest) SetAssigneeId(v float32) {
	o.AssigneeId = &v
}

// GetDone returns the Done field value if set, zero value otherwise.
func (o *AddTaskRequest) GetDone() float32 {
	if o == nil || IsNil(o.Done) {
		var ret float32
		return ret
	}
	return *o.Done
}

// GetDoneOk returns a tuple with the Done field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTaskRequest) GetDoneOk() (*float32, bool) {
	if o == nil || IsNil(o.Done) {
		return nil, false
	}
	return o.Done, true
}

// HasDone returns a boolean if a field has been set.
func (o *AddTaskRequest) HasDone() bool {
	if o != nil && !IsNil(o.Done) {
		return true
	}

	return false
}

// SetDone gets a reference to the given float32 and assigns it to the Done field.
func (o *AddTaskRequest) SetDone(v float32) {
	o.Done = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *AddTaskRequest) GetDueDate() string {
	if o == nil || IsNil(o.DueDate) {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddTaskRequest) GetDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *AddTaskRequest) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *AddTaskRequest) SetDueDate(v string) {
	o.DueDate = &v
}

func (o AddTaskRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddTaskRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	toSerialize["project_id"] = o.ProjectId
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ParentTaskId) {
		toSerialize["parent_task_id"] = o.ParentTaskId
	}
	if !IsNil(o.AssigneeId) {
		toSerialize["assignee_id"] = o.AssigneeId
	}
	if !IsNil(o.Done) {
		toSerialize["done"] = o.Done
	}
	if !IsNil(o.DueDate) {
		toSerialize["due_date"] = o.DueDate
	}
	return toSerialize, nil
}

func (o *AddTaskRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"project_id",
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

	varAddTaskRequest := _AddTaskRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddTaskRequest)

	if err != nil {
		return err
	}

	*o = AddTaskRequest(varAddTaskRequest)

	return err
}

type NullableAddTaskRequest struct {
	value *AddTaskRequest
	isSet bool
}

func (v NullableAddTaskRequest) Get() *AddTaskRequest {
	return v.value
}

func (v *NullableAddTaskRequest) Set(val *AddTaskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddTaskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddTaskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddTaskRequest(val *AddTaskRequest) *NullableAddTaskRequest {
	return &NullableAddTaskRequest{value: val, isSet: true}
}

func (v NullableAddTaskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddTaskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


