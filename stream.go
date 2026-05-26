package arpc

import (
	"context"
)

// Stream .
type Stream struct {
	id            uint64
	cli           *Client
	method        string
	local         bool
	chData        chan *Message
	stateRecv     int32
	stateSend     int32
	stateCloseCnt int32
}

func (s *Stream) Id() uint64 { _ = "STUB: not implemented"; return 0 }

func (s *Stream) onMessage(msg *Message) { _ = "STUB: not implemented"; return }

func (s *Stream) CloseRecv() { _ = "STUB: not implemented"; return }

func (s *Stream) CloseRecvContext(ctx context.Context) { _ = "STUB: not implemented"; return }

func (s *Stream) CloseSend() { _ = "STUB: not implemented"; return }

func (s *Stream) CloseSendContext(ctx context.Context) { _ = "STUB: not implemented"; return }

func (s *Stream) Recv(v interface{}) error { _ = "STUB: not implemented"; return nil }

func (s *Stream) RecvContext(ctx context.Context, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stream) RecvWith(ctx context.Context, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stream) newMessage(v interface{}, args ...interface{}) *Message {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stream) Send(v interface{}, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stream) SendContext(ctx context.Context, v interface{}, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stream) SendWith(ctx context.Context, v interface{}, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stream) SendAndClose(v interface{}, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stream) SendAndCloseContext(ctx context.Context, v interface{}, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stream) SendAndCloseWith(ctx context.Context, v interface{}, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stream) checkStateAndSend(ctx context.Context, v interface{}, eof bool, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stream) send(ctx context.Context, v interface{}, eof bool, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// c.Handler.OnOverstock(c, msg)

func (s *Stream) halfClose() { _ = "STUB: not implemented"; return }

// NewStream creates a stream.
func (client *Client) NewStream(method string) *Stream { _ = "STUB: not implemented"; return nil }

func (client *Client) newStream(method string, id uint64, local bool) *Stream {
	_ = "STUB: not implemented"
	return nil
}
