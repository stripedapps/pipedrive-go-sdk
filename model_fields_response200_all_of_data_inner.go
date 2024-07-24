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
)

// checks if the FieldsResponse200AllOfDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FieldsResponse200AllOfDataInner{}

// FieldsResponse200AllOfDataInner struct for FieldsResponse200AllOfDataInner
type FieldsResponse200AllOfDataInner struct {
	// The ID of the field. Value is `null` in case of subfields.
	Id *int32 `json:"id,omitempty"`
	// The key of the field. For custom fields this is generated upon creation.
	Key *string `json:"key,omitempty"`
	// The name of the field
	Name *string `json:"name,omitempty"`
	// The order number of the field
	OrderNr *int32 `json:"order_nr,omitempty"`
	// The type of the field<table><tr><th>Value</th><th>Description</th></tr><tr><td>`address`</td><td>Address field</td></tr><tr><td>`date`</td><td>Date (format YYYY-MM-DD)</td></tr><tr><td>`daterange`</td><td>Date-range field (has a start date and end date value, both YYYY-MM-DD)</td></tr><tr><td>`double`</td><td>Numeric value</td></tr><tr><td>`enum`</td><td>Options field with a single possible chosen option</td></tr><tr></tr><tr><td>`monetary`</td><td>Monetary field (has a numeric value and a currency value)</td></tr><tr><td>`org`</td><td>Organization field (contains an organization ID which is stored on the same account)</td></tr><tr><td>`people`</td><td>Person field (contains a person ID which is stored on the same account)</td></tr><tr><td>`phone`</td><td>Phone field (up to 255 numbers and/or characters)</td></tr><tr><td>`set`</td><td>Options field with a possibility of having multiple chosen options</td></tr><tr><td>`text`</td><td>Long text (up to 65k characters)</td></tr><tr><td>`time`</td><td>Time field (format HH:MM:SS)</td></tr><tr><td>`timerange`</td><td>Time-range field (has a start time and end time value, both HH:MM:SS)</td></tr><tr><td>`user`</td><td>User field (contains a user ID of another Pipedrive user)</td></tr><tr><td>`varchar`</td><td>Text (up to 255 characters)</td></tr><tr><td>`varchar_auto`</td><td>Autocomplete text (up to 255 characters)</td></tr><tr><td>`visible_to`</td><td>System field that keeps item's visibility setting</td></tr></table>
	FieldType *string `json:"field_type,omitempty"`
	// The creation time of the field
	AddTime *time.Time `json:"add_time,omitempty"`
	// The update time of the field
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// The ID of the user who created or most recently updated the field, only applicable for custom fields
	LastUpdatedByUserId *int32 `json:"last_updated_by_user_id,omitempty"`
	// The ID of the user who created the field
	CreatedByUserId *int32 `json:"created_by_user_id,omitempty"`
	// The active flag of the field
	ActiveFlag *bool `json:"active_flag,omitempty"`
	// The edit flag of the field
	EditFlag *bool `json:"edit_flag,omitempty"`
	// Not used
	IndexVisibleFlag *bool `json:"index_visible_flag,omitempty"`
	// Not used
	DetailsVisibleFlag *bool `json:"details_visible_flag,omitempty"`
	// Not used
	AddVisibleFlag *bool `json:"add_visible_flag,omitempty"`
	// Not used
	ImportantFlag *bool `json:"important_flag,omitempty"`
	// Whether or not the field of an item can be edited in bulk
	BulkEditAllowed *bool `json:"bulk_edit_allowed,omitempty"`
	// Whether or not items can be searched by this field
	SearchableFlag *bool `json:"searchable_flag,omitempty"`
	// Whether or not items can be filtered by this field
	FilteringAllowed *bool `json:"filtering_allowed,omitempty"`
	// Whether or not items can be sorted by this field
	SortableFlag *bool `json:"sortable_flag,omitempty"`
	// Whether or not the field is mandatory
	MandatoryFlag *bool `json:"mandatory_flag,omitempty"`
	// The options of the field. When there are no options, `null` is returned.
	Options []map[string]interface{} `json:"options,omitempty"`
	// The deleted options of the field. Only present when there is at least 1 deleted option.
	OptionsDeleted []map[string]interface{} `json:"options_deleted,omitempty"`
	// Whether or not the field is a subfield of another field. Only present if field is subfield.
	IsSubfield *bool `json:"is_subfield,omitempty"`
	// The subfields of the field. Only present when the field has subfields.
	Subfields []map[string]interface{} `json:"subfields,omitempty"`
}

