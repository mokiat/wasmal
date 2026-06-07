// Package wasmal provides Go WASM bindings for the W3C Web Audio API 1.1.
//
// It is a thin, direct mapping over JavaScript via syscall/js; all audio
// processing takes place in the browser's audio engine, not in Go.
// The package must be compiled with GOOS=js GOARCH=wasm.
package wasmal
