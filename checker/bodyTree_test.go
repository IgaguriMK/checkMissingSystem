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

//// GetTier ////

func TestGetTier(t *testing.T) {
	tt := []struct {
		Name string
		Tier Tier
	}{
		{"", SingleStar},
		{"A", BinaryStar},
		{"B", BinaryStar},
		{"1", Planet},
		{"13", Planet},
		{"a", Satellite},
	}

	for j, c := range tt {
		tier := BodyTree{Name: c.Name}.GetTier()

		if tier != c.Tier {
			t.Errorf("[%d] Mismatch tier: actual %v, tobe %v", j, tier, c.Tier)
		}
	}
}

//// Tier ////

func TestIndexName(t *testing.T) {
	tt := []struct {
		Index int
		Tier  Tier
		ToBe  string
	}{
		{0, SingleStar, ""},
		{1, BinaryStar, "A"},
		{4, BinaryStar, "D"},
		{1, Planet, "1"},
		{1, Satellite, "a"},
	}

	for j, c := range tt {
		actual := c.Tier.IndexName(c.Index)

		if actual != c.ToBe {
			t.Errorf("[%d] Mismatch IndexName: actual %q, tobe %q", j, actual, c.ToBe)
		}
	}
}

//// Missing ////

func TestMissing_None(t *testing.T) {
	bodies := []string{
		"",
	}

	tree := BuildTree(bodies)[0]
	actual := tree.Missing()

	if actual != false {
		t.Errorf("Should not detect missing: %+v", tree)
	}
}

func TestMissing_Simple_Zero(t *testing.T) {
	bodies := []string{
		"",
		"1",
		"2",
	}

	tree := BuildTree(bodies)[0]
	actual := tree.Missing()

	if actual != false {
		t.Errorf("Should not detect missing: %+v", tree)
	}
}

func TestMissing_Simple_Missing(t *testing.T) {
	bodies := []string{
		"",
		"1",
		"3",
	}

	trees := BuildTree(bodies)
	tree := trees[0]
	actual := tree.Missing()

	if actual != true {
		t.Errorf("Should detect missing: %+v\n    %+v", tree, trees)
	}
}
