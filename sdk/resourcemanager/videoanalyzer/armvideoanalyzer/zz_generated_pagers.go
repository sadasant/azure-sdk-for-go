//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvideoanalyzer

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// AccessPoliciesClientListPager provides operations for iterating over paged responses.
type AccessPoliciesClientListPager struct {
	client    *AccessPoliciesClient
	current   AccessPoliciesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AccessPoliciesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AccessPoliciesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AccessPoliciesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AccessPolicyEntityCollection.NextLink == nil || len(*p.current.AccessPolicyEntityCollection.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AccessPoliciesClientListResponse page.
func (p *AccessPoliciesClientListPager) PageResponse() AccessPoliciesClientListResponse {
	return p.current
}

// EdgeModulesClientListPager provides operations for iterating over paged responses.
type EdgeModulesClientListPager struct {
	client    *EdgeModulesClient
	current   EdgeModulesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, EdgeModulesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *EdgeModulesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *EdgeModulesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.EdgeModuleEntityCollection.NextLink == nil || len(*p.current.EdgeModuleEntityCollection.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current EdgeModulesClientListResponse page.
func (p *EdgeModulesClientListPager) PageResponse() EdgeModulesClientListResponse {
	return p.current
}

// LivePipelinesClientListPager provides operations for iterating over paged responses.
type LivePipelinesClientListPager struct {
	client    *LivePipelinesClient
	current   LivePipelinesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, LivePipelinesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *LivePipelinesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *LivePipelinesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.LivePipelineCollection.NextLink == nil || len(*p.current.LivePipelineCollection.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current LivePipelinesClientListResponse page.
func (p *LivePipelinesClientListPager) PageResponse() LivePipelinesClientListResponse {
	return p.current
}

// PipelineJobsClientListPager provides operations for iterating over paged responses.
type PipelineJobsClientListPager struct {
	client    *PipelineJobsClient
	current   PipelineJobsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PipelineJobsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PipelineJobsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PipelineJobsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PipelineJobCollection.NextLink == nil || len(*p.current.PipelineJobCollection.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PipelineJobsClientListResponse page.
func (p *PipelineJobsClientListPager) PageResponse() PipelineJobsClientListResponse {
	return p.current
}

// PipelineTopologiesClientListPager provides operations for iterating over paged responses.
type PipelineTopologiesClientListPager struct {
	client    *PipelineTopologiesClient
	current   PipelineTopologiesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PipelineTopologiesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PipelineTopologiesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PipelineTopologiesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PipelineTopologyCollection.NextLink == nil || len(*p.current.PipelineTopologyCollection.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PipelineTopologiesClientListResponse page.
func (p *PipelineTopologiesClientListPager) PageResponse() PipelineTopologiesClientListResponse {
	return p.current
}

// VideosClientListPager provides operations for iterating over paged responses.
type VideosClientListPager struct {
	client    *VideosClient
	current   VideosClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VideosClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *VideosClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *VideosClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VideoEntityCollection.NextLink == nil || len(*p.current.VideoEntityCollection.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current VideosClientListResponse page.
func (p *VideosClientListPager) PageResponse() VideosClientListResponse {
	return p.current
}
