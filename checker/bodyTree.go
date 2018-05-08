package checker

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

const (
	DummyName = "SYSTEM"
)

type BodyTree struct {
	Name   string
	Childs []BodyTree
}

func (bt BodyTree) GetAll() []string {
	res := make([]string, 0)
	return bt.getAllInternal(res, "")

}

func GetAllTrees(bts []BodyTree) []string {
	res := make([]string, 0)

	for _, bt := range bts {
		res = append(res, bt.GetAll()...)
	}

	return res
}

func (bt BodyTree) getAllInternal(res []string, prefix string) []string {
	name := bt.Name

	if prefix != "" {
		name = prefix + " " + name
	}

	res = append(res, name)

	for _, ch := range bt.Childs {
		res = ch.getAllInternal(res, name)
	}

	return res
}

func (bt BodyTree) Missing() bool {
	indexMap := make(map[string][]int)

	for _, c := range bt.Childs {
		pf, i := c.Index()

		if i < 0 {
			log.Printf("Index parse error: %q\n", bt.Name)
			return false
		}

		if len(pf) == 2 { // like "A "
			return true
		}

		if _, ok := indexMap[pf]; !ok {
			indexMap[pf] = make([]int, 0)
		}
	}

	for _, c := range bt.Childs {
		pf, ci := c.Index()

		indexMap[pf] = append(indexMap[pf], ci)
	}

	for _, cs := range indexMap {
		sort.Ints(cs)

		offset := 1

		for j, i := range cs {
			if i == 0 {
				offset = 0
				continue
			}

			if i != j+offset {
				return true
			}
		}
	}

	for _, c := range bt.Childs {
		if c.Missing() {
			return true
		}
	}

	return false
}

func CheckMissing(bts []BodyTree) bool {
	tempBt := BodyTree{
		Name:   DummyName,
		Childs: bts,
	}

	return tempBt.Missing()
}

func (bt BodyTree) Index() (string, int) {
	n := bt.Name

	if n == DummyName {
		return "", 0
	}

	if n == "" {
		return "", 1
	}

	ns := strings.Split(n, " ")
	indexStr := ns[len(ns)-1]

	prefix := ""
	if len(ns) > 1 {
		prefix = strings.Join(ns[:len(ns)-1], " ") + " "
	}

	if i, err := strconv.Atoi(indexStr); err == nil {
		if i <= 0 {
			panic("Negative indexStr")
		}

		return prefix, i
	}

	r := indexStr[0]
	switch {
	case 'A' <= r && r <= 'Z':
		return prefix, 1 + int(r-'A')
	case 'a' <= r && r <= 'z':
		return prefix, 1 + int(r-'a')
	}

	return n, -1
}

func (bt BodyTree) GetTier() Tier {
	if bt.Name == "" {
		return SingleStar
	}

	ns := strings.Split(bt.Name, " ")
	indexStr := ns[len(ns)-1]

	if i, err := strconv.Atoi(indexStr); err == nil {
		if i <= 0 {
			panic("Negative indexStr")
		}

		return Planet
	}

	r := indexStr[0]
	switch {
	case 'A' <= r && r <= 'Z':
		return BinaryStar
	case 'a' <= r && r <= 'z':
		return Satellite
	}

	panic(fmt.Sprintf("Invalid index name: %q in %q", indexStr, bt.Name))
}

type Tier int

const (
	SingleStar Tier = iota
	BinaryStar
	Planet
	Satellite
)

func (t Tier) String() string {
	switch t {
	case SingleStar:
		return "SingleStar"
	case BinaryStar:
		return "BinaryStar"
	case Planet:
		return "Planet"
	case Satellite:
		return "Satellite"
	}

	panic(fmt.Sprintf("Illegal Tier value %d", int(t)))
}

func (t Tier) IndexName(index int) string {
	if t == SingleStar {
		return ""
	}

	switch t {
	case BinaryStar:
		r := 'A' + rune(index) - 1
		return string([]rune{r})
	case Planet:
		return strconv.Itoa(index)
	case Satellite:
		r := 'a' + rune(index) - 1
		return string([]rune{r})
	}

	panic("Should not reach")
}

func (bt BodyTree) String() string {
	if len(bt.Childs) == 0 {
		return fmt.Sprintf("%q{}", bt.Name)
	}

	cs := make([]string, 0, len(bt.Childs))
	for _, c := range bt.Childs {
		s := c.String()
		cs = append(cs, s)
	}

	return fmt.Sprintf(
		"%q{ %s }",
		bt.Name,
		strings.Join(cs, ", "),
	)
}
