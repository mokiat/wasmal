package wasmal

import "syscall/js"

type object interface {
	ref() js.Value
}

type goObject struct {
	jsValue js.Value
}

func (g goObject) ref() js.Value {
	return g.jsValue
}
