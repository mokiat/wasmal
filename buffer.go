package wasmal

import "github.com/mokiat/gog/opt"

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
	AudioScheduledSourceNode

	Buffer() AudioBuffer
	SetBuffer(buffer AudioBuffer)
	Detune() AudioParam
	Loop() bool
	SetLoop(loop bool)
	LoopStart() float64
	SetLoopStart(start float64)
	LoopEnd() float64
	SetLoopEnd(end float64)
	PlaybackRate() AudioParam

	StartDetailed(when, offset float64, duration opt.T[float64])
}

var _ AudioBuffer = goAudioBuffer{}

type goAudioBuffer struct {
	goObject
}

func (g goAudioBuffer) SampleRate() float64 {
	return g.jsValue.Get("sampleRate").Float()
}

func (g goAudioBuffer) Length() uint {
	return uint(g.jsValue.Get("length").Int())
}

func (g goAudioBuffer) Duration() float64 {
	return g.jsValue.Get("duration").Float()
}

func (g goAudioBuffer) NumberOfChannels() uint {
	return uint(g.jsValue.Get("numberOfChannels").Int())
}

func (g goAudioBuffer) GetChannelData(channel uint) Float32Array {
	panic("TODO")
}

var _ AudioBufferSourceNode = goAudioBufferSourceNode{}

type goAudioBufferSourceNode struct {
	goAudioScheduledSourceNode
}

func (g goAudioBufferSourceNode) Buffer() AudioBuffer {
	jsValue := g.jsValue.Get("buffer")
	return goAudioBuffer{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g goAudioBufferSourceNode) SetBuffer(buffer AudioBuffer) {
	g.jsValue.Set("buffer", buffer.ref())
}

func (g goAudioBufferSourceNode) Detune() AudioParam {
	panic("TODO")
}

func (g goAudioBufferSourceNode) Loop() bool {
	panic("TODO")
}

func (g goAudioBufferSourceNode) SetLoop(loop bool) {
	panic("TODO")
}

func (g goAudioBufferSourceNode) LoopStart() float64 {
	panic("TODO")
}

func (g goAudioBufferSourceNode) SetLoopStart(start float64) {
	panic("TODO")
}

func (g goAudioBufferSourceNode) LoopEnd() float64 {
	panic("TODO")
}

func (g goAudioBufferSourceNode) SetLoopEnd(end float64) {
	panic("TODO")
}

func (g goAudioBufferSourceNode) PlaybackRate() AudioParam {
	panic("TODO")
}

func (g goAudioBufferSourceNode) StartDetailed(when, offset float64, duration opt.T[float64]) {
	panic("TODO")
}
