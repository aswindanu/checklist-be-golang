package middlewares

import (
	"encoding/json"

	"golang-startermodules/logger"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

// SetCompleteLogMiddlware Middleware for logging request and response
func SetCompleteLogMiddlware(e *echo.Echo) {

	log, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {

		var bodyJSON interface{}
		var bodyRESP interface{}

		json.Unmarshal(reqBody, &bodyJSON)
		json.Unmarshal(resBody, &bodyRESP)

		log.Logger.WithFields(logrus.Fields{
			"remote_ip": c.RealIP(),
			"protocol":  c.Request().Proto,
			"host":      c.Request().Host,
			"uri":       c.Request().RequestURI,
			"headers":   c.Request().Header,
			"request":   bodyJSON,
			"response":  bodyRESP,
		}).Info("REQUEST LOG")
	}))
}
