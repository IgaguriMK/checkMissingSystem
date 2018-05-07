package checker

import "testing"

func TestSuffixes_Ok(t *testing.T) {
	src := []string{
		"Foo",
		"Foo 1",
		"Foo 2",
		"Foo 2 a",
	}

	actual, ok := Suffixes("Foo", src)
	if !ok {
		t.Fatal("Should return true")
	}

	checkSlice(
		t,
		actual,
		[]string{
			"",
			"1",
			"2",
			"2 a",
		},
	)
}

func TestSuffixes_NG(t *testing.T) {
	src := []string{
		"Foo",
		"Hoge",
		"Foo 2",
		"Foo 2 a",
	}

	_, ok := Suffixes("Foo", src)
	if ok {
		t.Fatal("Should return false")
	}
}
