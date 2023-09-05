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

// checks if the OrganizationsCollectionResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationsCollectionResponseObject{}

// OrganizationsCollectionResponseObject struct for OrganizationsCollectionResponseObject
type OrganizationsCollectionResponseObject struct {
	// The full address of the organization
	Address *string `json:"address,omitempty"`
	// The sub-premise of the organization location
	AddressSubpremise *string `json:"address_subpremise,omitempty"`
	// The street number of the organization location
	AddressStreetNumber *string `json:"address_street_number,omitempty"`
	// The route of the organization location
	AddressRoute *string `json:"address_route,omitempty"`
	// The sub-locality of the organization location
	AddressSublocality *string `json:"address_sublocality,omitempty"`
	// The locality of the organization location
	AddressLocality *string `json:"address_locality,omitempty"`
	// The level 1 admin area of the organization location
	AddressAdminAreaLevel1 *string `json:"address_admin_area_level_1,omitempty"`
	// The level 2 admin area of the organization location
	AddressAdminAreaLevel2 *string `json:"address_admin_area_level_2,omitempty"`
	// The country of the organization location
	AddressCountry *string `json:"address_country,omitempty"`
	// The postal code of the organization location
	AddressPostalCode *string `json:"address_postal_code,omitempty"`
	// The formatted organization location
	AddressFormattedAddress *string `json:"address_formatted_address,omitempty"`
	// The ID of the organization
	Id *int32 `json:"id,omitempty"`
	// Whether the organization is active or not
	ActiveFlag *bool `json:"active_flag,omitempty"`
	// The ID of the owner
	OwnerId *int32 `json:"owner_id,omitempty"`
	// The name of the organization
	Name *string `json:"name,omitempty"`
	// The last updated date and time of the organization. Format: YYYY-MM-DD HH:MM:SS
	UpdateTime *string `json:"update_time,omitempty"`
	// The date and time this organization was deleted. Format: YYYY-MM-DD HH:MM:SS
	DeleteTime NullableString `json:"delete_time,omitempty"`
	// The date and time when the organization was added/created. Format: YYYY-MM-DD HH:MM:SS
	AddTime *string `json:"add_time,omitempty"`
	// The visibility group ID of who can see the organization
	VisibleTo *string `json:"visible_to,omitempty"`
	// The label assigned to the organization
	Label NullableInt32 `json:"label,omitempty"`
	// The BCC email associated with the organization
	CcEmail *string `json:"cc_email,omitempty"`
}

// NewOrganizationsCollectionResponseObject instantiates a new OrganizationsCollectionResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationsCollectionResponseObject() *OrganizationsCollectionResponseObject {
	this := OrganizationsCollectionResponseObject{}
	return &this
}

// NewOrganizationsCollectionResponseObjectWithDefaults instantiates a new OrganizationsCollectionResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationsCollectionResponseObjectWithDefaults() *OrganizationsCollectionResponseObject {
	this := OrganizationsCollectionResponseObject{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *OrganizationsCollectionResponseObject) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsCollectionResponseObject) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *OrganizationsCollectionResponseObject) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *OrganizationsCollectionResponseObject) SetAddress(v string) {
	o.Address = &v
}

// GetAddressSubpremise returns the AddressSubpremise field value if set, zero value otherwise.
func (o *OrganizationsCollectionResponseObject) GetAddressSubpremise() string {
	if o == nil || IsNil(o.AddressSubpremise) {
		var ret string
		return ret
	}
	return *o.AddressSubpremise
}

// GetAddressSubpremiseOk returns a tuple with the AddressSubpremise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsCollectionResponseObject) GetAddressSubpremiseOk() (*string, bool) {
	if o == nil || IsNil(o.AddressSubpremise) {
		return nil, false
	}
	return o.AddressSubpremise, true
}

// HasAddressSubpremise returns a boolean if a field has been set.
func (o *OrganizationsCollectionResponseObject) HasAddressSubpremise() bool {
	if o != nil && !IsNil(o.AddressSubpremise) {
		return true
	}

	return false
}

