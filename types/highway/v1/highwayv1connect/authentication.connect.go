// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: highway/v1/authentication.proto

package highwayv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/sonrhq/core/types/highway/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// AuthenticationName is the fully-qualified name of the Authentication service.
	AuthenticationName = "sonrhq.highway.v1.Authentication"
)

// AuthenticationClient is a client for the sonrhq.highway.v1.Authentication service.
type AuthenticationClient interface {
	Keygen(context.Context, *connect_go.Request[v1.KeygenRequest]) (*connect_go.Response[v1.KeygenResponse], error)
	Login(context.Context, *connect_go.Request[v1.LoginRequest]) (*connect_go.Response[v1.LoginResponse], error)
	QueryDocument(context.Context, *connect_go.Request[v1.QueryDocumentRequest]) (*connect_go.Response[v1.QueryDocumentResponse], error)
	QueryService(context.Context, *connect_go.Request[v1.QueryServiceRequest]) (*connect_go.Response[v1.QueryServiceResponse], error)
}

// NewAuthenticationClient constructs a client for the sonrhq.highway.v1.Authentication service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAuthenticationClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AuthenticationClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &authenticationClient{
		keygen: connect_go.NewClient[v1.KeygenRequest, v1.KeygenResponse](
			httpClient,
			baseURL+"/sonrhq.highway.v1.Authentication/Keygen",
			opts...,
		),
		login: connect_go.NewClient[v1.LoginRequest, v1.LoginResponse](
			httpClient,
			baseURL+"/sonrhq.highway.v1.Authentication/Login",
			opts...,
		),
		queryDocument: connect_go.NewClient[v1.QueryDocumentRequest, v1.QueryDocumentResponse](
			httpClient,
			baseURL+"/sonrhq.highway.v1.Authentication/QueryDocument",
			opts...,
		),
		queryService: connect_go.NewClient[v1.QueryServiceRequest, v1.QueryServiceResponse](
			httpClient,
			baseURL+"/sonrhq.highway.v1.Authentication/QueryService",
			opts...,
		),
	}
}

// authenticationClient implements AuthenticationClient.
type authenticationClient struct {
	keygen        *connect_go.Client[v1.KeygenRequest, v1.KeygenResponse]
	login         *connect_go.Client[v1.LoginRequest, v1.LoginResponse]
	queryDocument *connect_go.Client[v1.QueryDocumentRequest, v1.QueryDocumentResponse]
	queryService  *connect_go.Client[v1.QueryServiceRequest, v1.QueryServiceResponse]
}

// Keygen calls sonrhq.highway.v1.Authentication.Keygen.
func (c *authenticationClient) Keygen(ctx context.Context, req *connect_go.Request[v1.KeygenRequest]) (*connect_go.Response[v1.KeygenResponse], error) {
	return c.keygen.CallUnary(ctx, req)
}

// Login calls sonrhq.highway.v1.Authentication.Login.
func (c *authenticationClient) Login(ctx context.Context, req *connect_go.Request[v1.LoginRequest]) (*connect_go.Response[v1.LoginResponse], error) {
	return c.login.CallUnary(ctx, req)
}

// QueryDocument calls sonrhq.highway.v1.Authentication.QueryDocument.
func (c *authenticationClient) QueryDocument(ctx context.Context, req *connect_go.Request[v1.QueryDocumentRequest]) (*connect_go.Response[v1.QueryDocumentResponse], error) {
	return c.queryDocument.CallUnary(ctx, req)
}

// QueryService calls sonrhq.highway.v1.Authentication.QueryService.
func (c *authenticationClient) QueryService(ctx context.Context, req *connect_go.Request[v1.QueryServiceRequest]) (*connect_go.Response[v1.QueryServiceResponse], error) {
	return c.queryService.CallUnary(ctx, req)
}

// AuthenticationHandler is an implementation of the sonrhq.highway.v1.Authentication service.
type AuthenticationHandler interface {
	Keygen(context.Context, *connect_go.Request[v1.KeygenRequest]) (*connect_go.Response[v1.KeygenResponse], error)
	Login(context.Context, *connect_go.Request[v1.LoginRequest]) (*connect_go.Response[v1.LoginResponse], error)
	QueryDocument(context.Context, *connect_go.Request[v1.QueryDocumentRequest]) (*connect_go.Response[v1.QueryDocumentResponse], error)
	QueryService(context.Context, *connect_go.Request[v1.QueryServiceRequest]) (*connect_go.Response[v1.QueryServiceResponse], error)
}

// NewAuthenticationHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAuthenticationHandler(svc AuthenticationHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/sonrhq.highway.v1.Authentication/Keygen", connect_go.NewUnaryHandler(
		"/sonrhq.highway.v1.Authentication/Keygen",
		svc.Keygen,
		opts...,
	))
	mux.Handle("/sonrhq.highway.v1.Authentication/Login", connect_go.NewUnaryHandler(
		"/sonrhq.highway.v1.Authentication/Login",
		svc.Login,
		opts...,
	))
	mux.Handle("/sonrhq.highway.v1.Authentication/QueryDocument", connect_go.NewUnaryHandler(
		"/sonrhq.highway.v1.Authentication/QueryDocument",
		svc.QueryDocument,
		opts...,
	))
	mux.Handle("/sonrhq.highway.v1.Authentication/QueryService", connect_go.NewUnaryHandler(
		"/sonrhq.highway.v1.Authentication/QueryService",
		svc.QueryService,
		opts...,
	))
	return "/sonrhq.highway.v1.Authentication/", mux
}

// UnimplementedAuthenticationHandler returns CodeUnimplemented from all methods.
type UnimplementedAuthenticationHandler struct{}

func (UnimplementedAuthenticationHandler) Keygen(context.Context, *connect_go.Request[v1.KeygenRequest]) (*connect_go.Response[v1.KeygenResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sonrhq.highway.v1.Authentication.Keygen is not implemented"))
}

func (UnimplementedAuthenticationHandler) Login(context.Context, *connect_go.Request[v1.LoginRequest]) (*connect_go.Response[v1.LoginResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sonrhq.highway.v1.Authentication.Login is not implemented"))
}

func (UnimplementedAuthenticationHandler) QueryDocument(context.Context, *connect_go.Request[v1.QueryDocumentRequest]) (*connect_go.Response[v1.QueryDocumentResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sonrhq.highway.v1.Authentication.QueryDocument is not implemented"))
}

func (UnimplementedAuthenticationHandler) QueryService(context.Context, *connect_go.Request[v1.QueryServiceRequest]) (*connect_go.Response[v1.QueryServiceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sonrhq.highway.v1.Authentication.QueryService is not implemented"))
}