package particles

import "particleLife/cmd/wasm/pkg/particle"

type ParticleGroup struct {
	particles           [][]particle.Particle
	maxHeight, maxWidth float64
	x, y                float64
	rule                map[string]float64
}

func NewParticleGroup(maxHeight, maxWidth float64) ParticleGroup {
	return ParticleGroup{
		maxHeight: maxHeight,
		maxWidth:  maxWidth,
		rule:      make(map[string]float64),
	}
}
