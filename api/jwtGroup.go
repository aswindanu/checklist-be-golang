package api

import (
	"golang-starterapi/handlers"

	"github.com/labstack/echo"
)

func JwtGroup(g *echo.Group) {
	g.GET("/main", handlers.MainJwt)
}
