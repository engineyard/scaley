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
