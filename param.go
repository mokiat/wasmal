package wasmal

// AudioParam as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/AudioParam
type AudioParam interface {
	object

	// Value is the current value of the AudioParam.
	Value() float32

	// SetValue sets the value of the AudioParam.
	SetValue(value float32)

	// DefaultValue is the default value of the AudioParam.
	DefaultValue() float32

	// MinValue is the minimum value of the AudioParam.
	MinValue() float32

	// MaxValue is the maximum value of the AudioParam.
	MaxValue() float32

	// SetValueAtTime schedules a parameter change at a specific time.
	SetValueAtTime(value float32, startTime float64)

	// LinearRampToValueAtTime schedules a parameter change to a value at a
	// specific time, following a linear ramp.
	LinearRampToValueAtTime(value float32, endTime float64)

	// ExponentialRampToValueAtTime schedules a parameter change to a value at a
	// specific time, following an exponential ramp.
	ExponentialRampToValueAtTime(value float32, endTime float64)

	// SetTargetAtTime schedules a parameter change to a target value at a specific
	// time, following an exponential approach.
	SetTargetAtTime(target float32, startTime float64, timeConstant float32)

	// CancelScheduledValues cancels all scheduled parameter changes with a start time
	// greater than or equal to the given time.
	CancelScheduledValues(cancelTime float64)

	// CancelAndHoldAtTime cancels all scheduled parameter changes with a start
	// time greater than or equal to the given time, and holds the parameter value
	// at the value it had at the given time until the next scheduled change.
	CancelAndHoldAtTime(cancelTime float64)
}

var _ AudioParam = (*goAudioParam)(nil)

type goAudioParam struct {
	goObject
}

func (g *goAudioParam) Value() float32 {
	return float32(g.jsValue.Get("value").Float())
}

func (g *goAudioParam) SetValue(value float32) {
	g.jsValue.Set("value", value)
}

func (g *goAudioParam) DefaultValue() float32 {
	return float32(g.jsValue.Get("defaultValue").Float())
}

func (g *goAudioParam) MinValue() float32 {
	return float32(g.jsValue.Get("minValue").Float())
}

func (g *goAudioParam) MaxValue() float32 {
	return float32(g.jsValue.Get("maxValue").Float())
}

func (g *goAudioParam) SetValueAtTime(value float32, startTime float64) {
	g.jsValue.Call("setValueAtTime", value, startTime)
}

func (g *goAudioParam) LinearRampToValueAtTime(value float32, endTime float64) {
	g.jsValue.Call("linearRampToValueAtTime", value, endTime)
}

func (g *goAudioParam) ExponentialRampToValueAtTime(value float32, endTime float64) {
	g.jsValue.Call("exponentialRampToValueAtTime", value, endTime)
}

func (g *goAudioParam) SetTargetAtTime(target float32, startTime float64, timeConstant float32) {
	g.jsValue.Call("setTargetAtTime", target, startTime, timeConstant)
}

func (g *goAudioParam) CancelScheduledValues(startTime float64) {
	g.jsValue.Call("cancelScheduledValues", startTime)
}

func (g *goAudioParam) CancelAndHoldAtTime(cancelTime float64) {
	g.jsValue.Call("cancelAndHoldAtTime", cancelTime)
}
