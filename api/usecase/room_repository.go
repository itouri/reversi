package usecase

import "../domain"

type RoomRepository interface {
	FindAll() (domain.Rooms, error)
	FindByRoomID(string) (domain.Room, error)
	UpsertRoomWithPlayers(string, domain.Players) error
	InsertRoom(domain.Room) (domain.Rooms, error)
	DeleteRoom(string) error
}
