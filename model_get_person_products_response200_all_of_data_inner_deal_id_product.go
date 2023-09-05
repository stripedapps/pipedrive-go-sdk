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

// checks if the GetPersonProductsResponse200AllOfDataInnerDEALIDProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPersonProductsResponse200AllOfDataInnerDEALIDProduct{}

// GetPersonProductsResponse200AllOfDataInnerDEALIDProduct struct for GetPersonProductsResponse200AllOfDataInnerDEALIDProduct
type GetPersonProductsResponse200AllOfDataInnerDEALIDProduct struct {
	// The ID of the product
	Id *int32 `json:"id,omitempty"`
	// The ID of the company
	CompanyId *int32 `json:"company_id,omitempty"`
	// The name of the product
	Name *string `json:"name,omitempty"`
	// The product code
	Code *string `json:"code,omitempty"`
	// The description of the product
	Description *string `json:"description,omitempty"`
	// The unit in which this product is sold
	Unit *string `json:"unit,omitempty"`
	// The tax percentage
	Tax *float32 `json:"tax,omitempty"`
	// The category of the product
	Category *string `json:"category,omitempty"`
	// Whether this product will be made active or not
	ActiveFlag *bool `json:"active_flag,omitempty"`
	// Whether this product can be selected in deals or not
	Selectable *bool `json:"selectable,omitempty"`
	// The first letter of the product name
	FirstChar *string `json:"first_char,omitempty"`
	VisibleTo *string `json:"visible_to,omitempty"`
	// The ID of the user who will be marked as the owner of this product. When omitted, the authorized user ID will be used.
	OwnerId *int32 `json:"owner_id,omitempty"`
	// The count of files
	FilesCount *int32 `json:"files_count,omitempty"`
	// The date and time when the product was added to the deal
	AddTime *string `json:"add_time,omitempty"`
	// The date and time when the product was updated to the deal
	UpdateTime *string `json:"update_time,omitempty"`
	// The ID of the deal
	DealId *int32 `json:"deal_id,omitempty"`
}

// NewGetPersonProductsResponse200AllOfDataInnerDEALIDProduct instantiates a new GetPersonProductsResponse200AllOfDataInnerDEALIDProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPersonProductsResponse200AllOfDataInnerDEALIDProduct() *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct {
	this := GetPersonProductsResponse200AllOfDataInnerDEALIDProduct{}
	var tax float32 = 0
	this.Tax = &tax
	var activeFlag bool = true
	this.ActiveFlag = &activeFlag
	var selectable bool = true
	this.Selectable = &selectable
	return &this
}

// NewGetPersonProductsResponse200AllOfDataInnerDEALIDProductWithDefaults instantiates a new GetPersonProductsResponse200AllOfDataInnerDEALIDProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPersonProductsResponse200AllOfDataInnerDEALIDProductWithDefaults() *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct {
	this := GetPersonProductsResponse200AllOfDataInnerDEALIDProduct{}
	var tax float32 = 0
	this.Tax = &tax
	var activeFlag bool = true
	this.ActiveFlag = &activeFlag
	var selectable bool = true
	this.Selectable = &selectable
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) SetId(v int32) {
	o.Id = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetCompanyId() int32 {
	if o == nil || IsNil(o.CompanyId) {
		var ret int32
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetCompanyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CompanyId) {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) HasCompanyId() bool {
	if o != nil && !IsNil(o.CompanyId) {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given int32 and assigns it to the CompanyId field.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) SetCompanyId(v int32) {
	o.CompanyId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) SetDescription(v string) {
	o.Description = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) SetUnit(v string) {
	o.Unit = &v
}

// GetTax returns the Tax field value if set, zero value otherwise.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetTax() float32 {
	if o == nil || IsNil(o.Tax) {
		var ret float32
		return ret
	}
	return *o.Tax
}

// GetTaxOk returns a tuple with the Tax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetTaxOk() (*float32, bool) {
	if o == nil || IsNil(o.Tax) {
		return nil, false
	}
	return o.Tax, true
}

// HasTax returns a boolean if a field has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) HasTax() bool {
	if o != nil && !IsNil(o.Tax) {
		return true
	}

	return false
}

