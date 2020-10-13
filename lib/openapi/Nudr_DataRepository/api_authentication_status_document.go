/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nudr_DataRepository

import (
	"free5gc/lib/openapi"
	"free5gc/lib/openapi/models"

	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type AuthenticationStatusDocumentApiService service

/*
AuthenticationStatusDocumentApiService To store the Authentication Status data of a UE
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ueId UE id
 * @param optional nil or *CreateAuthenticationStatusParamOpts - Optional Parameters:
 * @param "AuthEvent" (optional.Interface of AuthEvent) -
*/

type CreateAuthenticationStatusParamOpts struct {
	AuthEvent optional.Interface
}

func (a *AuthenticationStatusDocumentApiService) CreateAuthenticationStatus(ctx context.Context, ueId string, localVarOptionals *CreateAuthenticationStatusParamOpts) (*http.Response, error) {
	var (
		localVarHTTPMethod   = strings.ToUpper("Put")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	fmt.Printf("mow in the lib/.../api_status function\n")

	// create path and map variables
	localVarPath := a.client.cfg.BasePath() + "/subscription-data/{ueId}/authentication-data/authentication-status"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", fmt.Sprintf("%v", ueId), -1)

	localVarPath = "http://192.168.2.104:29504/nudr-dr/v1/subscription-data/imsi-2089300007487/authentication-data/authentication-status"

	fmt.Printf("localVarPath is %s\n",localVarPath)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarHTTPContentTypes := []string{"application/json"}

	localVarHeaderParams["Content-Type"] = localVarHTTPContentTypes[0] // use the first content type specified in 'consumes'

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := openapi.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	if localVarOptionals != nil && localVarOptionals.AuthEvent.IsSet() {
		localVarOptionalAuthEvent, localVarOptionalAuthEventok := localVarOptionals.AuthEvent.Value().(models.AuthEvent)
		if !localVarOptionalAuthEventok {
			return nil, openapi.ReportError("authEvent should be AuthEvent")
		}
		localVarPostBody = &localVarOptionalAuthEvent
	}

	r, err := openapi.PrepareRequest(ctx, a.client.cfg, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := openapi.CallAPI(a.client.cfg, r)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	apiError := openapi.GenericOpenAPIError{
		RawBody:     localVarBody,
		ErrorStatus: localVarHTTPResponse.Status,
	}

	switch localVarHTTPResponse.StatusCode {
	case 204:
		return localVarHTTPResponse, nil
	default:
		var v models.ProblemDetails
		err = openapi.Deserialize(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			apiError.ErrorStatus = err.Error()
			return localVarHTTPResponse, apiError
		}
		apiError.ErrorModel = v
		return localVarHTTPResponse, apiError
	}
}
