package models

type Player struct {
	PlayerId string `validate:"required"`
	Name     string `validate:"required"`
}
