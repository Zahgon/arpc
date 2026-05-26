// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package arpc

import (
	"context"
	"io"
	"net"
	"sync"
	"time"

	"github.com/lesismal/arpc/codec"
)

// DefaultHandler is the default Handler used by arpc
var DefaultHandler Handler = NewHandler()

// HandlerFunc defines message handler.
type HandlerFunc func(*Context)

// StreamHandlerFunc defines stream handler.
type StreamHandlerFunc func(*Stream)

// AsyncHandlerFunc defines callback of Client.CallAsync.
type AsyncHandlerFunc func(*Context, error)

type asyncHandler struct {
	timer   *time.Timer
	handler AsyncHandlerFunc
}

var (
	emptyAsyncHandler = asyncHandler{}
	asyncHandlerPool  = sync.Pool{
		New: func() interface{} {
			return &asyncHandler{}
		},
	}
)

func getAsyncHandler(t *time.Timer, h AsyncHandlerFunc) *asyncHandler {
	_ = "STUB: not implemented"
	return nil
}

func putAsyncHandler(ah *asyncHandler) { _ = "STUB: not implemented"; return }

// routerHandler saves all middleware and method/router handler funcs
// for every method by register order,
// all the funcs will be called one by one for every message.
type routerHandler struct {
	async    bool
	handlers []HandlerFunc
}

// streamHandler saves all stream handler and middleware funcs.
// for every method by register order,
// all the funcs will be called one by one for every message.
type streamHandler struct {
	async   bool
	handler StreamHandlerFunc
}

