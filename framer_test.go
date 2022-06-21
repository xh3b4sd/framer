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
		fra []Frame
	}{
		// case 0
		{
			sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			fra: []Frame{
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
		// case 1
		{
			sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 1, 0, 0, 0, time.UTC),
			fra: []Frame{
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
		// case 2
		{
			sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 16, 0, 0, 0, time.UTC),
			fra: []Frame{
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
		// case 3
		{
			sta: time.Date(2022, time.March, 26, 22, 16, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 23, 0, 0, 0, time.UTC),
			fra: []Frame{
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		// case 4
		{
			sta: time.Date(2022, time.March, 26, 22, 16, 3, 22, time.UTC),
			end: time.Date(2022, time.March, 26, 22, 21, 7, 37, time.UTC),
			fra: []Frame{
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		// case 5
		{
			sta: time.Date(2022, time.March, 26, 16, 3, 22, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 21, 7, 37, 0, time.UTC),
			fra: []Frame{
				{
					Sta: time.Date(2022, time.March, 26, 0, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		// case 6
		{
			sta: time.Date(2022, time.March, 24, 10, 7, 18, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 21, 7, 37, 0, time.UTC),
			fra: []Frame{
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
		// case 7
		{
			sta: time.Date(2022, time.March, 24, 23, 59, 59, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 23, 59, 59, 0, time.UTC),
			fra: []Frame{
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
		// case 8
		{
			sta: time.Date(2022, time.March, 24, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 23, 59, 59, 0, time.UTC),
			fra: []Frame{
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
		// case 9
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

			var fra []Frame
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
		fra []Frame
	}{
		// case 0
		{
			sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			fra: []Frame{
				{
					Sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		// case 1
		{
			sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 1, 0, 0, 0, time.UTC),
			fra: []Frame{
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
		// case 2
		{
			sta: time.Date(2022, time.March, 26, 22, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 16, 0, 0, 0, time.UTC),
			fra: []Frame{
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
		// case 3
		{
			sta: time.Date(2022, time.March, 26, 22, 16, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 23, 0, 0, 0, time.UTC),
			fra: []Frame{
				{
					Sta: time.Date(2022, time.March, 26, 22, 16, 0, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 23, 0, 0, 0, time.UTC),
				},
			},
		},
		// case 4
		{
			sta: time.Date(2022, time.March, 26, 22, 16, 3, 22, time.UTC),
			end: time.Date(2022, time.March, 26, 22, 21, 7, 37, time.UTC),
			fra: []Frame{
				{
					Sta: time.Date(2022, time.March, 26, 22, 16, 3, 22, time.UTC),
					End: time.Date(2022, time.March, 26, 22, 21, 7, 37, time.UTC),
				},
			},
		},
		// case 5
		{
			sta: time.Date(2022, time.March, 26, 16, 3, 22, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 21, 7, 37, 0, time.UTC),
			fra: []Frame{
				{
					Sta: time.Date(2022, time.March, 26, 16, 3, 22, 0, time.UTC),
					End: time.Date(2022, time.March, 26, 21, 7, 37, 0, time.UTC),
				},
			},
		},
		// case 6
		{
			sta: time.Date(2022, time.March, 24, 10, 7, 18, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 21, 7, 37, 0, time.UTC),
			fra: []Frame{
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
		// case 7
		{
			sta: time.Date(2022, time.March, 24, 23, 59, 59, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 23, 59, 59, 0, time.UTC),
			fra: []Frame{
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
		// case 8
		{
			sta: time.Date(2022, time.March, 24, 0, 0, 0, 0, time.UTC),
			end: time.Date(2022, time.March, 26, 23, 59, 59, 0, time.UTC),
			fra: []Frame{
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
		// case 9
		{
			sta: time.Date(2022, time.March, 26, 23, 59, 59, 0, time.UTC),
			end: time.Date(2022, time.March, 24, 0, 0, 0, 0, time.UTC),
			fra: nil,
		},
		// case 10
		{
			sta: time.Date(2022, time.March, 26, 23, 59, 59, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
			fra: []Frame{
				{
					Sta: time.Date(2022, time.March, 26, 23, 59, 59, 0, time.UTC),
					End: time.Date(2022, time.March, 27, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		// case 11
		{
			sta: time.Date(2022, time.March, 26, 23, 59, 59, 0, time.UTC),
			end: time.Date(2022, time.March, 27, 16, 0, 0, 0, time.UTC),
			fra: []Frame{
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

			var fra []Frame
			{
				fra = f.Exa()
			}

			if !reflect.DeepEqual(tc.fra, fra) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.fra, fra))
			}
		})
	}
}
