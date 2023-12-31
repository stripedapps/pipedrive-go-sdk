/*
Pipedrive API v1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// OrganizationRelationshipsAPIService OrganizationRelationshipsAPI service
type OrganizationRelationshipsAPIService service

type ApiAddOrganizationRelationshipRequest struct {
	ctx context.Context
	ApiService *OrganizationRelationshipsAPIService
	addOrganizationRelationshipRequest *AddOrganizationRelationshipRequest
}

func (r ApiAddOrganizationRelationshipRequest) AddOrganizationRelationshipRequest(addOrganizationRelationshipRequest AddOrganizationRelationshipRequest) ApiAddOrganizationRelationshipRequest {
	r.addOrganizationRelationshipRequest = &addOrganizationRelationshipRequest
	return r
}

func (r ApiAddOrganizationRelationshipRequest) Execute() (*AddOrganizationRelationshipResponse200, *http.Response, error) {
	return r.ApiService.AddOrganizationRelationshipExecute(r)
}

/*
AddOrganizationRelationship Create an organization relationship

Creates and returns an organization relationship.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddOrganizationRelationshipRequest
*/
func (a *OrganizationRelationshipsAPIService) AddOrganizationRelationship(ctx context.Context) ApiAddOrganizationRelationshipRequest {
	return ApiAddOrganizationRelationshipRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AddOrganizationRelationshipResponse200
func (a *OrganizationRelationshipsAPIService) AddOrganizationRelationshipExecute(r ApiAddOrganizationRelationshipRequest) (*AddOrganizationRelationshipResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AddOrganizationRelationshipResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationRelationshipsAPIService.AddOrganizationRelationship")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizationRelationships"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.addOrganizationRelationshipRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteOrganizationRelationshipRequest struct {
	ctx context.Context
	ApiService *OrganizationRelationshipsAPIService
	id int32
}

func (r ApiDeleteOrganizationRelationshipRequest) Execute() (*DeleteOrganizationRelationshipResponse200, *http.Response, error) {
	return r.ApiService.DeleteOrganizationRelationshipExecute(r)
}

/*
DeleteOrganizationRelationship Delete an organization relationship

Deletes an organization relationship and returns the deleted ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the organization relationship
 @return ApiDeleteOrganizationRelationshipRequest
*/
func (a *OrganizationRelationshipsAPIService) DeleteOrganizationRelationship(ctx context.Context, id int32) ApiDeleteOrganizationRelationshipRequest {
	return ApiDeleteOrganizationRelationshipRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return DeleteOrganizationRelationshipResponse200
func (a *OrganizationRelationshipsAPIService) DeleteOrganizationRelationshipExecute(r ApiDeleteOrganizationRelationshipRequest) (*DeleteOrganizationRelationshipResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DeleteOrganizationRelationshipResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationRelationshipsAPIService.DeleteOrganizationRelationship")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizationRelationships/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetOrganizationRelationshipRequest struct {
	ctx context.Context
	ApiService *OrganizationRelationshipsAPIService
	id int32
	orgId *int32
}

// The ID of the base organization for the returned calculated values
func (r ApiGetOrganizationRelationshipRequest) OrgId(orgId int32) ApiGetOrganizationRelationshipRequest {
	r.orgId = &orgId
	return r
}

func (r ApiGetOrganizationRelationshipRequest) Execute() (*GetOrganizationRelationshipResponse200, *http.Response, error) {
	return r.ApiService.GetOrganizationRelationshipExecute(r)
}

/*
GetOrganizationRelationship Get one organization relationship

Finds and returns an organization relationship from its ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the organization relationship
 @return ApiGetOrganizationRelationshipRequest
*/
func (a *OrganizationRelationshipsAPIService) GetOrganizationRelationship(ctx context.Context, id int32) ApiGetOrganizationRelationshipRequest {
	return ApiGetOrganizationRelationshipRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GetOrganizationRelationshipResponse200
func (a *OrganizationRelationshipsAPIService) GetOrganizationRelationshipExecute(r ApiGetOrganizationRelationshipRequest) (*GetOrganizationRelationshipResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetOrganizationRelationshipResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationRelationshipsAPIService.GetOrganizationRelationship")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizationRelationships/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.orgId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "org_id", r.orgId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetOrganizationRelationshipsRequest struct {
	ctx context.Context
	ApiService *OrganizationRelationshipsAPIService
	orgId *int32
}

// The ID of the organization to get relationships for
func (r ApiGetOrganizationRelationshipsRequest) OrgId(orgId int32) ApiGetOrganizationRelationshipsRequest {
	r.orgId = &orgId
	return r
}

func (r ApiGetOrganizationRelationshipsRequest) Execute() (*GetOrganizationRelationshipsResponse200, *http.Response, error) {
	return r.ApiService.GetOrganizationRelationshipsExecute(r)
}

/*
GetOrganizationRelationships Get all relationships for organization

Gets all of the relationships for a supplied organization ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetOrganizationRelationshipsRequest
*/
func (a *OrganizationRelationshipsAPIService) GetOrganizationRelationships(ctx context.Context) ApiGetOrganizationRelationshipsRequest {
	return ApiGetOrganizationRelationshipsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetOrganizationRelationshipsResponse200
func (a *OrganizationRelationshipsAPIService) GetOrganizationRelationshipsExecute(r ApiGetOrganizationRelationshipsRequest) (*GetOrganizationRelationshipsResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetOrganizationRelationshipsResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationRelationshipsAPIService.GetOrganizationRelationships")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizationRelationships"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.orgId == nil {
		return localVarReturnValue, nil, reportError("orgId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "org_id", r.orgId, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateOrganizationRelationshipRequest struct {
	ctx context.Context
	ApiService *OrganizationRelationshipsAPIService
	id int32
	organizationRelationship *OrganizationRelationship
}

func (r ApiUpdateOrganizationRelationshipRequest) OrganizationRelationship(organizationRelationship OrganizationRelationship) ApiUpdateOrganizationRelationshipRequest {
	r.organizationRelationship = &organizationRelationship
	return r
}

func (r ApiUpdateOrganizationRelationshipRequest) Execute() (*UpdateOrganizationRelationshipResponse200, *http.Response, error) {
	return r.ApiService.UpdateOrganizationRelationshipExecute(r)
}

/*
UpdateOrganizationRelationship Update an organization relationship

Updates and returns an organization relationship.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the organization relationship
 @return ApiUpdateOrganizationRelationshipRequest
*/
func (a *OrganizationRelationshipsAPIService) UpdateOrganizationRelationship(ctx context.Context, id int32) ApiUpdateOrganizationRelationshipRequest {
	return ApiUpdateOrganizationRelationshipRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return UpdateOrganizationRelationshipResponse200
func (a *OrganizationRelationshipsAPIService) UpdateOrganizationRelationshipExecute(r ApiUpdateOrganizationRelationshipRequest) (*UpdateOrganizationRelationshipResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpdateOrganizationRelationshipResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrganizationRelationshipsAPIService.UpdateOrganizationRelationship")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizationRelationships/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.organizationRelationship
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_token", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
