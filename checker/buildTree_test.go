package checker

import "testing"

func TestBuildTree_Single(t *testing.T) {
	tobe := []string{
		"Foo A",
	}

	trees := BuildTree("Foo", tobe)
	actual := make([]string, 0)

	for _, tree := range trees {
		names := tree.GetAll("Foo")
		actual = append(actual, names...)
	}

	checkSlice(actual, tobe, t)
}

func TestBuildTree_Simple(t *testing.T) {
	tobe := []string{
		"Foo A",
		"Foo B",
	}

	trees := BuildTree("Foo", tobe)
	actual := make([]string, 0)

	for _, tree := range trees {
		names := tree.GetAll("Foo")
		actual = append(actual, names...)
	}

	checkSlice(actual, tobe, t)
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

	actual := trees[0].GetAll("Foo")
	checkSlice(actual, tobe, t)

	// Level 2
	trees = trees[0].Childs
	tobe = tobe[1:]

	if len(trees) != 1 {
		t.Fatalf("Unexpected trees size: %d", len(trees))
	}

	actual = trees[0].GetAll("Foo A")
	checkSlice(actual, tobe, t)

	// Level 3
	trees = trees[0].Childs
	tobe = tobe[1:]

	if len(trees) != 1 {
		t.Fatalf("Unexpected trees size: %d", len(trees))
	}

	actual = trees[0].GetAll("Foo A 1")
	checkSlice(actual, tobe, t)

	// Level 4
	trees = trees[0].Childs
	if len(trees) > 0 {
		t.Fatalf("Unexpected Childs exists: %+v", trees)
	}
}
