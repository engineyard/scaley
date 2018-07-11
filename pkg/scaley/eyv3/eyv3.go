// Package eyv3 contains services and interfaces for working with verison 3
// of the Engine Yard Core API.
package eyv3

import (
	"net/url"

	"github.com/engineyard/eycore/core"
)

// Reader is an interface that describes low-level API read operations
type Reader interface {
	Get(string, url.Values) ([]byte, error)
}

// Writer is an interface that describes low-level API write operations
type Writer interface {
	Post(string, url.Values, core.Body) ([]byte, error)
	Put(string, url.Values, core.Body) ([]byte, error)
}

// Deleter is an interface that describes low-level API delete operations
type Deleter interface {
	Delete(string, url.Values) ([]byte, error)
}

// Client is an interface that describes a fully-featured low-level API driver
type Client interface {
	Reader
	Writer
	Deleter
}

// Copyright © 2018 Engine Yard, Inc.
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
