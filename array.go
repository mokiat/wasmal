package wasmal

// Float32Array as described here:
// https://webidl.spec.whatwg.org/#idl-Float32Array
type Float32Array interface {
	Get(index int) float32
	Set(index int, value float32)
}

var _ Float32Array = (*goFloat32Array)(nil)

type goFloat32Array struct {
	goObject
}

// Get returns the value at the given index.
func (g goFloat32Array) Get(index int) float32 {
	return float32(g.jsValue.Index(index).Float())
}

// Set sets the value at the given index.
func (g goFloat32Array) Set(index int, value float32) {
	g.jsValue.SetIndex(index, value)
}
