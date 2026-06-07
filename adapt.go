package wasmal

import (
	"syscall/js"
)

// CleanupFunc is a function that must be called to release resources.
// This is specific to the WebAssembly adapter and is not part of the official API.
type CleanupFunc func()

// Promise represents a JavaScript Promise object in Go.
type Promise[T any] interface {

	// Then registers a callback to be called when the Promise is resolved with
	// a value of type T.
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
