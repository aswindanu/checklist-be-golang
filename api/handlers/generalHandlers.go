package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/thedevsaddam/govalidator"
)

func ValidateRequest(c echo.Context, rules govalidator.MapData, data interface{}) map[string]interface{} {
	var err map[string]interface{}

	opts := govalidator.Options{
		Request: c.Request(),
		Data:    data,
		Rules:   rules,
		// RequiredDefault: true,
	}

	v := govalidator.New(opts)

	e := v.ValidateJSON()

	if len(e) > 0 {
		err = map[string]interface{}{"validation_errors": e}
	}

	return err
}

func ValidateQueryStr(c echo.Context, rules govalidator.MapData) map[string]interface{} {
	var err map[string]interface{}

	opts := govalidator.Options{
		Request: c.Request(),
		Rules:   rules,
		// RequiredDefault: true,
	}

	v := govalidator.New(opts)

	e := v.Validate()

	if len(e) > 0 {
		err = map[string]interface{}{"validation_errors": e}
	}

	return err
}

func Yallo(c echo.Context) error {
	return c.String(http.StatusOK, "yallo from the web side!")
}