// Handler defines net message handler interface.
type Handler interface {
	// Clone returns a copy of Handler.
	Clone() Handler

	// LogTag returns log tag value.
	LogTag() string
	// SetLogTag sets log tag.
	SetLogTag(tag string)

	// HandleConnected registers handler which will be called when client connected.
	HandleConnected(onConnected func(*Client))
	// OnConnected will be called when client is connected.
	OnConnected(c *Client)

	// HandleDisconnected registers handler which will be called when client is disconnected.
	HandleDisconnected(onDisConnected func(*Client))
	// OnDisconnected will be called when client is disconnected.
	OnDisconnected(c *Client)

	// MaxReconnectTimes returns client's max reconnect times.
	MaxReconnectTimes() int
	// SetMaxReconnectTimes sets client's max reconnect times for.
	SetMaxReconnectTimes(n int)

	// HandleOverstock registers handler which will be called when client send queue is overstock.
	HandleOverstock(onOverstock func(c *Client, m *Message))
	// OnOverstock will be called when client chSend is full.
	OnOverstock(c *Client, m *Message)

	// HandleMessageDone registers handler which will be called when message dropped.
	HandleMessageDone(onMessageDone func(c *Client, m *Message))
	// OnMessageDone will be called when message is dropped.
	OnMessageDone(c *Client, m *Message)

	// HandleMessageDropped registers handler which will be called when message dropped.
	HandleMessageDropped(onOverstock func(c *Client, m *Message))
	// OnOverstock will be called when message is dropped.
	OnMessageDropped(c *Client, m *Message)

	// HandleSessionMiss registers handler which will be called when async message seq not found.
	HandleSessionMiss(onSessionMiss func(c *Client, m *Message))
	// OnSessionMiss will be called when async message seq not found.
	OnSessionMiss(c *Client, m *Message)

	// HandleContextDone registers handler which will be called when message dropped.
	HandleContextDone(onContextDone func(ctx *Context))
	// OnContextDone will be called when message is dropped.
	OnContextDone(ctx *Context)

	// BeforeRecv registers handler which will be called before Recv.
	BeforeRecv(h func(net.Conn) error)
	// BeforeSend registers handler which will be called before Send.
	BeforeSend(h func(net.Conn) error)

	// BatchRecv returns BatchRecv flag.
	BatchRecv() bool
	// SetBatchRecv sets BatchRecv flag.
	SetBatchRecv(batch bool)
	// BatchSend returns BatchSend flag.
	BatchSend() bool
	// SetBatchSend sets BatchSend flag.
	SetBatchSend(batch bool)

	// AsyncWrite returns AsyncWrite flag.
	AsyncWrite() bool
	// SetAsyncWrite sets AsyncWrite flag.
	SetAsyncWrite(async bool)

	// AsyncResponse returns AsyncResponse flag.
	AsyncResponse() bool
	// SetAsyncResponse sets AsyncResponse flag.
	SetAsyncResponse(async bool)

	// WrapReader wraps net.Conn to Read data with io.Reader.
	WrapReader(conn net.Conn) io.Reader
	// SetReaderWrapper registers reader wrapper for net.Conn.
	SetReaderWrapper(wrapper func(conn net.Conn) io.Reader)

	// Recv reads a message from a client.
	Recv(c *Client) (*Message, error)
	// Send writes buffer data to a connection.
	Send(c net.Conn, buffer []byte) (int, error)
	// SendN writes multiple buffer data to a connection.
	SendN(conn net.Conn, buffers net.Buffers) (int, error)

	// RecvBufferSize returns client's recv buffer size.
	RecvBufferSize() int
	// SetRecvBufferSize sets client's recv buffer size.
	SetRecvBufferSize(size int)

	// SendBufferSize returns client's send buffer size.
	SendBufferSize() int
	// SetSendBufferSize sets client's send buffer size.
	SetSendBufferSize(size int)

	// ReadTimeout returns client's read timeout.
	ReadTimeout() time.Duration
	// SetReadTimeout sets client's read timeout.
	SetReadTimeout(timeout time.Duration)

	// WriteTimeout returns client's write timeout.
	WriteTimeout() time.Duration
	// SetWriteTimeout sets client's write timeout.
	SetWriteTimeout(timeout time.Duration)

	// SendQueueSize returns client's send queue channel capacity.
	SendQueueSize() int
	// SetSendQueueSize sets client's send queue channel capacity.
	SetSendQueueSize(size int)

	// StreamQueueSize returns stream queue channel capacity.
	StreamQueueSize() int
	// SetStreamQueueSize sets stream queue channel capacity.
	SetStreamQueueSize(size int)

	// MaxBodyLen returns max body length of a message.
	MaxBodyLen() int
	// SetMaxBodyLen sets max body length of a message.
	SetMaxBodyLen(l int)

	// Use registers method/router handler middleware.
	Use(h HandlerFunc)

	// UseCoder registers message coding middleware,
	// coder.Encode will be called before message send,
	// coder.Decode will be called after message recv.
	UseCoder(coder MessageCoder)

	// Coders returns coding middlewares.
	Coders() []MessageCoder

	// Handle registers method/router handler.
	//
	// If pass a Boolean value of "true", the handler will be called asynchronously in a new goroutine,
	// Else the handler will be called synchronously in the client's reading goroutine one by one.
	Handle(m string, h HandlerFunc, args ...interface{})

	// HandleNotFound registers "" method/router handler,
	// It will be called when mothod/router is not found.
	HandleNotFound(h HandlerFunc)

	// HandleStream registers method/router stream handler.
	HandleStream(m string, h StreamHandlerFunc, args ...interface{})

	// OnMessage finds method/router middlewares and handler, then call them one by one.
	OnMessage(c *Client, m *Message)

	// Malloc makes a buffer by size.
	Malloc(size int) []byte
	// HandleMalloc registers buffer maker.
	HandleMalloc(f func(size int) []byte)

	// Append append bytes to buffer.
	Append(b []byte, more ...byte) []byte
	// HandleAppend registers buffer appender.
	HandleAppend(f func(b []byte, more ...byte) []byte)

	// Free release a buffer.
	Free([]byte)
	// HandleFree registers buffer releaser.
	HandleFree(f func(buf []byte))

	// EnablePool registers handlers for pool operation for Context and Message and Message.Buffer
	EnablePool(enable bool)

	Context() (context.Context, context.CancelFunc)
	SetContext(ctx context.Context, cancel context.CancelFunc)
	Cancel()

	// NewMessage creates a Message.
	NewMessage(cmd byte, method string, v interface{}, isError bool, isAsync bool, seq uint64, codec codec.Codec, values map[interface{}]interface{}) *Message

	// NewMessageWithBuffer creates a message with the buffer and manage the message by the pool.
	// The buffer arg should be managed by a pool if EnablePool(true) .
	NewMessageWithBuffer(buffer []byte) *Message

	// SetAsyncExecutor sets executor.
	SetAsyncExecutor(executor func(f func()))
	// AsyncExecute executes a func
	AsyncExecute(f func())
}

