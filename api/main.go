package main

import (
	"./handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())

	e.GET("/api/v1/rooms", handlers.GetRooms)
	e.POST("/api/v1/rooms", handlers.PostRooms)
	// e.PUT("/api/v1/rooms/:id", handlers.UpdateRooms)

	// Start server
	//e.Run(standard.New(":1323"))
	e.Start(":12345")
}
