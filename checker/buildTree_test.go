package checker

import "testing"

func TestBuildTree_Single(t *testing.T) {
	tobe := []string{
		"Foo A",
	}

	trees := BuildTree("Foo", tobe)
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
		"Foo A",
		"Foo B",
	}

	trees := BuildTree("Foo", tobe)
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
		"Foo A",
		"Foo A 1",
		"Foo A 1 a",
	}

	trees := BuildTree("Foo", tobe)

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
