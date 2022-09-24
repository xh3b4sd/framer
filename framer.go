package framer

import (
	"time"
)

type Framer struct {
	lef time.Time
	sta time.Time
	end time.Time
	rig time.Time
	len time.Duration
	tic time.Duration
}

func New(con Config) *Framer {
	{
		con = con.Ensure()
	}

	{
		con.Verify()
	}

	return &Framer{
		lef: con.Sta.UTC(),
		sta: con.Sta.UTC().Add(-con.Tic),
		end: con.Sta.UTC().Add(+con.Len).Add(-con.Tic),
		rig: con.End.UTC(),
		len: con.Len,
		tic: con.Tic,
	}
}

func (f *Framer) Conf() Config {
	return Config{
		Sta: f.sta,
		End: f.end,
		Len: f.len,
		Tic: f.tic,
	}
}

func (f *Framer) Firs() bool {
	return !f.sta.After(f.lef)
}

func (f *Framer) Last() bool {
	return !f.end.Before(f.rig)
}

func (f *Framer) List() Frames {
	var fra Frames

	for !f.Last() {
		fra = append(fra, f.Next())
	}

	return fra
}

func (f *Framer) Next() Frame {
	if f.Last() {
		return Frame{}
	}

	{
		f.sta = f.sta.Add(f.tic)
		f.end = f.end.Add(f.tic)
	}

	var sta time.Time
	var end time.Time
	{
		sta = f.sta
		end = f.end
	}

	if end.After(f.end) {
		end = f.end
	}

	return Frame{
		Sta: sta,
		End: end,
	}
}

func (f *Framer) Prev() Frame {
	if f.Firs() {
		return Frame{}
	}

	{
		f.sta = f.sta.Add(-f.tic)
		f.end = f.end.Add(-f.tic)
	}

	var sta time.Time
	var end time.Time
	{
		sta = f.sta
		end = f.end
	}

	if sta.Before(f.sta) {
		sta = f.sta
	}

	return Frame{
		Sta: sta,
		End: end,
	}
}
