package wasmal

// DynamicsCompressorNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/DynamicsCompressorNode
type DynamicsCompressorNode interface {
	AudioNode

	Threshold() AudioParam
	Knee() AudioParam
	Ratio() AudioParam
	Reduction() float64
	Attack() AudioParam
	Release() AudioParam
}

var _ DynamicsCompressorNode = goDynamicsCompressorNode{}

type goDynamicsCompressorNode struct {
	goAudioNode
}

func (g goDynamicsCompressorNode) Threshold() AudioParam {
	jsValue := g.jsValue.Get("threshold")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g goDynamicsCompressorNode) Knee() AudioParam {
	jsValue := g.jsValue.Get("knee")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g goDynamicsCompressorNode) Ratio() AudioParam {
	jsValue := g.jsValue.Get("ratio")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g goDynamicsCompressorNode) Reduction() float64 {
	return g.jsValue.Get("reduction").Float()
}

func (g goDynamicsCompressorNode) Attack() AudioParam {
	jsValue := g.jsValue.Get("attack")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g goDynamicsCompressorNode) Release() AudioParam {
	jsValue := g.jsValue.Get("release")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}
