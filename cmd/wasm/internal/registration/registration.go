package registration

import "syscall/js"

func (c JsCallBack) RegistrationFunction() {
	c.particles.GenerateGroups(1, 1, 2)
	js.Global().Set("update", js.FuncOf(c.Update))

}

func (c JsCallBack) Update(this js.Value, args []js.Value) interface{} {
	c.canvas.Clear()
	c.particles.MoveAll()
	c.canvas.DrawAll(c.particles.GetAllParticle())
	js.Global().Call("requestAnimationFrame", js.FuncOf(c.Update))
	return nil
}
