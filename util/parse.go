package util

import (
	"log"
	"strconv"
)

func ParseBool(value string) bool {
	parsed, err := strconv.ParseBool(value)
	if err != nil {
		log.Println(err)
		return false
	}
	return parsed
}

func ParseInt(value string) (int, error) {
	parsed, err := strconv.Atoi(value)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return parsed, nil
}
