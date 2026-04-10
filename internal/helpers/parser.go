package helpers

import (
	"strconv"
)

func StringToInt(txt string) (int, bool) {
	val, err := strconv.Atoi(txt)

	if err != nil {
		return 0, false
	}

	return val, true
}
