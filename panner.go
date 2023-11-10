package wasmal

// PannerNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/PannerNode
type PannerNode interface {
	AudioNode

	ConeInnerAngle() float64
	SetConeInnerAngle(angle float64)
	ConeOuterAngle() float64
	SetConeOuterAngle(angle float64)
	DistanceModel() PannerDistanceModel
	SetDistanceModel(model PannerDistanceModel)
	MaxDistance() float64
	SetMaxDistance(distance float64)
	OrientationX() AudioParam
	OrientationY() AudioParam
	OrientationZ() AudioParam
	PanningModel() PannerPanningModel
	SetPanningModel(model PannerPanningModel)
	PositionX() AudioParam
	PositionY() AudioParam
	PositionZ() AudioParam
	RefDistance() float64
	SetRefDistance(distance float64)
	RolloffFactor() float64
	SetRolloffFactor(factor float64)
}

type PannerDistanceModel string

const (
	PannerDistanceModelLinear      PannerDistanceModel = "linear"
	PannerDistanceModelInverse     PannerDistanceModel = "inverse"
	PannerDistanceModelExponential PannerDistanceModel = "exponential"
)

type PannerPanningModel string

const (
	PannerPanningModelEqualPower PannerPanningModel = "equalpower"
	PannerPanningModelHRTF       PannerPanningModel = "HRTF"
)

// StereoPannerNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/StereoPannerNode
type StereoPannerNode interface {
	AudioNode
	Pan() AudioParam
}

var _ PannerNode = goPannerNode{}

type goPannerNode struct {
	goAudioNode
}

func (g goPannerNode) ConeInnerAngle() float64 {
	return g.jsValue.Get("coneInnerAngle").Float()
}

func (g goPannerNode) SetConeInnerAngle(angle float64) {
	g.jsValue.Set("coneInnerAngle", angle)
}

func (g goPannerNode) ConeOuterAngle() float64 {
	return g.jsValue.Get("coneOuterAngle").Float()
}

func (g goPannerNode) SetConeOuterAngle(angle float64) {
	g.jsValue.Set("coneOuterAngle", angle)
}

func (g goPannerNode) DistanceModel() PannerDistanceModel {
	return PannerDistanceModel(g.jsValue.Get("distanceModel").String())
}

func (g goPannerNode) SetDistanceModel(model PannerDistanceModel) {
	g.jsValue.Set("distanceModel", string(model))
}

func (g goPannerNode) MaxDistance() float64 {
	return g.jsValue.Get("maxDistance").Float()
}

func (g goPannerNode) SetMaxDistance(distance float64) {
	g.jsValue.Set("maxDistance", distance)
}

func (g goPannerNode) OrientationX() AudioParam {
	jsValue := g.jsValue.Get("orientationX")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g goPannerNode) OrientationY() AudioParam {
	jsValue := g.jsValue.Get("orientationY")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g goPannerNode) OrientationZ() AudioParam {
	jsValue := g.jsValue.Get("orientationZ")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g goPannerNode) PanningModel() PannerPanningModel {
	return PannerPanningModel(g.jsValue.Get("panningModel").String())
}

func (g goPannerNode) SetPanningModel(model PannerPanningModel) {
	g.jsValue.Set("panningModel", string(model))
}

func (g goPannerNode) PositionX() AudioParam {
	jsValue := g.jsValue.Get("positionX")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g goPannerNode) PositionY() AudioParam {
	jsValue := g.jsValue.Get("positionY")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g goPannerNode) PositionZ() AudioParam {
	jsValue := g.jsValue.Get("positionZ")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g goPannerNode) RefDistance() float64 {
	return g.jsValue.Get("refDistance").Float()
}

func (g goPannerNode) SetRefDistance(distance float64) {
	g.jsValue.Set("refDistance", distance)
}

func (g goPannerNode) RolloffFactor() float64 {
	return g.jsValue.Get("rolloffFactor").Float()
}

func (g goPannerNode) SetRolloffFactor(factor float64) {
	g.jsValue.Set("rolloffFactor", factor)
}

var _ StereoPannerNode = goStereoPannerNode{}

type goStereoPannerNode struct {
	goAudioNode
}

func (g goStereoPannerNode) Pan() AudioParam {
	jsValue := g.jsValue.Get("pan")
	return goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}
