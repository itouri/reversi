package controller

type Context interface {
	Param(string) string
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{}) error
	String(int, string) error
	NoContent(int) error
}