// NewFieldsResponse200AllOfDataInner instantiates a new FieldsResponse200AllOfDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFieldsResponse200AllOfDataInner() *FieldsResponse200AllOfDataInner {
	this := FieldsResponse200AllOfDataInner{}
	return &this
}

// NewFieldsResponse200AllOfDataInnerWithDefaults instantiates a new FieldsResponse200AllOfDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFieldsResponse200AllOfDataInnerWithDefaults() *FieldsResponse200AllOfDataInner {
	this := FieldsResponse200AllOfDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfDataInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfDataInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *FieldsResponse200AllOfDataInner) SetId(v int32) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfDataInner) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfDataInner) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *FieldsResponse200AllOfDataInner) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfDataInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfDataInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FieldsResponse200AllOfDataInner) SetName(v string) {
	o.Name = &v
}

// GetOrderNr returns the OrderNr field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfDataInner) GetOrderNr() int32 {
	if o == nil || IsNil(o.OrderNr) {
		var ret int32
		return ret
	}
	return *o.OrderNr
}

// GetOrderNrOk returns a tuple with the OrderNr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfDataInner) GetOrderNrOk() (*int32, bool) {
	if o == nil || IsNil(o.OrderNr) {
		return nil, false
	}
	return o.OrderNr, true
}

// HasOrderNr returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasOrderNr() bool {
	if o != nil && !IsNil(o.OrderNr) {
		return true
	}

	return false
}

// SetOrderNr gets a reference to the given int32 and assigns it to the OrderNr field.
func (o *FieldsResponse200AllOfDataInner) SetOrderNr(v int32) {
	o.OrderNr = &v
}

// GetFieldType returns the FieldType field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfDataInner) GetFieldType() string {
	if o == nil || IsNil(o.FieldType) {
		var ret string
		return ret
	}
	return *o.FieldType
}

// GetFieldTypeOk returns a tuple with the FieldType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfDataInner) GetFieldTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FieldType) {
		return nil, false
	}
	return o.FieldType, true
}

// HasFieldType returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasFieldType() bool {
	if o != nil && !IsNil(o.FieldType) {
		return true
	}

	return false
}

// SetFieldType gets a reference to the given string and assigns it to the FieldType field.
func (o *FieldsResponse200AllOfDataInner) SetFieldType(v string) {
	o.FieldType = &v
}

// GetAddTime returns the AddTime field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfDataInner) GetAddTime() time.Time {
	if o == nil || IsNil(o.AddTime) {
		var ret time.Time
		return ret
	}
	return *o.AddTime
}

// GetAddTimeOk returns a tuple with the AddTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfDataInner) GetAddTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AddTime) {
		return nil, false
	}
	return o.AddTime, true
}

// HasAddTime returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasAddTime() bool {
	if o != nil && !IsNil(o.AddTime) {
		return true
	}

	return false
}

// SetAddTime gets a reference to the given time.Time and assigns it to the AddTime field.
func (o *FieldsResponse200AllOfDataInner) SetAddTime(v time.Time) {
	o.AddTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfDataInner) GetUpdateTime() time.Time {
	if o == nil || IsNil(o.UpdateTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfDataInner) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given time.Time and assigns it to the UpdateTime field.
func (o *FieldsResponse200AllOfDataInner) SetUpdateTime(v time.Time) {
	o.UpdateTime = &v
}

// GetLastUpdatedByUserId returns the LastUpdatedByUserId field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfDataInner) GetLastUpdatedByUserId() int32 {
	if o == nil || IsNil(o.LastUpdatedByUserId) {
		var ret int32
		return ret
	}
	return *o.LastUpdatedByUserId
}

