package controllers

type Algorithim struct {
	Application
}

func (c Algorithim) Index() {
	return c.Render()
}
