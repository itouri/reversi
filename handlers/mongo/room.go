package mongo

import (
	"../../db"
	"../../models"
)

func GetAllRooms() {
	dbb := &db.DbBase{}
	dbb.Collection(models.Player.String())
}
