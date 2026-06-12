package wasmal

// IIRFilterNode as described here:
// https://www.w3.org/TR/webaudio-1.1/#IIRFilterNode
type IIRFilterNode interface {
	AudioNode
}

var _ IIRFilterNode = (*goIIRFilterNode)(nil)

type goIIRFilterNode struct {
	goAudioNode
}
