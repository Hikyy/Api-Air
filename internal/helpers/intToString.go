package helpers

import "strconv"

func TransformStringToInt(value string) (int, error) {
	returnedValue, err := strconv.Atoi(value)
	if err != nil {
		return returnedValue, err
	}
	return returnedValue, err
}
