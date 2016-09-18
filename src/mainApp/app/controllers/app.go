package controllers

import "github.com/revel/revel"
import "log"

type App struct {
	Application
}

/*func (c App) Index() revel.Result {
	greeting := "Alex Agudelo"
	return c.Render(greeting)
}*/

/*func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("Your name is required!")
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.Render(myName)
}*/

func (c App) Survey() revel.Result {
	loggedIn := false
	if c.connected() != nil {
		loggedIn = true
	}
	return c.Render(loggedIn)
}

func (c App) SurveyResult() revel.Result {
	var streetAddress, zip, city, state string
	var typeHome []string
	var numRooms, numBathroom []int
	var airConditioning, washerDryerHome, dishWasher, bath, parkingSpot, washerDryerBuilding, elevator, costReturnLow, costReturnHigh, amountSizeLow, amountSizeHigh, amountMaxCommuteLow, amountMaxCommuteHigh int
	c.Params.Bind(&streetAddress, "streetAddress")
	c.Params.Bind(&zip, "zip")
	c.Params.Bind(&city, "city")
	c.Params.Bind(&state, "state")
	c.Params.Bind(&amountMaxCommuteLow, "amountMaxCommuteLow")
	c.Params.Bind(&amountMaxCommuteHigh, "amountMaxCommuteHigh")
	c.Params.Bind(&costReturnLow, "costReturnLow")
	c.Params.Bind(&costReturnHigh, "costReturnHigh")
	c.Params.Bind(&amountSizeLow, "amountSizeLow")
	c.Params.Bind(&amountSizeHigh, "amountSizeHigh")
	c.Params.Bind(&typeHome, "typeHome")
	c.Params.Bind(&numRooms, "numRooms")
	c.Params.Bind(&numBathroom, "numBathroom")
	c.Params.Bind(&washerDryerHome, "washerDryerHome")
	c.Params.Bind(&dishWasher, "dishWasher")
	c.Params.Bind(&bath, "bath")
	c.Params.Bind(&parkingSpot, "parkingSpot")
	c.Params.Bind(&washerDryerBuilding, "washerDryerBuilding")
	c.Params.Bind(&elevator, "elevator")
	c.Params.Bind(&airConditioning, "airConditioning")
	log.Println(costReturnLow)
	log.Println(costReturnHigh)
    log.Println(amountSizeHigh)
    log.Println(amountSizeLow)
    log.Println(amountMaxCommuteHigh)
    log.Println(amountMaxCommuteLow)


	c.Validation.Required(streetAddress).Message("You need a destination!")
	c.Validation.Required(zip).Message("You need a Zip Code")
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Survey)
	}
	return c.Render(streetAddress, zip, state, city, typeHome, numRooms, numBathroom, airConditioning, parkingSpot, amountSizeHigh, costReturnLow, amountMaxCommuteHigh)
}
