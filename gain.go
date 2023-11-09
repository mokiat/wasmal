package wasmal

// GainNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/GainNode
type GainNode interface {
	AudioNode

	Gain() AudioParam
}

var _ GainNode = goGainNode{}

type goGainNode struct {
	goAudioNode
}

func (g goGainNode) Gain() AudioParam {
	jsValue := g.jsValue.Get("gain")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}
