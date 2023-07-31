package registration

import (
	"strconv"
	"syscall/js"
)

func (c *JsCallBack) RegistrationFunction() {
	js.Global().Set("update", js.FuncOf(c.Update))
	js.Global().Set("setup", js.FuncOf(c.Setup))
}

func (c *JsCallBack) Setup(this js.Value, args []js.Value) interface{} {
	doc := js.Global().Get("document")

	if err := c.setCounters(doc); err != nil {
		return err
	}

	if err := c.setForce(doc); err != nil {
		return err
	}

	c.isSetup = true
	return nil
}

func (c *JsCallBack) setCounters(doc js.Value) error {
	countRed, err := strconv.Atoi(doc.Call("getElementById", "countRed").Get("value").String())
	if err != nil {
		return err
	}

	countBlue, err := strconv.Atoi(doc.Call("getElementById", "countBlue").Get("value").String())
	if err != nil {
		return err
	}

	countBlack, err := strconv.Atoi(doc.Call("getElementById", "countBlack").Get("value").String())
	if err != nil {
		return err
	}

	c.particles.GenerateGroups(countBlue, countRed, countBlack)
	return nil
}

func (c *JsCallBack) setForce(doc js.Value) error {
	colors := [3]string{"red", "blue", "black"}

	for _, firstColor := range colors {
		for _, secondColor := range colors {
			coupleName := firstColor + secondColor
			value, err := strconv.ParseFloat(doc.Call("getElementById", coupleName).Get("value").String(), 64)
			if err != nil {
				return err
			}

			c.particles.SetRules(coupleName, value)
		}
	}

	return nil
}

func (c *JsCallBack) Update(this js.Value, args []js.Value) interface{} {
	if !c.isSetup {
		return nil
	}

	var d func(this js.Value, args []js.Value) interface{}

	d = func(this js.Value, args []js.Value) interface{} {
		c.canvas.Clear()
		c.particles.MoveAll()
		c.canvas.DrawAll(c.particles.GetAllParticle())

		js.Global().Call("requestAnimationFrame", js.FuncOf(d))
		return nil
	}

	js.Global().Call("requestAnimationFrame", js.FuncOf(d))

	return nil
}
