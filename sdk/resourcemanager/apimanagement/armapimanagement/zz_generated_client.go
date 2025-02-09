//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// Client contains the methods for the APIManagementClient group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewClient creates a new instance of Client with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *Client {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &Client{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginPerformConnectivityCheckAsync - Performs a connectivity check between the API Management service and a given destination,
// and returns metrics for the connection, as well as errors encountered while trying to establish it.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// connectivityCheckRequestParams - Connectivity Check request parameters.
// options - ClientBeginPerformConnectivityCheckAsyncOptions contains the optional parameters for the Client.BeginPerformConnectivityCheckAsync
// method.
func (client *Client) BeginPerformConnectivityCheckAsync(ctx context.Context, resourceGroupName string, serviceName string, connectivityCheckRequestParams ConnectivityCheckRequest, options *ClientBeginPerformConnectivityCheckAsyncOptions) (ClientPerformConnectivityCheckAsyncPollerResponse, error) {
	resp, err := client.performConnectivityCheckAsync(ctx, resourceGroupName, serviceName, connectivityCheckRequestParams, options)
	if err != nil {
		return ClientPerformConnectivityCheckAsyncPollerResponse{}, err
	}
	result := ClientPerformConnectivityCheckAsyncPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("Client.PerformConnectivityCheckAsync", "location", resp, client.pl)
	if err != nil {
		return ClientPerformConnectivityCheckAsyncPollerResponse{}, err
	}
	result.Poller = &ClientPerformConnectivityCheckAsyncPoller{
		pt: pt,
	}
	return result, nil
}

// PerformConnectivityCheckAsync - Performs a connectivity check between the API Management service and a given destination,
// and returns metrics for the connection, as well as errors encountered while trying to establish it.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *Client) performConnectivityCheckAsync(ctx context.Context, resourceGroupName string, serviceName string, connectivityCheckRequestParams ConnectivityCheckRequest, options *ClientBeginPerformConnectivityCheckAsyncOptions) (*http.Response, error) {
	req, err := client.performConnectivityCheckAsyncCreateRequest(ctx, resourceGroupName, serviceName, connectivityCheckRequestParams, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// performConnectivityCheckAsyncCreateRequest creates the PerformConnectivityCheckAsync request.
func (client *Client) performConnectivityCheckAsyncCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, connectivityCheckRequestParams ConnectivityCheckRequest, options *ClientBeginPerformConnectivityCheckAsyncOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/connectivityCheck"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, connectivityCheckRequestParams)
}
