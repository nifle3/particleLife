package update

func (u *Updater) Update(quit chan bool) {
	for {
		select {
		case <-quit:
			break

		default:
			u.particles.MoveAllParticle()
			u.canvas.DrawAll(u.particles.GetAllCoordinates())
		}
	}
}
