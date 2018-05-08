package checker

import "testing"

func TestBuildTree_NoName(t *testing.T) {
	tobe := []string{
		"",
	}

	trees := BuildTree(tobe)
	actual := make([]string, 0)

	for _, tree := range trees {
		names := tree.GetAll()
		actual = append(actual, names...)
	}
	checkSlice(
		t,
		actual,
		[]string{
			"",
		},
	)
}

func TestBuildTree_NoNameWithPlanet(t *testing.T) {
	tobe := []string{
		"",
		"1",
		"2",
	}

	trees := BuildTree(tobe)

	if len(trees) != 1 {
		t.Fatalf("Mismatch trees length: actual %d, tobe 1", len(trees))
	}

	tree := trees[0]

	if tree.Name != "" {
		t.Errorf("Mismatch tree name: actual %q, tobe %q", tree.Name, "")
	}

	checkSlice(
		t,
		tree.GetAll(),
		[]string{
			"",
			"1",
			"2",
		},
	)
}

func TestBuildTree_Single(t *testing.T) {
	tobe := []string{
		"A",
	}

	trees := BuildTree(tobe)
	actual := make([]string, 0)

	for _, tree := range trees {
		names := tree.GetAll()
		actual = append(actual, names...)
	}
	checkSlice(
		t,
		actual,
		[]string{
			"A",
		},
	)
}

func TestBuildTree_Simple(t *testing.T) {
	tobe := []string{
		"A",
		"B",
	}

	trees := BuildTree(tobe)
	actual := make([]string, 0)

	for _, tree := range trees {
		names := tree.GetAll()
		actual = append(actual, names...)
	}

	checkSlice(
		t,
		actual,
		[]string{
			"A",
			"B",
		},
	)
}

func TestBuildTree_SimpleTree(t *testing.T) {
	tobe := []string{
		"A",
		"A 1",
		"A 1 a",
	}

	trees := BuildTree(tobe)

	// Level 1
	if len(trees) != 1 {
		t.Fatalf("Unexpected trees size: %d", len(trees))
	}

	actual := trees[0].GetAll()
	checkSlice(
		t,
		actual,
		[]string{
			"A",
			"A 1",
			"A 1 a",
		},
	)

	// Level 2
	trees = trees[0].Childs
	tobe = tobe[1:]

	if len(trees) != 1 {
		t.Fatalf("Unexpected trees size: %d", len(trees))
	}

	actual = trees[0].GetAll()
	checkSlice(
		t,
		actual,
		[]string{
			"1",
			"1 a",
		},
	)

	// Level 3
	trees = trees[0].Childs
	tobe = tobe[1:]

	if len(trees) != 1 {
		t.Fatalf("Unexpected trees size: %d", len(trees))
	}

	actual = trees[0].GetAll()
	checkSlice(
		t,
		actual,
		[]string{
			"a",
		},
	)

	// Level 4
	trees = trees[0].Childs
	if len(trees) > 0 {
		t.Fatalf("Unexpected Childs exists: %+v", trees)
	}
}

func TestBuildTree_Twin(t *testing.T) {
	tobe := []string{
		"A",
		"B",
		"AB 1",
	}

	trees := BuildTree(tobe)
	actual := make([]string, 0)

	for _, tree := range trees {
		names := tree.GetAll()
		actual = append(actual, names...)
	}

	checkSlice(
		t,
		actual,
		[]string{
			"A",
			"B",
			"AB 1",
		},
	)
}

func TestBuildTree_Binary(t *testing.T) {
	src := []string{
		"A",
		"A 1",
		"AB 1",
		"AB 2",
		"B",
		"B 1",
	}

	trees := BuildTree(src)

	if len(trees) != 4 {
		t.Fatalf("Mismatch trees length: actual %d, tobe 4\n    %+v", len(trees), trees)
	}
}
