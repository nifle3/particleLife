package update

type Canvas interface {
	DrawAll([][3]int)
}

type GroupParticle interface {
	MoveAllParticle()
	GetAllCoordinates() [][3]int
}

type Updater struct {
	particles GroupParticle
	canvas    Canvas
}

func NewUpdater(particles GroupParticle, canvas Canvas) *Updater {
	return &Updater{
		particles: particles,
		canvas:    canvas,
	}
}
