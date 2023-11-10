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
	jsValue := g.jsValue.Call("createBufferSource")
	return goAudioBufferSourceNode{
		goAudioScheduledSourceNode: goAudioScheduledSourceNode{
			goAudioNode: goAudioNode{
				goObject: goObject{
					jsValue: jsValue,
				},
			},
		},
	}
}

func (g goBaseAudioContext) CreateConvolver() ConvolverNode {
	jsValue := g.jsValue.Call("createConvolver")
	return goConvolverNode{
		goAudioNode: goAudioNode{
			goObject: goObject{
				jsValue: jsValue,
			},
		},
	}
}

func (g goBaseAudioContext) CreateDelay() DelayNode {
	jsValue := g.jsValue.Call("createDelay")
	return goDelayNode{
		goAudioNode: goAudioNode{
			goObject: goObject{
				jsValue: jsValue,
			},
		},
	}
}

func (g goBaseAudioContext) CreateDynamicsCompressor() DynamicsCompressorNode {
	jsValue := g.jsValue.Call("createDynamicsCompressor")
	return goDynamicsCompressorNode{
		goAudioNode: goAudioNode{
			goObject: goObject{
				jsValue: jsValue,
			},
		},
	}
}

func (g goBaseAudioContext) CreateGain() GainNode {
	jsValue := g.jsValue.Call("createGain")
	return goGainNode{
		goAudioNode: goAudioNode{
			goObject: goObject{
				jsValue: jsValue,
			},
		},
	}
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
	jsValue := g.jsValue.Call("createPanner")
	return goPannerNode{
		goAudioNode: goAudioNode{
			goObject: goObject{
				jsValue: jsValue,
			},
		},
	}
}

func (g goBaseAudioContext) CreateStereoPanner() StereoPannerNode {
	jsValue := g.jsValue.Call("createStereoPanner")
	return goStereoPannerNode{
		goAudioNode: goAudioNode{
			goObject: goObject{
				jsValue: jsValue,
			},
		},
	}
}

func (g goBaseAudioContext) DecodeAudioData(data []byte) Promise[AudioBuffer] {
	arrayBuffer := js.Global().Get("ArrayBuffer").New(len(data))
	uint8Array := js.Global().Get("Uint8Array").New(arrayBuffer)
	js.CopyBytesToJS(uint8Array, data)

	jsPromise := g.jsValue.Call("decodeAudioData", arrayBuffer)
	return goPromise[AudioBuffer]{
		goObject: goObject{
			jsValue: jsPromise,
		},
		convert: func(value js.Value) AudioBuffer {
			return goAudioBuffer{
				goObject: goObject{
					jsValue: value,
				},
			}
		},
	}
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
