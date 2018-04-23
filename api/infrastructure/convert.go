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

// func PostRooms(c echo.Context) error {

// }

// func PutRooms(c echo.Context) error {

// }

// func ExitRoom(c echo.Context) error {

// }
