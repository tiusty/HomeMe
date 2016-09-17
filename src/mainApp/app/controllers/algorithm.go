package controllers

import (
	"github.com/revel/revel"
)

type Algorithm struct {
	Application
}

func (c Algorithm) Index(streetAddress, zip, state, city, rentBuy string) revel.Result {
	c.Validation.Required(streetAddress).Message("You need a destination!")
    c.Validation.Required(zip).Message("You need a Zip Code")
    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(App.Survey)
    }
    c.RenderArgs["streetAddressInitial"] = streetAddress
    c.RenderArgs["zipInitial"] = zip
    c.RenderArgs["stateInitial"] = state
    c.RenderArgs["cityInitial"] = city
    if(rentBuy == "rent") {
        return c.RenderTemplate("Algorithm/rentingSurvey.html")
    } else {
        return c.Redirect(App.Survey)
    }

}

func (c Algorithm) RentingSurvey() revel.Result {
    return c.Render()
}
