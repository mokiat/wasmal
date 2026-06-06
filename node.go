package wasmal

// AudioNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/AudioNode
type AudioNode interface {
	object

	// Context returns the AudioContext which owns this node.
	Context() BaseAudioContext

	// NumberOfInputs returns the number of inputs feeding into this node.
	NumberOfInputs() uint32

	// NumberOfOutputs returns the number of outputs leading out of this node.
	NumberOfOutputs() uint32

	// ConnectToNode connects this node to another AudioNode.
	ConnectToNode(destination AudioNode)

	// ConnectToParam connects this node to an AudioParam.
	ConnectToParam(destination AudioParam)

	// Disconnect disconnects this node from all outputs.
	Disconnect()

	// DisconnectFromNode disconnects this node from the specified destination node.
	DisconnectFromNode(destination AudioNode)

	// DisconnectFromParam disconnects this node from the specified destination param.
	DisconnectFromParam(destination AudioParam)
}

var _ AudioNode = (*goAudioNode)(nil)

type goAudioNode struct {
	goObject
}

func (g *goAudioNode) Context() BaseAudioContext {
	jsValue := g.jsValue.Get("context")
	return &goBaseAudioContext{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goAudioNode) NumberOfInputs() uint32 {
	return uint32(g.jsValue.Get("numberOfInputs").Int())
}

func (g *goAudioNode) NumberOfOutputs() uint32 {
	return uint32(g.jsValue.Get("numberOfOutputs").Int())
}

func (g *goAudioNode) ConnectToNode(destination AudioNode) {
	g.jsValue.Call("connect", destination.ref())
}

func (g *goAudioNode) ConnectToParam(destination AudioParam) {
	g.jsValue.Call("connect", destination.ref())
}

func (g *goAudioNode) Disconnect() {
	g.jsValue.Call("disconnect")
}

func (g *goAudioNode) DisconnectFromNode(destination AudioNode) {
	g.jsValue.Call("disconnect", destination.ref())
}

func (g *goAudioNode) DisconnectFromParam(destination AudioParam) {
	g.jsValue.Call("disconnect", destination.ref())
}
