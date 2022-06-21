package framer

import (
	"time"
)

const (
	day = 24 * time.Hour
)

type Frame struct {
	Sta time.Time
	End time.Time
}
