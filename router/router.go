package router

import (
	"golang-starterapi"
	"golang-starterapi/middlewares"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	// router groups
	jwtGroup := e.Group("/jwt")

	// set all middlewares
	middlewares.SetMainMiddlewares(e)
	middlewares.SetCompleteLogMiddlware(e)

	middlewares.SetJwtAdminMiddlewares(jwtGroup)

	// set main routes
	api.MainGroup(e)

	// set group routes
	api.JwtGroup(jwtGroup)

	return e
}
