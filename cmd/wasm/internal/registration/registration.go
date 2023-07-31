package registration

import (
	"strconv"
	"syscall/js"
)

func (c JsCallBack) RegistrationFunction() {
	js.Global().Set("update", js.FuncOf(c.Update))
	js.Global().Set("setup", js.FuncOf(c.Setup))
}

func (c JsCallBack) Setup(this js.Value, args []js.Value) interface{} {
	doc := js.Global().Get("document")

	countRed, err := strconv.Atoi(doc.Call("getElementById", "countRed").Get("value").String())
	countBlue, err := strconv.Atoi(doc.Call("getElementById", "countBlue").Get("value").String())
	countBlack, err := strconv.Atoi(doc.Call("getElementById", "countBlack").Get("value").String())

	if err != nil {
		return nil
	}

	c.particles.GenerateGroups(countBlue, countRed, countBlack)

	return nil
}

func (c JsCallBack) Update(this js.Value, args []js.Value) interface{} {
	c.canvas.Clear()
	c.particles.MoveAll()
	c.canvas.DrawAll(c.particles.GetAllParticle())

	js.Global().Call("requestAnimationFrame", js.FuncOf(c.Update))

	return nil
}
