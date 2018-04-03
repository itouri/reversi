package handlers

import (
	"log"
	"net/http"

	"../../db"
	"../../models"
	"../util"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// TODO グローバル変数にしていいのかな
var dbRoom *mgo.Collection

func init() {
	room := &models.Room{}
	dbb := &db.DbBase{}
	dbRoom = dbb.Collection(room.String())
}

type Req struct {
	RoomID     string `json:"room_id"`
	PlayerID   string `json:"player_id"`
	PlayerName string `json:"player_name"`
}

func GetRooms(c echo.Context) error {
	rooms := []models.Room{}
	err := dbRoom.Find(nil).All(&rooms)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, rooms)
}

func PostRooms(c echo.Context) error {
	req := new(Req)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if req.PlayerName == "" {
		log.Println("player_name is required")
		return c.String(http.StatusBadRequest, "player_name is required")
	}

	playerID := uuid.Must(uuid.NewV4()).String()
	roomID := uuid.Must(uuid.NewV4()).String()

	room := &models.Room{
		RoomID: roomID,
		Players: []models.Player{
			models.Player{
				ID:   playerID,
				Name: req.PlayerName,
			},
		},
	}
	err := dbRoom.Insert(room)
	if err != nil {
		return c.NoContent(http.StatusOK)
	}

	type res struct {
		RoomID   string `json:"room_id"`
		PlayerID string `json:"player_id"`
	}

	return c.JSON(http.StatusOK, res{roomID, playerID})
}

func PutRooms(c echo.Context) error {
	req := new(Req)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if req.RoomID == "" {
		log.Println("room_id is required")
		return c.String(http.StatusBadRequest, "room_id is required")
	}

	if req.PlayerName == "" {
		log.Println("player_name is required")
		return c.String(http.StatusBadRequest, "player_name is required")
	}

	player_id := uuid.Must(uuid.NewV4()).String()

	query := bson.M{"room_id": req.RoomID}

	room := &models.Room{}
	// TODO upsertのやりかたが間違ってる
	err := dbRoom.Find(query).One(room)
	if err != nil {
		log.Println(err)
		return c.NoContent(http.StatusOK)
	}

	if len(room.Players) > 1 {
		return c.String(http.StatusBadRequest, "This room is full")
	}

	player := models.Player{player_id, req.PlayerName}
	room.Players = append(room.Players, player)

	upsert := bson.M{"$set": bson.M{"players": room.Players}}
	_, err = dbRoom.Upsert(query, upsert)
	if err != nil {
		return c.NoContent(http.StatusOK)
	}
	// Stringで返してもAngular側で受け取れない！subsriberがJSONでmapしてるから(多分)
	return c.JSON(http.StatusOK, player_id)
}

func ExitRoom(c echo.Context) error {
	// TODO いちいちコレ書くの面倒
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
	err := dbRoom.Find(query).One(room)
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
			room.Players = util.Unset(room.Players, i)
		}
	}

	// TODO
	upsert := bson.M{"$set": bson.M{"players": room.Players}}
	_, err = dbRoom.Upsert(query, upsert)
	if err != nil {
		return c.NoContent(http.StatusOK)
	}
	return c.NoContent(http.StatusOK)
}

func deleteRoom(roomID string) error {
	query := bson.M{"room_id": roomID}
	return dbRoom.Remove(query)
}
