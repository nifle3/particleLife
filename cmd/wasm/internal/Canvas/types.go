package Canvas

import "syscall/js"

type Canvas struct {
	width   int
	height  int
	context js.Value
}

func NewCanvas() *Canvas {
	canvas := js.Global().Get("document").Call("getElementById", "myCanvas")
	ctx := canvas.Call("getContext", "2d")
	return &Canvas{
		width:   canvas.Get("width").Int(),
		height:  canvas.Get("height").Int(),
		context: ctx,
	}
}
