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

// newFloat32Array creates a new JavaScript Float32Array that holds
// a copy of the specified data, transferred in a single bulk copy.
func newFloat32Array(data []float32) js.Value {
	return newTypedArray("Float32Array", data)
}

// newFloat64Array creates a new JavaScript Float64Array that holds
// a copy of the specified data, transferred in a single bulk copy.
func newFloat64Array(data []float64) js.Value {
	return newTypedArray("Float64Array", data)
}

// newTypedArray creates a new instance of the specified JavaScript typed
// array constructor and copies data into it in a single bulk copy. The
// constructor must have an element type that matches T.
func newTypedArray[T DataTypes](constructor string, data []T) js.Value {
	jsArray := js.Global().Get(constructor).New(len(data))
	uint8View := js.Global().Get("Uint8Array").New(jsArray.Get("buffer"))
	js.CopyBytesToJS(uint8View, asByteSlice(data))
	return jsArray
}

// asByteSlice returns a []byte representation for the
// specified arbitrary slice type.
//
// This utility function is related to the following issues:
// https://github.com/golang/go/issues/32402
// https://github.com/golang/go/issues/31980
func asByteSlice[T DataTypes](data []T) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(data))), byteSize(data))
}

// byteSize returns the number of bytes that would be
// needed to represent data once it is converted to a
// byte slice through asByteSlice.
func byteSize[T DataTypes](data []T) int {
	var zero T
	return len(data) * int(unsafe.Sizeof(zero))
}
