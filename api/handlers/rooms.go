package handlers

import (
	"../../db"
	"github.com/labstack/echo"
)

func GetRooms(c echo.Context) error {
	db.Set()
	return
}

func PostRooms(c echo.Context) error {

}

func PutRooms(c echo.Context) error {

}
