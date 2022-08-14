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

func (f Frames) Dur(dur time.Duration) Frames {
	max := ceiling(f.Max().End, dur)
	min := f.Min().Sta.Truncate(dur)

	var fra []Frame

	s := min
	e := min.Add(dur)

	for {
		fra = append(fra, Frame{Sta: s, End: e})

		if e.Equal(max) {
			break
		}

		s = s.Add(dur)
		e = e.Add(dur)
	}

	return fra
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

func ceiling(t time.Time, d time.Duration) time.Time {
	f := t.Truncate(d)

	if f.Equal(t) {
		return t
	}

	return f.Add(d)
}
