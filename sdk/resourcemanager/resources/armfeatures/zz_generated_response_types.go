//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfeatures

import "net/http"

// ClientGetResponse contains the response from method Client.Get.
type ClientGetResponse struct {
	ClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClientGetResult contains the result from method Client.Get.
type ClientGetResult struct {
	FeatureResult
}

// ClientListAllResponse contains the response from method Client.ListAll.
type ClientListAllResponse struct {
	ClientListAllResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClientListAllResult contains the result from method Client.ListAll.
type ClientListAllResult struct {
	FeatureOperationsListResult
}

// ClientListResponse contains the response from method Client.List.
type ClientListResponse struct {
	ClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClientListResult contains the result from method Client.List.
type ClientListResult struct {
	FeatureOperationsListResult
}

// ClientRegisterResponse contains the response from method Client.Register.
type ClientRegisterResponse struct {
	ClientRegisterResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClientRegisterResult contains the result from method Client.Register.
type ClientRegisterResult struct {
	FeatureResult
}

// ClientUnregisterResponse contains the response from method Client.Unregister.
type ClientUnregisterResponse struct {
	ClientUnregisterResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ClientUnregisterResult contains the result from method Client.Unregister.
type ClientUnregisterResult struct {
	FeatureResult
}

// FeatureClientListOperationsResponse contains the response from method FeatureClient.ListOperations.
type FeatureClientListOperationsResponse struct {
	FeatureClientListOperationsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// FeatureClientListOperationsResult contains the result from method FeatureClient.ListOperations.
type FeatureClientListOperationsResult struct {
	OperationListResult
}

// SubscriptionFeatureRegistrationsClientCreateOrUpdateResponse contains the response from method SubscriptionFeatureRegistrationsClient.CreateOrUpdate.
type SubscriptionFeatureRegistrationsClientCreateOrUpdateResponse struct {
	SubscriptionFeatureRegistrationsClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionFeatureRegistrationsClientCreateOrUpdateResult contains the result from method SubscriptionFeatureRegistrationsClient.CreateOrUpdate.
type SubscriptionFeatureRegistrationsClientCreateOrUpdateResult struct {
	SubscriptionFeatureRegistration
}

// SubscriptionFeatureRegistrationsClientDeleteResponse contains the response from method SubscriptionFeatureRegistrationsClient.Delete.
type SubscriptionFeatureRegistrationsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionFeatureRegistrationsClientGetResponse contains the response from method SubscriptionFeatureRegistrationsClient.Get.
type SubscriptionFeatureRegistrationsClientGetResponse struct {
	SubscriptionFeatureRegistrationsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionFeatureRegistrationsClientGetResult contains the result from method SubscriptionFeatureRegistrationsClient.Get.
type SubscriptionFeatureRegistrationsClientGetResult struct {
	SubscriptionFeatureRegistration
}

// SubscriptionFeatureRegistrationsClientListAllBySubscriptionResponse contains the response from method SubscriptionFeatureRegistrationsClient.ListAllBySubscription.
type SubscriptionFeatureRegistrationsClientListAllBySubscriptionResponse struct {
	SubscriptionFeatureRegistrationsClientListAllBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionFeatureRegistrationsClientListAllBySubscriptionResult contains the result from method SubscriptionFeatureRegistrationsClient.ListAllBySubscription.
type SubscriptionFeatureRegistrationsClientListAllBySubscriptionResult struct {
	SubscriptionFeatureRegistrationList
}

// SubscriptionFeatureRegistrationsClientListBySubscriptionResponse contains the response from method SubscriptionFeatureRegistrationsClient.ListBySubscription.
type SubscriptionFeatureRegistrationsClientListBySubscriptionResponse struct {
	SubscriptionFeatureRegistrationsClientListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubscriptionFeatureRegistrationsClientListBySubscriptionResult contains the result from method SubscriptionFeatureRegistrationsClient.ListBySubscription.
type SubscriptionFeatureRegistrationsClientListBySubscriptionResult struct {
	SubscriptionFeatureRegistrationList
}
