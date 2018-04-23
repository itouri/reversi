package database

import "../../domain"

type MongoHandler interface {
	Find(interface{}) error
	FindAll(string) (domain.Rooms, error)
}
