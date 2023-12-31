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

// checks if the GetFilesResponse200DataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFilesResponse200DataInner{}

// GetFilesResponse200DataInner The file data
type GetFilesResponse200DataInner struct {
	// The ID of the file
	Id *int32 `json:"id,omitempty"`
	// The ID of the user to associate the file with
	UserId *int32 `json:"user_id,omitempty"`
	// The ID of the deal to associate the file with
	DealId *int32 `json:"deal_id,omitempty"`
	// The ID of the person to associate the file with
	PersonId *int32 `json:"person_id,omitempty"`
	// The ID of the organization to associate the file with
	OrgId *int32 `json:"org_id,omitempty"`
	// The ID of the product to associate the file with
	ProductId *int32 `json:"product_id,omitempty"`
	// The ID of the activity to associate the file with
	ActivityId *int32 `json:"activity_id,omitempty"`
	// The ID of the lead to associate the file with
	LeadId *string `json:"lead_id,omitempty"`
	// The date and time when the file was added/created. Format: YYYY-MM-DD HH:MM:SS
	AddTime *string `json:"add_time,omitempty"`
	// The last updated date and time of the file. Format: YYYY-MM-DD HH:MM:SS
	UpdateTime *string `json:"update_time,omitempty"`
	// The original name of the file
	FileName *string `json:"file_name,omitempty"`
	// The size of the file
	FileSize *int32 `json:"file_size,omitempty"`
	// Whether the user is active or not. false = Not activated, true = Activated
	ActiveFlag *bool `json:"active_flag,omitempty"`
	// Whether the file was uploaded as inline or not
	InlineFlag *bool `json:"inline_flag,omitempty"`
	// The location type to send the file to. Only googledrive is supported at the moment.
	RemoteLocation *string `json:"remote_location,omitempty"`
	// The ID of the remote item
	RemoteId *string `json:"remote_id,omitempty"`
	// The ID of the inline attachment
	Cid *string `json:"cid,omitempty"`
	// The location of the cloud storage
	S3Bucket *string `json:"s3_bucket,omitempty"`
	// The ID of the mail message to associate the file with
	MailMessageId *string `json:"mail_message_id,omitempty"`
	// The ID of the mail template to associate the file with
	MailTemplateId *string `json:"mail_template_id,omitempty"`
	// The name of the deal associated with the file
	DealName *string `json:"deal_name,omitempty"`
	// The name of the person associated with the file
	PersonName *string `json:"person_name,omitempty"`
	// The name of the organization associated with the file
	OrgName *string `json:"org_name,omitempty"`
	// The name of the product associated with the file
	ProductName *string `json:"product_name,omitempty"`
	// The name of the lead associated with the file
	LeadName *string `json:"lead_name,omitempty"`
	// The URL of the download file
	Url *string `json:"url,omitempty"`
	// The visible name of the file
	Name *string `json:"name,omitempty"`
	// The description of the file
	Description *string `json:"description,omitempty"`
}

// NewGetFilesResponse200DataInner instantiates a new GetFilesResponse200DataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFilesResponse200DataInner() *GetFilesResponse200DataInner {
	this := GetFilesResponse200DataInner{}
	return &this
}

// NewGetFilesResponse200DataInnerWithDefaults instantiates a new GetFilesResponse200DataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFilesResponse200DataInnerWithDefaults() *GetFilesResponse200DataInner {
	this := GetFilesResponse200DataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetFilesResponse200DataInner) SetId(v int32) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *GetFilesResponse200DataInner) SetUserId(v int32) {
	o.UserId = &v
}

// GetDealId returns the DealId field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetDealId() int32 {
	if o == nil || IsNil(o.DealId) {
		var ret int32
		return ret
	}
	return *o.DealId
}

// GetDealIdOk returns a tuple with the DealId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetDealIdOk() (*int32, bool) {
	if o == nil || IsNil(o.DealId) {
		return nil, false
	}
	return o.DealId, true
}

// HasDealId returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasDealId() bool {
	if o != nil && !IsNil(o.DealId) {
		return true
	}

	return false
}

