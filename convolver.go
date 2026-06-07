package wasmal

// ConvolverNode as described here:
// https://www.w3.org/TR/webaudio-1.1/#ConvolverNode
type ConvolverNode interface {
	AudioNode

	// Buffer returns the AudioBuffer containing the impulse response used by this
	// ConvolverNode.
	Buffer() AudioBuffer

	// SetBuffer sets the AudioBuffer containing the impulse response used by this
	// ConvolverNode.
	SetBuffer(buffer AudioBuffer)

	// Normalize returns whether the impulse response is scaled by an equal-power
	// normalization.
	Normalize() bool

	// SetNormalize sets whether the impulse response is scaled by an equal-power
	// normalization.
	SetNormalize(normalize bool)
}

var _ ConvolverNode = (*goConvolverNode)(nil)

type goConvolverNode struct {
	goAudioNode
}

func (g *goConvolverNode) Buffer() AudioBuffer {
	jsValue := g.jsValue.Get("buffer")
	return &goAudioBuffer{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goConvolverNode) SetBuffer(buffer AudioBuffer) {
	g.jsValue.Set("buffer", buffer.ref())
}

func (g *goConvolverNode) Normalize() bool {
	return g.jsValue.Get("normalize").Bool()
}

func (g *goConvolverNode) SetNormalize(normalize bool) {
	g.jsValue.Set("normalize", normalize)
}
