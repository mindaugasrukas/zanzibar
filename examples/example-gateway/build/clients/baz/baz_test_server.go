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

// Package bazClient is generated code used to make or handle TChannel calls using Thrift.
package bazClient

import (
	"context"
	"errors"

	"github.com/uber/zanzibar/runtime"
	"go.uber.org/thriftrw/wire"

	clientsBazBase "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/base"
	clientsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/baz"
)

// SecondServiceEchoFunc is the handler function for "Echo" method of thrift service "SecondService".
type SecondServiceEchoFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SecondService_Echo_Args,
) (string, map[string]string, error)

// NewSecondServiceEchoHandler wraps a handler function so it can be registered with a thrift server.
func NewSecondServiceEchoHandler(f SecondServiceEchoFunc) zanzibar.TChannelHandler {
	return &SecondServiceEchoHandler{f}
}

// SecondServiceEchoHandler handles the "Echo" method call of thrift service "SecondService".
type SecondServiceEchoHandler struct {
	echo SecondServiceEchoFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SecondServiceEchoHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SecondService_Echo_Args
	var res clientsBazBaz.SecondService_Echo_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	r, respHeaders, err := h.echo(ctx, reqHeaders, &req)

	if err != nil {
		return false, nil, nil, err
	}
	res.Success = r

	return err == nil, &res, respHeaders, nil
}

// SimpleServiceCallFunc is the handler function for "Call" method of thrift service "SimpleService".
type SimpleServiceCallFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SimpleService_Call_Args,
) (map[string]string, error)

// NewSimpleServiceCallHandler wraps a handler function so it can be registered with a thrift server.
func NewSimpleServiceCallHandler(f SimpleServiceCallFunc) zanzibar.TChannelHandler {
	return &SimpleServiceCallHandler{f}
}

// SimpleServiceCallHandler handles the "Call" method call of thrift service "SimpleService".
type SimpleServiceCallHandler struct {
	call SimpleServiceCallFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SimpleServiceCallHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SimpleService_Call_Args
	var res clientsBazBaz.SimpleService_Call_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	respHeaders, err := h.call(ctx, reqHeaders, &req)

	if err != nil {
		switch v := err.(type) {
		case *clientsBazBaz.AuthErr:
			if v == nil {
				return false, nil, nil, errors.New(
					"Handler for Call returned non-nil error type *AuthErr but nil value",
				)
			}
			res.AuthErr = v
		default:
			return false, nil, nil, err
		}
	}

	return err == nil, &res, respHeaders, nil
}

// SimpleServiceCompareFunc is the handler function for "Compare" method of thrift service "SimpleService".
type SimpleServiceCompareFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsBazBaz.SimpleService_Compare_Args,
) (*clientsBazBase.BazResponse, map[string]string, error)

// NewSimpleServiceCompareHandler wraps a handler function so it can be registered with a thrift server.
func NewSimpleServiceCompareHandler(f SimpleServiceCompareFunc) zanzibar.TChannelHandler {
	return &SimpleServiceCompareHandler{f}
}

// SimpleServiceCompareHandler handles the "Compare" method call of thrift service "SimpleService".
type SimpleServiceCompareHandler struct {
	compare SimpleServiceCompareFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SimpleServiceCompareHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SimpleService_Compare_Args
	var res clientsBazBaz.SimpleService_Compare_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	r, respHeaders, err := h.compare(ctx, reqHeaders, &req)

	if err != nil {
		switch v := err.(type) {
		case *clientsBazBaz.AuthErr:
			if v == nil {
				return false, nil, nil, errors.New(
					"Handler for Compare returned non-nil error type *AuthErr but nil value",
				)
			}
			res.AuthErr = v
		default:
			return false, nil, nil, err
		}
	} else {
		res.Success = r
	}

	return err == nil, &res, respHeaders, nil
}

// SimpleServicePingFunc is the handler function for "Ping" method of thrift service "SimpleService".
type SimpleServicePingFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
) (*clientsBazBase.BazResponse, map[string]string, error)

// NewSimpleServicePingHandler wraps a handler function so it can be registered with a thrift server.
func NewSimpleServicePingHandler(f SimpleServicePingFunc) zanzibar.TChannelHandler {
	return &SimpleServicePingHandler{f}
}

// SimpleServicePingHandler handles the "Ping" method call of thrift service "SimpleService".
type SimpleServicePingHandler struct {
	ping SimpleServicePingFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SimpleServicePingHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SimpleService_Ping_Args
	var res clientsBazBaz.SimpleService_Ping_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	r, respHeaders, err := h.ping(ctx, reqHeaders)

	if err != nil {
		return false, nil, nil, err
	}
	res.Success = r

	return err == nil, &res, respHeaders, nil
}

// SimpleServiceSillyNoopFunc is the handler function for "SillyNoop" method of thrift service "SimpleService".
type SimpleServiceSillyNoopFunc func(
	ctx context.Context,
	reqHeaders map[string]string,
) (map[string]string, error)

// NewSimpleServiceSillyNoopHandler wraps a handler function so it can be registered with a thrift server.
func NewSimpleServiceSillyNoopHandler(f SimpleServiceSillyNoopFunc) zanzibar.TChannelHandler {
	return &SimpleServiceSillyNoopHandler{f}
}

// SimpleServiceSillyNoopHandler handles the "SillyNoop" method call of thrift service "SimpleService".
type SimpleServiceSillyNoopHandler struct {
	sillynoop SimpleServiceSillyNoopFunc
}

// Handle parses request from wire value and calls corresponding handler function.
func (h *SimpleServiceSillyNoopHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	var req clientsBazBaz.SimpleService_SillyNoop_Args
	var res clientsBazBaz.SimpleService_SillyNoop_Result

	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	respHeaders, err := h.sillynoop(ctx, reqHeaders)

	if err != nil {
		switch v := err.(type) {
		case *clientsBazBaz.AuthErr:
			if v == nil {
				return false, nil, nil, errors.New(
					"Handler for SillyNoop returned non-nil error type *AuthErr but nil value",
				)
			}
			res.AuthErr = v
		case *clientsBazBase.ServerErr:
			if v == nil {
				return false, nil, nil, errors.New(
					"Handler for SillyNoop returned non-nil error type *ServerErr but nil value",
				)
			}
			res.ServerErr = v
		default:
			return false, nil, nil, err
		}
	}

	return err == nil, &res, respHeaders, nil
}
