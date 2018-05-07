package checker

import (
	"sort"
	"testing"
)

func TestGetAll_Single_NoName(t *testing.T) {
	tree := MaxTree{
		Name: "",
	}

	actual := tree.GetAll("Foo")

	tobe := []string{
		"Foo",
	}

	checkSlice(actual, tobe, t)
}

func TestGetAll_Single_HasName(t *testing.T) {
	tree := MaxTree{
		Name: "A",
	}

	actual := tree.GetAll("Foo")

	tobe := []string{
		"Foo A",
	}

	checkSlice(actual, tobe, t)
}

func TestGetAll_OneChild_NoName(t *testing.T) {
	tree := MaxTree{
		Name: "",
		Childs: []MaxTree{
			MaxTree{
				Name: "1",
			},
		},
	}

	actual := tree.GetAll("Foo")

	tobe := []string{
		"Foo",
		"Foo 1",
	}

	checkSlice(actual, tobe, t)
}

//// Util ////

func checkSlice(actual, tobe []string, t *testing.T) {
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
			t.Fatalf("Mismatch result: actual: %q, tobe: %q", a, tb)
		}
	}
}
