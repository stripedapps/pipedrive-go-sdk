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

// checks if the UpdatePipelineResponse200AllOfData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePipelineResponse200AllOfData{}

// UpdatePipelineResponse200AllOfData The pipeline object
type UpdatePipelineResponse200AllOfData struct {
	// The ID of the pipeline
	Id *int32 `json:"id,omitempty"`
	// The name of the pipeline
	Name *string `json:"name,omitempty"`
	// The pipeline title displayed in the URL
	UrlTitle *string `json:"url_title,omitempty"`
	// Defines the order of pipelines. First order (`order_nr=0`) is the default pipeline.
	OrderNr *int32 `json:"order_nr,omitempty"`
	// Whether this pipeline will be made inactive (hidden) or active
	Active *bool `json:"active,omitempty"`
	// Whether deal probability is disabled or enabled for this pipeline
	DealProbability *bool `json:"deal_probability,omitempty"`
	// The pipeline creation time. Format: YYYY-MM-DD HH:MM:SS.
	AddTime *string `json:"add_time,omitempty"`
	// The pipeline update time. Format: YYYY-MM-DD HH:MM:SS.
	UpdateTime *string `json:"update_time,omitempty"`
	// A boolean that shows if the pipeline is selected from a filter or not
	Selected *bool `json:"selected,omitempty"`
}

// NewUpdatePipelineResponse200AllOfData instantiates a new UpdatePipelineResponse200AllOfData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePipelineResponse200AllOfData() *UpdatePipelineResponse200AllOfData {
	this := UpdatePipelineResponse200AllOfData{}
	return &this
}

// NewUpdatePipelineResponse200AllOfDataWithDefaults instantiates a new UpdatePipelineResponse200AllOfData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePipelineResponse200AllOfDataWithDefaults() *UpdatePipelineResponse200AllOfData {
	this := UpdatePipelineResponse200AllOfData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdatePipelineResponse200AllOfData) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePipelineResponse200AllOfData) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdatePipelineResponse200AllOfData) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UpdatePipelineResponse200AllOfData) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdatePipelineResponse200AllOfData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePipelineResponse200AllOfData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdatePipelineResponse200AllOfData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdatePipelineResponse200AllOfData) SetName(v string) {
	o.Name = &v
}

// GetUrlTitle returns the UrlTitle field value if set, zero value otherwise.
func (o *UpdatePipelineResponse200AllOfData) GetUrlTitle() string {
	if o == nil || IsNil(o.UrlTitle) {
		var ret string
		return ret
	}
	return *o.UrlTitle
}

// GetUrlTitleOk returns a tuple with the UrlTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePipelineResponse200AllOfData) GetUrlTitleOk() (*string, bool) {
	if o == nil || IsNil(o.UrlTitle) {
		return nil, false
	}
	return o.UrlTitle, true
}

// HasUrlTitle returns a boolean if a field has been set.
func (o *UpdatePipelineResponse200AllOfData) HasUrlTitle() bool {
	if o != nil && !IsNil(o.UrlTitle) {
		return true
	}

	return false
}

// SetUrlTitle gets a reference to the given string and assigns it to the UrlTitle field.
func (o *UpdatePipelineResponse200AllOfData) SetUrlTitle(v string) {
	o.UrlTitle = &v
}

// GetOrderNr returns the OrderNr field value if set, zero value otherwise.
func (o *UpdatePipelineResponse200AllOfData) GetOrderNr() int32 {
	if o == nil || IsNil(o.OrderNr) {
		var ret int32
		return ret
	}
	return *o.OrderNr
}

// GetOrderNrOk returns a tuple with the OrderNr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePipelineResponse200AllOfData) GetOrderNrOk() (*int32, bool) {
	if o == nil || IsNil(o.OrderNr) {
		return nil, false
	}
	return o.OrderNr, true
}

// HasOrderNr returns a boolean if a field has been set.
func (o *UpdatePipelineResponse200AllOfData) HasOrderNr() bool {
	if o != nil && !IsNil(o.OrderNr) {
		return true
	}

	return false
}

// SetOrderNr gets a reference to the given int32 and assigns it to the OrderNr field.
func (o *UpdatePipelineResponse200AllOfData) SetOrderNr(v int32) {
	o.OrderNr = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *UpdatePipelineResponse200AllOfData) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePipelineResponse200AllOfData) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *UpdatePipelineResponse200AllOfData) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *UpdatePipelineResponse200AllOfData) SetActive(v bool) {
	o.Active = &v
}

