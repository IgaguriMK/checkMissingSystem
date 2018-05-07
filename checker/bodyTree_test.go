package checker

import (
	"testing"
)

func TestGetAll_Single_NoName(t *testing.T) {
	tree := BodyTree{
		Name: "",
	}

	actual := tree.GetAll("Foo")

	tobe := []string{
		"Foo",
	}

	checkSlice(actual, tobe, t)
}

func TestGetAll_Single_HasName(t *testing.T) {
	tree := BodyTree{
		Name: "A",
	}

	actual := tree.GetAll("Foo")

	tobe := []string{
		"Foo A",
	}

	checkSlice(actual, tobe, t)
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

	actual := tree.GetAll("Foo")

	tobe := []string{
		"Foo",
		"Foo 1",
	}

	checkSlice(actual, tobe, t)
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

	actual := tree.GetAll("Foo")

	tobe := []string{
		"Foo A",
		"Foo A 1",
	}

	checkSlice(actual, tobe, t)
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

	actual := tree.GetAll("Foo")

	tobe := []string{
		"Foo A",
		"Foo A 1",
		"Foo A 1 a",
		"Foo A 1 a a",
	}

	checkSlice(actual, tobe, t)
}
