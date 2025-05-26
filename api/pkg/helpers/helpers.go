package helpers

import (
	"log"
	"time"
)

func TimeHostNow(tz string) time.Time {
	// you can change Asia/Jakarta with your own location.
	// check on this https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	location, err := time.LoadLocation(tz)
	if err != nil {
		log.Printf("Error get time, cause:%+v\n", err)
	}
	now := time.Now()
	timeInLoc := now.In(location)
	return timeInLoc
}
