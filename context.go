package wasmal

import "syscall/js"

const DefaultSampleRate = 44100

// BaseAudioContext as described here:
// https://www.w3.org/TR/webaudio-1.1/#BaseAudioContext
type BaseAudioContext interface {
	object

	// Destination is a read-only property that returns an AudioDestinationNode
	// representing the final destination of all audio in the context.
	Destination() AudioDestinationNode

	// SampleRate returns the sample rate (in samples per second) at which the
	// AudioContext handles audio.
	SampleRate() float32

	// CurrentTime returns the current time of the audio context in seconds.
	CurrentTime() float64

	// Listener returns an AudioListener object which can be used to control the
	// position and orientation of the listener in 3D space.
	Listener() AudioListener

	// State returns the current state of the audio context.
	State() AudioContextState

	// CreateBuffer creates an empty AudioBuffer with the specified number of
	// channels, length in sample-frames, and sample rate.
	CreateBuffer(numChannels, length uint32, sampleRate float32) AudioBuffer

	// CreateBufferSource creates an AudioBufferSourceNode, which can be used to
	// play audio data contained within an AudioBuffer.
	CreateBufferSource() AudioBufferSourceNode

	// CreateConvolver creates a ConvolverNode, which can be used to apply a
	// convolution effect given an impulse response.
	CreateConvolver() ConvolverNode

	// CreateDelay creates a DelayNode, which can be used to delay the incoming
	// audio signal by a certain amount of time.
	CreateDelay(maxDelayTime float64) DelayNode

	// CreateDynamicsCompressor creates a DynamicsCompressorNode, which can be used
	// to apply dynamic range compression to the audio signal.
	CreateDynamicsCompressor() DynamicsCompressorNode

	// CreateGain creates a GainNode, which can be used to control the volume of the
	// audio signal.
	CreateGain() GainNode

	// CreateOscillator creates an OscillatorNode, which can be used to generate
	// periodic waveforms such as sine, square, sawtooth, and triangle waves.
	CreateOscillator() OscillatorNode

	// CreatePanner creates a PannerNode, which can be used to spatialize the
	// audio signal in 3D space.
	CreatePanner() PannerNode

	// CreateStereoPanner creates a StereoPannerNode, which can be used to pan the
	// audio signal left or right in the stereo field.
	CreateStereoPanner() StereoPannerNode

	// DecodeAudioData takes raw data representing an audio file and decodes it
	// asynchronously, returning a Promise that resolves to an AudioBuffer
	// containing the decoded audio data.
	DecodeAudioData(data []byte) Promise[AudioBuffer]
}

// AudioContextState as described here:
// https://www.w3.org/TR/webaudio-1.1/#enumdef-audiocontextstate
type AudioContextState string

const (
	AudioContextStateSuspended AudioContextState = "suspended"
	AudioContextStateRunning   AudioContextState = "running"
	AudioContextStateClosed    AudioContextState = "closed"
)

// AudioContext as described here:
// https://www.w3.org/TR/webaudio-1.1/#AudioContext
type AudioContext interface {
	BaseAudioContext

	// BaseLatency returns the base latency of the audio context in seconds.
	// This is the amount of time that elapses between the moment an audio signal
	// is generated and the moment it is heard at the destination.
	BaseLatency() float64

	// OutputLatency returns the output latency of the audio hardware in seconds.
	OutputLatency() float64

	// Close closes the audio context, releasing any system audio resources that
	// it uses.
	Close() Promise[struct{}]

	// Resume resumes the progression of time in the audio context, allowing audio
	// to play if it was previously suspended.
	Resume() Promise[struct{}]

	// Suspend suspends the progression of time in the audio context, pausing
	// audio playback and processing until Resume() is called.
	Suspend() Promise[struct{}]
}

var _ BaseAudioContext = (*goBaseAudioContext)(nil)

type goBaseAudioContext struct {
	goObject
}

func (g *goBaseAudioContext) Destination() AudioDestinationNode {
	jsValue := g.jsValue.Get("destination")
	return &goAudioDestinationNode{
		goAudioNode: goAudioNode{
			goObject: goObject{
				jsValue: jsValue,
			},
		},
	}
}

func (g *goBaseAudioContext) SampleRate() float32 {
	return float32(g.jsValue.Get("sampleRate").Float())
}

func (g *goBaseAudioContext) CurrentTime() float64 {
	return g.jsValue.Get("currentTime").Float()
}

