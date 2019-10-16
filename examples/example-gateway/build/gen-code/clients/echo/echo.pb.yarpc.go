// Code generated by protoc-gen-yarpc-go. DO NOT EDIT.
// source: echo.proto

package echo

import (
	"context"
	"io/ioutil"
	"reflect"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"go.uber.org/fx"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/encoding/protobuf"
	"go.uber.org/yarpc/encoding/protobuf/reflection"
)

var _ = ioutil.NopCloser

// EchoYARPCClient is the YARPC client-side interface for the Echo service.
type EchoYARPCClient interface {
	Echo(context.Context, *Request, ...yarpc.CallOption) (*Response, error)
}

func newEchoYARPCClient(clientConfig transport.ClientConfig, anyResolver jsonpb.AnyResolver, options ...protobuf.ClientOption) EchoYARPCClient {
	return &_EchoYARPCCaller{protobuf.NewStreamClient(
		protobuf.ClientParams{
			ServiceName:  "echo.Echo",
			ClientConfig: clientConfig,
			AnyResolver:  anyResolver,
			Options:      options,
		},
	)}
}

// NewEchoYARPCClient builds a new YARPC client for the Echo service.
func NewEchoYARPCClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) EchoYARPCClient {
	return newEchoYARPCClient(clientConfig, nil, options...)
}

// EchoYARPCServer is the YARPC server-side interface for the Echo service.
type EchoYARPCServer interface {
	Echo(context.Context, *Request) (*Response, error)
}

type buildEchoYARPCProceduresParams struct {
	Server      EchoYARPCServer
	AnyResolver jsonpb.AnyResolver
}

func buildEchoYARPCProcedures(params buildEchoYARPCProceduresParams) []transport.Procedure {
	handler := &_EchoYARPCHandler{params.Server}
	return protobuf.BuildProcedures(
		protobuf.BuildProceduresParams{
			ServiceName: "echo.Echo",
			UnaryHandlerParams: []protobuf.BuildProceduresUnaryHandlerParams{
				{
					MethodName: "Echo",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:      handler.Echo,
							NewRequest:  newEchoServiceEchoYARPCRequest,
							AnyResolver: params.AnyResolver,
						},
					),
				},
			},
			OnewayHandlerParams: []protobuf.BuildProceduresOnewayHandlerParams{},
			StreamHandlerParams: []protobuf.BuildProceduresStreamHandlerParams{},
		},
	)
}

// BuildEchoYARPCProcedures prepares an implementation of the Echo service for YARPC registration.
func BuildEchoYARPCProcedures(server EchoYARPCServer) []transport.Procedure {
	return buildEchoYARPCProcedures(buildEchoYARPCProceduresParams{Server: server})
}

// FxEchoYARPCClientParams defines the input
// for NewFxEchoYARPCClient. It provides the
// paramaters to get a EchoYARPCClient in an
// Fx application.
type FxEchoYARPCClientParams struct {
	fx.In

	Provider    yarpc.ClientConfig
	AnyResolver jsonpb.AnyResolver `name:"yarpcfx" optional:"true"`
}

// FxEchoYARPCClientResult defines the output
// of NewFxEchoYARPCClient. It provides a
// EchoYARPCClient to an Fx application.
type FxEchoYARPCClientResult struct {
	fx.Out

	Client EchoYARPCClient

	// We are using an fx.Out struct here instead of just returning a client
	// so that we can add more values or add named versions of the client in
	// the future without breaking any existing code.
}

// NewFxEchoYARPCClient provides a EchoYARPCClient
// to an Fx application using the given name for routing.
//
//  fx.Provide(
//    echo.NewFxEchoYARPCClient("service-name"),
//    ...
//  )
func NewFxEchoYARPCClient(name string, options ...protobuf.ClientOption) interface{} {
	return func(params FxEchoYARPCClientParams) FxEchoYARPCClientResult {
		return FxEchoYARPCClientResult{
			Client: newEchoYARPCClient(params.Provider.ClientConfig(name), params.AnyResolver, options...),
		}
	}
}

