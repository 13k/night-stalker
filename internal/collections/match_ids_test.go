package collections_test

import (
	"testing"

	nscol "github.com/13k/night-stalker/internal/collections"
	nspb "github.com/13k/night-stalker/internal/protocol"
)

func TestNewMatchIDsFromString(t *testing.T) {
	testCases := []struct {
		String   string
		Sep      string
		Expected nscol.MatchIDs
		Err      string
	}{
		{
			String:   "",
			Sep:      "",
			Expected: nil,
			Err:      "",
		},
		{
			String:   "",
			Sep:      ",",
			Expected: nil,
			Err:      "",
		},
		{
			String:   "1,2,3",
			Sep:      ",",
			Expected: nscol.MatchIDs{1, 2, 3},
			Err:      "",
		},
		{
			String:   "123",
			Sep:      "",
			Expected: nscol.MatchIDs{1, 2, 3},
			Err:      "",
		},
		{
			String:   "1 2 3",
			Sep:      "",
			Expected: nil,
			Err:      `strconv.ParseUint: parsing " ": invalid syntax`,
		},
		{
			String:   "1 2 3",
			Sep:      " ",
			Expected: nscol.MatchIDs{1, 2, 3},
			Err:      "",
		},
		{
			String:   "abc",
			Sep:      "",
			Expected: nil,
			Err:      `strconv.ParseUint: parsing "a": invalid syntax`,
		},
		{
			String:   "-1,2,3",
			Sep:      ",",
			Expected: nil,
			Err:      `strconv.ParseUint: parsing "-1": invalid syntax`,
		},
	}

	for testCaseIdx, testCase := range testCases {
		actual, err := nscol.NewMatchIDsFromString(testCase.String, testCase.Sep)

		if testCase.Err != "" {
			if err == nil {
				t.Fatalf("case %d: expected error %q", testCaseIdx, testCase.Err)
			}

			if err.Error() != testCase.Err {
				t.Fatalf("case %d: expected error %q, got %q", testCaseIdx, testCase.Err, err)
			}
		} else {
			if err != nil {
				t.Fatalf("case %d: expected no error, got %q", testCaseIdx, err)
			}

			if testCase.Expected == nil {
				if actual != nil {
					t.Fatalf("case %d: expected nil", testCaseIdx)
				}
			} else {
				if actual == nil {
					t.Fatalf("case %d: expected non-nil", testCaseIdx)
				}
			}

			expectedLen := len(testCase.Expected)
			actualLen := len(actual)

			if actualLen != expectedLen {
				t.Fatalf("case %d: expected len %d, got %d", testCaseIdx, expectedLen, actualLen)
			}

			for i, expectedID := range testCase.Expected {
				actualID := actual[i]

				if actualID != expectedID {
					t.Fatalf("case %d: index %d: expected %d, got %d", testCaseIdx, i, expectedID, actualID)
				}
			}
		}
	}
}

