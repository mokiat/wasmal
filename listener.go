package wasmal

// AudioListener as described here:
// https://www.w3.org/TR/webaudio-1.1/#AudioListener
type AudioListener interface {
	object

	// PositionX returns an AudioParam representing the x coordinate of the
	// position of the listener.
	PositionX() AudioParam

	// PositionY returns an AudioParam representing the y coordinate of the
	// position of the listener.
	PositionY() AudioParam

	// PositionZ returns an AudioParam representing the z coordinate of the
	// position of the listener.
	PositionZ() AudioParam

	// ForwardX returns an AudioParam representing the x coordinate of the
	// forward orientation of the listener.
	ForwardX() AudioParam

	// ForwardY returns an AudioParam representing the y coordinate of the
	// forward orientation of the listener.
	ForwardY() AudioParam

	// ForwardZ returns an AudioParam representing the z coordinate of the
	// forward orientation of the listener.
	ForwardZ() AudioParam

	// UpX returns an AudioParam representing the x coordinate of the up
	// orientation of the listener.
	UpX() AudioParam

	// UpY returns an AudioParam representing the y coordinate of the up
	// orientation of the listener.
	UpY() AudioParam

	// UpZ returns an AudioParam representing the z coordinate of the up
	// orientation of the listener.
	UpZ() AudioParam

	// SetPosition sets the position of the listener in 3D space. It is
	// equivalent to setting the values of PositionX, PositionY and PositionZ,
	// but requires a single call.
	SetPosition(x, y, z float32)

	// SetOrientation sets the forward and up directions of the listener in 3D
	// space. It is equivalent to setting the values of ForwardX, ForwardY,
	// ForwardZ, UpX, UpY and UpZ, but requires a single call.
	SetOrientation(x, y, z, xUp, yUp, zUp float32)
}

var _ AudioListener = (*goAudioListener)(nil)

type goAudioListener struct {
	goObject
}

func (g *goAudioListener) PositionX() AudioParam {
	jsValue := g.jsValue.Get("positionX")
	return &goAudioParam{
		jsValue: jsValue,
	}
}

func (g *goAudioListener) PositionY() AudioParam {
	jsValue := g.jsValue.Get("positionY")
	return &goAudioParam{
		jsValue: jsValue,
	}
}

func (g *goAudioListener) PositionZ() AudioParam {
	jsValue := g.jsValue.Get("positionZ")
	return &goAudioParam{
		jsValue: jsValue,
	}
}

func (g *goAudioListener) ForwardX() AudioParam {
	jsValue := g.jsValue.Get("forwardX")
	return &goAudioParam{
		jsValue: jsValue,
	}
}

func (g *goAudioListener) ForwardY() AudioParam {
	jsValue := g.jsValue.Get("forwardY")
	return &goAudioParam{
		jsValue: jsValue,
	}
}

func (g *goAudioListener) ForwardZ() AudioParam {
	jsValue := g.jsValue.Get("forwardZ")
	return &goAudioParam{
		jsValue: jsValue,
	}
}

func (g *goAudioListener) UpX() AudioParam {
	jsValue := g.jsValue.Get("upX")
	return &goAudioParam{
		jsValue: jsValue,
	}
}

func (g *goAudioListener) UpY() AudioParam {
	jsValue := g.jsValue.Get("upY")
	return &goAudioParam{
		jsValue: jsValue,
	}
}

func (g *goAudioListener) UpZ() AudioParam {
	jsValue := g.jsValue.Get("upZ")
	return &goAudioParam{
		jsValue: jsValue,
	}
}

func (g *goAudioListener) SetPosition(x, y, z float32) {
	g.jsValue.Call("setPosition", x, y, z)
}

func (g *goAudioListener) SetOrientation(x, y, z, xUp, yUp, zUp float32) {
	g.jsValue.Call("setOrientation", x, y, z, xUp, yUp, zUp)
}
