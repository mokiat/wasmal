package wasmal

// AudioDestinationNode as described here:
// https://www.w3.org/TR/webaudio-1.1/#AudioDestinationNode
type AudioDestinationNode interface {
	AudioNode

	// MaxChannelCount returns the maximum number of channels that the
	// destination node can handle.
	MaxChannelCount() uint32
}

var _ AudioDestinationNode = (*goAudioDestinationNode)(nil)

type goAudioDestinationNode struct {
	goAudioNode
}

func (g *goAudioDestinationNode) MaxChannelCount() uint32 {
	return uint32(g.jsValue.Get("maxChannelCount").Int())
}
