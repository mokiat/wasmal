package wasmal

// DelayNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/DelayNode
type DelayNode interface {
	AudioNode

	DelayTime() AudioParam
}

var _ DelayNode = goDelayNode{}

type goDelayNode struct {
	goAudioNode
}

func (g goDelayNode) DelayTime() AudioParam {
	jsValue := g.jsValue.Get("delayTime")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}
