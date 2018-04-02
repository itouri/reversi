package util

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func WriteCookie(c echo.Context, key string, value string) {
	cookie := new(http.Cookie)
	cookie.Name = key
	cookie.Value = value
	// 1day
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	log.Println("writeCookie: ", key, value)
}

func ReadCookie(c echo.Context, key string) (string, error) {
	cookie, err := c.Cookie(key)
	if err != nil {
		return "", err
	}
	log.Println("readCookie: ", key, cookie.Value)
	return cookie.Value, nil
}
