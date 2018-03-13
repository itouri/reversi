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
	room := &models.Room{}
	rooms := []models.Room{}
	err := dbb.Collection(room.String()).Find(nil).All(&rooms)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, rooms)
}

func PostRooms(c echo.Context) error {
	dbb := &db.DbBase{}

	name := c.QueryParam("name")
	if name == "" {
		return c.String(http.StatusBadRequest, "Name is required")
	}

	roomID := uuid.Must(uuid.NewV4()).String()
	room := &models.Room{
		RoomId: roomID,
		PlayerIds: []string{
			name,
		},
	}
	err := dbb.Collection(room.String()).Insert(room)
	if err != nil {
		return c.NoContent(http.StatusOK)
	}
	return c.NoContent(http.StatusOK)
}

func PutRooms(c echo.Context) error {
	dbb := &db.DbBase{}
	room := &models.Room{}

	roomID := c.QueryParam("room_id")
	if roomID == "" {
		return c.String(http.StatusBadRequest, "room_id is required")
	}

	name := c.QueryParam("name")
	if name == "" {
		return c.String(http.StatusBadRequest, "Name is required")
	}

	query := bson.M{"room_id": roomID}

	// TODO upsertのやりかたが間違ってる
	err := dbb.Collection(room.String()).Find(query).One(room)
	if err != nil {
		return c.NoContent(http.StatusOK)
	}

	if len(room.PlayerIds) > 1 {
		return c.String(http.StatusBadRequest, "This room is full")
	}
	room.PlayerIds = append(room.PlayerIds, name)

	upsert := bson.M{"$set": bson.M{"player_ids": room.PlayerIds}}
	_, err = dbb.Collection(room.String()).Upsert(query, upsert)
	if err != nil {
		return c.NoContent(http.StatusOK)
	}
	return c.NoContent(http.StatusOK)
}
