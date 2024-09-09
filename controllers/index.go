package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/srazash/RaceTracker/views"
)

func Index(c echo.Context) error {
	return Render(c, http.StatusOK, views.Layout("demo", views.Index()))
}

func H1(c echo.Context) error {
	return c.String(http.StatusOK, "I live!")
}
