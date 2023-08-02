package particles

import (
	"math"
	"math/rand"
	"particleLife/cmd/wasm/pkg/particle"
	"sync"
)

func (group *ParticleGroup) GenerateGroups(countBlue, countRed, countBlack int) {
	group.particles = [][]particle.Particle{
		group.generateArray("blue", countBlue),
		group.generateArray("red", countRed),
		group.generateArray("black", countBlack),
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
			group.moveTwoGroup(group.particles[i], group.particles[j])
		}
	}
}

func (group *ParticleGroup) moveTwoGroup(first, second []particle.Particle) {
	var firstWaitGroup sync.WaitGroup
	for i := 0; i < len(first); i++ {
		firstWaitGroup.Add(1)

		go func(i int) {
			defer firstWaitGroup.Done()

			fx := 0.0
			fy := 0.0

			for j := 0; j < len(second); j++ {

				dx := first[i].X - second[j].X
				dy := first[i].Y - second[j].Y
				d := math.Sqrt(dx*dx + dy*dy)

				if d != 0 {
					g := group.getRules(first[i].Color + second[j].Color)

					f := g / d
					fx += f * dx
					fy += f * dy
				}

				if first[i].Vx >= 80 || first[i].Vx <= -80 {
					first[i].Vx += fx
				} else {
					first[i].Vx -= fx

					first[i].X -= first[i].Vx
				}

				if first[i].Vy >= 80 || first[i].Vy <= -80 {
					first[i].Vy += fy
				} else {
					first[i].Vy -= fy

					first[i].Y -= first[i].Vy
				}

				if first[i].X < 0 || first[i].X > group.maxWidth {
					first[i].Vx *= -1
				}

				if first[i].Y < 0 || first[i].Y > group.maxHeight {
					first[i].Vy *= -1
				}
			}
		}(i)
	}

	firstWaitGroup.Wait()
}

func (group *ParticleGroup) GetAllParticle() [][]particle.Particle {
	return group.particles
}

func (group *ParticleGroup) SetRules(coupleName string, value float64) {
	group.rule[coupleName] = value
}

func (group *ParticleGroup) getRules(ruleName string) float64 {
	g := 1.0

	value, ok := group.rule[ruleName]
	if ok {
		g = value
	}

	return g
}
