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
