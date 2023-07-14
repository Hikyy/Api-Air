package helpers

import (
	"fmt"
	"strconv"
	"time"
)

func TimeStampConverter(data uint) time.Time {
	conv := strconv.FormatUint(uint64(data), 10)
	i, err := strconv.ParseInt(conv, 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	tmUTC2 := tm.In(time.FixedZone("UTC+2", 2*60*60)) // Convertir en UTC+2

	fmt.Println(tmUTC2)
	return tmUTC2
}
