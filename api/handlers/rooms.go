package handlers

import (
	"log"
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

	player_name := c.QueryParam("player_name")
	if player_name == "" {
		return c.String(http.StatusBadRequest, "player_name is required")
	}

	roomID := uuid.Must(uuid.NewV4()).String()
	room := &models.Room{
		RoomID: roomID,
		PlayerNames: []string{
			player_name,
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
		log.Println("room_id is required")
		return c.String(http.StatusBadRequest, "room_id is required")
	}

	player_name := c.QueryParam("player_name")
	if player_name == "" {
		log.Println("player_name is required")
		return c.String(http.StatusBadRequest, "player_name is required")
	}

	query := bson.M{"room_id": roomID}

	// TODO upsertのやりかたが間違ってる
	err := dbb.Collection(room.String()).Find(query).One(room)
	if err != nil {
		log.Println(err)
		return c.NoContent(http.StatusOK)
	}

	if len(room.PlayerNames) > 1 {
		return c.String(http.StatusBadRequest, "This room is full")
	}
	room.PlayerNames = append(room.PlayerNames, player_name)

	upsert := bson.M{"$set": bson.M{"player_names": room.PlayerNames}}
	_, err = dbb.Collection(room.String()).Upsert(query, upsert)
	if err != nil {
		return c.NoContent(http.StatusOK)
	}
	return c.NoContent(http.StatusOK)
}
