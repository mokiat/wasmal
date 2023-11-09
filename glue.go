package wasmal

import (
	"errors"
	"syscall/js"
)

type Promise[T any] interface {
	Then(cb func(value T)) Promise[T]
	Catch(cb func(err error)) Promise[T]
}

type Float32Array interface {
}

var _ Promise[struct{}] = goPromise[struct{}]{}

type goPromise[T any] struct {
	goObject
	convert func(value js.Value) T
}

func (g goPromise[T]) Then(cb func(value T)) Promise[T] {
	var jsFunc js.Func
	jsFunc = js.FuncOf(func(this js.Value, args []js.Value) any {
		defer jsFunc.Release()
		cb(g.convert(args[0]))
		return nil
	})
	g.jsValue.Call("then", jsFunc)
	return g
}

func (g goPromise[T]) Catch(cb func(err error)) Promise[T] {
	var jsFunc js.Func
	jsFunc = js.FuncOf(func(this js.Value, args []js.Value) any {
		defer jsFunc.Release()
		cb(errors.New(args[0].String()))
		return nil
	})
	g.jsValue.Call("catch", jsFunc)
	return g
}
