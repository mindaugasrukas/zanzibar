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

package echogatewayservicegeneratedmock

import (
	"github.com/golang/mock/gomock"
	module "github.com/uber/zanzibar/examples/example-gateway/build/services/echo-gateway/module"
	zanzibar "github.com/uber/zanzibar/runtime"

	echoendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/tchannel/echo"
	echoendpointmodule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/tchannel/echo/module"
)

// MockNodes contains mock  dependencies
type MockNodes struct {
}

// InitializeDependenciesMock fully initializes all dependencies in the dep tree
// for the echo-gateway service with leaf nodes being mocks
func InitializeDependenciesMock(
	g *zanzibar.Gateway,
	ctrl *gomock.Controller,
) (*module.DependenciesTree, *module.Dependencies, *MockNodes) {
	tree := &module.DependenciesTree{}

	mockNodes := &MockNodes{}
	initializedDefaultDependencies := &zanzibar.DefaultDependencies{
		ContextLogger: g.ContextLogger,
		Logger:        g.Logger,
		Scope:         g.AllHostScope,
		Config:        g.Config,
		Channel:       g.Channel,
		Tracer:        g.Tracer,
	}

	initializedEndpointDependencies := &module.EndpointDependenciesNodes{}
	tree.Endpoint = initializedEndpointDependencies
	initializedEndpointDependencies.Echo = echoendpointgenerated.NewEndpoint(&echoendpointmodule.Dependencies{
		Default: initializedDefaultDependencies,
	})

	dependencies := &module.Dependencies{
		Default: initializedDefaultDependencies,
		Endpoint: &module.EndpointDependencies{
			Echo: initializedEndpointDependencies.Echo,
		},
	}

	return tree, dependencies, mockNodes
}
