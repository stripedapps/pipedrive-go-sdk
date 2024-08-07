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
	"reflect"
)


// GoalsAPIService GoalsAPI service
type GoalsAPIService service

type ApiAddGoalRequest struct {
	ctx context.Context
	ApiService *GoalsAPIService
	addGoalRequest *AddGoalRequest
}

func (r ApiAddGoalRequest) AddGoalRequest(addGoalRequest AddGoalRequest) ApiAddGoalRequest {
	r.addGoalRequest = &addGoalRequest
	return r
}

func (r ApiAddGoalRequest) Execute() (*AddOrUpdateGoalResponse200, *http.Response, error) {
	return r.ApiService.AddGoalExecute(r)
}

/*
AddGoal Add a new goal

Adds a new goal. Along with adding a new goal, a report is created to track the progress of your goal.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddGoalRequest
*/
func (a *GoalsAPIService) AddGoal(ctx context.Context) ApiAddGoalRequest {
	return ApiAddGoalRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AddOrUpdateGoalResponse200
func (a *GoalsAPIService) AddGoalExecute(r ApiAddGoalRequest) (*AddOrUpdateGoalResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AddOrUpdateGoalResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoalsAPIService.AddGoal")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/goals"

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
	localVarPostBody = r.addGoalRequest
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

type ApiDeleteGoalRequest struct {
	ctx context.Context
	ApiService *GoalsAPIService
	id string
}

func (r ApiDeleteGoalRequest) Execute() (*DeleteGoalResponse200, *http.Response, error) {
	return r.ApiService.DeleteGoalExecute(r)
}

/*
DeleteGoal Delete existing goal

Marks a goal as deleted.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the goal
 @return ApiDeleteGoalRequest
*/
func (a *GoalsAPIService) DeleteGoal(ctx context.Context, id string) ApiDeleteGoalRequest {
	return ApiDeleteGoalRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return DeleteGoalResponse200
func (a *GoalsAPIService) DeleteGoalExecute(r ApiDeleteGoalRequest) (*DeleteGoalResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DeleteGoalResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoalsAPIService.DeleteGoal")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/goals/{id}"
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

type ApiGetGoalResultRequest struct {
	ctx context.Context
	ApiService *GoalsAPIService
	id string
	periodStart *string
	periodEnd *string
}

// The start date of the period for which to find the goal&#39;s progress. Format: YYYY-MM-DD. This date must be the same or after the goal duration start date. 
func (r ApiGetGoalResultRequest) PeriodStart(periodStart string) ApiGetGoalResultRequest {
	r.periodStart = &periodStart
	return r
}

// The end date of the period for which to find the goal&#39;s progress. Format: YYYY-MM-DD. This date must be the same or before the goal duration end date. 
func (r ApiGetGoalResultRequest) PeriodEnd(periodEnd string) ApiGetGoalResultRequest {
	r.periodEnd = &periodEnd
	return r
}

func (r ApiGetGoalResultRequest) Execute() (*GetGoalResultResponse200, *http.Response, error) {
	return r.ApiService.GetGoalResultExecute(r)
}

/*
GetGoalResult Get result of a goal

Gets the progress of a goal for the specified period.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the goal that the results are looked for
 @return ApiGetGoalResultRequest
*/
func (a *GoalsAPIService) GetGoalResult(ctx context.Context, id string) ApiGetGoalResultRequest {
	return ApiGetGoalResultRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GetGoalResultResponse200
func (a *GoalsAPIService) GetGoalResultExecute(r ApiGetGoalResultRequest) (*GetGoalResultResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetGoalResultResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoalsAPIService.GetGoalResult")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/goals/{id}/results"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.periodStart == nil {
		return localVarReturnValue, nil, reportError("periodStart is required and must be specified")
	}
	if r.periodEnd == nil {
		return localVarReturnValue, nil, reportError("periodEnd is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "period.start", r.periodStart, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "period.end", r.periodEnd, "")
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

type ApiGetGoalsRequest struct {
	ctx context.Context
	ApiService *GoalsAPIService
	typeName *string
	title *string
	isActive *bool
	assigneeId *int32
	assigneeType *string
	expectedOutcomeTarget *float32
	expectedOutcomeTrackingMetric *string
	expectedOutcomeCurrencyId *int32
	typeParamsPipelineId *[]int32
	typeParamsStageId *int32
	typeParamsActivityTypeId *[]int32
	periodStart *string
	periodEnd *string
}

// The type of the goal. If provided, everyone&#39;s goals will be returned.
func (r ApiGetGoalsRequest) TypeName(typeName string) ApiGetGoalsRequest {
	r.typeName = &typeName
	return r
}

// The title of the goal
func (r ApiGetGoalsRequest) Title(title string) ApiGetGoalsRequest {
	r.title = &title
	return r
}

// Whether the goal is active or not
func (r ApiGetGoalsRequest) IsActive(isActive bool) ApiGetGoalsRequest {
	r.isActive = &isActive
	return r
}

// The ID of the user who&#39;s goal to fetch. When omitted, only your goals will be returned.
func (r ApiGetGoalsRequest) AssigneeId(assigneeId int32) ApiGetGoalsRequest {
	r.assigneeId = &assigneeId
	return r
}

// The type of the goal&#39;s assignee. If provided, everyone&#39;s goals will be returned.
func (r ApiGetGoalsRequest) AssigneeType(assigneeType string) ApiGetGoalsRequest {
	r.assigneeType = &assigneeType
	return r
}

// The numeric value of the outcome. If provided, everyone&#39;s goals will be returned.
func (r ApiGetGoalsRequest) ExpectedOutcomeTarget(expectedOutcomeTarget float32) ApiGetGoalsRequest {
	r.expectedOutcomeTarget = &expectedOutcomeTarget
	return r
}

// The tracking metric of the expected outcome of the goal. If provided, everyone&#39;s goals will be returned.
func (r ApiGetGoalsRequest) ExpectedOutcomeTrackingMetric(expectedOutcomeTrackingMetric string) ApiGetGoalsRequest {
	r.expectedOutcomeTrackingMetric = &expectedOutcomeTrackingMetric
	return r
}

// The numeric ID of the goal&#39;s currency. Only applicable to goals with &#x60;expected_outcome.tracking_metric&#x60; with value &#x60;sum&#x60;. If provided, everyone&#39;s goals will be returned.
func (r ApiGetGoalsRequest) ExpectedOutcomeCurrencyId(expectedOutcomeCurrencyId int32) ApiGetGoalsRequest {
	r.expectedOutcomeCurrencyId = &expectedOutcomeCurrencyId
	return r
}

// An array of pipeline IDs or &#x60;null&#x60; for all pipelines. If provided, everyone&#39;s goals will be returned.
func (r ApiGetGoalsRequest) TypeParamsPipelineId(typeParamsPipelineId []int32) ApiGetGoalsRequest {
	r.typeParamsPipelineId = &typeParamsPipelineId
	return r
}

// The ID of the stage. Applicable to only &#x60;deals_progressed&#x60; type of goals. If provided, everyone&#39;s goals will be returned.
func (r ApiGetGoalsRequest) TypeParamsStageId(typeParamsStageId int32) ApiGetGoalsRequest {
	r.typeParamsStageId = &typeParamsStageId
	return r
}

// An array of IDs or &#x60;null&#x60; for all activity types. Only applicable for &#x60;activities_completed&#x60; and/or &#x60;activities_added&#x60; types of goals. If provided, everyone&#39;s goals will be returned.
func (r ApiGetGoalsRequest) TypeParamsActivityTypeId(typeParamsActivityTypeId []int32) ApiGetGoalsRequest {
	r.typeParamsActivityTypeId = &typeParamsActivityTypeId
	return r
}

// The start date of the period for which to find goals. Date in format of YYYY-MM-DD. When &#x60;period.start&#x60; is provided, &#x60;period.end&#x60; must be provided too.
func (r ApiGetGoalsRequest) PeriodStart(periodStart string) ApiGetGoalsRequest {
	r.periodStart = &periodStart
	return r
}

// The end date of the period for which to find goals. Date in format of YYYY-MM-DD.
func (r ApiGetGoalsRequest) PeriodEnd(periodEnd string) ApiGetGoalsRequest {
	r.periodEnd = &periodEnd
	return r
}

func (r ApiGetGoalsRequest) Execute() (*GetGoalsResponse200, *http.Response, error) {
	return r.ApiService.GetGoalsExecute(r)
}

/*
GetGoals Find goals

Returns data about goals based on criteria. For searching, append `{searchField}={searchValue}` to the URL, where `searchField` can be any one of the lowest-level fields in dot-notation (e.g. `type.params.pipeline_id`; `title`). `searchValue` should be the value you are looking for on that field. Additionally, `is_active=<true|false>` can be provided to search for only active/inactive goals. When providing `period.start`, `period.end` must also be provided and vice versa.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetGoalsRequest
*/
func (a *GoalsAPIService) GetGoals(ctx context.Context) ApiGetGoalsRequest {
	return ApiGetGoalsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetGoalsResponse200
func (a *GoalsAPIService) GetGoalsExecute(r ApiGetGoalsRequest) (*GetGoalsResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetGoalsResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoalsAPIService.GetGoals")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/goals/find"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.typeName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type.name", r.typeName, "")
	}
	if r.title != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "title", r.title, "")
	}
	if r.isActive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is_active", r.isActive, "")
	} else {
		var defaultValue bool = true
		r.isActive = &defaultValue
	}
	if r.assigneeId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "assignee.id", r.assigneeId, "")
	}
	if r.assigneeType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "assignee.type", r.assigneeType, "")
	}
	if r.expectedOutcomeTarget != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "expected_outcome.target", r.expectedOutcomeTarget, "")
	}
	if r.expectedOutcomeTrackingMetric != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "expected_outcome.tracking_metric", r.expectedOutcomeTrackingMetric, "")
	}
	if r.expectedOutcomeCurrencyId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "expected_outcome.currency_id", r.expectedOutcomeCurrencyId, "")
	}
	if r.typeParamsPipelineId != nil {
		t := *r.typeParamsPipelineId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "type.params.pipeline_id", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "type.params.pipeline_id", t, "multi")
		}
	}
	if r.typeParamsStageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type.params.stage_id", r.typeParamsStageId, "")
	}
	if r.typeParamsActivityTypeId != nil {
		t := *r.typeParamsActivityTypeId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "type.params.activity_type_id", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "type.params.activity_type_id", t, "multi")
		}
	}
	if r.periodStart != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period.start", r.periodStart, "")
	}
	if r.periodEnd != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "period.end", r.periodEnd, "")
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

type ApiUpdateGoalRequest struct {
	ctx context.Context
	ApiService *GoalsAPIService
	id string
	basicGoalRequest *BasicGoalRequest
}

func (r ApiUpdateGoalRequest) BasicGoalRequest(basicGoalRequest BasicGoalRequest) ApiUpdateGoalRequest {
	r.basicGoalRequest = &basicGoalRequest
	return r
}

func (r ApiUpdateGoalRequest) Execute() (*AddOrUpdateGoalResponse200, *http.Response, error) {
	return r.ApiService.UpdateGoalExecute(r)
}

/*
UpdateGoal Update existing goal

Updates an existing goal.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the goal
 @return ApiUpdateGoalRequest
*/
func (a *GoalsAPIService) UpdateGoal(ctx context.Context, id string) ApiUpdateGoalRequest {
	return ApiUpdateGoalRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AddOrUpdateGoalResponse200
func (a *GoalsAPIService) UpdateGoalExecute(r ApiUpdateGoalRequest) (*AddOrUpdateGoalResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AddOrUpdateGoalResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GoalsAPIService.UpdateGoal")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/goals/{id}"
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
	localVarPostBody = r.basicGoalRequest
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
