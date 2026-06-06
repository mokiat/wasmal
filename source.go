package wasmal

import "syscall/js"

// AudioScheduledSourceNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/AudioScheduledSourceNode
type AudioScheduledSourceNode interface {
	AudioNode

	// Start starts playback of the audio source at the specified time.
	//
	// If it is zero or before [AudioContext.CurrentTime], the audio will start immediately.
	Start(when float64)

	// Stop stops playback of the audio source at the specified time.
	//
	// If it is zero or before [AudioContext.CurrentTime], the audio will stop immediately.
	Stop(when float64)

	// SetOnEnded sets a callback that will be called when the audio source has finished playing.
	//
	// The returned CleanupFunc is not per spec but is needed due to the way Go's js.Func works.
	// It should be called when the callback is no longer needed to release resources.
	// It is important that the callback is not called after it has been released. You may want
	// to call SetOnEnded with a nil callback to ensure that the callback is not called after it
	// has been released.
	SetOnEnded(onEnded func()) CleanupFunc
}

// AudioBufferSourceNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/AudioBufferSourceNode
type AudioBufferSourceNode interface {
	AudioScheduledSourceNode

	// Buffer returns the AudioBuffer that is being used as the source of the audio data.
	Buffer() AudioBuffer

	// SetBuffer sets the AudioBuffer that is being used as the source of the audio data.
	SetBuffer(buffer AudioBuffer)

	// PlaybackRate returns an AudioParam that controls the rate at which the audio is played back.
	PlaybackRate() AudioParam

	// Detune returns an AudioParam that controls the detuning of the audio source in cents.
	Detune() AudioParam

	// Loop returns a boolean indicating whether the audio source should loop within
	// [AudioBufferSourceNode.LoopStart] and [AudioBufferSourceNode.LoopEnd].
	Loop() bool

	// SetLoop sets a boolean indicating whether the audio source should loop.
	SetLoop(loop bool)

	// LoopStart returns the time in seconds within the audio buffer at which the
	// loop should start.
	LoopStart() float64

	// SetLoopStart sets the time in seconds within the audio buffer at which the
	// loop should start.
	SetLoopStart(start float64)

	// LoopEnd returns the time in seconds within the audio buffer at which the
	// loop should end.
	LoopEnd() float64

	// SetLoopEnd sets the time in seconds within the audio buffer at which the
	// loop should end.
	SetLoopEnd(end float64)

	// StartOffset starts playback of the audio source at the specified time and
	// with the specified offset.
	StartOffset(when, offset float64)

	// StartOffsetDuration starts playback of the audio source at the specified
	// time, with the specified offset and duration.
	StartOffsetDuration(when, offset, duration float64)
}

var _ AudioScheduledSourceNode = (*goAudioScheduledSourceNode)(nil)

type goAudioScheduledSourceNode struct {
	goAudioNode

	onEnded js.Func
}

func (g *goAudioScheduledSourceNode) Start(when float64) {
	g.jsValue.Call("start", when)
}

func (g *goAudioScheduledSourceNode) Stop(when float64) {
	g.jsValue.Call("stop", when)
}

func (g *goAudioScheduledSourceNode) SetOnEnded(cb func()) CleanupFunc {
	onEnded := js.FuncOf(func(this js.Value, args []js.Value) any {
		cb()
		return nil
	})
	g.jsValue.Set("onended", onEnded)
	return onEnded.Release
}

var _ AudioBufferSourceNode = (*goAudioBufferSourceNode)(nil)

type goAudioBufferSourceNode struct {
	goAudioScheduledSourceNode
}

func (g *goAudioBufferSourceNode) Buffer() AudioBuffer {
	jsValue := g.jsValue.Get("buffer")
	return &goAudioBuffer{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goAudioBufferSourceNode) SetBuffer(buffer AudioBuffer) {
	g.jsValue.Set("buffer", buffer.ref())
}

func (g *goAudioBufferSourceNode) PlaybackRate() AudioParam {
	jsValue := g.jsValue.Get("playbackRate")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goAudioBufferSourceNode) Detune() AudioParam {
	jsValue := g.jsValue.Get("detune")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goAudioBufferSourceNode) Loop() bool {
	return g.jsValue.Get("loop").Bool()
}

func (g *goAudioBufferSourceNode) SetLoop(loop bool) {
	g.jsValue.Set("loop", loop)
}

func (g *goAudioBufferSourceNode) LoopStart() float64 {
	return g.jsValue.Get("loopStart").Float()
}

func (g *goAudioBufferSourceNode) SetLoopStart(start float64) {
	g.jsValue.Set("loopStart", start)
}

func (g *goAudioBufferSourceNode) LoopEnd() float64 {
	return g.jsValue.Get("loopEnd").Float()
}

func (g *goAudioBufferSourceNode) SetLoopEnd(end float64) {
	g.jsValue.Set("loopEnd", end)
}

func (g *goAudioBufferSourceNode) StartOffset(when, offset float64) {
	g.jsValue.Call("start", when, offset)
}

func (g *goAudioBufferSourceNode) StartOffsetDuration(when, offset, duration float64) {
	g.jsValue.Call("start", when, offset, duration)
}
