package wasmal

// ConvolverNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/ConvolverNode
type ConvolverNode interface {
	AudioNode

	Buffer() AudioBuffer
	SetBuffer(buffer AudioBuffer)
	Normalize() bool
	SetNormalize(normalize bool)
}

var _ ConvolverNode = goConvolverNode{}

type goConvolverNode struct {
	goAudioNode
}

func (g goConvolverNode) Buffer() AudioBuffer {
	jsValue := g.jsValue.Get("buffer")
	return goAudioBuffer{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g goConvolverNode) SetBuffer(buffer AudioBuffer) {
	g.jsValue.Set("buffer", buffer.ref())
}

func (g goConvolverNode) Normalize() bool {
	return g.jsValue.Get("normalize").Bool()
}

func (g goConvolverNode) SetNormalize(normalize bool) {
	g.jsValue.Set("normalize", normalize)
}
