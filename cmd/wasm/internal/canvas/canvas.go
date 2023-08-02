package canvas

import (
	"particleLife/cmd/wasm/pkg/particle"
	"sync"
)

func (can Canvas) DrawAll(particles [][]particle.Particle) {
	var wg sync.WaitGroup

	for _, group := range particles {
		wg.Add(1)

		go func(group []particle.Particle) {
			defer wg.Done()

			for _, value := range group {
				can.context.Set("fillStyle", value.Color)
				can.context.Call("fillRect", value.X, value.Y, 5, 5)
			}
		}(group)
	}

	wg.Wait()
}

func (can Canvas) Clear() {
	can.context.Call("clearRect", 0, 0, can.width, can.height)
}