// GetLastUpdatedByUserIdOk returns a tuple with the LastUpdatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfDataInner) GetLastUpdatedByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LastUpdatedByUserId) {
		return nil, false
	}
	return o.LastUpdatedByUserId, true
}

// HasLastUpdatedByUserId returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasLastUpdatedByUserId() bool {
	if o != nil && !IsNil(o.LastUpdatedByUserId) {
		return true
	}

	return false
}

// SetLastUpdatedByUserId gets a reference to the given int32 and assigns it to the LastUpdatedByUserId field.
func (o *FieldsResponse200AllOfDataInner) SetLastUpdatedByUserId(v int32) {
	o.LastUpdatedByUserId = &v
}

// GetCreatedByUserId returns the CreatedByUserId field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfDataInner) GetCreatedByUserId() int32 {
	if o == nil || IsNil(o.CreatedByUserId) {
		var ret int32
		return ret
	}
	return *o.CreatedByUserId
}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfDataInner) GetCreatedByUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedByUserId) {
		return nil, false
	}
	return o.CreatedByUserId, true
}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasCreatedByUserId() bool {
	if o != nil && !IsNil(o.CreatedByUserId) {
		return true
	}

	return false
}

// SetCreatedByUserId gets a reference to the given int32 and assigns it to the CreatedByUserId field.
func (o *FieldsResponse200AllOfDataInner) SetCreatedByUserId(v int32) {
	o.CreatedByUserId = &v
}

// GetActiveFlag returns the ActiveFlag field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfDataInner) GetActiveFlag() bool {
	if o == nil || IsNil(o.ActiveFlag) {
		var ret bool
		return ret
	}
	return *o.ActiveFlag
}

// GetActiveFlagOk returns a tuple with the ActiveFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfDataInner) GetActiveFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.ActiveFlag) {
		return nil, false
	}
	return o.ActiveFlag, true
}

// HasActiveFlag returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasActiveFlag() bool {
	if o != nil && !IsNil(o.ActiveFlag) {
		return true
	}

	return false
}

// SetActiveFlag gets a reference to the given bool and assigns it to the ActiveFlag field.
func (o *FieldsResponse200AllOfDataInner) SetActiveFlag(v bool) {
	o.ActiveFlag = &v
}

// GetEditFlag returns the EditFlag field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfDataInner) GetEditFlag() bool {
	if o == nil || IsNil(o.EditFlag) {
		var ret bool
		return ret
	}
	return *o.EditFlag
}

// GetEditFlagOk returns a tuple with the EditFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfDataInner) GetEditFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.EditFlag) {
		return nil, false
	}
	return o.EditFlag, true
}

// HasEditFlag returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasEditFlag() bool {
	if o != nil && !IsNil(o.EditFlag) {
		return true
	}

	return false
}

// SetEditFlag gets a reference to the given bool and assigns it to the EditFlag field.
func (o *FieldsResponse200AllOfDataInner) SetEditFlag(v bool) {
	o.EditFlag = &v
}

// GetIndexVisibleFlag returns the IndexVisibleFlag field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfDataInner) GetIndexVisibleFlag() bool {
	if o == nil || IsNil(o.IndexVisibleFlag) {
		var ret bool
		return ret
	}
	return *o.IndexVisibleFlag
}

// GetIndexVisibleFlagOk returns a tuple with the IndexVisibleFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfDataInner) GetIndexVisibleFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.IndexVisibleFlag) {
		return nil, false
	}
	return o.IndexVisibleFlag, true
}

// HasIndexVisibleFlag returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasIndexVisibleFlag() bool {
	if o != nil && !IsNil(o.IndexVisibleFlag) {
		return true
	}

	return false
}