// SetAddressSubpremise gets a reference to the given string and assigns it to the AddressSubpremise field.
func (o *OrganizationsCollectionResponseObject) SetAddressSubpremise(v string) {
	o.AddressSubpremise = &v
}

// GetAddressStreetNumber returns the AddressStreetNumber field value if set, zero value otherwise.
func (o *OrganizationsCollectionResponseObject) GetAddressStreetNumber() string {
	if o == nil || IsNil(o.AddressStreetNumber) {
		var ret string
		return ret
	}
	return *o.AddressStreetNumber
}

// GetAddressStreetNumberOk returns a tuple with the AddressStreetNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsCollectionResponseObject) GetAddressStreetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AddressStreetNumber) {
		return nil, false
	}
	return o.AddressStreetNumber, true
}

// HasAddressStreetNumber returns a boolean if a field has been set.
func (o *OrganizationsCollectionResponseObject) HasAddressStreetNumber() bool {
	if o != nil && !IsNil(o.AddressStreetNumber) {
		return true
	}

	return false
}

// SetAddressStreetNumber gets a reference to the given string and assigns it to the AddressStreetNumber field.
func (o *OrganizationsCollectionResponseObject) SetAddressStreetNumber(v string) {
	o.AddressStreetNumber = &v
}

// GetAddressRoute returns the AddressRoute field value if set, zero value otherwise.
func (o *OrganizationsCollectionResponseObject) GetAddressRoute() string {
	if o == nil || IsNil(o.AddressRoute) {
		var ret string
		return ret
	}
	return *o.AddressRoute
}

// GetAddressRouteOk returns a tuple with the AddressRoute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsCollectionResponseObject) GetAddressRouteOk() (*string, bool) {
	if o == nil || IsNil(o.AddressRoute) {
		return nil, false
	}
	return o.AddressRoute, true
}

// HasAddressRoute returns a boolean if a field has been set.
func (o *OrganizationsCollectionResponseObject) HasAddressRoute() bool {
	if o != nil && !IsNil(o.AddressRoute) {
		return true
	}

	return false
}

// SetAddressRoute gets a reference to the given string and assigns it to the AddressRoute field.
func (o *OrganizationsCollectionResponseObject) SetAddressRoute(v string) {
	o.AddressRoute = &v
}

// GetAddressSublocality returns the AddressSublocality field value if set, zero value otherwise.
func (o *OrganizationsCollectionResponseObject) GetAddressSublocality() string {
	if o == nil || IsNil(o.AddressSublocality) {
		var ret string
		return ret
	}
	return *o.AddressSublocality
}

// GetAddressSublocalityOk returns a tuple with the AddressSublocality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsCollectionResponseObject) GetAddressSublocalityOk() (*string, bool) {
	if o == nil || IsNil(o.AddressSublocality) {
		return nil, false
	}
	return o.AddressSublocality, true
}

// HasAddressSublocality returns a boolean if a field has been set.
func (o *OrganizationsCollectionResponseObject) HasAddressSublocality() bool {
	if o != nil && !IsNil(o.AddressSublocality) {
		return true
	}

	return false
}

// SetAddressSublocality gets a reference to the given string and assigns it to the AddressSublocality field.
func (o *OrganizationsCollectionResponseObject) SetAddressSublocality(v string) {
	o.AddressSublocality = &v
}

// GetAddressLocality returns the AddressLocality field value if set, zero value otherwise.
func (o *OrganizationsCollectionResponseObject) GetAddressLocality() string {
	if o == nil || IsNil(o.AddressLocality) {
		var ret string
		return ret
	}
	return *o.AddressLocality
}

// GetAddressLocalityOk returns a tuple with the AddressLocality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsCollectionResponseObject) GetAddressLocalityOk() (*string, bool) {
	if o == nil || IsNil(o.AddressLocality) {
		return nil, false
	}
	return o.AddressLocality, true
}

// HasAddressLocality returns a boolean if a field has been set.
func (o *OrganizationsCollectionResponseObject) HasAddressLocality() bool {
	if o != nil && !IsNil(o.AddressLocality) {
		return true
	}

	return false
}

