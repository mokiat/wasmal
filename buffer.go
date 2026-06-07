package wasmal

// AudioBuffer as described here:
// https://www.w3.org/TR/webaudio-1.1/#AudioBuffer
type AudioBuffer interface {
	object

	// SampleRate is the sample rate of the PCM audio data in the buffer, in samples per second.
	SampleRate() float32

	// Length is the length of the PCM audio data in the buffer, in sample-frames.
	Length() uint32

	// Duration is the duration of the PCM audio data in the buffer, in seconds.
	Duration() float64

	// NumberOfChannels is the number of discrete audio channels described by the buffer.
	NumberOfChannels() uint32

	// GetChannelData returns a Float32Array containing the PCM audio data associated with the specified channel.
	GetChannelData(channel uint32) Float32Array
}

var _ AudioBuffer = (*goAudioBuffer)(nil)

type goAudioBuffer struct {
	goObject
}

func (g *goAudioBuffer) SampleRate() float32 {
	return float32(g.jsValue.Get("sampleRate").Float())
}

func (g *goAudioBuffer) Length() uint32 {
	return uint32(g.jsValue.Get("length").Int())
}

func (g *goAudioBuffer) Duration() float64 {
	return g.jsValue.Get("duration").Float()
}

func (g *goAudioBuffer) NumberOfChannels() uint32 {
	return uint32(g.jsValue.Get("numberOfChannels").Int())
}

func (g *goAudioBuffer) GetChannelData(channel uint32) Float32Array {
	jsValue := g.jsValue.Call("getChannelData", channel)
	return &goFloat32Array{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}
