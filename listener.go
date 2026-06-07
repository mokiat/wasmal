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
}

var _ AudioListener = (*goAudioListener)(nil)

type goAudioListener struct {
	goObject
}

func (g *goAudioListener) PositionX() AudioParam {
	jsValue := g.jsValue.Get("positionX")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goAudioListener) PositionY() AudioParam {
	jsValue := g.jsValue.Get("positionY")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goAudioListener) PositionZ() AudioParam {
	jsValue := g.jsValue.Get("positionZ")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goAudioListener) ForwardX() AudioParam {
	jsValue := g.jsValue.Get("forwardX")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goAudioListener) ForwardY() AudioParam {
	jsValue := g.jsValue.Get("forwardY")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goAudioListener) ForwardZ() AudioParam {
	jsValue := g.jsValue.Get("forwardZ")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goAudioListener) UpX() AudioParam {
	jsValue := g.jsValue.Get("upX")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goAudioListener) UpY() AudioParam {
	jsValue := g.jsValue.Get("upY")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goAudioListener) UpZ() AudioParam {
	jsValue := g.jsValue.Get("upZ")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}
