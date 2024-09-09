package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/srazash/RaceTracker/controllers"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.Static("/", "static")

	e.GET("/", controllers.Index)
	e.GET("/h1", controllers.H1)

	e.Logger.Fatal(e.Start(":1234"))
}
