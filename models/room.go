package models

type Room struct {
	PlayerIds      []string `validate:"required"`
	SpectatorNames []string `validate:"required"`
}

func (r *Room) String() string {
	return "room"
}
