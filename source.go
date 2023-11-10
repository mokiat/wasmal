package wasmal

// AudioScheduledSourceNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/AudioScheduledSourceNode
type AudioScheduledSourceNode interface {
	AudioNode

	Start(when float64)
	Stop(when float64)
}

var _ AudioScheduledSourceNode = goAudioScheduledSourceNode{}

type goAudioScheduledSourceNode struct {
	goAudioNode
}

func (g goAudioScheduledSourceNode) Start(when float64) {
	g.jsValue.Call("start", when)
}

func (g goAudioScheduledSourceNode) Stop(when float64) {
	g.jsValue.Call("stop", when)
}
