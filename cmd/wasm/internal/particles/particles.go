package particles

import (
	"math"
	"math/rand"
	"particleLife/cmd/wasm/pkg/particle"
)

func (group *ParticleGroup) GenerateGroups(countBlue, countRed, countBlack int) {
	group.particles = [][]particle.Particle{
		group.generateArray("blue", countBlue),
		group.generateArray("red", countRed),
		//group.generateArray("black", countBlack),
	}
}

func (group *ParticleGroup) generateArray(color string, count int) []particle.Particle {
	particles := make([]particle.Particle, 0, count)

	for i := 0; i < count; i++ {
		x := 10 + rand.Float64()*(group.maxWidth-20.0)
		y := 10 + rand.Float64()*(group.maxHeight-20.0)
		particles = append(particles, particle.NewParticle(color, x, y))
	}

	return particles
}

func (group *ParticleGroup) MoveAll() {
	for i := 0; i < len(group.particles); i++ {
		for j := 0; j < len(group.particles); j++ {
			group.moveTwoGroup(group.particles[i], group.particles[j], 0.01)
		}
	}
}

func (group *ParticleGroup) moveTwoGroup(first, second []particle.Particle, g float64) {
	for i := 0; i < len(first); i++ {
		fx := 0.0
		fy := 0.0

		for j := 0; j < len(second); j++ {
			dx := first[i].X - second[j].X
			dy := first[i].Y - second[j].Y
			d := math.Sqrt(dx*dx + dy*dy)

			if d != 0 {
				f := g / d
				fx += f * dx
				fy += f * dy
			}
			first[i].Vx += fx
			first[i].Vy += fy

			first[i].X -= first[i].Vx
			first[i].Y -= first[i].Vy

			if first[i].X < 0 || first[i].X > group.maxWidth {
				first[i].Vx *= -1
			}

			if first[i].Y < 0 || first[i].Y > group.maxHeight {
				first[i].Vy *= -1
			}
		}

	}
}

func (group *ParticleGroup) GetAllParticle() [][]particle.Particle {
	return group.particles
}
