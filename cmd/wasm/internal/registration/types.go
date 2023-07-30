package registration

type Canvas interface {
	DrawAll()
	Clear()
}

type ParticleGroup interface {
	MoveAll()
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
