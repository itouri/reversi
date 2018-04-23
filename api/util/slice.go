package util

import "../domain"

func Remove(strings []string, search string) []string {
	result := []string{}
	for _, v := range strings {
		if v != search {
			result = append(result, v)
		}
	}
	return result
}

func Unset(s []domain.Player, i int) []domain.Player {
	if i >= len(s) {
		return s
	}
	return append(s[:i], s[i+1:]...)
}
