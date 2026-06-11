package wasmal

import (
	"syscall/js"
)

// Float32Array as described here:
// https://webidl.spec.whatwg.org/#idl-Float32Array
type Float32Array interface {
	Get(index int) float32
	Set(index int, value float32)
	CopyFrom(data []float32) // not per spec
}

var _ Float32Array = (*goFloat32Array)(nil)

type goFloat32Array struct {
	goObject
}

func (g *goFloat32Array) Get(index int) float32 {
	return float32(g.jsValue.Index(index).Float())
}

func (g *goFloat32Array) Set(index int, value float32) {
	g.jsValue.SetIndex(index, value)
}

func (g *goFloat32Array) CopyFrom(data []float32) {
	if len(data) == 0 {
		return
	}
	uint8View := js.Global().Get("Uint8Array").New(
		g.jsValue.Get("buffer"),
		g.jsValue.Get("byteOffset"),
		len(data)*4,
	)
	byteData := asByteSlice(data)
	js.CopyBytesToJS(uint8View, byteData)
}
