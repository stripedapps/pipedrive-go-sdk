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

// checks if the ProductField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductField{}

// ProductField struct for ProductField
type ProductField struct {
	// The name of the field
	Name string `json:"name"`
	// When `field_type` is either `set` or `enum`, possible options must be supplied as a JSON-encoded sequential array, for example:</br>`[{\"label\":\"red\"}, {\"label\":\"blue\"}, {\"label\":\"lilac\"}]`
	Options []map[string]interface{} `json:"options,omitempty"`
	// The type of the field<table><tr><th>Value</th><th>Description</th></tr><tr><td>`varchar`</td><td>Text (up to 255 characters)</td><tr><td>`varchar_auto`</td><td>Autocomplete text (up to 255 characters)</td><tr><td>`text`</td><td>Long text (up to 65k characters)</td><tr><td>`double`</td><td>Numeric value</td><tr><td>`monetary`</td><td>Monetary field (has a numeric value and a currency value)</td><tr><td>`date`</td><td>Date (format YYYY-MM-DD)</td><tr><td>`set`</td><td>Options field with a possibility of having multiple chosen options</td><tr><td>`enum`</td><td>Options field with a single possible chosen option</td><tr><td>`user`</td><td>User field (contains a user ID of another Pipedrive user)</td><tr><td>`org`</td><td>Organization field (contains an organization ID which is stored on the same account)</td><tr><td>`people`</td><td>Person field (contains a product ID which is stored on the same account)</td><tr><td>`phone`</td><td>Phone field (up to 255 numbers and/or characters)</td><tr><td>`time`</td><td>Time field (format HH:MM:SS)</td><tr><td>`timerange`</td><td>Time-range field (has a start time and end time value, both HH:MM:SS)</td><tr><td>`daterange`</td><td>Date-range field (has a start date and end date value, both YYYY-MM-DD)</td><tr><td>`address`</td><td>Address field</dd></table>
	FieldType string `json:"field_type"`
	// The ID of the product field
	Id *int32 `json:"id,omitempty"`
	// The key of the product field
	Key *string `json:"key,omitempty"`
	// The position (index) of the product field in the detail view
	OrderNr *int32 `json:"order_nr,omitempty"`
	// The product field creation time. Format: YYYY-MM-DD HH:MM:SS
	AddTime *string `json:"add_time,omitempty"`
	// The product field last update time. Format: YYYY-MM-DD HH:MM:SS
	UpdateTime *string `json:"update_time,omitempty"`
	// The ID of the last user to update the product field
	LastUpdatedByUserId *int32 `json:"last_updated_by_user_id,omitempty"`
	// The ID of the user who created the product field
	CreatedByUserId *int32 `json:"created_by_user_id,omitempty"`
	// Whether or not the product field is currently active
	ActiveFlag *bool `json:"active_flag,omitempty"`
	// Whether or not the product field name and metadata is editable
	EditFlag *bool `json:"edit_flag,omitempty"`
	// Whether or not the product field is visible in the Add Product Modal
	AddVisibleFlag *bool `json:"add_visible_flag,omitempty"`
	// Whether or not the product field is marked as important
	ImportantFlag *bool `json:"important_flag,omitempty"`
	// Whether or not the product field data can be edited
	BulkEditAllowed *bool `json:"bulk_edit_allowed,omitempty"`
	// Whether or not the product field is searchable
	SearchableFlag *bool `json:"searchable_flag,omitempty"`
	// Whether or not the product field value can be used when filtering searches
	FilteringAllowed *bool `json:"filtering_allowed,omitempty"`
	// Whether or not the product field is sortable
	SortableFlag *bool `json:"sortable_flag,omitempty"`
	// Whether or not the product field is mandatory when creating products
	MandatoryFlag *bool `json:"mandatory_flag,omitempty"`
}

type _ProductField ProductField

// NewProductField instantiates a new ProductField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductField(name string, fieldType string) *ProductField {
	this := ProductField{}
	this.Name = name
	this.FieldType = fieldType
	return &this
}

// NewProductFieldWithDefaults instantiates a new ProductField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductFieldWithDefaults() *ProductField {
	this := ProductField{}
	return &this
}

// GetName returns the Name field value
func (o *ProductField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProductField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProductField) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ProductField) GetOptions() []map[string]interface{} {
	if o == nil || IsNil(o.Options) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductField) GetOptionsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ProductField) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []map[string]interface{} and assigns it to the Options field.
func (o *ProductField) SetOptions(v []map[string]interface{}) {
	o.Options = v
}

// GetFieldType returns the FieldType field value
func (o *ProductField) GetFieldType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldType
}

// GetFieldTypeOk returns a tuple with the FieldType field value
// and a boolean to check if the value has been set.
func (o *ProductField) GetFieldTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldType, true
}

