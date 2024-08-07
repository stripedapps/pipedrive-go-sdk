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
)


// OauthAPIService OauthAPI service
type OauthAPIService service

type ApiAuthorizeRequest struct {
	ctx context.Context
	ApiService *OauthAPIService
	clientId *string
	redirectUri *string
	state *string
}

// The client ID provided to you by the Pipedrive Marketplace when you register your app
func (r ApiAuthorizeRequest) ClientId(clientId string) ApiAuthorizeRequest {
	r.clientId = &clientId
	return r
}

// The callback URL you provided when you registered your app. Authorization code will be sent to that URL (if it matches with the value you entered in the registration form) if a user approves the app install. Or, if a customer declines, the corresponding error will also be sent to this URL.
func (r ApiAuthorizeRequest) RedirectUri(redirectUri string) ApiAuthorizeRequest {
	r.redirectUri = &redirectUri
	return r
}

// You may pass any random string as the state parameter and the same string will be returned to your app after a user authorizes access. It may be used to store the user&#39;s session ID from your app or distinguish different responses. Using state may increase security; see RFC-6749. 
func (r ApiAuthorizeRequest) State(state string) ApiAuthorizeRequest {
	r.state = &state
	return r
}

func (r ApiAuthorizeRequest) Execute() (*http.Response, error) {
	return r.ApiService.AuthorizeExecute(r)
}

/*
Authorize Requesting authorization

Authorize a user by redirecting them to the Pipedrive OAuth authorization page and request their permissions to act on their behalf. This step is necessary to implement only when you allow app installation outside of the Marketplace.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuthorizeRequest
*/
func (a *OauthAPIService) Authorize(ctx context.Context) ApiAuthorizeRequest {
	return ApiAuthorizeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *OauthAPIService) AuthorizeExecute(r ApiAuthorizeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthAPIService.Authorize")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/authorize"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientId == nil {
		return nil, reportError("clientId is required and must be specified")
	}
	if r.redirectUri == nil {
		return nil, reportError("redirectUri is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "client_id", r.clientId, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "redirect_uri", r.redirectUri, "")
	if r.state != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "state", r.state, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetTokensRequest struct {
	ctx context.Context
	ApiService *OauthAPIService
	authorization *string
	grantType *string
	code *string
	redirectUri *string
}

// Base 64 encoded string containing the &#x60;client_id&#x60; and &#x60;client_secret&#x60; values. The header value should be &#x60;Basic &lt;base64(client_id:client_secret)&gt;&#x60;.
func (r ApiGetTokensRequest) Authorization(authorization string) ApiGetTokensRequest {
	r.authorization = &authorization
	return r
}

// Since you are trying to exchange an authorization code for a pair of tokens, you must use the value \\\&quot;authorization_code\\\&quot;
func (r ApiGetTokensRequest) GrantType(grantType string) ApiGetTokensRequest {
	r.grantType = &grantType
	return r
}

// The authorization code that you received after the user confirmed app installation
func (r ApiGetTokensRequest) Code(code string) ApiGetTokensRequest {
	r.code = &code
	return r
}

// The callback URL you provided when you registered your app
func (r ApiGetTokensRequest) RedirectUri(redirectUri string) ApiGetTokensRequest {
	r.redirectUri = &redirectUri
	return r
}

func (r ApiGetTokensRequest) Execute() (*GetTokensResponse200, *http.Response, error) {
	return r.ApiService.GetTokensExecute(r)
}

/*
GetTokens Getting the tokens

After the customer has confirmed the app installation, you will need to exchange the `authorization_code` to a pair of access and refresh tokens. Using an access token, you can access the user's data through the API.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetTokensRequest
*/
func (a *OauthAPIService) GetTokens(ctx context.Context) ApiGetTokensRequest {
	return ApiGetTokensRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetTokensResponse200
func (a *OauthAPIService) GetTokensExecute(r ApiGetTokensRequest) (*GetTokensResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetTokensResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthAPIService.GetTokens")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	if r.grantType != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "grant_type", r.grantType, "")
	}
	if r.code != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "code", r.code, "")
	}
	if r.redirectUri != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "redirect_uri", r.redirectUri, "")
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

type ApiRefreshTokensRequest struct {
	ctx context.Context
	ApiService *OauthAPIService
	authorization *string
	grantType *string
	refreshToken *string
}

// Base 64 encoded string containing the &#x60;client_id&#x60; and &#x60;client_secret&#x60; values. The header value should be &#x60;Basic &lt;base64(client_id:client_secret)&gt;&#x60;.
func (r ApiRefreshTokensRequest) Authorization(authorization string) ApiRefreshTokensRequest {
	r.authorization = &authorization
	return r
}

// Since you are to refresh your access_token, you must use the value \\\&quot;refresh_token\\\&quot;
func (r ApiRefreshTokensRequest) GrantType(grantType string) ApiRefreshTokensRequest {
	r.grantType = &grantType
	return r
}

// The refresh token that you received after you exchanged the authorization code
func (r ApiRefreshTokensRequest) RefreshToken(refreshToken string) ApiRefreshTokensRequest {
	r.refreshToken = &refreshToken
	return r
}

func (r ApiRefreshTokensRequest) Execute() (*GetTokensResponse200, *http.Response, error) {
	return r.ApiService.RefreshTokensExecute(r)
}

/*
RefreshTokens Refreshing the tokens

The `access_token` has a lifetime. After a period of time, which was returned to you in `expires_in` JSON property, the `access_token` will be invalid, and you can no longer use it to get data from our API. To refresh the `access_token`, you must use the `refresh_token`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRefreshTokensRequest
*/
func (a *OauthAPIService) RefreshTokens(ctx context.Context) ApiRefreshTokensRequest {
	return ApiRefreshTokensRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetTokensResponse200
func (a *OauthAPIService) RefreshTokensExecute(r ApiRefreshTokensRequest) (*GetTokensResponse200, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetTokensResponse200
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OauthAPIService.RefreshTokens")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/token/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

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
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	if r.grantType != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "grant_type", r.grantType, "")
	}
	if r.refreshToken != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "refresh_token", r.refreshToken, "")
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
