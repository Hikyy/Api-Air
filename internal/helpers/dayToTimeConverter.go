package helpers

import "time"

func ConvertStringToStartOfDay(dateString string) (string, error) {
	layout := "2006-01-02" // Customize the layout based on your input string format

	date, err := time.Parse(layout, dateString)
	if err != nil {
		return "", err
	}

	// Set time to the start of the day
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

	startOfDayToMarshallString, _ := startOfDay.MarshalText()
	if err != nil {
		return "", nil
	}
	return string(startOfDayToMarshallString), nil
}

func ConvertStringToEndOfDay(dateString string) (string, error) {
	layout := "2006-01-02" // Customize the layout based on your input string format

	date, err := time.Parse(layout, dateString)
	if err != nil {
		return "", err
	}

	// Set time to the start of the next day and subtract 1 second
	endOfDay := time.Date(date.Year(), date.Month(), date.Day()+1, 0, 0, 0, 0, date.Location()).Add(-time.Second)

	endOfDayToMarshallString, _ := endOfDay.MarshalText()
	if err != nil {
		return "", nil
	}

	return string(endOfDayToMarshallString), nil
}
