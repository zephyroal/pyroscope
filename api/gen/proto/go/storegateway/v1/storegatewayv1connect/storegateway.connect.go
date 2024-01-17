// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: storegateway/v1/storegateway.proto

package storegatewayv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v11 "github.com/grafana/pyroscope/api/gen/proto/go/ingester/v1"
	v1 "github.com/grafana/pyroscope/api/gen/proto/go/storegateway/v1"
	v12 "github.com/grafana/pyroscope/api/gen/proto/go/types/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// StoreGatewayServiceName is the fully-qualified name of the StoreGatewayService service.
	StoreGatewayServiceName = "storegateway.v1.StoreGatewayService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// StoreGatewayServiceMergeProfilesStacktracesProcedure is the fully-qualified name of the
	// StoreGatewayService's MergeProfilesStacktraces RPC.
	StoreGatewayServiceMergeProfilesStacktracesProcedure = "/storegateway.v1.StoreGatewayService/MergeProfilesStacktraces"
	// StoreGatewayServiceMergeProfilesLabelsProcedure is the fully-qualified name of the
	// StoreGatewayService's MergeProfilesLabels RPC.
	StoreGatewayServiceMergeProfilesLabelsProcedure = "/storegateway.v1.StoreGatewayService/MergeProfilesLabels"
	// StoreGatewayServiceMergeProfilesPprofProcedure is the fully-qualified name of the
	// StoreGatewayService's MergeProfilesPprof RPC.
	StoreGatewayServiceMergeProfilesPprofProcedure = "/storegateway.v1.StoreGatewayService/MergeProfilesPprof"
	// StoreGatewayServiceMergeSpanProfileProcedure is the fully-qualified name of the
	// StoreGatewayService's MergeSpanProfile RPC.
	StoreGatewayServiceMergeSpanProfileProcedure = "/storegateway.v1.StoreGatewayService/MergeSpanProfile"
	// StoreGatewayServiceProfileTypesProcedure is the fully-qualified name of the StoreGatewayService's
	// ProfileTypes RPC.
	StoreGatewayServiceProfileTypesProcedure = "/storegateway.v1.StoreGatewayService/ProfileTypes"
	// StoreGatewayServiceLabelValuesProcedure is the fully-qualified name of the StoreGatewayService's
	// LabelValues RPC.
	StoreGatewayServiceLabelValuesProcedure = "/storegateway.v1.StoreGatewayService/LabelValues"
	// StoreGatewayServiceLabelNamesProcedure is the fully-qualified name of the StoreGatewayService's
	// LabelNames RPC.
	StoreGatewayServiceLabelNamesProcedure = "/storegateway.v1.StoreGatewayService/LabelNames"
	// StoreGatewayServiceSeriesProcedure is the fully-qualified name of the StoreGatewayService's
	// Series RPC.
	StoreGatewayServiceSeriesProcedure = "/storegateway.v1.StoreGatewayService/Series"
	// StoreGatewayServiceBlockMetadataProcedure is the fully-qualified name of the
	// StoreGatewayService's BlockMetadata RPC.
	StoreGatewayServiceBlockMetadataProcedure = "/storegateway.v1.StoreGatewayService/BlockMetadata"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	storeGatewayServiceServiceDescriptor                        = v1.File_storegateway_v1_storegateway_proto.Services().ByName("StoreGatewayService")
	storeGatewayServiceMergeProfilesStacktracesMethodDescriptor = storeGatewayServiceServiceDescriptor.Methods().ByName("MergeProfilesStacktraces")
	storeGatewayServiceMergeProfilesLabelsMethodDescriptor      = storeGatewayServiceServiceDescriptor.Methods().ByName("MergeProfilesLabels")
	storeGatewayServiceMergeProfilesPprofMethodDescriptor       = storeGatewayServiceServiceDescriptor.Methods().ByName("MergeProfilesPprof")
	storeGatewayServiceMergeSpanProfileMethodDescriptor         = storeGatewayServiceServiceDescriptor.Methods().ByName("MergeSpanProfile")
	storeGatewayServiceProfileTypesMethodDescriptor             = storeGatewayServiceServiceDescriptor.Methods().ByName("ProfileTypes")
	storeGatewayServiceLabelValuesMethodDescriptor              = storeGatewayServiceServiceDescriptor.Methods().ByName("LabelValues")
	storeGatewayServiceLabelNamesMethodDescriptor               = storeGatewayServiceServiceDescriptor.Methods().ByName("LabelNames")
	storeGatewayServiceSeriesMethodDescriptor                   = storeGatewayServiceServiceDescriptor.Methods().ByName("Series")
	storeGatewayServiceBlockMetadataMethodDescriptor            = storeGatewayServiceServiceDescriptor.Methods().ByName("BlockMetadata")
)

