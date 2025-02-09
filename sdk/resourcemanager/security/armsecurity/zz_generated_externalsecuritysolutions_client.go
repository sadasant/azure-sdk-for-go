//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ExternalSecuritySolutionsClient contains the methods for the ExternalSecuritySolutions group.
// Don't use this type directly, use NewExternalSecuritySolutionsClient() instead.
type ExternalSecuritySolutionsClient struct {
	host           string
	subscriptionID string
	ascLocation    string
	pl             runtime.Pipeline
}

// NewExternalSecuritySolutionsClient creates a new instance of ExternalSecuritySolutionsClient with the specified values.
// subscriptionID - Azure subscription ID
// ascLocation - The location where ASC stores the data of the subscription. can be retrieved from Get locations
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewExternalSecuritySolutionsClient(subscriptionID string, ascLocation string, credential azcore.TokenCredential, options *arm.ClientOptions) *ExternalSecuritySolutionsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ExternalSecuritySolutionsClient{
		subscriptionID: subscriptionID,
		ascLocation:    ascLocation,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Gets a specific external Security Solution.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// externalSecuritySolutionsName - Name of an external security solution.
// options - ExternalSecuritySolutionsClientGetOptions contains the optional parameters for the ExternalSecuritySolutionsClient.Get
// method.
func (client *ExternalSecuritySolutionsClient) Get(ctx context.Context, resourceGroupName string, externalSecuritySolutionsName string, options *ExternalSecuritySolutionsClientGetOptions) (ExternalSecuritySolutionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, externalSecuritySolutionsName, options)
	if err != nil {
		return ExternalSecuritySolutionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExternalSecuritySolutionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExternalSecuritySolutionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExternalSecuritySolutionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, externalSecuritySolutionsName string, options *ExternalSecuritySolutionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/ExternalSecuritySolutions/{externalSecuritySolutionsName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.ascLocation == "" {
		return nil, errors.New("parameter client.ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(client.ascLocation))
	if externalSecuritySolutionsName == "" {
		return nil, errors.New("parameter externalSecuritySolutionsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{externalSecuritySolutionsName}", url.PathEscape(externalSecuritySolutionsName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExternalSecuritySolutionsClient) getHandleResponse(resp *http.Response) (ExternalSecuritySolutionsClientGetResponse, error) {
	result := ExternalSecuritySolutionsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExternalSecuritySolution); err != nil {
		return ExternalSecuritySolutionsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets a list of external security solutions for the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ExternalSecuritySolutionsClientListOptions contains the optional parameters for the ExternalSecuritySolutionsClient.List
// method.
func (client *ExternalSecuritySolutionsClient) List(options *ExternalSecuritySolutionsClientListOptions) *ExternalSecuritySolutionsClientListPager {
	return &ExternalSecuritySolutionsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ExternalSecuritySolutionsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExternalSecuritySolutionList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ExternalSecuritySolutionsClient) listCreateRequest(ctx context.Context, options *ExternalSecuritySolutionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/externalSecuritySolutions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExternalSecuritySolutionsClient) listHandleResponse(resp *http.Response) (ExternalSecuritySolutionsClientListResponse, error) {
	result := ExternalSecuritySolutionsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExternalSecuritySolutionList); err != nil {
		return ExternalSecuritySolutionsClientListResponse{}, err
	}
	return result, nil
}

// ListByHomeRegion - Gets a list of external Security Solutions for the subscription and location.
// If the operation fails it returns an *azcore.ResponseError type.
// options - ExternalSecuritySolutionsClientListByHomeRegionOptions contains the optional parameters for the ExternalSecuritySolutionsClient.ListByHomeRegion
// method.
func (client *ExternalSecuritySolutionsClient) ListByHomeRegion(options *ExternalSecuritySolutionsClientListByHomeRegionOptions) *ExternalSecuritySolutionsClientListByHomeRegionPager {
	return &ExternalSecuritySolutionsClientListByHomeRegionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByHomeRegionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ExternalSecuritySolutionsClientListByHomeRegionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExternalSecuritySolutionList.NextLink)
		},
	}
}

// listByHomeRegionCreateRequest creates the ListByHomeRegion request.
func (client *ExternalSecuritySolutionsClient) listByHomeRegionCreateRequest(ctx context.Context, options *ExternalSecuritySolutionsClientListByHomeRegionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/ExternalSecuritySolutions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if client.ascLocation == "" {
		return nil, errors.New("parameter client.ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(client.ascLocation))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByHomeRegionHandleResponse handles the ListByHomeRegion response.
func (client *ExternalSecuritySolutionsClient) listByHomeRegionHandleResponse(resp *http.Response) (ExternalSecuritySolutionsClientListByHomeRegionResponse, error) {
	result := ExternalSecuritySolutionsClientListByHomeRegionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExternalSecuritySolutionList); err != nil {
		return ExternalSecuritySolutionsClientListByHomeRegionResponse{}, err
	}
	return result, nil
}
