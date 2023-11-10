package wasmal

// OscillatorNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/OscillatorNode
type OscillatorNode interface {
	AudioScheduledSourceNode

	Frequency() AudioParam
	Detune() AudioParam
	Type() OscillatorType
	SetType(oType OscillatorType)
}

// OscillatorType as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/OscillatorNode/type
type OscillatorType string

const (
	OscillatorTypeSine     OscillatorType = "sine"
	OscillatorTypeSquare   OscillatorType = "square"
	OscillatorTypeSawTooth OscillatorType = "sawtooth"
	OscillatorTypeTriangle OscillatorType = "triangle"
	OscillatorTypeCustom   OscillatorType = "custom"
)

var _ OscillatorNode = goOscillatorNode{}

type goOscillatorNode struct {
	goAudioScheduledSourceNode
}

func (g goOscillatorNode) Frequency() AudioParam {
	jsValue := g.jsValue.Get("frequency")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g goOscillatorNode) Detune() AudioParam {
	jsValue := g.jsValue.Get("detune")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g goOscillatorNode) Type() OscillatorType {
	return OscillatorType(g.jsValue.Get("type").String())
}

func (g goOscillatorNode) SetType(oType OscillatorType) {
	g.jsValue.Set("type", string(oType))
}
