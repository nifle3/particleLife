package Canvas

func (c *Canvas) DrawAll(coordinates [][3]int) {
	c.context.Call("clearRect", 0, 0, c.width, c.height)

	for _, value := range coordinates {
		c.context.Set("fillStyle", c.getColor(value[2]))
		c.context.Call("fillRect", value[0], value[1], 10, 10)
	}
}

func (c *Canvas) getColor(colorCode int) string {
	switch colorCode {
	case 0:
		return "white"
	case 1:
		return "red"
	case 2:
		return "yellow"
	default:
		return "black"
	}
}
