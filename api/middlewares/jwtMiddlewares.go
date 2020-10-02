package middlewares

import (
	"fmt"
	"golang-starterapi/models"
	"golang-starterbootstrap"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//SetJwtMiddlewares Set Only JWT for For User
func SetJwtMiddlewares(g *echo.Group) {

	secret := bootstrap.App.DBConfig.String("jwt_secret")
	// validate jwt token
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(secret),
	}))

	// validate payload related with user
	g.Use(validateJwtUser)
}

//SetJwtAdminMiddlewares Set Only JWT for For Admin
func SetJwtAdminMiddlewares(g *echo.Group) {

	secret := bootstrap.App.DBConfig.String("jwt_secret")

	// validate jwt token
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(secret),
	}))

	// validate payload related with admin type of token
	g.Use(validateJwtAdmin)
}

//Setting Middleware To Get Access Either ADMIN or MERCHANT
func SetJwtGeneralMiddlewares(g *echo.Group) {
	secret := bootstrap.App.DBConfig.String("jwt_secret")

	// validate jwt token
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(secret),
	}))

	// validate payload related with admin type of token
	g.Use(ValidateGeneralJwt)
}

// validateJwtAdmin
// Middleware for validating access to Admin only resources
func validateJwtAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		user := c.Get("user")
		token := user.(*jwt.Token)

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if claims["is_admin"] == true {
				return next(c)
			} else {
				return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
			}
		}

		return echo.NewHTTPError(http.StatusForbidden, "Invalid Token")
	}
}

func validateJwtUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			user := models.User{}
			mid, _ := strconv.Atoi(fmt.Sprintf("%s", claims["jti"]))
			_, err := user.FindByID(mid)

			if err != nil {
				return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
			}

			c.Set("user", user)

			return next(c)
		}

		return echo.NewHTTPError(http.StatusForbidden, "Invalid Token")
	}
}

//ValidateGeneralJwt Use this method to Get Data Either ADMIN or MERCHANT
func ValidateGeneralJwt(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if claims["is_admin"] == true {
				return next(c)
			} else {
				user := models.User{}
				mid, _ := strconv.Atoi(fmt.Sprintf("%s", claims["jti"]))
				_, err := user.FindByID(mid)

				if err != nil {
					return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
				}

				c.Set("merchant", user)
			}
			return next(c)
		}

		return echo.NewHTTPError(http.StatusForbidden, "Invalid Token")
	}
}
