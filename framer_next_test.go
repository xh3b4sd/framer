package framer

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_Framer_Next_1h(t *testing.T) {
	testCases := []struct {
		sta time.Time
		end time.Time
		tic time.Duration
		nex int
		pre int
		fir bool
		las bool
		fra Frame
	}{
		// Case 0
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			nex: 1,
			fir: true,
			las: false,
			fra: Frame{
				Sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
				End: time.Date(2022, time.March, 26, 1, 0, 0, 0, time.UTC),
			},
		},
		// Case 1
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			nex: 4,
			las: false,
			fra: Frame{
				Sta: time.Date(2022, time.March, 26, 3, 0, 0, 0, time.UTC),
				End: time.Date(2022, time.March, 26, 4, 0, 0, 0, time.UTC),
			},
		},
		// Case 2
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			nex: 4,
			pre: 1,
			las: false,
			fra: Frame{
				Sta: time.Date(2022, time.March, 26, 2, 0, 0, 0, time.UTC),
				End: time.Date(2022, time.March, 26, 3, 0, 0, 0, time.UTC),
			},
		},
		// Case 3
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			nex: 24,
			las: true,
			fra: Frame{
				Sta: time.Date(2022, time.March, 26, 23, 0, 0, 0, time.UTC),
				End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			},
		},
		// Case 4
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			nex: 24,
			pre: 4,
			las: false,
			fra: Frame{
				Sta: time.Date(2022, time.March, 26, 19, 0, 0, 0, time.UTC),
				End: time.Date(2022, time.March, 26, 20, 0, 0, 0, time.UTC),
			},
		},
		// Case 5
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			nex: 25,
			las: true,
			fra: Frame{},
		},
		// Case 6
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			tic: 10 * time.Minute,
			nex: 1,
			fir: true,
			las: false,
			fra: Frame{
				Sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
				End: time.Date(2022, time.March, 26, 1, 0, 0, 0, time.UTC),
			},
		},
		// Case 7
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			tic: 10 * time.Minute,
			nex: 7,
			las: false,
			fra: Frame{
				Sta: time.Date(2022, time.March, 26, 1, 0, 0, 0, time.UTC),
				End: time.Date(2022, time.March, 26, 2, 0, 0, 0, time.UTC),
			},
		},
		// Case 8
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			tic: 10 * time.Minute,
			nex: 8,
			las: false,
			fra: Frame{
				Sta: time.Date(2022, time.March, 26, 1, 10, 0, 0, time.UTC),
				End: time.Date(2022, time.March, 26, 2, 10, 0, 0, time.UTC),
			},
		},
		// Case 9
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			tic: 10 * time.Minute,
			nex: 138,
			las: false,
			fra: Frame{
				Sta: time.Date(2022, time.March, 26, 22, 50, 0, 0, time.UTC),
				End: time.Date(2022, time.March, 26, 23, 50, 0, 0, time.UTC),
			},
		},
		// Case 10
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			tic: 10 * time.Minute,
			nex: 139,
			las: true,
			fra: Frame{
				Sta: time.Date(2022, time.March, 26, 23, 0, 0, 0, time.UTC),
				End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			},
		},
		// Case 11
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			tic: 10 * time.Minute,
			nex: 140,
			las: true,
			fra: Frame{},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var f *Framer
			{
				f = New(Config{
					Sta: tc.sta,
					End: tc.end,
					Len: time.Hour,
					Tic: tc.tic,
				})
			}

			var fra Frame
			for i := 0; i < tc.nex; i++ {
				fra = f.Next()
			}
			for i := 0; i < tc.pre; i++ {
				fra = f.Prev()
			}

			// We want to call f.Curr() multiple times in order to make sure the
			// value remains the same. We cannot inline the calls because the
			// golinter and/or golangci complains and no ignore directive seemed
			// to work.
			{
				one := f.Curr()
				two := f.Curr()
				thr := f.Curr()
				fou := f.Curr()
				if !tc.las && (fra != one || fra != two || thr != fou) {
					t.Fatal("current frame must never change")
				}
			}

			var fir bool
			{
				fir = f.Firs()
			}

			var las bool
			{
				las = f.Last()
			}

			if !reflect.DeepEqual(tc.fra, fra) {
				t.Fatalf("fra\n\n%s\n", cmp.Diff(tc.fra, fra))
			}
			if !reflect.DeepEqual(tc.fir, fir) {
				t.Fatalf("fir\n\n%s\n", cmp.Diff(tc.fir, fir))
			}
			if !reflect.DeepEqual(tc.las, las) {
				t.Fatalf("las\n\n%s\n", cmp.Diff(tc.las, las))
			}
		})
	}
}

