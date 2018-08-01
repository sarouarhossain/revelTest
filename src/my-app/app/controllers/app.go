package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

type Result struct {
	Name string
	Age  int
	Job  string
}

func (c App) Index() revel.Result {
	rtdata := Result{"Roman", 26, "Engineer"}
	return c.RenderJSON(rtdata)
	//return c.Render(greeting)
}

func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("Name is required!")
	c.Validation.MinSize(myName, 3).Message("Name is not long enough")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.Render(myName)
}
