package api

import (
	"golang-starterapi/handlers"
	"golang-starterapi/middlewares"

	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	e.GET("/login", handlers.LoginUser)
	e.GET("/login/admin", handlers.LoginAdmin)

	e.GET("/yallo", handlers.Yallo)
	e.GET("/cats/:data", handlers.GetCats)

	e.POST("/cats", handlers.AddCat)
	e.POST("/dogs", handlers.AddDog)
	e.POST("/hamsters", handlers.AddHamster)

	//API Group For Normal User
	g := e.Group("/users")
	middlewares.SetJwtMiddlewares(g)
	g.GET("", handlers.GetUserById)
	g.GET("/list", handlers.GetUsers)
	g.POST("", handlers.AddUser)
	g.PUT("", handlers.EditUser)
	g.DELETE("", handlers.DeleteUser)

	//API Group For Admin
	g = e.Group("admin/users")
	middlewares.SetJwtAdminMiddlewares(g)
	g.GET("", handlers.GetUserById)
	g.GET("/list", handlers.GetUsers)
	g.POST("", handlers.AddUser)
	g.PUT("", handlers.EditUser)
	g.DELETE("", handlers.DeleteUser)

	//API Group For Either Admin or User
	g = e.Group("general/users")
	middlewares.SetJwtGeneralMiddlewares(g)
	g.GET("", handlers.GetUserById)
	g.GET("/list", handlers.GetUsers)
	g.POST("", handlers.AddUser)
	g.PUT("", handlers.EditUser)
	g.DELETE("", handlers.DeleteUser)
}
