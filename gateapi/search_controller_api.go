/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
)

// Linger please
var (
	_ context.Context
)

type SearchControllerApiService service


/* SearchControllerApiService Search infrastructure
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param type_ type
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "q" (string) q
     @param "platform" (string) platform
     @param "pageSize" (int32) pageSize
     @param "page" (int32) page
     @param "allowShortQuery" (bool) allowShortQuery
     @param "xRateLimitApp" (string) X-RateLimit-App
 @return []HashMap*/
func (a *SearchControllerApiService) SearchUsingGET(ctx context.Context, type_ string, localVarOptionals map[string]interface{}) ([]HashMap,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  []HashMap
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["q"], "string", "q"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["platform"], "string", "platform"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["pageSize"], "int32", "pageSize"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["page"], "int32", "page"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["allowShortQuery"], "bool", "allowShortQuery"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xRateLimitApp"], "string", "xRateLimitApp"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["q"].(string); localVarOk {
		localVarQueryParams.Add("q", parameterToString(localVarTempParam, ""))
	}
	localVarQueryParams.Add("type", parameterToString(type_, ""))
	if localVarTempParam, localVarOk := localVarOptionals["platform"].(string); localVarOk {
		localVarQueryParams.Add("platform", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["pageSize"].(int32); localVarOk {
		localVarQueryParams.Add("pageSize", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["page"].(int32); localVarOk {
		localVarQueryParams.Add("page", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["allowShortQuery"].(bool); localVarOk {
		localVarQueryParams.Add("allowShortQuery", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"*/*",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarTempParam, localVarOk := localVarOptionals["xRateLimitApp"].(string); localVarOk {
		localVarHeaderParams["X-RateLimit-App"] = parameterToString(localVarTempParam, "")
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

