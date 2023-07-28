package Canvas

import "syscall/js"

type Canvas struct {
	canvas  js.Value
	context js.Value
}

func NewCanvas() *Canvas {
	canvas := js.Global().Get("document").Call("getElementById", "myCanvas")
	ctx := canvas.Call("getContext", "2d")
	return &Canvas{
		canvas:  canvas,
		context: ctx,
	}
}
