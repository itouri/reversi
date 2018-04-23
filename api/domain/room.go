package domain

type Rooms []Rooms

type Room struct {
	RoomID         string   `bson:"room_id" validate:"required"`
	Players        []Player `bson:"players" validate:"required"`
	SpectatorNames []string `bson:"spector_names"`
}

func (r *Room) String() string {
	return "room"
}
