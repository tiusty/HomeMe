package controllers

import (
	"github.com/revel/revel"
)

type Algorithim struct {
	Application
}

func (c Algorithim) Index() revel.Result {
	return c.Render()
}
