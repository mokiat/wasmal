package wasmal

// AudioDestinationNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/AudioDestinationNode
type AudioDestinationNode interface {
	AudioNode

	MaxChannelCount() uint32
}

var _ AudioDestinationNode = goAudioDestinationNode{}

type goAudioDestinationNode struct {
	goAudioNode
}

func (g goAudioDestinationNode) MaxChannelCount() uint32 {
	return uint32(g.jsValue.Get("maxChannelCount").Int())
}
