package framer

import (
	"time"
)

type Frames []Frame

func (f Frames) Contains(fra Frame) bool {
	for _, v := range f {
		if v.Sta.Equal(fra.Sta) && v.End.Equal(fra.End) {
			return true
		}
	}

	return false
}

func (f Frames) Hour() Frames {
	max := ceiling(f.Max().End, time.Hour)
	min := f.Min().Sta.Truncate(time.Hour)

	var fra []Frame

	s := min
	e := min.Add(time.Hour)

	for {
		fra = append(fra, Frame{Sta: s, End: e})

		if e.Equal(max) {
			break
		}

		s = s.Add(time.Hour)
		e = e.Add(time.Hour)
	}

	return fra
}

func (f Frames) Max() Frame {
	var max Frame

	for _, v := range f {
		if v.End.After(max.End) {
			max = v
		}
	}

	return max
}

func (f Frames) Min() Frame {
	min := f[0]

	for _, v := range f {
		if v.Sta.Before(min.Sta) {
			min = v
		}
	}

	return min
}

func (f Frames) Remove(rem Frames) Frames {
	var cle Frames

	for _, v := range f {
		if !rem.Contains(v) {
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
