package wasmal

import "github.com/mokiat/gog"

// AudioParam as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/AudioParam
type AudioParam interface {
	object

	DefaultValue() float64
	MaxValue() float64
	MinValue() float64
	Value() float64
	SetValue(value float64)

	SetValueAtTime(value, startTime float64)
	LinearRampToValueAtTime(value, endTime float64)
	ExponentialRampToValueAtTime(value, endTime float64)
	SetTargetAtTime(target, startTime, timeConstant float64)
	SetValueCurveAtTime(values []float64, startTime, duration float64)
	CancelScheduledValues(startTime float64)
	CancelAndHoldAtTime(cancelTime float64)
}

type goAudioParam struct {
	goObject
}

func (g goAudioParam) DefaultValue() float64 {
	return g.jsValue.Get("defaultValue").Float()
}

func (g goAudioParam) MaxValue() float64 {
	return g.jsValue.Get("maxValue").Float()
}

func (g goAudioParam) MinValue() float64 {
	return g.jsValue.Get("minValue").Float()
}

func (g goAudioParam) Value() float64 {
	return g.jsValue.Get("value").Float()
}

func (g goAudioParam) SetValue(value float64) {
	g.jsValue.Set("value", value)
}

func (g goAudioParam) SetValueAtTime(value, startTime float64) {
	g.jsValue.Call("setValueAtTime", value, startTime)
}

func (g goAudioParam) LinearRampToValueAtTime(value, endTime float64) {
	g.jsValue.Call("linearRampToValueAtTime", value, endTime)
}

func (g goAudioParam) ExponentialRampToValueAtTime(value, endTime float64) {
	g.jsValue.Call("exponentialRampToValueAtTime", value, endTime)
}

func (g goAudioParam) SetTargetAtTime(target, startTime, timeConstant float64) {
	g.jsValue.Call("setTargetAtTime", target, startTime, timeConstant)
}

func (g goAudioParam) SetValueCurveAtTime(values []float64, startTime, duration float64) {
	anyValues := gog.Map(values, func(v float64) any { return v })
	g.jsValue.Call("setValueCurveAtTime", anyValues, startTime, duration)
}

func (g goAudioParam) CancelScheduledValues(startTime float64) {
	g.jsValue.Call("cancelScheduledValues", startTime)
}

func (g goAudioParam) CancelAndHoldAtTime(cancelTime float64) {
	g.jsValue.Call("cancelAndHoldAtTime", cancelTime)
}