// SetDealId gets a reference to the given int32 and assigns it to the DealId field.
func (o *GetFilesResponse200DataInner) SetDealId(v int32) {
	o.DealId = &v
}

// GetPersonId returns the PersonId field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetPersonId() int32 {
	if o == nil || IsNil(o.PersonId) {
		var ret int32
		return ret
	}
	return *o.PersonId
}

// GetPersonIdOk returns a tuple with the PersonId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetPersonIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PersonId) {
		return nil, false
	}
	return o.PersonId, true
}

// HasPersonId returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasPersonId() bool {
	if o != nil && !IsNil(o.PersonId) {
		return true
	}

	return false
}

// SetPersonId gets a reference to the given int32 and assigns it to the PersonId field.
func (o *GetFilesResponse200DataInner) SetPersonId(v int32) {
	o.PersonId = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetOrgId() int32 {
	if o == nil || IsNil(o.OrgId) {
		var ret int32
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetOrgIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given int32 and assigns it to the OrgId field.
func (o *GetFilesResponse200DataInner) SetOrgId(v int32) {
	o.OrgId = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetProductId() int32 {
	if o == nil || IsNil(o.ProductId) {
		var ret int32
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetProductIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given int32 and assigns it to the ProductId field.
func (o *GetFilesResponse200DataInner) SetProductId(v int32) {
	o.ProductId = &v
}

// GetActivityId returns the ActivityId field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetActivityId() int32 {
	if o == nil || IsNil(o.ActivityId) {
		var ret int32
		return ret
	}
	return *o.ActivityId
}

// GetActivityIdOk returns a tuple with the ActivityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetActivityIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ActivityId) {
		return nil, false
	}
	return o.ActivityId, true
}

// HasActivityId returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasActivityId() bool {
	if o != nil && !IsNil(o.ActivityId) {
		return true
	}

	return false
}

// SetActivityId gets a reference to the given int32 and assigns it to the ActivityId field.
func (o *GetFilesResponse200DataInner) SetActivityId(v int32) {
	o.ActivityId = &v
}

// GetLeadId returns the LeadId field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetLeadId() string {
	if o == nil || IsNil(o.LeadId) {
		var ret string
		return ret
	}
	return *o.LeadId
}

// GetLeadIdOk returns a tuple with the LeadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetLeadIdOk() (*string, bool) {
	if o == nil || IsNil(o.LeadId) {
		return nil, false
	}
	return o.LeadId, true
}

// HasLeadId returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasLeadId() bool {
	if o != nil && !IsNil(o.LeadId) {
		return true
	}

	return false
}

// SetLeadId gets a reference to the given string and assigns it to the LeadId field.
func (o *GetFilesResponse200DataInner) SetLeadId(v string) {
	o.LeadId = &v
}

// GetAddTime returns the AddTime field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetAddTime() string {
	if o == nil || IsNil(o.AddTime) {
		var ret string
		return ret
	}
	return *o.AddTime
}

// GetAddTimeOk returns a tuple with the AddTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetAddTimeOk() (*string, bool) {
	if o == nil || IsNil(o.AddTime) {
		return nil, false
	}
	return o.AddTime, true
}

// HasAddTime returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasAddTime() bool {
	if o != nil && !IsNil(o.AddTime) {
		return true
	}

	return false
}

