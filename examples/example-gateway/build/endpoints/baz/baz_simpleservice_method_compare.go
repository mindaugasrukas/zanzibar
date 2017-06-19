// Code generated by zanzibar
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package baz

import (
	"context"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"

	clientsBazBase "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/base"
	clientsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/baz"
	endpointsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/baz/baz"
)

// CompareHandler is the handler for "/baz/compare"
type CompareHandler struct {
	Clients *clients.Clients
}

// NewCompareEndpoint creates a handler
func NewCompareEndpoint(
	gateway *zanzibar.Gateway,
) *CompareHandler {
	return &CompareHandler{
		Clients: gateway.Clients.(*clients.Clients),
	}
}

// HandleRequest handles "/baz/compare".
func (handler *CompareHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {
	var requestBody endpointsBazBaz.SimpleService_Compare_Args
	if ok := req.ReadAndUnmarshalBody(&requestBody); !ok {
		return
	}

	workflow := CompareEndpoint{
		Clients: handler.Clients,
		Logger:  req.Logger,
		Request: req,
	}

	response, cliRespHeaders, err := workflow.Handle(ctx, req.Header, &requestBody)
	if err != nil {
		switch errValue := err.(type) {

		case *endpointsBazBaz.AuthErr:
			res.WriteJSON(
				403, cliRespHeaders, errValue,
			)
			return

		case *endpointsBazBaz.OtherAuthErr:
			res.WriteJSON(
				403, cliRespHeaders, errValue,
			)
			return

		default:
			req.Logger.Warn("Workflow for endpoint returned error",
				zap.String("error", errValue.Error()),
			)
			res.SendErrorString(500, "Unexpected server error")
			return
		}
	}

	res.WriteJSON(200, cliRespHeaders, response)
}

// CompareEndpoint calls thrift client Baz.Compare
type CompareEndpoint struct {
	Clients *clients.Clients
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w CompareEndpoint) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsBazBaz.SimpleService_Compare_Args,
) (*endpointsBazBaz.BazResponse, zanzibar.Header, error) {
	clientRequest := convertToCompareClientRequest(r)

	clientHeaders := map[string]string{}

	clientRespBody, _, err := w.Clients.Baz.Compare(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		case *clientsBazBaz.AuthErr:
			serverErr := convertCompareAuthErr(
				errValue,
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, nil, serverErr

		case *clientsBazBaz.OtherAuthErr:
			serverErr := convertCompareOtherAuthErr(
				errValue,
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, nil, serverErr

		default:
			w.Logger.Warn("Could not make client request",
				zap.String("error", errValue.Error()),
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.

	// TODO: Add support for TChannel Headers with a switch here
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertCompareClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertToCompareClientRequest(in *endpointsBazBaz.SimpleService_Compare_Args) *clientsBazBaz.SimpleService_Compare_Args {
	out := &clientsBazBaz.SimpleService_Compare_Args{}

	if in.Arg1 != nil {
		out.Arg1 = &clientsBazBaz.BazRequest{}
		out.Arg1.B1 = bool(in.Arg1.B1)
		out.Arg1.S2 = string(in.Arg1.S2)
		out.Arg1.I3 = int32(in.Arg1.I3)
	} else {
		out.Arg1 = nil
	}
	if in.Arg2 != nil {
		out.Arg2 = &clientsBazBaz.BazRequest{}
		out.Arg2.B1 = bool(in.Arg2.B1)
		out.Arg2.S2 = string(in.Arg2.S2)
		out.Arg2.I3 = int32(in.Arg2.I3)
	} else {
		out.Arg2 = nil
	}

	return out
}

func convertCompareAuthErr(
	clientError *clientsBazBaz.AuthErr,
) *endpointsBazBaz.AuthErr {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBazBaz.AuthErr{}
	return serverError
}
func convertCompareOtherAuthErr(
	clientError *clientsBazBaz.OtherAuthErr,
) *endpointsBazBaz.OtherAuthErr {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBazBaz.OtherAuthErr{}
	return serverError
}

func convertCompareClientResponse(in *clientsBazBase.BazResponse) *endpointsBazBaz.BazResponse {
	out := &endpointsBazBaz.BazResponse{}

	out.Message = string(in.Message)

	return out
}