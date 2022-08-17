package framer

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_Framer_Next_Hour(t *testing.T) {
	testCases := []struct {
		sta time.Time
		end time.Time
		cou int
		las bool
		fra Frame
	}{
		// Case 0
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			cou: 1,
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
			cou: 4,
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
			cou: 24,
			las: true,
			fra: Frame{
				Sta: time.Date(2022, time.March, 26, 23, 0, 0, 0, time.UTC),
				End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			},
		},
		// Case 3
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			cou: 25,
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
					Dur: time.Hour,
				})
			}

			var fra Frame
			for i := 0; i < tc.cou; i++ {
				fra = f.Next()
			}

			var las bool
			{
				las = f.Last()
			}

			if !reflect.DeepEqual(tc.fra, fra) {
				t.Fatalf("fra\n\n%s\n", cmp.Diff(tc.fra, fra))
			}
			if !reflect.DeepEqual(tc.las, las) {
				t.Fatalf("las\n\n%s\n", cmp.Diff(tc.las, las))
			}
		})
	}
}

func Test_Framer_Next_10_Milliseconds(t *testing.T) {
	testCases := []struct {
		sta time.Time
		end time.Time
		cou int
		las bool
		fra Frame
	}{
		// Case 0
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 0, 0, 0, 1e8, time.UTC),
			cou: 1,
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
			cou: 7,
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
			cou: 10,
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
			cou: 11,
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
					Dur: 10 * time.Millisecond,
				})
			}

			var fra Frame
			for i := 0; i < tc.cou; i++ {
				fra = f.Next()
			}

			var las bool
			{
				las = f.Last()
			}

			if !reflect.DeepEqual(tc.fra, fra) {
				t.Fatalf("fra\n\n%s\n", cmp.Diff(tc.fra, fra))
			}
			if !reflect.DeepEqual(tc.las, las) {
				t.Fatalf("las\n\n%s\n", cmp.Diff(tc.las, las))
			}
		})
	}
}
