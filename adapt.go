package wasmal

import (
	"syscall/js"
	"unsafe"
)

// CleanupFunc is a function that must be called to release resources.
// This is specific to the WebAssembly adapter and is not part of the official API.
type CleanupFunc func()

// Promise represents a JavaScript Promise object in Go.
type Promise[T any] interface {

	// Then registers a callback to be called when the Promise is resolved with
	// a value of type T.
	//
	// If the Promise is rejected, the callback is not invoked and the rejection
	// is silently discarded. Register a callback through Catch to observe errors.
	//
	// The returned CleanupFunc must be called to release resources associated
	// with the callback, regardless of whether the Promise is resolved or rejected.
	Then(cb func(value T)) CleanupFunc

	// Catch registers a callback to be called when the Promise is rejected with
	// an error.
	//
	// The returned CleanupFunc must be called to release resources associated
	// with the callback, regardless of whether the Promise is resolved or rejected.
	Catch(cb func(err error)) CleanupFunc
}

var _ Promise[struct{}] = (*goPromise[struct{}])(nil)

type goPromise[T any] struct {
	goObject
	convert func(value js.Value) T
}

func (g *goPromise[T]) Then(cb func(value T)) CleanupFunc {
	jsFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		cb(g.convert(args[0]))
		return nil
	})
	g.jsValue.Call("then", jsFunc, noopCallback)
	return jsFunc.Release
}

func (g *goPromise[T]) Catch(cb func(err error)) CleanupFunc {
	jsFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		cb(js.Error{
			Value: args[0],
		})
		return js.Undefined()
	})
	g.jsValue.Call("catch", jsFunc)
	return jsFunc.Release
}

var noopCallback = js.FuncOf(func(this js.Value, args []js.Value) any {
	return nil
})

// DataTypes represents allowed data slice types.
type DataTypes interface {
	~int8 | ~uint8 | ~int16 | ~uint16 | ~int32 | ~uint32 | ~float32 | ~float64
}

// asByteSlice returns a []byte representation for the
// specified arbitrary slice type.
//
// This utility function is related to the following issues:
// https://github.com/golang/go/issues/32402
// https://github.com/golang/go/issues/31980
func asByteSlice[T DataTypes](data []T) []byte {
	if len(data) == 0 {
		return nil
	}
	dataSize := byteSize(data)
	return unsafe.Slice((*byte)(unsafe.Pointer(&data[0])), dataSize)
}

// byteSize returns the number of bytes that would be
// needed to represent data once it is converted to a
// byte slice through asByteSlice.
func byteSize[T DataTypes](data []T) int {
	if len(data) == 0 {
		return 0
	}
	return len(data) * int(unsafe.Sizeof(data[0]))
}