// SetAddressLocality gets a reference to the given string and assigns it to the AddressLocality field.
func (o *OrganizationsCollectionResponseObject) SetAddressLocality(v string) {
	o.AddressLocality = &v
}

// GetAddressAdminAreaLevel1 returns the AddressAdminAreaLevel1 field value if set, zero value otherwise.
func (o *OrganizationsCollectionResponseObject) GetAddressAdminAreaLevel1() string {
	if o == nil || IsNil(o.AddressAdminAreaLevel1) {
		var ret string
		return ret
	}
	return *o.AddressAdminAreaLevel1
}

// GetAddressAdminAreaLevel1Ok returns a tuple with the AddressAdminAreaLevel1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsCollectionResponseObject) GetAddressAdminAreaLevel1Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressAdminAreaLevel1) {
		return nil, false
	}
	return o.AddressAdminAreaLevel1, true
}

// HasAddressAdminAreaLevel1 returns a boolean if a field has been set.
func (o *OrganizationsCollectionResponseObject) HasAddressAdminAreaLevel1() bool {
	if o != nil && !IsNil(o.AddressAdminAreaLevel1) {
		return true
	}

	return false
}

// SetAddressAdminAreaLevel1 gets a reference to the given string and assigns it to the AddressAdminAreaLevel1 field.
func (o *OrganizationsCollectionResponseObject) SetAddressAdminAreaLevel1(v string) {
	o.AddressAdminAreaLevel1 = &v
}

// GetAddressAdminAreaLevel2 returns the AddressAdminAreaLevel2 field value if set, zero value otherwise.
func (o *OrganizationsCollectionResponseObject) GetAddressAdminAreaLevel2() string {
	if o == nil || IsNil(o.AddressAdminAreaLevel2) {
		var ret string
		return ret
	}
	return *o.AddressAdminAreaLevel2
}

// GetAddressAdminAreaLevel2Ok returns a tuple with the AddressAdminAreaLevel2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsCollectionResponseObject) GetAddressAdminAreaLevel2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressAdminAreaLevel2) {
		return nil, false
	}
	return o.AddressAdminAreaLevel2, true
}

// HasAddressAdminAreaLevel2 returns a boolean if a field has been set.
func (o *OrganizationsCollectionResponseObject) HasAddressAdminAreaLevel2() bool {
	if o != nil && !IsNil(o.AddressAdminAreaLevel2) {
		return true
	}

	return false
}

// SetAddressAdminAreaLevel2 gets a reference to the given string and assigns it to the AddressAdminAreaLevel2 field.
func (o *OrganizationsCollectionResponseObject) SetAddressAdminAreaLevel2(v string) {
	o.AddressAdminAreaLevel2 = &v
}

// GetAddressCountry returns the AddressCountry field value if set, zero value otherwise.
func (o *OrganizationsCollectionResponseObject) GetAddressCountry() string {
	if o == nil || IsNil(o.AddressCountry) {
		var ret string
		return ret
	}
	return *o.AddressCountry
}

// GetAddressCountryOk returns a tuple with the AddressCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsCollectionResponseObject) GetAddressCountryOk() (*string, bool) {
	if o == nil || IsNil(o.AddressCountry) {
		return nil, false
	}
	return o.AddressCountry, true
}

// HasAddressCountry returns a boolean if a field has been set.
func (o *OrganizationsCollectionResponseObject) HasAddressCountry() bool {
	if o != nil && !IsNil(o.AddressCountry) {
		return true
	}

	return false
}

// SetAddressCountry gets a reference to the given string and assigns it to the AddressCountry field.
func (o *OrganizationsCollectionResponseObject) SetAddressCountry(v string) {
	o.AddressCountry = &v
}

// GetAddressPostalCode returns the AddressPostalCode field value if set, zero value otherwise.
func (o *OrganizationsCollectionResponseObject) GetAddressPostalCode() string {
	if o == nil || IsNil(o.AddressPostalCode) {
		var ret string
		return ret
	}
	return *o.AddressPostalCode
}

