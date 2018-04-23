package database

import "../../domain"

type MongoHandler interface {
	Find(interface{}) (interface{}, error)
	FindOne(interface{}) (interface{}, error)
	FindAll(string) (*domain.Rooms, error)
	Insert(string, interface{}) error
	Upsert(string, interface{}, interface{}) error
	Delete(string, interface{}) error
}
