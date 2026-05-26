// Copyright 2020 lesismal. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package arpc

import (
	"sync"
)

type Allocator interface {
	Malloc(size int) []byte
	Realloc(buf []byte, size int) []byte
	Append(buf []byte, more ...byte) []byte
	AppendString(buf []byte, more string) []byte
	Free(buf []byte)
}

// DefaultAllocator .
var DefaultAllocator Allocator = New(64, 64)

// BufferPool .
type BufferPool struct {
	Debug bool
	mux   sync.Mutex

	smallSize int
	bigSize   int
	smallPool *sync.Pool
	bigPool   *sync.Pool

	allocCnt    uint64
	freeCnt     uint64
	allocStacks map[uintptr]string
}

// New .
func New(smallSize, bigSize int) Allocator { _ = "STUB: not implemented"; return *new(Allocator) }

// Debug:       true,

// Malloc .
func (bp *BufferPool) Malloc(size int) []byte { _ = "STUB: not implemented"; return nil }

// Realloc .
func (bp *BufferPool) Realloc(buf []byte, size int) []byte { _ = "STUB: not implemented"; return nil }

func (bp *BufferPool) reallocDebug(buf []byte, size int) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Append .
func (bp *BufferPool) Append(buf []byte, more ...byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

// AppendString .
func (bp *BufferPool) AppendString(buf []byte, more string) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (bp *BufferPool) appendStringDebug(buf []byte, more string) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Free .
func (bp *BufferPool) Free(buf []byte) { _ = "STUB: not implemented"; return }

func (bp *BufferPool) addAllocStack(ptr uintptr) { _ = "STUB: not implemented"; return }

func (bp *BufferPool) deleteAllocStack(ptr uintptr) { _ = "STUB: not implemented"; return }

func (bp *BufferPool) LogDebugInfo() { _ = "STUB: not implemented"; return }

// fmt.Println("---------------------------------------------------------")
// fmt.Println("Free")
// for s, n := range bp.freeStacks {
// 	fmt.Println("num:", n)
// 	fmt.Println("stack:\n", s)
// 	totalFree += n
// 	fmt.Println("---------------------------------------------------------")
// }

// NativeAllocator definition.
type NativeAllocator struct{}

// Malloc .
func (a *NativeAllocator) Malloc(size int) []byte { _ = "STUB: not implemented"; return nil }

// Realloc .
func (a *NativeAllocator) Realloc(buf []byte, size int) []byte {
	_ = "STUB: not implemented"
	return nil
}

// Free .
func (a *NativeAllocator) Free(buf []byte) {
	_ = "STUB: not implemented"

	// Malloc exports default package method.
	return
}

func Malloc(size int) []byte { _ = "STUB: not implemented"; return nil }

// Realloc exports default package method.
func Realloc(buf []byte, size int) []byte { _ = "STUB: not implemented"; return nil }

// Append exports default package method.
func Append(buf []byte, more ...byte) []byte { _ = "STUB: not implemented"; return nil }

// AppendString exports default package method.
func AppendString(buf []byte, more string) []byte { _ = "STUB: not implemented"; return nil }

// Free exports default package method.
func Free(buf []byte) { _ = "STUB: not implemented"; return }

// SetDebug .
func SetDebug(enable bool) { _ = "STUB: not implemented"; return }

// LogDebugInfo .
func LogDebugInfo() { _ = "STUB: not implemented"; return }

func getBufferPtr(buf []byte) uintptr { _ = "STUB: not implemented"; return 0 }

func getStack() string { _ = "STUB: not implemented"; return "" }
