package wasmal

// BiquadFilterNode as described here:
// https://www.w3.org/TR/webaudio-1.1/#BiquadFilterNode
type BiquadFilterNode interface {
	AudioNode

	// Type returns the type of the filter.
	Type() BiquadFilterType

	// SetType sets the type of the filter.
	SetType(typ BiquadFilterType)

	// Frequency returns the AudioParam representing the frequency of the filter.
	Frequency() AudioParam

	// Detune returns the AudioParam representing the detune of the filter.
	Detune() AudioParam

	// Q returns the AudioParam representing the Q factor of the filter.
	Q() AudioParam

	// Gain returns the AudioParam representing the gain of the filter.
	Gain() AudioParam
}

// BiquadFilterType as described here:
// https://www.w3.org/TR/webaudio-1.1/#enumdef-biquadfiltertype
type BiquadFilterType string

const (
	BiquadFilterTypeLowpass   BiquadFilterType = "lowpass"
	BiquadFilterTypeHighpass  BiquadFilterType = "highpass"
	BiquadFilterTypeBandpass  BiquadFilterType = "bandpass"
	BiquadFilterTypeLowshelf  BiquadFilterType = "lowshelf"
	BiquadFilterTypeHighshelf BiquadFilterType = "highshelf"
	BiquadFilterTypePeaking   BiquadFilterType = "peaking"
	BiquadFilterTypeNotch     BiquadFilterType = "notch"
	BiquadFilterTypeAllpass   BiquadFilterType = "allpass"
)

var _ BiquadFilterNode = (*goBiquadFilterNode)(nil)

type goBiquadFilterNode struct {
	goAudioNode
}

func (g *goBiquadFilterNode) Type() BiquadFilterType {
	return BiquadFilterType(g.jsValue.Get("type").String())
}

func (g *goBiquadFilterNode) SetType(typ BiquadFilterType) {
	g.jsValue.Set("type", string(typ))
}

func (g *goBiquadFilterNode) Frequency() AudioParam {
	jsValue := g.jsValue.Get("frequency")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goBiquadFilterNode) Detune() AudioParam {
	jsValue := g.jsValue.Get("detune")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goBiquadFilterNode) Q() AudioParam {
	jsValue := g.jsValue.Get("Q")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goBiquadFilterNode) Gain() AudioParam {
	jsValue := g.jsValue.Get("gain")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}