// GetDealProbability returns the DealProbability field value if set, zero value otherwise.
func (o *UpdatePipelineResponse200AllOfData) GetDealProbability() bool {
	if o == nil || IsNil(o.DealProbability) {
		var ret bool
		return ret
	}
	return *o.DealProbability
}

// GetDealProbabilityOk returns a tuple with the DealProbability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePipelineResponse200AllOfData) GetDealProbabilityOk() (*bool, bool) {
	if o == nil || IsNil(o.DealProbability) {
		return nil, false
	}
	return o.DealProbability, true
}

// HasDealProbability returns a boolean if a field has been set.
func (o *UpdatePipelineResponse200AllOfData) HasDealProbability() bool {
	if o != nil && !IsNil(o.DealProbability) {
		return true
	}

	return false
}

// SetDealProbability gets a reference to the given bool and assigns it to the DealProbability field.
func (o *UpdatePipelineResponse200AllOfData) SetDealProbability(v bool) {
	o.DealProbability = &v
}

// GetAddTime returns the AddTime field value if set, zero value otherwise.
func (o *UpdatePipelineResponse200AllOfData) GetAddTime() string {
	if o == nil || IsNil(o.AddTime) {
		var ret string
		return ret
	}
	return *o.AddTime
}

// GetAddTimeOk returns a tuple with the AddTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePipelineResponse200AllOfData) GetAddTimeOk() (*string, bool) {
	if o == nil || IsNil(o.AddTime) {
		return nil, false
	}
	return o.AddTime, true
}

// HasAddTime returns a boolean if a field has been set.
func (o *UpdatePipelineResponse200AllOfData) HasAddTime() bool {
	if o != nil && !IsNil(o.AddTime) {
		return true
	}

	return false
}

// SetAddTime gets a reference to the given string and assigns it to the AddTime field.
func (o *UpdatePipelineResponse200AllOfData) SetAddTime(v string) {
	o.AddTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *UpdatePipelineResponse200AllOfData) GetUpdateTime() string {
	if o == nil || IsNil(o.UpdateTime) {
		var ret string
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePipelineResponse200AllOfData) GetUpdateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *UpdatePipelineResponse200AllOfData) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given string and assigns it to the UpdateTime field.
func (o *UpdatePipelineResponse200AllOfData) SetUpdateTime(v string) {
	o.UpdateTime = &v
}

// GetSelected returns the Selected field value if set, zero value otherwise.
func (o *UpdatePipelineResponse200AllOfData) GetSelected() bool {
	if o == nil || IsNil(o.Selected) {
		var ret bool
		return ret
	}
	return *o.Selected
}

// GetSelectedOk returns a tuple with the Selected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePipelineResponse200AllOfData) GetSelectedOk() (*bool, bool) {
	if o == nil || IsNil(o.Selected) {
		return nil, false
	}
	return o.Selected, true
}

// HasSelected returns a boolean if a field has been set.
func (o *UpdatePipelineResponse200AllOfData) HasSelected() bool {
	if o != nil && !IsNil(o.Selected) {
		return true
	}

	return false
}

// SetSelected gets a reference to the given bool and assigns it to the Selected field.
func (o *UpdatePipelineResponse200AllOfData) SetSelected(v bool) {
	o.Selected = &v
}

func (o UpdatePipelineResponse200AllOfData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePipelineResponse200AllOfData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.UrlTitle) {
		toSerialize["url_title"] = o.UrlTitle
	}
	if !IsNil(o.OrderNr) {
		toSerialize["order_nr"] = o.OrderNr
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.DealProbability) {
		toSerialize["deal_probability"] = o.DealProbability
	}
	if !IsNil(o.AddTime) {
		toSerialize["add_time"] = o.AddTime
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["update_time"] = o.UpdateTime
	}
	if !IsNil(o.Selected) {
		toSerialize["selected"] = o.Selected
	}
	return toSerialize, nil
}

type NullableUpdatePipelineResponse200AllOfData struct {
	value *UpdatePipelineResponse200AllOfData
	isSet bool
}

func (v NullableUpdatePipelineResponse200AllOfData) Get() *UpdatePipelineResponse200AllOfData {
	return v.value
}

func (v *NullableUpdatePipelineResponse200AllOfData) Set(val *UpdatePipelineResponse200AllOfData) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePipelineResponse200AllOfData) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePipelineResponse200AllOfData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePipelineResponse200AllOfData(val *UpdatePipelineResponse200AllOfData) *NullableUpdatePipelineResponse200AllOfData {
	return &NullableUpdatePipelineResponse200AllOfData{value: val, isSet: true}
}

func (v NullableUpdatePipelineResponse200AllOfData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePipelineResponse200AllOfData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