// SetIndexVisibleFlag gets a reference to the given bool and assigns it to the IndexVisibleFlag field.
func (o *FieldsResponse200AllOfDataInner) SetIndexVisibleFlag(v bool) {
	o.IndexVisibleFlag = &v
}

// GetDetailsVisibleFlag returns the DetailsVisibleFlag field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfDataInner) GetDetailsVisibleFlag() bool {
	if o == nil || IsNil(o.DetailsVisibleFlag) {
		var ret bool
		return ret
	}
	return *o.DetailsVisibleFlag
}

// GetDetailsVisibleFlagOk returns a tuple with the DetailsVisibleFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfDataInner) GetDetailsVisibleFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.DetailsVisibleFlag) {
		return nil, false
	}
	return o.DetailsVisibleFlag, true
}

// HasDetailsVisibleFlag returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasDetailsVisibleFlag() bool {
	if o != nil && !IsNil(o.DetailsVisibleFlag) {
		return true
	}

	return false
}

// SetDetailsVisibleFlag gets a reference to the given bool and assigns it to the DetailsVisibleFlag field.
func (o *FieldsResponse200AllOfDataInner) SetDetailsVisibleFlag(v bool) {
	o.DetailsVisibleFlag = &v
}

// GetAddVisibleFlag returns the AddVisibleFlag field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfDataInner) GetAddVisibleFlag() bool {
	if o == nil || IsNil(o.AddVisibleFlag) {
		var ret bool
		return ret
	}
	return *o.AddVisibleFlag
}

// GetAddVisibleFlagOk returns a tuple with the AddVisibleFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfDataInner) GetAddVisibleFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.AddVisibleFlag) {
		return nil, false
	}
	return o.AddVisibleFlag, true
}

// HasAddVisibleFlag returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasAddVisibleFlag() bool {
	if o != nil && !IsNil(o.AddVisibleFlag) {
		return true
	}

	return false
}

// SetAddVisibleFlag gets a reference to the given bool and assigns it to the AddVisibleFlag field.
func (o *FieldsResponse200AllOfDataInner) SetAddVisibleFlag(v bool) {
	o.AddVisibleFlag = &v
}

// GetImportantFlag returns the ImportantFlag field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfDataInner) GetImportantFlag() bool {
	if o == nil || IsNil(o.ImportantFlag) {
		var ret bool
		return ret
	}
	return *o.ImportantFlag
}

// GetImportantFlagOk returns a tuple with the ImportantFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfDataInner) GetImportantFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.ImportantFlag) {
		return nil, false
	}
	return o.ImportantFlag, true
}

// HasImportantFlag returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasImportantFlag() bool {
	if o != nil && !IsNil(o.ImportantFlag) {
		return true
	}

	return false
}

// SetImportantFlag gets a reference to the given bool and assigns it to the ImportantFlag field.
func (o *FieldsResponse200AllOfDataInner) SetImportantFlag(v bool) {
	o.ImportantFlag = &v
}

// GetBulkEditAllowed returns the BulkEditAllowed field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfDataInner) GetBulkEditAllowed() bool {
	if o == nil || IsNil(o.BulkEditAllowed) {
		var ret bool
		return ret
	}
	return *o.BulkEditAllowed
}

// GetBulkEditAllowedOk returns a tuple with the BulkEditAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfDataInner) GetBulkEditAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.BulkEditAllowed) {
		return nil, false
	}
	return o.BulkEditAllowed, true
}

// HasBulkEditAllowed returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasBulkEditAllowed() bool {
	if o != nil && !IsNil(o.BulkEditAllowed) {
		return true
	}

	return false
}

// SetBulkEditAllowed gets a reference to the given bool and assigns it to the BulkEditAllowed field.
func (o *FieldsResponse200AllOfDataInner) SetBulkEditAllowed(v bool) {
	o.BulkEditAllowed = &v
}

