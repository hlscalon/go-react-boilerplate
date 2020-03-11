package utils

import (
	"strconv"
)

func StrToInt(str string) (int, error) {
	if str == "" {
		return 0, nil
	}

	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return i, nil
}
