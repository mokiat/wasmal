package wasmal

// GainNode as described here:
// https://www.w3.org/TR/webaudio-1.1/#GainNode
type GainNode interface {
	AudioNode

	// Gain returns an AudioParam representing the amount of gain to apply.
	Gain() AudioParam
}

var _ GainNode = (*goGainNode)(nil)

type goGainNode struct {
	goAudioNode
}

func (g *goGainNode) Gain() AudioParam {
	jsValue := g.jsValue.Get("gain")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}