// StoreGatewayServiceClient is a client for the storegateway.v1.StoreGatewayService service.
type StoreGatewayServiceClient interface {
	MergeProfilesStacktraces(context.Context) *connect.BidiStreamForClient[v11.MergeProfilesStacktracesRequest, v11.MergeProfilesStacktracesResponse]
	MergeProfilesLabels(context.Context) *connect.BidiStreamForClient[v11.MergeProfilesLabelsRequest, v11.MergeProfilesLabelsResponse]
	MergeProfilesPprof(context.Context) *connect.BidiStreamForClient[v11.MergeProfilesPprofRequest, v11.MergeProfilesPprofResponse]
	MergeSpanProfile(context.Context) *connect.BidiStreamForClient[v11.MergeSpanProfileRequest, v11.MergeSpanProfileResponse]
	// Deprecated: ProfileType call is deprecated in the store components
	// TODO: Remove this call in release v1.4
	ProfileTypes(context.Context, *connect.Request[v11.ProfileTypesRequest]) (*connect.Response[v11.ProfileTypesResponse], error)
	LabelValues(context.Context, *connect.Request[v12.LabelValuesRequest]) (*connect.Response[v12.LabelValuesResponse], error)
	LabelNames(context.Context, *connect.Request[v12.LabelNamesRequest]) (*connect.Response[v12.LabelNamesResponse], error)
	Series(context.Context, *connect.Request[v11.SeriesRequest]) (*connect.Response[v11.SeriesResponse], error)
	BlockMetadata(context.Context, *connect.Request[v11.BlockMetadataRequest]) (*connect.Response[v11.BlockMetadataResponse], error)
}

