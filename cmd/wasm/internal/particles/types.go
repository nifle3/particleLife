package particles

import "particleLife/cmd/wasm/pkg/particle"

type ParticleGroup struct {
	particles           [][]particle.Particle
	maxHeight, maxWidth float64
	x, y                float64
}

func NewParticleGroup(maxHeight, maxWidth, x, y float64) ParticleGroup {
	return ParticleGroup{
		maxHeight: maxHeight,
		maxWidth:  maxWidth,
		x:         x,
		y:         y,
	}
}
