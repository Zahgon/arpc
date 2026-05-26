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
	"github.com/lesismal/arpc/util"
)

const (
	// TimeZero represents zero time.
	TimeZero time.Duration = 0
	// TimeForever represents forever time.
	TimeForever time.Duration = 1<<63 - 1
)

// DialerFunc defines the dialer used by arpc Client to connect to the server.
type DialerFunc func() (net.Conn, error)

// rpcSession represents an active calling session.
type rpcSession struct {
	seq  uint64
	done chan *Message
}

// newSession creates rpcSession
func newSession(seq uint64) *rpcSession { _ = "STUB: not implemented"; return nil }

// Client represents an arpc Client.
// There may be multiple outstanding Calls or Notifys associated
// with a single Client, and a Client may be used by
// multiple goroutines simultaneously.
type Client struct {
	// 64-aligned on 32-bit
	seq uint64

	Conn    net.Conn
	Codec   codec.Codec
	Handler Handler
	Reader  io.Reader
	Dialer  DialerFunc
	Head    Header

	running      bool
	reconnecting bool

	mux             sync.Mutex
	sessionMap      map[uint64]*rpcSession
	asyncHandlerMap map[uint64]*asyncHandler
	streamLocalMap  map[uint64]*Stream
	streamRemoteMap map[uint64]*Stream

	chSend  chan *Message
	chClose chan util.Empty

	onStop func(*Client)

	values map[interface{}]interface{}
	// UserData interface{}
}

// SetState sets running state, should be used only for non-blocking conn.
func (c *Client) SetState(running bool) { _ = "STUB: not implemented"; return }

// Get returns value for key.
func (c *Client) Get(key interface{}) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Set sets key-value pair.
func (c *Client) Set(key interface{}, value interface{}) { _ = "STUB: not implemented"; return }

// Delete deletes key-value pair
func (c *Client) Delete(key interface{}) { _ = "STUB: not implemented"; return }

// Ping .
func (c *Client) Ping() { _ = "STUB: not implemented"; return }

// Pong .
func (c *Client) Pong() { _ = "STUB: not implemented"; return }

// Ping .
func (c *Client) Keepalive(interval time.Duration) { _ = "STUB: not implemented"; return }

// NewMessage creates a Message by client's seq, handler and codec.
func (c *Client) NewMessage(cmd byte, method string, v interface{}, args ...interface{}) *Message {
	_ = "STUB: not implemented"
	return nil
}

