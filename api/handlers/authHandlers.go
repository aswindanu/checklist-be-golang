package handlers

import (
	"log"
	"net/http"
	"golang-starterapi/models"
	"golang-starterbootstrap"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/thedevsaddam/govalidator"
)

type response struct {
	Token     string  `json:"token"`
	ExpiresIn float64 `json:"expires_in"`
}

type JwtClaims struct {
	Name    string `json:"name"`
	IsAdmin bool   `json:"is_admin"`
	jwt.StandardClaims
}

func LoginUser(c echo.Context) error {
	var (
		user models.User
	)

	username := c.QueryParam("username")
	password := c.QueryParam("password")

	rules := govalidator.MapData{
		"username": []string{"required"},
		"password": []string{"required"},
	}

	vld := ValidateQueryStr(c, rules)
	if vld != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, vld)
	}

	if bootstrap.App.DB.Where("username = ?", username).Where("password = ?", password).Find(&user).RecordNotFound() {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid username or password")
	} else {

		// create jwt token
		token, err := createJwtToken(strconv.FormatUint(user.ID, 10), "user")
		if err != nil {
			log.Println("Error Creating User JWT token", err)
			return c.String(http.StatusInternalServerError, "something went wrong")
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "You were logged in!",
			"token":   token,
		})
	}

	return c.String(http.StatusUnauthorized, "Your username or password were wrong")
}

func LoginAdmin(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	rules := govalidator.MapData{
		"username": []string{"required"},
		"password": []string{"required"},
	}

	vld := ValidateQueryStr(c, rules)
	if vld != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, vld)
	}
	adminUsername, adminPassword := bootstrap.App.AppConfig.String("admin_username"), bootstrap.App.AppConfig.String("admin_password")

	if username == adminUsername && password == adminPassword {
		// create jwt token
		token, err := createJwtToken(username, "admin")
		if err != nil {
			log.Println("Error Creating User JWT token", err)
			return c.String(http.StatusInternalServerError, "something went wrong")
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "You were logged in!",
			"token":   token,
		})
	}

	return c.String(http.StatusUnauthorized, "Your username or password were wrong")
}

func createJwtToken(uname string, jtype string) (string, error) {
	var (
		claim    JwtClaims
		lifeTime int64 = time.Now().Add(24 * time.Hour).Unix()
	)

	if jtype == "admin" {
		claim = JwtClaims{
			uname,
			true,
			jwt.StandardClaims{
				Id:        uname,
				ExpiresAt: lifeTime,
			},
		}
	} else {
		claim = JwtClaims{
			uname,
			false,
			jwt.StandardClaims{
				Id:        uname,
				ExpiresAt: lifeTime,
			},
		}
	}

	secret := bootstrap.App.AppConfig.String("jwt_secret")
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)
	token, err := rawToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}
