package canvas

import "syscall/js"

type Canvas struct {
	context       js.Value
	width, height float64
}

func NewCanvas(ctx js.Value, width, height float64) Canvas {
	return Canvas{
		context: ctx,
		width:   width,
		height:  height,
	}
}
