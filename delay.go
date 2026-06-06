package wasmal

// DelayNode as described here:
// https://www.w3.org/TR/webaudio-1.1/#DelayNode
type DelayNode interface {
	AudioNode

	// DelayTime returns the delay time parameter of this DelayNode.
	DelayTime() AudioParam
}

var _ DelayNode = (*goDelayNode)(nil)

type goDelayNode struct {
	goAudioNode
}

func (g *goDelayNode) DelayTime() AudioParam {
	jsValue := g.jsValue.Get("delayTime")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}
