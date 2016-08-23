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
	c.Validation.Required(myName).Message("Your name is required!")
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")
	
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.Render(myName)
}

func (c App) Survey() revel.Result{
	return c.Render()
}

func (c App) SurveyResult(streetAddress, zip string) revel.Result {
	c.Validation.Required(streetAddress).Message("You need a destination!")
	c.Validation.Required(zip).Message("You need a Zip Code")	
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Survey)
	}	
	return c.Render(streetAddress, zip)
}