// GetAddressPostalCodeOk returns a tuple with the AddressPostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsCollectionResponseObject) GetAddressPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AddressPostalCode) {
		return nil, false
	}
	return o.AddressPostalCode, true
}

// HasAddressPostalCode returns a boolean if a field has been set.
func (o *OrganizationsCollectionResponseObject) HasAddressPostalCode() bool {
	if o != nil && !IsNil(o.AddressPostalCode) {
		return true
	}

	return false
}

// SetAddressPostalCode gets a reference to the given string and assigns it to the AddressPostalCode field.
func (o *OrganizationsCollectionResponseObject) SetAddressPostalCode(v string) {
	o.AddressPostalCode = &v
}

// GetAddressFormattedAddress returns the AddressFormattedAddress field value if set, zero value otherwise.
func (o *OrganizationsCollectionResponseObject) GetAddressFormattedAddress() string {
	if o == nil || IsNil(o.AddressFormattedAddress) {
		var ret string
		return ret
	}
	return *o.AddressFormattedAddress
}

// GetAddressFormattedAddressOk returns a tuple with the AddressFormattedAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsCollectionResponseObject) GetAddressFormattedAddressOk() (*string, bool) {
	if o == nil || IsNil(o.AddressFormattedAddress) {
		return nil, false
	}
	return o.AddressFormattedAddress, true
}

// HasAddressFormattedAddress returns a boolean if a field has been set.
func (o *OrganizationsCollectionResponseObject) HasAddressFormattedAddress() bool {
	if o != nil && !IsNil(o.AddressFormattedAddress) {
		return true
	}

	return false
}

// SetAddressFormattedAddress gets a reference to the given string and assigns it to the AddressFormattedAddress field.
func (o *OrganizationsCollectionResponseObject) SetAddressFormattedAddress(v string) {
	o.AddressFormattedAddress = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationsCollectionResponseObject) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsCollectionResponseObject) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationsCollectionResponseObject) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OrganizationsCollectionResponseObject) SetId(v int32) {
	o.Id = &v
}

// GetActiveFlag returns the ActiveFlag field value if set, zero value otherwise.
func (o *OrganizationsCollectionResponseObject) GetActiveFlag() bool {
	if o == nil || IsNil(o.ActiveFlag) {
		var ret bool
		return ret
	}
	return *o.ActiveFlag
}

// GetActiveFlagOk returns a tuple with the ActiveFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsCollectionResponseObject) GetActiveFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.ActiveFlag) {
		return nil, false
	}
	return o.ActiveFlag, true
}

// HasActiveFlag returns a boolean if a field has been set.
func (o *OrganizationsCollectionResponseObject) HasActiveFlag() bool {
	if o != nil && !IsNil(o.ActiveFlag) {
		return true
	}

	return false
}

// SetActiveFlag gets a reference to the given bool and assigns it to the ActiveFlag field.
func (o *OrganizationsCollectionResponseObject) SetActiveFlag(v bool) {
	o.ActiveFlag = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *OrganizationsCollectionResponseObject) GetOwnerId() int32 {
	if o == nil || IsNil(o.OwnerId) {
		var ret int32
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsCollectionResponseObject) GetOwnerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *OrganizationsCollectionResponseObject) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given int32 and assigns it to the OwnerId field.
func (o *OrganizationsCollectionResponseObject) SetOwnerId(v int32) {
	o.OwnerId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrganizationsCollectionResponseObject) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsCollectionResponseObject) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrganizationsCollectionResponseObject) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrganizationsCollectionResponseObject) SetName(v string) {
	o.Name = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *OrganizationsCollectionResponseObject) GetUpdateTime() string {
	if o == nil || IsNil(o.UpdateTime) {
		var ret string
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsCollectionResponseObject) GetUpdateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *OrganizationsCollectionResponseObject) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given string and assigns it to the UpdateTime field.
func (o *OrganizationsCollectionResponseObject) SetUpdateTime(v string) {
	o.UpdateTime = &v
}

