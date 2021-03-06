package usecase

import "github.com/itouri/reversi/api/domain"

type RoomRepository interface {
	FindAll() (*[]domain.Room, error)
	FindByRoomID(string) (*domain.Room, error)
	UpsertRoomWithPlayers(string, domain.Players) error
	InsertRoom(domain.Room) error
	DeleteRoom(string) error
}
