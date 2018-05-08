package controller

import (
	"log"
	"net/http"

	"github.com/itouri/reversi/api/domain"
	"github.com/itouri/reversi/api/interfaces/database"
	"github.com/itouri/reversi/api/usecase"
	uuid "github.com/satori/go.uuid"
)

type RoomController struct {
	Interactor usecase.RoomInteractor
}

type Req struct {
	RoomID     string `json:"room_id"`
	PlayerID   string `json:"player_id"`
	PlayerName string `json:"player_name"`
}

func NewRoomController(mongoHandler database.MongoHandler) *RoomController {
	return &RoomController{
		Interactor: usecase.RoomInteractor{
			RoomRepository: &database.RoomRepository{
				MongoHandler: mongoHandler,
				Collection:   "room",
			},
		},
	}
}

func (rc *RoomController) GetRooms(c Context) error {
	rooms, err := rc.Interactor.FindAll()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, rooms)
}

func (rc *RoomController) PostRooms(c Context) error {
	req := new(Req)
	if err := c.Bind(req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if req.PlayerName == "" {
		log.Println("player_name is required")
		return c.String(http.StatusBadRequest, "player_name is required")
	}

	playerID := uuid.Must(uuid.NewV4(), nil).String()
	roomID := uuid.Must(uuid.NewV4(), nil).String()

	room := &domain.Room{
		RoomID: roomID,
		Players: []domain.Player{
			domain.Player{
				ID:   playerID,
				Name: req.PlayerName,
			},
		},
	}
	err := rc.Interactor.CreateRoom(*room)
	if err != nil {
		return c.NoContent(http.StatusOK)
	}

	type res struct {
		RoomID   string `json:"room_id"`
		PlayerID string `json:"player_id"`
	}

	return c.JSON(http.StatusOK, res{roomID, playerID})
}

func (rc *RoomController) PutRooms(c Context) error {
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

	playerID := uuid.Must(uuid.NewV4(), nil).String()
	player := domain.Player{playerID, req.PlayerName}
	return rc.Interactor.AddPlayerToRoom(req.RoomID, player)
}

func (rc *RoomController) ExitRoom(c Context) error {
	roomID := c.Param("room_id")
	if roomID == "" {
		log.Println("room_id is required")
		return c.String(http.StatusBadRequest, "room_id is required")
	}

	playerID := c.Param("player_id")
	if playerID == "" {
		log.Println("player_id is required")
		return c.String(http.StatusBadRequest, "player_id is required")
	}
	return rc.Interactor.DeletePlayerFromRoom(roomID, playerID)
}