// SetTax gets a reference to the given float32 and assigns it to the Tax field.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) SetTax(v float32) {
	o.Tax = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) SetCategory(v string) {
	o.Category = &v
}

// GetActiveFlag returns the ActiveFlag field value if set, zero value otherwise.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetActiveFlag() bool {
	if o == nil || IsNil(o.ActiveFlag) {
		var ret bool
		return ret
	}
	return *o.ActiveFlag
}

// GetActiveFlagOk returns a tuple with the ActiveFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetActiveFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.ActiveFlag) {
		return nil, false
	}
	return o.ActiveFlag, true
}

// HasActiveFlag returns a boolean if a field has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) HasActiveFlag() bool {
	if o != nil && !IsNil(o.ActiveFlag) {
		return true
	}

	return false
}

// SetActiveFlag gets a reference to the given bool and assigns it to the ActiveFlag field.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) SetActiveFlag(v bool) {
	o.ActiveFlag = &v
}

// GetSelectable returns the Selectable field value if set, zero value otherwise.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetSelectable() bool {
	if o == nil || IsNil(o.Selectable) {
		var ret bool
		return ret
	}
	return *o.Selectable
}

// GetSelectableOk returns a tuple with the Selectable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetSelectableOk() (*bool, bool) {
	if o == nil || IsNil(o.Selectable) {
		return nil, false
	}
	return o.Selectable, true
}

// HasSelectable returns a boolean if a field has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) HasSelectable() bool {
	if o != nil && !IsNil(o.Selectable) {
		return true
	}

	return false
}

// SetSelectable gets a reference to the given bool and assigns it to the Selectable field.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) SetSelectable(v bool) {
	o.Selectable = &v
}

// GetFirstChar returns the FirstChar field value if set, zero value otherwise.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetFirstChar() string {
	if o == nil || IsNil(o.FirstChar) {
		var ret string
		return ret
	}
	return *o.FirstChar
}

// GetFirstCharOk returns a tuple with the FirstChar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetFirstCharOk() (*string, bool) {
	if o == nil || IsNil(o.FirstChar) {
		return nil, false
	}
	return o.FirstChar, true
}

// HasFirstChar returns a boolean if a field has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) HasFirstChar() bool {
	if o != nil && !IsNil(o.FirstChar) {
		return true
	}

	return false
}

// SetFirstChar gets a reference to the given string and assigns it to the FirstChar field.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) SetFirstChar(v string) {
	o.FirstChar = &v
}

// GetVisibleTo returns the VisibleTo field value if set, zero value otherwise.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetVisibleTo() string {
	if o == nil || IsNil(o.VisibleTo) {
		var ret string
		return ret
	}
	return *o.VisibleTo
}

// GetVisibleToOk returns a tuple with the VisibleTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetVisibleToOk() (*string, bool) {
	if o == nil || IsNil(o.VisibleTo) {
		return nil, false
	}
	return o.VisibleTo, true
}

// HasVisibleTo returns a boolean if a field has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) HasVisibleTo() bool {
	if o != nil && !IsNil(o.VisibleTo) {
		return true
	}

	return false
}

// SetVisibleTo gets a reference to the given string and assigns it to the VisibleTo field.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) SetVisibleTo(v string) {
	o.VisibleTo = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetOwnerId() int32 {
	if o == nil || IsNil(o.OwnerId) {
		var ret int32
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetOwnerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given int32 and assigns it to the OwnerId field.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) SetOwnerId(v int32) {
	o.OwnerId = &v
}

// GetFilesCount returns the FilesCount field value if set, zero value otherwise.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetFilesCount() int32 {
	if o == nil || IsNil(o.FilesCount) {
		var ret int32
		return ret
	}
	return *o.FilesCount
}

// GetFilesCountOk returns a tuple with the FilesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetFilesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.FilesCount) {
		return nil, false
	}
	return o.FilesCount, true
}

