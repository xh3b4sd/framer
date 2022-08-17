package framer

import (
	"time"
)

type Frame struct {
	Sta time.Time
	End time.Time
}

func (f Frame) Empty() bool {
	return f.Sta.IsZero() && f.End.IsZero()
}
