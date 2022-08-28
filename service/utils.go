package service

import (
	"fmt"
	"strconv"
)

func GetID(idString string) (id int, err error) {
	id, err = strconv.Atoi(idString)
	if id <= 0 {
		return 0, fmt.Errorf("invalid id")
	}

	return id, err
}
