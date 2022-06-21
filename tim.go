package framer

import "time"

func timday(tim time.Time) time.Time {
	return tim.UTC().Truncate(24 * time.Hour)
}
