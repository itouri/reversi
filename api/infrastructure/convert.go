package infrastructure

import (
	"../interfaces/controller"
	"github.com/labstack/echo"
)

var roomController *controller.RoomController

func init() {
	roomController = controller.NewRoomController(NewMongoHandler())
}

func GetRooms(c echo.Context) error {
	return roomController.GetRooms(c)
}

func PostRooms(c echo.Context) error {
	return roomController.PostRooms(c)
}

func PutRooms(c echo.Context) error {
	return roomController.PutRooms(c)
}

// func ExitRoom(c echo.Context) error {

// }
