package framer

import (
	"time"
)

type Frames []Frame

// Con returns true if Frames contains Frame, based on the respective start and
// end time.
func (f Frames) Con(fra Frame) bool {
	for _, v := range f {
		if v.Sta.Equal(fra.Sta) && v.End.Equal(fra.End) {
			return true
		}
	}

	return false
}

// Len returns Frames restructured based on l. If Frames were a list of 5
// hourly frames and l were 10 minutes, then Len would return 30 frames of 10
// minutes length each. Optionally a tick size can be provided, which defaults
// to l.
func (f Frames) Len(l time.Duration, t ...time.Duration) Frames {
	s := f.Min().Sta.Truncate(l)
	e := cei(f.Max().End, l)

	if len(t) == 0 {
		t = append(t, l)
	}

	var n *Framer
	{
		n = New(Config{
			Sta: s,
			End: e,
			Len: l,
			Tic: t[0],
		})
	}

	return n.List()
}

// Max returns the Frame within Frames having the latest end time.
func (f Frames) Max() Frame {
	var max Frame

	for _, v := range f {
		if v.End.After(max.End) {
			max = v
		}
	}

	return max
}

// Min returns the Frame within Frames having the earliest start time.
func (f Frames) Min() Frame {
	min := f[0]

	for _, v := range f {
		if v.Sta.Before(min.Sta) {
			min = v
		}
	}

	return min
}

// Rem returns a new copy of Frames without any frame contained in rem.
func (f Frames) Rem(rem Frames) Frames {
	var cle Frames

	for _, v := range f {
		if !rem.Con(v) {
			cle = append(cle, v)
		}
	}

	return cle
}

func cei(t time.Time, d time.Duration) time.Time {
	f := t.Truncate(d)

	if f.Equal(t) {
		return t
	}

	return f.Add(d)
}