// HasFilesCount returns a boolean if a field has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) HasFilesCount() bool {
	if o != nil && !IsNil(o.FilesCount) {
		return true
	}

	return false
}

// SetFilesCount gets a reference to the given int32 and assigns it to the FilesCount field.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) SetFilesCount(v int32) {
	o.FilesCount = &v
}

// GetAddTime returns the AddTime field value if set, zero value otherwise.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetAddTime() string {
	if o == nil || IsNil(o.AddTime) {
		var ret string
		return ret
	}
	return *o.AddTime
}

// GetAddTimeOk returns a tuple with the AddTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetAddTimeOk() (*string, bool) {
	if o == nil || IsNil(o.AddTime) {
		return nil, false
	}
	return o.AddTime, true
}

// HasAddTime returns a boolean if a field has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) HasAddTime() bool {
	if o != nil && !IsNil(o.AddTime) {
		return true
	}

	return false
}

// SetAddTime gets a reference to the given string and assigns it to the AddTime field.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) SetAddTime(v string) {
	o.AddTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetUpdateTime() string {
	if o == nil || IsNil(o.UpdateTime) {
		var ret string
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetUpdateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given string and assigns it to the UpdateTime field.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) SetUpdateTime(v string) {
	o.UpdateTime = &v
}

// GetDealId returns the DealId field value if set, zero value otherwise.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetDealId() int32 {
	if o == nil || IsNil(o.DealId) {
		var ret int32
		return ret
	}
	return *o.DealId
}

// GetDealIdOk returns a tuple with the DealId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) GetDealIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DealId) {
		return nil, false
	}
	return o.DealId, true
}

// HasDealId returns a boolean if a field has been set.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) HasDealId() bool {
	if o != nil && !IsNil(o.DealId) {
		return true
	}

	return false
}

// SetDealId gets a reference to the given int32 and assigns it to the DealId field.
func (o *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) SetDealId(v int32) {
	o.DealId = &v
}

func (o GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CompanyId) {
		toSerialize["company_id"] = o.CompanyId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !IsNil(o.Tax) {
		toSerialize["tax"] = o.Tax
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.ActiveFlag) {
		toSerialize["active_flag"] = o.ActiveFlag
	}
	if !IsNil(o.Selectable) {
		toSerialize["selectable"] = o.Selectable
	}
	if !IsNil(o.FirstChar) {
		toSerialize["first_char"] = o.FirstChar
	}
	if !IsNil(o.VisibleTo) {
		toSerialize["visible_to"] = o.VisibleTo
	}
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	if !IsNil(o.FilesCount) {
		toSerialize["files_count"] = o.FilesCount
	}
	if !IsNil(o.AddTime) {
		toSerialize["add_time"] = o.AddTime
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["update_time"] = o.UpdateTime
	}
	if !IsNil(o.DealId) {
		toSerialize["deal_id"] = o.DealId
	}
	return toSerialize, nil
}

type NullableGetPersonProductsResponse200AllOfDataInnerDEALIDProduct struct {
	value *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct
	isSet bool
}

func (v NullableGetPersonProductsResponse200AllOfDataInnerDEALIDProduct) Get() *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct {
	return v.value
}

func (v *NullableGetPersonProductsResponse200AllOfDataInnerDEALIDProduct) Set(val *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPersonProductsResponse200AllOfDataInnerDEALIDProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPersonProductsResponse200AllOfDataInnerDEALIDProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPersonProductsResponse200AllOfDataInnerDEALIDProduct(val *GetPersonProductsResponse200AllOfDataInnerDEALIDProduct) *NullableGetPersonProductsResponse200AllOfDataInnerDEALIDProduct {
	return &NullableGetPersonProductsResponse200AllOfDataInnerDEALIDProduct{value: val, isSet: true}
}

func (v NullableGetPersonProductsResponse200AllOfDataInnerDEALIDProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPersonProductsResponse200AllOfDataInnerDEALIDProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