// handler represents a default Handler implementation.
type handler struct {
	logtag            string
	batchRecv         bool
	batchSend         bool
	asyncWrite        bool
	asyncResponse     bool
	recvBufferSize    int
	sendBufferSize    int
	readTimeout       time.Duration
	writeTimeout      time.Duration
	sendQueueSize     int
	streamQueueSize   int
	maxBodyLen        int
	maxReconnectTimes int

	onConnected      func(*Client)
	onDisConnected   func(*Client)
	onOverstock      func(c *Client, m *Message)
	onMessageDone    func(c *Client, m *Message)
	onMessageDropped func(c *Client, m *Message)
	onSessionMiss    func(c *Client, m *Message)
	onContextDone    func(ctx *Context)

	beforeRecv func(net.Conn) error
	beforeSend func(net.Conn) error
	malloc     func(int) []byte
	append     func([]byte, ...byte) []byte
	free       func([]byte)

	wrapReader func(conn net.Conn) io.Reader

	routes  map[string]*routerHandler
	streams map[string]*streamHandler

	middles   []HandlerFunc
	msgCoders []MessageCoder

	ctx    context.Context
	cancel context.CancelFunc

	executor func(f func())
}

func (h *handler) Clone() Handler { _ = "STUB: not implemented"; return *new(Handler) }

func (h *handler) LogTag() string { _ = "STUB: not implemented"; return "" }

func (h *handler) SetLogTag(tag string) { _ = "STUB: not implemented"; return }

func (h *handler) HandleConnected(onConnected func(*Client)) { _ = "STUB: not implemented"; return }

func (h *handler) OnConnected(c *Client) { _ = "STUB: not implemented"; return }

func (h *handler) HandleDisconnected(onDisConnected func(*Client)) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) OnDisconnected(c *Client) { _ = "STUB: not implemented"; return }

func (h *handler) MaxReconnectTimes() int { _ = "STUB: not implemented"; return 0 }

func (h *handler) SetMaxReconnectTimes(n int) { _ = "STUB: not implemented"; return }

func (h *handler) HandleOverstock(onOverstock func(c *Client, m *Message)) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) OnOverstock(c *Client, m *Message) { _ = "STUB: not implemented"; return }

func (h *handler) HandleMessageDropped(onMessageDropped func(c *Client, m *Message)) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) OnMessageDropped(c *Client, m *Message) { _ = "STUB: not implemented"; return }

func (h *handler) HandleMessageDone(onMessageDone func(c *Client, m *Message)) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) OnMessageDone(c *Client, m *Message) { _ = "STUB: not implemented"; return }

func (h *handler) HandleSessionMiss(onSessionMiss func(c *Client, m *Message)) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) OnSessionMiss(c *Client, m *Message) { _ = "STUB: not implemented"; return }

func (h *handler) HandleContextDone(onContextDone func(ctx *Context)) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) OnContextDone(ctx *Context) { _ = "STUB: not implemented"; return }

func (h *handler) BeforeRecv(hb func(net.Conn) error) { _ = "STUB: not implemented"; return }

func (h *handler) BeforeSend(hs func(net.Conn) error) { _ = "STUB: not implemented"; return }

func (h *handler) BatchRecv() bool { _ = "STUB: not implemented"; return false }

func (h *handler) SetBatchRecv(batch bool) { _ = "STUB: not implemented"; return }

func (h *handler) BatchSend() bool { _ = "STUB: not implemented"; return false }

func (h *handler) SetBatchSend(batch bool) { _ = "STUB: not implemented"; return }

func (h *handler) AsyncWrite() bool { _ = "STUB: not implemented"; return false }

func (h *handler) SetAsyncWrite(async bool) { _ = "STUB: not implemented"; return }

func (h *handler) AsyncResponse() bool { _ = "STUB: not implemented"; return false }

func (h *handler) SetAsyncResponse(async bool) { _ = "STUB: not implemented"; return }

func (h *handler) WrapReader(conn net.Conn) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

func (h *handler) SetReaderWrapper(wrapper func(conn net.Conn) io.Reader) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) RecvBufferSize() int { _ = "STUB: not implemented"; return 0 }

func (h *handler) SetRecvBufferSize(size int) { _ = "STUB: not implemented"; return }

func (h *handler) SendBufferSize() int { _ = "STUB: not implemented"; return 0 }

func (h *handler) SetSendBufferSize(size int) { _ = "STUB: not implemented"; return }

func (h *handler) ReadTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (h *handler) SetReadTimeout(timeout time.Duration) { _ = "STUB: not implemented"; return }

func (h *handler) WriteTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (h *handler) SetWriteTimeout(timeout time.Duration) { _ = "STUB: not implemented"; return }

