package registration

import "syscall/js"

func (c JsCallBack) RegistrationFunction() {
	js.Global().Set("update", js.FuncOf(c.Update))

}

func (c JsCallBack) Update(this js.Value, args []js.Value) interface{} {
	c.canvas.Clear()
	c.particles.MoveAll()
	c.canvas.DrawAll()
	js.Global().Call("requestAnimationFrame", c.Update)
	return nil
}
