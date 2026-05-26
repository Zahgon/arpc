// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package arpc

import (
	"sync"

	"github.com/lesismal/arpc/codec"
)

const (
	// CmdNone is invalid
	CmdNone byte = 0

	// CmdRequest the other side should response to a request message
	CmdRequest byte = 1

	// CmdResponse the other side should not response to a request message
	CmdResponse byte = 2

	// CmdNotify the other side should not response to a request message
	CmdNotify byte = 3

	// CmdPing .
	CmdPing byte = 4

	// CmdPong .
	CmdPong byte = 5

	// CmdStream .
	CmdStream byte = 6
)

const (
	// HeaderIndexBodyLenBegin .
	HeaderIndexBodyLenBegin = 0
	// HeaderIndexBodyLenEnd .
	HeaderIndexBodyLenEnd = 4
	// HeaderIndexReserved .
	HeaderIndexReserved = 4
	// HeaderIndexCmd .
	HeaderIndexCmd = 5
	// HeaderIndexFlag .
	HeaderIndexFlag = 6
	// HeaderIndexMethodLen .
	HeaderIndexMethodLen = 7
	// HeaderIndexSeqBegin .
	HeaderIndexSeqBegin = 8
	// HeaderIndexSeqEnd .
	HeaderIndexSeqEnd = 16
	// HeaderFlagMaskError .
	HeaderFlagMaskError byte = 0x01
	// HeaderFlagMaskAsync .
	HeaderFlagMaskAsync byte = 0x02

	HeaderStreamLocalBitIndex = 7
	HeaderStreamEOFBitIndex   = 6
	HeaderStreamLocalBit      = byte(0x1) << HeaderStreamLocalBitIndex
	HeaderStreamEOFBit        = byte(0x1) << HeaderStreamEOFBitIndex
	HeaderStreamFlagBitMask   = HeaderStreamLocalBit | HeaderStreamEOFBit
	HeaderCmdBitMask          = ^HeaderStreamFlagBitMask
)

const (
	// HeadLen represents Message head length.
	HeadLen int = 16

	// MaxMethodLen limits Message method length.
	MaxMethodLen int = 127

	// DefaultMaxBodyLen limits Message body length.
	DefaultMaxBodyLen int = 1024*1024*64 - 16
)

var (
	// PingMessage .
	PingMessage = newMessage(CmdPing, "", nil, false, false, 0, nil, nil, nil)

	// PongMessage .
	PongMessage = newMessage(CmdPong, "", nil, false, false, 0, nil, nil, nil)
)

// Header defines Message head
type Header []byte

// BodyLen returns Message body length.
func (h Header) BodyLen() int { _ = "STUB: not implemented"; return 0 }

