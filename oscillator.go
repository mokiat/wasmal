package wasmal

// OscillatorNode as described here:
// https://www.w3.org/TR/webaudio-1.1/#OscillatorNode
type OscillatorNode interface {
	AudioScheduledSourceNode

	// Type returns the type of the oscillator. The default value is "sine".
	Type() OscillatorType

	// SetType sets the type of the oscillator.
	SetType(oType OscillatorType)

	// Frequency returns an AudioParam representing the frequency of the oscillator.
	Frequency() AudioParam

	// Detune returns an AudioParam representing the detuning of the oscillator in cents.
	Detune() AudioParam
}

// OscillatorType as described here:
// https://www.w3.org/TR/webaudio-1.1/#enumdef-oscillatortype
type OscillatorType string

const (
	OscillatorTypeSine     OscillatorType = "sine"
	OscillatorTypeSquare   OscillatorType = "square"
	OscillatorTypeSawtooth OscillatorType = "sawtooth"
	OscillatorTypeTriangle OscillatorType = "triangle"
	OscillatorTypeCustom   OscillatorType = "custom"
)

var _ OscillatorNode = (*goOscillatorNode)(nil)

type goOscillatorNode struct {
	goAudioScheduledSourceNode
}

func (g *goOscillatorNode) Type() OscillatorType {
	return OscillatorType(g.jsValue.Get("type").String())
}

func (g *goOscillatorNode) SetType(oType OscillatorType) {
	g.jsValue.Set("type", string(oType))
}

func (g *goOscillatorNode) Frequency() AudioParam {
	jsValue := g.jsValue.Get("frequency")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goOscillatorNode) Detune() AudioParam {
	jsValue := g.jsValue.Get("detune")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}
