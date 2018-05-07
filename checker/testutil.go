package checker

import (
	"sort"
	"strings"
	"testing"
)

//// Util ////

func checkSlice(t *testing.T, actual, tobe []string) {
	sort.Strings(actual)
	sort.Strings(tobe)

	if len(actual) < len(tobe) {
		t.Fatalf("Too short results: %d < %d", len(actual), len(tobe))
	}

	if len(actual) > len(tobe) {
		t.Fatalf("Too long results: %d < %d", len(actual), len(tobe))
	}

	for i := 0; i < len(actual); i++ {
		a := actual[i]
		tb := tobe[i]

		if a != tb {
			t.Fatalf("Mismatch result:\n    actual: %+v\n    tobe: %+v", quoteSlice(actual), quoteSlice(tobe))
		}
	}
}

func quoteSlice(strs []string) string {
	return "[\"" + strings.Join(strs, "\", \"") + "\"]"
}