func (g *goBaseAudioContext) Listener() AudioListener {
	jsValue := g.jsValue.Get("listener")
	return &goAudioListener{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goBaseAudioContext) State() AudioContextState {
	return AudioContextState(g.jsValue.Get("state").String())
}

func (g *goBaseAudioContext) CreateBuffer(numChannels, length uint32, sampleRate float32) AudioBuffer {
	jsValue := g.jsValue.Call("createBuffer", numChannels, length, sampleRate)
	return &goAudioBuffer{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goBaseAudioContext) CreateBufferSource() AudioBufferSourceNode {
	jsValue := g.jsValue.Call("createBufferSource")
	return &goAudioBufferSourceNode{
		goAudioScheduledSourceNode: goAudioScheduledSourceNode{
			goAudioNode: goAudioNode{
				goObject: goObject{
					jsValue: jsValue,
				},
			},
		},
	}
}

func (g *goBaseAudioContext) CreateConvolver() ConvolverNode {
	jsValue := g.jsValue.Call("createConvolver")
	return &goConvolverNode{
		goAudioNode: goAudioNode{
			goObject: goObject{
				jsValue: jsValue,
			},
		},
	}
}

func (g *goBaseAudioContext) CreateDelay(maxDelayTime float64) DelayNode {
	jsValue := g.jsValue.Call("createDelay", maxDelayTime)
	return &goDelayNode{
		goAudioNode: goAudioNode{
			goObject: goObject{
				jsValue: jsValue,
			},
		},
	}
}

func (g *goBaseAudioContext) CreateDynamicsCompressor() DynamicsCompressorNode {
	jsValue := g.jsValue.Call("createDynamicsCompressor")
	return &goDynamicsCompressorNode{
		goAudioNode: goAudioNode{
			goObject: goObject{
				jsValue: jsValue,
			},
		},
	}
}

func (g *goBaseAudioContext) CreateGain() GainNode {
	jsValue := g.jsValue.Call("createGain")
	return &goGainNode{
		goAudioNode: goAudioNode{
			goObject: goObject{
				jsValue: jsValue,
			},
		},
	}
}

func (g *goBaseAudioContext) CreateOscillator() OscillatorNode {
	jsValue := g.jsValue.Call("createOscillator")
	return &goOscillatorNode{
		goAudioScheduledSourceNode: goAudioScheduledSourceNode{
			goAudioNode: goAudioNode{
				goObject: goObject{
					jsValue: jsValue,
				},
			},
		},
	}
}

func (g *goBaseAudioContext) CreatePanner() PannerNode {
	jsValue := g.jsValue.Call("createPanner")
	return &goPannerNode{
		goAudioNode: goAudioNode{
			goObject: goObject{
				jsValue: jsValue,
			},
		},
	}
}

func (g *goBaseAudioContext) CreateStereoPanner() StereoPannerNode {
	jsValue := g.jsValue.Call("createStereoPanner")
	return &goStereoPannerNode{
		goAudioNode: goAudioNode{
			goObject: goObject{
				jsValue: jsValue,
			},
		},
	}
}

func (g *goBaseAudioContext) DecodeAudioData(data []byte) Promise[AudioBuffer] {
	arrayBuffer := js.Global().Get("ArrayBuffer").New(len(data))
	uint8Array := js.Global().Get("Uint8Array").New(arrayBuffer)
	js.CopyBytesToJS(uint8Array, data)

	jsPromise := g.jsValue.Call("decodeAudioData", arrayBuffer)
	return &goPromise[AudioBuffer]{
		goObject: goObject{
			jsValue: jsPromise,
		},
		convert: func(value js.Value) AudioBuffer {
			return &goAudioBuffer{
				goObject: goObject{
					jsValue: value,
				},
			}
		},
	}
}

var _ AudioContext = (*goAudioContext)(nil)

type goAudioContext struct {
	goBaseAudioContext
}

func NewAudioContext() AudioContext {
	jsValue := js.Global().Get("AudioContext").New()
	return &goAudioContext{
		goBaseAudioContext: goBaseAudioContext{
			goObject: goObject{
				jsValue: jsValue,
			},
		},
	}
}

func (g *goAudioContext) BaseLatency() float64 {
	return g.jsValue.Get("baseLatency").Float()
}

func (g *goAudioContext) OutputLatency() float64 {
	return g.jsValue.Get("outputLatency").Float()
}

func (g *goAudioContext) Close() Promise[struct{}] {
	jsPromise := g.jsValue.Call("close")
	return &goPromise[struct{}]{
		goObject: goObject{
			jsValue: jsPromise,
		},
		convert: func(value js.Value) struct{} {
			return struct{}{}
		},
	}
}

func (g *goAudioContext) Resume() Promise[struct{}] {
	jsPromise := g.jsValue.Call("resume")
	return &goPromise[struct{}]{
		goObject: goObject{
			jsValue: jsPromise,
		},
		convert: func(value js.Value) struct{} {
			return struct{}{}
		},
	}
}

func (g *goAudioContext) Suspend() Promise[struct{}] {
	jsPromise := g.jsValue.Call("suspend")
	return &goPromise[struct{}]{
		goObject: goObject{
			jsValue: jsPromise,
		},
		convert: func(value js.Value) struct{} {
			return struct{}{}
		},
	}
}