// FxEchoYARPCProceduresParams defines the input
// for NewFxEchoYARPCProcedures. It provides the
// paramaters to get EchoYARPCServer procedures in an
// Fx application.
type FxEchoYARPCProceduresParams struct {
	fx.In

	Server      EchoYARPCServer
	AnyResolver jsonpb.AnyResolver `name:"yarpcfx" optional:"true"`
}

// FxEchoYARPCProceduresResult defines the output
// of NewFxEchoYARPCProcedures. It provides
// EchoYARPCServer procedures to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group.
// Dig 1.2 or newer must be used for this feature to work.
type FxEchoYARPCProceduresResult struct {
	fx.Out

	Procedures     []transport.Procedure `group:"yarpcfx"`
	ReflectionMeta reflection.ServerMeta `group:"yarpcfx"`
}

// NewFxEchoYARPCProcedures provides EchoYARPCServer procedures to an Fx application.
// It expects a EchoYARPCServer to be present in the container.
//
//  fx.Provide(
//    echo.NewFxEchoYARPCProcedures(),
//    ...
//  )
func NewFxEchoYARPCProcedures() interface{} {
	return func(params FxEchoYARPCProceduresParams) FxEchoYARPCProceduresResult {
		return FxEchoYARPCProceduresResult{
			Procedures: buildEchoYARPCProcedures(buildEchoYARPCProceduresParams{
				Server:      params.Server,
				AnyResolver: params.AnyResolver,
			}),
			ReflectionMeta: reflection.ServerMeta{
				ServiceName:     "echo.Echo",
				FileDescriptors: yarpcFileDescriptorClosure08134aea513e0001,
			},
		}
	}
}

type _EchoYARPCCaller struct {
	streamClient protobuf.StreamClient
}

func (c *_EchoYARPCCaller) Echo(ctx context.Context, request *Request, options ...yarpc.CallOption) (*Response, error) {
	responseMessage, err := c.streamClient.Call(ctx, "Echo", request, newEchoServiceEchoYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*Response)
	if !ok {
		return nil, protobuf.CastError(emptyEchoServiceEchoYARPCResponse, responseMessage)
	}
	return response, err
}

type _EchoYARPCHandler struct {
	server EchoYARPCServer
}

func (h *_EchoYARPCHandler) Echo(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *Request
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*Request)
		if !ok {
			return nil, protobuf.CastError(emptyEchoServiceEchoYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.Echo(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func newEchoServiceEchoYARPCRequest() proto.Message {
	return &Request{}
}

func newEchoServiceEchoYARPCResponse() proto.Message {
	return &Response{}
}

var (
	emptyEchoServiceEchoYARPCRequest  = &Request{}
	emptyEchoServiceEchoYARPCResponse = &Response{}
)

var yarpcFileDescriptorClosure08134aea513e0001 = [][]byte{
	// echo.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4d, 0xce, 0xc8,
		0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x94, 0xb9, 0xd8, 0x83, 0x52,
		0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53,
		0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x25, 0x15, 0x2e, 0x8e, 0xa0, 0xd4, 0xe2,
		0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xdc, 0xaa, 0x8c, 0x74, 0xb9, 0x58, 0x5c, 0x93, 0x33, 0xf2, 0x85,
		0x54, 0xa1, 0x34, 0xaf, 0x1e, 0xd8, 0x36, 0xa8, 0xf1, 0x52, 0x7c, 0x30, 0x2e, 0xc4, 0xa0, 0x24,
		0x36, 0xb0, 0x33, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x1c, 0xf0, 0xfd, 0x94, 0x00,
		0x00, 0x00,
	},
}

func init() {
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) EchoYARPCClient {
			return NewEchoYARPCClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
}
