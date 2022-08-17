package framer

import (
	"time"
)

type Config struct {
	Sta time.Time
	End time.Time
	Dur time.Duration
}

type Framer struct {
	sta time.Time
	end time.Time
	dur time.Duration
	// poi is a pointer to the latest end time from which the next frame can be
	// constructured.
	poi time.Time
}

func New(con Config) *Framer {
	if con.Sta.IsZero() {
		panic("Config.Sta must not be empty")
	}
	if con.End.IsZero() {
		panic("Config.End must not be empty")
	}
	if con.Dur == 0 {
		panic("Config.Dur must not be empty")
	}

	return &Framer{
		sta: con.Sta.UTC(),
		end: con.End.UTC(),
		dur: con.Dur,
		poi: con.Sta.UTC(),
	}
}

func (f *Framer) Last() bool {
	return !f.poi.Before(f.end)
}

func (f *Framer) List() Frames {
	return fra(f.sta, f.end, f.dur)
}

func (f *Framer) Next() Frame {
	if f.Last() {
		return Frame{}
	}

	var sta time.Time
	var end time.Time
	{
		sta = f.poi
		end = sta.Add(f.dur).Truncate(f.dur)
	}

	if end.After(f.end) {
		end = f.end
	}

	{
		f.poi = end
	}

	return Frame{
		Sta: sta,
		End: end,
	}
}
