package framer

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_Framer_List_1h(t *testing.T) {
	testCases := []struct {
		sta time.Time
		end time.Time
		fra Frames
		tic time.Duration
	}{
		// Case 0
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 1, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 1, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 2, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 2, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 3, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 3, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 4, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 4, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 5, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 5, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 6, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 6, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 7, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 7, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 8, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 8, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 9, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 9, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 10, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 10, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 11, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 11, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 12, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 12, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 13, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 13, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 14, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 14, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 15, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 15, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 16, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 16, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 17, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 17, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 18, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 18, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 19, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 19, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 20, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 20, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 21, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 21, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 23, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 23, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		// Case 1
		{
			sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 23, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 23, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		// Case 2
		{
			sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 1, 0, 0, 0, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 23, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 23, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 1, 0, 0, 0, time.UTC),
				},
			},
		},
		// Case 3
		{
			sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			tic: 15 * time.Minute,
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 23, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 22, 15, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 23, 15, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 22, 30, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 23, 30, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 22, 45, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 23, 45, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 23, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		// Case 4
		{
			sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			tic: 30 * time.Minute,
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 23, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 22, 30, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 23, 30, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 23, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
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
					Len: time.Hour,
					Tic: tc.tic,
				})
			}

			var fra Frames
			{
				fra = f.List()
			}

			if !reflect.DeepEqual(tc.fra, fra) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.fra, fra))
			}
		})
	}
}

func Test_Framer_List_10ms(t *testing.T) {
	testCases := []struct {
		sta time.Time
		end time.Time
		fra Frames
		tic time.Duration
	}{
		// Case 0
		{
			sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 0, 0, 0, 1e8, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 0, 0, 0, 1e7, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 1e7, time.UTC),
					End: time.Date(2022, time.March, 26, 0, 0, 0, 2e7, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 2e7, time.UTC),
					End: time.Date(2022, time.March, 26, 0, 0, 0, 3e7, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 3e7, time.UTC),
					End: time.Date(2022, time.March, 26, 0, 0, 0, 4e7, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 4e7, time.UTC),
					End: time.Date(2022, time.March, 26, 0, 0, 0, 5e7, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 5e7, time.UTC),
					End: time.Date(2022, time.March, 26, 0, 0, 0, 6e7, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 6e7, time.UTC),
					End: time.Date(2022, time.March, 26, 0, 0, 0, 7e7, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 7e7, time.UTC),
					End: time.Date(2022, time.March, 26, 0, 0, 0, 8e7, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 8e7, time.UTC),
					End: time.Date(2022, time.March, 26, 0, 0, 0, 9e7, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 9e7, time.UTC),
					End: time.Date(2022, time.March, 26, 0, 0, 0, 1e8, time.UTC),
				},
			},
		},
		// Case 1
		{
			sta: time.Date(2022, time.March, 27, 4, 12, 36, 35e6, time.UTC),
			end: time.Date(2022, time.March, 27, 4, 12, 36, 85e6, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 27, 4, 12, 36, 35e6, time.UTC),
					End: time.Date(2022, time.March, 27, 4, 12, 36, 45e6, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 27, 4, 12, 36, 45e6, time.UTC),
					End: time.Date(2022, time.March, 27, 4, 12, 36, 55e6, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 27, 4, 12, 36, 55e6, time.UTC),
					End: time.Date(2022, time.March, 27, 4, 12, 36, 65e6, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 27, 4, 12, 36, 65e6, time.UTC),
					End: time.Date(2022, time.March, 27, 4, 12, 36, 75e6, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 27, 4, 12, 36, 75e6, time.UTC),
					End: time.Date(2022, time.March, 27, 4, 12, 36, 85e6, time.UTC),
				},
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

			var fra Frames
			{
				fra = f.List()
			}

			if !reflect.DeepEqual(tc.fra, fra) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.fra, fra))
			}
		})
	}
}
