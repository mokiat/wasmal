package wasmal

type IIRFilterNode interface {
	AudioNode
}

var _ IIRFilterNode = (*goIIRFilterNode)(nil)

type goIIRFilterNode struct {
	goAudioNode
}
