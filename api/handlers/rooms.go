package handlers

import (
	"log"
	"net/http"
	"time"

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
		log.Println("player_name is required")
		return c.String(http.StatusBadRequest, "player_name is required")
	}

	player_id := c.QueryParam("player_id")
	if player_id == "" {
		player_id = uuid.Must(uuid.NewV4()).String()
	}

	roomID := uuid.Must(uuid.NewV4()).String()
	room := &models.Room{
		RoomID: roomID,
		Players: []models.Player{
			models.Player{
				ID:   player_id,
				Name: player_name,
			},
		},
	}
	err := dbb.Collection(room.String()).Insert(room)
	if err != nil {
		return c.NoContent(http.StatusOK)
	}

	type res struct {
		RoomID   string
		PlayerID string
	}

	return c.JSON(http.StatusOK, res{roomID, player_id})
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

	player_id := c.QueryParam("player_id")
	if player_id == "" {
		player_id = uuid.Must(uuid.NewV4()).String()
	}

	query := bson.M{"room_id": roomID}

	// TODO upsertのやりかたが間違ってる
	err := dbb.Collection(room.String()).Find(query).One(room)
	if err != nil {
		log.Println(err)
		return c.NoContent(http.StatusOK)
	}

	if len(room.Players) > 1 {
		return c.String(http.StatusBadRequest, "This room is full")
	}

	player := models.Player{player_id, player_name}
	room.Players = append(room.Players, player)

	upsert := bson.M{"$set": bson.M{"players": room.Players}}
	_, err = dbb.Collection(room.String()).Upsert(query, upsert)
	if err != nil {
		return c.NoContent(http.StatusOK)
	}
	// Stringで返してもAngular側で受け取れない！subsriberがJSONでmapしてるから(多分)
	return c.JSON(http.StatusOK, player_id)
}

func ExitRoom(c echo.Context) error {
	// TODO いちいちコレ書くの面倒
	dbb := &db.DbBase{}
	room := &models.Room{}

	roomID := c.Param("room_id")
	if roomID == "" {
		log.Println("room_id is required")
		return c.String(http.StatusBadRequest, "room_id is required")
	}

	player_id := c.Param("player_id")
	if player_id == "" {
		log.Println("player_id is required")
		return c.String(http.StatusBadRequest, "player_id is required")
	}

	query := bson.M{"room_id": roomID}

	// TODO upsertのやりかたが間違ってる
	err := dbb.Collection(room.String()).Find(query).One(room)
	if err != nil {
		log.Println(err)
		return c.NoContent(http.StatusOK)
	}

	if len(room.Players) == 1 {
		deleteRoom(roomID)
		return c.NoContent(http.StatusOK)
	}

	for i, player := range room.Players {
		if player.ID == player_id {
			room.Players = unset(room.Players, i)
		}
	}

	// TODO
	upsert := bson.M{"$set": bson.M{"players": room.Players}}
	_, err = dbb.Collection(room.String()).Upsert(query, upsert)
	if err != nil {
		return c.NoContent(http.StatusOK)
	}
	return c.NoContent(http.StatusOK)
}

func deleteRoom(roomID string) error {
	dbb := &db.DbBase{}
	room := &models.Room{}

	query := bson.M{"room_id": roomID}

	err := dbb.Collection(room.String()).Remove(query)
	if err != nil {
		return err
	}

	return nil
}

// TODO utilへ
func remove(strings []string, search string) []string {
	result := []string{}
	for _, v := range strings {
		if v != search {
			result = append(result, v)
		}
	}
	return result
}

func unset(s []models.Player, i int) []models.Player {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}

func writeCookie(c echo.Context, key string, value string) {
	cookie := new(http.Cookie)
	cookie.Name = key
	cookie.Value = value
	// 1day
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	log.Println("writeCookie: ", key, value)
}

func readCookie(c echo.Context, key string) (string, error) {
	cookie, err := c.Cookie(key)
	if err != nil {
		return "", err
	}
	log.Println("readCookie: ", key, cookie.Value)
	return cookie.Value, nil
}