func Test_Framer_Next_10ms(t *testing.T) {
	testCases := []struct {
		sta time.Time
		end time.Time
		tic time.Duration
		nex int
		fir bool
		las bool
		fra Frame
	}{
		// Case 0
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 0, 0, 0, 1e8, time.UTC),
			nex: 1,
			fir: true,
			las: false,
			fra: Frame{
				Sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
				End: time.Date(2022, time.March, 26, 0, 0, 0, 1e7, time.UTC),
			},
		},
		// Case 1
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 0, 0, 0, 1e8, time.UTC),
			nex: 7,
			las: false,
			fra: Frame{
				Sta: time.Date(2022, time.March, 26, 0, 0, 0, 6e7, time.UTC),
				End: time.Date(2022, time.March, 26, 0, 0, 0, 7e7, time.UTC),
			},
		},
		// Case 2
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 0, 0, 0, 1e8, time.UTC),
			nex: 10,
			las: true,
			fra: Frame{
				Sta: time.Date(2022, time.March, 26, 0, 0, 0, 9e7, time.UTC),
				End: time.Date(2022, time.March, 26, 0, 0, 0, 1e8, time.UTC),
			},
		},
		// Case 3
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 0, 0, 0, 1e8, time.UTC),
			nex: 11,
			las: true,
			fra: Frame{},
		},
		// Case 4
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 0, 0, 0, 1e8, time.UTC),
			tic: 2 * time.Millisecond,
			nex: 1,
			fir: true,
			las: false,
			fra: Frame{
				Sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
				End: time.Date(2022, time.March, 26, 0, 0, 0, 1e7, time.UTC),
			},
		},
		// Case 5
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 0, 0, 0, 1e8, time.UTC),
			tic: 2 * time.Millisecond,
			nex: 2,
			las: false,
			fra: Frame{
				Sta: time.Date(2022, time.March, 26, 0, 0, 0, 0+2e6, time.UTC),
				End: time.Date(2022, time.March, 26, 0, 0, 0, 1e7+2e6, time.UTC),
			},
		},
		// Case 6
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 0, 0, 0, 1e8, time.UTC),
			tic: 2 * time.Millisecond,
			nex: 6,
			las: false,
			fra: Frame{
				Sta: time.Date(2022, time.March, 26, 0, 0, 0, 1e7, time.UTC),
				End: time.Date(2022, time.March, 26, 0, 0, 0, 2e7, time.UTC),
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var f *Framer
			{
				f = New(Config{
					Sta: tc.sta,
					End: tc.end,
					Len: 10 * time.Millisecond,
					Tic: tc.tic,
				})
			}

			var fra Frame
			for i := 0; i < tc.nex; i++ {
				fra = f.Next()
			}

			// We want to call f.Curr() multiple times in order to make sure the
			// value remains the same. We cannot inline the calls because the
			// golinter and/or golangci complains and no ignore directive seemed
			// to work.
			{
				one := f.Curr()
				two := f.Curr()
				thr := f.Curr()
				fou := f.Curr()
				if !tc.las && (fra != one || fra != two || thr != fou) {
					t.Fatal("current frame must never change")
				}
			}

			var fir bool
			{
				fir = f.Firs()
			}

			var las bool
			{
				las = f.Last()
			}

			if !reflect.DeepEqual(tc.fra, fra) {
				t.Fatalf("fra\n\n%s\n", cmp.Diff(tc.fra, fra))
			}
			if !reflect.DeepEqual(tc.fir, fir) {
				t.Fatalf("fir\n\n%s\n", cmp.Diff(tc.fir, fir))
			}
			if !reflect.DeepEqual(tc.las, las) {
				t.Fatalf("las\n\n%s\n", cmp.Diff(tc.las, las))
			}
		})
	}
}
