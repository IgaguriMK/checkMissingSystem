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

func TestMissing_Simple_CompleteSatellite(t *testing.T) {
	bodies := []string{
		"",
		"1",
		"2",
		"3",
		"3 a",
		"3 b",
		"3 c",
	}

	trees := BuildTree(bodies)
	tree := trees[0]
	actual := tree.Missing()

	if actual != false {
		t.Errorf("Should not detect missing: %+v\n    %+v", tree, trees)
	}
}

func TestMissing_Simple_MissingSatellite(t *testing.T) {
	bodies := []string{
		"",
		"1",
		"2",
		"3",
		"3 a",
		"3 b",
		"3 c",
	}

	trees := BuildTree(bodies)
	tree := trees[0]
	actual := tree.Missing()

	if actual != false {
		t.Errorf("Should detect missing: %+v\n    %+v", tree, trees)
	}
}

func TestMissing_Long(t *testing.T) {
	bodies := []string{
		"",
		"1",
		"10",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
	}

	trees := BuildTree(bodies)
	tree := trees[0]
	actual := tree.Missing()

	if actual != false {
		t.Errorf("Should not detect missing: %+v\n    %+v", tree, trees)
	}
}

//// CheckMissing ///

func TestCheckMissing_Complete(t *testing.T) {
	src := []string{
		"A",
		"A 1",
		"A 2",
		"A 3",
		"A 4",
		"B",
		"B 1",
		"B 1 a",
		"B 1 b",
		"B 1 c",
		"B 1 c a",
		"B 2",
		"B 3",
		"AB 1",
		"AB 1 a",
		"AB 2",
		"AB 2 a",
		"AB 2 b",
		"AB 2 c",
		"AB 2 d",
		"AB 2 d a",
		"AB 3",
	}

	trees := BuildTree(src)
	actual := CheckMissing(trees)

	if actual != false {
		t.Errorf("Should not detect missing: %+v", trees)
	}
}

func TestCheckMissing_MissingStar(t *testing.T) {
	src := []string{
		"A",
		"A 1",
		"A 2",
		"A 3",
		"A 4",
		"B 1",
		"B 1 a",
		"B 1 b",
		"B 1 c",
		"B 1 c a",
		"B 2",
		"B 3",
		"AB 1",
		"AB 1 a",
		"AB 2",
		"AB 2 a",
		"AB 2 b",
		"AB 2 c",
		"AB 2 d",
		"AB 2 d a",
		"AB 3",
	}

	trees := BuildTree(src)
	actual := CheckMissing(trees)

	if actual != true {
		t.Errorf("Should detect missing: %+v", trees)
	}
}
