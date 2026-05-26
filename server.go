// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package arpc

import (
	"context"
	"net"
	"sync"

	"github.com/lesismal/arpc/codec"
	"github.com/lesismal/arpc/util"
)

// Server represents an arpc Server.
type Server struct {
	Accepted int64
	CurrLoad int64
	MaxLoad  int64

	// 64-aligned on 32-bit
	seq uint64

	Codec   codec.Codec
	Handler Handler

	Listener net.Listener

	mux sync.Mutex

	running bool
	chStop  chan error
	clients map[*Client]util.Empty
}

// Serve starts service with listener.
func (s *Server) Serve(ln net.Listener) error { _ = "STUB: not implemented"; return nil }

// Run starts tcp service on addr.
func (s *Server) Run(addr string) error { _ = "STUB: not implemented"; return nil }

// defer log.Info("%v Stopped", s.Handler.LogTag())

func (s *Server) Broadcast(method string, v interface{}, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) BroadcastWithFilter(method string, v interface{}, filter func(*Client) bool, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (s *Server) ForEach(h func(*Client)) { _ = "STUB: not implemented"; return }

func (s *Server) ForEachWithFilter(h func(*Client), filter func(*Client) bool) {
	_ = "STUB: not implemented"
	return
}

// Stop stops service.
func (s *Server) Stop() error { _ = "STUB: not implemented"; return nil }

// Shutdown shutdown service.
func (s *Server) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// NewMessage creates a Message.
func (s *Server) NewMessage(cmd byte, method string, v interface{}, args ...interface{}) *Message {
	_ = "STUB: not implemented"
	return nil
}

func (s *Server) addLoad() int64 { _ = "STUB: not implemented"; return 0 }

func (s *Server) subLoad() int64 { _ = "STUB: not implemented"; return 0 }

func (s *Server) addClient(c *Client) { _ = "STUB: not implemented"; return }

func (s *Server) deleteClient(c *Client) { _ = "STUB: not implemented"; return }

func (s *Server) clearClients() { _ = "STUB: not implemented"; return }

func (s *Server) runLoop() error { _ = "STUB: not implemented"; return nil }

// NewServer creates an arpc Server.
func NewServer() *Server { _ = "STUB: not implemented"; return nil }