// message creates a Message by body length.
func (h Header) message(handler Handler) (*Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// msg := &Message{Buffer: handler.Malloc(HeadLen + bodyLen)}

var (
	messagePool = sync.Pool{
		New: func() interface{} {
			return &Message{}
		},
	}

	emptyMessage = Message{}
)

// Message represents an arpc Message.
type Message struct {
	// 64-aligned on 32-bit
	ref int32

	Buffer []byte

	handler Handler
	values  map[interface{}]interface{}
}

// Retain increment the reference count and returns the current value.
func (m *Message) Retain() int32 { _ = "STUB: not implemented"; return 0 }

// Release decrement the reference count and returns the current value.
func (m *Message) Release() int32 { _ = "STUB: not implemented"; return 0 }

// ResetAttrs resets reserved/cmd/flag/methodLen to 0.
func (m *Message) ResetAttrs() { _ = "STUB: not implemented"; return }

// Payback put Message to the pool.
func (m *Message) Payback() { _ = "STUB: not implemented"; return }

// Len returns total length of buffer.
func (m *Message) Len() int { _ = "STUB: not implemented"; return 0 }

// Cmd returns cmd.
func (m *Message) Cmd() byte { _ = "STUB: not implemented"; return 0 }

// SetCmd sets cmd.
func (m *Message) SetCmd(cmd byte) { _ = "STUB: not implemented"; return }

// // IsStream represents whether it's a stream message.
// func (m *Message) IsStream() bool {
// 	return m.Buffer[HeaderIndexCmd]&HeaderStreamBit > 0
// }

// // SetStream sets the flag for a stream message.
// func (m *Message) SetStream(isStream bool) {
// 	if isStream {
// 		m.Buffer[HeaderIndexCmd] |= HeaderStreamBit
// 	} else {
// 		m.Buffer[HeaderIndexCmd] &= (^HeaderStreamBit)
// 	}
// }

// IsStream represents whether it's a stream message.
func (m *Message) IsStreamLocal() bool { _ = "STUB: not implemented"; return false }

// SetStream sets the flag for a stream message.
func (m *Message) SetStreamLocal(local bool) { _ = "STUB: not implemented"; return }

// IsStream represents whether it's a stream's last message and the stream is EOF and closed.
func (m *Message) IsStreamEOF() bool { _ = "STUB: not implemented"; return false }

// SetStream sets the flag for a stream's last message and mark the stream is EOF and closed.
func (m *Message) SetStreamEOF(eof bool) { _ = "STUB: not implemented"; return }

// IsError returns error flag.
func (m *Message) IsError() bool { _ = "STUB: not implemented"; return false }

// SetError sets error flag.
func (m *Message) SetError(isError bool) { _ = "STUB: not implemented"; return }

// Error returns error.
func (m *Message) Error() error { _ = "STUB: not implemented"; return nil }

// IsAsync returns async flag.
func (m *Message) IsAsync() bool { _ = "STUB: not implemented"; return false }

// SetAsync sets async flag.
func (m *Message) SetAsync(isAsync bool) { _ = "STUB: not implemented"; return }

// Values returns values.
func (m *Message) Values() map[interface{}]interface{} {
	_ = "STUB: not implemented"

	// SetFlagBit sets flag bit value by index.
	return nil
}

func (m *Message) SetFlagBit(index int, value bool) error { _ = "STUB: not implemented"; return nil }

// case 8, 9:
// 	if value {
// 		m.Buffer[HeaderIndexFlag] |= (0x1 << (index - 2))
// 	} else {
// 		m.Buffer[HeaderIndexFlag] &= (^(0x1 << (index - 2)))
// 	}
// 	return nil

// IsFlagBitSet returns flag bit value.
func (m *Message) IsFlagBitSet(index int) bool { _ = "STUB: not implemented"; return false }

// case 8, 9:
// 	return (m.Buffer[HeaderIndexFlag] & (0x1 << (index - 2))) != 0

// MethodLen returns method length.
func (m *Message) MethodLen() int { _ = "STUB: not implemented"; return 0 }

// SetMethodLen sets method length.
func (m *Message) SetMethodLen(l int) { _ = "STUB: not implemented"; return }

// Method returns method.
func (m *Message) Method() string { _ = "STUB: not implemented"; return "" }

func (m *Message) method() string { _ = "STUB: not implemented"; return "" }

// BodyLen returns body length.
func (m *Message) BodyLen() int { _ = "STUB: not implemented"; return 0 }

// SetBodyLen sets body length.
func (m *Message) SetBodyLen(l int) { _ = "STUB: not implemented"; return }

// Seq returns sequence number.
func (m *Message) Seq() uint64 { _ = "STUB: not implemented"; return 0 }

// SetSeq sets sequence number.
func (m *Message) SetSeq(seq uint64) { _ = "STUB: not implemented"; return }

// Data returns payload data after method.
func (m *Message) Data() []byte { _ = "STUB: not implemented"; return nil }

// Get returns value for key.
func (m *Message) Get(key interface{}) (interface{}, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// Set sets key-value pair.
func (m *Message) Set(key interface{}, value interface{}) { _ = "STUB: not implemented"; return }

// NewMessage creates a Message.
func NewMessage(cmd byte, method string, v interface{}, isError bool, isAsync bool, seq uint64, h Handler, codec codec.Codec, values map[interface{}]interface{}) *Message {
	_ = "STUB: not implemented"
	return nil
}

// newMessage creates a Message.
func newMessage(cmd byte, method string, v interface{}, isError bool, isAsync bool, seq uint64, h Handler, codec codec.Codec, values map[interface{}]interface{}) *Message {
	_ = "STUB: not implemented"
	return nil
}

// msg = &Message{Buffer: h.Malloc(HeadLen + bodyLen), values: values}

func checkMethod(method string) error { _ = "STUB: not implemented"; return nil }

// MessageCoder defines Message coding middleware interface.
type MessageCoder interface {
	// Encode wrap message before send to client
	Encode(*Client, *Message) *Message
	// Decode unwrap message between recv and handle
	Decode(*Client, *Message) *Message
}
