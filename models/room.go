package models

type Room struct {
	RoomID         string   `bson:"room_id" validate:"required"`
	PlayerNames    []string `bson:"player_names" validate:"required"`
	SpectatorNames []string `bson:"spector_names"`
}

func (r *Room) String() string {
	return "room"
}
