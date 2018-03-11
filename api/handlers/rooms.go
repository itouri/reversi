package handlers

import (
	"net/http"

	"../../db"
	"../../models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
)

func GetRooms(c echo.Context) error {
	dbb := &db.DbBase{}
	rooms := []models.Room{}
	err := dbb.Collection(rooms[0].String()).Find(nil).All(rooms)
	if err != nil {
		return c.NoContent(http.StatusOK)
	}
	return c.JSON(http.StatusOK, rooms)
}

func PostRooms(c echo.Context) error {
	dbb := &db.DbBase{}
	room := &models.Room{}
	roomID := uuid.Must(uuid.NewV4()).String()
	selector := bson.M{"room_id": roomID}
	err := dbb.Collection(rooms[0].String()).Find(nil).All(rooms)
	if err != nil {
		return c.NoContent(http.StatusOK)
	}
	return c.NoContent(http.StatusOK)
}

func PutRooms(c echo.Context) error {

}
