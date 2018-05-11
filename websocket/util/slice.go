package util

import (
	"github.com/itouri/reversi/websocket/models"
)

func Remove(strings []string, search string) []string {
	result := []string{}
	for _, v := range strings {
		if v != search {
			result = append(result, v)
		}
	}
	return result
}

func Unset(s []models.Player, i int) []models.Player {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}
