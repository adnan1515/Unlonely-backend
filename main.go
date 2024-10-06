package main

import (
	"net/http"
	"rest/controller"

	log "rest/logging"
	"rest/persistence"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	_, err := persistence.Connect()
	if err != nil {
		log.Error("Failed to connect to Db")
	}

	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/", hello)
	e.POST("/login", controller.Login)
	e.POST("/register", controller.Register)
	e.GET("/user", controller.User)
	e.GET("/logout", controller.Logout)
	e.GET("/getRecomm", controller.GetRecomms)
	log.Info("Starting server...")
	log.Error(e.Start(":1323"))
}
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
