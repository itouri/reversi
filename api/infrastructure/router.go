package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())

	// パスの最後に / をつけるといけない
	e.GET("/api/v1/rooms", GetRooms)
	e.POST("/api/v1/rooms", PostRooms)
	e.PUT("/api/v1/rooms", PutRooms)

	// e.DELETE("/api/v1/rooms/:room_id/:player_id", ExitRoom)

	e.Start(":12345")
}
