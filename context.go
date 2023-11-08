package wasmal

import "syscall/js"

const DefaultSampleRate = 44100

// BaseAudioContext as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext
type BaseAudioContext interface {
	object

	CurrentTime() float64
	Destination() AudioDestinationNode
	Listener() AudioListener
	SampleRate() float64
	State() AudioContextState

	CreateBuffer(numChannels, length, sampleRate uint) AudioBuffer
	CreateBufferSource() AudioBufferSourceNode
	CreateConvolver() ConvolverNode
	CreateDelay() DelayNode
	CreateDynamicsCompressor() DynamicsCompressorNode
	CreateGain() GainNode
	CreateOscillator() OscillatorNode
	CreatePanner() PannerNode
	CreateStereoPanner() StereoPannerNode
	DecodeAudioData(data []byte) Promise[AudioBuffer]
}

// AudioContextState as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/BaseAudioContext/state
type AudioContextState string

const (
	AudioContextStateSuspended AudioContextState = "suspended"
	AudioContextStateRunning   AudioContextState = "running"
	AudioContextStateClosed    AudioContextState = "closed"
)

// AudioContext as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/AudioContext
type AudioContext interface {
	BaseAudioContext

	BaseLatency() float64
	OutputLatency() float64

	Close() Promise[struct{}]
	Resume() Promise[struct{}]
	Suspend() Promise[struct{}]
}

func NewAudioContext() AudioContext {
	jsValue := js.Global().Get("AudioContext").New()
	return goAudioContext{
		goBaseAudioContext: goBaseAudioContext{
			goObject: goObject{
				jsValue: jsValue,
			},
		},
	}
}

var _ BaseAudioContext = goBaseAudioContext{}

type goBaseAudioContext struct {
	goObject
}

func (g goBaseAudioContext) CurrentTime() float64 {
	return g.jsValue.Get("currentTime").Float()
}

func (g goBaseAudioContext) Destination() AudioDestinationNode {
	jsValue := g.jsValue.Get("destination")
	return goAudioDestinationNode{
		goAudioNode: goAudioNode{
			goObject: goObject{
				jsValue: jsValue,
			},
		},
	}
}

func (g goBaseAudioContext) Listener() AudioListener {
	panic("TODO")
}

func (g goBaseAudioContext) SampleRate() float64 {
	return g.jsValue.Get("sampleRate").Float()
}

func (g goBaseAudioContext) State() AudioContextState {
	return AudioContextState(g.jsValue.Get("state").String())
}

func (g goBaseAudioContext) CreateBuffer(numChannels, length, sampleRate uint) AudioBuffer {
	panic("TODO")
}

func (g goBaseAudioContext) CreateBufferSource() AudioBufferSourceNode {
	panic("TODO")
}

func (g goBaseAudioContext) CreateConvolver() ConvolverNode {
	panic("TODO")
}

func (g goBaseAudioContext) CreateDelay() DelayNode {
	panic("TODO")
}

func (g goBaseAudioContext) CreateDynamicsCompressor() DynamicsCompressorNode {
	panic("TODO")
}

func (g goBaseAudioContext) CreateGain() GainNode {
	panic("TODO")
}

func (g goBaseAudioContext) CreateOscillator() OscillatorNode {
	jsValue := g.jsValue.Call("createOscillator")
	return goOscillatorNode{
		goAudioScheduledSourceNode: goAudioScheduledSourceNode{
			goAudioNode: goAudioNode{
				goObject: goObject{
					jsValue: jsValue,
				},
			},
		},
	}
}

func (g goBaseAudioContext) CreatePanner() PannerNode {
	panic("TODO")
}

func (g goBaseAudioContext) CreateStereoPanner() StereoPannerNode {
	panic("TODO")
}

func (g goBaseAudioContext) DecodeAudioData(data []byte) Promise[AudioBuffer] {
	panic("TODO")
}

var _ AudioContext = goAudioContext{}

type goAudioContext struct {
	goBaseAudioContext
}

func (g goAudioContext) BaseLatency() float64 {
	return g.jsValue.Get("baseLatency").Float()
}

func (g goAudioContext) OutputLatency() float64 {
	return g.jsValue.Get("outputLatency").Float()
}

func (g goAudioContext) Close() Promise[struct{}] {
	panic("TODO")
}

func (g goAudioContext) Resume() Promise[struct{}] {
	panic("TODO")
}

func (g goAudioContext) Suspend() Promise[struct{}] {
	panic("TODO")
}
