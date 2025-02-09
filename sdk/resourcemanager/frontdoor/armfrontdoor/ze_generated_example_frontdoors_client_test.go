//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfrontdoor_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/frontdoor/armfrontdoor"
)

// x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorListAll.json
func ExampleFrontDoorsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armfrontdoor.NewFrontDoorsClient("<subscription-id>", cred, nil)
	pager := client.List(nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorList.json
func ExampleFrontDoorsClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armfrontdoor.NewFrontDoorsClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		nil)
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorGet.json
func ExampleFrontDoorsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armfrontdoor.NewFrontDoorsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<front-door-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.FrontDoorsClientGetResult)
}

// x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorCreate.json
func ExampleFrontDoorsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armfrontdoor.NewFrontDoorsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<front-door-name>",
		armfrontdoor.FrontDoor{
			Location: to.StringPtr("<location>"),
			Tags: map[string]*string{
				"tag1": to.StringPtr("value1"),
				"tag2": to.StringPtr("value2"),
			},
			Properties: &armfrontdoor.Properties{
				BackendPools: []*armfrontdoor.BackendPool{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armfrontdoor.BackendPoolProperties{
							Backends: []*armfrontdoor.Backend{
								{
									Address:   to.StringPtr("<address>"),
									HTTPPort:  to.Int32Ptr(80),
									HTTPSPort: to.Int32Ptr(443),
									Priority:  to.Int32Ptr(2),
									Weight:    to.Int32Ptr(1),
								},
								{
									Address:                    to.StringPtr("<address>"),
									HTTPPort:                   to.Int32Ptr(80),
									HTTPSPort:                  to.Int32Ptr(443),
									Priority:                   to.Int32Ptr(1),
									PrivateLinkApprovalMessage: to.StringPtr("<private-link-approval-message>"),
									PrivateLinkLocation:        to.StringPtr("<private-link-location>"),
									PrivateLinkResourceID:      to.StringPtr("<private-link-resource-id>"),
									Weight:                     to.Int32Ptr(2),
								},
								{
									Address:                    to.StringPtr("<address>"),
									HTTPPort:                   to.Int32Ptr(80),
									HTTPSPort:                  to.Int32Ptr(443),
									Priority:                   to.Int32Ptr(1),
									PrivateLinkAlias:           to.StringPtr("<private-link-alias>"),
									PrivateLinkApprovalMessage: to.StringPtr("<private-link-approval-message>"),
									Weight:                     to.Int32Ptr(1),
								}},
							HealthProbeSettings: &armfrontdoor.SubResource{
								ID: to.StringPtr("<id>"),
							},
							LoadBalancingSettings: &armfrontdoor.SubResource{
								ID: to.StringPtr("<id>"),
							},
						},
					}},
				BackendPoolsSettings: &armfrontdoor.BackendPoolsSettings{
					EnforceCertificateNameCheck: armfrontdoor.EnforceCertificateNameCheckEnabledState("Enabled").ToPtr(),
					SendRecvTimeoutSeconds:      to.Int32Ptr(60),
				},
				EnabledState: armfrontdoor.FrontDoorEnabledState("Enabled").ToPtr(),
				FrontendEndpoints: []*armfrontdoor.FrontendEndpoint{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armfrontdoor.FrontendEndpointProperties{
							HostName:                    to.StringPtr("<host-name>"),
							SessionAffinityEnabledState: armfrontdoor.SessionAffinityEnabledState("Enabled").ToPtr(),
							SessionAffinityTTLSeconds:   to.Int32Ptr(60),
							WebApplicationFirewallPolicyLink: &armfrontdoor.FrontendEndpointUpdateParametersWebApplicationFirewallPolicyLink{
								ID: to.StringPtr("<id>"),
							},
						},
					},
					{
						Name: to.StringPtr("<name>"),
						Properties: &armfrontdoor.FrontendEndpointProperties{
							HostName: to.StringPtr("<host-name>"),
						},
					}},
				HealthProbeSettings: []*armfrontdoor.HealthProbeSettingsModel{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armfrontdoor.HealthProbeSettingsProperties{
							Path:              to.StringPtr("<path>"),
							EnabledState:      armfrontdoor.HealthProbeEnabled("Enabled").ToPtr(),
							HealthProbeMethod: armfrontdoor.FrontDoorHealthProbeMethod("HEAD").ToPtr(),
							IntervalInSeconds: to.Int32Ptr(120),
							Protocol:          armfrontdoor.FrontDoorProtocol("Http").ToPtr(),
						},
					}},
				LoadBalancingSettings: []*armfrontdoor.LoadBalancingSettingsModel{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armfrontdoor.LoadBalancingSettingsProperties{
							SampleSize:                to.Int32Ptr(4),
							SuccessfulSamplesRequired: to.Int32Ptr(2),
						},
					}},
				RoutingRules: []*armfrontdoor.RoutingRule{
					{
						Name: to.StringPtr("<name>"),
						Properties: &armfrontdoor.RoutingRuleProperties{
							AcceptedProtocols: []*armfrontdoor.FrontDoorProtocol{
								armfrontdoor.FrontDoorProtocol("Http").ToPtr()},
							EnabledState: armfrontdoor.RoutingRuleEnabledState("Enabled").ToPtr(),
							FrontendEndpoints: []*armfrontdoor.SubResource{
								{
									ID: to.StringPtr("<id>"),
								},
								{
									ID: to.StringPtr("<id>"),
								}},
							PatternsToMatch: []*string{
								to.StringPtr("/*")},
							RouteConfiguration: &armfrontdoor.ForwardingConfiguration{
								ODataType: to.StringPtr("<odata-type>"),
								BackendPool: &armfrontdoor.SubResource{
									ID: to.StringPtr("<id>"),
								},
							},
							RulesEngine: &armfrontdoor.SubResource{
								ID: to.StringPtr("<id>"),
							},
							WebApplicationFirewallPolicyLink: &armfrontdoor.RoutingRuleUpdateParametersWebApplicationFirewallPolicyLink{
								ID: to.StringPtr("<id>"),
							},
						},
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.FrontDoorsClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorDelete.json
func ExampleFrontDoorsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armfrontdoor.NewFrontDoorsClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<front-door-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2020-05-01/examples/FrontdoorValidateCustomDomain.json
func ExampleFrontDoorsClient_ValidateCustomDomain() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armfrontdoor.NewFrontDoorsClient("<subscription-id>", cred, nil)
	res, err := client.ValidateCustomDomain(ctx,
		"<resource-group-name>",
		"<front-door-name>",
		armfrontdoor.ValidateCustomDomainInput{
			HostName: to.StringPtr("<host-name>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.FrontDoorsClientValidateCustomDomainResult)
}
