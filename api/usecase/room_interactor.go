package usecase

import (
	"fmt"

	"github.com/itouri/reversi/api/domain"
	"github.com/itouri/reversi/api/util"
)

type RoomInteractor struct {
	RoomRepository RoomRepository
}

func (ri *RoomInteractor) FindAll() (*[]domain.Room, error) {
	return ri.RoomRepository.FindAll()
}

// 最初に見つかったものだけ返す
func (ri *RoomInteractor) FindByRoomID(roomID string) (*domain.Room, error) {
	return ri.RoomRepository.FindByRoomID(roomID)
}

func (ri *RoomInteractor) AddPlayerToRoom(roomID string, player domain.Player) error {
	room, err := ri.RoomRepository.FindByRoomID(roomID)
	if err != nil {
		return err
	}
	if len(room.Players) > 1 {
		return fmt.Errorf("This room is full")
	}
	room.Players = append(room.Players, player)
	// upsert って MySQL と MongoDB で共通化できなくないか
	// TODO mongo非依存
	return ri.RoomRepository.UpsertRoomWithPlayers(roomID, room.Players)
}

func (ri *RoomInteractor) DeletePlayerFromRoom(roomID string, playerID string) error {
	room, err := ri.RoomRepository.FindByRoomID(roomID)
	if err != nil {
		return err
	}
	if len(room.Players) == 1 {
		return ri.DeleteRoom(roomID)
	}

	for i, roomPlayer := range room.Players {
		if roomPlayer.ID == playerID {
			room.Players = util.Unset(room.Players, i)
		}
	}
	return ri.RoomRepository.UpsertRoomWithPlayers(roomID, room.Players)
}

func (ri *RoomInteractor) CreateRoom(room domain.Room) error {
	return ri.RoomRepository.InsertRoom(room)
}

func (ri *RoomInteractor) DeleteRoom(roomID string) error {
	return ri.RoomRepository.DeleteRoom(roomID)
}
