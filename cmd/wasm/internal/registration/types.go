package registration

import "particleLife/cmd/wasm/pkg/particle"

type Canvas interface {
	DrawAll(particle []particle.Particle)
	Clear()
}

type ParticleGroup interface {
	MoveAll()
	GetAllParticle() []particle.Particle
}

type JsCallBack struct {
	canvas    Canvas
	particles ParticleGroup
}

func NewJsCallBack(canvas Canvas, particles ParticleGroup) JsCallBack {
	return JsCallBack{
		canvas:    canvas,
		particles: particles,
	}
}
