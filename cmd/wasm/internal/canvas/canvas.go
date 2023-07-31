package canvas

import "particleLife/cmd/wasm/pkg/particle"

func (can Canvas) DrawAll(particles [][]particle.Particle) {
	for _, group := range particles {
		for _, value := range group {
			can.context.Set("fillStyle", value.Color)
			can.context.Call("fillRect", value.X, value.Y, 5, 5)
		}
	}
}

func (can Canvas) Clear() {
	can.context.Call("clearRect", 0, 0, can.width, can.height)
}
