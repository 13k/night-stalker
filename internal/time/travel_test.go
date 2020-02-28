package time_test

import (
	"testing"
	"time"

	nstime "github.com/13k/night-stalker/internal/time"
)

func TestTravel_Ago(t *testing.T) {
	testCases := []struct {
		From     time.Time
		Args     [3]int
		Expected time.Time
	}{
		{
			From:     time.Date(2000, 01, 02, 03, 04, 05, 0, time.UTC),
			Args:     [3]int{1, 0, 0},
			Expected: time.Date(1999, 01, 02, 03, 04, 05, 0, time.UTC),
		},
		{
			From:     time.Date(2000, 01, 02, 03, 04, 05, 0, time.UTC),
			Args:     [3]int{-1, 0, 0},
			Expected: time.Date(1999, 01, 02, 03, 04, 05, 0, time.UTC),
		},
		{
			From:     time.Date(2000, 01, 02, 03, 04, 05, 0, time.UTC),
			Args:     [3]int{1, 1, 0},
			Expected: time.Date(1998, 12, 02, 03, 04, 05, 0, time.UTC),
		},
		{
			From:     time.Date(2000, 01, 02, 03, 04, 05, 0, time.UTC),
			Args:     [3]int{-1, -1, 0},
			Expected: time.Date(1998, 12, 02, 03, 04, 05, 0, time.UTC),
		},
		{
			From:     time.Date(2000, 01, 02, 03, 04, 05, 0, time.UTC),
			Args:     [3]int{1, 1, 1},
			Expected: time.Date(1998, 12, 01, 03, 04, 05, 0, time.UTC),
		},
		{
			From:     time.Date(2000, 01, 02, 03, 04, 05, 0, time.UTC),
			Args:     [3]int{-1, -1, -1},
			Expected: time.Date(1998, 12, 01, 03, 04, 05, 0, time.UTC),
		},
	}

	for testCaseIdx, testCase := range testCases {
		args := testCase.Args
		subject := nstime.Travel{testCase.From}
		actual := subject.Ago(args[0], args[1], args[2])

		if !actual.Equal(testCase.Expected) {
			t.Fatalf("case %d: expected %v, got %v", testCaseIdx, testCase.Expected, actual)
		}
	}
}

func TestTravel_YearsAgo(t *testing.T) {
	testCases := []struct {
		From     time.Time
		Years    int
		Expected time.Time
	}{
		{
			From:     time.Date(2000, 01, 02, 03, 04, 05, 0, time.UTC),
			Years:    1,
			Expected: time.Date(1999, 01, 02, 03, 04, 05, 0, time.UTC),
		},
		{
			From:     time.Date(2000, 01, 02, 03, 04, 05, 0, time.UTC),
			Years:    -1,
			Expected: time.Date(1999, 01, 02, 03, 04, 05, 0, time.UTC),
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := nstime.Travel{testCase.From}
		actual := subject.YearsAgo(testCase.Years)

		if !actual.Equal(testCase.Expected) {
			t.Fatalf("case %d: expected %v, got %v", testCaseIdx, testCase.Expected, actual)
		}
	}
}

func TestTravel_MonthsAgo(t *testing.T) {
	testCases := []struct {
		From     time.Time
		Months   int
		Expected time.Time
	}{
		{
			From:     time.Date(2000, 01, 02, 03, 04, 05, 0, time.UTC),
			Months:   1,
			Expected: time.Date(1999, 12, 02, 03, 04, 05, 0, time.UTC),
		},
		{
			From:     time.Date(2000, 01, 02, 03, 04, 05, 0, time.UTC),
			Months:   -1,
			Expected: time.Date(1999, 12, 02, 03, 04, 05, 0, time.UTC),
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := nstime.Travel{testCase.From}
		actual := subject.MonthsAgo(testCase.Months)

		if !actual.Equal(testCase.Expected) {
			t.Fatalf("case %d: expected %v, got %v", testCaseIdx, testCase.Expected, actual)
		}
	}
}

func TestTravel_DaysAgo(t *testing.T) {
	testCases := []struct {
		From     time.Time
		Days     int
		Expected time.Time
	}{
		{
			From:     time.Date(2000, 01, 02, 03, 04, 05, 0, time.UTC),
			Days:     3,
			Expected: time.Date(1999, 12, 30, 03, 04, 05, 0, time.UTC),
		},
		{
			From:     time.Date(2000, 01, 02, 03, 04, 05, 0, time.UTC),
			Days:     -3,
			Expected: time.Date(1999, 12, 30, 03, 04, 05, 0, time.UTC),
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := nstime.Travel{testCase.From}
		actual := subject.DaysAgo(testCase.Days)

		if !actual.Equal(testCase.Expected) {
			t.Fatalf("case %d: expected %v, got %v", testCaseIdx, testCase.Expected, actual)
		}
	}
}

func TestTravel_BeginningOfDay(t *testing.T) {
	testCases := []struct {
		From     time.Time
		Expected time.Time
	}{
		{
			From:     time.Date(2000, 01, 02, 00, 00, 00, 0, time.UTC),
			Expected: time.Date(2000, 01, 02, 00, 00, 00, 0, time.UTC),
		},
		{
			From:     time.Date(2000, 01, 02, 03, 04, 05, 0, time.UTC),
			Expected: time.Date(2000, 01, 02, 00, 00, 00, 0, time.UTC),
		},
		{
			From:     time.Date(2000, 01, 02, 23, 59, 59, 0, time.UTC),
			Expected: time.Date(2000, 01, 02, 00, 00, 00, 0, time.UTC),
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := nstime.Travel{testCase.From}
		actual := subject.BeginningOfDay()

		if !actual.Equal(testCase.Expected) {
			t.Fatalf("case %d: expected %v, got %v", testCaseIdx, testCase.Expected, actual)
		}
	}
}
