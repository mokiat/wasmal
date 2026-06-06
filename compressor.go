package wasmal

// DynamicsCompressorNode as described here:
// https://www.w3.org/TR/webaudio-1.1/#DynamicsCompressorNode
type DynamicsCompressorNode interface {
	AudioNode

	// Threshold is the decibel value above which the compression will start
	// taking effect.
	Threshold() AudioParam

	// Knee is the decibel value representing the range above the threshold where
	// the curve smoothly transitions to the compressed portion.
	Knee() AudioParam

	// Ratio is the amount of change, in dB, needed in the input for a 1 dB change
	// in the output.
	Ratio() AudioParam

	// Reduction is the amount of gain reduction currently applied by the compressor
	// to the signal, in dB.
	Reduction() float32

	// Attack is the amount of time, in seconds, required to reduce the gain by 10 dB.
	Attack() AudioParam

	// Release is the amount of time, in seconds, required to increase the gain by 10 dB.
	Release() AudioParam
}

var _ DynamicsCompressorNode = (*goDynamicsCompressorNode)(nil)

type goDynamicsCompressorNode struct {
	goAudioNode
}

func (g *goDynamicsCompressorNode) Threshold() AudioParam {
	jsValue := g.jsValue.Get("threshold")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goDynamicsCompressorNode) Knee() AudioParam {
	jsValue := g.jsValue.Get("knee")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goDynamicsCompressorNode) Ratio() AudioParam {
	jsValue := g.jsValue.Get("ratio")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goDynamicsCompressorNode) Reduction() float32 {
	return float32(g.jsValue.Get("reduction").Float())
}

func (g *goDynamicsCompressorNode) Attack() AudioParam {
	jsValue := g.jsValue.Get("attack")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goDynamicsCompressorNode) Release() AudioParam {
	jsValue := g.jsValue.Get("release")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}
