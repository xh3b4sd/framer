package framer

import (
	"time"

	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Sta time.Time
	End time.Time
}

type Framer struct {
	sta time.Time
	end time.Time
}

func New(con Config) (*Framer, error) {
	if con.Sta.IsZero() {
		return nil, tracer.Maskf(invalidConfigError, "%T.Sta must not be empty", con)
	}
	if con.End.IsZero() {
		return nil, tracer.Maskf(invalidConfigError, "%T.End must not be empty", con)
	}

	f := &Framer{
		sta: con.Sta,
		end: con.End,
	}

	return f, nil
}

func (f *Framer) Day() []Frame {
	var sta time.Time
	var end time.Time
	{
		sta = f.sta.UTC()
		end = f.end.UTC()
	}

	if !sta.Before(end) {
		return nil
	}

	var fra []Frame

	s := timday(sta)
	e := timday(sta.Add(day))

	for {
		fra = append(fra, Frame{Sta: s, End: e})

		if e.After(end) {
			break
		}

		s = s.Add(day)
		e = e.Add(day)
	}

	return fra
}

func (f *Framer) Exa() []Frame {
	var dfr []Frame
	{
		dfr = f.Lat()
	}

	{
		if len(dfr) == 0 {
			return nil
		}
	}

	{
		dfr[0].Sta = f.sta

		if len(dfr) > 0 {
			if dfr[len(dfr)-1].Sta == f.end {
				dfr = dfr[:len(dfr)-1]
			} else {
				dfr[len(dfr)-1].End = f.end
			}
		}
	}

	return dfr
}

func (f *Framer) Lat() []Frame {
	var dfr []Frame
	{
		dfr = f.Day()
	}

	{
		if len(dfr) == 0 {
			return nil
		}
	}

	{
		if len(dfr) > 0 {
			if dfr[len(dfr)-1].Sta == f.end {
				dfr = dfr[:len(dfr)-1]
			} else {
				dfr[len(dfr)-1].End = f.end
			}
		}
	}

	return dfr
}