// GetSearchableFlag returns the SearchableFlag field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfDataInner) GetSearchableFlag() bool {
	if o == nil || IsNil(o.SearchableFlag) {
		var ret bool
		return ret
	}
	return *o.SearchableFlag
}

// GetSearchableFlagOk returns a tuple with the SearchableFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfDataInner) GetSearchableFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.SearchableFlag) {
		return nil, false
	}
	return o.SearchableFlag, true
}

// HasSearchableFlag returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasSearchableFlag() bool {
	if o != nil && !IsNil(o.SearchableFlag) {
		return true
	}

	return false
}

// SetSearchableFlag gets a reference to the given bool and assigns it to the SearchableFlag field.
func (o *FieldsResponse200AllOfDataInner) SetSearchableFlag(v bool) {
	o.SearchableFlag = &v
}

// GetFilteringAllowed returns the FilteringAllowed field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfDataInner) GetFilteringAllowed() bool {
	if o == nil || IsNil(o.FilteringAllowed) {
		var ret bool
		return ret
	}
	return *o.FilteringAllowed
}

// GetFilteringAllowedOk returns a tuple with the FilteringAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfDataInner) GetFilteringAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.FilteringAllowed) {
		return nil, false
	}
	return o.FilteringAllowed, true
}

// HasFilteringAllowed returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasFilteringAllowed() bool {
	if o != nil && !IsNil(o.FilteringAllowed) {
		return true
	}

	return false
}

// SetFilteringAllowed gets a reference to the given bool and assigns it to the FilteringAllowed field.
func (o *FieldsResponse200AllOfDataInner) SetFilteringAllowed(v bool) {
	o.FilteringAllowed = &v
}

// GetSortableFlag returns the SortableFlag field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfDataInner) GetSortableFlag() bool {
	if o == nil || IsNil(o.SortableFlag) {
		var ret bool
		return ret
	}
	return *o.SortableFlag
}

// GetSortableFlagOk returns a tuple with the SortableFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfDataInner) GetSortableFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.SortableFlag) {
		return nil, false
	}
	return o.SortableFlag, true
}

// HasSortableFlag returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasSortableFlag() bool {
	if o != nil && !IsNil(o.SortableFlag) {
		return true
	}

	return false
}

// SetSortableFlag gets a reference to the given bool and assigns it to the SortableFlag field.
func (o *FieldsResponse200AllOfDataInner) SetSortableFlag(v bool) {
	o.SortableFlag = &v
}

// GetMandatoryFlag returns the MandatoryFlag field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfDataInner) GetMandatoryFlag() bool {
	if o == nil || IsNil(o.MandatoryFlag) {
		var ret bool
		return ret
	}
	return *o.MandatoryFlag
}

// GetMandatoryFlagOk returns a tuple with the MandatoryFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfDataInner) GetMandatoryFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.MandatoryFlag) {
		return nil, false
	}
	return o.MandatoryFlag, true
}

// HasMandatoryFlag returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasMandatoryFlag() bool {
	if o != nil && !IsNil(o.MandatoryFlag) {
		return true
	}

	return false
}

// SetMandatoryFlag gets a reference to the given bool and assigns it to the MandatoryFlag field.
func (o *FieldsResponse200AllOfDataInner) SetMandatoryFlag(v bool) {
	o.MandatoryFlag = &v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FieldsResponse200AllOfDataInner) GetOptions() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FieldsResponse200AllOfDataInner) GetOptionsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []map[string]interface{} and assigns it to the Options field.
func (o *FieldsResponse200AllOfDataInner) SetOptions(v []map[string]interface{}) {
	o.Options = v
}

// GetOptionsDeleted returns the OptionsDeleted field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfDataInner) GetOptionsDeleted() []map[string]interface{} {
	if o == nil || IsNil(o.OptionsDeleted) {
		var ret []map[string]interface{}
		return ret
	}
	return o.OptionsDeleted
}

