package models

type Room struct {
	PlayerIds      [2]string `validate:"required"`
	SpectatorNames []string  `validate:"required"`
}
