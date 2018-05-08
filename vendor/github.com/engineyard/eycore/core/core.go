// Copyright © 2017 Engine Yard, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package core provides common interfaces and functions used throughout the
// eycore packages.
package core

import (
	"net/url"
)

// The Body interface describes an object that can be easily converted to
// a POST/PUT body for HTTP requests.
type Body interface {
	Body() []byte
}

// The Client interface describes any object that konws how to handle
// interactions with the Engine Yard Core API.
type Client interface {
	Get(string, url.Values) ([]byte, error)
	Post(string, url.Values, Body) ([]byte, error)
	Put(string, url.Values, Body) ([]byte, error)
	Delete(string, url.Values) ([]byte, error)
}

// Error is a specific error type for errors returned from the Engine Yard
// API.
type Error struct {
	ErrorString string
}

// Error is the full error string from the API for a given operation.
func (err *Error) Error() string {
	return err.ErrorString
}

// NewError instantiates a Error.
func NewError(errorString string) *Error {
	return &Error{
		ErrorString: errorString,
	}
}