// Call makes an rpc call with a timeout.
// Call will block waiting for the server's response until timeout.
func (c *Client) Call(method string, req interface{}, rsp interface{}, timeout time.Duration, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// if timeout < 0 {
// 	timeout = TimeForever
// }

// c.Handler.OnOverstock(c, msg)

// c.Handler.OnOverstock(c, msg)

// CallWith uses context to make rpc call.
// CallWith blocks to wait for a response from the server until it times out.
func (c *Client) CallWith(ctx context.Context, method string, req interface{}, rsp interface{}, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// CallContext uses context to make rpc call.
// CallContext blocks to wait for a response from the server until it times out.
func (c *Client) CallContext(ctx context.Context, method string, req interface{}, rsp interface{}, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// c.Handler.OnOverstock(c, msg)

// c.Handler.OnOverstock(c, msg)

// CallAsync makes an asynchronous rpc call with timeout.
// CallAsync will not block waiting for the server's response,
// But the handler will be called if the response arrives before the timeout.
func (c *Client) CallAsync(method string, req interface{}, handler AsyncHandlerFunc, timeout time.Duration, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Notify makes a notify with timeout.
// A notify does not need a response from the server.
func (c *Client) Notify(method string, data interface{}, timeout time.Duration, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// NotifyWith use context to make rpc notify.
// A notify does not need a response from the server.
func (c *Client) NotifyWith(ctx context.Context, method string, data interface{}, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// NotifyContext use context to make rpc notify.
// A notify does not need a response from the server.
func (c *Client) NotifyContext(ctx context.Context, method string, data interface{}, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// c.Handler.OnOverstock(c, msg)

// c.Handler.OnOverstock(c, msg)

// PushMsg pushes a msg to Client's send queue with timeout.
func (c *Client) PushMsg(msg *Message, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// c.Handler.OnOverstock(c, msg)

// Restart stops and restarts a Client.
func (c *Client) Restart() error { _ = "STUB: not implemented"; return nil }

// Stop stops a Client.
func (c *Client) Stop() { _ = "STUB: not implemented"; return }

func (c *Client) closeAndClean() { _ = "STUB: not implemented"; return }

// CheckState checks Client's state.
func (c *Client) CheckState() error { _ = "STUB: not implemented"; return nil }

func (c *Client) checkCallArgs(method string, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) checkCallAsyncArgs(method string, handler AsyncHandlerFunc, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) checkNotifyArgs(method string, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) checkStateAndMethod(method string) error { _ = "STUB: not implemented"; return nil }

func (c *Client) pushMessage(msg *Message, timer *time.Timer) error {
	_ = "STUB: not implemented"
	return nil
}

// c.Handler.OnOverstock(c, msg)

// c.Handler.OnOverstock(c, msg)

// c.Handler.OnOverstock(c, msg)

func (c *Client) newRequestMessage(cmd byte, method string, v interface{}, isError bool, isAsync bool, args ...interface{}) *Message {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) parseResponse(msg *Message, rsp interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// case *error:
// 	*vt = msg.Error()

func (c *Client) addSession(seq uint64, session *rpcSession) { _ = "STUB: not implemented"; return }

func (c *Client) deleteSession(seq uint64) *rpcSession { _ = "STUB: not implemented"; return nil }

func (c *Client) getSession(seq uint64) (*rpcSession, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *Client) clearSession() { _ = "STUB: not implemented"; return }

func (c *Client) dropMessage(msg *Message) { _ = "STUB: not implemented"; return }

func (c *Client) addAsyncHandler(seq uint64, ah *asyncHandler) { _ = "STUB: not implemented"; return }

func (c *Client) deleteAsyncHandler(seq uint64) { _ = "STUB: not implemented"; return }

func (c *Client) getAndDeleteAsyncHandler(seq uint64) (*asyncHandler, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *Client) clearAsyncHandler() { _ = "STUB: not implemented"; return }

func (c *Client) deleteStream(id uint64, local bool) { _ = "STUB: not implemented"; return }

func (c *Client) getStreamAndPushMsg(id uint64, local, done bool) (stream *Stream, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (c *Client) clearStream() { _ = "STUB: not implemented"; return }

func (c *Client) run() { _ = "STUB: not implemented"; return }

func (c *Client) initReader() { _ = "STUB: not implemented"; return }

func (c *Client) recvLoop() { _ = "STUB: not implemented"; return }

// if c.running {
// 	log.Info("%v\t%v\tReconnect Start", c.Handler.LogTag(), addr)
// }

func (c *Client) sendLoop() { _ = "STUB: not implemented"; return }

func (c *Client) normalSendLoop() { _ = "STUB: not implemented"; return }

// clear msg in send queue

func (c *Client) batchSendLoop() { _ = "STUB: not implemented"; return }

// clear msg in send queue

func newClientWithConn(conn net.Conn, codec codec.Codec, handler Handler, onStop func(*Client)) *Client {
	_ = "STUB: not implemented"
	return nil
}

// NewClient creates a Client.
func NewClient(dialer DialerFunc, args ...interface{}) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ClientPool represents an arpc Client Pool.
type ClientPool struct {
	size    uint64
	round   uint64
	clients []*Client
}

// Size returns Client number.
func (pool *ClientPool) Size() int { _ = "STUB: not implemented"; return 0 }

// Get returns a Client by index.
func (pool *ClientPool) Get(index int) *Client { _ = "STUB: not implemented"; return nil }

// Next returns a Client by round robin.
func (pool *ClientPool) Next() *Client { _ = "STUB: not implemented"; return nil }

// Handler returns Handler.
func (pool *ClientPool) Handler() Handler { _ = "STUB: not implemented"; return *new(Handler) }

// Stop stops all clients.
func (pool *ClientPool) Stop() { _ = "STUB: not implemented"; return }

// NewClientPool creates a ClientPool.
func NewClientPool(dialer DialerFunc, size int, args ...interface{}) (*ClientPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewClientPoolFromDialers creates a ClientPool by multiple dialers.
func NewClientPoolFromDialers(dialers []DialerFunc, args ...interface{}) (*ClientPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
