package particle

type Particle struct {
	Color  string
	Vx, Vy float64
	X, Y   float64
}

func NewParticle(color string, x, y float64) Particle {
	return Particle{
		Color: color,
		X:     x,
		Y:     y,
		Vx:    0,
		Vy:    0,
	}
}
