package wasmal

// PannerNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/PannerNode
type PannerNode interface {
	AudioNode
}

// StereoPannerNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/StereoPannerNode
type StereoPannerNode interface {
	AudioNode
	Pan() AudioParam
}

var _ StereoPannerNode = goStereoPannerNode{}

type goStereoPannerNode struct {
	goAudioNode
}

func (g goStereoPannerNode) Pan() AudioParam {
	jsValue := g.jsValue.Get("pan")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}
