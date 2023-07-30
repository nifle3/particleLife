package particles

import "particleLife/cmd/wasm/pkg/particle"

type ParticleGroup struct {
	particles           [][]particle.Particle
	maxHeight, maxWidth float64
}
