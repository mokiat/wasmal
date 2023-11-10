package wasmal

type AudioListener interface {
	object

	PositionX() AudioParam
	PositionY() AudioParam
	PositionZ() AudioParam
	ForwardX() AudioParam
	ForwardY() AudioParam
	ForwardZ() AudioParam
	UpX() AudioParam
	UpY() AudioParam
	UpZ() AudioParam
}

var _ AudioListener = goAudioListener{}

type goAudioListener struct {
	goObject
}

func (g goAudioListener) PositionX() AudioParam {
	jsValue := g.jsValue.Get("positionX")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g goAudioListener) PositionY() AudioParam {
	jsValue := g.jsValue.Get("positionY")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g goAudioListener) PositionZ() AudioParam {
	jsValue := g.jsValue.Get("positionZ")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g goAudioListener) ForwardX() AudioParam {
	jsValue := g.jsValue.Get("forwardX")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g goAudioListener) ForwardY() AudioParam {
	jsValue := g.jsValue.Get("forwardY")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g goAudioListener) ForwardZ() AudioParam {
	jsValue := g.jsValue.Get("forwardZ")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g goAudioListener) UpX() AudioParam {
	jsValue := g.jsValue.Get("upX")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g goAudioListener) UpY() AudioParam {
	jsValue := g.jsValue.Get("upY")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g goAudioListener) UpZ() AudioParam {
	jsValue := g.jsValue.Get("upZ")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}
