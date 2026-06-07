package wasmal

// PannerNode as described here:
// https://www.w3.org/TR/webaudio-1.1/#PannerNode
type PannerNode interface {
	AudioNode

	// PanningModel returns the current panning model used by the PannerNode.
	PanningModel() PanningModelType

	// SetPanningModel sets the panning model used by the PannerNode. The default
	// value is "equalpower".
	SetPanningModel(model PanningModelType)

	// PositionX returns an AudioParam representing the x-axis position of the
	// audio source in 3D space.
	PositionX() AudioParam

	// PositionY returns an AudioParam representing the y-axis position of the
	// audio source in 3D space.
	PositionY() AudioParam

	// PositionZ returns an AudioParam representing the z-axis position of the
	// audio source in 3D space.
	PositionZ() AudioParam

	// OrientationX returns an AudioParam representing the x-axis orientation of the
	// audio source in 3D space.
	OrientationX() AudioParam

	// OrientationY returns an AudioParam representing the y-axis orientation of the
	// audio source in 3D space.
	OrientationY() AudioParam

	// OrientationZ returns an AudioParam representing the z-axis orientation of the
	// audio source in 3D space.
	OrientationZ() AudioParam

	// DistanceModel returns the current distance model used by the PannerNode.
	DistanceModel() DistanceModelType

	// SetDistanceModel sets the distance model used by the PannerNode. The default
	// value is "inverse".
	SetDistanceModel(model DistanceModelType)

	// RefDistance returns the reference distance for reducing volume as the audio
	// source moves further from the listener. The default value is 1.
	RefDistance() float64

	// SetRefDistance sets the reference distance for reducing volume as the audio
	// source moves further from the listener. The default value is 1.
	SetRefDistance(distance float64)

	// MaxDistance returns the maximum distance between the audio source and the
	// listener, after which the volume will not be reduced any further.
	MaxDistance() float64

	// SetMaxDistance sets the maximum distance between the audio source and the
	// listener, after which the volume will not be reduced any further.
	SetMaxDistance(distance float64)

	// RolloffFactor returns the rolloff factor for reducing volume as the audio
	// source moves further from the listener. The default value is 1.
	RolloffFactor() float64

	// SetRolloffFactor sets the rolloff factor for reducing volume as the audio
	// source moves further from the listener. The default value is 1.
	SetRolloffFactor(factor float64)

	// ConeInnerAngle returns the inner angle of the cone in degrees.
	// The default value is 360.
	ConeInnerAngle() float64

	// SetConeInnerAngle sets the inner angle of the cone in degrees.
	// The default value is 360.
	SetConeInnerAngle(angle float64)

	// ConeOuterAngle returns the outer angle of the cone in degrees. The default
	// value is 360.
	ConeOuterAngle() float64

	// SetConeOuterAngle sets the outer angle of the cone in degrees. The default
	// value is 360.
	SetConeOuterAngle(angle float64)

	// ConeOuterGain returns the gain outside the outer angle of the cone. The
	// default value is 0.
	ConeOuterGain() float64

	// SetConeOuterGain sets the gain outside the outer angle of the cone. The
	// default value is 0.
	SetConeOuterGain(gain float64)
}

// DistanceModelType as described here:
// https://www.w3.org/TR/webaudio-1.1/#enumdef-distancemodeltype
type DistanceModelType string

const (
	DistanceModelTypeLinear      DistanceModelType = "linear"
	DistanceModelTypeInverse     DistanceModelType = "inverse"
	DistanceModelTypeExponential DistanceModelType = "exponential"
)

// PanningModelType as described here:
// https://www.w3.org/TR/webaudio-1.1/#enumdef-panningmodeltype
type PanningModelType string

const (
	PanningModelTypeEqualPower PanningModelType = "equalpower"
	PanningModelTypeHRTF       PanningModelType = "HRTF"
)

// StereoPannerNode as described here:
// https://www.w3.org/TR/webaudio-1.1/#stereopannernode
type StereoPannerNode interface {
	AudioNode

	// Pan returns an AudioParam representing the left-to-right pan control.
	Pan() AudioParam
}

var _ PannerNode = (*goPannerNode)(nil)

type goPannerNode struct {
	goAudioNode
}

func (g *goPannerNode) PanningModel() PanningModelType {
	return PanningModelType(g.jsValue.Get("panningModel").String())
}

func (g *goPannerNode) SetPanningModel(model PanningModelType) {
	g.jsValue.Set("panningModel", string(model))
}

func (g *goPannerNode) PositionX() AudioParam {
	jsValue := g.jsValue.Get("positionX")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goPannerNode) PositionY() AudioParam {
	jsValue := g.jsValue.Get("positionY")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goPannerNode) PositionZ() AudioParam {
	jsValue := g.jsValue.Get("positionZ")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goPannerNode) OrientationX() AudioParam {
	jsValue := g.jsValue.Get("orientationX")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goPannerNode) OrientationY() AudioParam {
	jsValue := g.jsValue.Get("orientationY")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goPannerNode) OrientationZ() AudioParam {
	jsValue := g.jsValue.Get("orientationZ")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}

func (g *goPannerNode) DistanceModel() DistanceModelType {
	return DistanceModelType(g.jsValue.Get("distanceModel").String())
}

func (g *goPannerNode) SetDistanceModel(model DistanceModelType) {
	g.jsValue.Set("distanceModel", string(model))
}

func (g *goPannerNode) RefDistance() float64 {
	return g.jsValue.Get("refDistance").Float()
}

func (g *goPannerNode) SetRefDistance(distance float64) {
	g.jsValue.Set("refDistance", distance)
}

func (g *goPannerNode) MaxDistance() float64 {
	return g.jsValue.Get("maxDistance").Float()
}

func (g *goPannerNode) SetMaxDistance(distance float64) {
	g.jsValue.Set("maxDistance", distance)
}

func (g *goPannerNode) RolloffFactor() float64 {
	return g.jsValue.Get("rolloffFactor").Float()
}

func (g *goPannerNode) SetRolloffFactor(factor float64) {
	g.jsValue.Set("rolloffFactor", factor)
}

func (g *goPannerNode) ConeInnerAngle() float64 {
	return g.jsValue.Get("coneInnerAngle").Float()
}

func (g *goPannerNode) SetConeInnerAngle(angle float64) {
	g.jsValue.Set("coneInnerAngle", angle)
}

func (g *goPannerNode) ConeOuterAngle() float64 {
	return g.jsValue.Get("coneOuterAngle").Float()
}

func (g *goPannerNode) SetConeOuterAngle(angle float64) {
	g.jsValue.Set("coneOuterAngle", angle)
}

func (g *goPannerNode) ConeOuterGain() float64 {
	return g.jsValue.Get("coneOuterGain").Float()
}

func (g *goPannerNode) SetConeOuterGain(gain float64) {
	g.jsValue.Set("coneOuterGain", gain)
}

var _ StereoPannerNode = (*goStereoPannerNode)(nil)

type goStereoPannerNode struct {
	goAudioNode
}

func (g *goStereoPannerNode) Pan() AudioParam {
	jsValue := g.jsValue.Get("pan")
	return &goAudioParam{
		goObject: goObject{
			jsValue: jsValue,
		},
	}
}