// SetAddTime gets a reference to the given string and assigns it to the AddTime field.
func (o *GetFilesResponse200DataInner) SetAddTime(v string) {
	o.AddTime = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetUpdateTime() string {
	if o == nil || IsNil(o.UpdateTime) {
		var ret string
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetUpdateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given string and assigns it to the UpdateTime field.
func (o *GetFilesResponse200DataInner) SetUpdateTime(v string) {
	o.UpdateTime = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *GetFilesResponse200DataInner) SetFileName(v string) {
	o.FileName = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetFileSize() int32 {
	if o == nil || IsNil(o.FileSize) {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetFileSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *GetFilesResponse200DataInner) SetFileSize(v int32) {
	o.FileSize = &v
}

// GetActiveFlag returns the ActiveFlag field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetActiveFlag() bool {
	if o == nil || IsNil(o.ActiveFlag) {
		var ret bool
		return ret
	}
	return *o.ActiveFlag
}

// GetActiveFlagOk returns a tuple with the ActiveFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetActiveFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.ActiveFlag) {
		return nil, false
	}
	return o.ActiveFlag, true
}

// HasActiveFlag returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasActiveFlag() bool {
	if o != nil && !IsNil(o.ActiveFlag) {
		return true
	}

	return false
}

// SetActiveFlag gets a reference to the given bool and assigns it to the ActiveFlag field.
func (o *GetFilesResponse200DataInner) SetActiveFlag(v bool) {
	o.ActiveFlag = &v
}

// GetInlineFlag returns the InlineFlag field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetInlineFlag() bool {
	if o == nil || IsNil(o.InlineFlag) {
		var ret bool
		return ret
	}
	return *o.InlineFlag
}

// GetInlineFlagOk returns a tuple with the InlineFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetInlineFlagOk() (*bool, bool) {
	if o == nil || IsNil(o.InlineFlag) {
		return nil, false
	}
	return o.InlineFlag, true
}

// HasInlineFlag returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasInlineFlag() bool {
	if o != nil && !IsNil(o.InlineFlag) {
		return true
	}

	return false
}

// SetInlineFlag gets a reference to the given bool and assigns it to the InlineFlag field.
func (o *GetFilesResponse200DataInner) SetInlineFlag(v bool) {
	o.InlineFlag = &v
}

// GetRemoteLocation returns the RemoteLocation field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetRemoteLocation() string {
	if o == nil || IsNil(o.RemoteLocation) {
		var ret string
		return ret
	}
	return *o.RemoteLocation
}

// GetRemoteLocationOk returns a tuple with the RemoteLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetRemoteLocationOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteLocation) {
		return nil, false
	}
	return o.RemoteLocation, true
}

// HasRemoteLocation returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasRemoteLocation() bool {
	if o != nil && !IsNil(o.RemoteLocation) {
		return true
	}

	return false
}

// SetRemoteLocation gets a reference to the given string and assigns it to the RemoteLocation field.
func (o *GetFilesResponse200DataInner) SetRemoteLocation(v string) {
	o.RemoteLocation = &v
}

// GetRemoteId returns the RemoteId field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetRemoteId() string {
	if o == nil || IsNil(o.RemoteId) {
		var ret string
		return ret
	}
	return *o.RemoteId
}

// GetRemoteIdOk returns a tuple with the RemoteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetRemoteIdOk() (*string, bool) {
	if o == nil || IsNil(o.RemoteId) {
		return nil, false
	}
	return o.RemoteId, true
}

// HasRemoteId returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasRemoteId() bool {
	if o != nil && !IsNil(o.RemoteId) {
		return true
	}

	return false
}

// SetRemoteId gets a reference to the given string and assigns it to the RemoteId field.
func (o *GetFilesResponse200DataInner) SetRemoteId(v string) {
	o.RemoteId = &v
}

// GetCid returns the Cid field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetCid() string {
	if o == nil || IsNil(o.Cid) {
		var ret string
		return ret
	}
	return *o.Cid
}

// GetCidOk returns a tuple with the Cid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetCidOk() (*string, bool) {
	if o == nil || IsNil(o.Cid) {
		return nil, false
	}
	return o.Cid, true
}

// HasCid returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasCid() bool {
	if o != nil && !IsNil(o.Cid) {
		return true
	}

	return false
}

// SetCid gets a reference to the given string and assigns it to the Cid field.
func (o *GetFilesResponse200DataInner) SetCid(v string) {
	o.Cid = &v
}

// GetS3Bucket returns the S3Bucket field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetS3Bucket() string {
	if o == nil || IsNil(o.S3Bucket) {
		var ret string
		return ret
	}
	return *o.S3Bucket
}