func TestMatchIDs_ToInterfaces(t *testing.T) {
	testCases := []struct {
		Subject  nscol.MatchIDs
		Expected []interface{}
	}{
		{
			Subject:  nil,
			Expected: nil,
		},
		{
			Subject:  nscol.MatchIDs{},
			Expected: nil,
		},
		{
			Subject: nscol.MatchIDs{1, 2, 3},
			Expected: []interface{}{
				nspb.MatchID(1),
				nspb.MatchID(2),
				nspb.MatchID(3),
			},
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := testCase.Subject
		actual := subject.ToInterfaces()

		if testCase.Expected == nil {
			if actual != nil {
				t.Fatalf("case %d: expected nil", testCaseIdx)
			}
		} else {
			if actual == nil {
				t.Fatalf("case %d: expected non-nil", testCaseIdx)
			}
		}

		expectedLen := len(testCase.Expected)
		actualLen := len(actual)

		if actualLen != expectedLen {
			t.Fatalf("case %d: expected len %d, got %d", testCaseIdx, expectedLen, actualLen)
		}

		for i, expectedID := range testCase.Expected {
			actualID := actual[i]

			if actualID != expectedID {
				t.Fatalf(
					"case %d: index %d: expected %#+[2]v [%[2]T], got %#+[3]v [%[3]T]",
					testCaseIdx,
					i,
					expectedID,
					actualID,
				)
			}
		}
	}
}

func TestMatchIDs_Join(t *testing.T) {
	testCases := []struct {
		Subject  nscol.MatchIDs
		Sep      string
		Expected string
	}{
		{
			Subject:  nil,
			Sep:      "",
			Expected: "",
		},
		{
			Subject:  nil,
			Sep:      "",
			Expected: "",
		},
		{
			Subject:  nscol.MatchIDs{},
			Sep:      "",
			Expected: "",
		},
		{
			Subject:  nscol.MatchIDs{},
			Sep:      "",
			Expected: "",
		},
		{
			Subject:  nscol.MatchIDs{},
			Sep:      ",",
			Expected: "",
		},
		{
			Subject:  nscol.MatchIDs{1},
			Sep:      "",
			Expected: "1",
		},
		{
			Subject:  nscol.MatchIDs{1, 1, 1},
			Sep:      "",
			Expected: "111",
		},
		{
			Subject:  nscol.MatchIDs{1, 1, 2, 2, 3, 3},
			Sep:      ", ",
			Expected: "1, 1, 2, 2, 3, 3",
		},
		{
			Subject:  nscol.MatchIDs{1, 1, 3, 3, 5, 5},
			Sep:      "",
			Expected: "113355",
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := testCase.Subject
		expected := testCase.Expected
		actual := subject.Join(testCase.Sep)

		if actual != expected {
			t.Fatalf("case %d: expected %q, got %q", testCaseIdx, expected, actual)
		}
	}
}

func TestMatchIDs_AddUnique(t *testing.T) {
	testCases := []struct {
		Subject  nscol.MatchIDs
		Add      nscol.MatchIDs
		Expected nscol.MatchIDs
	}{
		{
			Subject:  nil,
			Add:      nil,
			Expected: nil,
		},
		{
			Subject:  nil,
			Add:      nscol.MatchIDs{},
			Expected: nil,
		},
		{
			Subject:  nscol.MatchIDs{},
			Add:      nil,
			Expected: nscol.MatchIDs{},
		},
		{
			Subject:  nscol.MatchIDs{},
			Add:      nscol.MatchIDs{},
			Expected: nscol.MatchIDs{},
		},
		{
			Subject:  nscol.MatchIDs{},
			Add:      nscol.MatchIDs{1},
			Expected: nscol.MatchIDs{1},
		},
		{
			Subject:  nscol.MatchIDs{1},
			Add:      nscol.MatchIDs{1},
			Expected: nscol.MatchIDs{1},
		},
		{
			Subject:  nscol.MatchIDs{1, 1, 1},
			Add:      nscol.MatchIDs{1, 1, 1},
			Expected: nscol.MatchIDs{1, 1, 1},
		},
		{
			Subject:  nscol.MatchIDs{1, 1, 2, 2, 3, 3},
			Add:      nscol.MatchIDs{1, 2, 3},
			Expected: nscol.MatchIDs{1, 1, 2, 2, 3, 3},
		},
		{
			Subject:  nscol.MatchIDs{1, 1, 3, 3, 5, 5},
			Add:      nscol.MatchIDs{1, 2, 3, 4, 5, 6},
			Expected: nscol.MatchIDs{1, 1, 3, 3, 5, 5, 2, 4, 6},
		},
	}

	for testCaseIdx, testCase := range testCases {
		subject := testCase.Subject
		actual := subject.AddUnique(testCase.Add...)

		if testCase.Expected == nil {
			if actual != nil {
				t.Fatalf("case %d: expected nil", testCaseIdx)
			}
		} else {
			if actual == nil {
				t.Fatalf("case %d: expected non-nil", testCaseIdx)
			}
		}

		expectedLen := len(testCase.Expected)
		actualLen := len(actual)

		if actualLen != expectedLen {
			t.Fatalf("case %d: expected len %d, got %d", testCaseIdx, expectedLen, actualLen)
		}

		for i, expectedID := range testCase.Expected {
			actualID := actual[i]

			if actualID != expectedID {
				t.Fatalf("case %d: index %d: expected %d, got %d", testCaseIdx, i, expectedID, actualID)
			}
		}
	}
}
