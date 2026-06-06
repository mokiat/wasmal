package wasmal

// CleanupFunc is a function that must be called to release resources.
// This is specific to the WebAssembly adapter and is not part of the official API.
type CleanupFunc func()