func (h *handler) SendQueueSize() int { _ = "STUB: not implemented"; return 0 }

func (h *handler) SetSendQueueSize(size int) { _ = "STUB: not implemented"; return }

func (h *handler) StreamQueueSize() int { _ = "STUB: not implemented"; return 0 }

func (h *handler) SetStreamQueueSize(size int) { _ = "STUB: not implemented"; return }

func (h *handler) MaxBodyLen() int { _ = "STUB: not implemented"; return 0 }

func (h *handler) SetMaxBodyLen(l int) { _ = "STUB: not implemented"; return }

func (h *handler) Use(cb HandlerFunc) { _ = "STUB: not implemented"; return }

func (h *handler) UseCoder(coder MessageCoder) { _ = "STUB: not implemented"; return }

func (h *handler) Coders() []MessageCoder { _ = "STUB: not implemented"; return nil }

func (h *handler) Handle(method string, cb HandlerFunc, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) HandleNotFound(cb HandlerFunc) { _ = "STUB: not implemented"; return }

func (h *handler) handle(method string, cb HandlerFunc, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) HandleStream(method string, cb StreamHandlerFunc, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) Recv(c *Client) (*Message, error) { _ = "STUB: not implemented"; return nil, nil }

func (h *handler) Send(conn net.Conn, buffer []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (h *handler) SendN(conn net.Conn, buffers net.Buffers) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (h *handler) OnMessage(c *Client, msg *Message) { _ = "STUB: not implemented"; return }

func (h *handler) Malloc(size int) []byte { _ = "STUB: not implemented"; return nil }

func (h *handler) HandleMalloc(f func(int) []byte) { _ = "STUB: not implemented"; return }

func (h *handler) Append(b []byte, more ...byte) []byte { _ = "STUB: not implemented"; return nil }

func (h *handler) HandleAppend(f func(b []byte, more ...byte) []byte) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) Free(b []byte) { _ = "STUB: not implemented"; return }

func (h *handler) HandleFree(f func([]byte)) { _ = "STUB: not implemented"; return }

func (h *handler) EnablePool(enable bool) { _ = "STUB: not implemented"; return }