// SetFieldType sets field value
func (o *ProductField) SetFieldType(v string) {
	o.FieldType = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductField) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductField) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductField) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ProductField) SetId(v int32) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ProductField) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductField) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ProductField) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ProductField) SetKey(v string) {
	o.Key = &v
}

// GetOrderNr returns the OrderNr field value if set, zero value otherwise.
func (o *ProductField) GetOrderNr() int32 {
	if o == nil || IsNil(o.OrderNr) {
		var ret int32
		return ret
	}
	return *o.OrderNr
}

// GetOrderNrOk returns a tuple with the OrderNr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductField) GetOrderNrOk() (*int32, bool) {
	if o == nil || IsNil(o.OrderNr) {
		return nil, false
	}
	return o.OrderNr, true
}

// HasOrderNr returns a boolean if a field has been set.
func (o *ProductField) HasOrderNr() bool {
	if o != nil && !IsNil(o.OrderNr) {
		return true
	}

	return false
}

// SetOrderNr gets a reference to the given int32 and assigns it to the OrderNr field.
func (o *ProductField) SetOrderNr(v int32) {
	o.OrderNr = &v
}

// GetAddTime returns the AddTime field value if set, zero value otherwise.
func (o *ProductField) GetAddTime() string {
	if o == nil || IsNil(o.AddTime) {
		var ret string
		return ret
	}
	return *o.AddTime
}

// GetAddTimeOk returns a tuple with the AddTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductField) GetAddTimeOk() (*string, bool) {
	if o == nil || IsNil(o.AddTime) {
		return nil, false
	}
	return o.AddTime, true
}

// HasAddTime returns a boolean if a field has been set.
func (o *ProductField) HasAddTime() bool {
	if o != nil && !IsNil(o.AddTime) {
		return true
	}

	return false
}

// SetAddTime gets a reference to the given string and assigns it to the AddTime field.
func (o *ProductField) SetAddTime(v string) {
	o.AddTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *ProductField) GetUpdateTime() string {
	if o == nil || IsNil(o.UpdateTime) {
		var ret string
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductField) GetUpdateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *ProductField) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given string and assigns it to the UpdateTime field.
func (o *ProductField) SetUpdateTime(v string) {
	o.UpdateTime = &v
}

// GetLastUpdatedByUserId returns the LastUpdatedByUserId field value if set, zero value otherwise.
func (o *ProductField) GetLastUpdatedByUserId() int32 {
	if o == nil || IsNil(o.LastUpdatedByUserId) {
		var ret int32
		return ret
	}
	return *o.LastUpdatedByUserId
}

// GetLastUpdatedByUserIdOk returns a tuple with the LastUpdatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductField) GetLastUpdatedByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LastUpdatedByUserId) {
		return nil, false
	}
	return o.LastUpdatedByUserId, true
}

// HasLastUpdatedByUserId returns a boolean if a field has been set.
func (o *ProductField) HasLastUpdatedByUserId() bool {
	if o != nil && !IsNil(o.LastUpdatedByUserId) {
		return true
	}

	return false
}

// SetLastUpdatedByUserId gets a reference to the given int32 and assigns it to the LastUpdatedByUserId field.
func (o *ProductField) SetLastUpdatedByUserId(v int32) {
	o.LastUpdatedByUserId = &v
}

// GetCreatedByUserId returns the CreatedByUserId field value if set, zero value otherwise.
func (o *ProductField) GetCreatedByUserId() int32 {
	if o == nil || IsNil(o.CreatedByUserId) {
		var ret int32
		return ret
	}
	return *o.CreatedByUserId
}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductField) GetCreatedByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedByUserId) {
		return nil, false
	}
	return o.CreatedByUserId, true
}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *ProductField) HasCreatedByUserId() bool {
	if o != nil && !IsNil(o.CreatedByUserId) {
		return true
	}

	return false
}

// SetCreatedByUserId gets a reference to the given int32 and assigns it to the CreatedByUserId field.
func (o *ProductField) SetCreatedByUserId(v int32) {
	o.CreatedByUserId = &v
}

// GetActiveFlag returns the ActiveFlag field value if set, zero value otherwise.
func (o *ProductField) GetActiveFlag() bool {
	if o == nil || IsNil(o.ActiveFlag) {
		var ret bool
		return ret
	}
	return *o.ActiveFlag
}

// GetActiveFlagOk returns a tuple with the ActiveFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductField) GetActiveFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.ActiveFlag) {
		return nil, false
	}
	return o.ActiveFlag, true
}

// HasActiveFlag returns a boolean if a field has been set.
func (o *ProductField) HasActiveFlag() bool {
	if o != nil && !IsNil(o.ActiveFlag) {
		return true
	}

	return false
}

