package models

type Room struct {
	RoomId         string   `bson:"room_id" validate:"required"`
	PlayerIds      []string `bson:"player_ids" validate:"required"`
	SpectatorNames []string `bson:"spector_names"`
}

func (r *Room) String() string {
	return "room"
}
