package controller

type Context interface {
	Param(string) string
	Bind(interface{}) error
	JSON(int, interface{}) error
	String(int, string) error
	NoContent(int) error
}