// GetOptionsDeletedOk returns a tuple with the OptionsDeleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfDataInner) GetOptionsDeletedOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.OptionsDeleted) {
		return nil, false
	}
	return o.OptionsDeleted, true
}

// HasOptionsDeleted returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasOptionsDeleted() bool {
	if o != nil && !IsNil(o.OptionsDeleted) {
		return true
	}

	return false
}

// SetOptionsDeleted gets a reference to the given []map[string]interface{} and assigns it to the OptionsDeleted field.
func (o *FieldsResponse200AllOfDataInner) SetOptionsDeleted(v []map[string]interface{}) {
	o.OptionsDeleted = v
}

// GetIsSubfield returns the IsSubfield field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfDataInner) GetIsSubfield() bool {
	if o == nil || IsNil(o.IsSubfield) {
		var ret bool
		return ret
	}
	return *o.IsSubfield
}

// GetIsSubfieldOk returns a tuple with the IsSubfield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfDataInner) GetIsSubfieldOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSubfield) {
		return nil, false
	}
	return o.IsSubfield, true
}

// HasIsSubfield returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasIsSubfield() bool {
	if o != nil && !IsNil(o.IsSubfield) {
		return true
	}

	return false
}

// SetIsSubfield gets a reference to the given bool and assigns it to the IsSubfield field.
func (o *FieldsResponse200AllOfDataInner) SetIsSubfield(v bool) {
	o.IsSubfield = &v
}

// GetSubfields returns the Subfields field value if set, zero value otherwise.
func (o *FieldsResponse200AllOfDataInner) GetSubfields() []map[string]interface{} {
	if o == nil || IsNil(o.Subfields) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Subfields
}

// GetSubfieldsOk returns a tuple with the Subfields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FieldsResponse200AllOfDataInner) GetSubfieldsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Subfields) {
		return nil, false
	}
	return o.Subfields, true
}

// HasSubfields returns a boolean if a field has been set.
func (o *FieldsResponse200AllOfDataInner) HasSubfields() bool {
	if o != nil && !IsNil(o.Subfields) {
		return true
	}

	return false
}

// SetSubfields gets a reference to the given []map[string]interface{} and assigns it to the Subfields field.
func (o *FieldsResponse200AllOfDataInner) SetSubfields(v []map[string]interface{}) {
	o.Subfields = v
}

func (o FieldsResponse200AllOfDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FieldsResponse200AllOfDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OrderNr) {
		toSerialize["order_nr"] = o.OrderNr
	}
	if !IsNil(o.FieldType) {
		toSerialize["field_type"] = o.FieldType
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
	if !IsNil(o.IndexVisibleFlag) {
		toSerialize["index_visible_flag"] = o.IndexVisibleFlag
	}
	if !IsNil(o.DetailsVisibleFlag) {
		toSerialize["details_visible_flag"] = o.DetailsVisibleFlag
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
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.OptionsDeleted) {
		toSerialize["options_deleted"] = o.OptionsDeleted
	}
	if !IsNil(o.IsSubfield) {
		toSerialize["is_subfield"] = o.IsSubfield
	}
	if !IsNil(o.Subfields) {
		toSerialize["subfields"] = o.Subfields
	}
	return toSerialize, nil
}

type NullableFieldsResponse200AllOfDataInner struct {
	value *FieldsResponse200AllOfDataInner
	isSet bool
}

func (v NullableFieldsResponse200AllOfDataInner) Get() *FieldsResponse200AllOfDataInner {
	return v.value
}

func (v *NullableFieldsResponse200AllOfDataInner) Set(val *FieldsResponse200AllOfDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableFieldsResponse200AllOfDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableFieldsResponse200AllOfDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFieldsResponse200AllOfDataInner(val *FieldsResponse200AllOfDataInner) *NullableFieldsResponse200AllOfDataInner {
	return &NullableFieldsResponse200AllOfDataInner{value: val, isSet: true}
}

func (v NullableFieldsResponse200AllOfDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFieldsResponse200AllOfDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