// NewStoreGatewayServiceClient constructs a client for the storegateway.v1.StoreGatewayService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewStoreGatewayServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) StoreGatewayServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &storeGatewayServiceClient{
		mergeProfilesStacktraces: connect.NewClient[v11.MergeProfilesStacktracesRequest, v11.MergeProfilesStacktracesResponse](
			httpClient,
			baseURL+StoreGatewayServiceMergeProfilesStacktracesProcedure,
			connect.WithSchema(storeGatewayServiceMergeProfilesStacktracesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		mergeProfilesLabels: connect.NewClient[v11.MergeProfilesLabelsRequest, v11.MergeProfilesLabelsResponse](
			httpClient,
			baseURL+StoreGatewayServiceMergeProfilesLabelsProcedure,
			connect.WithSchema(storeGatewayServiceMergeProfilesLabelsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		mergeProfilesPprof: connect.NewClient[v11.MergeProfilesPprofRequest, v11.MergeProfilesPprofResponse](
			httpClient,
			baseURL+StoreGatewayServiceMergeProfilesPprofProcedure,
			connect.WithSchema(storeGatewayServiceMergeProfilesPprofMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		mergeSpanProfile: connect.NewClient[v11.MergeSpanProfileRequest, v11.MergeSpanProfileResponse](
			httpClient,
			baseURL+StoreGatewayServiceMergeSpanProfileProcedure,
			connect.WithSchema(storeGatewayServiceMergeSpanProfileMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		profileTypes: connect.NewClient[v11.ProfileTypesRequest, v11.ProfileTypesResponse](
			httpClient,
			baseURL+StoreGatewayServiceProfileTypesProcedure,
			connect.WithSchema(storeGatewayServiceProfileTypesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		labelValues: connect.NewClient[v12.LabelValuesRequest, v12.LabelValuesResponse](
			httpClient,
			baseURL+StoreGatewayServiceLabelValuesProcedure,
			connect.WithSchema(storeGatewayServiceLabelValuesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		labelNames: connect.NewClient[v12.LabelNamesRequest, v12.LabelNamesResponse](
			httpClient,
			baseURL+StoreGatewayServiceLabelNamesProcedure,
			connect.WithSchema(storeGatewayServiceLabelNamesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		series: connect.NewClient[v11.SeriesRequest, v11.SeriesResponse](
			httpClient,
			baseURL+StoreGatewayServiceSeriesProcedure,
			connect.WithSchema(storeGatewayServiceSeriesMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		blockMetadata: connect.NewClient[v11.BlockMetadataRequest, v11.BlockMetadataResponse](
			httpClient,
			baseURL+StoreGatewayServiceBlockMetadataProcedure,
			connect.WithSchema(storeGatewayServiceBlockMetadataMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// storeGatewayServiceClient implements StoreGatewayServiceClient.
type storeGatewayServiceClient struct {
	mergeProfilesStacktraces *connect.Client[v11.MergeProfilesStacktracesRequest, v11.MergeProfilesStacktracesResponse]
	mergeProfilesLabels      *connect.Client[v11.MergeProfilesLabelsRequest, v11.MergeProfilesLabelsResponse]
	mergeProfilesPprof       *connect.Client[v11.MergeProfilesPprofRequest, v11.MergeProfilesPprofResponse]
	mergeSpanProfile         *connect.Client[v11.MergeSpanProfileRequest, v11.MergeSpanProfileResponse]
	profileTypes             *connect.Client[v11.ProfileTypesRequest, v11.ProfileTypesResponse]
	labelValues              *connect.Client[v12.LabelValuesRequest, v12.LabelValuesResponse]
	labelNames               *connect.Client[v12.LabelNamesRequest, v12.LabelNamesResponse]
	series                   *connect.Client[v11.SeriesRequest, v11.SeriesResponse]
	blockMetadata            *connect.Client[v11.BlockMetadataRequest, v11.BlockMetadataResponse]
}

// MergeProfilesStacktraces calls storegateway.v1.StoreGatewayService.MergeProfilesStacktraces.
func (c *storeGatewayServiceClient) MergeProfilesStacktraces(ctx context.Context) *connect.BidiStreamForClient[v11.MergeProfilesStacktracesRequest, v11.MergeProfilesStacktracesResponse] {
	return c.mergeProfilesStacktraces.CallBidiStream(ctx)
}

// MergeProfilesLabels calls storegateway.v1.StoreGatewayService.MergeProfilesLabels.
func (c *storeGatewayServiceClient) MergeProfilesLabels(ctx context.Context) *connect.BidiStreamForClient[v11.MergeProfilesLabelsRequest, v11.MergeProfilesLabelsResponse] {
	return c.mergeProfilesLabels.CallBidiStream(ctx)
}

// MergeProfilesPprof calls storegateway.v1.StoreGatewayService.MergeProfilesPprof.
func (c *storeGatewayServiceClient) MergeProfilesPprof(ctx context.Context) *connect.BidiStreamForClient[v11.MergeProfilesPprofRequest, v11.MergeProfilesPprofResponse] {
	return c.mergeProfilesPprof.CallBidiStream(ctx)
}

// MergeSpanProfile calls storegateway.v1.StoreGatewayService.MergeSpanProfile.
func (c *storeGatewayServiceClient) MergeSpanProfile(ctx context.Context) *connect.BidiStreamForClient[v11.MergeSpanProfileRequest, v11.MergeSpanProfileResponse] {
	return c.mergeSpanProfile.CallBidiStream(ctx)
}

// ProfileTypes calls storegateway.v1.StoreGatewayService.ProfileTypes.
func (c *storeGatewayServiceClient) ProfileTypes(ctx context.Context, req *connect.Request[v11.ProfileTypesRequest]) (*connect.Response[v11.ProfileTypesResponse], error) {
	return c.profileTypes.CallUnary(ctx, req)
}

// LabelValues calls storegateway.v1.StoreGatewayService.LabelValues.
func (c *storeGatewayServiceClient) LabelValues(ctx context.Context, req *connect.Request[v12.LabelValuesRequest]) (*connect.Response[v12.LabelValuesResponse], error) {
	return c.labelValues.CallUnary(ctx, req)
}

// LabelNames calls storegateway.v1.StoreGatewayService.LabelNames.
func (c *storeGatewayServiceClient) LabelNames(ctx context.Context, req *connect.Request[v12.LabelNamesRequest]) (*connect.Response[v12.LabelNamesResponse], error) {
	return c.labelNames.CallUnary(ctx, req)
}

// Series calls storegateway.v1.StoreGatewayService.Series.
func (c *storeGatewayServiceClient) Series(ctx context.Context, req *connect.Request[v11.SeriesRequest]) (*connect.Response[v11.SeriesResponse], error) {
	return c.series.CallUnary(ctx, req)
}

// BlockMetadata calls storegateway.v1.StoreGatewayService.BlockMetadata.
func (c *storeGatewayServiceClient) BlockMetadata(ctx context.Context, req *connect.Request[v11.BlockMetadataRequest]) (*connect.Response[v11.BlockMetadataResponse], error) {
	return c.blockMetadata.CallUnary(ctx, req)
}

// StoreGatewayServiceHandler is an implementation of the storegateway.v1.StoreGatewayService
// service.
type StoreGatewayServiceHandler interface {
	MergeProfilesStacktraces(context.Context, *connect.BidiStream[v11.MergeProfilesStacktracesRequest, v11.MergeProfilesStacktracesResponse]) error
	MergeProfilesLabels(context.Context, *connect.BidiStream[v11.MergeProfilesLabelsRequest, v11.MergeProfilesLabelsResponse]) error
	MergeProfilesPprof(context.Context, *connect.BidiStream[v11.MergeProfilesPprofRequest, v11.MergeProfilesPprofResponse]) error
	MergeSpanProfile(context.Context, *connect.BidiStream[v11.MergeSpanProfileRequest, v11.MergeSpanProfileResponse]) error
	// Deprecated: ProfileType call is deprecated in the store components
	// TODO: Remove this call in release v1.4
	ProfileTypes(context.Context, *connect.Request[v11.ProfileTypesRequest]) (*connect.Response[v11.ProfileTypesResponse], error)
	LabelValues(context.Context, *connect.Request[v12.LabelValuesRequest]) (*connect.Response[v12.LabelValuesResponse], error)
	LabelNames(context.Context, *connect.Request[v12.LabelNamesRequest]) (*connect.Response[v12.LabelNamesResponse], error)
	Series(context.Context, *connect.Request[v11.SeriesRequest]) (*connect.Response[v11.SeriesResponse], error)
	BlockMetadata(context.Context, *connect.Request[v11.BlockMetadataRequest]) (*connect.Response[v11.BlockMetadataResponse], error)
}

// NewStoreGatewayServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewStoreGatewayServiceHandler(svc StoreGatewayServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	storeGatewayServiceMergeProfilesStacktracesHandler := connect.NewBidiStreamHandler(
		StoreGatewayServiceMergeProfilesStacktracesProcedure,
		svc.MergeProfilesStacktraces,
		connect.WithSchema(storeGatewayServiceMergeProfilesStacktracesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	storeGatewayServiceMergeProfilesLabelsHandler := connect.NewBidiStreamHandler(
		StoreGatewayServiceMergeProfilesLabelsProcedure,
		svc.MergeProfilesLabels,
		connect.WithSchema(storeGatewayServiceMergeProfilesLabelsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	storeGatewayServiceMergeProfilesPprofHandler := connect.NewBidiStreamHandler(
		StoreGatewayServiceMergeProfilesPprofProcedure,
		svc.MergeProfilesPprof,
		connect.WithSchema(storeGatewayServiceMergeProfilesPprofMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	storeGatewayServiceMergeSpanProfileHandler := connect.NewBidiStreamHandler(
		StoreGatewayServiceMergeSpanProfileProcedure,
		svc.MergeSpanProfile,
		connect.WithSchema(storeGatewayServiceMergeSpanProfileMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	storeGatewayServiceProfileTypesHandler := connect.NewUnaryHandler(
		StoreGatewayServiceProfileTypesProcedure,
		svc.ProfileTypes,
		connect.WithSchema(storeGatewayServiceProfileTypesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	storeGatewayServiceLabelValuesHandler := connect.NewUnaryHandler(
		StoreGatewayServiceLabelValuesProcedure,
		svc.LabelValues,
		connect.WithSchema(storeGatewayServiceLabelValuesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	storeGatewayServiceLabelNamesHandler := connect.NewUnaryHandler(
		StoreGatewayServiceLabelNamesProcedure,
		svc.LabelNames,
		connect.WithSchema(storeGatewayServiceLabelNamesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	storeGatewayServiceSeriesHandler := connect.NewUnaryHandler(
		StoreGatewayServiceSeriesProcedure,
		svc.Series,
		connect.WithSchema(storeGatewayServiceSeriesMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	storeGatewayServiceBlockMetadataHandler := connect.NewUnaryHandler(
		StoreGatewayServiceBlockMetadataProcedure,
		svc.BlockMetadata,
		connect.WithSchema(storeGatewayServiceBlockMetadataMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/storegateway.v1.StoreGatewayService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case StoreGatewayServiceMergeProfilesStacktracesProcedure:
			storeGatewayServiceMergeProfilesStacktracesHandler.ServeHTTP(w, r)
		case StoreGatewayServiceMergeProfilesLabelsProcedure:
			storeGatewayServiceMergeProfilesLabelsHandler.ServeHTTP(w, r)
		case StoreGatewayServiceMergeProfilesPprofProcedure:
			storeGatewayServiceMergeProfilesPprofHandler.ServeHTTP(w, r)
		case StoreGatewayServiceMergeSpanProfileProcedure:
			storeGatewayServiceMergeSpanProfileHandler.ServeHTTP(w, r)
		case StoreGatewayServiceProfileTypesProcedure:
			storeGatewayServiceProfileTypesHandler.ServeHTTP(w, r)
		case StoreGatewayServiceLabelValuesProcedure:
			storeGatewayServiceLabelValuesHandler.ServeHTTP(w, r)
		case StoreGatewayServiceLabelNamesProcedure:
			storeGatewayServiceLabelNamesHandler.ServeHTTP(w, r)
		case StoreGatewayServiceSeriesProcedure:
			storeGatewayServiceSeriesHandler.ServeHTTP(w, r)
		case StoreGatewayServiceBlockMetadataProcedure:
			storeGatewayServiceBlockMetadataHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedStoreGatewayServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedStoreGatewayServiceHandler struct{}

func (UnimplementedStoreGatewayServiceHandler) MergeProfilesStacktraces(context.Context, *connect.BidiStream[v11.MergeProfilesStacktracesRequest, v11.MergeProfilesStacktracesResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("storegateway.v1.StoreGatewayService.MergeProfilesStacktraces is not implemented"))
}

func (UnimplementedStoreGatewayServiceHandler) MergeProfilesLabels(context.Context, *connect.BidiStream[v11.MergeProfilesLabelsRequest, v11.MergeProfilesLabelsResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("storegateway.v1.StoreGatewayService.MergeProfilesLabels is not implemented"))
}

func (UnimplementedStoreGatewayServiceHandler) MergeProfilesPprof(context.Context, *connect.BidiStream[v11.MergeProfilesPprofRequest, v11.MergeProfilesPprofResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("storegateway.v1.StoreGatewayService.MergeProfilesPprof is not implemented"))
}

func (UnimplementedStoreGatewayServiceHandler) MergeSpanProfile(context.Context, *connect.BidiStream[v11.MergeSpanProfileRequest, v11.MergeSpanProfileResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("storegateway.v1.StoreGatewayService.MergeSpanProfile is not implemented"))
}

func (UnimplementedStoreGatewayServiceHandler) ProfileTypes(context.Context, *connect.Request[v11.ProfileTypesRequest]) (*connect.Response[v11.ProfileTypesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("storegateway.v1.StoreGatewayService.ProfileTypes is not implemented"))
}

func (UnimplementedStoreGatewayServiceHandler) LabelValues(context.Context, *connect.Request[v12.LabelValuesRequest]) (*connect.Response[v12.LabelValuesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("storegateway.v1.StoreGatewayService.LabelValues is not implemented"))
}

func (UnimplementedStoreGatewayServiceHandler) LabelNames(context.Context, *connect.Request[v12.LabelNamesRequest]) (*connect.Response[v12.LabelNamesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("storegateway.v1.StoreGatewayService.LabelNames is not implemented"))
}

func (UnimplementedStoreGatewayServiceHandler) Series(context.Context, *connect.Request[v11.SeriesRequest]) (*connect.Response[v11.SeriesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("storegateway.v1.StoreGatewayService.Series is not implemented"))
}

func (UnimplementedStoreGatewayServiceHandler) BlockMetadata(context.Context, *connect.Request[v11.BlockMetadataRequest]) (*connect.Response[v11.BlockMetadataResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("storegateway.v1.StoreGatewayService.BlockMetadata is not implemented"))
}
