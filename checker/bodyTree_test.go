package checker

import (
	"testing"
)

//// GetAll ////

func TestGetAll_Single_NoName(t *testing.T) {
	tree := BodyTree{
		Name: "",
	}

	actual := tree.GetAll()

	tobe := []string{
		"",
	}

	checkSlice(t, actual, tobe)
}

func TestGetAll_Single_HasName(t *testing.T) {
	tree := BodyTree{
		Name: "A",
	}

	actual := tree.GetAll()

	tobe := []string{
		"A",
	}

	checkSlice(t, actual, tobe)
}

func TestGetAll_OneChild_NoName(t *testing.T) {
	tree := BodyTree{
		Name: "",
		Childs: []BodyTree{
			BodyTree{
				Name: "1",
			},
		},
	}

	actual := tree.GetAll()

	tobe := []string{
		"",
		"1",
	}

	checkSlice(t, actual, tobe)
}

func TestGetAll_OneChild_HasName(t *testing.T) {
	tree := BodyTree{
		Name: "A",
		Childs: []BodyTree{
			BodyTree{
				Name: "1",
			},
		},
	}

	actual := tree.GetAll()

	tobe := []string{
		"A",
		"A 1",
	}

	checkSlice(t, actual, tobe)
}

func TestGetAll_LongTree(t *testing.T) {
	tree := BodyTree{
		Name: "A",
		Childs: []BodyTree{
			BodyTree{
				Name: "1",
				Childs: []BodyTree{
					BodyTree{
						Name: "a",
						Childs: []BodyTree{
							BodyTree{
								Name: "a",
							},
						},
					},
				},
			},
		},
	}

	actual := tree.GetAll()

	tobe := []string{
		"A",
		"A 1",
		"A 1 a",
		"A 1 a a",
	}

	checkSlice(t, actual, tobe)
}

//// Index ////

func TestIndex(t *testing.T) {
	tt := []struct {
		Name   string
		Prefix string
		Index  int
	}{
		{"", "", 0},
		{"1", "", 1},
		{"10", "", 10},
		{"A", "", 1},
		{"a", "", 1},
		{"AB 1", "AB ", 1},
	}

	for j, c := range tt {
		prefix, index := BodyTree{Name: c.Name}.Index()

		if prefix != c.Prefix {
			t.Errorf("[%d] Mismatch Prefix: actual %q, tobe %q", j, prefix, c.Prefix)
		}
		if index != c.Index {
			t.Errorf("[%d] Mismatch index: actual %d, tobe %d", j, index, c.Index)
		}
	}
}

//// Mismatch ////
