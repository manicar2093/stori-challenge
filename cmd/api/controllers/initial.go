package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Initial struct{}

func NewInitial() *Initial {
	return &Initial{}
}

func (c *Initial) SetUpRoutes(group *echo.Group) {
	group.GET("/", c.GetHelloHtmlMessage)
	group.GET("/json", c.GetHelloJsonMessage)
}

func (c *Initial) GetHelloHtmlMessage(ctx echo.Context) error {
	return ctx.HTML(http.StatusOK, "<b>Welcome to the starter template!</b>")
}

//	@Summary		Gets a json welcome message
//	@Description	Return a json content with a welcome message
//	@Tags			initial
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	echo.Map	"Demo data"
//	@Failure		500	"Something unidentified has occurred"
//	@Router			/ [get]
func (c *Initial) GetHelloJsonMessage(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "welcome to the starter template!",
	})
}
