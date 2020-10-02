package handlers

import (
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func MainJwt(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		log.Println("User Name: ", claims["name"], "User ID: ", claims["jti"])
		return c.String(http.StatusOK, "you are on the top secret jwt page!")
	}

	return c.String(http.StatusForbidden, "Forbidden")
}