// GetDeleteTime returns the DeleteTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationsCollectionResponseObject) GetDeleteTime() string {
	if o == nil || IsNil(o.DeleteTime.Get()) {
		var ret string
		return ret
	}
	return *o.DeleteTime.Get()
}

// GetDeleteTimeOk returns a tuple with the DeleteTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationsCollectionResponseObject) GetDeleteTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeleteTime.Get(), o.DeleteTime.IsSet()
}

// HasDeleteTime returns a boolean if a field has been set.
func (o *OrganizationsCollectionResponseObject) HasDeleteTime() bool {
	if o != nil && o.DeleteTime.IsSet() {
		return true
	}

	return false
}

// SetDeleteTime gets a reference to the given NullableString and assigns it to the DeleteTime field.
func (o *OrganizationsCollectionResponseObject) SetDeleteTime(v string) {
	o.DeleteTime.Set(&v)
}
// SetDeleteTimeNil sets the value for DeleteTime to be an explicit nil
func (o *OrganizationsCollectionResponseObject) SetDeleteTimeNil() {
	o.DeleteTime.Set(nil)
}

// UnsetDeleteTime ensures that no value is present for DeleteTime, not even an explicit nil
func (o *OrganizationsCollectionResponseObject) UnsetDeleteTime() {
	o.DeleteTime.Unset()
}

// GetAddTime returns the AddTime field value if set, zero value otherwise.
func (o *OrganizationsCollectionResponseObject) GetAddTime() string {
	if o == nil || IsNil(o.AddTime) {
		var ret string
		return ret
	}
	return *o.AddTime
}

// GetAddTimeOk returns a tuple with the AddTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsCollectionResponseObject) GetAddTimeOk() (*string, bool) {
	if o == nil || IsNil(o.AddTime) {
		return nil, false
	}
	return o.AddTime, true
}

// HasAddTime returns a boolean if a field has been set.
func (o *OrganizationsCollectionResponseObject) HasAddTime() bool {
	if o != nil && !IsNil(o.AddTime) {
		return true
	}

	return false
}

// SetAddTime gets a reference to the given string and assigns it to the AddTime field.
func (o *OrganizationsCollectionResponseObject) SetAddTime(v string) {
	o.AddTime = &v
}

// GetVisibleTo returns the VisibleTo field value if set, zero value otherwise.
func (o *OrganizationsCollectionResponseObject) GetVisibleTo() string {
	if o == nil || IsNil(o.VisibleTo) {
		var ret string
		return ret
	}
	return *o.VisibleTo
}

// GetVisibleToOk returns a tuple with the VisibleTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsCollectionResponseObject) GetVisibleToOk() (*string, bool) {
	if o == nil || IsNil(o.VisibleTo) {
		return nil, false
	}
	return o.VisibleTo, true
}

// HasVisibleTo returns a boolean if a field has been set.
func (o *OrganizationsCollectionResponseObject) HasVisibleTo() bool {
	if o != nil && !IsNil(o.VisibleTo) {
		return true
	}

	return false
}

// SetVisibleTo gets a reference to the given string and assigns it to the VisibleTo field.
func (o *OrganizationsCollectionResponseObject) SetVisibleTo(v string) {
	o.VisibleTo = &v
}

// GetLabel returns the Label field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationsCollectionResponseObject) GetLabel() int32 {
	if o == nil || IsNil(o.Label.Get()) {
		var ret int32
		return ret
	}
	return *o.Label.Get()
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OrganizationsCollectionResponseObject) GetLabelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Label.Get(), o.Label.IsSet()
}

// HasLabel returns a boolean if a field has been set.
func (o *OrganizationsCollectionResponseObject) HasLabel() bool {
	if o != nil && o.Label.IsSet() {
		return true
	}

	return false
}

// SetLabel gets a reference to the given NullableInt32 and assigns it to the Label field.
func (o *OrganizationsCollectionResponseObject) SetLabel(v int32) {
	o.Label.Set(&v)
}
// SetLabelNil sets the value for Label to be an explicit nil
func (o *OrganizationsCollectionResponseObject) SetLabelNil() {
	o.Label.Set(nil)
}

