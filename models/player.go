package models

type Player struct {
	ID   string `validate:"required"`
	Name string `validate:"required"`
}

func (p *Player) String() string {
	return "player"
}