// SetActiveFlag gets a reference to the given bool and assigns it to the ActiveFlag field.
func (o *ProductField) SetActiveFlag(v bool) {
	o.ActiveFlag = &v
}

// GetEditFlag returns the EditFlag field value if set, zero value otherwise.
func (o *ProductField) GetEditFlag() bool {
	if o == nil || IsNil(o.EditFlag) {
		var ret bool
		return ret
	}
	return *o.EditFlag
}

// GetEditFlagOk returns a tuple with the EditFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductField) GetEditFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.EditFlag) {
		return nil, false
	}
	return o.EditFlag, true
}

// HasEditFlag returns a boolean if a field has been set.
func (o *ProductField) HasEditFlag() bool {
	if o != nil && !IsNil(o.EditFlag) {
		return true
	}

	return false
}

// SetEditFlag gets a reference to the given bool and assigns it to the EditFlag field.
func (o *ProductField) SetEditFlag(v bool) {
	o.EditFlag = &v
}

// GetAddVisibleFlag returns the AddVisibleFlag field value if set, zero value otherwise.
func (o *ProductField) GetAddVisibleFlag() bool {
	if o == nil || IsNil(o.AddVisibleFlag) {
		var ret bool
		return ret
	}
	return *o.AddVisibleFlag
}

// GetAddVisibleFlagOk returns a tuple with the AddVisibleFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductField) GetAddVisibleFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.AddVisibleFlag) {
		return nil, false
	}
	return o.AddVisibleFlag, true
}

// HasAddVisibleFlag returns a boolean if a field has been set.
func (o *ProductField) HasAddVisibleFlag() bool {
	if o != nil && !IsNil(o.AddVisibleFlag) {
		return true
	}

	return false
}

// SetAddVisibleFlag gets a reference to the given bool and assigns it to the AddVisibleFlag field.
func (o *ProductField) SetAddVisibleFlag(v bool) {
	o.AddVisibleFlag = &v
}

// GetImportantFlag returns the ImportantFlag field value if set, zero value otherwise.
func (o *ProductField) GetImportantFlag() bool {
	if o == nil || IsNil(o.ImportantFlag) {
		var ret bool
		return ret
	}
	return *o.ImportantFlag
}

// GetImportantFlagOk returns a tuple with the ImportantFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductField) GetImportantFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.ImportantFlag) {
		return nil, false
	}
	return o.ImportantFlag, true
}

// HasImportantFlag returns a boolean if a field has been set.
func (o *ProductField) HasImportantFlag() bool {
	if o != nil && !IsNil(o.ImportantFlag) {
		return true
	}

	return false
}

// SetImportantFlag gets a reference to the given bool and assigns it to the ImportantFlag field.
func (o *ProductField) SetImportantFlag(v bool) {
	o.ImportantFlag = &v
}

// GetBulkEditAllowed returns the BulkEditAllowed field value if set, zero value otherwise.
func (o *ProductField) GetBulkEditAllowed() bool {
	if o == nil || IsNil(o.BulkEditAllowed) {
		var ret bool
		return ret
	}
	return *o.BulkEditAllowed
}

// GetBulkEditAllowedOk returns a tuple with the BulkEditAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductField) GetBulkEditAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.BulkEditAllowed) {
		return nil, false
	}
	return o.BulkEditAllowed, true
}

// HasBulkEditAllowed returns a boolean if a field has been set.
func (o *ProductField) HasBulkEditAllowed() bool {
	if o != nil && !IsNil(o.BulkEditAllowed) {
		return true
	}

	return false
}

// SetBulkEditAllowed gets a reference to the given bool and assigns it to the BulkEditAllowed field.
func (o *ProductField) SetBulkEditAllowed(v bool) {
	o.BulkEditAllowed = &v
}

// GetSearchableFlag returns the SearchableFlag field value if set, zero value otherwise.
func (o *ProductField) GetSearchableFlag() bool {
	if o == nil || IsNil(o.SearchableFlag) {
		var ret bool
		return ret
	}
	return *o.SearchableFlag
}

// GetSearchableFlagOk returns a tuple with the SearchableFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductField) GetSearchableFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.SearchableFlag) {
		return nil, false
	}
	return o.SearchableFlag, true
}

// HasSearchableFlag returns a boolean if a field has been set.
func (o *ProductField) HasSearchableFlag() bool {
	if o != nil && !IsNil(o.SearchableFlag) {
		return true
	}

	return false
}

// SetSearchableFlag gets a reference to the given bool and assigns it to the SearchableFlag field.
func (o *ProductField) SetSearchableFlag(v bool) {
	o.SearchableFlag = &v
}

// GetFilteringAllowed returns the FilteringAllowed field value if set, zero value otherwise.
func (o *ProductField) GetFilteringAllowed() bool {
	if o == nil || IsNil(o.FilteringAllowed) {
		var ret bool
		return ret
	}
	return *o.FilteringAllowed
}

