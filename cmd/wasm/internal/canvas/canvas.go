package canvas

import "particleLife/cmd/wasm/pkg/particle"

func (can Canvas) DrawAll(particles []particle.Particle) {
	for _, value := range particles {
		can.context.Set("fillStyle", value.Color)
		can.context.Call("fillRect", value.X, value.Y, 10, 10)
	}
}

func (can Canvas) Clear() {
	can.context.Call("clearRect", 0, 0, can.width, can.height)
}
