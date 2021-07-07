// Copyright 2021 Google LLC
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

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	resourcespb "github.com/dictav/go-genproto-googleads/pb/v8/resources"
	servicespb "github.com/dictav/go-genproto-googleads/pb/v8/services"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	statuspb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

var newCampaignExperimentClientHook clientHook

// CampaignExperimentCallOptions contains the retry settings for each method of CampaignExperimentClient.
type CampaignExperimentCallOptions struct {
	GetCampaignExperiment             []gax.CallOption
	CreateCampaignExperiment          []gax.CallOption
	MutateCampaignExperiments         []gax.CallOption
	GraduateCampaignExperiment        []gax.CallOption
	PromoteCampaignExperiment         []gax.CallOption
	EndCampaignExperiment             []gax.CallOption
	ListCampaignExperimentAsyncErrors []gax.CallOption
}

func defaultCampaignExperimentGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCampaignExperimentCallOptions() *CampaignExperimentCallOptions {
	return &CampaignExperimentCallOptions{
		GetCampaignExperiment: []gax.CallOption{
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
		CreateCampaignExperiment: []gax.CallOption{
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
		MutateCampaignExperiments: []gax.CallOption{
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
		GraduateCampaignExperiment: []gax.CallOption{
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
		PromoteCampaignExperiment: []gax.CallOption{
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
		EndCampaignExperiment: []gax.CallOption{
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
		ListCampaignExperimentAsyncErrors: []gax.CallOption{
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

// internalCampaignExperimentClient is an interface that defines the methods availaible from Google Ads API.
type internalCampaignExperimentClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetCampaignExperiment(context.Context, *servicespb.GetCampaignExperimentRequest, ...gax.CallOption) (*resourcespb.CampaignExperiment, error)
	CreateCampaignExperiment(context.Context, *servicespb.CreateCampaignExperimentRequest, ...gax.CallOption) (*CreateCampaignExperimentOperation, error)
	CreateCampaignExperimentOperation(name string) *CreateCampaignExperimentOperation
	MutateCampaignExperiments(context.Context, *servicespb.MutateCampaignExperimentsRequest, ...gax.CallOption) (*servicespb.MutateCampaignExperimentsResponse, error)
	GraduateCampaignExperiment(context.Context, *servicespb.GraduateCampaignExperimentRequest, ...gax.CallOption) (*servicespb.GraduateCampaignExperimentResponse, error)
	PromoteCampaignExperiment(context.Context, *servicespb.PromoteCampaignExperimentRequest, ...gax.CallOption) (*PromoteCampaignExperimentOperation, error)
	PromoteCampaignExperimentOperation(name string) *PromoteCampaignExperimentOperation
	EndCampaignExperiment(context.Context, *servicespb.EndCampaignExperimentRequest, ...gax.CallOption) error
	ListCampaignExperimentAsyncErrors(context.Context, *servicespb.ListCampaignExperimentAsyncErrorsRequest, ...gax.CallOption) *StatusIterator
}

// CampaignExperimentClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// CampaignExperimentService manages the life cycle of campaign experiments.
// It is used to create new experiments from drafts, modify experiment
// properties, promote changes in an experiment back to its base campaign,
// graduate experiments into new stand-alone campaigns, and to remove an
// experiment.
//
// An experiment consists of two variants or arms - the base campaign and the
// experiment campaign, directing a fixed share of traffic to each arm.
// A campaign experiment is created from a draft of changes to the base campaign
// and will be a snapshot of changes in the draft at the time of creation.
type CampaignExperimentClient struct {
	// The internal transport-dependent client.
	internalClient internalCampaignExperimentClient

	// The call options for this service.
	CallOptions *CampaignExperimentCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CampaignExperimentClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CampaignExperimentClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *CampaignExperimentClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetCampaignExperiment returns the requested campaign experiment in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *CampaignExperimentClient) GetCampaignExperiment(ctx context.Context, req *servicespb.GetCampaignExperimentRequest, opts ...gax.CallOption) (*resourcespb.CampaignExperiment, error) {
	return c.internalClient.GetCampaignExperiment(ctx, req, opts...)
}

// CreateCampaignExperiment creates a campaign experiment based on a campaign draft. The draft campaign
// will be forked into a real campaign (called the experiment campaign) that
// will begin serving ads if successfully created.
//
// The campaign experiment is created immediately with status INITIALIZING.
// This method return a long running operation that tracks the forking of the
// draft campaign. If the forking fails, a list of errors can be retrieved
// using the ListCampaignExperimentAsyncErrors method. The operation’s
// metadata will be a StringValue containing the resource name of the created
// campaign experiment.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CampaignExperimentError (at )
// DatabaseError (at )
// DateError (at )
// DateRangeError (at )
// FieldError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
func (c *CampaignExperimentClient) CreateCampaignExperiment(ctx context.Context, req *servicespb.CreateCampaignExperimentRequest, opts ...gax.CallOption) (*CreateCampaignExperimentOperation, error) {
	return c.internalClient.CreateCampaignExperiment(ctx, req, opts...)
}

// CreateCampaignExperimentOperation returns a new CreateCampaignExperimentOperation from a given name.
// The name must be that of a previously created CreateCampaignExperimentOperation, possibly from a different process.
func (c *CampaignExperimentClient) CreateCampaignExperimentOperation(name string) *CreateCampaignExperimentOperation {
	return c.internalClient.CreateCampaignExperimentOperation(name)
}

// MutateCampaignExperiments updates campaign experiments. Operation statuses are returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CampaignExperimentError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *CampaignExperimentClient) MutateCampaignExperiments(ctx context.Context, req *servicespb.MutateCampaignExperimentsRequest, opts ...gax.CallOption) (*servicespb.MutateCampaignExperimentsResponse, error) {
	return c.internalClient.MutateCampaignExperiments(ctx, req, opts...)
}

// GraduateCampaignExperiment graduates a campaign experiment to a full campaign. The base and experiment
// campaigns will start running independently with their own budgets.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CampaignExperimentError (at )
// HeaderError (at )
// InternalError (at )
// MutateError (at )
// QuotaError (at )
// RequestError (at )
func (c *CampaignExperimentClient) GraduateCampaignExperiment(ctx context.Context, req *servicespb.GraduateCampaignExperimentRequest, opts ...gax.CallOption) (*servicespb.GraduateCampaignExperimentResponse, error) {
	return c.internalClient.GraduateCampaignExperiment(ctx, req, opts...)
}

// PromoteCampaignExperiment promotes the changes in a experiment campaign back to the base campaign.
//
// The campaign experiment is updated immediately with status PROMOTING.
// This method return a long running operation that tracks the promoting of
// the experiment campaign. If the promoting fails, a list of errors can be
// retrieved using the ListCampaignExperimentAsyncErrors method.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *CampaignExperimentClient) PromoteCampaignExperiment(ctx context.Context, req *servicespb.PromoteCampaignExperimentRequest, opts ...gax.CallOption) (*PromoteCampaignExperimentOperation, error) {
	return c.internalClient.PromoteCampaignExperiment(ctx, req, opts...)
}

// PromoteCampaignExperimentOperation returns a new PromoteCampaignExperimentOperation from a given name.
// The name must be that of a previously created PromoteCampaignExperimentOperation, possibly from a different process.
func (c *CampaignExperimentClient) PromoteCampaignExperimentOperation(name string) *PromoteCampaignExperimentOperation {
	return c.internalClient.PromoteCampaignExperimentOperation(name)
}

// EndCampaignExperiment immediately ends a campaign experiment, changing the experiment’s scheduled
// end date and without waiting for end of day. End date is updated to be the
// time of the request.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CampaignExperimentError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *CampaignExperimentClient) EndCampaignExperiment(ctx context.Context, req *servicespb.EndCampaignExperimentRequest, opts ...gax.CallOption) error {
	return c.internalClient.EndCampaignExperiment(ctx, req, opts...)
}

// ListCampaignExperimentAsyncErrors returns all errors that occurred during CampaignExperiment create or
// promote (whichever occurred last).
// Supports standard list paging.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *CampaignExperimentClient) ListCampaignExperimentAsyncErrors(ctx context.Context, req *servicespb.ListCampaignExperimentAsyncErrorsRequest, opts ...gax.CallOption) *StatusIterator {
	return c.internalClient.ListCampaignExperimentAsyncErrors(ctx, req, opts...)
}

// campaignExperimentGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type campaignExperimentGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing CampaignExperimentClient
	CallOptions **CampaignExperimentCallOptions

	// The gRPC API client.
	campaignExperimentClient servicespb.CampaignExperimentServiceClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewCampaignExperimentClient creates a new campaign experiment service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// CampaignExperimentService manages the life cycle of campaign experiments.
// It is used to create new experiments from drafts, modify experiment
// properties, promote changes in an experiment back to its base campaign,
// graduate experiments into new stand-alone campaigns, and to remove an
// experiment.
//
// An experiment consists of two variants or arms - the base campaign and the
// experiment campaign, directing a fixed share of traffic to each arm.
// A campaign experiment is created from a draft of changes to the base campaign
// and will be a snapshot of changes in the draft at the time of creation.
func NewCampaignExperimentClient(ctx context.Context, opts ...option.ClientOption) (*CampaignExperimentClient, error) {
	clientOpts := defaultCampaignExperimentGRPCClientOptions()
	if newCampaignExperimentClientHook != nil {
		hookOpts, err := newCampaignExperimentClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CampaignExperimentClient{CallOptions: defaultCampaignExperimentCallOptions()}

	c := &campaignExperimentGRPCClient{
		connPool:                 connPool,
		disableDeadlines:         disableDeadlines,
		campaignExperimentClient: servicespb.NewCampaignExperimentServiceClient(connPool),
		CallOptions:              &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *campaignExperimentGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *campaignExperimentGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *campaignExperimentGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *campaignExperimentGRPCClient) GetCampaignExperiment(ctx context.Context, req *servicespb.GetCampaignExperimentRequest, opts ...gax.CallOption) (*resourcespb.CampaignExperiment, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetCampaignExperiment[0:len((*c.CallOptions).GetCampaignExperiment):len((*c.CallOptions).GetCampaignExperiment)], opts...)
	var resp *resourcespb.CampaignExperiment
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.campaignExperimentClient.GetCampaignExperiment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *campaignExperimentGRPCClient) CreateCampaignExperiment(ctx context.Context, req *servicespb.CreateCampaignExperimentRequest, opts ...gax.CallOption) (*CreateCampaignExperimentOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateCampaignExperiment[0:len((*c.CallOptions).CreateCampaignExperiment):len((*c.CallOptions).CreateCampaignExperiment)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.campaignExperimentClient.CreateCampaignExperiment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateCampaignExperimentOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *campaignExperimentGRPCClient) MutateCampaignExperiments(ctx context.Context, req *servicespb.MutateCampaignExperimentsRequest, opts ...gax.CallOption) (*servicespb.MutateCampaignExperimentsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateCampaignExperiments[0:len((*c.CallOptions).MutateCampaignExperiments):len((*c.CallOptions).MutateCampaignExperiments)], opts...)
	var resp *servicespb.MutateCampaignExperimentsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.campaignExperimentClient.MutateCampaignExperiments(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *campaignExperimentGRPCClient) GraduateCampaignExperiment(ctx context.Context, req *servicespb.GraduateCampaignExperimentRequest, opts ...gax.CallOption) (*servicespb.GraduateCampaignExperimentResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "campaign_experiment", url.QueryEscape(req.GetCampaignExperiment())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GraduateCampaignExperiment[0:len((*c.CallOptions).GraduateCampaignExperiment):len((*c.CallOptions).GraduateCampaignExperiment)], opts...)
	var resp *servicespb.GraduateCampaignExperimentResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.campaignExperimentClient.GraduateCampaignExperiment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *campaignExperimentGRPCClient) PromoteCampaignExperiment(ctx context.Context, req *servicespb.PromoteCampaignExperimentRequest, opts ...gax.CallOption) (*PromoteCampaignExperimentOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "campaign_experiment", url.QueryEscape(req.GetCampaignExperiment())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).PromoteCampaignExperiment[0:len((*c.CallOptions).PromoteCampaignExperiment):len((*c.CallOptions).PromoteCampaignExperiment)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.campaignExperimentClient.PromoteCampaignExperiment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &PromoteCampaignExperimentOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *campaignExperimentGRPCClient) EndCampaignExperiment(ctx context.Context, req *servicespb.EndCampaignExperimentRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "campaign_experiment", url.QueryEscape(req.GetCampaignExperiment())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).EndCampaignExperiment[0:len((*c.CallOptions).EndCampaignExperiment):len((*c.CallOptions).EndCampaignExperiment)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.campaignExperimentClient.EndCampaignExperiment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *campaignExperimentGRPCClient) ListCampaignExperimentAsyncErrors(ctx context.Context, req *servicespb.ListCampaignExperimentAsyncErrorsRequest, opts ...gax.CallOption) *StatusIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListCampaignExperimentAsyncErrors[0:len((*c.CallOptions).ListCampaignExperimentAsyncErrors):len((*c.CallOptions).ListCampaignExperimentAsyncErrors)], opts...)
	it := &StatusIterator{}
	req = proto.Clone(req).(*servicespb.ListCampaignExperimentAsyncErrorsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*statuspb.Status, string, error) {
		var resp *servicespb.ListCampaignExperimentAsyncErrorsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.campaignExperimentClient.ListCampaignExperimentAsyncErrors(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetErrors(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()
	return it
}

// CreateCampaignExperimentOperation manages a long-running operation from CreateCampaignExperiment.
type CreateCampaignExperimentOperation struct {
	lro *longrunning.Operation
}

// CreateCampaignExperimentOperation returns a new CreateCampaignExperimentOperation from a given name.
// The name must be that of a previously created CreateCampaignExperimentOperation, possibly from a different process.
func (c *campaignExperimentGRPCClient) CreateCampaignExperimentOperation(name string) *CreateCampaignExperimentOperation {
	return &CreateCampaignExperimentOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateCampaignExperimentOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.WaitWithInterval(ctx, nil, time.Minute, opts...)
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CreateCampaignExperimentOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CreateCampaignExperimentOperation) Metadata() (*servicespb.CreateCampaignExperimentMetadata, error) {
	var meta servicespb.CreateCampaignExperimentMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateCampaignExperimentOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateCampaignExperimentOperation) Name() string {
	return op.lro.Name()
}

// PromoteCampaignExperimentOperation manages a long-running operation from PromoteCampaignExperiment.
type PromoteCampaignExperimentOperation struct {
	lro *longrunning.Operation
}

// PromoteCampaignExperimentOperation returns a new PromoteCampaignExperimentOperation from a given name.
// The name must be that of a previously created PromoteCampaignExperimentOperation, possibly from a different process.
func (c *campaignExperimentGRPCClient) PromoteCampaignExperimentOperation(name string) *PromoteCampaignExperimentOperation {
	return &PromoteCampaignExperimentOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *PromoteCampaignExperimentOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.WaitWithInterval(ctx, nil, time.Minute, opts...)
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *PromoteCampaignExperimentOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *PromoteCampaignExperimentOperation) Metadata() (*emptypb.Empty, error) {
	var meta emptypb.Empty
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *PromoteCampaignExperimentOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *PromoteCampaignExperimentOperation) Name() string {
	return op.lro.Name()
}
