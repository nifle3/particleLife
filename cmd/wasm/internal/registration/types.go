package registration

import "particleLife/cmd/wasm/pkg/particle"

type Canvas interface {
	DrawAll(particle [][]particle.Particle)
	Clear()
}

type ParticleGroup interface {
	MoveAll()
	GetAllParticle() [][]particle.Particle
	GenerateGroups(countBlue, countRed, countBlack int)
	SetRules(string, float64)
}

type JsCallBack struct {
	canvas    Canvas
	particles ParticleGroup
	isSetup   bool
}

func NewJsCallBack(canvas Canvas, particles ParticleGroup) JsCallBack {
	return JsCallBack{
		canvas:    canvas,
		particles: particles,
		isSetup:   false,
	}
}
