package wasmal

// GainNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/GainNode
type GainNode interface {
	AudioNode

	Gain() AudioParam
}
