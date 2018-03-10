package models

type Player struct {
	DbBase   `bson:omitempty`
	PlayerId string `bson:player_id`
	Name     string `bson:name`
}
