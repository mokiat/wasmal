# wasmal

Go WASM bindings for the [W3C Web Audio API 1.1](https://www.w3.org/TR/webaudio-1.1/).

The library is a thin, direct mapping over JavaScript via `syscall/js`. All audio processing takes place in the browser's audio engine — no audio work happens in Go itself.

> **Warning:** The project is in early development and the API may change in backward-incompatible ways.

## Requirements

- Compiled with `GOOS=js GOARCH=wasm`
- A browser environment that supports the Web Audio API

If you are new to Go and WebAssembly, see the [official WebAssembly with Go documentation](https://github.com/golang/go/wiki/WebAssembly).

## Installation

```sh
go get github.com/mokiat/wasmal@latest
```

## Usage

Create an `AudioContext`, build a node graph, and connect it to the destination:

```go
ctx := wasmal.NewAudioContext()

osc := ctx.CreateOscillator()
osc.Frequency().SetValue(440) // A4

gain := ctx.CreateGain()
gain.Gain().SetValue(0.5)

osc.Connect(gain)
gain.Connect(ctx.Destination())
osc.Start(0)
```

Asynchronous operations return a `Promise[T]`. Register callbacks and call the returned `CleanupFunc` when they are no longer needed to release the underlying JavaScript function object:

```go
cleanup := ctx.Resume().Then(func(_ struct{}) {
    // audio is now playing
})
defer cleanup()
```

## Supported API

| Interface | Factory |
|-----------|---------|
| `AudioContext` / `BaseAudioContext` | `NewAudioContext()` |
| `AudioDestinationNode` | `ctx.Destination()` |
| `AudioBufferSourceNode` | `ctx.CreateBufferSource()` |
| `OscillatorNode` | `ctx.CreateOscillator()` |
| `GainNode` | `ctx.CreateGain()` |
| `BiquadFilterNode` | `ctx.CreateBiquadFilter()` |
| `ConvolverNode` | `ctx.CreateConvolver()` |
| `DelayNode` | `ctx.CreateDelay(maxDelay)` |
| `DynamicsCompressorNode` | `ctx.CreateDynamicsCompressor()` |
| `PannerNode` | `ctx.CreatePanner()` |
| `StereoPannerNode` | `ctx.CreateStereoPanner()` |
| `AudioBuffer` | `ctx.CreateBuffer(channels, length, sampleRate)` |
| `AudioParam` | exposed as fields on each node |
| `AudioListener` | `ctx.Listener()` |