// UnsetLabel ensures that no value is present for Label, not even an explicit nil
func (o *OrganizationsCollectionResponseObject) UnsetLabel() {
	o.Label.Unset()
}

// GetCcEmail returns the CcEmail field value if set, zero value otherwise.
func (o *OrganizationsCollectionResponseObject) GetCcEmail() string {
	if o == nil || IsNil(o.CcEmail) {
		var ret string
		return ret
	}
	return *o.CcEmail
}

// GetCcEmailOk returns a tuple with the CcEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationsCollectionResponseObject) GetCcEmailOk() (*string, bool) {
	if o == nil || IsNil(o.CcEmail) {
		return nil, false
	}
	return o.CcEmail, true
}

// HasCcEmail returns a boolean if a field has been set.
func (o *OrganizationsCollectionResponseObject) HasCcEmail() bool {
	if o != nil && !IsNil(o.CcEmail) {
		return true
	}

	return false
}

// SetCcEmail gets a reference to the given string and assigns it to the CcEmail field.
func (o *OrganizationsCollectionResponseObject) SetCcEmail(v string) {
	o.CcEmail = &v
}

func (o OrganizationsCollectionResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationsCollectionResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.AddressSubpremise) {
		toSerialize["address_subpremise"] = o.AddressSubpremise
	}
	if !IsNil(o.AddressStreetNumber) {
		toSerialize["address_street_number"] = o.AddressStreetNumber
	}
	if !IsNil(o.AddressRoute) {
		toSerialize["address_route"] = o.AddressRoute
	}
	if !IsNil(o.AddressSublocality) {
		toSerialize["address_sublocality"] = o.AddressSublocality
	}
	if !IsNil(o.AddressLocality) {
		toSerialize["address_locality"] = o.AddressLocality
	}
	if !IsNil(o.AddressAdminAreaLevel1) {
		toSerialize["address_admin_area_level_1"] = o.AddressAdminAreaLevel1
	}
	if !IsNil(o.AddressAdminAreaLevel2) {
		toSerialize["address_admin_area_level_2"] = o.AddressAdminAreaLevel2
	}
	if !IsNil(o.AddressCountry) {
		toSerialize["address_country"] = o.AddressCountry
	}
	if !IsNil(o.AddressPostalCode) {
		toSerialize["address_postal_code"] = o.AddressPostalCode
	}
	if !IsNil(o.AddressFormattedAddress) {
		toSerialize["address_formatted_address"] = o.AddressFormattedAddress
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ActiveFlag) {
		toSerialize["active_flag"] = o.ActiveFlag
	}
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["update_time"] = o.UpdateTime
	}
	if o.DeleteTime.IsSet() {
		toSerialize["delete_time"] = o.DeleteTime.Get()
	}
	if !IsNil(o.AddTime) {
		toSerialize["add_time"] = o.AddTime
	}
	if !IsNil(o.VisibleTo) {
		toSerialize["visible_to"] = o.VisibleTo
	}
	if o.Label.IsSet() {
		toSerialize["label"] = o.Label.Get()
	}
	if !IsNil(o.CcEmail) {
		toSerialize["cc_email"] = o.CcEmail
	}
	return toSerialize, nil
}

type NullableOrganizationsCollectionResponseObject struct {
	value *OrganizationsCollectionResponseObject
	isSet bool
}

func (v NullableOrganizationsCollectionResponseObject) Get() *OrganizationsCollectionResponseObject {
	return v.value
}

func (v *NullableOrganizationsCollectionResponseObject) Set(val *OrganizationsCollectionResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationsCollectionResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationsCollectionResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationsCollectionResponseObject(val *OrganizationsCollectionResponseObject) *NullableOrganizationsCollectionResponseObject {
	return &NullableOrganizationsCollectionResponseObject{value: val, isSet: true}
}

func (v NullableOrganizationsCollectionResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationsCollectionResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


