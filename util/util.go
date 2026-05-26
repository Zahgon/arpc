// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package util

import (
	acodec "github.com/lesismal/arpc/codec"
)

// Empty struct
type Empty struct{}

// Recover handles panic and logs stack info
func Recover() { _ = "STUB: not implemented"; return }

// Safe wraps a function-calling with panic recovery
func Safe(call func()) { _ = "STUB: not implemented"; return }

// StrToBytes hacks string to []byte
func StrToBytes(s string) []byte { _ = "STUB: not implemented"; return nil }

// BytesToStr hacks []byte to string
func BytesToStr(b []byte) string { _ = "STUB: not implemented"; return "" }

// ValueToBytes converts values to []byte
func ValueToBytes(codec acodec.Codec, v interface{}) []byte { _ = "STUB: not implemented"; return nil }

// BytesToValue converts []byte to values
func BytesToValue(codec acodec.Codec, data []byte, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
