package wasmal

// AudioNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/AudioNode
type AudioNode interface {
	object

	Context() BaseAudioContext
	NumberOfInputs() uint
	NumberOfOutputs() uint

	ConnectNode(destination AudioNode)
	ConnectParam(destination AudioParam)
	Disconnect()
	DisconnectNode(destination AudioNode)
	DisconnectParam(destination AudioParam)
}

var _ AudioNode = goAudioNode{}

type goAudioNode struct {
	goObject
}

func (g goAudioNode) Context() BaseAudioContext {
	jsValue := g.jsValue.Get("context")
	return goBaseAudioContext{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g goAudioNode) NumberOfInputs() uint {
	return uint(g.jsValue.Get("numberOfInputs").Int())
}

func (g goAudioNode) NumberOfOutputs() uint {
	return uint(g.jsValue.Get("numberOfOutputs").Int())
}

func (g goAudioNode) ConnectNode(destination AudioNode) {
	g.jsValue.Call("connect", destination.ref())
}

func (g goAudioNode) ConnectParam(destination AudioParam) {
	g.jsValue.Call("connect", destination.ref())
}

func (g goAudioNode) Disconnect() {
	g.jsValue.Call("disconnect")
}

func (g goAudioNode) DisconnectNode(destination AudioNode) {
	g.jsValue.Call("disconnect", destination.ref())
}

func (g goAudioNode) DisconnectParam(destination AudioParam) {
	g.jsValue.Call("disconnect", destination.ref())
}