// GetFilteringAllowedOk returns a tuple with the FilteringAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductField) GetFilteringAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.FilteringAllowed) {
		return nil, false
	}
	return o.FilteringAllowed, true
}

// HasFilteringAllowed returns a boolean if a field has been set.
func (o *ProductField) HasFilteringAllowed() bool {
	if o != nil && !IsNil(o.FilteringAllowed) {
		return true
	}

	return false
}

// SetFilteringAllowed gets a reference to the given bool and assigns it to the FilteringAllowed field.
func (o *ProductField) SetFilteringAllowed(v bool) {
	o.FilteringAllowed = &v
}

// GetSortableFlag returns the SortableFlag field value if set, zero value otherwise.
func (o *ProductField) GetSortableFlag() bool {
	if o == nil || IsNil(o.SortableFlag) {
		var ret bool
		return ret
	}
	return *o.SortableFlag
}

// GetSortableFlagOk returns a tuple with the SortableFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductField) GetSortableFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.SortableFlag) {
		return nil, false
	}
	return o.SortableFlag, true
}

// HasSortableFlag returns a boolean if a field has been set.
func (o *ProductField) HasSortableFlag() bool {
	if o != nil && !IsNil(o.SortableFlag) {
		return true
	}

	return false
}

// SetSortableFlag gets a reference to the given bool and assigns it to the SortableFlag field.
func (o *ProductField) SetSortableFlag(v bool) {
	o.SortableFlag = &v
}

// GetMandatoryFlag returns the MandatoryFlag field value if set, zero value otherwise.
func (o *ProductField) GetMandatoryFlag() bool {
	if o == nil || IsNil(o.MandatoryFlag) {
		var ret bool
		return ret
	}
	return *o.MandatoryFlag
}

// GetMandatoryFlagOk returns a tuple with the MandatoryFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductField) GetMandatoryFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.MandatoryFlag) {
		return nil, false
	}
	return o.MandatoryFlag, true
}

// HasMandatoryFlag returns a boolean if a field has been set.
func (o *ProductField) HasMandatoryFlag() bool {
	if o != nil && !IsNil(o.MandatoryFlag) {
		return true
	}

	return false
}

// SetMandatoryFlag gets a reference to the given bool and assigns it to the MandatoryFlag field.
func (o *ProductField) SetMandatoryFlag(v bool) {
	o.MandatoryFlag = &v
}

func (o ProductField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	toSerialize["field_type"] = o.FieldType
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.OrderNr) {
		toSerialize["order_nr"] = o.OrderNr
	}
	if !IsNil(o.AddTime) {
		toSerialize["add_time"] = o.AddTime
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["update_time"] = o.UpdateTime
	}
	if !IsNil(o.LastUpdatedByUserId) {
		toSerialize["last_updated_by_user_id"] = o.LastUpdatedByUserId
	}
	if !IsNil(o.CreatedByUserId) {
		toSerialize["created_by_user_id"] = o.CreatedByUserId
	}
	if !IsNil(o.ActiveFlag) {
		toSerialize["active_flag"] = o.ActiveFlag
	}
	if !IsNil(o.EditFlag) {
		toSerialize["edit_flag"] = o.EditFlag
	}
	if !IsNil(o.AddVisibleFlag) {
		toSerialize["add_visible_flag"] = o.AddVisibleFlag
	}
	if !IsNil(o.ImportantFlag) {
		toSerialize["important_flag"] = o.ImportantFlag
	}
	if !IsNil(o.BulkEditAllowed) {
		toSerialize["bulk_edit_allowed"] = o.BulkEditAllowed
	}
	if !IsNil(o.SearchableFlag) {
		toSerialize["searchable_flag"] = o.SearchableFlag
	}
	if !IsNil(o.FilteringAllowed) {
		toSerialize["filtering_allowed"] = o.FilteringAllowed
	}
	if !IsNil(o.SortableFlag) {
		toSerialize["sortable_flag"] = o.SortableFlag
	}
	if !IsNil(o.MandatoryFlag) {
		toSerialize["mandatory_flag"] = o.MandatoryFlag
	}
	return toSerialize, nil
}

func (o *ProductField) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"field_type",
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

	varProductField := _ProductField{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProductField)

	if err != nil {
		return err
	}

	*o = ProductField(varProductField)

	return err
}

type NullableProductField struct {
	value *ProductField
	isSet bool
}

func (v NullableProductField) Get() *ProductField {
	return v.value
}

func (v *NullableProductField) Set(val *ProductField) {
	v.value = val
	v.isSet = true
}

func (v NullableProductField) IsSet() bool {
	return v.isSet
}

func (v *NullableProductField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductField(val *ProductField) *NullableProductField {
	return &NullableProductField{value: val, isSet: true}
}

func (v NullableProductField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


