// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
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

package barendpoint

import (
	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar/module"
	zanzibar "github.com/uber/zanzibar/runtime"
)

// Endpoint registers a request handler on a gateway
type Endpoint interface {
	Register(*zanzibar.Gateway) error
}

// NewEndpoint returns a collection of endpoints that can be registered on
// a gateway
func NewEndpoint(deps *module.Dependencies) Endpoint {
	return &EndpointHandlers{
		BarArgNotStructHandler:                     NewBarArgNotStructHandler(deps),
		BarArgWithHeadersHandler:                   NewBarArgWithHeadersHandler(deps),
		BarArgWithQueryParamsHandler:               NewBarArgWithQueryParamsHandler(deps),
		BarArgWithNestedQueryParamsHandler:         NewBarArgWithNestedQueryParamsHandler(deps),
		BarArgWithUntaggedNestedQueryParamsHandler: NewBarArgWithUntaggedNestedQueryParamsHandler(deps),
		BarArgWithQueryHeaderHandler:               NewBarArgWithQueryHeaderHandler(deps),
		BarArgWithParamsHandler:                    NewBarArgWithParamsHandler(deps),
		BarArgWithManyQueryParamsHandler:           NewBarArgWithManyQueryParamsHandler(deps),
		BarMissingArgHandler:                       NewBarMissingArgHandler(deps),
		BarNoRequestHandler:                        NewBarNoRequestHandler(deps),
		BarNormalHandler:                           NewBarNormalHandler(deps),
		BarTooManyArgsHandler:                      NewBarTooManyArgsHandler(deps),
		BarHelloWorldHandler:                       NewBarHelloWorldHandler(deps),
		BarListAndEnumHandler:                      NewBarListAndEnumHandler(deps),
		BarArgWithParamsAndDuplicateFieldsHandler:  NewBarArgWithParamsAndDuplicateFieldsHandler(deps),
	}
}

// EndpointHandlers is a collection of individual endpoint handlers
type EndpointHandlers struct {
	BarArgNotStructHandler                     *BarArgNotStructHandler
	BarArgWithHeadersHandler                   *BarArgWithHeadersHandler
	BarArgWithQueryParamsHandler               *BarArgWithQueryParamsHandler
	BarArgWithNestedQueryParamsHandler         *BarArgWithNestedQueryParamsHandler
	BarArgWithUntaggedNestedQueryParamsHandler *BarArgWithUntaggedNestedQueryParamsHandler
	BarArgWithQueryHeaderHandler               *BarArgWithQueryHeaderHandler
	BarArgWithParamsHandler                    *BarArgWithParamsHandler
	BarArgWithManyQueryParamsHandler           *BarArgWithManyQueryParamsHandler
	BarMissingArgHandler                       *BarMissingArgHandler
	BarNoRequestHandler                        *BarNoRequestHandler
	BarNormalHandler                           *BarNormalHandler
	BarTooManyArgsHandler                      *BarTooManyArgsHandler
	BarHelloWorldHandler                       *BarHelloWorldHandler
	BarListAndEnumHandler                      *BarListAndEnumHandler
	BarArgWithParamsAndDuplicateFieldsHandler  *BarArgWithParamsAndDuplicateFieldsHandler
}

// Register registers the endpoint handlers with the gateway
func (handlers *EndpointHandlers) Register(gateway *zanzibar.Gateway) error {
	err0 := handlers.BarArgNotStructHandler.Register(gateway)
	if err0 != nil {
		return err0
	}
	err1 := handlers.BarArgWithHeadersHandler.Register(gateway)
	if err1 != nil {
		return err1
	}
	err2 := handlers.BarArgWithQueryParamsHandler.Register(gateway)
	if err2 != nil {
		return err2
	}
	err3 := handlers.BarArgWithNestedQueryParamsHandler.Register(gateway)
	if err3 != nil {
		return err3
	}
	err4 := handlers.BarArgWithUntaggedNestedQueryParamsHandler.Register(gateway)
	if err4 != nil {
		return err4
	}
	err5 := handlers.BarArgWithQueryHeaderHandler.Register(gateway)
	if err5 != nil {
		return err5
	}
	err6 := handlers.BarArgWithParamsHandler.Register(gateway)
	if err6 != nil {
		return err6
	}
	err7 := handlers.BarArgWithManyQueryParamsHandler.Register(gateway)
	if err7 != nil {
		return err7
	}
	err8 := handlers.BarMissingArgHandler.Register(gateway)
	if err8 != nil {
		return err8
	}
	err9 := handlers.BarNoRequestHandler.Register(gateway)
	if err9 != nil {
		return err9
	}
	err10 := handlers.BarNormalHandler.Register(gateway)
	if err10 != nil {
		return err10
	}
	err11 := handlers.BarTooManyArgsHandler.Register(gateway)
	if err11 != nil {
		return err11
	}
	err12 := handlers.BarHelloWorldHandler.Register(gateway)
	if err12 != nil {
		return err12
	}
	err13 := handlers.BarListAndEnumHandler.Register(gateway)
	if err13 != nil {
		return err13
	}
	err14 := handlers.BarArgWithParamsAndDuplicateFieldsHandler.Register(gateway)
	if err14 != nil {
		return err14
	}
	return nil
}
