package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

 

func (c App) Index() revel.Result {
	greeting := "Alex Agudelo"
	return c.Render(greeting)
}

func (c App) Hello(myName string) revel.Result {
	return c.Render(myName)
}
