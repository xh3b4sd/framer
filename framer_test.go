package framer

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_Framer_Day(t *testing.T) {
	testCases := []struct {
		sta time.Time
		end time.Time
		fra Frames
	}{
		// Case 0
		{
			sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		// Case 1
		{
			sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 1, 0, 0, 0, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 28, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		// Case 2
		{
			sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 16, 0, 0, 0, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 28, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		// Case 3
		{
			sta: time.Date(2022, time.March, 26, 22, 16, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 23, 0, 0, 0, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		// Case 4
		{
			sta: time.Date(2022, time.March, 26, 22, 16, 3, 22, time.UTC),
			end: time.Date(2022, time.March, 26, 22, 21, 7, 37, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		// Case 5
		{
			sta: time.Date(2022, time.March, 26, 16, 3, 22, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 21, 7, 37, 0, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		// Case 6
		{
			sta: time.Date(2022, time.March, 24, 10, 7, 18, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 21, 7, 37, 0, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 24, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 25, 0, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 25, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		// Case 7
		{
			sta: time.Date(2022, time.March, 24, 23, 59, 59, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 23, 59, 59, 0, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 24, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 25, 0, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 25, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		// Case 8
		{
			sta: time.Date(2022, time.March, 24, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 23, 59, 59, 0, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 24, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 25, 0, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 25, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		// Case 9
		{
			sta: time.Date(2022, time.March, 26, 23, 59, 59, 0, time.UTC),
			end: time.Date(2022, time.March, 24, 0, 0, 0, 0, time.UTC),
			fra: nil,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var f Interface
			{
				c := Config{
					Sta: tc.sta,
					End: tc.end,
				}

				f, err = New(c)
				if err != nil {
					t.Fatal(err)
				}
			}

			var fra Frames
			{
				fra = f.Day()
			}

			if !reflect.DeepEqual(tc.fra, fra) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.fra, fra))
			}
		})
	}
}

func Test_Framer_Exa(t *testing.T) {
	testCases := []struct {
		sta time.Time
		end time.Time
		fra Frames
	}{
		// Case 0
		{
			sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		// Case 1
		{
			sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 1, 0, 0, 0, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 1, 0, 0, 0, time.UTC),
				},
			},
		},
		// Case 2
		{
			sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 16, 0, 0, 0, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 16, 0, 0, 0, time.UTC),
				},
			},
		},
		// Case 3
		{
			sta: time.Date(2022, time.March, 26, 22, 16, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 23, 0, 0, 0, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 26, 22, 16, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 23, 0, 0, 0, time.UTC),
				},
			},
		},
		// Case 4
		{
			sta: time.Date(2022, time.March, 26, 22, 16, 3, 22, time.UTC),
			end: time.Date(2022, time.March, 26, 22, 21, 7, 37, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 26, 22, 16, 3, 22, time.UTC),
					End: time.Date(2022, time.March, 26, 22, 21, 7, 37, time.UTC),
				},
			},
		},
		// Case 5
		{
			sta: time.Date(2022, time.March, 26, 16, 3, 22, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 21, 7, 37, 0, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 26, 16, 3, 22, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 21, 7, 37, 0, time.UTC),
				},
			},
		},
		// Case 6
		{
			sta: time.Date(2022, time.March, 24, 10, 7, 18, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 21, 7, 37, 0, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 24, 10, 7, 18, 0, time.UTC),
					End: time.Date(2022, time.March, 25, 0, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 25, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 21, 7, 37, 0, time.UTC),
				},
			},
		},
		// Case 7
		{
			sta: time.Date(2022, time.March, 24, 23, 59, 59, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 23, 59, 59, 0, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 24, 23, 59, 59, 0, time.UTC),
					End: time.Date(2022, time.March, 25, 0, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 25, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 23, 59, 59, 0, time.UTC),
				},
			},
		},
		// Case 8
		{
			sta: time.Date(2022, time.March, 24, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 23, 59, 59, 0, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 24, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 25, 0, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 25, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 23, 59, 59, 0, time.UTC),
				},
			},
		},
		// Case 9
		{
			sta: time.Date(2022, time.March, 26, 23, 59, 59, 0, time.UTC),
			end: time.Date(2022, time.March, 24, 0, 0, 0, 0, time.UTC),
			fra: nil,
		},
		// Case 10
		{
			sta: time.Date(2022, time.March, 26, 23, 59, 59, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 26, 23, 59, 59, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		// Case 11
		{
			sta: time.Date(2022, time.March, 26, 23, 59, 59, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 16, 0, 0, 0, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 26, 23, 59, 59, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
				{
					Sta: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 16, 0, 0, 0, time.UTC),
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var f Interface
			{
				c := Config{
					Sta: tc.sta,
					End: tc.end,
				}

				f, err = New(c)
				if err != nil {
					t.Fatal(err)
				}
			}

			var fra Frames
			{
				fra = f.Exa()
			}

			if !reflect.DeepEqual(tc.fra, fra) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.fra, fra))
			}
		})
	}
}

func Test_Framer_Hour(t *testing.T) {
	testCases := []struct {
		sta time.Time
		end time.Time
		fra Frames
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
			sta: time.Date(2022, time.March, 26, 22, 16, 3, 22, time.UTC),
			end: time.Date(2022, time.March, 26, 22, 21, 7, 37, time.UTC),
			fra: Frames{
				{
					Sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 23, 0, 0, 0, time.UTC),
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var f Interface
			{
				c := Config{
					Sta: tc.sta,
					End: tc.end,
				}

				f, err = New(c)
				if err != nil {
					t.Fatal(err)
				}
			}

			var fra Frames
			{
				fra = f.Exa().Hour()
			}

			if !reflect.DeepEqual(tc.fra, fra) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.fra, fra))
			}
		})
	}
}
