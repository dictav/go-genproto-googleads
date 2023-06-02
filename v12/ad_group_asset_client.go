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
	servicespb "github.com/dictav/go-genproto-googleads/pb/v12/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var newAdGroupAssetClientHook clientHook

// AdGroupAssetCallOptions contains the retry settings for each method of AdGroupAssetClient.
type AdGroupAssetCallOptions struct {
	MutateAdGroupAssets []gax.CallOption
}

func defaultAdGroupAssetGRPCClientOptions() []option.ClientOption {
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

func defaultAdGroupAssetCallOptions() *AdGroupAssetCallOptions {
	return &AdGroupAssetCallOptions{
		MutateAdGroupAssets: []gax.CallOption{
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

// internalAdGroupAssetClient is an interface that defines the methods available from Google Ads API.
type internalAdGroupAssetClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	MutateAdGroupAssets(context.Context, *servicespb.MutateAdGroupAssetsRequest, ...gax.CallOption) (*servicespb.MutateAdGroupAssetsResponse, error)
}

// AdGroupAssetClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage ad group assets.
type AdGroupAssetClient struct {
	// The internal transport-dependent client.
	internalClient internalAdGroupAssetClient

	// The call options for this service.
	CallOptions *AdGroupAssetCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AdGroupAssetClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AdGroupAssetClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *AdGroupAssetClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// MutateAdGroupAssets creates, updates, or removes ad group assets. Operation statuses are
// returned.
//
// List of thrown errors:
// AssetLinkError (at )
// AuthenticationError (at )
// AuthorizationError (at )
// ContextError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// NotAllowlistedError (at )
// QuotaError (at )
// RequestError (at )
func (c *AdGroupAssetClient) MutateAdGroupAssets(ctx context.Context, req *servicespb.MutateAdGroupAssetsRequest, opts ...gax.CallOption) (*servicespb.MutateAdGroupAssetsResponse, error) {
	return c.internalClient.MutateAdGroupAssets(ctx, req, opts...)
}

// adGroupAssetGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type adGroupAssetGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing AdGroupAssetClient
	CallOptions **AdGroupAssetCallOptions

	// The gRPC API client.
	adGroupAssetClient servicespb.AdGroupAssetServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewAdGroupAssetClient creates a new ad group asset service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage ad group assets.
func NewAdGroupAssetClient(ctx context.Context, opts ...option.ClientOption) (*AdGroupAssetClient, error) {
	clientOpts := defaultAdGroupAssetGRPCClientOptions()
	if newAdGroupAssetClientHook != nil {
		hookOpts, err := newAdGroupAssetClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := AdGroupAssetClient{CallOptions: defaultAdGroupAssetCallOptions()}

	c := &adGroupAssetGRPCClient{
		connPool:           connPool,
		adGroupAssetClient: servicespb.NewAdGroupAssetServiceClient(connPool),
		CallOptions:        &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *adGroupAssetGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *adGroupAssetGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *adGroupAssetGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *adGroupAssetGRPCClient) MutateAdGroupAssets(ctx context.Context, req *servicespb.MutateAdGroupAssetsRequest, opts ...gax.CallOption) (*servicespb.MutateAdGroupAssetsResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).MutateAdGroupAssets[0:len((*c.CallOptions).MutateAdGroupAssets):len((*c.CallOptions).MutateAdGroupAssets)], opts...)
	var resp *servicespb.MutateAdGroupAssetsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.adGroupAssetClient.MutateAdGroupAssets(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
