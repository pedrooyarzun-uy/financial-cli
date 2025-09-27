package validators

import (
	"errors"
	"strconv"
)

func CheckAmount(amountTxt string) (float64, error) {
	if amountTxt == "" {
		return 0, errors.New("Amount can't be empty")
	}

	amount, err := strconv.ParseFloat(amountTxt, 64)
	if err != nil {
		return 0, errors.New("Only numbers and numbers with . accepted")
	}

	if amount == 0.0 {
		return 0, errors.New("Amount can't be zero")
	}

	return amount, nil
}
