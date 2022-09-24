package framer

import (
	"math"
	"time"
)

type Config struct {
	Sta time.Time
	End time.Time
	// Len is the frame length moving along a timeline. Frames produced by a
	// framer instance will always be Len long.
	Len time.Duration
	// Tic is the amount of time any next frame is moving forward. By default
	// Tic will be set to Len.
	Tic time.Duration
}

func (c Config) Ensure() Config {
	if c.Tic == 0 {
		c.Tic = c.Len
	}

	return c
}

func (c Config) Verify() {
	if c.Sta.IsZero() {
		panic("Config.Sta must not be empty")
	}
	if c.End.IsZero() {
		panic("Config.End must not be empty")
	}
	if c.Len == 0 {
		panic("Config.Len must not be empty")
	}
	if c.Tic != 0 && math.Mod(float64(c.Len), float64(c.Tic)) != 0 {
		panic("Config.Tic must be a smaller multiple of Config.Len")
	}
}
