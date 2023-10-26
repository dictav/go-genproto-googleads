// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package googleads

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	servicespb "github.com/dictav/go-genproto-googleads/pb/v14/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var newAssetGroupClientHook clientHook

// AssetGroupCallOptions contains the retry settings for each method of AssetGroupClient.
type AssetGroupCallOptions struct {
	MutateAssetGroups []gax.CallOption
}

func defaultAssetGroupGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultAssetGroupCallOptions() *AssetGroupCallOptions {
	return &AssetGroupCallOptions{
		MutateAssetGroups: []gax.CallOption{
			gax.WithTimeout(14400000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalAssetGroupClient is an interface that defines the methods available from Google Ads API.
type internalAssetGroupClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateAssetGroups(context.Context, *servicespb.MutateAssetGroupsRequest, ...gax.CallOption) (*servicespb.MutateAssetGroupsResponse, error)
}

// AssetGroupClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage asset group
type AssetGroupClient struct {
	// The internal transport-dependent client.
	internalClient internalAssetGroupClient

	// The call options for this service.
	CallOptions *AssetGroupCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AssetGroupClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AssetGroupClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *AssetGroupClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateAssetGroups creates, updates or removes asset groups. Operation statuses are
// returned.
func (c *AssetGroupClient) MutateAssetGroups(ctx context.Context, req *servicespb.MutateAssetGroupsRequest, opts ...gax.CallOption) (*servicespb.MutateAssetGroupsResponse, error) {
	return c.internalClient.MutateAssetGroups(ctx, req, opts...)
}

// assetGroupGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type assetGroupGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing AssetGroupClient
	CallOptions **AssetGroupCallOptions

	// The gRPC API client.
	assetGroupClient servicespb.AssetGroupServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewAssetGroupClient creates a new asset group service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage asset group
func NewAssetGroupClient(ctx context.Context, opts ...option.ClientOption) (*AssetGroupClient, error) {
	clientOpts := defaultAssetGroupGRPCClientOptions()
	if newAssetGroupClientHook != nil {
		hookOpts, err := newAssetGroupClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := AssetGroupClient{CallOptions: defaultAssetGroupCallOptions()}

	c := &assetGroupGRPCClient{
		connPool:         connPool,
		assetGroupClient: servicespb.NewAssetGroupServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *assetGroupGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *assetGroupGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *assetGroupGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *assetGroupGRPCClient) MutateAssetGroups(ctx context.Context, req *servicespb.MutateAssetGroupsRequest, opts ...gax.CallOption) (*servicespb.MutateAssetGroupsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateAssetGroups[0:len((*c.CallOptions).MutateAssetGroups):len((*c.CallOptions).MutateAssetGroups)], opts...)
	var resp *servicespb.MutateAssetGroupsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.assetGroupClient.MutateAssetGroups(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