// GetS3BucketOk returns a tuple with the S3Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetS3BucketOk() (*string, bool) {
	if o == nil || IsNil(o.S3Bucket) {
		return nil, false
	}
	return o.S3Bucket, true
}

// HasS3Bucket returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasS3Bucket() bool {
	if o != nil && !IsNil(o.S3Bucket) {
		return true
	}

	return false
}

// SetS3Bucket gets a reference to the given string and assigns it to the S3Bucket field.
func (o *GetFilesResponse200DataInner) SetS3Bucket(v string) {
	o.S3Bucket = &v
}

// GetMailMessageId returns the MailMessageId field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetMailMessageId() string {
	if o == nil || IsNil(o.MailMessageId) {
		var ret string
		return ret
	}
	return *o.MailMessageId
}

// GetMailMessageIdOk returns a tuple with the MailMessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetMailMessageIdOk() (*string, bool) {
	if o == nil || IsNil(o.MailMessageId) {
		return nil, false
	}
	return o.MailMessageId, true
}

// HasMailMessageId returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasMailMessageId() bool {
	if o != nil && !IsNil(o.MailMessageId) {
		return true
	}

	return false
}

// SetMailMessageId gets a reference to the given string and assigns it to the MailMessageId field.
func (o *GetFilesResponse200DataInner) SetMailMessageId(v string) {
	o.MailMessageId = &v
}

// GetMailTemplateId returns the MailTemplateId field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetMailTemplateId() string {
	if o == nil || IsNil(o.MailTemplateId) {
		var ret string
		return ret
	}
	return *o.MailTemplateId
}

// GetMailTemplateIdOk returns a tuple with the MailTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetMailTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.MailTemplateId) {
		return nil, false
	}
	return o.MailTemplateId, true
}

// HasMailTemplateId returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasMailTemplateId() bool {
	if o != nil && !IsNil(o.MailTemplateId) {
		return true
	}

	return false
}

// SetMailTemplateId gets a reference to the given string and assigns it to the MailTemplateId field.
func (o *GetFilesResponse200DataInner) SetMailTemplateId(v string) {
	o.MailTemplateId = &v
}

// GetDealName returns the DealName field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetDealName() string {
	if o == nil || IsNil(o.DealName) {
		var ret string
		return ret
	}
	return *o.DealName
}

// GetDealNameOk returns a tuple with the DealName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetDealNameOk() (*string, bool) {
	if o == nil || IsNil(o.DealName) {
		return nil, false
	}
	return o.DealName, true
}

// HasDealName returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasDealName() bool {
	if o != nil && !IsNil(o.DealName) {
		return true
	}

	return false
}

// SetDealName gets a reference to the given string and assigns it to the DealName field.
func (o *GetFilesResponse200DataInner) SetDealName(v string) {
	o.DealName = &v
}

// GetPersonName returns the PersonName field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetPersonName() string {
	if o == nil || IsNil(o.PersonName) {
		var ret string
		return ret
	}
	return *o.PersonName
}

// GetPersonNameOk returns a tuple with the PersonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetPersonNameOk() (*string, bool) {
	if o == nil || IsNil(o.PersonName) {
		return nil, false
	}
	return o.PersonName, true
}

// HasPersonName returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasPersonName() bool {
	if o != nil && !IsNil(o.PersonName) {
		return true
	}

	return false
}

