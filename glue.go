package wasmal

type Promise[T any] interface {
	Then(cb func(value T))
	Catch(cb func(err error))
}

type Float32Array interface {
}
