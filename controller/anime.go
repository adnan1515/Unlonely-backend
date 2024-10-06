package controller

import (
	"net/http"
	"rest/service"

	"github.com/labstack/echo/v4"
)

func GetRecomms(c echo.Context) error {
	recomm := service.GetRecommendations()
	c.JSON(http.StatusAccepted, recomm)
	return nil
}
