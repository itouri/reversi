package database

import "../../domain"

type RoomRepository struct {
	MongoHandler
	Collection string
}

func (r *RoomRepository) FindAll() (domain.Rooms, error) {
	return r.MongoHandler.FindAll(r.Collection)
}

func (r *RoomRepository) FindByRoomID(string) (domain.Room, error) {

}

func (r *RoomRepository) UpsertRoomWithPlayers(string, domain.Players) error {

}

func (r *RoomRepository) InsertRoom(domain.Room) error {

}

func (r *RoomRepository) DeleteRoom(string) error {

}
