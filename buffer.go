package wasmal

// AudioBuffer as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/AudioBuffer
type AudioBuffer interface {
	object

	SampleRate() float64
	Length() uint
	Duration() float64
	NumberOfChannels() uint

	GetChannelData(channel uint) Float32Array
}

// AudioBufferSourceNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/AudioBufferSourceNode
type AudioBufferSourceNode interface {
}
