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

// checks if the PipelineDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PipelineDetails{}

// PipelineDetails The pipeline object
type PipelineDetails struct {
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
	DealsSummary *PipelineDetailsAllOfDealsSummary `json:"deals_summary,omitempty"`
}

// NewPipelineDetails instantiates a new PipelineDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineDetails() *PipelineDetails {
	this := PipelineDetails{}
	return &this
}

// NewPipelineDetailsWithDefaults instantiates a new PipelineDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineDetailsWithDefaults() *PipelineDetails {
	this := PipelineDetails{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PipelineDetails) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineDetails) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PipelineDetails) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PipelineDetails) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PipelineDetails) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineDetails) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PipelineDetails) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PipelineDetails) SetName(v string) {
	o.Name = &v
}

// GetUrlTitle returns the UrlTitle field value if set, zero value otherwise.
func (o *PipelineDetails) GetUrlTitle() string {
	if o == nil || IsNil(o.UrlTitle) {
		var ret string
		return ret
	}
	return *o.UrlTitle
}

// GetUrlTitleOk returns a tuple with the UrlTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineDetails) GetUrlTitleOk() (*string, bool) {
	if o == nil || IsNil(o.UrlTitle) {
		return nil, false
	}
	return o.UrlTitle, true
}

// HasUrlTitle returns a boolean if a field has been set.
func (o *PipelineDetails) HasUrlTitle() bool {
	if o != nil && !IsNil(o.UrlTitle) {
		return true
	}

	return false
}

// SetUrlTitle gets a reference to the given string and assigns it to the UrlTitle field.
func (o *PipelineDetails) SetUrlTitle(v string) {
	o.UrlTitle = &v
}

// GetOrderNr returns the OrderNr field value if set, zero value otherwise.
func (o *PipelineDetails) GetOrderNr() int32 {
	if o == nil || IsNil(o.OrderNr) {
		var ret int32
		return ret
	}
	return *o.OrderNr
}

// GetOrderNrOk returns a tuple with the OrderNr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineDetails) GetOrderNrOk() (*int32, bool) {
	if o == nil || IsNil(o.OrderNr) {
		return nil, false
	}
	return o.OrderNr, true
}

// HasOrderNr returns a boolean if a field has been set.
func (o *PipelineDetails) HasOrderNr() bool {
	if o != nil && !IsNil(o.OrderNr) {
		return true
	}

	return false
}

// SetOrderNr gets a reference to the given int32 and assigns it to the OrderNr field.
func (o *PipelineDetails) SetOrderNr(v int32) {
	o.OrderNr = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *PipelineDetails) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineDetails) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *PipelineDetails) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *PipelineDetails) SetActive(v bool) {
	o.Active = &v
}

// GetDealProbability returns the DealProbability field value if set, zero value otherwise.
func (o *PipelineDetails) GetDealProbability() bool {
	if o == nil || IsNil(o.DealProbability) {
		var ret bool
		return ret
	}
	return *o.DealProbability
}

// GetDealProbabilityOk returns a tuple with the DealProbability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineDetails) GetDealProbabilityOk() (*bool, bool) {
	if o == nil || IsNil(o.DealProbability) {
		return nil, false
	}
	return o.DealProbability, true
}

// HasDealProbability returns a boolean if a field has been set.
func (o *PipelineDetails) HasDealProbability() bool {
	if o != nil && !IsNil(o.DealProbability) {
		return true
	}

	return false
}

// SetDealProbability gets a reference to the given bool and assigns it to the DealProbability field.
func (o *PipelineDetails) SetDealProbability(v bool) {
	o.DealProbability = &v
}

// GetAddTime returns the AddTime field value if set, zero value otherwise.
func (o *PipelineDetails) GetAddTime() string {
	if o == nil || IsNil(o.AddTime) {
		var ret string
		return ret
	}
	return *o.AddTime
}

// GetAddTimeOk returns a tuple with the AddTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineDetails) GetAddTimeOk() (*string, bool) {
	if o == nil || IsNil(o.AddTime) {
		return nil, false
	}
	return o.AddTime, true
}

// HasAddTime returns a boolean if a field has been set.
func (o *PipelineDetails) HasAddTime() bool {
	if o != nil && !IsNil(o.AddTime) {
		return true
	}

	return false
}

// SetAddTime gets a reference to the given string and assigns it to the AddTime field.
func (o *PipelineDetails) SetAddTime(v string) {
	o.AddTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *PipelineDetails) GetUpdateTime() string {
	if o == nil || IsNil(o.UpdateTime) {
		var ret string
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineDetails) GetUpdateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *PipelineDetails) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given string and assigns it to the UpdateTime field.
func (o *PipelineDetails) SetUpdateTime(v string) {
	o.UpdateTime = &v
}

// GetSelected returns the Selected field value if set, zero value otherwise.
func (o *PipelineDetails) GetSelected() bool {
	if o == nil || IsNil(o.Selected) {
		var ret bool
		return ret
	}
	return *o.Selected
}

// GetSelectedOk returns a tuple with the Selected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineDetails) GetSelectedOk() (*bool, bool) {
	if o == nil || IsNil(o.Selected) {
		return nil, false
	}
	return o.Selected, true
}

// HasSelected returns a boolean if a field has been set.
func (o *PipelineDetails) HasSelected() bool {
	if o != nil && !IsNil(o.Selected) {
		return true
	}

	return false
}

// SetSelected gets a reference to the given bool and assigns it to the Selected field.
func (o *PipelineDetails) SetSelected(v bool) {
	o.Selected = &v
}

// GetDealsSummary returns the DealsSummary field value if set, zero value otherwise.
func (o *PipelineDetails) GetDealsSummary() PipelineDetailsAllOfDealsSummary {
	if o == nil || IsNil(o.DealsSummary) {
		var ret PipelineDetailsAllOfDealsSummary
		return ret
	}
	return *o.DealsSummary
}

// GetDealsSummaryOk returns a tuple with the DealsSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineDetails) GetDealsSummaryOk() (*PipelineDetailsAllOfDealsSummary, bool) {
	if o == nil || IsNil(o.DealsSummary) {
		return nil, false
	}
	return o.DealsSummary, true
}

// HasDealsSummary returns a boolean if a field has been set.
func (o *PipelineDetails) HasDealsSummary() bool {
	if o != nil && !IsNil(o.DealsSummary) {
		return true
	}

	return false
}

// SetDealsSummary gets a reference to the given PipelineDetailsAllOfDealsSummary and assigns it to the DealsSummary field.
func (o *PipelineDetails) SetDealsSummary(v PipelineDetailsAllOfDealsSummary) {
	o.DealsSummary = &v
}

func (o PipelineDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PipelineDetails) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DealsSummary) {
		toSerialize["deals_summary"] = o.DealsSummary
	}
	return toSerialize, nil
}

type NullablePipelineDetails struct {
	value *PipelineDetails
	isSet bool
}

func (v NullablePipelineDetails) Get() *PipelineDetails {
	return v.value
}

func (v *NullablePipelineDetails) Set(val *PipelineDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineDetails(val *PipelineDetails) *NullablePipelineDetails {
	return &NullablePipelineDetails{value: val, isSet: true}
}

func (v NullablePipelineDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