// SetPersonName gets a reference to the given string and assigns it to the PersonName field.
func (o *GetFilesResponse200DataInner) SetPersonName(v string) {
	o.PersonName = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetOrgName() string {
	if o == nil || IsNil(o.OrgName) {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetOrgNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrgName) {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasOrgName() bool {
	if o != nil && !IsNil(o.OrgName) {
		return true
	}

	return false
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *GetFilesResponse200DataInner) SetOrgName(v string) {
	o.OrgName = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetProductName() string {
	if o == nil || IsNil(o.ProductName) {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.ProductName) {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasProductName() bool {
	if o != nil && !IsNil(o.ProductName) {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *GetFilesResponse200DataInner) SetProductName(v string) {
	o.ProductName = &v
}

// GetLeadName returns the LeadName field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetLeadName() string {
	if o == nil || IsNil(o.LeadName) {
		var ret string
		return ret
	}
	return *o.LeadName
}

// GetLeadNameOk returns a tuple with the LeadName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetLeadNameOk() (*string, bool) {
	if o == nil || IsNil(o.LeadName) {
		return nil, false
	}
	return o.LeadName, true
}

// HasLeadName returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasLeadName() bool {
	if o != nil && !IsNil(o.LeadName) {
		return true
	}

	return false
}

// SetLeadName gets a reference to the given string and assigns it to the LeadName field.
func (o *GetFilesResponse200DataInner) SetLeadName(v string) {
	o.LeadName = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GetFilesResponse200DataInner) SetUrl(v string) {
	o.Url = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetFilesResponse200DataInner) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetFilesResponse200DataInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetFilesResponse200DataInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetFilesResponse200DataInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetFilesResponse200DataInner) SetDescription(v string) {
	o.Description = &v
}

func (o GetFilesResponse200DataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFilesResponse200DataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.DealId) {
		toSerialize["deal_id"] = o.DealId
	}
	if !IsNil(o.PersonId) {
		toSerialize["person_id"] = o.PersonId
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.ProductId) {
		toSerialize["product_id"] = o.ProductId
	}
	if !IsNil(o.ActivityId) {
		toSerialize["activity_id"] = o.ActivityId
	}
	if !IsNil(o.LeadId) {
		toSerialize["lead_id"] = o.LeadId
	}
	if !IsNil(o.AddTime) {
		toSerialize["add_time"] = o.AddTime
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["update_time"] = o.UpdateTime
	}
	if !IsNil(o.FileName) {
		toSerialize["file_name"] = o.FileName
	}
	if !IsNil(o.FileSize) {
		toSerialize["file_size"] = o.FileSize
	}
	if !IsNil(o.ActiveFlag) {
		toSerialize["active_flag"] = o.ActiveFlag
	}
	if !IsNil(o.InlineFlag) {
		toSerialize["inline_flag"] = o.InlineFlag
	}
	if !IsNil(o.RemoteLocation) {
		toSerialize["remote_location"] = o.RemoteLocation
	}
	if !IsNil(o.RemoteId) {
		toSerialize["remote_id"] = o.RemoteId
	}
	if !IsNil(o.Cid) {
		toSerialize["cid"] = o.Cid
	}
	if !IsNil(o.S3Bucket) {
		toSerialize["s3_bucket"] = o.S3Bucket
	}
	if !IsNil(o.MailMessageId) {
		toSerialize["mail_message_id"] = o.MailMessageId
	}
	if !IsNil(o.MailTemplateId) {
		toSerialize["mail_template_id"] = o.MailTemplateId
	}
	if !IsNil(o.DealName) {
		toSerialize["deal_name"] = o.DealName
	}
	if !IsNil(o.PersonName) {
		toSerialize["person_name"] = o.PersonName
	}
	if !IsNil(o.OrgName) {
		toSerialize["org_name"] = o.OrgName
	}
	if !IsNil(o.ProductName) {
		toSerialize["product_name"] = o.ProductName
	}
	if !IsNil(o.LeadName) {
		toSerialize["lead_name"] = o.LeadName
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableGetFilesResponse200DataInner struct {
	value *GetFilesResponse200DataInner
	isSet bool
}

func (v NullableGetFilesResponse200DataInner) Get() *GetFilesResponse200DataInner {
	return v.value
}

func (v *NullableGetFilesResponse200DataInner) Set(val *GetFilesResponse200DataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFilesResponse200DataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFilesResponse200DataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFilesResponse200DataInner(val *GetFilesResponse200DataInner) *NullableGetFilesResponse200DataInner {
	return &NullableGetFilesResponse200DataInner{value: val, isSet: true}
}

func (v NullableGetFilesResponse200DataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFilesResponse200DataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


