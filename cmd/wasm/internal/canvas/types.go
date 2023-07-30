package canvas

import "syscall/js"

type Canvas struct {
	context       js.Value
	width, height float64
}

func NewCanvas(ctx js.Value) Canvas {
	return Canvas{
		context: ctx,
	}
}
