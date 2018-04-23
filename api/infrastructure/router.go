package infrastructure

import (
	"../interfaces/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())

	roomController := controller.NewRoomController(NewMongoHandler())

	// パスの最後に / をつけるといけない
	e.GET("/api/v1/rooms", roomController.GetRooms)
	e.POST("/api/v1/rooms", roomController.PostRooms)
	e.PUT("/api/v1/rooms", roomController.PutRooms)

	e.DELETE("/api/v1/rooms/:room_id/:player_id", roomController.ExitRoom)

	e.Start(":12345")
}
