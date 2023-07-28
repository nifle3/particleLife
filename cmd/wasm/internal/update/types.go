package update

type Canvas interface {
	DrawAll([][2]int, []string)
}

type GroupParticle interface {
	MoveAllParticle()
	GetAllCoordinates() ([][2]int, []string)
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
