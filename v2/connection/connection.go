//
// DISCLAIMER
//
// Copyright 2020-2021 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//
// Author Adam Janikowski
// Author Tomasz Mielech
//

package connection

import (
	"context"
	"io"
)

type Wrapper func(c Connection) Connection

type Factory func() (Connection, error)

type Connection interface {
	NewRequest(method string, urls ...string) (Request, error)
	NewRequestWithEndpoint(endpoint string, method string, urls ...string) (Request, error)

	Do(ctx context.Context, request Request, output interface{}) (Response, error)
	Stream(ctx context.Context, request Request) (Response, io.ReadCloser, error)

	GetEndpoint() Endpoint
	SetEndpoint(e Endpoint) error

	GetAuthentication() Authentication
	SetAuthentication(a Authentication) error

	Decoder(contentType string) Decoder
}

type Request interface {
	Method() string
	URL() string

	Endpoint() string

	SetBody(i interface{}) error
	AddHeader(key, value string)
	AddQuery(key, value string)

	GetHeader(key string) (string, bool)
	GetQuery(key string) (string, bool)

	SetFragment(s string)
}

type Response interface {
	Code() int
	Response() interface{}
	Endpoint() string
	Content() string
}
