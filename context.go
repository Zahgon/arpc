// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package arpc

import (
	"sync"
	"time"
)

var (
	contextPool = sync.Pool{
		New: func() interface{} {
			return &Context{}
		},
	}

	emptyContext = Context{}
)

// Context represents an arpc Call's context.
type Context struct {
	Client  *Client
	Message *Message

	index       int
	handlers    []HandlerFunc
	responseErr interface{}
}

func (ctx *Context) Release() { _ = "STUB: not implemented"; return }

func (ctx *Context) ResponseError() interface{} { _ = "STUB: not implemented"; return nil }

// Get returns value for key.
func (ctx *Context) Get(key interface{}) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Set sets key-value pair.
func (ctx *Context) Set(key interface{}, value interface{}) { _ = "STUB: not implemented"; return }

// Values returns values.
func (ctx *Context) Values() map[interface{}]interface{} { _ = "STUB: not implemented"; return nil }

// Body returns body.
func (ctx *Context) Body() []byte { _ = "STUB: not implemented"; return nil }

// Bind parses the body data and stores the result
// in the value pointed to by v.
func (ctx *Context) Bind(v interface{}) error { _ = "STUB: not implemented"; return nil }

// case *error:
// 	*vt = errors.New(util.BytesToStr(data))

// Write responses a Message to the Client.
func (ctx *Context) Write(v interface{}) error { _ = "STUB: not implemented"; return nil }

// WriteWithTimeout responses a Message to the Client with timeout.
func (ctx *Context) WriteWithTimeout(v interface{}, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// Error responses an error Message to the Client.
func (ctx *Context) Error(v interface{}) error { _ = "STUB: not implemented"; return nil }

// Next calls next middleware or method/router handler.
func (ctx *Context) Next() { _ = "STUB: not implemented"; return }

// Abort stops the one-by-one-calling of middlewares and method/router handler.
func (ctx *Context) Abort() { _ = "STUB: not implemented"; return }

// Deadline implements stdlib's Context.
func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	_ = "STUB: not implemented"

	// Done implements stdlib's Context.
	return *new(time.Time), false
}

func (ctx *Context) Done() <-chan struct{} {
	_ = "STUB: not implemented"

	// Err implements stdlib's Context.
	return nil
}

func (ctx *Context) Err() error {
	_ = "STUB: not implemented"

	// Value returns the value associated with this context for key, implements stdlib's Context.
	return nil
}

func (ctx *Context) Value(key interface{}) interface{} { _ = "STUB: not implemented"; return nil }

func (ctx *Context) write(v interface{}, isError bool, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *Context) writeDirectly(v interface{}, isError bool) error {
	_ = "STUB: not implemented"
	return nil
}

func newContext(cli *Client, msg *Message, handlers []HandlerFunc) *Context {
	_ = "STUB: not implemented"
	return nil
}