func (h *handler) Context() (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

func (h *handler) SetContext(ctx context.Context, cancel context.CancelFunc) {
	_ = "STUB: not implemented"
	return
}

func (h *handler) Cancel() { _ = "STUB: not implemented"; return }

func (h *handler) NewMessage(cmd byte, method string, v interface{}, isError bool, isAsync bool, seq uint64, codec codec.Codec, values map[interface{}]interface{}) *Message {
	_ = "STUB: not implemented"
	return nil
}

func (h *handler) NewMessageWithBuffer(buffer []byte) *Message {
	_ = "STUB: not implemented"
	return nil
}

// SetAsyncExecutor sets executor for message.
func (h *handler) SetAsyncExecutor(executor func(f func())) { _ = "STUB: not implemented"; return }

// AsyncExecute executes a func.
func (h *handler) AsyncExecute(f func()) { _ = "STUB: not implemented"; return }

// NewHandler returns a default Handler implementation.
func NewHandler() Handler { _ = "STUB: not implemented"; return *new(Handler) }

// SetHandler sets default Handler.
func SetHandler(h Handler) {
	_ = "STUB: not implemented"

	// SetLogTag sets DefaultHandler's log tag.
	return
}

func SetLogTag(tag string) { _ = "STUB: not implemented"; return }

// HandleConnected registers default handler which will be called when client connected.
func HandleConnected(onConnected func(*Client)) { _ = "STUB: not implemented"; return }

// HandleDisconnected registers default handler which will be called when client disconnected.
func HandleDisconnected(onDisConnected func(*Client)) { _ = "STUB: not implemented"; return }

// HandleOverstock registers default handler which will be called when client send queue is overstock.
func HandleOverstock(onOverstock func(c *Client, m *Message)) { _ = "STUB: not implemented"; return }

// HandleMessageDropped registers default handler which will be called when message dropped.
func HandleMessageDropped(onOverstock func(c *Client, m *Message)) {
	_ = "STUB: not implemented"
	return
}

// HandleSessionMiss registers default handler which will be called when async message seq not found.
func HandleSessionMiss(onSessionMiss func(c *Client, m *Message)) {
	_ = "STUB: not implemented"
	return
}

// BeforeRecv registers default handler which will be called before Recv.
func BeforeRecv(h func(net.Conn) error) { _ = "STUB: not implemented"; return }

// BeforeSend registers default handler which will be called before Send.
func BeforeSend(h func(net.Conn) error) { _ = "STUB: not implemented"; return }

// BatchRecv returns default BatchRecv flag.
func BatchRecv() bool { _ = "STUB: not implemented"; return false }

// SetBatchRecv sets default BatchRecv flag.
func SetBatchRecv(batch bool) { _ = "STUB: not implemented"; return }

// BatchSend returns default BatchSend flag.
func BatchSend() bool { _ = "STUB: not implemented"; return false }

// SetBatchSend sets default BatchSend flag.
func SetBatchSend(batch bool) { _ = "STUB: not implemented"; return }

// AsyncResponse returns default AsyncResponse flag.
func AsyncResponse() bool { _ = "STUB: not implemented"; return false }

// SetAsyncResponse sets default AsyncResponse flag.
func SetAsyncResponse(async bool) { _ = "STUB: not implemented"; return }

// SetReaderWrapper registers default reader wrapper for net.Conn.
func SetReaderWrapper(wrapper func(conn net.Conn) io.Reader) { _ = "STUB: not implemented"; return }

// RecvBufferSize returns default client's read buffer size.
func RecvBufferSize() int { _ = "STUB: not implemented"; return 0 }

// SetRecvBufferSize sets default client's read buffer size.
func SetRecvBufferSize(size int) { _ = "STUB: not implemented"; return }

// SendBufferSize returns default client's read buffer size.
func SendBufferSize() int { _ = "STUB: not implemented"; return 0 }

// SetSendBufferSize sets default client's read buffer size.
func SetSendBufferSize(size int) { _ = "STUB: not implemented"; return }

// ReadTimeout returns client's read timeout.
func ReadTimeout() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// SetReadTimeout sets client's read timeout.
func SetReadTimeout(timeout time.Duration) { _ = "STUB: not implemented"; return }

// WriteTimeout returns client's write timeout.
func WriteTimeout() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// SetWriteTimeout sets client's write timeout.
func SetWriteTimeout(timeout time.Duration) { _ = "STUB: not implemented"; return }

// SendQueueSize returns default client's send queue channel capacity.
func SendQueueSize() int { _ = "STUB: not implemented"; return 0 }

// SetSendQueueSize sets default client's send queue channel capacity.
func SetSendQueueSize(size int) { _ = "STUB: not implemented"; return }

// StreamQueueSize returns default stream queue channel capacity.
func StreamQueueSize() int { _ = "STUB: not implemented"; return 0 }

// SetStreamQueueSize sets default stream queue channel capacity.
func SetStreamQueueSize(size int) { _ = "STUB: not implemented"; return }

func MaxBodyLen() int { _ = "STUB: not implemented"; return 0 }

func SetMaxBodyLen(l int) { _ = "STUB: not implemented"; return }

// Use registers default method/router handler middleware.
func Use(h HandlerFunc) { _ = "STUB: not implemented"; return }

// UseCoder registers default message coding middleware,
// coder.Encode will be called before message send,
// coder.Decode will be called after message recv.
func UseCoder(coder MessageCoder) { _ = "STUB: not implemented"; return }

// Handle registers default method/router handler.
//
// If pass a Boolean value of "true", the handler will be called asynchronously in a new goroutine,
// Else the handler will be called synchronously in the client's reading goroutine one by one.
func Handle(m string, h HandlerFunc, args ...interface{}) { _ = "STUB: not implemented"; return }

// HandleNotFound registers default "" method/router handler,
// It will be called when mothod/router is not found.
func HandleNotFound(h HandlerFunc) { _ = "STUB: not implemented"; return }

// HandleMalloc registers default buffer maker.
func HandleMalloc(f func(int) []byte) { _ = "STUB: not implemented"; return }

// HandleFree registers buffer releaser.
func HandleFree(f func([]byte)) { _ = "STUB: not implemented"; return }

// EnablePool registers handlers for pool operation for Context and Message and Message.Buffer
func EnablePool(enable bool) { _ = "STUB: not implemented"; return }

// SetAsyncExecutor sets executor.
// AsyncExecute executes a func
func SetAsyncExecutor(executor func(f func())) { _ = "STUB: not implemented"; return }

// AsyncExecute executes a func.
func AsyncExecute(f func()) { _ = "STUB: not implemented"; return }
