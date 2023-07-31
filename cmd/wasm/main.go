package main

import (
	"fmt"
	"particleLife/cmd/wasm/internal/canvas"
	"particleLife/cmd/wasm/internal/particles"
	"particleLife/cmd/wasm/internal/registration"
	"syscall/js"
)

func main() {
	c := make(chan bool)
	fmt.Println("Wasm init")

	canvasJs := js.Global().Get("document").Call("getElementById", "myCanvas")

	canvasRect := canvasJs.Call("getBoundingClientRect")
	x := canvasRect.Get("left").Float()
	y := canvasRect.Get("top").Float()

	maxHeight := canvasJs.Get("height").Float()
	maxWidth := canvasJs.Get("width").Float()
	particlesGroup := particles.NewParticleGroup(maxHeight, maxWidth, x, y)

	ctx := canvasJs.Call("getContext", "2d")
	canv := canvas.NewCanvas(ctx, maxWidth, maxHeight)

	jsCallBack := registration.NewJsCallBack(canv, &particlesGroup)
	jsCallBack.RegistrationFunction()

	<-c
}
