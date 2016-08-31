package controllers

import "github.com/revel/revel"

type UserAuth struct {
	*revel.Controller
}

func (c UserAuth) Login() revel.Result {
	return c.Render();
}
