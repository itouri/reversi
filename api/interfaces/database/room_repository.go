package database

import (
	"../../domain"
	"gopkg.in/mgo.v2/bson"
)

type RoomRepository struct {
	MongoHandler
	Collection string
}

func (r *RoomRepository) FindAll() (*domain.Rooms, error) {
	return r.MongoHandler.FindAll(r.Collection)
}

func (r *RoomRepository) FindByRoomID(roomID string) (*domain.Room, error) {
	query := bson.M{"room_id": roomID}
	room := new(domain.Room)
	err := r.MongoHandler.FindOne(query, room)
	if err != nil {
		return nil, err
	}
	// room, ok := res.(*domain.Room)
	// if !ok {
	// 	return nil, fmt.Errorf("room is not type of room")
	// }
	return room, nil
}

func (r *RoomRepository) UpsertRoomWithPlayers(roomID string, players domain.Players) error {
	query := bson.M{"room_id": roomID}
	upsert := bson.M{"$set": bson.M{"players": players}}
	return r.MongoHandler.Upsert(r.Collection, query, upsert)
}

func (r *RoomRepository) InsertRoom(room domain.Room) error {
	return r.MongoHandler.Insert(r.Collection, room)
}

func (r *RoomRepository) DeleteRoom(roomID string) error {
	query := bson.M{"room_id": roomID}
	return r.MongoHandler.Delete(r.Collection, query)
}
